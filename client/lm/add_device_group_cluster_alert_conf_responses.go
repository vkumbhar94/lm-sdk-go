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

// AddDeviceGroupClusterAlertConfReader is a Reader for the AddDeviceGroupClusterAlertConf structure.
type AddDeviceGroupClusterAlertConfReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddDeviceGroupClusterAlertConfReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddDeviceGroupClusterAlertConfOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewAddDeviceGroupClusterAlertConfDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewAddDeviceGroupClusterAlertConfOK creates a AddDeviceGroupClusterAlertConfOK with default headers values
func NewAddDeviceGroupClusterAlertConfOK() *AddDeviceGroupClusterAlertConfOK {
	return &AddDeviceGroupClusterAlertConfOK{}
}

/*AddDeviceGroupClusterAlertConfOK handles this case with default header values.

successful operation
*/
type AddDeviceGroupClusterAlertConfOK struct {
	Payload *models.DeviceClusterAlertConfig
}

func (o *AddDeviceGroupClusterAlertConfOK) Error() string {
	return fmt.Sprintf("[POST /device/groups/{deviceGroupId}/clusterAlertConf][%d] addDeviceGroupClusterAlertConfOK  %+v", 200, o.Payload)
}

func (o *AddDeviceGroupClusterAlertConfOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceClusterAlertConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewAddDeviceGroupClusterAlertConfDefault creates a AddDeviceGroupClusterAlertConfDefault with default headers values
func NewAddDeviceGroupClusterAlertConfDefault(code int) *AddDeviceGroupClusterAlertConfDefault {
	return &AddDeviceGroupClusterAlertConfDefault{
		_statusCode: code,
	}
}

/*AddDeviceGroupClusterAlertConfDefault handles this case with default header values.

Error
*/
type AddDeviceGroupClusterAlertConfDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the add device group cluster alert conf default response
func (o *AddDeviceGroupClusterAlertConfDefault) Code() int {
	return o._statusCode
}

func (o *AddDeviceGroupClusterAlertConfDefault) Error() string {
	return fmt.Sprintf("[POST /device/groups/{deviceGroupId}/clusterAlertConf][%d] addDeviceGroupClusterAlertConf default  %+v", o._statusCode, o.Payload)
}

func (o *AddDeviceGroupClusterAlertConfDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
