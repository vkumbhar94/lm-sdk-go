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

// UpdateNetscanReader is a Reader for the UpdateNetscan structure.
type UpdateNetscanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetscanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateNetscanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateNetscanDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateNetscanOK creates a UpdateNetscanOK with default headers values
func NewUpdateNetscanOK() *UpdateNetscanOK {
	return &UpdateNetscanOK{}
}

/*UpdateNetscanOK handles this case with default header values.

successful operation
*/
type UpdateNetscanOK struct {
	Payload models.Netscan
}

func (o *UpdateNetscanOK) Error() string {
	return fmt.Sprintf("[PUT /setting/netscans/{id}][%d] updateNetscanOK  %+v", 200, o.Payload)
}

func (o *UpdateNetscanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload as interface type
	payload, err := models.UnmarshalNetscan(response.Body(), consumer)
	if err != nil {
		return err
	}
	o.Payload = payload

	return nil
}

// NewUpdateNetscanDefault creates a UpdateNetscanDefault with default headers values
func NewUpdateNetscanDefault(code int) *UpdateNetscanDefault {
	return &UpdateNetscanDefault{
		_statusCode: code,
	}
}

/*UpdateNetscanDefault handles this case with default header values.

Error
*/
type UpdateNetscanDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update netscan default response
func (o *UpdateNetscanDefault) Code() int {
	return o._statusCode
}

func (o *UpdateNetscanDefault) Error() string {
	return fmt.Sprintf("[PUT /setting/netscans/{id}][%d] updateNetscan default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateNetscanDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
