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

// UpdateReportGroupByIDReader is a Reader for the UpdateReportGroupByID structure.
type UpdateReportGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateReportGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateReportGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateReportGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateReportGroupByIDOK creates a UpdateReportGroupByIDOK with default headers values
func NewUpdateReportGroupByIDOK() *UpdateReportGroupByIDOK {
	return &UpdateReportGroupByIDOK{}
}

/*UpdateReportGroupByIDOK handles this case with default header values.

successful operation
*/
type UpdateReportGroupByIDOK struct {
	Payload *models.ReportGroup
}

func (o *UpdateReportGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /report/groups/{id}][%d] updateReportGroupByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateReportGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateReportGroupByIDDefault creates a UpdateReportGroupByIDDefault with default headers values
func NewUpdateReportGroupByIDDefault(code int) *UpdateReportGroupByIDDefault {
	return &UpdateReportGroupByIDDefault{
		_statusCode: code,
	}
}

/*UpdateReportGroupByIDDefault handles this case with default header values.

Error
*/
type UpdateReportGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update report group by Id default response
func (o *UpdateReportGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateReportGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /report/groups/{id}][%d] updateReportGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateReportGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
