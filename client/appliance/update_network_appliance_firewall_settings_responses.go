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

// UpdateNetworkApplianceFirewallSettingsReader is a Reader for the UpdateNetworkApplianceFirewallSettings structure.
type UpdateNetworkApplianceFirewallSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceFirewallSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceFirewallSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/appliance/firewall/settings] updateNetworkApplianceFirewallSettings", response, response.Code())
	}
}

// NewUpdateNetworkApplianceFirewallSettingsOK creates a UpdateNetworkApplianceFirewallSettingsOK with default headers values
func NewUpdateNetworkApplianceFirewallSettingsOK() *UpdateNetworkApplianceFirewallSettingsOK {
	return &UpdateNetworkApplianceFirewallSettingsOK{}
}

/*
UpdateNetworkApplianceFirewallSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceFirewallSettingsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance firewall settings o k response has a 2xx status code
func (o *UpdateNetworkApplianceFirewallSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance firewall settings o k response has a 3xx status code
func (o *UpdateNetworkApplianceFirewallSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance firewall settings o k response has a 4xx status code
func (o *UpdateNetworkApplianceFirewallSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance firewall settings o k response has a 5xx status code
func (o *UpdateNetworkApplianceFirewallSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance firewall settings o k response a status code equal to that given
func (o *UpdateNetworkApplianceFirewallSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network appliance firewall settings o k response
func (o *UpdateNetworkApplianceFirewallSettingsOK) Code() int {
	return 200
}

func (o *UpdateNetworkApplianceFirewallSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/settings][%d] updateNetworkApplianceFirewallSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallSettingsOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/settings][%d] updateNetworkApplianceFirewallSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceFirewallSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkApplianceFirewallSettingsBody update network appliance firewall settings body
// Example: {"spoofingProtection":{"ipSourceGuard":{"mode":"block"}}}
swagger:model UpdateNetworkApplianceFirewallSettingsBody
*/
type UpdateNetworkApplianceFirewallSettingsBody struct {

	// spoofing protection
	SpoofingProtection *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection `json:"spoofingProtection,omitempty"`
}

// Validate validates this update network appliance firewall settings body
func (o *UpdateNetworkApplianceFirewallSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSpoofingProtection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallSettingsBody) validateSpoofingProtection(formats strfmt.Registry) error {
	if swag.IsZero(o.SpoofingProtection) { // not required
		return nil
	}

	if o.SpoofingProtection != nil {
		if err := o.SpoofingProtection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance firewall settings body based on the context it is used
func (o *UpdateNetworkApplianceFirewallSettingsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateSpoofingProtection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallSettingsBody) contextValidateSpoofingProtection(ctx context.Context, formats strfmt.Registry) error {

	if o.SpoofingProtection != nil {

		if swag.IsZero(o.SpoofingProtection) { // not required
			return nil
		}

		if err := o.SpoofingProtection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallSettingsBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection Spoofing protection settings
swagger:model UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection
*/
type UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection struct {

	// ip source guard
	IPSourceGuard *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard `json:"ipSourceGuard,omitempty"`
}

// Validate validates this update network appliance firewall settings params body spoofing protection
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPSourceGuard(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection) validateIPSourceGuard(formats strfmt.Registry) error {
	if swag.IsZero(o.IPSourceGuard) { // not required
		return nil
	}

	if o.IPSourceGuard != nil {
		if err := o.IPSourceGuard.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection" + "." + "ipSourceGuard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection" + "." + "ipSourceGuard")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance firewall settings params body spoofing protection based on the context it is used
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPSourceGuard(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection) contextValidateIPSourceGuard(ctx context.Context, formats strfmt.Registry) error {

	if o.IPSourceGuard != nil {

		if swag.IsZero(o.IPSourceGuard) { // not required
			return nil
		}

		if err := o.IPSourceGuard.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection" + "." + "ipSourceGuard")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceFirewallSettings" + "." + "spoofingProtection" + "." + "ipSourceGuard")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard IP source address spoofing settings
swagger:model UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard
*/
type UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard struct {

	// Mode of protection
	// Enum: [block log]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this update network appliance firewall settings params body spoofing protection IP source guard
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIpSourceGuardTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["block","log"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIpSourceGuardTypeModePropEnum = append(updateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIpSourceGuardTypeModePropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuardModeBlock captures enum value "block"
	UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuardModeBlock string = "block"

	// UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuardModeLog captures enum value "log"
	UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuardModeLog string = "log"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIpSourceGuardTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("updateNetworkApplianceFirewallSettings"+"."+"spoofingProtection"+"."+"ipSourceGuard"+"."+"mode", "body", o.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance firewall settings params body spoofing protection IP source guard based on context it is used
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallSettingsParamsBodySpoofingProtectionIPSourceGuard
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}