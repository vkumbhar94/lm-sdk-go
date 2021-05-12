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

// GetNetflowFlowListReader is a Reader for the GetNetflowFlowList structure.
type GetNetflowFlowListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetflowFlowListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetNetflowFlowListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewGetNetflowFlowListDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetNetflowFlowListOK creates a GetNetflowFlowListOK with default headers values
func NewGetNetflowFlowListOK() *GetNetflowFlowListOK {
	return &GetNetflowFlowListOK{}
}

/*GetNetflowFlowListOK handles this case with default header values.

successful operation
*/
type GetNetflowFlowListOK struct {
	Payload *models.FlowRecordPaginationResponse
}

func (o *GetNetflowFlowListOK) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowListOK  %+v", 200, o.Payload)
}

func (o *GetNetflowFlowListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.FlowRecordPaginationResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetNetflowFlowListDefault creates a GetNetflowFlowListDefault with default headers values
func NewGetNetflowFlowListDefault(code int) *GetNetflowFlowListDefault {
	return &GetNetflowFlowListDefault{
		_statusCode: code,
	}
}

/*GetNetflowFlowListDefault handles this case with default header values.

Error
*/
type GetNetflowFlowListDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get netflow flow list default response
func (o *GetNetflowFlowListDefault) Code() int {
	return o._statusCode
}

func (o *GetNetflowFlowListDefault) Error() string {
	return fmt.Sprintf("[GET /device/devices/{id}/flows][%d] getNetflowFlowList default  %+v", o._statusCode, o.Payload)
}

func (o *GetNetflowFlowListDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
