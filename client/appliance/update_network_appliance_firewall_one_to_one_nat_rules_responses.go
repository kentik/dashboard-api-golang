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

// UpdateNetworkApplianceFirewallOneToOneNatRulesReader is a Reader for the UpdateNetworkApplianceFirewallOneToOneNatRules structure.
type UpdateNetworkApplianceFirewallOneToOneNatRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceFirewallOneToOneNatRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/appliance/firewall/oneToOneNatRules] updateNetworkApplianceFirewallOneToOneNatRules", response, response.Code())
	}
}

// NewUpdateNetworkApplianceFirewallOneToOneNatRulesOK creates a UpdateNetworkApplianceFirewallOneToOneNatRulesOK with default headers values
func NewUpdateNetworkApplianceFirewallOneToOneNatRulesOK() *UpdateNetworkApplianceFirewallOneToOneNatRulesOK {
	return &UpdateNetworkApplianceFirewallOneToOneNatRulesOK{}
}

/*
UpdateNetworkApplianceFirewallOneToOneNatRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceFirewallOneToOneNatRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance firewall one to one nat rules o k response has a 2xx status code
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance firewall one to one nat rules o k response has a 3xx status code
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance firewall one to one nat rules o k response has a 4xx status code
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance firewall one to one nat rules o k response has a 5xx status code
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance firewall one to one nat rules o k response a status code equal to that given
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network appliance firewall one to one nat rules o k response
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) Code() int {
	return 200
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/oneToOneNatRules][%d] updateNetworkApplianceFirewallOneToOneNatRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/oneToOneNatRules][%d] updateNetworkApplianceFirewallOneToOneNatRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkApplianceFirewallOneToOneNatRulesBody update network appliance firewall one to one nat rules body
// Example: {"rules":[{"allowedInbound":[{"allowedIps":["10.82.112.0/24","10.82.0.0/16"],"destinationPorts":["80"],"protocol":"tcp"},{"allowedIps":["10.81.110.5","10.81.0.0/16"],"destinationPorts":["8080"],"protocol":"udp"}],"lanIp":"192.168.128.22","name":"Service behind NAT","publicIp":"146.12.3.33","uplink":"internet1"}]}
swagger:model UpdateNetworkApplianceFirewallOneToOneNatRulesBody
*/
type UpdateNetworkApplianceFirewallOneToOneNatRulesBody struct {

	// An array of 1:1 nat rules
	// Required: true
	Rules []*UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0 `json:"rules"`
}

// Validate validates this update network appliance firewall one to one nat rules body
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesBody) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkApplianceFirewallOneToOneNatRules"+"."+"rules", "body", o.Rules); err != nil {
		return err
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceFirewallOneToOneNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallOneToOneNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance firewall one to one nat rules body based on the context it is used
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {

			if swag.IsZero(o.Rules[i]) { // not required
				return nil
			}

			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceFirewallOneToOneNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallOneToOneNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallOneToOneNatRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0 update network appliance firewall one to one nat rules params body rules items0
swagger:model UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0
*/
type UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0 struct {

	// The ports this mapping will provide access on, and the remote IPs that will be allowed access to the resource
	AllowedInbound []*UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0 `json:"allowedInbound"`

	// The IP address of the server or device that hosts the internal resource that you wish to make available on the WAN
	// Required: true
	LanIP *string `json:"lanIp"`

	// A descriptive name for the rule
	Name string `json:"name,omitempty"`

	// The IP address that will be used to access the internal resource from the WAN
	PublicIP string `json:"publicIp,omitempty"`

	// The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
	// Enum: [internet1 internet2]
	Uplink string `json:"uplink,omitempty"`
}

// Validate validates this update network appliance firewall one to one nat rules params body rules items0
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAllowedInbound(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLanIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUplink(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) validateAllowedInbound(formats strfmt.Registry) error {
	if swag.IsZero(o.AllowedInbound) { // not required
		return nil
	}

	for i := 0; i < len(o.AllowedInbound); i++ {
		if swag.IsZero(o.AllowedInbound[i]) { // not required
			continue
		}

		if o.AllowedInbound[i] != nil {
			if err := o.AllowedInbound[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedInbound" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedInbound" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) validateLanIP(formats strfmt.Registry) error {

	if err := validate.Required("lanIp", "body", o.LanIP); err != nil {
		return err
	}

	return nil
}

var updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0TypeUplinkPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internet1","internet2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0TypeUplinkPropEnum = append(updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0TypeUplinkPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0UplinkInternet1 captures enum value "internet1"
	UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0UplinkInternet1 string = "internet1"

	// UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0UplinkInternet2 captures enum value "internet2"
	UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0UplinkInternet2 string = "internet2"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) validateUplinkEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0TypeUplinkPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) validateUplink(formats strfmt.Registry) error {
	if swag.IsZero(o.Uplink) { // not required
		return nil
	}

	// value enum
	if err := o.validateUplinkEnum("uplink", "body", o.Uplink); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update network appliance firewall one to one nat rules params body rules items0 based on the context it is used
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAllowedInbound(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) contextValidateAllowedInbound(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.AllowedInbound); i++ {

		if o.AllowedInbound[i] != nil {

			if swag.IsZero(o.AllowedInbound[i]) { // not required
				return nil
			}

			if err := o.AllowedInbound[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("allowedInbound" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("allowedInbound" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0 update network appliance firewall one to one nat rules params body rules items0 allowed inbound items0
swagger:model UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0
*/
type UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0 struct {

	// An array of ranges of WAN IP addresses that are allowed to make inbound connections on the specified ports or port ranges, or 'any'
	AllowedIps []string `json:"allowedIps"`

	// An array of ports or port ranges that will be forwarded to the host on the LAN
	DestinationPorts []string `json:"destinationPorts"`

	// Either of the following: 'tcp', 'udp', 'icmp-ping' or 'any'
	// Enum: [any icmp-ping tcp udp]
	Protocol string `json:"protocol,omitempty"`
}

// Validate validates this update network appliance firewall one to one nat rules params body rules items0 allowed inbound items0
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0TypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["any","icmp-ping","tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0TypeProtocolPropEnum = append(updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0TypeProtocolPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolAny captures enum value "any"
	UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolAny string = "any"

	// UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolIcmpDashPing captures enum value "icmp-ping"
	UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolIcmpDashPing string = "icmp-ping"

	// UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolTCP captures enum value "tcp"
	UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolTCP string = "tcp"

	// UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolUDP captures enum value "udp"
	UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0ProtocolUDP string = "udp"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0TypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(o.Protocol) { // not required
		return nil
	}

	// value enum
	if err := o.validateProtocolEnum("protocol", "body", o.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance firewall one to one nat rules params body rules items0 allowed inbound items0 based on context it is used
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallOneToOneNatRulesParamsBodyRulesItems0AllowedInboundItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
