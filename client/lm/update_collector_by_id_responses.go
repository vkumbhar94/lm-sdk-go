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

// UpdateCollectorByIDReader is a Reader for the UpdateCollectorByID structure.
type UpdateCollectorByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCollectorByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCollectorByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateCollectorByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCollectorByIDOK creates a UpdateCollectorByIDOK with default headers values
func NewUpdateCollectorByIDOK() *UpdateCollectorByIDOK {
	return &UpdateCollectorByIDOK{}
}

/*UpdateCollectorByIDOK handles this case with default header values.

successful operation
*/
type UpdateCollectorByIDOK struct {
	Payload *models.Collector
}

func (o *UpdateCollectorByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/collector/collectors/{id}][%d] updateCollectorByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateCollectorByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Collector)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectorByIDDefault creates a UpdateCollectorByIDDefault with default headers values
func NewUpdateCollectorByIDDefault(code int) *UpdateCollectorByIDDefault {
	return &UpdateCollectorByIDDefault{
		_statusCode: code,
	}
}

/*UpdateCollectorByIDDefault handles this case with default header values.

Error
*/
type UpdateCollectorByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update collector by Id default response
func (o *UpdateCollectorByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCollectorByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/collector/collectors/{id}][%d] updateCollectorById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCollectorByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
