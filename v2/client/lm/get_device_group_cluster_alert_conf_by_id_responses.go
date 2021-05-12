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

// GetDeviceGroupClusterAlertConfByIDReader is a Reader for the GetDeviceGroupClusterAlertConfByID structure.
type GetDeviceGroupClusterAlertConfByIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceGroupClusterAlertConfByIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetDeviceGroupClusterAlertConfByIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetDeviceGroupClusterAlertConfByIDDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetDeviceGroupClusterAlertConfByIDOK creates a GetDeviceGroupClusterAlertConfByIDOK with default headers values
func NewGetDeviceGroupClusterAlertConfByIDOK() *GetDeviceGroupClusterAlertConfByIDOK {
	return &GetDeviceGroupClusterAlertConfByIDOK{}
}

/*GetDeviceGroupClusterAlertConfByIDOK handles this case with default header values.

successful operation
*/
type GetDeviceGroupClusterAlertConfByIDOK struct {
	Payload *models.DeviceClusterAlertConfig
}

func (o *GetDeviceGroupClusterAlertConfByIDOK) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/clusterAlertConf/{id}][%d] getDeviceGroupClusterAlertConfByIdOK  %+v", 200, o.Payload)
}

func (o *GetDeviceGroupClusterAlertConfByIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DeviceClusterAlertConfig)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetDeviceGroupClusterAlertConfByIDDefault creates a GetDeviceGroupClusterAlertConfByIDDefault with default headers values
func NewGetDeviceGroupClusterAlertConfByIDDefault(code int) *GetDeviceGroupClusterAlertConfByIDDefault {
	return &GetDeviceGroupClusterAlertConfByIDDefault{
		_statusCode: code,
	}
}

/*GetDeviceGroupClusterAlertConfByIDDefault handles this case with default header values.

Error
*/
type GetDeviceGroupClusterAlertConfByIDDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get device group cluster alert conf by Id default response
func (o *GetDeviceGroupClusterAlertConfByIDDefault) Code() int {
	return o._statusCode
}

func (o *GetDeviceGroupClusterAlertConfByIDDefault) Error() string {
	return fmt.Sprintf("[GET /device/groups/{deviceGroupId}/clusterAlertConf/{id}][%d] getDeviceGroupClusterAlertConfById default  %+v", o._statusCode, o.Payload)
}

func (o *GetDeviceGroupClusterAlertConfByIDDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
