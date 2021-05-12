// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/v2/models"
)

// GetWebsiteAlertListByWebsiteIDReader is a Reader for the GetWebsiteAlertListByWebsiteID structure.
type GetWebsiteAlertListByWebsiteIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebsiteAlertListByWebsiteIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetWebsiteAlertListByWebsiteIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetWebsiteAlertListByWebsiteIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetWebsiteAlertListByWebsiteIDOK creates a GetWebsiteAlertListByWebsiteIDOK with default headers values
func NewGetWebsiteAlertListByWebsiteIDOK() *GetWebsiteAlertListByWebsiteIDOK {
	return &GetWebsiteAlertListByWebsiteIDOK{}
}

/*GetWebsiteAlertListByWebsiteIDOK handles this case with default header values.

successful operation
*/
type GetWebsiteAlertListByWebsiteIDOK struct {
	Payload *models.AlertPaginationResponse
}

func (o *GetWebsiteAlertListByWebsiteIDOK) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/alerts][%d] getWebsiteAlertListByWebsiteIdOK  %+v", 200, o.Payload)
}

func (o *GetWebsiteAlertListByWebsiteIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebsiteAlertListByWebsiteIDDefault creates a GetWebsiteAlertListByWebsiteIDDefault with default headers values
func NewGetWebsiteAlertListByWebsiteIDDefault(code int) *GetWebsiteAlertListByWebsiteIDDefault {
	return &GetWebsiteAlertListByWebsiteIDDefault{
		_statusCode: code,
	}
}

/*GetWebsiteAlertListByWebsiteIDDefault handles this case with default header values.

Error
*/
type GetWebsiteAlertListByWebsiteIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get website alert list by website Id default response
func (o *GetWebsiteAlertListByWebsiteIDDefault) Code() int {
	return o._statusCode
}

func (o *GetWebsiteAlertListByWebsiteIDDefault) Error() string {
	return fmt.Sprintf("[GET /website/websites/{id}/alerts][%d] getWebsiteAlertListByWebsiteId default  %+v", o._statusCode, o.Payload)
}

func (o *GetWebsiteAlertListByWebsiteIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
