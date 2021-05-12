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

// NewGetDeviceConfigSourceConfigByIDParams creates a new GetDeviceConfigSourceConfigByIDParams object
// with the default values initialized.
func NewGetDeviceConfigSourceConfigByIDParams() *GetDeviceConfigSourceConfigByIDParams {
	var (
		formatDefault     = string("json")
		startEpochDefault = int64(0)
	)
	return &GetDeviceConfigSourceConfigByIDParams{
		Format:     &formatDefault,
		StartEpoch: &startEpochDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceConfigSourceConfigByIDParamsWithTimeout creates a new GetDeviceConfigSourceConfigByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceConfigSourceConfigByIDParamsWithTimeout(timeout time.Duration) *GetDeviceConfigSourceConfigByIDParams {
	var (
		formatDefault     = string("json")
		startEpochDefault = int64(0)
	)
	return &GetDeviceConfigSourceConfigByIDParams{
		Format:     &formatDefault,
		StartEpoch: &startEpochDefault,

		timeout: timeout,
	}
}

// NewGetDeviceConfigSourceConfigByIDParamsWithContext creates a new GetDeviceConfigSourceConfigByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceConfigSourceConfigByIDParamsWithContext(ctx context.Context) *GetDeviceConfigSourceConfigByIDParams {
	var (
		formatDefault     = string("json")
		startEpochDefault = int64(0)
	)
	return &GetDeviceConfigSourceConfigByIDParams{
		Format:     &formatDefault,
		StartEpoch: &startEpochDefault,

		Context: ctx,
	}
}

// NewGetDeviceConfigSourceConfigByIDParamsWithHTTPClient creates a new GetDeviceConfigSourceConfigByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceConfigSourceConfigByIDParamsWithHTTPClient(client *http.Client) *GetDeviceConfigSourceConfigByIDParams {
	var (
		formatDefault     = string("json")
		startEpochDefault = int64(0)
	)
	return &GetDeviceConfigSourceConfigByIDParams{
		Format:     &formatDefault,
		StartEpoch: &startEpochDefault,
		HTTPClient: client,
	}
}

/*GetDeviceConfigSourceConfigByIDParams contains all the parameters to send to the API endpoint
for the get device config source config by Id operation typically these are written to a http.Request
*/
type GetDeviceConfigSourceConfigByIDParams struct {

	/*DeviceID*/
	DeviceID int32
	/*Fields*/
	Fields *string
	/*Format*/
	Format *string
	/*HdsID*/
	HdsID int32
	/*ID*/
	ID string
	/*InstanceID*/
	InstanceID int32
	/*StartEpoch*/
	StartEpoch *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithTimeout(timeout time.Duration) *GetDeviceConfigSourceConfigByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithContext(ctx context.Context) *GetDeviceConfigSourceConfigByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithHTTPClient(client *http.Client) *GetDeviceConfigSourceConfigByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithDeviceID(deviceID int32) *GetDeviceConfigSourceConfigByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithFields(fields *string) *GetDeviceConfigSourceConfigByIDParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFormat adds the format to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithFormat(format *string) *GetDeviceConfigSourceConfigByIDParams {
	o.SetFormat(format)
	return o
}

// SetFormat adds the format to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetFormat(format *string) {
	o.Format = format
}

// WithHdsID adds the hdsID to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithHdsID(hdsID int32) *GetDeviceConfigSourceConfigByIDParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithID(id string) *GetDeviceConfigSourceConfigByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetID(id string) {
	o.ID = id
}

// WithInstanceID adds the instanceID to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithInstanceID(instanceID int32) *GetDeviceConfigSourceConfigByIDParams {
	o.SetInstanceID(instanceID)
	return o
}

// SetInstanceID adds the instanceId to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetInstanceID(instanceID int32) {
	o.InstanceID = instanceID
}

// WithStartEpoch adds the startEpoch to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) WithStartEpoch(startEpoch *int64) *GetDeviceConfigSourceConfigByIDParams {
	o.SetStartEpoch(startEpoch)
	return o
}

// SetStartEpoch adds the startEpoch to the get device config source config by Id params
func (o *GetDeviceConfigSourceConfigByIDParams) SetStartEpoch(startEpoch *int64) {
	o.StartEpoch = startEpoch
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceConfigSourceConfigByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

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

	if o.Format != nil {

		// query param format
		var qrFormat string
		if o.Format != nil {
			qrFormat = *o.Format
		}
		qFormat := qrFormat
		if qFormat != "" {
			if err := r.SetQueryParam("format", qFormat); err != nil {
				return err
			}
		}

	}

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	// path param instanceId
	if err := r.SetPathParam("instanceId", swag.FormatInt32(o.InstanceID)); err != nil {
		return err
	}

	if o.StartEpoch != nil {

		// query param startEpoch
		var qrStartEpoch int64
		if o.StartEpoch != nil {
			qrStartEpoch = *o.StartEpoch
		}
		qStartEpoch := swag.FormatInt64(qrStartEpoch)
		if qStartEpoch != "" {
			if err := r.SetQueryParam("startEpoch", qStartEpoch); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
