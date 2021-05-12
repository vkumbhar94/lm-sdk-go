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

// NewPatchNetscanParams creates a new PatchNetscanParams object
// with the default values initialized.
func NewPatchNetscanParams() *PatchNetscanParams {
	var ()
	return &PatchNetscanParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchNetscanParamsWithTimeout creates a new PatchNetscanParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchNetscanParamsWithTimeout(timeout time.Duration) *PatchNetscanParams {
	var ()
	return &PatchNetscanParams{

		timeout: timeout,
	}
}

// NewPatchNetscanParamsWithContext creates a new PatchNetscanParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchNetscanParamsWithContext(ctx context.Context) *PatchNetscanParams {
	var ()
	return &PatchNetscanParams{

		Context: ctx,
	}
}

// NewPatchNetscanParamsWithHTTPClient creates a new PatchNetscanParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchNetscanParamsWithHTTPClient(client *http.Client) *PatchNetscanParams {
	var ()
	return &PatchNetscanParams{
		HTTPClient: client,
	}
}

/*PatchNetscanParams contains all the parameters to send to the API endpoint
for the patch netscan operation typically these are written to a http.Request
*/
type PatchNetscanParams struct {

	/*Body*/
	Body models.Netscan
	/*ID*/
	ID int32
	/*Reason*/
	Reason *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch netscan params
func (o *PatchNetscanParams) WithTimeout(timeout time.Duration) *PatchNetscanParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch netscan params
func (o *PatchNetscanParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch netscan params
func (o *PatchNetscanParams) WithContext(ctx context.Context) *PatchNetscanParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch netscan params
func (o *PatchNetscanParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch netscan params
func (o *PatchNetscanParams) WithHTTPClient(client *http.Client) *PatchNetscanParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch netscan params
func (o *PatchNetscanParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch netscan params
func (o *PatchNetscanParams) WithBody(body models.Netscan) *PatchNetscanParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch netscan params
func (o *PatchNetscanParams) SetBody(body models.Netscan) {
	o.Body = body
}

// WithID adds the id to the patch netscan params
func (o *PatchNetscanParams) WithID(id int32) *PatchNetscanParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the patch netscan params
func (o *PatchNetscanParams) SetID(id int32) {
	o.ID = id
}

// WithReason adds the reason to the patch netscan params
func (o *PatchNetscanParams) WithReason(reason *string) *PatchNetscanParams {
	o.SetReason(reason)
	return o
}

// SetReason adds the reason to the patch netscan params
func (o *PatchNetscanParams) SetReason(reason *string) {
	o.Reason = reason
}

// WriteToRequest writes these params to a swagger request
func (o *PatchNetscanParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
