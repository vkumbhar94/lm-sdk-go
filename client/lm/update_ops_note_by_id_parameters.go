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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewUpdateOpsNoteByIDParams creates a new UpdateOpsNoteByIDParams object
// with the default values initialized.
func NewUpdateOpsNoteByIDParams() *UpdateOpsNoteByIDParams {
	var ()
	return &UpdateOpsNoteByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateOpsNoteByIDParamsWithTimeout creates a new UpdateOpsNoteByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateOpsNoteByIDParamsWithTimeout(timeout time.Duration) *UpdateOpsNoteByIDParams {
	var ()
	return &UpdateOpsNoteByIDParams{

		timeout: timeout,
	}
}

// NewUpdateOpsNoteByIDParamsWithContext creates a new UpdateOpsNoteByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateOpsNoteByIDParamsWithContext(ctx context.Context) *UpdateOpsNoteByIDParams {
	var ()
	return &UpdateOpsNoteByIDParams{

		Context: ctx,
	}
}

// NewUpdateOpsNoteByIDParamsWithHTTPClient creates a new UpdateOpsNoteByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateOpsNoteByIDParamsWithHTTPClient(client *http.Client) *UpdateOpsNoteByIDParams {
	var ()
	return &UpdateOpsNoteByIDParams{
		HTTPClient: client,
	}
}

/*UpdateOpsNoteByIDParams contains all the parameters to send to the API endpoint
for the update ops note by Id operation typically these are written to a http.Request
*/
type UpdateOpsNoteByIDParams struct {

	/*Body*/
	Body *models.OpsNote
	/*ID*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) WithTimeout(timeout time.Duration) *UpdateOpsNoteByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) WithContext(ctx context.Context) *UpdateOpsNoteByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) WithHTTPClient(client *http.Client) *UpdateOpsNoteByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) WithBody(body *models.OpsNote) *UpdateOpsNoteByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) SetBody(body *models.OpsNote) {
	o.Body = body
}

// WithID adds the id to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) WithID(id string) *UpdateOpsNoteByIDParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update ops note by Id params
func (o *UpdateOpsNoteByIDParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateOpsNoteByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
