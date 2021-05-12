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

// DeleteWebsiteByIDReader is a Reader for the DeleteWebsiteByID structure.
type DeleteWebsiteByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteWebsiteByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDeleteWebsiteByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteWebsiteByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteWebsiteByIDOK creates a DeleteWebsiteByIDOK with default headers values
func NewDeleteWebsiteByIDOK() *DeleteWebsiteByIDOK {
	return &DeleteWebsiteByIDOK{}
}

/*DeleteWebsiteByIDOK handles this case with default header values.

successful operation
*/
type DeleteWebsiteByIDOK struct {
	Payload interface{}
}

func (o *DeleteWebsiteByIDOK) Error() string {
	return fmt.Sprintf("[DELETE /website/websites/{id}][%d] deleteWebsiteByIdOK  %+v", 200, o.Payload)
}

func (o *DeleteWebsiteByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteWebsiteByIDDefault creates a DeleteWebsiteByIDDefault with default headers values
func NewDeleteWebsiteByIDDefault(code int) *DeleteWebsiteByIDDefault {
	return &DeleteWebsiteByIDDefault{
		_statusCode: code,
	}
}

/*DeleteWebsiteByIDDefault handles this case with default header values.

Error
*/
type DeleteWebsiteByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the delete website by Id default response
func (o *DeleteWebsiteByIDDefault) Code() int {
	return o._statusCode
}

func (o *DeleteWebsiteByIDDefault) Error() string {
	return fmt.Sprintf("[DELETE /website/websites/{id}][%d] deleteWebsiteById default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteWebsiteByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
