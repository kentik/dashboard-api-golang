// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateDeviceApplianceRadioSettingsReader is a Reader for the UpdateDeviceApplianceRadioSettings structure.
type UpdateDeviceApplianceRadioSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceApplianceRadioSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceApplianceRadioSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}/appliance/radio/settings] updateDeviceApplianceRadioSettings", response, response.Code())
	}
}

// NewUpdateDeviceApplianceRadioSettingsOK creates a UpdateDeviceApplianceRadioSettingsOK with default headers values
func NewUpdateDeviceApplianceRadioSettingsOK() *UpdateDeviceApplianceRadioSettingsOK {
	return &UpdateDeviceApplianceRadioSettingsOK{}
}

/*
UpdateDeviceApplianceRadioSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceApplianceRadioSettingsOK struct {
	Payload *UpdateDeviceApplianceRadioSettingsOKBody
}

// IsSuccess returns true when this update device appliance radio settings o k response has a 2xx status code
func (o *UpdateDeviceApplianceRadioSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device appliance radio settings o k response has a 3xx status code
func (o *UpdateDeviceApplianceRadioSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device appliance radio settings o k response has a 4xx status code
func (o *UpdateDeviceApplianceRadioSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device appliance radio settings o k response has a 5xx status code
func (o *UpdateDeviceApplianceRadioSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device appliance radio settings o k response a status code equal to that given
func (o *UpdateDeviceApplianceRadioSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device appliance radio settings o k response
func (o *UpdateDeviceApplianceRadioSettingsOK) Code() int {
	return 200
}

func (o *UpdateDeviceApplianceRadioSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/appliance/radio/settings][%d] updateDeviceApplianceRadioSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceApplianceRadioSettingsOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/appliance/radio/settings][%d] updateDeviceApplianceRadioSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceApplianceRadioSettingsOK) GetPayload() *UpdateDeviceApplianceRadioSettingsOKBody {
	return o.Payload
}

func (o *UpdateDeviceApplianceRadioSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateDeviceApplianceRadioSettingsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceApplianceRadioSettingsBody update device appliance radio settings body
// Example: {"fiveGhzSettings":{"channel":149,"channelWidth":20,"targetPower":15},"rfProfileId":"1234","twoFourGhzSettings":{"channel":11,"targetPower":21}}
swagger:model UpdateDeviceApplianceRadioSettingsBody
*/
type UpdateDeviceApplianceRadioSettingsBody struct {

	// five ghz settings
	FiveGhzSettings *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings `json:"fiveGhzSettings,omitempty"`

	// The ID of an RF profile to assign to the device. If the value of this parameter is null, the appropriate basic RF profile (indoor or outdoor) will be assigned to the device. Assigning an RF profile will clear ALL manually configured overrides on the device (channel width, channel, power).
	RfProfileID string `json:"rfProfileId,omitempty"`

	// two four ghz settings
	TwoFourGhzSettings *UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
}

