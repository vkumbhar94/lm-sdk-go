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

// ExecuteDebugCommandReader is a Reader for the ExecuteDebugCommand structure.
type ExecuteDebugCommandReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExecuteDebugCommandReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExecuteDebugCommandOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewExecuteDebugCommandDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewExecuteDebugCommandOK creates a ExecuteDebugCommandOK with default headers values
func NewExecuteDebugCommandOK() *ExecuteDebugCommandOK {
	return &ExecuteDebugCommandOK{}
}

/*ExecuteDebugCommandOK handles this case with default header values.

successful operation
*/
type ExecuteDebugCommandOK struct {
	Payload *models.Debug
}

func (o *ExecuteDebugCommandOK) Error() string {
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommandOK  %+v", 200, o.Payload)
}

func (o *ExecuteDebugCommandOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Debug)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewExecuteDebugCommandDefault creates a ExecuteDebugCommandDefault with default headers values
func NewExecuteDebugCommandDefault(code int) *ExecuteDebugCommandDefault {
	return &ExecuteDebugCommandDefault{
		_statusCode: code,
	}
}

/*ExecuteDebugCommandDefault handles this case with default header values.

Error
*/
type ExecuteDebugCommandDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the execute debug command default response
func (o *ExecuteDebugCommandDefault) Code() int {
	return o._statusCode
}

func (o *ExecuteDebugCommandDefault) Error() string {
	return fmt.Sprintf("[POST /debug][%d] executeDebugCommand default  %+v", o._statusCode, o.Payload)
}

func (o *ExecuteDebugCommandDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
