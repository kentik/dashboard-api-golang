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

// UpdateDeviceSwitchRoutingInterfaceReader is a Reader for the UpdateDeviceSwitchRoutingInterface structure.
type UpdateDeviceSwitchRoutingInterfaceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceSwitchRoutingInterfaceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceSwitchRoutingInterfaceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}/switch/routing/interfaces/{interfaceId}] updateDeviceSwitchRoutingInterface", response, response.Code())
	}
}

// NewUpdateDeviceSwitchRoutingInterfaceOK creates a UpdateDeviceSwitchRoutingInterfaceOK with default headers values
func NewUpdateDeviceSwitchRoutingInterfaceOK() *UpdateDeviceSwitchRoutingInterfaceOK {
	return &UpdateDeviceSwitchRoutingInterfaceOK{}
}

/*
UpdateDeviceSwitchRoutingInterfaceOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceSwitchRoutingInterfaceOK struct {
	Payload *UpdateDeviceSwitchRoutingInterfaceOKBody
}

// IsSuccess returns true when this update device switch routing interface o k response has a 2xx status code
func (o *UpdateDeviceSwitchRoutingInterfaceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device switch routing interface o k response has a 3xx status code
func (o *UpdateDeviceSwitchRoutingInterfaceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device switch routing interface o k response has a 4xx status code
func (o *UpdateDeviceSwitchRoutingInterfaceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device switch routing interface o k response has a 5xx status code
func (o *UpdateDeviceSwitchRoutingInterfaceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device switch routing interface o k response a status code equal to that given
func (o *UpdateDeviceSwitchRoutingInterfaceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device switch routing interface o k response
func (o *UpdateDeviceSwitchRoutingInterfaceOK) Code() int {
	return 200
}

func (o *UpdateDeviceSwitchRoutingInterfaceOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/switch/routing/interfaces/{interfaceId}][%d] updateDeviceSwitchRoutingInterfaceOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceSwitchRoutingInterfaceOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/switch/routing/interfaces/{interfaceId}][%d] updateDeviceSwitchRoutingInterfaceOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceSwitchRoutingInterfaceOK) GetPayload() *UpdateDeviceSwitchRoutingInterfaceOKBody {
	return o.Payload
}

func (o *UpdateDeviceSwitchRoutingInterfaceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateDeviceSwitchRoutingInterfaceOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceBody update device switch routing interface body
// Example: {"defaultGateway":"192.168.1.1","interfaceIp":"192.168.1.2","ipv6":{"address":"1:2:3:4::1","assignmentMode":"static","gateway":"1:2:3:4::2","prefix":"1:2:3:4::/48"},"multicastRouting":"disabled","name":"L3 interface","ospfSettings":{"area":"0","cost":1,"isPassiveEnabled":true},"ospfV3":{"area":"1","cost":2,"isPassiveEnabled":true},"subnet":"192.168.1.0/24","vlanId":100}
swagger:model UpdateDeviceSwitchRoutingInterfaceBody
*/
type UpdateDeviceSwitchRoutingInterfaceBody struct {

	// The next hop for any traffic that isn't going to a directly connected subnet or over a static route.         This IP address must exist in a subnet with a routed interface. Required if this is the first IPv4 interface.
	DefaultGateway string `json:"defaultGateway,omitempty"`

	// The IP address this switch will use for layer 3 routing on this VLAN or subnet. This cannot be the same         as the switch's management IP.
	InterfaceIP string `json:"interfaceIp,omitempty"`

	// ipv6
	IPV6 *UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6 `json:"ipv6,omitempty"`

	// Enable multicast support if, multicast routing between VLANs is required. Options are:         'disabled', 'enabled' or 'IGMP snooping querier'. Default is 'disabled'.
	// Enum: [IGMP snooping querier disabled enabled]
	MulticastRouting string `json:"multicastRouting,omitempty"`

	// A friendly name or description for the interface or VLAN.
	Name string `json:"name,omitempty"`

	// ospf settings
	OspfSettings *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings `json:"ospfSettings,omitempty"`

	// ospf v3
	OspfV3 *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3 `json:"ospfV3,omitempty"`

	// The network that this routed interface is on, in CIDR notation (ex. 10.1.1.0/24).
	Subnet string `json:"subnet,omitempty"`

	// The VLAN this routed interface is on. VLAN must be between 1 and 4094.
	VlanID int64 `json:"vlanId,omitempty"`
}

