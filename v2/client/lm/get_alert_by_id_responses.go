// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/logicmonitor/lm-sdk-go/v2/models"
)

// GetAlertByIDReader is a Reader for the GetAlertByID structure.
type GetAlertByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlertByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAlertByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertByIDOK creates a GetAlertByIDOK with default headers values
func NewGetAlertByIDOK() *GetAlertByIDOK {
	return &GetAlertByIDOK{}
}

/*GetAlertByIDOK handles this case with default header values.

successful operation
*/
type GetAlertByIDOK struct {
	Payload *models.Alert
}

func (o *GetAlertByIDOK) Error() string {
	return fmt.Sprintf("[GET /alert/alerts/{id}][%d] getAlertByIdOK  %+v", 200, o.Payload)
}

func (o *GetAlertByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Alert)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertByIDDefault creates a GetAlertByIDDefault with default headers values
func NewGetAlertByIDDefault(code int) *GetAlertByIDDefault {
	return &GetAlertByIDDefault{
		_statusCode: code,
	}
}

/*GetAlertByIDDefault handles this case with default header values.

Error
*/
type GetAlertByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alert by Id default response
func (o *GetAlertByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertByIDDefault) Error() string {
	return fmt.Sprintf("[GET /alert/alerts/{id}][%d] getAlertById default  %+v", o._statusCode, o.Payload)
}

func (o *GetAlertByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
