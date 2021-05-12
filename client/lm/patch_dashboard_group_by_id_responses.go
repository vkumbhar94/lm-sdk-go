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

// PatchDashboardGroupByIDReader is a Reader for the PatchDashboardGroupByID structure.
type PatchDashboardGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PatchDashboardGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewPatchDashboardGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewPatchDashboardGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewPatchDashboardGroupByIDOK creates a PatchDashboardGroupByIDOK with default headers values
func NewPatchDashboardGroupByIDOK() *PatchDashboardGroupByIDOK {
	return &PatchDashboardGroupByIDOK{}
}

/*PatchDashboardGroupByIDOK handles this case with default header values.

successful operation
*/
type PatchDashboardGroupByIDOK struct {
	Payload *models.DashboardGroup
}

func (o *PatchDashboardGroupByIDOK) Error() string {
	return fmt.Sprintf("[PATCH /dashboard/groups/{id}][%d] patchDashboardGroupByIdOK  %+v", 200, o.Payload)
}

func (o *PatchDashboardGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DashboardGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPatchDashboardGroupByIDDefault creates a PatchDashboardGroupByIDDefault with default headers values
func NewPatchDashboardGroupByIDDefault(code int) *PatchDashboardGroupByIDDefault {
	return &PatchDashboardGroupByIDDefault{
		_statusCode: code,
	}
}

/*PatchDashboardGroupByIDDefault handles this case with default header values.

Error
*/
type PatchDashboardGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the patch dashboard group by Id default response
func (o *PatchDashboardGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *PatchDashboardGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PATCH /dashboard/groups/{id}][%d] patchDashboardGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *PatchDashboardGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
