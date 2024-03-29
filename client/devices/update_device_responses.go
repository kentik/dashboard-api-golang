// Code generated by go-swagger; DO NOT EDIT.

package devices

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

// UpdateDeviceReader is a Reader for the UpdateDevice structure.
type UpdateDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}] updateDevice", response, response.Code())
	}
}

// NewUpdateDeviceOK creates a UpdateDeviceOK with default headers values
func NewUpdateDeviceOK() *UpdateDeviceOK {
	return &UpdateDeviceOK{}
}

/*
UpdateDeviceOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device o k response has a 2xx status code
func (o *UpdateDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device o k response has a 3xx status code
func (o *UpdateDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device o k response has a 4xx status code
func (o *UpdateDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device o k response has a 5xx status code
func (o *UpdateDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device o k response a status code equal to that given
func (o *UpdateDeviceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device o k response
func (o *UpdateDeviceOK) Code() int {
	return 200
}

func (o *UpdateDeviceOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}][%d] updateDeviceOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}][%d] updateDeviceOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceBody update device body
// Example: {"lat":37.4180951010362,"lng":-122.098531723022,"mac":"00:11:22:33:44:55","name":"My AP","serial":"Q234-ABCD-5678","tags":["recently-added"]}
swagger:model UpdateDeviceBody
*/
type UpdateDeviceBody struct {

	// The address of a device
	Address string `json:"address,omitempty"`

	// The floor plan to associate to this device. null disassociates the device from the floorplan.
	FloorPlanID string `json:"floorPlanId,omitempty"`

	// The latitude of a device
	Lat float32 `json:"lat,omitempty"`

	// The longitude of a device
	Lng float32 `json:"lng,omitempty"`

	// Whether or not to set the latitude and longitude of a device based on the new address. Only applies when lat and lng are not specified.
	MoveMapMarker bool `json:"moveMapMarker,omitempty"`

	// The name of a device
	Name string `json:"name,omitempty"`

	// The notes for the device. String. Limited to 255 characters.
	Notes string `json:"notes,omitempty"`

	// The ID of a switch template to bind to the device (for available switch templates, see the 'Switch Templates' endpoint). Use null to unbind the switch device from the current profile. For a device to be bindable to a switch template, it must (1) be a switch, and (2) belong to a network that is bound to a configuration template.
	SwitchProfileID string `json:"switchProfileId,omitempty"`

	// The list of tags of a device
	Tags []string `json:"tags"`
}

// Validate validates this update device body
func (o *UpdateDeviceBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device body based on context it is used
func (o *UpdateDeviceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
