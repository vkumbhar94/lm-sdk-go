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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewAckAlertByIDParams creates a new AckAlertByIDParams object
// with the default values initialized.
func NewAckAlertByIDParams() *AckAlertByIDParams {
	var ()
	return &AckAlertByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAckAlertByIDParamsWithTimeout creates a new AckAlertByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAckAlertByIDParamsWithTimeout(timeout time.Duration) *AckAlertByIDParams {
	var ()
	return &AckAlertByIDParams{

		timeout: timeout,
	}
}

// NewAckAlertByIDParamsWithContext creates a new AckAlertByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewAckAlertByIDParamsWithContext(ctx context.Context) *AckAlertByIDParams {
	var ()
	return &AckAlertByIDParams{

		Context: ctx,
	}
}

// NewAckAlertByIDParamsWithHTTPClient creates a new AckAlertByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAckAlertByIDParamsWithHTTPClient(client *http.Client) *AckAlertByIDParams {
	var ()
	return &AckAlertByIDParams{
		HTTPClient: client,
	}
}

/*AckAlertByIDParams contains all the parameters to send to the API endpoint
for the ack alert by Id operation typically these are written to a http.Request
*/
type AckAlertByIDParams struct {

	/*Body*/
	Body *models.AlertAck
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ack alert by Id params
func (o *AckAlertByIDParams) WithTimeout(timeout time.Duration) *AckAlertByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ack alert by Id params
func (o *AckAlertByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ack alert by Id params
func (o *AckAlertByIDParams) WithContext(ctx context.Context) *AckAlertByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ack alert by Id params
func (o *AckAlertByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ack alert by Id params
func (o *AckAlertByIDParams) WithHTTPClient(client *http.Client) *AckAlertByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ack alert by Id params
func (o *AckAlertByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the ack alert by Id params
func (o *AckAlertByIDParams) WithBody(body *models.AlertAck) *AckAlertByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the ack alert by Id params
func (o *AckAlertByIDParams) SetBody(body *models.AlertAck) {
	o.Body = body
}

// WithID adds the id to the ack alert by Id params
func (o *AckAlertByIDParams) WithID(id string) *AckAlertByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ack alert by Id params
func (o *AckAlertByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *AckAlertByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
