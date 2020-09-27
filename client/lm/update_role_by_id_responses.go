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

// UpdateRoleByIDReader is a Reader for the UpdateRoleByID structure.
type UpdateRoleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateRoleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateRoleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateRoleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateRoleByIDOK creates a UpdateRoleByIDOK with default headers values
func NewUpdateRoleByIDOK() *UpdateRoleByIDOK {
	return &UpdateRoleByIDOK{}
}

/*UpdateRoleByIDOK handles this case with default header values.

successful operation
*/
type UpdateRoleByIDOK struct {
	Payload *models.Role
}

func (o *UpdateRoleByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/roles/{id}][%d] updateRoleByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateRoleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Role)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateRoleByIDDefault creates a UpdateRoleByIDDefault with default headers values
func NewUpdateRoleByIDDefault(code int) *UpdateRoleByIDDefault {
	return &UpdateRoleByIDDefault{
		_statusCode: code,
	}
}

/*UpdateRoleByIDDefault handles this case with default header values.

Error
*/
type UpdateRoleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update role by Id default response
func (o *UpdateRoleByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateRoleByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/roles/{id}][%d] updateRoleById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateRoleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
