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

	models "github.com/vkumbhar94/lm-sdk-go/v2/models"
)

// NewPatchCollectorByIDParams creates a new PatchCollectorByIDParams object
// with the default values initialized.
func NewPatchCollectorByIDParams() *PatchCollectorByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchCollectorByIDParamsWithTimeout creates a new PatchCollectorByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchCollectorByIDParamsWithTimeout(timeout time.Duration) *PatchCollectorByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,

		timeout: timeout,
	}
}

// NewPatchCollectorByIDParamsWithContext creates a new PatchCollectorByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchCollectorByIDParamsWithContext(ctx context.Context) *PatchCollectorByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,

		Context: ctx,
	}
}

// NewPatchCollectorByIDParamsWithHTTPClient creates a new PatchCollectorByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchCollectorByIDParamsWithHTTPClient(client *http.Client) *PatchCollectorByIDParams {
	var (
		collectorLoadBalancedDefault        = bool(false)
		forceUpdateFailedOverDevicesDefault = bool(false)
	)
	return &PatchCollectorByIDParams{
		CollectorLoadBalanced:        &collectorLoadBalancedDefault,
		ForceUpdateFailedOverDevices: &forceUpdateFailedOverDevicesDefault,
		HTTPClient:                   client,
	}
}

/*PatchCollectorByIDParams contains all the parameters to send to the API endpoint
for the patch collector by Id operation typically these are written to a http.Request
*/
type PatchCollectorByIDParams struct {

	/*Body*/
	Body *models.Collector
	/*CollectorLoadBalanced*/
	CollectorLoadBalanced *bool
	/*ForceUpdateFailedOverDevices*/
	ForceUpdateFailedOverDevices *bool
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithTimeout(timeout time.Duration) *PatchCollectorByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithContext(ctx context.Context) *PatchCollectorByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithHTTPClient(client *http.Client) *PatchCollectorByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithBody(body *models.Collector) *PatchCollectorByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetBody(body *models.Collector) {
	o.Body = body
}

// WithCollectorLoadBalanced adds the collectorLoadBalanced to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithCollectorLoadBalanced(collectorLoadBalanced *bool) *PatchCollectorByIDParams {
	o.SetCollectorLoadBalanced(collectorLoadBalanced)
	return o
}

// SetCollectorLoadBalanced adds the collectorLoadBalanced to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetCollectorLoadBalanced(collectorLoadBalanced *bool) {
	o.CollectorLoadBalanced = collectorLoadBalanced
}

// WithForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) *PatchCollectorByIDParams {
	o.SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices)
	return o
}

// SetForceUpdateFailedOverDevices adds the forceUpdateFailedOverDevices to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetForceUpdateFailedOverDevices(forceUpdateFailedOverDevices *bool) {
	o.ForceUpdateFailedOverDevices = forceUpdateFailedOverDevices
}

// WithID adds the id to the patch collector by Id params
func (o *PatchCollectorByIDParams) WithID(id int32) *PatchCollectorByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch collector by Id params
func (o *PatchCollectorByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *PatchCollectorByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if o.CollectorLoadBalanced != nil {

		// query param collectorLoadBalanced
		var qrCollectorLoadBalanced bool
		if o.CollectorLoadBalanced != nil {
			qrCollectorLoadBalanced = *o.CollectorLoadBalanced
		}
		qCollectorLoadBalanced := swag.FormatBool(qrCollectorLoadBalanced)
		if qCollectorLoadBalanced != "" {
			if err := r.SetQueryParam("collectorLoadBalanced", qCollectorLoadBalanced); err != nil {
				return err
			}
		}

	}

	if o.ForceUpdateFailedOverDevices != nil {

		// query param forceUpdateFailedOverDevices
		var qrForceUpdateFailedOverDevices bool
		if o.ForceUpdateFailedOverDevices != nil {
			qrForceUpdateFailedOverDevices = *o.ForceUpdateFailedOverDevices
		}
		qForceUpdateFailedOverDevices := swag.FormatBool(qrForceUpdateFailedOverDevices)
		if qForceUpdateFailedOverDevices != "" {
			if err := r.SetQueryParam("forceUpdateFailedOverDevices", qForceUpdateFailedOverDevices); err != nil {
				return err
			}
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
