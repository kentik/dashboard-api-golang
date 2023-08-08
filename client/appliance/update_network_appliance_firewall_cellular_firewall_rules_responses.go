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

// UpdateNetworkApplianceFirewallCellularFirewallRulesReader is a Reader for the UpdateNetworkApplianceFirewallCellularFirewallRules structure.
type UpdateNetworkApplianceFirewallCellularFirewallRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceFirewallCellularFirewallRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/appliance/firewall/cellularFirewallRules] updateNetworkApplianceFirewallCellularFirewallRules", response, response.Code())
	}
}

// NewUpdateNetworkApplianceFirewallCellularFirewallRulesOK creates a UpdateNetworkApplianceFirewallCellularFirewallRulesOK with default headers values
func NewUpdateNetworkApplianceFirewallCellularFirewallRulesOK() *UpdateNetworkApplianceFirewallCellularFirewallRulesOK {
	return &UpdateNetworkApplianceFirewallCellularFirewallRulesOK{}
}

/*
UpdateNetworkApplianceFirewallCellularFirewallRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceFirewallCellularFirewallRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance firewall cellular firewall rules o k response has a 2xx status code
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance firewall cellular firewall rules o k response has a 3xx status code
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance firewall cellular firewall rules o k response has a 4xx status code
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance firewall cellular firewall rules o k response has a 5xx status code
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance firewall cellular firewall rules o k response a status code equal to that given
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network appliance firewall cellular firewall rules o k response
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) Code() int {
	return 200
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/cellularFirewallRules][%d] updateNetworkApplianceFirewallCellularFirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/cellularFirewallRules][%d] updateNetworkApplianceFirewallCellularFirewallRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkApplianceFirewallCellularFirewallRulesBody update network appliance firewall cellular firewall rules body
// Example: {"rules":[{"comment":"Allow TCP traffic to subnet with HTTP servers.","destCidr":"192.168.1.0/24","destPort":"443","policy":"allow","protocol":"tcp","srcCidr":"Any","srcPort":"Any","syslogEnabled":false}]}
swagger:model UpdateNetworkApplianceFirewallCellularFirewallRulesBody
*/
type UpdateNetworkApplianceFirewallCellularFirewallRulesBody struct {

	// An ordered array of the firewall rules (not including the default rule)
	Rules []*UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0 `json:"rules"`
}

// Validate validates this update network appliance firewall cellular firewall rules body
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesBody) validateRules(formats strfmt.Registry) error {
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
					return ve.ValidateName("updateNetworkApplianceFirewallCellularFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallCellularFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance firewall cellular firewall rules body based on the context it is used
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {

			if swag.IsZero(o.Rules[i]) { // not required
				return nil
			}

			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceFirewallCellularFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallCellularFirewallRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallCellularFirewallRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0 update network appliance firewall cellular firewall rules params body rules items0
swagger:model UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0
*/
type UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0 struct {

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
	// Enum: [any icmp icmp6 tcp udp]
	Protocol *string `json:"protocol"`

	// Comma-separated list of source IP address(es) (in IP or CIDR notation), or 'any' (note: FQDN not supported for source addresses)
	// Required: true
	SrcCidr *string `json:"srcCidr"`

	// Comma-separated list of source port(s) (integer in the range 1-65535), or 'any'
	SrcPort string `json:"srcPort,omitempty"`

	// Log this rule to syslog (true or false, boolean value) - only applicable if a syslog has been configured (optional)
	SyslogEnabled bool `json:"syslogEnabled,omitempty"`
}

// Validate validates this update network appliance firewall cellular firewall rules params body rules items0
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
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

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) validateDestCidr(formats strfmt.Registry) error {

	if err := validate.Required("destCidr", "body", o.DestCidr); err != nil {
		return err
	}

	return nil
}

var updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["allow","deny"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum = append(updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0PolicyAllow captures enum value "allow"
	UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0PolicyAllow string = "allow"

	// UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0PolicyDeny captures enum value "deny"
	UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0PolicyDeny string = "deny"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) validatePolicyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypePolicyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) validatePolicy(formats strfmt.Registry) error {

	if err := validate.Required("policy", "body", o.Policy); err != nil {
		return err
	}

	// value enum
	if err := o.validatePolicyEnum("policy", "body", *o.Policy); err != nil {
		return err
	}

	return nil
}

var updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","icmp","icmp6","tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum = append(updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolAny captures enum value "any"
	UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolAny string = "any"

	// UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolIcmp captures enum value "icmp"
	UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolIcmp string = "icmp"

	// UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolIcmp6 captures enum value "icmp6"
	UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolIcmp6 string = "icmp6"

	// UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolTCP captures enum value "tcp"
	UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolTCP string = "tcp"

	// UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolUDP captures enum value "udp"
	UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0ProtocolUDP string = "udp"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0TypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) validateProtocol(formats strfmt.Registry) error {

	if err := validate.Required("protocol", "body", o.Protocol); err != nil {
		return err
	}

	// value enum
	if err := o.validateProtocolEnum("protocol", "body", *o.Protocol); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) validateSrcCidr(formats strfmt.Registry) error {

	if err := validate.Required("srcCidr", "body", o.SrcCidr); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance firewall cellular firewall rules params body rules items0 based on context it is used
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallCellularFirewallRulesParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
