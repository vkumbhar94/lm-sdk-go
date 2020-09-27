// Code generated by go-swagger; DO NOT EDIT.

package lm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/vkumbhar94/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/vkumbhar94/lm-sdk-go/models"
)

// GetDeviceGroupDatasourceAlertSettingReader is a Reader for the GetDeviceGroupDatasourceAlertSetting structure.
type GetDeviceGroupDatasourceAlertSettingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceGroupDatasourceAlertSettingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceGroupDatasourceAlertSettingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceGroupDatasourceAlertSettingDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceGroupDatasourceAlertSettingOK creates a GetDeviceGroupDatasourceAlertSettingOK with default headers values
func NewGetDeviceGroupDatasourceAlertSettingOK() *GetDeviceGroupDatasourceAlertSettingOK {
	return &GetDeviceGroupDatasourceAlertSettingOK{}
}

/*GetDeviceGroupDatasourceAlertSettingOK handles this case with default header values.

successful operation
*/
type GetDeviceGroupDatasourceAlertSettingOK struct {
	Payload *models.DeviceGroupDataSourceAlertConfig
}

func (o *GetDeviceGroupDatasourceAlertSettingOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings][%d] getDeviceGroupDatasourceAlertSettingOK  %+v", 200, o.Payload)
}

func (o *GetDeviceGroupDatasourceAlertSettingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceGroupDataSourceAlertConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceGroupDatasourceAlertSettingDefault creates a GetDeviceGroupDatasourceAlertSettingDefault with default headers values
func NewGetDeviceGroupDatasourceAlertSettingDefault(code int) *GetDeviceGroupDatasourceAlertSettingDefault {
	return &GetDeviceGroupDatasourceAlertSettingDefault{
		_statusCode: code,
	}
}

/*GetDeviceGroupDatasourceAlertSettingDefault handles this case with default header values.

Error
*/
type GetDeviceGroupDatasourceAlertSettingDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device group datasource alert setting default response
func (o *GetDeviceGroupDatasourceAlertSettingDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceGroupDatasourceAlertSettingDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/datasources/{dsId}/alertsettings][%d] getDeviceGroupDatasourceAlertSetting default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceGroupDatasourceAlertSettingDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
