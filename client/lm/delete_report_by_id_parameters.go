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

// NewDeleteReportByIDParams creates a new DeleteReportByIDParams object
// with the default values initialized.
func NewDeleteReportByIDParams() *DeleteReportByIDParams {
	var ()
	return &DeleteReportByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteReportByIDParamsWithTimeout creates a new DeleteReportByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteReportByIDParamsWithTimeout(timeout time.Duration) *DeleteReportByIDParams {
	var ()
	return &DeleteReportByIDParams{

		timeout: timeout,
	}
}

// NewDeleteReportByIDParamsWithContext creates a new DeleteReportByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteReportByIDParamsWithContext(ctx context.Context) *DeleteReportByIDParams {
	var ()
	return &DeleteReportByIDParams{

		Context: ctx,
	}
}

// NewDeleteReportByIDParamsWithHTTPClient creates a new DeleteReportByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteReportByIDParamsWithHTTPClient(client *http.Client) *DeleteReportByIDParams {
	var ()
	return &DeleteReportByIDParams{
		HTTPClient: client,
	}
}

/*DeleteReportByIDParams contains all the parameters to send to the API endpoint
for the delete report by Id operation typically these are written to a http.Request
*/
type DeleteReportByIDParams struct {

	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete report by Id params
func (o *DeleteReportByIDParams) WithTimeout(timeout time.Duration) *DeleteReportByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete report by Id params
func (o *DeleteReportByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete report by Id params
func (o *DeleteReportByIDParams) WithContext(ctx context.Context) *DeleteReportByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete report by Id params
func (o *DeleteReportByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete report by Id params
func (o *DeleteReportByIDParams) WithHTTPClient(client *http.Client) *DeleteReportByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete report by Id params
func (o *DeleteReportByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete report by Id params
func (o *DeleteReportByIDParams) WithID(id int32) *DeleteReportByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete report by Id params
func (o *DeleteReportByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteReportByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
