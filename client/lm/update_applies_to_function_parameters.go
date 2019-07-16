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

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// NewUpdateAppliesToFunctionParams creates a new UpdateAppliesToFunctionParams object
// with the default values initialized.
func NewUpdateAppliesToFunctionParams() *UpdateAppliesToFunctionParams {
	var (
		ignoreReferenceDefault = bool(false)
	)
	return &UpdateAppliesToFunctionParams{
		IgnoreReference: &ignoreReferenceDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAppliesToFunctionParamsWithTimeout creates a new UpdateAppliesToFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAppliesToFunctionParamsWithTimeout(timeout time.Duration) *UpdateAppliesToFunctionParams {
	var (
		ignoreReferenceDefault = bool(false)
	)
	return &UpdateAppliesToFunctionParams{
		IgnoreReference: &ignoreReferenceDefault,

		timeout: timeout,
	}
}

// NewUpdateAppliesToFunctionParamsWithContext creates a new UpdateAppliesToFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAppliesToFunctionParamsWithContext(ctx context.Context) *UpdateAppliesToFunctionParams {
	var (
		ignoreReferenceDefault = bool(false)
	)
	return &UpdateAppliesToFunctionParams{
		IgnoreReference: &ignoreReferenceDefault,

		Context: ctx,
	}
}

// NewUpdateAppliesToFunctionParamsWithHTTPClient creates a new UpdateAppliesToFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAppliesToFunctionParamsWithHTTPClient(client *http.Client) *UpdateAppliesToFunctionParams {
	var (
		ignoreReferenceDefault = bool(false)
	)
	return &UpdateAppliesToFunctionParams{
		IgnoreReference: &ignoreReferenceDefault,
		HTTPClient:      client,
	}
}

/*UpdateAppliesToFunctionParams contains all the parameters to send to the API endpoint
for the update applies to function operation typically these are written to a http.Request
*/
type UpdateAppliesToFunctionParams struct {

	/*Body*/
	Body *models.AppliesToFunction
	/*ID*/
	ID int32
	/*IgnoreReference*/
	IgnoreReference *bool
	/*Reason*/
	Reason *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update applies to function params
func (o *UpdateAppliesToFunctionParams) WithTimeout(timeout time.Duration) *UpdateAppliesToFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update applies to function params
func (o *UpdateAppliesToFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update applies to function params
func (o *UpdateAppliesToFunctionParams) WithContext(ctx context.Context) *UpdateAppliesToFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update applies to function params
func (o *UpdateAppliesToFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update applies to function params
func (o *UpdateAppliesToFunctionParams) WithHTTPClient(client *http.Client) *UpdateAppliesToFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update applies to function params
func (o *UpdateAppliesToFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update applies to function params
func (o *UpdateAppliesToFunctionParams) WithBody(body *models.AppliesToFunction) *UpdateAppliesToFunctionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update applies to function params
func (o *UpdateAppliesToFunctionParams) SetBody(body *models.AppliesToFunction) {
	o.Body = body
}

// WithID adds the id to the update applies to function params
func (o *UpdateAppliesToFunctionParams) WithID(id int32) *UpdateAppliesToFunctionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update applies to function params
func (o *UpdateAppliesToFunctionParams) SetID(id int32) {
	o.ID = id
}

// WithIgnoreReference adds the ignoreReference to the update applies to function params
func (o *UpdateAppliesToFunctionParams) WithIgnoreReference(ignoreReference *bool) *UpdateAppliesToFunctionParams {
	o.SetIgnoreReference(ignoreReference)
	return o
}

// SetIgnoreReference adds the ignoreReference to the update applies to function params
func (o *UpdateAppliesToFunctionParams) SetIgnoreReference(ignoreReference *bool) {
	o.IgnoreReference = ignoreReference
}

// WithReason adds the reason to the update applies to function params
func (o *UpdateAppliesToFunctionParams) WithReason(reason *string) *UpdateAppliesToFunctionParams {
	o.SetReason(reason)
	return o
}

// SetReason adds the reason to the update applies to function params
func (o *UpdateAppliesToFunctionParams) SetReason(reason *string) {
	o.Reason = reason
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAppliesToFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.IgnoreReference != nil {

		// query param ignoreReference
		var qrIgnoreReference bool
		if o.IgnoreReference != nil {
			qrIgnoreReference = *o.IgnoreReference
		}
		qIgnoreReference := swag.FormatBool(qrIgnoreReference)
		if qIgnoreReference != "" {
			if err := r.SetQueryParam("ignoreReference", qIgnoreReference); err != nil {
				return err
			}
		}

	}

	if o.Reason != nil {

		// query param reason
		var qrReason string
		if o.Reason != nil {
			qrReason = *o.Reason
		}
		qReason := qrReason
		if qReason != "" {
			if err := r.SetQueryParam("reason", qReason); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
