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

// GetDeviceDatasourceInstanceAlertSettingByIDReader is a Reader for the GetDeviceDatasourceInstanceAlertSettingByID structure.
type GetDeviceDatasourceInstanceAlertSettingByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceDatasourceInstanceAlertSettingByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceDatasourceInstanceAlertSettingByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceDatasourceInstanceAlertSettingByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceDatasourceInstanceAlertSettingByIDOK creates a GetDeviceDatasourceInstanceAlertSettingByIDOK with default headers values
func NewGetDeviceDatasourceInstanceAlertSettingByIDOK() *GetDeviceDatasourceInstanceAlertSettingByIDOK {
	return &GetDeviceDatasourceInstanceAlertSettingByIDOK{}
}

/*GetDeviceDatasourceInstanceAlertSettingByIDOK handles this case with default header values.

successful operation
*/
type GetDeviceDatasourceInstanceAlertSettingByIDOK struct {
	Payload *models.DeviceDataSourceInstanceAlertSetting
}

func (o *GetDeviceDatasourceInstanceAlertSettingByIDOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] getDeviceDatasourceInstanceAlertSettingByIdOK  %+v", 200, o.Payload)
}

func (o *GetDeviceDatasourceInstanceAlertSettingByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceDataSourceInstanceAlertSetting)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceDatasourceInstanceAlertSettingByIDDefault creates a GetDeviceDatasourceInstanceAlertSettingByIDDefault with default headers values
func NewGetDeviceDatasourceInstanceAlertSettingByIDDefault(code int) *GetDeviceDatasourceInstanceAlertSettingByIDDefault {
	return &GetDeviceDatasourceInstanceAlertSettingByIDDefault{
		_statusCode: code,
	}
}

/*GetDeviceDatasourceInstanceAlertSettingByIDDefault handles this case with default header values.

Error
*/
type GetDeviceDatasourceInstanceAlertSettingByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device datasource instance alert setting by Id default response
func (o *GetDeviceDatasourceInstanceAlertSettingByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceDatasourceInstanceAlertSettingByIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{deviceId}/devicedatasources/{hdsId}/instances/{instanceId}/alertsettings/{id}][%d] getDeviceDatasourceInstanceAlertSettingById default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceDatasourceInstanceAlertSettingByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
