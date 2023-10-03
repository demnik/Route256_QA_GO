// Code generated by go-swagger; DO NOT EDIT.

package act_device_api_service

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/ozonmp/act-device-api/homework/hw2/test/http_clients/generated_client/models"
)

// ActDeviceAPIServiceUpdateDeviceV1Reader is a Reader for the ActDeviceAPIServiceUpdateDeviceV1 structure.
type ActDeviceAPIServiceUpdateDeviceV1Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ActDeviceAPIServiceUpdateDeviceV1Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewActDeviceAPIServiceUpdateDeviceV1OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewActDeviceAPIServiceUpdateDeviceV1Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewActDeviceAPIServiceUpdateDeviceV1OK creates a ActDeviceAPIServiceUpdateDeviceV1OK with default headers values
func NewActDeviceAPIServiceUpdateDeviceV1OK() *ActDeviceAPIServiceUpdateDeviceV1OK {
	return &ActDeviceAPIServiceUpdateDeviceV1OK{}
}

/*
ActDeviceAPIServiceUpdateDeviceV1OK describes a response with status code 200, with default header values.

A successful response.
*/
type ActDeviceAPIServiceUpdateDeviceV1OK struct {
	Payload *models.V1UpdateDeviceV1Response
}

// IsSuccess returns true when this act device Api service update device v1 o k response has a 2xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this act device Api service update device v1 o k response has a 3xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this act device Api service update device v1 o k response has a 4xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this act device Api service update device v1 o k response has a 5xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1OK) IsServerError() bool {
	return false
}

// IsCode returns true when this act device Api service update device v1 o k response a status code equal to that given
func (o *ActDeviceAPIServiceUpdateDeviceV1OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the act device Api service update device v1 o k response
func (o *ActDeviceAPIServiceUpdateDeviceV1OK) Code() int {
	return 200
}

func (o *ActDeviceAPIServiceUpdateDeviceV1OK) Error() string {
	return fmt.Sprintf("[PUT /api/v1/devices/{deviceId}][%d] actDeviceApiServiceUpdateDeviceV1OK  %+v", 200, o.Payload)
}

func (o *ActDeviceAPIServiceUpdateDeviceV1OK) String() string {
	return fmt.Sprintf("[PUT /api/v1/devices/{deviceId}][%d] actDeviceApiServiceUpdateDeviceV1OK  %+v", 200, o.Payload)
}

func (o *ActDeviceAPIServiceUpdateDeviceV1OK) GetPayload() *models.V1UpdateDeviceV1Response {
	return o.Payload
}

func (o *ActDeviceAPIServiceUpdateDeviceV1OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.V1UpdateDeviceV1Response)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewActDeviceAPIServiceUpdateDeviceV1Default creates a ActDeviceAPIServiceUpdateDeviceV1Default with default headers values
func NewActDeviceAPIServiceUpdateDeviceV1Default(code int) *ActDeviceAPIServiceUpdateDeviceV1Default {
	return &ActDeviceAPIServiceUpdateDeviceV1Default{
		_statusCode: code,
	}
}

/*
ActDeviceAPIServiceUpdateDeviceV1Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type ActDeviceAPIServiceUpdateDeviceV1Default struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this act device Api service update device v1 default response has a 2xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this act device Api service update device v1 default response has a 3xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this act device Api service update device v1 default response has a 4xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this act device Api service update device v1 default response has a 5xx status code
func (o *ActDeviceAPIServiceUpdateDeviceV1Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this act device Api service update device v1 default response a status code equal to that given
func (o *ActDeviceAPIServiceUpdateDeviceV1Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the act device Api service update device v1 default response
func (o *ActDeviceAPIServiceUpdateDeviceV1Default) Code() int {
	return o._statusCode
}

func (o *ActDeviceAPIServiceUpdateDeviceV1Default) Error() string {
	return fmt.Sprintf("[PUT /api/v1/devices/{deviceId}][%d] ActDeviceApiService_UpdateDeviceV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ActDeviceAPIServiceUpdateDeviceV1Default) String() string {
	return fmt.Sprintf("[PUT /api/v1/devices/{deviceId}][%d] ActDeviceApiService_UpdateDeviceV1 default  %+v", o._statusCode, o.Payload)
}

func (o *ActDeviceAPIServiceUpdateDeviceV1Default) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *ActDeviceAPIServiceUpdateDeviceV1Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
ActDeviceAPIServiceUpdateDeviceV1Body act device API service update device v1 body
swagger:model ActDeviceAPIServiceUpdateDeviceV1Body
*/
type ActDeviceAPIServiceUpdateDeviceV1Body struct {

	// platform
	Platform string `json:"platform,omitempty"`

	// user Id
	UserID string `json:"userId,omitempty"`
}

// Validate validates this act device API service update device v1 body
func (o *ActDeviceAPIServiceUpdateDeviceV1Body) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this act device API service update device v1 body based on context it is used
func (o *ActDeviceAPIServiceUpdateDeviceV1Body) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ActDeviceAPIServiceUpdateDeviceV1Body) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ActDeviceAPIServiceUpdateDeviceV1Body) UnmarshalBinary(b []byte) error {
	var res ActDeviceAPIServiceUpdateDeviceV1Body
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}