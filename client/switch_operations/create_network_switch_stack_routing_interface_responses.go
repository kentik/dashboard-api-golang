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

// CreateNetworkSwitchStackRoutingInterfaceReader is a Reader for the CreateNetworkSwitchStackRoutingInterface structure.
type CreateNetworkSwitchStackRoutingInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchStackRoutingInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchStackRoutingInterfaceCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces] createNetworkSwitchStackRoutingInterface", response, response.Code())
	}
}

// NewCreateNetworkSwitchStackRoutingInterfaceCreated creates a CreateNetworkSwitchStackRoutingInterfaceCreated with default headers values
func NewCreateNetworkSwitchStackRoutingInterfaceCreated() *CreateNetworkSwitchStackRoutingInterfaceCreated {
	return &CreateNetworkSwitchStackRoutingInterfaceCreated{}
}

/*
CreateNetworkSwitchStackRoutingInterfaceCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchStackRoutingInterfaceCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network switch stack routing interface created response has a 2xx status code
func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network switch stack routing interface created response has a 3xx status code
func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network switch stack routing interface created response has a 4xx status code
func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network switch stack routing interface created response has a 5xx status code
func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network switch stack routing interface created response a status code equal to that given
func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network switch stack routing interface created response
func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) Code() int {
	return 201
}

func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces][%d] createNetworkSwitchStackRoutingInterfaceCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces][%d] createNetworkSwitchStackRoutingInterfaceCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchStackRoutingInterfaceCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkSwitchStackRoutingInterfaceBody create network switch stack routing interface body
// Example: {"defaultGateway":"192.168.1.1","interfaceIp":"192.168.1.2","ipv6":{"address":"1:2:3:4::1","assignmentMode":"static","gateway":"1:2:3:4::2","prefix":"1:2:3:4::/48"},"multicastRouting":"disabled","name":"L3 interface","ospfSettings":{"area":"0","cost":1,"isPassiveEnabled":true},"ospfV3":{"area":"1","cost":2,"isPassiveEnabled":true},"subnet":"192.168.1.0/24","vlanId":100}
swagger:model CreateNetworkSwitchStackRoutingInterfaceBody
*/
type CreateNetworkSwitchStackRoutingInterfaceBody struct {

	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route. This IP address must exist in a subnet with a routed interface.
	DefaultGateway string `json:"defaultGateway,omitempty"`

	// The IP address this switch stack will use for layer 3 routing on this VLAN or subnet. This cannot be the same as the switch's management IP.
	InterfaceIP string `json:"interfaceIp,omitempty"`

	// ipv6
	IPV6 *CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6 `json:"ipv6,omitempty"`

	// Enable multicast support if, multicast routing between VLANs is required. Options are, 'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	// Enum: [IGMP snooping querier disabled enabled]
	MulticastRouting string `json:"multicastRouting,omitempty"`

	// A friendly name or description for the interface or VLAN.
	// Required: true
	Name *string `json:"name"`

	// ospf settings
	OspfSettings *CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings `json:"ospfSettings,omitempty"`

	// The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet string `json:"subnet,omitempty"`

	// The VLAN this routed interface is on. VLAN must be between 1 and 4094.
	// Required: true
	VlanID *int64 `json:"vlanId"`
}

// Validate validates this create network switch stack routing interface body
func (o *CreateNetworkSwitchStackRoutingInterfaceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMulticastRouting(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOspfSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVlanID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchStackRoutingInterfaceBody) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

var createNetworkSwitchStackRoutingInterfaceBodyTypeMulticastRoutingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IGMP snooping querier","disabled","enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkSwitchStackRoutingInterfaceBodyTypeMulticastRoutingPropEnum = append(createNetworkSwitchStackRoutingInterfaceBodyTypeMulticastRoutingPropEnum, v)
	}
}

const (

	// CreateNetworkSwitchStackRoutingInterfaceBodyMulticastRoutingIGMPSnoopingQuerier captures enum value "IGMP snooping querier"
	CreateNetworkSwitchStackRoutingInterfaceBodyMulticastRoutingIGMPSnoopingQuerier string = "IGMP snooping querier"

	// CreateNetworkSwitchStackRoutingInterfaceBodyMulticastRoutingDisabled captures enum value "disabled"
	CreateNetworkSwitchStackRoutingInterfaceBodyMulticastRoutingDisabled string = "disabled"

	// CreateNetworkSwitchStackRoutingInterfaceBodyMulticastRoutingEnabled captures enum value "enabled"
	CreateNetworkSwitchStackRoutingInterfaceBodyMulticastRoutingEnabled string = "enabled"
)

// prop value enum
func (o *CreateNetworkSwitchStackRoutingInterfaceBody) validateMulticastRoutingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkSwitchStackRoutingInterfaceBodyTypeMulticastRoutingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkSwitchStackRoutingInterfaceBody) validateMulticastRouting(formats strfmt.Registry) error {
	if swag.IsZero(o.MulticastRouting) { // not required
		return nil
	}

	// value enum
	if err := o.validateMulticastRoutingEnum("createNetworkSwitchStackRoutingInterface"+"."+"multicastRouting", "body", o.MulticastRouting); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchStackRoutingInterfaceBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchStackRoutingInterface"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchStackRoutingInterfaceBody) validateOspfSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.OspfSettings) { // not required
		return nil
	}

	if o.OspfSettings != nil {
		if err := o.OspfSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ospfSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ospfSettings")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchStackRoutingInterfaceBody) validateVlanID(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchStackRoutingInterface"+"."+"vlanId", "body", o.VlanID); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network switch stack routing interface body based on the context it is used
func (o *CreateNetworkSwitchStackRoutingInterfaceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOspfSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchStackRoutingInterfaceBody) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {

		if swag.IsZero(o.IPV6) { // not required
			return nil
		}

		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchStackRoutingInterfaceBody) contextValidateOspfSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.OspfSettings != nil {

		if swag.IsZero(o.OspfSettings) { // not required
			return nil
		}

		if err := o.OspfSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ospfSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchStackRoutingInterface" + "." + "ospfSettings")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingInterfaceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingInterfaceBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchStackRoutingInterfaceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6 The IPv6 settings of the interface.
swagger:model CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6
*/
type CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6 struct {

	// The IPv6 address of the interface. Required if assignmentMode is 'static'. Must not be included if assignmentMode is 'eui-64'.
	Address string `json:"address,omitempty"`

	// The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	AssignmentMode string `json:"assignmentMode,omitempty"`

	// The IPv6 default gateway of the interface. Required if prefix is defined and this is the first interface with IPv6 configured for the stack.
	Gateway string `json:"gateway,omitempty"`

	// The IPv6 prefix of the interface. Required if IPv6 object is included.
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this create network switch stack routing interface params body IP v6
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch stack routing interface params body IP v6 based on context it is used
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchStackRoutingInterfaceParamsBodyIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings The OSPF routing settings of the interface.
swagger:model CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings
*/
type CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings struct {

	// The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an existing OSPF area. Defaults to 'disabled'.
	Area string `json:"area,omitempty"`

	// The path cost for this interface. Defaults to 1, but can be increased up to 65535 to give lower priority.
	Cost int64 `json:"cost,omitempty"`

	// When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled bool `json:"isPassiveEnabled,omitempty"`
}

// Validate validates this create network switch stack routing interface params body ospf settings
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch stack routing interface params body ospf settings based on context it is used
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchStackRoutingInterfaceParamsBodyOspfSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
