// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// GetNetflowPortListReader is a Reader for the GetNetflowPortList structure.
type GetNetflowPortListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetflowPortListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetflowPortListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetflowPortListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetflowPortListOK creates a GetNetflowPortListOK with default headers values
func NewGetNetflowPortListOK() *GetNetflowPortListOK {
	return &GetNetflowPortListOK{}
}

/*GetNetflowPortListOK handles this case with default header values.

successful operation
*/
type GetNetflowPortListOK struct {
	Payload *models.PortPaginationResponse
}

func (o *GetNetflowPortListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/ports][%d] getNetflowPortListOK  %+v", 200, o.Payload)
}

func (o *GetNetflowPortListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PortPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetflowPortListDefault creates a GetNetflowPortListDefault with default headers values
func NewGetNetflowPortListDefault(code int) *GetNetflowPortListDefault {
	return &GetNetflowPortListDefault{
		_statusCode: code,
	}
}

/*GetNetflowPortListDefault handles this case with default header values.

Error
*/
type GetNetflowPortListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get netflow port list default response
func (o *GetNetflowPortListDefault) Code() int {
	return o._statusCode
}

func (o *GetNetflowPortListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/ports][%d] getNetflowPortList default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetflowPortListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
