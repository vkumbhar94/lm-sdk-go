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

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// NewUpdateAPITokenByAdminIDParams creates a new UpdateAPITokenByAdminIDParams object
// with the default values initialized.
func NewUpdateAPITokenByAdminIDParams() *UpdateAPITokenByAdminIDParams {
	var ()
	return &UpdateAPITokenByAdminIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAPITokenByAdminIDParamsWithTimeout creates a new UpdateAPITokenByAdminIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAPITokenByAdminIDParamsWithTimeout(timeout time.Duration) *UpdateAPITokenByAdminIDParams {
	var ()
	return &UpdateAPITokenByAdminIDParams{

		timeout: timeout,
	}
}

// NewUpdateAPITokenByAdminIDParamsWithContext creates a new UpdateAPITokenByAdminIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAPITokenByAdminIDParamsWithContext(ctx context.Context) *UpdateAPITokenByAdminIDParams {
	var ()
	return &UpdateAPITokenByAdminIDParams{

		Context: ctx,
	}
}

// NewUpdateAPITokenByAdminIDParamsWithHTTPClient creates a new UpdateAPITokenByAdminIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAPITokenByAdminIDParamsWithHTTPClient(client *http.Client) *UpdateAPITokenByAdminIDParams {
	var ()
	return &UpdateAPITokenByAdminIDParams{
		HTTPClient: client,
	}
}

/*UpdateAPITokenByAdminIDParams contains all the parameters to send to the API endpoint
for the update Api token by admin Id operation typically these are written to a http.Request
*/
type UpdateAPITokenByAdminIDParams struct {

	/*AdminID*/
	AdminID int32
	/*ApitokenID*/
	ApitokenID int32
	/*Body*/
	Body *models.APIToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) WithTimeout(timeout time.Duration) *UpdateAPITokenByAdminIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) WithContext(ctx context.Context) *UpdateAPITokenByAdminIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) WithHTTPClient(client *http.Client) *UpdateAPITokenByAdminIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAdminID adds the adminID to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) WithAdminID(adminID int32) *UpdateAPITokenByAdminIDParams {
	o.SetAdminID(adminID)
	return o
}

// SetAdminID adds the adminId to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) SetAdminID(adminID int32) {
	o.AdminID = adminID
}

// WithApitokenID adds the apitokenID to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) WithApitokenID(apitokenID int32) *UpdateAPITokenByAdminIDParams {
	o.SetApitokenID(apitokenID)
	return o
}

// SetApitokenID adds the apitokenId to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) SetApitokenID(apitokenID int32) {
	o.ApitokenID = apitokenID
}

// WithBody adds the body to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) WithBody(body *models.APIToken) *UpdateAPITokenByAdminIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update Api token by admin Id params
func (o *UpdateAPITokenByAdminIDParams) SetBody(body *models.APIToken) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAPITokenByAdminIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param adminId
	if err := r.SetPathParam("adminId", swag.FormatInt32(o.AdminID)); err != nil {
		return err
	}

	// path param apitokenId
	if err := r.SetPathParam("apitokenId", swag.FormatInt32(o.ApitokenID)); err != nil {
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