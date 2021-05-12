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

// GetDeviceDatasourceInstanceSDTHistoryReader is a Reader for the GetDeviceDatasourceInstanceSDTHistory structure.
type GetDeviceDatasourceInstanceSDTHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceSDTHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceDatasourceInstanceSDTHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceDatasourceInstanceSDTHistoryDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceSDTHistoryOK creates a GetDeviceDatasourceInstanceSDTHistoryOK with default headers values
func NewGetDeviceDatasourceInstanceSDTHistoryOK() *GetDeviceDatasourceInstanceSDTHistoryOK {
	return &GetDeviceDatasourceInstanceSDTHistoryOK{}
}

/*GetDeviceDatasourceInstanceSDTHistoryOK handles this case with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceSDTHistoryOK struct {
	Payload *models.DeviceGroupSDTHistoryPaginationResponse
}

func (o *GetDeviceDatasourceInstanceSDTHistoryOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/historysdts][%d] getDeviceDatasourceInstanceSdtHistoryOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceSDTHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupSDTHistoryPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceSDTHistoryDefault creates a GetDeviceDatasourceInstanceSDTHistoryDefault with default headers values
func NewGetDeviceDatasourceInstanceSDTHistoryDefault(code int) *GetDeviceDatasourceInstanceSDTHistoryDefault {
	return &GetDeviceDatasourceInstanceSDTHistoryDefault{
		_statusCode: code,
	}
}

/*GetDeviceDatasourceInstanceSDTHistoryDefault handles this case with default header values.

Error
*/
type GetDeviceDatasourceInstanceSDTHistoryDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance SDT history default response
func (o *GetDeviceDatasourceInstanceSDTHistoryDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceSDTHistoryDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{id}/historysdts][%d] getDeviceDatasourceInstanceSDTHistory default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceSDTHistoryDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
