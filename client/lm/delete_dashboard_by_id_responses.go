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

// DeleteDashboardByIDReader is a Reader for the DeleteDashboardByID structure.
type DeleteDashboardByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteDashboardByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteDashboardByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteDashboardByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteDashboardByIDOK creates a DeleteDashboardByIDOK with default headers values
func NewDeleteDashboardByIDOK() *DeleteDashboardByIDOK {
	return &DeleteDashboardByIDOK{}
}

/*DeleteDashboardByIDOK handles this case with default header values.

successful operation
*/
type DeleteDashboardByIDOK struct {
	Payload interface{}
}

func (o *DeleteDashboardByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteDashboardByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteDashboardByIDDefault creates a DeleteDashboardByIDDefault with default headers values
func NewDeleteDashboardByIDDefault(code int) *DeleteDashboardByIDDefault {
	return &DeleteDashboardByIDDefault{
		_statusCode: code,
	}
}

/*DeleteDashboardByIDDefault handles this case with default header values.

Error
*/
type DeleteDashboardByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete dashboard by Id default response
func (o *DeleteDashboardByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteDashboardByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /dashboard/dashboards/{id}][%d] deleteDashboardById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteDashboardByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
