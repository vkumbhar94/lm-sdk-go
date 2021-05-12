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

// NewUpdateDeviceDatasourceInstanceByIDParams creates a new UpdateDeviceDatasourceInstanceByIDParams object
// with the default values initialized.
func NewUpdateDeviceDatasourceInstanceByIDParams() *UpdateDeviceDatasourceInstanceByIDParams {
	var (
		opTypeDefault = string("refresh")
	)
	return &UpdateDeviceDatasourceInstanceByIDParams{
		OpType: &opTypeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateDeviceDatasourceInstanceByIDParamsWithTimeout creates a new UpdateDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateDeviceDatasourceInstanceByIDParamsWithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceByIDParams {
	var (
		opTypeDefault = string("refresh")
	)
	return &UpdateDeviceDatasourceInstanceByIDParams{
		OpType: &opTypeDefault,

		timeout: timeout,
	}
}

// NewUpdateDeviceDatasourceInstanceByIDParamsWithContext creates a new UpdateDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateDeviceDatasourceInstanceByIDParamsWithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceByIDParams {
	var (
		opTypeDefault = string("refresh")
	)
	return &UpdateDeviceDatasourceInstanceByIDParams{
		OpType: &opTypeDefault,

		Context: ctx,
	}
}

// NewUpdateDeviceDatasourceInstanceByIDParamsWithHTTPClient creates a new UpdateDeviceDatasourceInstanceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateDeviceDatasourceInstanceByIDParamsWithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceByIDParams {
	var (
		opTypeDefault = string("refresh")
	)
	return &UpdateDeviceDatasourceInstanceByIDParams{
		OpType:     &opTypeDefault,
		HTTPClient: client,
	}
}

/*UpdateDeviceDatasourceInstanceByIDParams contains all the parameters to send to the API endpoint
for the update device datasource instance by Id operation typically these are written to a http.Request
*/
type UpdateDeviceDatasourceInstanceByIDParams struct {

	/*Body*/
	Body *models.DeviceDataSourceInstance
	/*DeviceID*/
	DeviceID int32
	/*HdsID
	  The device-datasource ID

	*/
	HdsID int32
	/*ID*/
	ID int32
	/*OpType*/
	OpType *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithTimeout(timeout time.Duration) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithContext(ctx context.Context) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithHTTPClient(client *http.Client) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithBody(body *models.DeviceDataSourceInstance) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetBody(body *models.DeviceDataSourceInstance) {
	o.Body = body
}

// WithDeviceID adds the deviceID to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithDeviceID(deviceID int32) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithHdsID adds the hdsID to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithHdsID(hdsID int32) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetHdsID(hdsID)
	return o
}

// SetHdsID adds the hdsId to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetHdsID(hdsID int32) {
	o.HdsID = hdsID
}

// WithID adds the id to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithID(id int32) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetID(id int32) {
	o.ID = id
}

// WithOpType adds the opType to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) WithOpType(opType *string) *UpdateDeviceDatasourceInstanceByIDParams {
	o.SetOpType(opType)
	return o
}

// SetOpType adds the opType to the update device datasource instance by Id params
func (o *UpdateDeviceDatasourceInstanceByIDParams) SetOpType(opType *string) {
	o.OpType = opType
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateDeviceDatasourceInstanceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param deviceId
	if err := r.SetPathParam("deviceId", swag.FormatInt32(o.DeviceID)); err != nil {
		return err
	}

	// path param hdsId
	if err := r.SetPathParam("hdsId", swag.FormatInt32(o.HdsID)); err != nil {
		return err
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt32(o.ID)); err != nil {
		return err
	}

	if o.OpType != nil {

		// query param opType
		var qrOpType string
		if o.OpType != nil {
			qrOpType = *o.OpType
		}
		qOpType := qrOpType
		if qOpType != "" {
			if err := r.SetQueryParam("opType", qOpType); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
