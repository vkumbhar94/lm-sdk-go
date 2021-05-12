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

// GetAlertRuleByIDReader is a Reader for the GetAlertRuleByID structure.
type GetAlertRuleByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetAlertRuleByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetAlertRuleByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetAlertRuleByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetAlertRuleByIDOK creates a GetAlertRuleByIDOK with default headers values
func NewGetAlertRuleByIDOK() *GetAlertRuleByIDOK {
	return &GetAlertRuleByIDOK{}
}

/*GetAlertRuleByIDOK handles this case with default header values.

successful operation
*/
type GetAlertRuleByIDOK struct {
	Payload *models.AlertRule
}

func (o *GetAlertRuleByIDOK) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules/{id}][%d] getAlertRuleByIdOK  %+v", 200, o.Payload)
}

func (o *GetAlertRuleByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.AlertRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetAlertRuleByIDDefault creates a GetAlertRuleByIDDefault with default headers values
func NewGetAlertRuleByIDDefault(code int) *GetAlertRuleByIDDefault {
	return &GetAlertRuleByIDDefault{
		_statusCode: code,
	}
}

/*GetAlertRuleByIDDefault handles this case with default header values.

Error
*/
type GetAlertRuleByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get alert rule by Id default response
func (o *GetAlertRuleByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetAlertRuleByIDDefault) Error() string {
	return fmt.Sprintf("[GET /setting/alert/rules/{id}][%d] getAlertRuleById default  %+v", o._statusCode, o.Payload)
}

func (o *GetAlertRuleByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
