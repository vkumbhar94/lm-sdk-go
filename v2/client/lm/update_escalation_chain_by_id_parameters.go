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

// NewUpdateEscalationChainByIDParams creates a new UpdateEscalationChainByIDParams object
// with the default values initialized.
func NewUpdateEscalationChainByIDParams() *UpdateEscalationChainByIDParams {
	var ()
	return &UpdateEscalationChainByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateEscalationChainByIDParamsWithTimeout creates a new UpdateEscalationChainByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateEscalationChainByIDParamsWithTimeout(timeout time.Duration) *UpdateEscalationChainByIDParams {
	var ()
	return &UpdateEscalationChainByIDParams{

		timeout: timeout,
	}
}

// NewUpdateEscalationChainByIDParamsWithContext creates a new UpdateEscalationChainByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateEscalationChainByIDParamsWithContext(ctx context.Context) *UpdateEscalationChainByIDParams {
	var ()
	return &UpdateEscalationChainByIDParams{

		Context: ctx,
	}
}

// NewUpdateEscalationChainByIDParamsWithHTTPClient creates a new UpdateEscalationChainByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateEscalationChainByIDParamsWithHTTPClient(client *http.Client) *UpdateEscalationChainByIDParams {
	var ()
	return &UpdateEscalationChainByIDParams{
		HTTPClient: client,
	}
}

/*UpdateEscalationChainByIDParams contains all the parameters to send to the API endpoint
for the update escalation chain by Id operation typically these are written to a http.Request
*/
type UpdateEscalationChainByIDParams struct {

	/*Body*/
	Body *models.EscalatingChain
	/*ID*/
	ID int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) WithTimeout(timeout time.Duration) *UpdateEscalationChainByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) WithContext(ctx context.Context) *UpdateEscalationChainByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) WithHTTPClient(client *http.Client) *UpdateEscalationChainByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) WithBody(body *models.EscalatingChain) *UpdateEscalationChainByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) SetBody(body *models.EscalatingChain) {
	o.Body = body
}

// WithID adds the id to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) WithID(id int32) *UpdateEscalationChainByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update escalation chain by Id params
func (o *UpdateEscalationChainByIDParams) SetID(id int32) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateEscalationChainByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
