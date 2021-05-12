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

// UpdateCollectorGroupByIDReader is a Reader for the UpdateCollectorGroupByID structure.
type UpdateCollectorGroupByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCollectorGroupByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateCollectorGroupByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateCollectorGroupByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateCollectorGroupByIDOK creates a UpdateCollectorGroupByIDOK with default headers values
func NewUpdateCollectorGroupByIDOK() *UpdateCollectorGroupByIDOK {
	return &UpdateCollectorGroupByIDOK{}
}

/*UpdateCollectorGroupByIDOK handles this case with default header values.

successful operation
*/
type UpdateCollectorGroupByIDOK struct {
	Payload *models.CollectorGroup
}

func (o *UpdateCollectorGroupByIDOK) Error() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupByIdOK  %+v", 200, o.Payload)
}

func (o *UpdateCollectorGroupByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CollectorGroup)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectorGroupByIDDefault creates a UpdateCollectorGroupByIDDefault with default headers values
func NewUpdateCollectorGroupByIDDefault(code int) *UpdateCollectorGroupByIDDefault {
	return &UpdateCollectorGroupByIDDefault{
		_statusCode: code,
	}
}

/*UpdateCollectorGroupByIDDefault handles this case with default header values.

Error
*/
type UpdateCollectorGroupByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update collector group by Id default response
func (o *UpdateCollectorGroupByIDDefault) Code() int {
	return o._statusCode
}

func (o *UpdateCollectorGroupByIDDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/collector/groups/{id}][%d] updateCollectorGroupById default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateCollectorGroupByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