// Validate validates this update device appliance radio settings body
func (o *UpdateDeviceApplianceRadioSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFiveGhzSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTwoFourGhzSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsBody) validateFiveGhzSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.FiveGhzSettings) { // not required
		return nil
	}

	if o.FiveGhzSettings != nil {
		if err := o.FiveGhzSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettings" + "." + "fiveGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettings" + "." + "fiveGhzSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsBody) validateTwoFourGhzSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.TwoFourGhzSettings) { // not required
		return nil
	}

	if o.TwoFourGhzSettings != nil {
		if err := o.TwoFourGhzSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettings" + "." + "twoFourGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettings" + "." + "twoFourGhzSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update device appliance radio settings body based on the context it is used
func (o *UpdateDeviceApplianceRadioSettingsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFiveGhzSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTwoFourGhzSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsBody) contextValidateFiveGhzSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.FiveGhzSettings != nil {

		if swag.IsZero(o.FiveGhzSettings) { // not required
			return nil
		}

		if err := o.FiveGhzSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettings" + "." + "fiveGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettings" + "." + "fiveGhzSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsBody) contextValidateTwoFourGhzSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.TwoFourGhzSettings != nil {

		if swag.IsZero(o.TwoFourGhzSettings) { // not required
			return nil
		}

		if err := o.TwoFourGhzSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettings" + "." + "twoFourGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettings" + "." + "twoFourGhzSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceApplianceRadioSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceApplianceRadioSettingsOKBody update device appliance radio settings o k body
swagger:model UpdateDeviceApplianceRadioSettingsOKBody
*/
type UpdateDeviceApplianceRadioSettingsOKBody struct {

	// five ghz settings
	FiveGhzSettings *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings `json:"fiveGhzSettings,omitempty"`

	// RF Profile ID
	RfProfileID string `json:"rfProfileId,omitempty"`

	// The device serial
	Serial string `json:"serial,omitempty"`

	// two four ghz settings
	TwoFourGhzSettings *UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings `json:"twoFourGhzSettings,omitempty"`
}

// Validate validates this update device appliance radio settings o k body
func (o *UpdateDeviceApplianceRadioSettingsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFiveGhzSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTwoFourGhzSettings(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsOKBody) validateFiveGhzSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.FiveGhzSettings) { // not required
		return nil
	}

	if o.FiveGhzSettings != nil {
		if err := o.FiveGhzSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "fiveGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "fiveGhzSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsOKBody) validateTwoFourGhzSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.TwoFourGhzSettings) { // not required
		return nil
	}

	if o.TwoFourGhzSettings != nil {
		if err := o.TwoFourGhzSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "twoFourGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "twoFourGhzSettings")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update device appliance radio settings o k body based on the context it is used
func (o *UpdateDeviceApplianceRadioSettingsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFiveGhzSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTwoFourGhzSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsOKBody) contextValidateFiveGhzSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.FiveGhzSettings != nil {

		if swag.IsZero(o.FiveGhzSettings) { // not required
			return nil
		}

		if err := o.FiveGhzSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "fiveGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "fiveGhzSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsOKBody) contextValidateTwoFourGhzSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.TwoFourGhzSettings != nil {

		if swag.IsZero(o.TwoFourGhzSettings) { // not required
			return nil
		}

		if err := o.TwoFourGhzSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "twoFourGhzSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceApplianceRadioSettingsOK" + "." + "twoFourGhzSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceApplianceRadioSettingsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings Manual radio settings for 5 GHz
swagger:model UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings
*/
type UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings struct {

	// Manual channel for 5 GHz
	// Enum: [36 40 44 48 52 56 60 64 100 104 108 112 116 120 124 128 132 136 140 144 149 153 157 161 165 169 173 177]
	Channel int64 `json:"channel,omitempty"`

	// Manual channel width for 5 GHz
	// Enum: [0 20 40 80 160]
	ChannelWidth int64 `json:"channelWidth,omitempty"`

	// Manual target power for 5 GHz
	TargetPower int64 `json:"targetPower,omitempty"`
}

// Validate validates this update device appliance radio settings o k body five ghz settings
func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165,169,173,177]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelPropEnum = append(updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelPropEnum, v)
	}
}

// prop value enum
func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) validateChannelEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(o.Channel) { // not required
		return nil
	}

	// value enum
	if err := o.validateChannelEnum("updateDeviceApplianceRadioSettingsOK"+"."+"fiveGhzSettings"+"."+"channel", "body", o.Channel); err != nil {
		return err
	}

	return nil
}

var updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelWidthPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,20,40,80,160]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelWidthPropEnum = append(updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelWidthPropEnum, v)
	}
}

// prop value enum
func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) validateChannelWidthEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, updateDeviceApplianceRadioSettingsOKBodyFiveGhzSettingsTypeChannelWidthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) validateChannelWidth(formats strfmt.Registry) error {
	if swag.IsZero(o.ChannelWidth) { // not required
		return nil
	}

	// value enum
	if err := o.validateChannelWidthEnum("updateDeviceApplianceRadioSettingsOK"+"."+"fiveGhzSettings"+"."+"channelWidth", "body", o.ChannelWidth); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device appliance radio settings o k body five ghz settings based on context it is used
func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceApplianceRadioSettingsOKBodyFiveGhzSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings Manual radio settings for 2.4 GHz
swagger:model UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings
*/
type UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings struct {

	// Manual channel for 2.4 GHz
	// Enum: [1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	Channel int64 `json:"channel,omitempty"`

	// Manual target power for 2.4 GHz
	TargetPower int64 `json:"targetPower,omitempty"`
}

// Validate validates this update device appliance radio settings o k body two four ghz settings
func (o *UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettingsTypeChannelPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9,10,11,12,13,14]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettingsTypeChannelPropEnum = append(updateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettingsTypeChannelPropEnum, v)
	}
}

