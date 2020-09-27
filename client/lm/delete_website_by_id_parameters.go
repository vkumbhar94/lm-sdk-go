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

// NewDeleteWebsiteByIDParams creates a new DeleteWebsiteByIDParams object
// with the default values initialized.
func NewDeleteWebsiteByIDParams() *DeleteWebsiteByIDParams {
	var ()
	return &DeleteWebsiteByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteWebsiteByIDParamsWithTimeout creates a new DeleteWebsiteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteWebsiteByIDParamsWithTimeout(timeout time.Duration) *DeleteWebsiteByIDParams {
	var ()
	return &DeleteWebsiteByIDParams{

		timeout: timeout,
	}
}

// NewDeleteWebsiteByIDParamsWithContext creates a new DeleteWebsiteByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteWebsiteByIDParamsWithContext(ctx context.Context) *DeleteWebsiteByIDParams {
	var ()
	return &DeleteWebsiteByIDParams{

		Context: ctx,
	}
}

// NewDeleteWebsiteByIDParamsWithHTTPClient creates a new DeleteWebsiteByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteWebsiteByIDParamsWithHTTPClient(client *http.Client) *DeleteWebsiteByIDParams {
	var ()
	return &DeleteWebsiteByIDParams{
		HTTPClient: client,
	}
}

/*DeleteWebsiteByIDParams contains all the parameters to send to the API endpoint
for the delete website by Id operation typically these are written to a http.Request
*/
type DeleteWebsiteByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete website by Id params
func (o *DeleteWebsiteByIDParams) WithTimeout(timeout time.Duration) *DeleteWebsiteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete website by Id params
func (o *DeleteWebsiteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete website by Id params
func (o *DeleteWebsiteByIDParams) WithContext(ctx context.Context) *DeleteWebsiteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete website by Id params
func (o *DeleteWebsiteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete website by Id params
func (o *DeleteWebsiteByIDParams) WithHTTPClient(client *http.Client) *DeleteWebsiteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete website by Id params
func (o *DeleteWebsiteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete website by Id params
func (o *DeleteWebsiteByIDParams) WithID(id int32) *DeleteWebsiteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete website by Id params
func (o *DeleteWebsiteByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteWebsiteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
