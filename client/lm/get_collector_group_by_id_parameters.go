// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/vkumbhar94/runtime"
	cr "github.com/vkumbhar94/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetCollectorGroupByIDParams creates a new GetCollectorGroupByIDParams object
// with the default values initialized.
func NewGetCollectorGroupByIDParams() *GetCollectorGroupByIDParams {
	var ()
	return &GetCollectorGroupByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCollectorGroupByIDParamsWithTimeout creates a new GetCollectorGroupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCollectorGroupByIDParamsWithTimeout(timeout time.Duration) *GetCollectorGroupByIDParams {
	var ()
	return &GetCollectorGroupByIDParams{

		timeout: timeout,
	}
}

// NewGetCollectorGroupByIDParamsWithContext creates a new GetCollectorGroupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCollectorGroupByIDParamsWithContext(ctx context.Context) *GetCollectorGroupByIDParams {
	var ()
	return &GetCollectorGroupByIDParams{

		Context: ctx,
	}
}

// NewGetCollectorGroupByIDParamsWithHTTPClient creates a new GetCollectorGroupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCollectorGroupByIDParamsWithHTTPClient(client *http.Client) *GetCollectorGroupByIDParams {
	var ()
	return &GetCollectorGroupByIDParams{
		HTTPClient: client,
	}
}

/*GetCollectorGroupByIDParams contains all the parameters to send to the API endpoint
for the get collector group by Id operation typically these are written to a http.Request
*/
type GetCollectorGroupByIDParams struct {

	/*Fields*/
	Fields *string
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) WithTimeout(timeout time.Duration) *GetCollectorGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) WithContext(ctx context.Context) *GetCollectorGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) WithHTTPClient(client *http.Client) *GetCollectorGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) WithFields(fields *string) *GetCollectorGroupByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) WithID(id int32) *GetCollectorGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get collector group by Id params
func (o *GetCollectorGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *GetCollectorGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
