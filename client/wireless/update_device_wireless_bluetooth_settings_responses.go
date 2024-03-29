// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateDeviceWirelessBluetoothSettingsReader is a Reader for the UpdateDeviceWirelessBluetoothSettings structure.
type UpdateDeviceWirelessBluetoothSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceWirelessBluetoothSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceWirelessBluetoothSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}/wireless/bluetooth/settings] updateDeviceWirelessBluetoothSettings", response, response.Code())
	}
}

// NewUpdateDeviceWirelessBluetoothSettingsOK creates a UpdateDeviceWirelessBluetoothSettingsOK with default headers values
func NewUpdateDeviceWirelessBluetoothSettingsOK() *UpdateDeviceWirelessBluetoothSettingsOK {
	return &UpdateDeviceWirelessBluetoothSettingsOK{}
}

/*
UpdateDeviceWirelessBluetoothSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceWirelessBluetoothSettingsOK struct {
	Payload *UpdateDeviceWirelessBluetoothSettingsOKBody
}

// IsSuccess returns true when this update device wireless bluetooth settings o k response has a 2xx status code
func (o *UpdateDeviceWirelessBluetoothSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device wireless bluetooth settings o k response has a 3xx status code
func (o *UpdateDeviceWirelessBluetoothSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device wireless bluetooth settings o k response has a 4xx status code
func (o *UpdateDeviceWirelessBluetoothSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device wireless bluetooth settings o k response has a 5xx status code
func (o *UpdateDeviceWirelessBluetoothSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device wireless bluetooth settings o k response a status code equal to that given
func (o *UpdateDeviceWirelessBluetoothSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device wireless bluetooth settings o k response
func (o *UpdateDeviceWirelessBluetoothSettingsOK) Code() int {
	return 200
}

func (o *UpdateDeviceWirelessBluetoothSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/wireless/bluetooth/settings][%d] updateDeviceWirelessBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceWirelessBluetoothSettingsOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/wireless/bluetooth/settings][%d] updateDeviceWirelessBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceWirelessBluetoothSettingsOK) GetPayload() *UpdateDeviceWirelessBluetoothSettingsOKBody {
	return o.Payload
}

func (o *UpdateDeviceWirelessBluetoothSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateDeviceWirelessBluetoothSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceWirelessBluetoothSettingsBody update device wireless bluetooth settings body
// Example: {"major":13,"minor":125,"uuid":"00000000-0000-0000-000-000000000000"}
swagger:model UpdateDeviceWirelessBluetoothSettingsBody
*/
type UpdateDeviceWirelessBluetoothSettingsBody struct {

	// Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Major int64 `json:"major,omitempty"`

	// Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor int64 `json:"minor,omitempty"`

	// Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this update device wireless bluetooth settings body
func (o *UpdateDeviceWirelessBluetoothSettingsBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device wireless bluetooth settings body based on context it is used
func (o *UpdateDeviceWirelessBluetoothSettingsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceWirelessBluetoothSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceWirelessBluetoothSettingsBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceWirelessBluetoothSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceWirelessBluetoothSettingsOKBody update device wireless bluetooth settings o k body
swagger:model UpdateDeviceWirelessBluetoothSettingsOKBody
*/
type UpdateDeviceWirelessBluetoothSettingsOKBody struct {

	// Desired major value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Major int64 `json:"major,omitempty"`

	// Desired minor value of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	Minor int64 `json:"minor,omitempty"`

	// Desired UUID of the beacon. If the value is set to null it will reset to Dashboard's automatically generated value.
	UUID string `json:"uuid,omitempty"`
}

// Validate validates this update device wireless bluetooth settings o k body
func (o *UpdateDeviceWirelessBluetoothSettingsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device wireless bluetooth settings o k body based on context it is used
func (o *UpdateDeviceWirelessBluetoothSettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceWirelessBluetoothSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceWirelessBluetoothSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceWirelessBluetoothSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
