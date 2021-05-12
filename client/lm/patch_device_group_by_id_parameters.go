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

// NewPatchDeviceGroupByIDParams creates a new PatchDeviceGroupByIDParams object
// with the default values initialized.
func NewPatchDeviceGroupByIDParams() *PatchDeviceGroupByIDParams {
	var ()
	return &PatchDeviceGroupByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchDeviceGroupByIDParamsWithTimeout creates a new PatchDeviceGroupByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchDeviceGroupByIDParamsWithTimeout(timeout time.Duration) *PatchDeviceGroupByIDParams {
	var ()
	return &PatchDeviceGroupByIDParams{

		timeout: timeout,
	}
}

// NewPatchDeviceGroupByIDParamsWithContext creates a new PatchDeviceGroupByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchDeviceGroupByIDParamsWithContext(ctx context.Context) *PatchDeviceGroupByIDParams {
	var ()
	return &PatchDeviceGroupByIDParams{

		Context: ctx,
	}
}

// NewPatchDeviceGroupByIDParamsWithHTTPClient creates a new PatchDeviceGroupByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchDeviceGroupByIDParamsWithHTTPClient(client *http.Client) *PatchDeviceGroupByIDParams {
	var ()
	return &PatchDeviceGroupByIDParams{
		HTTPClient: client,
	}
}

/*PatchDeviceGroupByIDParams contains all the parameters to send to the API endpoint
for the patch device group by Id operation typically these are written to a http.Request
*/
type PatchDeviceGroupByIDParams struct {

	/*Body*/
	Body *models.DeviceGroup
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithTimeout(timeout time.Duration) *PatchDeviceGroupByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithContext(ctx context.Context) *PatchDeviceGroupByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithHTTPClient(client *http.Client) *PatchDeviceGroupByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithBody(body *models.DeviceGroup) *PatchDeviceGroupByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetBody(body *models.DeviceGroup) {
	o.Body = body
}

// WithID adds the id to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) WithID(id int32) *PatchDeviceGroupByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch device group by Id params
func (o *PatchDeviceGroupByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchDeviceGroupByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
