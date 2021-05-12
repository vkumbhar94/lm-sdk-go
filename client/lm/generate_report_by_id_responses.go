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

// GenerateReportByIDReader is a Reader for the GenerateReportByID structure.
type GenerateReportByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateReportByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGenerateReportByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGenerateReportByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGenerateReportByIDOK creates a GenerateReportByIDOK with default headers values
func NewGenerateReportByIDOK() *GenerateReportByIDOK {
	return &GenerateReportByIDOK{}
}

/*GenerateReportByIDOK handles this case with default header values.

successful operation
*/
type GenerateReportByIDOK struct {
	Payload *models.GenerateReportResult
}

func (o *GenerateReportByIDOK) Error() string {
	return fmt.Sprintf("[POST /report/reports/{id}/executions][%d] generateReportByIdOK  %+v", 200, o.Payload)
}

func (o *GenerateReportByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenerateReportResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGenerateReportByIDDefault creates a GenerateReportByIDDefault with default headers values
func NewGenerateReportByIDDefault(code int) *GenerateReportByIDDefault {
	return &GenerateReportByIDDefault{
		_statusCode: code,
	}
}

/*GenerateReportByIDDefault handles this case with default header values.

Error
*/
type GenerateReportByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the generate report by Id default response
func (o *GenerateReportByIDDefault) Code() int {
	return o._statusCode
}

func (o *GenerateReportByIDDefault) Error() string {
	return fmt.Sprintf("[POST /report/reports/{id}/executions][%d] generateReportById default  %+v", o._statusCode, o.Payload)
}

func (o *GenerateReportByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
