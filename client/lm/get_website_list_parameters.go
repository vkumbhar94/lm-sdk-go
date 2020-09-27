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
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetWebsiteListParams creates a new GetWebsiteListParams object
// with the default values initialized.
func NewGetWebsiteListParams() *GetWebsiteListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWebsiteListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewGetWebsiteListParamsWithTimeout creates a new GetWebsiteListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetWebsiteListParamsWithTimeout(timeout time.Duration) *GetWebsiteListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWebsiteListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		timeout: timeout,
	}
}

// NewGetWebsiteListParamsWithContext creates a new GetWebsiteListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetWebsiteListParamsWithContext(ctx context.Context) *GetWebsiteListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWebsiteListParams{
		Offset: &offsetDefault,
		Size:   &sizeDefault,

		Context: ctx,
	}
}

// NewGetWebsiteListParamsWithHTTPClient creates a new GetWebsiteListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetWebsiteListParamsWithHTTPClient(client *http.Client) *GetWebsiteListParams {
	var (
		offsetDefault = int32(0)
		sizeDefault   = int32(50)
	)
	return &GetWebsiteListParams{
		Offset:     &offsetDefault,
		Size:       &sizeDefault,
		HTTPClient: client,
	}
}

/*GetWebsiteListParams contains all the parameters to send to the API endpoint
for the get website list operation typically these are written to a http.Request
*/
type GetWebsiteListParams struct {

	/*CollectorIds*/
	CollectorIds *string
	/*Fields*/
	Fields *string
	/*Filter*/
	Filter *string
	/*Offset*/
	Offset *int32
	/*Size*/
	Size *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get website list params
func (o *GetWebsiteListParams) WithTimeout(timeout time.Duration) *GetWebsiteListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get website list params
func (o *GetWebsiteListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get website list params
func (o *GetWebsiteListParams) WithContext(ctx context.Context) *GetWebsiteListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get website list params
func (o *GetWebsiteListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get website list params
func (o *GetWebsiteListParams) WithHTTPClient(client *http.Client) *GetWebsiteListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get website list params
func (o *GetWebsiteListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectorIds adds the collectorIds to the get website list params
func (o *GetWebsiteListParams) WithCollectorIds(collectorIds *string) *GetWebsiteListParams {
	o.SetCollectorIds(collectorIds)
	return o
}

// SetCollectorIds adds the collectorIds to the get website list params
func (o *GetWebsiteListParams) SetCollectorIds(collectorIds *string) {
	o.CollectorIds = collectorIds
}

// WithFields adds the fields to the get website list params
func (o *GetWebsiteListParams) WithFields(fields *string) *GetWebsiteListParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get website list params
func (o *GetWebsiteListParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFilter adds the filter to the get website list params
func (o *GetWebsiteListParams) WithFilter(filter *string) *GetWebsiteListParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the get website list params
func (o *GetWebsiteListParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithOffset adds the offset to the get website list params
func (o *GetWebsiteListParams) WithOffset(offset *int32) *GetWebsiteListParams {
	o.SetOffset(offset)
	return o
}

// SetOffset adds the offset to the get website list params
func (o *GetWebsiteListParams) SetOffset(offset *int32) {
	o.Offset = offset
}

// WithSize adds the size to the get website list params
func (o *GetWebsiteListParams) WithSize(size *int32) *GetWebsiteListParams {
	o.SetSize(size)
	return o
}

// SetSize adds the size to the get website list params
func (o *GetWebsiteListParams) SetSize(size *int32) {
	o.Size = size
}

// WriteToRequest writes these params to a swagger request
func (o *GetWebsiteListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.CollectorIds != nil {

		// query param collectorIds
		var qrCollectorIds string
		if o.CollectorIds != nil {
			qrCollectorIds = *o.CollectorIds
		}
		qCollectorIds := qrCollectorIds
		if qCollectorIds != "" {
			if err := r.SetQueryParam("collectorIds", qCollectorIds); err != nil {
				return err
			}
		}

	}

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	if o.Offset != nil {

		// query param offset
		var qrOffset int32
		if o.Offset != nil {
			qrOffset = *o.Offset
		}
		qOffset := swag.FormatInt32(qrOffset)
		if qOffset != "" {
			if err := r.SetQueryParam("offset", qOffset); err != nil {
				return err
			}
		}

	}

	if o.Size != nil {

		// query param size
		var qrSize int32
		if o.Size != nil {
			qrSize = *o.Size
		}
		qSize := swag.FormatInt32(qrSize)
		if qSize != "" {
			if err := r.SetQueryParam("size", qSize); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
