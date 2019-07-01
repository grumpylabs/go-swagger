FROM golang:alpine as build

RUN apk --no-cache add ca-certificates shared-mime-info mailcap git build-base &&\
  go get -u github.com/go-openapi/runtime &&\
  go get -u github.com/asaskevich/govalidator &&\
  go get -u golang.org/x/net/context &&\
  go get -u github.com/jessevdk/go-flags &&\
  go get -u golang.org/x/net/context/ctxhttp

RUN mkdir -p /go/src/github.com/go-swagger/go-swagger
ADD . /go/src/github.com/go-swagger/go-swagger
RUN cd /go/src/github.com/go-swagger/go-swagger/cmd/swagger && go build
#
FROM alpine
RUN adduser -S -D -H -h / go-swagger
USER go-swagger
COPY --from=build /go/src/github.com/go-swagger/go-swagger/cmd/swagger/swagger /
WORKDIR /
ENTRYPOINT ["/swagger"]