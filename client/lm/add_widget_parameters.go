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

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// NewAddWidgetParams creates a new AddWidgetParams object
// with the default values initialized.
func NewAddWidgetParams() *AddWidgetParams {
	var ()
	return &AddWidgetParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddWidgetParamsWithTimeout creates a new AddWidgetParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddWidgetParamsWithTimeout(timeout time.Duration) *AddWidgetParams {
	var ()
	return &AddWidgetParams{

		timeout: timeout,
	}
}

// NewAddWidgetParamsWithContext creates a new AddWidgetParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddWidgetParamsWithContext(ctx context.Context) *AddWidgetParams {
	var ()
	return &AddWidgetParams{

		Context: ctx,
	}
}

// NewAddWidgetParamsWithHTTPClient creates a new AddWidgetParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddWidgetParamsWithHTTPClient(client *http.Client) *AddWidgetParams {
	var ()
	return &AddWidgetParams{
		HTTPClient: client,
	}
}

/*AddWidgetParams contains all the parameters to send to the API endpoint
for the add widget operation typically these are written to a http.Request
*/
type AddWidgetParams struct {

	/*Body*/
	Body models.Widget

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add widget params
func (o *AddWidgetParams) WithTimeout(timeout time.Duration) *AddWidgetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add widget params
func (o *AddWidgetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add widget params
func (o *AddWidgetParams) WithContext(ctx context.Context) *AddWidgetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add widget params
func (o *AddWidgetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add widget params
func (o *AddWidgetParams) WithHTTPClient(client *http.Client) *AddWidgetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add widget params
func (o *AddWidgetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add widget params
func (o *AddWidgetParams) WithBody(body models.Widget) *AddWidgetParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add widget params
func (o *AddWidgetParams) SetBody(body models.Widget) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddWidgetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
