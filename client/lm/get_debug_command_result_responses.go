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

// GetDebugCommandResultReader is a Reader for the GetDebugCommandResult structure.
type GetDebugCommandResultReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDebugCommandResultReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDebugCommandResultOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDebugCommandResultDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDebugCommandResultOK creates a GetDebugCommandResultOK with default headers values
func NewGetDebugCommandResultOK() *GetDebugCommandResultOK {
	return &GetDebugCommandResultOK{}
}

/*GetDebugCommandResultOK handles this case with default header values.

successful operation
*/
type GetDebugCommandResultOK struct {
	Payload *models.Debug
}

func (o *GetDebugCommandResultOK) Error() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResultOK  %+v", 200, o.Payload)
}

func (o *GetDebugCommandResultOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Debug)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDebugCommandResultDefault creates a GetDebugCommandResultDefault with default headers values
func NewGetDebugCommandResultDefault(code int) *GetDebugCommandResultDefault {
	return &GetDebugCommandResultDefault{
		_statusCode: code,
	}
}

/*GetDebugCommandResultDefault handles this case with default header values.

Error
*/
type GetDebugCommandResultDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get debug command result default response
func (o *GetDebugCommandResultDefault) Code() int {
	return o._statusCode
}

func (o *GetDebugCommandResultDefault) Error() string {
	return fmt.Sprintf("[GET /debug/{id}][%d] getDebugCommandResult default  %+v", o._statusCode, o.Payload)
}

func (o *GetDebugCommandResultDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
