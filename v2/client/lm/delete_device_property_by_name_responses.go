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

// DeleteDevicePropertyByNameReader is a Reader for the DeleteDevicePropertyByName structure.
type DeleteDevicePropertyByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDevicePropertyByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDevicePropertyByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDevicePropertyByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDevicePropertyByNameOK creates a DeleteDevicePropertyByNameOK with default headers values
func NewDeleteDevicePropertyByNameOK() *DeleteDevicePropertyByNameOK {
	return &DeleteDevicePropertyByNameOK{}
}

/*DeleteDevicePropertyByNameOK handles this case with default header values.

successful operation
*/
type DeleteDevicePropertyByNameOK struct {
	Payload interface{}
}

func (o *DeleteDevicePropertyByNameOK) Error() string {
	return fmt.Sprintf("[DELETE /device/devices/{deviceId}/properties/{name}][%d] deleteDevicePropertyByNameOK  %+v", 200, o.Payload)
}

func (o *DeleteDevicePropertyByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDevicePropertyByNameDefault creates a DeleteDevicePropertyByNameDefault with default headers values
func NewDeleteDevicePropertyByNameDefault(code int) *DeleteDevicePropertyByNameDefault {
	return &DeleteDevicePropertyByNameDefault{
		_statusCode: code,
	}
}

/*DeleteDevicePropertyByNameDefault handles this case with default header values.

Error
*/
type DeleteDevicePropertyByNameDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete device property by name default response
func (o *DeleteDevicePropertyByNameDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDevicePropertyByNameDefault) Error() string {
	return fmt.Sprintf("[DELETE /device/devices/{deviceId}/properties/{name}][%d] deleteDevicePropertyByName default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDevicePropertyByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
