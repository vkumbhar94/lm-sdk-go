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

// GetNetflowEndpointListReader is a Reader for the GetNetflowEndpointList structure.
type GetNetflowEndpointListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetflowEndpointListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetflowEndpointListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetflowEndpointListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetflowEndpointListOK creates a GetNetflowEndpointListOK with default headers values
func NewGetNetflowEndpointListOK() *GetNetflowEndpointListOK {
	return &GetNetflowEndpointListOK{}
}

/*GetNetflowEndpointListOK handles this case with default header values.

successful operation
*/
type GetNetflowEndpointListOK struct {
	Payload *models.EndpointPaginationResponse
}

func (o *GetNetflowEndpointListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/endpoints][%d] getNetflowEndpointListOK  %+v", 200, o.Payload)
}

func (o *GetNetflowEndpointListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EndpointPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetflowEndpointListDefault creates a GetNetflowEndpointListDefault with default headers values
func NewGetNetflowEndpointListDefault(code int) *GetNetflowEndpointListDefault {
	return &GetNetflowEndpointListDefault{
		_statusCode: code,
	}
}

/*GetNetflowEndpointListDefault handles this case with default header values.

Error
*/
type GetNetflowEndpointListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get netflow endpoint list default response
func (o *GetNetflowEndpointListDefault) Code() int {
	return o._statusCode
}

func (o *GetNetflowEndpointListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/endpoints][%d] getNetflowEndpointList default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetflowEndpointListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
