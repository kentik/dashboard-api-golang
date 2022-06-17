// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// UpdateNetworkApplianceFirewallInboundFirewallRulesReader is a Reader for the UpdateNetworkApplianceFirewallInboundFirewallRules structure.
type UpdateNetworkApplianceFirewallInboundFirewallRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceFirewallInboundFirewallRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceFirewallInboundFirewallRulesOK creates a UpdateNetworkApplianceFirewallInboundFirewallRulesOK with default headers values
func NewUpdateNetworkApplianceFirewallInboundFirewallRulesOK() *UpdateNetworkApplianceFirewallInboundFirewallRulesOK {
	return &UpdateNetworkApplianceFirewallInboundFirewallRulesOK{}
}

/* UpdateNetworkApplianceFirewallInboundFirewallRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceFirewallInboundFirewallRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance firewall inbound firewall rules o k response has a 2xx status code
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance firewall inbound firewall rules o k response has a 3xx status code
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance firewall inbound firewall rules o k response has a 4xx status code
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance firewall inbound firewall rules o k response has a 5xx status code
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance firewall inbound firewall rules o k response a status code equal to that given
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/inboundFirewallRules][%d] updateNetworkApplianceFirewallInboundFirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/inboundFirewallRules][%d] updateNetworkApplianceFirewallInboundFirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkApplianceFirewallInboundFirewallRulesBody update network appliance firewall inbound firewall rules body
// Example: {"rules":[{"comment":"Allow TCP traffic to subnet with HTTP servers.","destCidr":"192.168.1.0/24","destPort":"443","policy":"allow","protocol":"tcp","srcCidr":"Any","srcPort":"Any","syslogEnabled":false}]}
swagger:model UpdateNetworkApplianceFirewallInboundFirewallRulesBody
*/
type UpdateNetworkApplianceFirewallInboundFirewallRulesBody struct {

	// An ordered array of the firewall rules (not including the default rule)
	Rules []*UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0 `json:"rules"`

	// Log the special default rule (boolean value - enable only if you've configured a syslog server) (optional)
	SyslogDefaultRule bool `json:"syslogDefaultRule,omitempty"`
}

// Validate validates this update network appliance firewall inbound firewall rules body
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesBody) validateRules(formats strfmt.Registry) error {
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
					return ve.ValidateName("updateNetworkApplianceFirewallInboundFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallInboundFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance firewall inbound firewall rules body based on the context it is used
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {
			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceFirewallInboundFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallInboundFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallInboundFirewallRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0 update network appliance firewall inbound firewall rules params body rules items0
swagger:model UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0
*/
type UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0 struct {

	// Description of the rule (optional)
	Comment string `json:"comment,omitempty"`

	// Comma-separated list of destination IP address(es) (in IP or CIDR notation), fully-qualified domain names (FQDN) or 'any'
	// Required: true
	DestCidr *string `json:"destCidr"`

	// Comma-separated list of destination port(s) (integer in the range 1-65535), or 'any'
	DestPort string `json:"destPort,omitempty"`

	// 'allow' or 'deny' traffic specified by this rule
	// Required: true
	// Enum: [allow deny]
	Policy *string `json:"policy"`

	// The type of protocol (must be 'tcp', 'udp', 'icmp', 'icmp6' or 'any')
	// Required: true
	// Enum: [tcp udp icmp icmp6 any]
	Protocol *string `json:"protocol"`

	// Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	// Required: true
	SrcCidr *string `json:"srcCidr"`

	// Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SrcPort string `json:"srcPort,omitempty"`

	// Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
	SyslogEnabled bool `json:"syslogEnabled,omitempty"`
}

// Validate validates this update network appliance firewall inbound firewall rules params body rules items0
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDestCidr(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePolicy(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSrcCidr(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) validateDestCidr(formats strfmt.Registry) error {

	if err := validate.Required("destCidr", "body", o.DestCidr); err != nil {
		return err
	}

	return nil
}

var updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum = append(updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0PolicyAllow captures enum value "allow"
	UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0PolicyAllow string = "allow"

	// UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0PolicyDeny captures enum value "deny"
	UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0PolicyDeny string = "deny"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) validatePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", o.Policy); err != nil {
		return err
	}

	// value enum
	if err := o.validatePolicyEnum("policy", "body", *o.Policy); err != nil {
		return err
	}

	return nil
}

var updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp","icmp","icmp6","any"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum = append(updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolTCP captures enum value "tcp"
	UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolTCP string = "tcp"

	// UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolUDP captures enum value "udp"
	UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolUDP string = "udp"

	// UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolIcmp captures enum value "icmp"
	UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolIcmp string = "icmp"

	// UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolIcmp6 captures enum value "icmp6"
	UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolIcmp6 string = "icmp6"

	// UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolAny captures enum value "any"
	UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0ProtocolAny string = "any"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", o.Protocol); err != nil {
		return err
	}

	// value enum
	if err := o.validateProtocolEnum("protocol", "body", *o.Protocol); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) validateSrcCidr(formats strfmt.Registry) error {

	if err := validate.Required("srcCidr", "body", o.SrcCidr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance firewall inbound firewall rules params body rules items0 based on context it is used
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallInboundFirewallRulesParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
