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

// NewGetDeviceDatasourceInstanceGroupListParams creates a new GetDeviceDatasourceInstanceGroupListParams object
// with the default values initialized.
func NewGetDeviceDatasourceInstanceGroupListParams() *GetDeviceDatasourceInstanceGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceInstanceGroupListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupListParamsWithTimeout creates a new GetDeviceDatasourceInstanceGroupListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetDeviceDatasourceInstanceGroupListParamsWithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceInstanceGroupListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetDeviceDatasourceInstanceGroupListParamsWithContext creates a new GetDeviceDatasourceInstanceGroupListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetDeviceDatasourceInstanceGroupListParamsWithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceInstanceGroupListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetDeviceDatasourceInstanceGroupListParamsWithHTTPClient creates a new GetDeviceDatasourceInstanceGroupListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetDeviceDatasourceInstanceGroupListParamsWithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetDeviceDatasourceInstanceGroupListParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetDeviceDatasourceInstanceGroupListParams contains all the parameters to send to the API endpoint
for the get device datasource instance group list operation typically these are written to a http.Request
*/
type GetDeviceDatasourceInstanceGroupListParams struct {

	/*DeviceDsID
	  The device-datasource ID you'd like to add an instance group for

	*/
	DeviceDsID int32
	/*DeviceID*/
	DeviceID int32
	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithTimeout(timeout time.Duration) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithContext(ctx context.Context) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithHTTPClient(client *http.Client) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceDsID adds the deviceDsID to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithDeviceDsID(deviceDsID int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetDeviceDsID(deviceDsID)
	return o
}

// SetDeviceDsID adds the deviceDsId to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetDeviceDsID(deviceDsID int32) {
	o.DeviceDsID = deviceDsID
}

// WithDeviceID adds the deviceID to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithDeviceID(deviceID int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetDeviceID(deviceID int32) {
	o.DeviceID = deviceID
}

// WithFields adds the fields to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithFields(fields *string) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithFilter(filter *string) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithOffset(offset *int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) WithSize(size *int32) *GetDeviceDatasourceInstanceGroupListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get device datasource instance group list params
func (o *GetDeviceDatasourceInstanceGroupListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetDeviceDatasourceInstanceGroupListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceDsId
	if err := r.SetPathParam("deviceDsId", swag.FormatInt32(o.DeviceDsID)); err != nil {
		return err
	}

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

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