// prop value enum
func (o *UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings) validateChannelEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, updateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettingsTypeChannelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(o.Channel) { // not required
		return nil
	}

	// value enum
	if err := o.validateChannelEnum("updateDeviceApplianceRadioSettingsOK"+"."+"twoFourGhzSettings"+"."+"channel", "body", o.Channel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device appliance radio settings o k body two four ghz settings based on context it is used
func (o *UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceApplianceRadioSettingsOKBodyTwoFourGhzSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings Manual radio settings for 5 GHz.
swagger:model UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings
*/
type UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings struct {

	// Sets a manual channel for 5 GHz. Can be '36', '40', '44', '48', '52', '56', '60', '64', '100', '104', '108', '112', '116', '120', '124', '128', '132', '136', '140', '144', '149', '153', '157', '161', '165', '169', '173' or '177' or null for using auto channel.
	// Enum: [36 40 44 48 52 56 60 64 100 104 108 112 116 120 124 128 132 136 140 144 149 153 157 161 165 169 173 177]
	Channel int64 `json:"channel,omitempty"`

	// Sets a manual channel width for 5 GHz. Can be '0', '20', '40', '80' or '160' or null for using auto channel width.
	// Enum: [0 20 40 80 160]
	ChannelWidth int64 `json:"channelWidth,omitempty"`

	// Set a manual target power for 5 GHz. Can be between '8' or '30' or null for using auto power range.
	TargetPower int64 `json:"targetPower,omitempty"`
}

// Validate validates this update device appliance radio settings params body five ghz settings
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateChannelWidth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[36,40,44,48,52,56,60,64,100,104,108,112,116,120,124,128,132,136,140,144,149,153,157,161,165,169,173,177]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelPropEnum = append(updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelPropEnum, v)
	}
}

// prop value enum
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) validateChannelEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(o.Channel) { // not required
		return nil
	}

	// value enum
	if err := o.validateChannelEnum("updateDeviceApplianceRadioSettings"+"."+"fiveGhzSettings"+"."+"channel", "body", o.Channel); err != nil {
		return err
	}

	return nil
}

var updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelWidthPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[0,20,40,80,160]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelWidthPropEnum = append(updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelWidthPropEnum, v)
	}
}

// prop value enum
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) validateChannelWidthEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, updateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettingsTypeChannelWidthPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) validateChannelWidth(formats strfmt.Registry) error {
	if swag.IsZero(o.ChannelWidth) { // not required
		return nil
	}

	// value enum
	if err := o.validateChannelWidthEnum("updateDeviceApplianceRadioSettings"+"."+"fiveGhzSettings"+"."+"channelWidth", "body", o.ChannelWidth); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device appliance radio settings params body five ghz settings based on context it is used
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceApplianceRadioSettingsParamsBodyFiveGhzSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings Manual radio settings for 2.4 GHz.
swagger:model UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings
*/
type UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings struct {

	// Sets a manual channel for 2.4 GHz. Can be '1', '2', '3', '4', '5', '6', '7', '8', '9', '10', '11', '12', '13' or '14' or null for using auto channel.
	// Enum: [1 2 3 4 5 6 7 8 9 10 11 12 13 14]
	Channel int64 `json:"channel,omitempty"`

	// Set a manual target power for 2.4 GHz. Can be between '5' or '30' or null for using auto power range.
	TargetPower int64 `json:"targetPower,omitempty"`
}

// Validate validates this update device appliance radio settings params body two four ghz settings
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateChannel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettingsTypeChannelPropEnum []interface{}

func init() {
	var res []int64
	if err := json.Unmarshal([]byte(`[1,2,3,4,5,6,7,8,9,10,11,12,13,14]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettingsTypeChannelPropEnum = append(updateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettingsTypeChannelPropEnum, v)
	}
}

// prop value enum
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings) validateChannelEnum(path, location string, value int64) error {
	if err := validate.EnumCase(path, location, value, updateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettingsTypeChannelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings) validateChannel(formats strfmt.Registry) error {
	if swag.IsZero(o.Channel) { // not required
		return nil
	}

	// value enum
	if err := o.validateChannelEnum("updateDeviceApplianceRadioSettings"+"."+"twoFourGhzSettings"+"."+"channel", "body", o.Channel); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device appliance radio settings params body two four ghz settings based on context it is used
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceApplianceRadioSettingsParamsBodyTwoFourGhzSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}