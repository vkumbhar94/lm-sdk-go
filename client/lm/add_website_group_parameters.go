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

// NewAddWebsiteGroupParams creates a new AddWebsiteGroupParams object
// with the default values initialized.
func NewAddWebsiteGroupParams() *AddWebsiteGroupParams {
	var ()
	return &AddWebsiteGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddWebsiteGroupParamsWithTimeout creates a new AddWebsiteGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddWebsiteGroupParamsWithTimeout(timeout time.Duration) *AddWebsiteGroupParams {
	var ()
	return &AddWebsiteGroupParams{

		timeout: timeout,
	}
}

// NewAddWebsiteGroupParamsWithContext creates a new AddWebsiteGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddWebsiteGroupParamsWithContext(ctx context.Context) *AddWebsiteGroupParams {
	var ()
	return &AddWebsiteGroupParams{

		Context: ctx,
	}
}

// NewAddWebsiteGroupParamsWithHTTPClient creates a new AddWebsiteGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddWebsiteGroupParamsWithHTTPClient(client *http.Client) *AddWebsiteGroupParams {
	var ()
	return &AddWebsiteGroupParams{
		HTTPClient: client,
	}
}

/*AddWebsiteGroupParams contains all the parameters to send to the API endpoint
for the add website group operation typically these are written to a http.Request
*/
type AddWebsiteGroupParams struct {

	/*Body*/
	Body *models.WebsiteGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add website group params
func (o *AddWebsiteGroupParams) WithTimeout(timeout time.Duration) *AddWebsiteGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add website group params
func (o *AddWebsiteGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add website group params
func (o *AddWebsiteGroupParams) WithContext(ctx context.Context) *AddWebsiteGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add website group params
func (o *AddWebsiteGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add website group params
func (o *AddWebsiteGroupParams) WithHTTPClient(client *http.Client) *AddWebsiteGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add website group params
func (o *AddWebsiteGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add website group params
func (o *AddWebsiteGroupParams) WithBody(body *models.WebsiteGroup) *AddWebsiteGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add website group params
func (o *AddWebsiteGroupParams) SetBody(body *models.WebsiteGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddWebsiteGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