// Validate validates this update device switch routing interface body
func (o *UpdateDeviceSwitchRoutingInterfaceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMulticastRouting(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOspfSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOspfV3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceBody) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

var updateDeviceSwitchRoutingInterfaceBodyTypeMulticastRoutingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["IGMP snooping querier","disabled","enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateDeviceSwitchRoutingInterfaceBodyTypeMulticastRoutingPropEnum = append(updateDeviceSwitchRoutingInterfaceBodyTypeMulticastRoutingPropEnum, v)
	}
}

const (

	// UpdateDeviceSwitchRoutingInterfaceBodyMulticastRoutingIGMPSnoopingQuerier captures enum value "IGMP snooping querier"
	UpdateDeviceSwitchRoutingInterfaceBodyMulticastRoutingIGMPSnoopingQuerier string = "IGMP snooping querier"

	// UpdateDeviceSwitchRoutingInterfaceBodyMulticastRoutingDisabled captures enum value "disabled"
	UpdateDeviceSwitchRoutingInterfaceBodyMulticastRoutingDisabled string = "disabled"

	// UpdateDeviceSwitchRoutingInterfaceBodyMulticastRoutingEnabled captures enum value "enabled"
	UpdateDeviceSwitchRoutingInterfaceBodyMulticastRoutingEnabled string = "enabled"
)

