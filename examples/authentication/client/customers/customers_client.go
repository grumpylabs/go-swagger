// Code generated by go-swagger; DO NOT EDIT.

package customers

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new customers API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for customers API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
Create creates a new customer to track
*/
func (a *Client) Create(params *CreateParams, authInfo runtime.ClientAuthInfoWriter) (*CreateCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "create",
		Method:             "POST",
		PathPattern:        "/customers",
		ProducesMediaTypes: []string{"application/keyauth.api.v1+json"},
		ConsumesMediaTypes: []string{"application/keyauth.api.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*CreateDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
GetID gets a customer Id given an s s n
*/
func (a *Client) GetID(params *GetIDParams, authInfo runtime.ClientAuthInfoWriter) (*GetIDOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetIDParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getId",
		Method:             "GET",
		PathPattern:        "/customers",
		ProducesMediaTypes: []string{"application/keyauth.api.v1+json"},
		ConsumesMediaTypes: []string{"application/keyauth.api.v1+json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetIDReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetIDOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*GetIDDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
