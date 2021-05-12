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

// NewPatchWidgetByIDParams creates a new PatchWidgetByIDParams object
// with the default values initialized.
func NewPatchWidgetByIDParams() *PatchWidgetByIDParams {
	var ()
	return &PatchWidgetByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchWidgetByIDParamsWithTimeout creates a new PatchWidgetByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchWidgetByIDParamsWithTimeout(timeout time.Duration) *PatchWidgetByIDParams {
	var ()
	return &PatchWidgetByIDParams{

		timeout: timeout,
	}
}

// NewPatchWidgetByIDParamsWithContext creates a new PatchWidgetByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchWidgetByIDParamsWithContext(ctx context.Context) *PatchWidgetByIDParams {
	var ()
	return &PatchWidgetByIDParams{

		Context: ctx,
	}
}

// NewPatchWidgetByIDParamsWithHTTPClient creates a new PatchWidgetByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchWidgetByIDParamsWithHTTPClient(client *http.Client) *PatchWidgetByIDParams {
	var ()
	return &PatchWidgetByIDParams{
		HTTPClient: client,
	}
}

/*PatchWidgetByIDParams contains all the parameters to send to the API endpoint
for the patch widget by Id operation typically these are written to a http.Request
*/
type PatchWidgetByIDParams struct {

	/*Body*/
	Body models.Widget
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithTimeout(timeout time.Duration) *PatchWidgetByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithContext(ctx context.Context) *PatchWidgetByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithHTTPClient(client *http.Client) *PatchWidgetByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithBody(body models.Widget) *PatchWidgetByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetBody(body models.Widget) {
	o.Body = body
}

// WithID adds the id to the patch widget by Id params
func (o *PatchWidgetByIDParams) WithID(id int32) *PatchWidgetByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch widget by Id params
func (o *PatchWidgetByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchWidgetByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
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
