// Code generated by go-swagger; DO NOT EDIT.

package todos

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// DestroyOneHandlerFunc turns a function with the right signature into a destroy one handler
type DestroyOneHandlerFunc func(DestroyOneParams) middleware.Responder

// Handle executing the request and returning a response
func (fn DestroyOneHandlerFunc) Handle(params DestroyOneParams) middleware.Responder {
	return fn(params)
}

// DestroyOneHandler interface for that can handle valid destroy one params
type DestroyOneHandler interface {
	Handle(DestroyOneParams) middleware.Responder
}

// NewDestroyOne creates a new http.Handler for the destroy one operation
func NewDestroyOne(ctx *middleware.Context, handler DestroyOneHandler) *DestroyOne {
	return &DestroyOne{Context: ctx, Handler: handler}
}

/*DestroyOne swagger:route DELETE /{id} todos destroyOne

DestroyOne destroy one API

*/
type DestroyOne struct {
	Context *middleware.Context
	Handler DestroyOneHandler
}

func (o *DestroyOne) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	var Params = NewDestroyOneParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}
