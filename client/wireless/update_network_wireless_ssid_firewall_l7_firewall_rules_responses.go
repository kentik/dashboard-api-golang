// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkWirelessSsidFirewallL7FirewallRulesReader is a Reader for the UpdateNetworkWirelessSsidFirewallL7FirewallRules structure.
type UpdateNetworkWirelessSsidFirewallL7FirewallRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules] updateNetworkWirelessSsidFirewallL7FirewallRules", response, response.Code())
	}
}

// NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesOK creates a UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK with default headers values
func NewUpdateNetworkWirelessSsidFirewallL7FirewallRulesOK() *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK {
	return &UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK{}
}

/*
UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network wireless ssid firewall l7 firewall rules o k response has a 2xx status code
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network wireless ssid firewall l7 firewall rules o k response has a 3xx status code
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network wireless ssid firewall l7 firewall rules o k response has a 4xx status code
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network wireless ssid firewall l7 firewall rules o k response has a 5xx status code
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network wireless ssid firewall l7 firewall rules o k response a status code equal to that given
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network wireless ssid firewall l7 firewall rules o k response
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) Code() int {
	return 200
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules][%d] updateNetworkWirelessSsidFirewallL7FirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/ssids/{number}/firewall/l7FirewallRules][%d] updateNetworkWirelessSsidFirewallL7FirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody update network wireless ssid firewall l7 firewall rules body
// Example: {"rules":[{"policy":"deny","type":"host","value":"google.com"},{"policy":"deny","type":"port","value":"23"},{"policy":"deny","type":"ipRange","value":"10.11.12.00/24"},{"policy":"deny","type":"ipRange","value":"10.11.12.00/24:5555"}]}
swagger:model UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody
*/
type UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody struct {

	// An array of L7 firewall rules for this SSID. Rules will get applied in the same order user has specified in request. Empty array will clear the L7 firewall rule configuration.
	Rules []*UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0 `json:"rules"`
}

// Validate validates this update network wireless ssid firewall l7 firewall rules body
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) validateRules(formats strfmt.Registry) error {
	if swag.IsZero(o.Rules) { // not required
		return nil
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessSsidFirewallL7FirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessSsidFirewallL7FirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network wireless ssid firewall l7 firewall rules body based on the context it is used
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {

			if swag.IsZero(o.Rules[i]) { // not required
				return nil
			}

			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessSsidFirewallL7FirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessSsidFirewallL7FirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidFirewallL7FirewallRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0 update network wireless ssid firewall l7 firewall rules params body rules items0
swagger:model UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0
*/
type UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0 struct {

	// 'Deny' traffic specified by this rule
	// Enum: [deny]
	Policy string `json:"policy,omitempty"`

	// Type of the L7 firewall rule. One of: 'application', 'applicationCategory', 'host', 'port', 'ipRange'
	// Enum: [application applicationCategory host ipRange port]
	Type string `json:"type,omitempty"`

	// The value of what needs to get blocked. Format of the value varies depending on type of the firewall rule selected.
	Value string `json:"value,omitempty"`
}

// Validate validates this update network wireless ssid firewall l7 firewall rules params body rules items0
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypePolicyPropEnum = append(updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypePolicyPropEnum, v)
	}
}

const (

	// UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0PolicyDeny captures enum value "deny"
	UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0PolicyDeny string = "deny"
)

// prop value enum
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) validatePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) validatePolicy(formats strfmt.Registry) error {
	if swag.IsZero(o.Policy) { // not required
		return nil
	}

	// value enum
	if err := o.validatePolicyEnum("policy", "body", o.Policy); err != nil {
		return err
	}

	return nil
}

var updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["application","applicationCategory","host","ipRange","port"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeTypePropEnum = append(updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeTypePropEnum, v)
	}
}

const (

	// UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeApplication captures enum value "application"
	UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeApplication string = "application"

	// UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeApplicationCategory captures enum value "applicationCategory"
	UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeApplicationCategory string = "applicationCategory"

	// UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeHost captures enum value "host"
	UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeHost string = "host"

	// UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeIPRange captures enum value "ipRange"
	UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeIPRange string = "ipRange"

	// UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypePort captures enum value "port"
	UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypePort string = "port"
)

// prop value enum
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network wireless ssid firewall l7 firewall rules params body rules items0 based on context it is used
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessSsidFirewallL7FirewallRulesParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
