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

// NewAddReportGroupParams creates a new AddReportGroupParams object
// with the default values initialized.
func NewAddReportGroupParams() *AddReportGroupParams {
	var ()
	return &AddReportGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAddReportGroupParamsWithTimeout creates a new AddReportGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAddReportGroupParamsWithTimeout(timeout time.Duration) *AddReportGroupParams {
	var ()
	return &AddReportGroupParams{

		timeout: timeout,
	}
}

// NewAddReportGroupParamsWithContext creates a new AddReportGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewAddReportGroupParamsWithContext(ctx context.Context) *AddReportGroupParams {
	var ()
	return &AddReportGroupParams{

		Context: ctx,
	}
}

// NewAddReportGroupParamsWithHTTPClient creates a new AddReportGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAddReportGroupParamsWithHTTPClient(client *http.Client) *AddReportGroupParams {
	var ()
	return &AddReportGroupParams{
		HTTPClient: client,
	}
}

/*AddReportGroupParams contains all the parameters to send to the API endpoint
for the add report group operation typically these are written to a http.Request
*/
type AddReportGroupParams struct {

	/*Body*/
	Body *models.ReportGroup

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the add report group params
func (o *AddReportGroupParams) WithTimeout(timeout time.Duration) *AddReportGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the add report group params
func (o *AddReportGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the add report group params
func (o *AddReportGroupParams) WithContext(ctx context.Context) *AddReportGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the add report group params
func (o *AddReportGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the add report group params
func (o *AddReportGroupParams) WithHTTPClient(client *http.Client) *AddReportGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the add report group params
func (o *AddReportGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the add report group params
func (o *AddReportGroupParams) WithBody(body *models.ReportGroup) *AddReportGroupParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the add report group params
func (o *AddReportGroupParams) SetBody(body *models.ReportGroup) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *AddReportGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
