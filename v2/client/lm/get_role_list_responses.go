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

// GetRoleListReader is a Reader for the GetRoleList structure.
type GetRoleListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetRoleListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetRoleListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetRoleListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetRoleListOK creates a GetRoleListOK with default headers values
func NewGetRoleListOK() *GetRoleListOK {
	return &GetRoleListOK{}
}

/*GetRoleListOK handles this case with default header values.

successful operation
*/
type GetRoleListOK struct {
	Payload *models.RolePaginationResponse
}

func (o *GetRoleListOK) Error() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleListOK  %+v", 200, o.Payload)
}

func (o *GetRoleListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RolePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetRoleListDefault creates a GetRoleListDefault with default headers values
func NewGetRoleListDefault(code int) *GetRoleListDefault {
	return &GetRoleListDefault{
		_statusCode: code,
	}
}

/*GetRoleListDefault handles this case with default header values.

Error
*/
type GetRoleListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get role list default response
func (o *GetRoleListDefault) Code() int {
	return o._statusCode
}

func (o *GetRoleListDefault) Error() string {
	return fmt.Sprintf("[GET /setting/roles][%d] getRoleList default  %+v", o._statusCode, o.Payload)
}

func (o *GetRoleListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
