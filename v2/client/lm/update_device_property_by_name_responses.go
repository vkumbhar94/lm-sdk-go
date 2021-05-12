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

// UpdateDevicePropertyByNameReader is a Reader for the UpdateDevicePropertyByName structure.
type UpdateDevicePropertyByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDevicePropertyByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewUpdateDevicePropertyByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewUpdateDevicePropertyByNameDefault(response.Code())
		if result.Code() == 429 {
			errResp := &models.ErrorResponse{
				ErrorCode: 429,
				ErrorDetail: map[string]interface{}{
					"x-rate-limit-limit":     response.GetHeader("x-rate-limit-limit"),
					"x-rate-limit-remaining": response.GetHeader("x-rate-limit-remaining"),
					"x-rate-limit-window":    response.GetHeader("x-rate-limit-window"),
				},
				ErrorMessage: "Customized response from argus sdk",
			}
			result.Payload = errResp
			return nil, result
		}
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewUpdateDevicePropertyByNameOK creates a UpdateDevicePropertyByNameOK with default headers values
func NewUpdateDevicePropertyByNameOK() *UpdateDevicePropertyByNameOK {
	return &UpdateDevicePropertyByNameOK{}
}

/*UpdateDevicePropertyByNameOK handles this case with default header values.

successful operation
*/
type UpdateDevicePropertyByNameOK struct {
	Payload *models.EntityProperty
}

func (o *UpdateDevicePropertyByNameOK) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByNameOK  %+v", 200, o.Payload)
}

func (o *UpdateDevicePropertyByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.EntityProperty)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateDevicePropertyByNameDefault creates a UpdateDevicePropertyByNameDefault with default headers values
func NewUpdateDevicePropertyByNameDefault(code int) *UpdateDevicePropertyByNameDefault {
	return &UpdateDevicePropertyByNameDefault{
		_statusCode: code,
	}
}

/*UpdateDevicePropertyByNameDefault handles this case with default header values.

Error
*/
type UpdateDevicePropertyByNameDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the update device property by name default response
func (o *UpdateDevicePropertyByNameDefault) Code() int {
	return o._statusCode
}

func (o *UpdateDevicePropertyByNameDefault) Error() string {
	return fmt.Sprintf("[PUT /device/devices/{deviceId}/properties/{name}][%d] updateDevicePropertyByName default  %+v", o._statusCode, o.Payload)
}

func (o *UpdateDevicePropertyByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
