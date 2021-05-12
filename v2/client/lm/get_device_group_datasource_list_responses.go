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

// GetDeviceGroupDatasourceListReader is a Reader for the GetDeviceGroupDatasourceList structure.
type GetDeviceGroupDatasourceListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceGroupDatasourceListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceGroupDatasourceListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceGroupDatasourceListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceGroupDatasourceListOK creates a GetDeviceGroupDatasourceListOK with default headers values
func NewGetDeviceGroupDatasourceListOK() *GetDeviceGroupDatasourceListOK {
	return &GetDeviceGroupDatasourceListOK{}
}

/*GetDeviceGroupDatasourceListOK handles this case with default header values.

successful operation
*/
type GetDeviceGroupDatasourceListOK struct {
	Payload *models.DeviceGroupDatasourcePaginationResponse
}

func (o *GetDeviceGroupDatasourceListOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources][%d] getDeviceGroupDatasourceListOK  %+v", 200, o.Payload)
}

func (o *GetDeviceGroupDatasourceListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupDatasourcePaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceGroupDatasourceListDefault creates a GetDeviceGroupDatasourceListDefault with default headers values
func NewGetDeviceGroupDatasourceListDefault(code int) *GetDeviceGroupDatasourceListDefault {
	return &GetDeviceGroupDatasourceListDefault{
		_statusCode: code,
	}
}

/*GetDeviceGroupDatasourceListDefault handles this case with default header values.

Error
*/
type GetDeviceGroupDatasourceListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device group datasource list default response
func (o *GetDeviceGroupDatasourceListDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceGroupDatasourceListDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources][%d] getDeviceGroupDatasourceList default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceGroupDatasourceListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
