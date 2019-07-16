// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetAppliesToFunctionByIDParams creates a new GetAppliesToFunctionByIDParams object
// with the default values initialized.
func NewGetAppliesToFunctionByIDParams() *GetAppliesToFunctionByIDParams {
	var ()
	return &GetAppliesToFunctionByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetAppliesToFunctionByIDParamsWithTimeout creates a new GetAppliesToFunctionByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetAppliesToFunctionByIDParamsWithTimeout(timeout time.Duration) *GetAppliesToFunctionByIDParams {
	var ()
	return &GetAppliesToFunctionByIDParams{

		timeout: timeout,
	}
}

// NewGetAppliesToFunctionByIDParamsWithContext creates a new GetAppliesToFunctionByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetAppliesToFunctionByIDParamsWithContext(ctx context.Context) *GetAppliesToFunctionByIDParams {
	var ()
	return &GetAppliesToFunctionByIDParams{

		Context: ctx,
	}
}

// NewGetAppliesToFunctionByIDParamsWithHTTPClient creates a new GetAppliesToFunctionByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetAppliesToFunctionByIDParamsWithHTTPClient(client *http.Client) *GetAppliesToFunctionByIDParams {
	var ()
	return &GetAppliesToFunctionByIDParams{
		HTTPClient: client,
	}
}

/*GetAppliesToFunctionByIDParams contains all the parameters to send to the API endpoint
for the get applies to function by Id operation typically these are written to a http.Request
*/
type GetAppliesToFunctionByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) WithTimeout(timeout time.Duration) *GetAppliesToFunctionByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) WithContext(ctx context.Context) *GetAppliesToFunctionByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) WithHTTPClient(client *http.Client) *GetAppliesToFunctionByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) WithID(id int32) *GetAppliesToFunctionByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get applies to function by Id params
func (o *GetAppliesToFunctionByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetAppliesToFunctionByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
