// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewChunkedParams creates a new ChunkedParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewChunkedParams() *ChunkedParams {
	return &ChunkedParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewChunkedParamsWithTimeout creates a new ChunkedParams object
// with the ability to set a timeout on a request.
func NewChunkedParamsWithTimeout(timeout time.Duration) *ChunkedParams {
	return &ChunkedParams{
		timeout: timeout,
	}
}

// NewChunkedParamsWithContext creates a new ChunkedParams object
// with the ability to set a context for a request.
func NewChunkedParamsWithContext(ctx context.Context) *ChunkedParams {
	return &ChunkedParams{
		Context: ctx,
	}
}

// NewChunkedParamsWithHTTPClient creates a new ChunkedParams object
// with the ability to set a custom HTTPClient for a request.
func NewChunkedParamsWithHTTPClient(client *http.Client) *ChunkedParams {
	return &ChunkedParams{
		HTTPClient: client,
	}
}

/* ChunkedParams contains all the parameters to send to the API endpoint
   for the chunked operation.

   Typically these are written to a http.Request.
*/
type ChunkedParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the chunked params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChunkedParams) WithDefaults() *ChunkedParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the chunked params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ChunkedParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the chunked params
func (o *ChunkedParams) WithTimeout(timeout time.Duration) *ChunkedParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the chunked params
func (o *ChunkedParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the chunked params
func (o *ChunkedParams) WithContext(ctx context.Context) *ChunkedParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the chunked params
func (o *ChunkedParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the chunked params
func (o *ChunkedParams) WithHTTPClient(client *http.Client) *ChunkedParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the chunked params
func (o *ChunkedParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *ChunkedParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
