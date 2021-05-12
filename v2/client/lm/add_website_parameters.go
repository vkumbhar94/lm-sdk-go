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

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// NewAddWebsiteParams creates a new AddWebsiteParams object
// with the default values initialized.
func NewAddWebsiteParams() *AddWebsiteParams {
	var ()
	return &AddWebsiteParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddWebsiteParamsWithTimeout creates a new AddWebsiteParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddWebsiteParamsWithTimeout(timeout time.Duration) *AddWebsiteParams {
	var ()
	return &AddWebsiteParams{

		timeout: timeout,
	}
}

// NewAddWebsiteParamsWithContext creates a new AddWebsiteParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddWebsiteParamsWithContext(ctx context.Context) *AddWebsiteParams {
	var ()
	return &AddWebsiteParams{

		Context: ctx,
	}
}

// NewAddWebsiteParamsWithHTTPClient creates a new AddWebsiteParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddWebsiteParamsWithHTTPClient(client *http.Client) *AddWebsiteParams {
	var ()
	return &AddWebsiteParams{
		HTTPClient: client,
	}
}

/*AddWebsiteParams contains all the parameters to send to the API endpoint
for the add website operation typically these are written to a http.Request
*/
type AddWebsiteParams struct {

	/*Body*/
	Body models.Website

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add website params
func (o *AddWebsiteParams) WithTimeout(timeout time.Duration) *AddWebsiteParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add website params
func (o *AddWebsiteParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add website params
func (o *AddWebsiteParams) WithContext(ctx context.Context) *AddWebsiteParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add website params
func (o *AddWebsiteParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add website params
func (o *AddWebsiteParams) WithHTTPClient(client *http.Client) *AddWebsiteParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add website params
func (o *AddWebsiteParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add website params
func (o *AddWebsiteParams) WithBody(body models.Website) *AddWebsiteParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add website params
func (o *AddWebsiteParams) SetBody(body models.Website) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddWebsiteParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