// prop value enum
func (o *UpdateDeviceSwitchRoutingInterfaceBody) validateMulticastRoutingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateDeviceSwitchRoutingInterfaceBodyTypeMulticastRoutingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceBody) validateMulticastRouting(formats strfmt.Registry) error {
	if swag.IsZero(o.MulticastRouting) { // not required
		return nil
	}

	// value enum
	if err := o.validateMulticastRoutingEnum("updateDeviceSwitchRoutingInterface"+"."+"multicastRouting", "body", o.MulticastRouting); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceBody) validateOspfSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.OspfSettings) { // not required
		return nil
	}

	if o.OspfSettings != nil {
		if err := o.OspfSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceBody) validateOspfV3(formats strfmt.Registry) error {
	if swag.IsZero(o.OspfV3) { // not required
		return nil
	}

	if o.OspfV3 != nil {
		if err := o.OspfV3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfV3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfV3")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update device switch routing interface body based on the context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOspfSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOspfV3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceBody) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {

		if swag.IsZero(o.IPV6) { // not required
			return nil
		}

		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceBody) contextValidateOspfSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.OspfSettings != nil {

		if swag.IsZero(o.OspfSettings) { // not required
			return nil
		}

		if err := o.OspfSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceBody) contextValidateOspfV3(ctx context.Context, formats strfmt.Registry) error {

	if o.OspfV3 != nil {

		if swag.IsZero(o.OspfV3) { // not required
			return nil
		}

		if err := o.OspfV3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfV3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterface" + "." + "ospfV3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceOKBody update device switch routing interface o k body
swagger:model UpdateDeviceSwitchRoutingInterfaceOKBody
*/
type UpdateDeviceSwitchRoutingInterfaceOKBody struct {

	// IPv4 default gateway
	DefaultGateway string `json:"defaultGateway,omitempty"`

	// The id
	InterfaceID string `json:"interfaceId,omitempty"`

	// IPv4 address
	InterfaceIP string `json:"interfaceIp,omitempty"`

	// ipv6
	IPV6 *UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6 `json:"ipv6,omitempty"`

	// Multicast routing status
	MulticastRouting string `json:"multicastRouting,omitempty"`

	// The name
	Name string `json:"name,omitempty"`

	// ospf settings
	OspfSettings *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings `json:"ospfSettings,omitempty"`

	// ospf v3
	OspfV3 *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3 `json:"ospfV3,omitempty"`

	// IPv4 subnet
	Subnet string `json:"subnet,omitempty"`

	// VLAN id
	VlanID int64 `json:"vlanId,omitempty"`
}

// Validate validates this update device switch routing interface o k body
func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOspfSettings(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOspfV3(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) validateOspfSettings(formats strfmt.Registry) error {
	if swag.IsZero(o.OspfSettings) { // not required
		return nil
	}

	if o.OspfSettings != nil {
		if err := o.OspfSettings.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) validateOspfV3(formats strfmt.Registry) error {
	if swag.IsZero(o.OspfV3) { // not required
		return nil
	}

	if o.OspfV3 != nil {
		if err := o.OspfV3.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfV3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfV3")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update device switch routing interface o k body based on the context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOspfSettings(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateOspfV3(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {

		if swag.IsZero(o.IPV6) { // not required
			return nil
		}

		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) contextValidateOspfSettings(ctx context.Context, formats strfmt.Registry) error {

	if o.OspfSettings != nil {

		if swag.IsZero(o.OspfSettings) { // not required
			return nil
		}

		if err := o.OspfSettings.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfSettings")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfSettings")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) contextValidateOspfV3(ctx context.Context, formats strfmt.Registry) error {

	if o.OspfV3 != nil {

		if swag.IsZero(o.OspfV3) { // not required
			return nil
		}

		if err := o.OspfV3.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfV3")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateDeviceSwitchRoutingInterfaceOK" + "." + "ospfV3")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6 IPv6 addressing
swagger:model UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6
*/
type UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6 struct {

	// IPv6 address
	Address string `json:"address,omitempty"`

	// Assignment mode
	AssignmentMode string `json:"assignmentMode,omitempty"`

	// IPv6 gateway
	Gateway string `json:"gateway,omitempty"`

	// IPv6 subnet
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this update device switch routing interface o k body IP v6
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device switch routing interface o k body IP v6 based on context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceOKBodyIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings IPv4 OSPF Settings
swagger:model UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings
*/
type UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings struct {

	// Area id
	Area string `json:"area,omitempty"`

	// OSPF Cost
	Cost int64 `json:"cost,omitempty"`

	// Disable sending Hello packets on this interface's IPv4 area
	IsPassiveEnabled bool `json:"isPassiveEnabled,omitempty"`
}

// Validate validates this update device switch routing interface o k body ospf settings
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device switch routing interface o k body ospf settings based on context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceOKBodyOspfSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3 IPv6 OSPF Settings
swagger:model UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3
*/
type UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3 struct {

	// Area id
	Area string `json:"area,omitempty"`

	// OSPF Cost
	Cost int64 `json:"cost,omitempty"`

	// Disable sending Hello packets on this interface's IPv6 area
	IsPassiveEnabled bool `json:"isPassiveEnabled,omitempty"`
}

// Validate validates this update device switch routing interface o k body ospf v3
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device switch routing interface o k body ospf v3 based on context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceOKBodyOspfV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6 The IPv6 settings of the interface.
swagger:model UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6
*/
type UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6 struct {

	// The IPv6 address of the interface. Required if assignmentMode is 'static'. Must not be included if           assignmentMode is 'eui-64'.
	Address string `json:"address,omitempty"`

	// The IPv6 assignment mode for the interface. Can be either 'eui-64' or 'static'.
	AssignmentMode string `json:"assignmentMode,omitempty"`

	// The IPv6 default gateway of the interface. Required if prefix is defined and this is the first           interface with IPv6 configured for the switch.
	Gateway string `json:"gateway,omitempty"`

	// The IPv6 prefix of the interface. Required if IPv6 object is included.
	Prefix string `json:"prefix,omitempty"`
}

// Validate validates this update device switch routing interface params body IP v6
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device switch routing interface params body IP v6 based on context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceParamsBodyIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings The OSPF routing settings of the interface.
swagger:model UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings
*/
type UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings struct {

	// The OSPF area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPF area. Defaults to 'disabled'.
	Area string `json:"area,omitempty"`

	// The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	Cost int64 `json:"cost,omitempty"`

	// When enabled, OSPF will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled bool `json:"isPassiveEnabled,omitempty"`
}

// Validate validates this update device switch routing interface params body ospf settings
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device switch routing interface params body ospf settings based on context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfSettings
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3 The OSPFv3 routing settings of the interface.
swagger:model UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3
*/
type UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3 struct {

	// The OSPFv3 area to which this interface should belong. Can be either 'disabled' or the identifier of an           existing OSPFv3 area. Defaults to 'disabled'.
	Area string `json:"area,omitempty"`

	// The path cost for this interface. Defaults to 1, but can be increased up to 65535           to give lower priority.
	Cost int64 `json:"cost,omitempty"`

	// When enabled, OSPFv3 will not run on the interface, but the subnet will still be advertised.
	IsPassiveEnabled bool `json:"isPassiveEnabled,omitempty"`
}

// Validate validates this update device switch routing interface params body ospf v3
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device switch routing interface params body ospf v3 based on context it is used
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingInterfaceParamsBodyOspfV3
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
