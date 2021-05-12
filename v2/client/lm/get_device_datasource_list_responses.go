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

// GetDeviceDatasourceListReader is a Reader for the GetDeviceDatasourceList structure.
type GetDeviceDatasourceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceDatasourceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceDatasourceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceListOK creates a GetDeviceDatasourceListOK with default headers values
func NewGetDeviceDatasourceListOK() *GetDeviceDatasourceListOK {
	return &GetDeviceDatasourceListOK{}
}

/*GetDeviceDatasourceListOK handles this case with default header values.

successful operation
*/
type GetDeviceDatasourceListOK struct {
	Payload *models.DeviceDatasourcePaginationResponse
}

func (o *GetDeviceDatasourceListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources][%d] getDeviceDatasourceListOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDatasourcePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceListDefault creates a GetDeviceDatasourceListDefault with default headers values
func NewGetDeviceDatasourceListDefault(code int) *GetDeviceDatasourceListDefault {
	return &GetDeviceDatasourceListDefault{
		_statusCode: code,
	}
}

/*GetDeviceDatasourceListDefault handles this case with default header values.

Error
*/
type GetDeviceDatasourceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource list default response
func (o *GetDeviceDatasourceListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources][%d] getDeviceDatasourceList default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
