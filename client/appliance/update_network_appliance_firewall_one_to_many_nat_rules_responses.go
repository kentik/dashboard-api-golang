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

// UpdateNetworkApplianceFirewallOneToManyNatRulesReader is a Reader for the UpdateNetworkApplianceFirewallOneToManyNatRules structure.
type UpdateNetworkApplianceFirewallOneToManyNatRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceFirewallOneToManyNatRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceFirewallOneToManyNatRulesOK creates a UpdateNetworkApplianceFirewallOneToManyNatRulesOK with default headers values
func NewUpdateNetworkApplianceFirewallOneToManyNatRulesOK() *UpdateNetworkApplianceFirewallOneToManyNatRulesOK {
	return &UpdateNetworkApplianceFirewallOneToManyNatRulesOK{}
}

/* UpdateNetworkApplianceFirewallOneToManyNatRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceFirewallOneToManyNatRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance firewall one to many nat rules o k response has a 2xx status code
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance firewall one to many nat rules o k response has a 3xx status code
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance firewall one to many nat rules o k response has a 4xx status code
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance firewall one to many nat rules o k response has a 5xx status code
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance firewall one to many nat rules o k response a status code equal to that given
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/oneToManyNatRules][%d] updateNetworkApplianceFirewallOneToManyNatRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/oneToManyNatRules][%d] updateNetworkApplianceFirewallOneToManyNatRulesOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkApplianceFirewallOneToManyNatRulesBody update network appliance firewall one to many nat rules body
// Example: {"rules":[{"portRules":[{"allowedIps":["any"],"localIp":"192.168.128.1","localPort":"443","name":"Rule 1","protocol":"tcp","publicPort":"9443"},{"allowedIps":["10.82.110.0/24","10.82.111.0/24"],"localIp":"192.168.128.1","localPort":"80","name":"Rule 2","protocol":"tcp","publicPort":"8080"}],"publicIp":"146.11.11.13","uplink":"internet1"}]}
swagger:model UpdateNetworkApplianceFirewallOneToManyNatRulesBody
*/
type UpdateNetworkApplianceFirewallOneToManyNatRulesBody struct {

	// An array of 1:Many nat rules
	// Required: true
	Rules []*UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0 `json:"rules"`
}

// Validate validates this update network appliance firewall one to many nat rules body
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesBody) validateRules(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkApplianceFirewallOneToManyNatRules"+"."+"rules", "body", o.Rules); err != nil {
		return err
	}

	for i := 0; i < len(o.Rules); i++ {
		if swag.IsZero(o.Rules[i]) { // not required
			continue
		}

		if o.Rules[i] != nil {
			if err := o.Rules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceFirewallOneToManyNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallOneToManyNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance firewall one to many nat rules body based on the context it is used
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesBody) contextValidateRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Rules); i++ {

		if o.Rules[i] != nil {
			if err := o.Rules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceFirewallOneToManyNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceFirewallOneToManyNatRules" + "." + "rules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallOneToManyNatRulesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0 update network appliance firewall one to many nat rules params body rules items0
swagger:model UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0
*/
type UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0 struct {

	// An array of associated forwarding rules
	// Required: true
	PortRules []*UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0 `json:"portRules"`

	// The IP address that will be used to access the internal resource from the WAN
	// Required: true
	PublicIP *string `json:"publicIp"`

	// The physical WAN interface on which the traffic will arrive ('internet1' or, if available, 'internet2')
	// Required: true
	// Enum: [internet1 internet2]
	Uplink *string `json:"uplink"`
}

// Validate validates this update network appliance firewall one to many nat rules params body rules items0
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePortRules(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePublicIP(formats); err != nil {
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

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) validatePortRules(formats strfmt.Registry) error {

	if err := validate.Required("portRules", "body", o.PortRules); err != nil {
		return err
	}

	for i := 0; i < len(o.PortRules); i++ {
		if swag.IsZero(o.PortRules[i]) { // not required
			continue
		}

		if o.PortRules[i] != nil {
			if err := o.PortRules[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("portRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) validatePublicIP(formats strfmt.Registry) error {

	if err := validate.Required("publicIp", "body", o.PublicIP); err != nil {
		return err
	}

	return nil
}

var updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0TypeUplinkPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["internet1","internet2"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0TypeUplinkPropEnum = append(updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0TypeUplinkPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0UplinkInternet1 captures enum value "internet1"
	UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0UplinkInternet1 string = "internet1"

	// UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0UplinkInternet2 captures enum value "internet2"
	UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0UplinkInternet2 string = "internet2"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) validateUplinkEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0TypeUplinkPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) validateUplink(formats strfmt.Registry) error {

	if err := validate.Required("uplink", "body", o.Uplink); err != nil {
		return err
	}

	// value enum
	if err := o.validateUplinkEnum("uplink", "body", *o.Uplink); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update network appliance firewall one to many nat rules params body rules items0 based on the context it is used
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePortRules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) contextValidatePortRules(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PortRules); i++ {

		if o.PortRules[i] != nil {
			if err := o.PortRules[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("portRules" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("portRules" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0 update network appliance firewall one to many nat rules params body rules items0 port rules items0
swagger:model UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0
*/
type UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0 struct {

	// Remote IP addresses or ranges that are permitted to access the internal resource via this port forwarding rule, or 'any'
	AllowedIps []string `json:"allowedIps"`

	// Local IP address to which traffic will be forwarded
	LocalIP string `json:"localIp,omitempty"`

	// Destination port of the forwarded traffic that will be sent from the MX to the specified host on the LAN. If you simply wish to forward the traffic without translating the port, this should be the same as the Public port
	LocalPort string `json:"localPort,omitempty"`

	// A description of the rule
	Name string `json:"name,omitempty"`

	// 'tcp' or 'udp'
	// Enum: [tcp udp]
	Protocol string `json:"protocol,omitempty"`

	// Destination port of the traffic that is arriving on the WAN
	PublicPort string `json:"publicPort,omitempty"`
}

// Validate validates this update network appliance firewall one to many nat rules params body rules items0 port rules items0
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProtocol(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0TypeProtocolPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","udp"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0TypeProtocolPropEnum = append(updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0TypeProtocolPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0ProtocolTCP captures enum value "tcp"
	UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0ProtocolTCP string = "tcp"

	// UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0ProtocolUDP captures enum value "udp"
	UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0ProtocolUDP string = "udp"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0) validateProtocolEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0TypeProtocolPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0) validateProtocol(formats strfmt.Registry) error {
	if swag.IsZero(o.Protocol) { // not required
		return nil
	}

	// value enum
	if err := o.validateProtocolEnum("protocol", "body", o.Protocol); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance firewall one to many nat rules params body rules items0 port rules items0 based on context it is used
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallOneToManyNatRulesParamsBodyRulesItems0PortRulesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}