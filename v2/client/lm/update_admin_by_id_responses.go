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

// UpdateAdminByIDReader is a Reader for the UpdateAdminByID structure.
type UpdateAdminByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateAdminByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateAdminByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateAdminByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateAdminByIDOK creates a UpdateAdminByIDOK with default headers values
func NewUpdateAdminByIDOK() *UpdateAdminByIDOK {
	return &UpdateAdminByIDOK{}
}

/*UpdateAdminByIDOK handles this case with default header values.

successful operation
*/
type UpdateAdminByIDOK struct {
	Payload *models.Admin
}

func (o *UpdateAdminByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/admins/{id}][%d] updateAdminByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateAdminByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Admin)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateAdminByIDDefault creates a UpdateAdminByIDDefault with default headers values
func NewUpdateAdminByIDDefault(code int) *UpdateAdminByIDDefault {
	return &UpdateAdminByIDDefault{
		_statusCode: code,
	}
}

/*UpdateAdminByIDDefault handles this case with default header values.

Error
*/
type UpdateAdminByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update admin by Id default response
func (o *UpdateAdminByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateAdminByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/admins/{id}][%d] updateAdminById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateAdminByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
