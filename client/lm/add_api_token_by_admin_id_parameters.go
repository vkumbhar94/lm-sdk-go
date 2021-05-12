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

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewAddAPITokenByAdminIDParams creates a new AddAPITokenByAdminIDParams object
// with the default values initialized.
func NewAddAPITokenByAdminIDParams() *AddAPITokenByAdminIDParams {
	var ()
	return &AddAPITokenByAdminIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAPITokenByAdminIDParamsWithTimeout creates a new AddAPITokenByAdminIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAPITokenByAdminIDParamsWithTimeout(timeout time.Duration) *AddAPITokenByAdminIDParams {
	var ()
	return &AddAPITokenByAdminIDParams{

		timeout: timeout,
	}
}

// NewAddAPITokenByAdminIDParamsWithContext creates a new AddAPITokenByAdminIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAPITokenByAdminIDParamsWithContext(ctx context.Context) *AddAPITokenByAdminIDParams {
	var ()
	return &AddAPITokenByAdminIDParams{

		Context: ctx,
	}
}

// NewAddAPITokenByAdminIDParamsWithHTTPClient creates a new AddAPITokenByAdminIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAPITokenByAdminIDParamsWithHTTPClient(client *http.Client) *AddAPITokenByAdminIDParams {
	var ()
	return &AddAPITokenByAdminIDParams{
		HTTPClient: client,
	}
}

/*AddAPITokenByAdminIDParams contains all the parameters to send to the API endpoint
for the add Api token by admin Id operation typically these are written to a http.Request
*/
type AddAPITokenByAdminIDParams struct {

	/*AdminID*/
	AdminID int32
	/*Body*/
	Body *models.APIToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithTimeout(timeout time.Duration) *AddAPITokenByAdminIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithContext(ctx context.Context) *AddAPITokenByAdminIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithHTTPClient(client *http.Client) *AddAPITokenByAdminIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminID adds the adminID to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithAdminID(adminID int32) *AddAPITokenByAdminIDParams {
	o.SetAdminID(adminID)
	return o
}

// SetAdminID adds the adminId to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetAdminID(adminID int32) {
	o.AdminID = adminID
}

// WithBody adds the body to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) WithBody(body *models.APIToken) *AddAPITokenByAdminIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add Api token by admin Id params
func (o *AddAPITokenByAdminIDParams) SetBody(body *models.APIToken) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAPITokenByAdminIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adminId
	if err := r.SetPathParam("adminId", swag.FormatInt32(o.AdminID)); err != nil {
		return err
	}

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
