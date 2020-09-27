// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/vkumbhar94/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// UpdateSDTByIDReader is a Reader for the UpdateSDTByID structure.
type UpdateSDTByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateSDTByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateSDTByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateSDTByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateSDTByIDOK creates a UpdateSDTByIDOK with default headers values
func NewUpdateSDTByIDOK() *UpdateSDTByIDOK {
	return &UpdateSDTByIDOK{}
}

/*UpdateSDTByIDOK handles this case with default header values.

successful operation
*/
type UpdateSDTByIDOK struct {
	Payload models.SDT
}

func (o *UpdateSDTByIDOK) Error() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSdtByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateSDTByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalSDT(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateSDTByIDDefault creates a UpdateSDTByIDDefault with default headers values
func NewUpdateSDTByIDDefault(code int) *UpdateSDTByIDDefault {
	return &UpdateSDTByIDDefault{
		_statusCode: code,
	}
}

/*UpdateSDTByIDDefault handles this case with default header values.

Error
*/
type UpdateSDTByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update SDT by Id default response
func (o *UpdateSDTByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateSDTByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /sdt/sdts/{id}][%d] updateSDTById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateSDTByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
