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

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/models"
)

// NewAddAppliesToFunctionParams creates a new AddAppliesToFunctionParams object
// with the default values initialized.
func NewAddAppliesToFunctionParams() *AddAppliesToFunctionParams {
	var ()
	return &AddAppliesToFunctionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddAppliesToFunctionParamsWithTimeout creates a new AddAppliesToFunctionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddAppliesToFunctionParamsWithTimeout(timeout time.Duration) *AddAppliesToFunctionParams {
	var ()
	return &AddAppliesToFunctionParams{

		timeout: timeout,
	}
}

// NewAddAppliesToFunctionParamsWithContext creates a new AddAppliesToFunctionParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddAppliesToFunctionParamsWithContext(ctx context.Context) *AddAppliesToFunctionParams {
	var ()
	return &AddAppliesToFunctionParams{

		Context: ctx,
	}
}

// NewAddAppliesToFunctionParamsWithHTTPClient creates a new AddAppliesToFunctionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddAppliesToFunctionParamsWithHTTPClient(client *http.Client) *AddAppliesToFunctionParams {
	var ()
	return &AddAppliesToFunctionParams{
		HTTPClient: client,
	}
}

/*AddAppliesToFunctionParams contains all the parameters to send to the API endpoint
for the add applies to function operation typically these are written to a http.Request
*/
type AddAppliesToFunctionParams struct {

	/*Body*/
	Body *models.AppliesToFunction

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add applies to function params
func (o *AddAppliesToFunctionParams) WithTimeout(timeout time.Duration) *AddAppliesToFunctionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add applies to function params
func (o *AddAppliesToFunctionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add applies to function params
func (o *AddAppliesToFunctionParams) WithContext(ctx context.Context) *AddAppliesToFunctionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add applies to function params
func (o *AddAppliesToFunctionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add applies to function params
func (o *AddAppliesToFunctionParams) WithHTTPClient(client *http.Client) *AddAppliesToFunctionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add applies to function params
func (o *AddAppliesToFunctionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add applies to function params
func (o *AddAppliesToFunctionParams) WithBody(body *models.AppliesToFunction) *AddAppliesToFunctionParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add applies to function params
func (o *AddAppliesToFunctionParams) SetBody(body *models.AppliesToFunction) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddAppliesToFunctionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
