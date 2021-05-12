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

// PatchReportGroupByIDReader is a Reader for the PatchReportGroupByID structure.
type PatchReportGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchReportGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchReportGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchReportGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchReportGroupByIDOK creates a PatchReportGroupByIDOK with default headers values
func NewPatchReportGroupByIDOK() *PatchReportGroupByIDOK {
	return &PatchReportGroupByIDOK{}
}

/*PatchReportGroupByIDOK handles this case with default header values.

successful operation
*/
type PatchReportGroupByIDOK struct {
	Payload *models.ReportGroup
}

func (o *PatchReportGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /report/groups/{id}][%d] patchReportGroupByIdOK  %+v", 200, o.Payload)
}

func (o *PatchReportGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ReportGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchReportGroupByIDDefault creates a PatchReportGroupByIDDefault with default headers values
func NewPatchReportGroupByIDDefault(code int) *PatchReportGroupByIDDefault {
	return &PatchReportGroupByIDDefault{
		_statusCode: code,
	}
}

/*PatchReportGroupByIDDefault handles this case with default header values.

Error
*/
type PatchReportGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch report group by Id default response
func (o *PatchReportGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchReportGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /report/groups/{id}][%d] patchReportGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchReportGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
