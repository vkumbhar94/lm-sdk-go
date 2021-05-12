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

// DeleteSDTByIDReader is a Reader for the DeleteSDTByID structure.
type DeleteSDTByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSDTByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteSDTByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteSDTByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteSDTByIDOK creates a DeleteSDTByIDOK with default headers values
func NewDeleteSDTByIDOK() *DeleteSDTByIDOK {
	return &DeleteSDTByIDOK{}
}

/*DeleteSDTByIDOK handles this case with default header values.

successful operation
*/
type DeleteSDTByIDOK struct {
	Payload interface{}
}

func (o *DeleteSDTByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /sdt/sdts/{id}][%d] deleteSdtByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteSDTByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSDTByIDDefault creates a DeleteSDTByIDDefault with default headers values
func NewDeleteSDTByIDDefault(code int) *DeleteSDTByIDDefault {
	return &DeleteSDTByIDDefault{
		_statusCode: code,
	}
}

/*DeleteSDTByIDDefault handles this case with default header values.

Error
*/
type DeleteSDTByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete SDT by Id default response
func (o *DeleteSDTByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteSDTByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /sdt/sdts/{id}][%d] deleteSDTById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteSDTByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
