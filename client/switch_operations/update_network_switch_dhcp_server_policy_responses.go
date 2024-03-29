// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// UpdateNetworkSwitchDhcpServerPolicyReader is a Reader for the UpdateNetworkSwitchDhcpServerPolicy structure.
type UpdateNetworkSwitchDhcpServerPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchDhcpServerPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchDhcpServerPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/switch/dhcpServerPolicy] updateNetworkSwitchDhcpServerPolicy", response, response.Code())
	}
}

// NewUpdateNetworkSwitchDhcpServerPolicyOK creates a UpdateNetworkSwitchDhcpServerPolicyOK with default headers values
func NewUpdateNetworkSwitchDhcpServerPolicyOK() *UpdateNetworkSwitchDhcpServerPolicyOK {
	return &UpdateNetworkSwitchDhcpServerPolicyOK{}
}

/*
UpdateNetworkSwitchDhcpServerPolicyOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchDhcpServerPolicyOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network switch dhcp server policy o k response has a 2xx status code
func (o *UpdateNetworkSwitchDhcpServerPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network switch dhcp server policy o k response has a 3xx status code
func (o *UpdateNetworkSwitchDhcpServerPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network switch dhcp server policy o k response has a 4xx status code
func (o *UpdateNetworkSwitchDhcpServerPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network switch dhcp server policy o k response has a 5xx status code
func (o *UpdateNetworkSwitchDhcpServerPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network switch dhcp server policy o k response a status code equal to that given
func (o *UpdateNetworkSwitchDhcpServerPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network switch dhcp server policy o k response
func (o *UpdateNetworkSwitchDhcpServerPolicyOK) Code() int {
	return 200
}

func (o *UpdateNetworkSwitchDhcpServerPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/dhcpServerPolicy][%d] updateNetworkSwitchDhcpServerPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchDhcpServerPolicyOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/dhcpServerPolicy][%d] updateNetworkSwitchDhcpServerPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchDhcpServerPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchDhcpServerPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSwitchDhcpServerPolicyBody update network switch dhcp server policy body
// Example: {"alerts":{"email":{"enabled":true}},"allowedServers":["00:50:56:00:00:01","00:50:56:00:00:02"],"arpInspection":{"enabled":true},"blockedServers":["00:50:56:00:00:03","00:50:56:00:00:04"],"defaultPolicy":"block"}
swagger:model UpdateNetworkSwitchDhcpServerPolicyBody
*/
type UpdateNetworkSwitchDhcpServerPolicyBody struct {

	// alerts
	Alerts *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts `json:"alerts,omitempty"`

	// List the MAC addresses of DHCP servers to permit on the network when defaultPolicy is set to block. An empty array will clear the entries.
	AllowedServers []string `json:"allowedServers"`

	// arp inspection
	ArpInspection *UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection `json:"arpInspection,omitempty"`

	// List the MAC addresses of DHCP servers to block on the network when defaultPolicy is set to allow. An empty array will clear the entries.
	BlockedServers []string `json:"blockedServers"`

	// 'allow' or 'block' new DHCP servers. Default value is 'allow'.
	// Enum: [allow block]
	DefaultPolicy string `json:"defaultPolicy,omitempty"`
}

// Validate validates this update network switch dhcp server policy body
func (o *UpdateNetworkSwitchDhcpServerPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateArpInspection(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchDhcpServerPolicyBody) validateAlerts(formats strfmt.Registry) error {
	if swag.IsZero(o.Alerts) { // not required
		return nil
	}

	if o.Alerts != nil {
		if err := o.Alerts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchDhcpServerPolicyBody) validateArpInspection(formats strfmt.Registry) error {
	if swag.IsZero(o.ArpInspection) { // not required
		return nil
	}

	if o.ArpInspection != nil {
		if err := o.ArpInspection.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "arpInspection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "arpInspection")
			}
			return err
		}
	}

	return nil
}

var updateNetworkSwitchDhcpServerPolicyBodyTypeDefaultPolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","block"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchDhcpServerPolicyBodyTypeDefaultPolicyPropEnum = append(updateNetworkSwitchDhcpServerPolicyBodyTypeDefaultPolicyPropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchDhcpServerPolicyBodyDefaultPolicyAllow captures enum value "allow"
	UpdateNetworkSwitchDhcpServerPolicyBodyDefaultPolicyAllow string = "allow"

	// UpdateNetworkSwitchDhcpServerPolicyBodyDefaultPolicyBlock captures enum value "block"
	UpdateNetworkSwitchDhcpServerPolicyBodyDefaultPolicyBlock string = "block"
)

// prop value enum
func (o *UpdateNetworkSwitchDhcpServerPolicyBody) validateDefaultPolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchDhcpServerPolicyBodyTypeDefaultPolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchDhcpServerPolicyBody) validateDefaultPolicy(formats strfmt.Registry) error {
	if swag.IsZero(o.DefaultPolicy) { // not required
		return nil
	}

	// value enum
	if err := o.validateDefaultPolicyEnum("updateNetworkSwitchDhcpServerPolicy"+"."+"defaultPolicy", "body", o.DefaultPolicy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update network switch dhcp server policy body based on the context it is used
func (o *UpdateNetworkSwitchDhcpServerPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateArpInspection(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchDhcpServerPolicyBody) contextValidateAlerts(ctx context.Context, formats strfmt.Registry) error {

	if o.Alerts != nil {

		if swag.IsZero(o.Alerts) { // not required
			return nil
		}

		if err := o.Alerts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkSwitchDhcpServerPolicyBody) contextValidateArpInspection(ctx context.Context, formats strfmt.Registry) error {

	if o.ArpInspection != nil {

		if swag.IsZero(o.ArpInspection) { // not required
			return nil
		}

		if err := o.ArpInspection.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "arpInspection")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "arpInspection")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchDhcpServerPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts Alert settings for DHCP servers
swagger:model UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts
*/
type UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts struct {

	// email
	Email *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail `json:"email,omitempty"`
}

// Validate validates this update network switch dhcp server policy params body alerts
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts) validateEmail(formats strfmt.Registry) error {
	if swag.IsZero(o.Email) { // not required
		return nil
	}

	if o.Email != nil {
		if err := o.Email.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts" + "." + "email")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts" + "." + "email")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network switch dhcp server policy params body alerts based on the context it is used
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEmail(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts) contextValidateEmail(ctx context.Context, formats strfmt.Registry) error {

	if o.Email != nil {

		if swag.IsZero(o.Email) { // not required
			return nil
		}

		if err := o.Email.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts" + "." + "email")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkSwitchDhcpServerPolicy" + "." + "alerts" + "." + "email")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlerts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail Email alert settings for DHCP servers
swagger:model UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail
*/
type UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail struct {

	// When enabled, send an email if a new DHCP server is seen. Default value is false.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network switch dhcp server policy params body alerts email
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch dhcp server policy params body alerts email based on context it is used
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchDhcpServerPolicyParamsBodyAlertsEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection Dynamic ARP Inspection settings
swagger:model UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection
*/
type UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection struct {

	// Enable or disable Dynamic ARP Inspection on the network. Default value is false.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network switch dhcp server policy params body arp inspection
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network switch dhcp server policy params body arp inspection based on context it is used
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchDhcpServerPolicyParamsBodyArpInspection
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
