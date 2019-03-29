// Code generated by go-swagger; DO NOT EDIT.

package external

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCurrentGenerationParams creates a new GetCurrentGenerationParams object
// with the default values initialized.
func NewGetCurrentGenerationParams() *GetCurrentGenerationParams {

	return &GetCurrentGenerationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCurrentGenerationParamsWithTimeout creates a new GetCurrentGenerationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCurrentGenerationParamsWithTimeout(timeout time.Duration) *GetCurrentGenerationParams {

	return &GetCurrentGenerationParams{

		timeout: timeout,
	}
}

// NewGetCurrentGenerationParamsWithContext creates a new GetCurrentGenerationParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCurrentGenerationParamsWithContext(ctx context.Context) *GetCurrentGenerationParams {

	return &GetCurrentGenerationParams{

		Context: ctx,
	}
}

// NewGetCurrentGenerationParamsWithHTTPClient creates a new GetCurrentGenerationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCurrentGenerationParamsWithHTTPClient(client *http.Client) *GetCurrentGenerationParams {

	return &GetCurrentGenerationParams{
		HTTPClient: client,
	}
}

/*GetCurrentGenerationParams contains all the parameters to send to the API endpoint
for the get current generation operation typically these are written to a http.Request
*/
type GetCurrentGenerationParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get current generation params
func (o *GetCurrentGenerationParams) WithTimeout(timeout time.Duration) *GetCurrentGenerationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get current generation params
func (o *GetCurrentGenerationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get current generation params
func (o *GetCurrentGenerationParams) WithContext(ctx context.Context) *GetCurrentGenerationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get current generation params
func (o *GetCurrentGenerationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get current generation params
func (o *GetCurrentGenerationParams) WithHTTPClient(client *http.Client) *GetCurrentGenerationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get current generation params
func (o *GetCurrentGenerationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCurrentGenerationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
