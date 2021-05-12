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

// DeleteDeviceDatasourceInstanceByIDReader is a Reader for the DeleteDeviceDatasourceInstanceByID structure.
type DeleteDeviceDatasourceInstanceByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDeviceDatasourceInstanceByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDeviceDatasourceInstanceByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDeviceDatasourceInstanceByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDeviceDatasourceInstanceByIDOK creates a DeleteDeviceDatasourceInstanceByIDOK with default headers values
func NewDeleteDeviceDatasourceInstanceByIDOK() *DeleteDeviceDatasourceInstanceByIDOK {
	return &DeleteDeviceDatasourceInstanceByIDOK{}
}

/*DeleteDeviceDatasourceInstanceByIDOK handles this case with default header values.

successful operation
*/
type DeleteDeviceDatasourceInstanceByIDOK struct {
	Payload interface{}
}

func (o *DeleteDeviceDatasourceInstanceByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}][%d] deleteDeviceDatasourceInstanceByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteDeviceDatasourceInstanceByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDeviceDatasourceInstanceByIDDefault creates a DeleteDeviceDatasourceInstanceByIDDefault with default headers values
func NewDeleteDeviceDatasourceInstanceByIDDefault(code int) *DeleteDeviceDatasourceInstanceByIDDefault {
	return &DeleteDeviceDatasourceInstanceByIDDefault{
		_statusCode: code,
	}
}

/*DeleteDeviceDatasourceInstanceByIDDefault handles this case with default header values.

Error
*/
type DeleteDeviceDatasourceInstanceByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete device datasource instance by Id default response
func (o *DeleteDeviceDatasourceInstanceByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDeviceDatasourceInstanceByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}][%d] deleteDeviceDatasourceInstanceById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDeviceDatasourceInstanceByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
