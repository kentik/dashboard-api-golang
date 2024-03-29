// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// UpdateDeviceCameraSenseReader is a Reader for the UpdateDeviceCameraSense structure.
type UpdateDeviceCameraSenseReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceCameraSenseReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceCameraSenseOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}/camera/sense] updateDeviceCameraSense", response, response.Code())
	}
}

// NewUpdateDeviceCameraSenseOK creates a UpdateDeviceCameraSenseOK with default headers values
func NewUpdateDeviceCameraSenseOK() *UpdateDeviceCameraSenseOK {
	return &UpdateDeviceCameraSenseOK{}
}

/*
UpdateDeviceCameraSenseOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceCameraSenseOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device camera sense o k response has a 2xx status code
func (o *UpdateDeviceCameraSenseOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device camera sense o k response has a 3xx status code
func (o *UpdateDeviceCameraSenseOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device camera sense o k response has a 4xx status code
func (o *UpdateDeviceCameraSenseOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device camera sense o k response has a 5xx status code
func (o *UpdateDeviceCameraSenseOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device camera sense o k response a status code equal to that given
func (o *UpdateDeviceCameraSenseOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device camera sense o k response
func (o *UpdateDeviceCameraSenseOK) Code() int {
	return 200
}

func (o *UpdateDeviceCameraSenseOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/camera/sense][%d] updateDeviceCameraSenseOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCameraSenseOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/camera/sense][%d] updateDeviceCameraSenseOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCameraSenseOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceCameraSenseOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceCameraSenseBody update device camera sense body
// Example: {"audioDetection":{"enabled":false},"mqttBrokerId":"1234","senseEnabled":true}
swagger:model UpdateDeviceCameraSenseBody
*/
type UpdateDeviceCameraSenseBody struct {

	// audio detection
	AudioDetection *UpdateDeviceCameraSenseParamsBodyAudioDetection `json:"audioDetection,omitempty"`

	// The ID of the object detection model
	DetectionModelID string `json:"detectionModelId,omitempty"`

	// The ID of the MQTT broker to be enabled on the camera. A value of null will disable MQTT on the camera
	MqttBrokerID string `json:"mqttBrokerId,omitempty"`

	// Boolean indicating if sense(license) is enabled(true) or disabled(false) on the camera
	SenseEnabled bool `json:"senseEnabled,omitempty"`
}

// Validate validates this update device camera sense body
func (o *UpdateDeviceCameraSenseBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAudioDetection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCameraSenseBody) validateAudioDetection(formats strfmt.Registry) error {
	if swag.IsZero(o.AudioDetection) { // not required
		return nil
	}

	if o.AudioDetection != nil {
		if err := o.AudioDetection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceCameraSense" + "." + "audioDetection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceCameraSense" + "." + "audioDetection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update device camera sense body based on the context it is used
func (o *UpdateDeviceCameraSenseBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAudioDetection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCameraSenseBody) contextValidateAudioDetection(ctx context.Context, formats strfmt.Registry) error {

	if o.AudioDetection != nil {

		if swag.IsZero(o.AudioDetection) { // not required
			return nil
		}

		if err := o.AudioDetection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceCameraSense" + "." + "audioDetection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceCameraSense" + "." + "audioDetection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCameraSenseBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCameraSenseBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCameraSenseBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCameraSenseParamsBodyAudioDetection The details of the audio detection config.
swagger:model UpdateDeviceCameraSenseParamsBodyAudioDetection
*/
type UpdateDeviceCameraSenseParamsBodyAudioDetection struct {

	// Boolean indicating if audio detection is enabled(true) or disabled(false) on the camera
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update device camera sense params body audio detection
func (o *UpdateDeviceCameraSenseParamsBodyAudioDetection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device camera sense params body audio detection based on context it is used
func (o *UpdateDeviceCameraSenseParamsBodyAudioDetection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCameraSenseParamsBodyAudioDetection) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCameraSenseParamsBodyAudioDetection) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCameraSenseParamsBodyAudioDetection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
