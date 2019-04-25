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

// NewGetRoleByIDParams creates a new GetRoleByIDParams object
// with the default values initialized.
func NewGetRoleByIDParams() *GetRoleByIDParams {
	var ()
	return &GetRoleByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetRoleByIDParamsWithTimeout creates a new GetRoleByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetRoleByIDParamsWithTimeout(timeout time.Duration) *GetRoleByIDParams {
	var ()
	return &GetRoleByIDParams{

		timeout: timeout,
	}
}

// NewGetRoleByIDParamsWithContext creates a new GetRoleByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetRoleByIDParamsWithContext(ctx context.Context) *GetRoleByIDParams {
	var ()
	return &GetRoleByIDParams{

		Context: ctx,
	}
}

// NewGetRoleByIDParamsWithHTTPClient creates a new GetRoleByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetRoleByIDParamsWithHTTPClient(client *http.Client) *GetRoleByIDParams {
	var ()
	return &GetRoleByIDParams{
		HTTPClient: client,
	}
}

/*GetRoleByIDParams contains all the parameters to send to the API endpoint
for the get role by Id operation typically these are written to a http.Request
*/
type GetRoleByIDParams struct {

	/*Fields*/
	Fields *string
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get role by Id params
func (o *GetRoleByIDParams) WithTimeout(timeout time.Duration) *GetRoleByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get role by Id params
func (o *GetRoleByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get role by Id params
func (o *GetRoleByIDParams) WithContext(ctx context.Context) *GetRoleByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get role by Id params
func (o *GetRoleByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get role by Id params
func (o *GetRoleByIDParams) WithHTTPClient(client *http.Client) *GetRoleByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get role by Id params
func (o *GetRoleByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get role by Id params
func (o *GetRoleByIDParams) WithFields(fields *string) *GetRoleByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get role by Id params
func (o *GetRoleByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get role by Id params
func (o *GetRoleByIDParams) WithID(id int32) *GetRoleByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get role by Id params
func (o *GetRoleByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetRoleByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}