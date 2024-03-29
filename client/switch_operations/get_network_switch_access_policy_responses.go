// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// GetNetworkSwitchAccessPolicyReader is a Reader for the GetNetworkSwitchAccessPolicy structure.
type GetNetworkSwitchAccessPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchAccessPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchAccessPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}] getNetworkSwitchAccessPolicy", response, response.Code())
	}
}

// NewGetNetworkSwitchAccessPolicyOK creates a GetNetworkSwitchAccessPolicyOK with default headers values
func NewGetNetworkSwitchAccessPolicyOK() *GetNetworkSwitchAccessPolicyOK {
	return &GetNetworkSwitchAccessPolicyOK{}
}

/*
GetNetworkSwitchAccessPolicyOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchAccessPolicyOK struct {
	Payload *GetNetworkSwitchAccessPolicyOKBody
}

// IsSuccess returns true when this get network switch access policy o k response has a 2xx status code
func (o *GetNetworkSwitchAccessPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch access policy o k response has a 3xx status code
func (o *GetNetworkSwitchAccessPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch access policy o k response has a 4xx status code
func (o *GetNetworkSwitchAccessPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch access policy o k response has a 5xx status code
func (o *GetNetworkSwitchAccessPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch access policy o k response a status code equal to that given
func (o *GetNetworkSwitchAccessPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch access policy o k response
func (o *GetNetworkSwitchAccessPolicyOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchAccessPolicyOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}][%d] getNetworkSwitchAccessPolicyOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchAccessPolicyOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/accessPolicies/{accessPolicyNumber}][%d] getNetworkSwitchAccessPolicyOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchAccessPolicyOK) GetPayload() *GetNetworkSwitchAccessPolicyOKBody {
	return o.Payload
}

func (o *GetNetworkSwitchAccessPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkSwitchAccessPolicyOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSwitchAccessPolicyOKBody get network switch access policy o k body
swagger:model GetNetworkSwitchAccessPolicyOKBody
*/
type GetNetworkSwitchAccessPolicyOKBody struct {

	// Access Type of the policy. Automatically 'Hybrid authentication' when hostMode is 'Multi-Domain'.
	// Enum: [802.1x Hybrid authentication MAC authentication bypass]
	AccessPolicyType string `json:"accessPolicyType,omitempty"`

	// dot1x
	Dot1x *GetNetworkSwitchAccessPolicyOKBodyDot1x `json:"dot1x,omitempty"`

	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	GuestPortBouncing bool `json:"guestPortBouncing,omitempty"`

	// ID for the guest VLAN allow unauthorized devices access to limited network resources
	GuestVlanID int64 `json:"guestVlanId,omitempty"`

	// Choose the Host Mode for the access policy.
	// Enum: [Multi-Auth Multi-Domain Multi-Host Single-Host]
	HostMode string `json:"hostMode,omitempty"`

	// Enabling this option will make switches execute 802.1X and MAC-bypass authentication simultaneously so that clients authenticate faster. Only required when accessPolicyType is 'Hybrid Authentication.
	IncreaseAccessSpeed bool `json:"increaseAccessSpeed,omitempty"`

	// Name of the access policy
	Name string `json:"name,omitempty"`

	// radius
	Radius *GetNetworkSwitchAccessPolicyOKBodyRadius `json:"radius,omitempty"`

	// Enable to send start, interim-update and stop messages to a configured RADIUS accounting server for tracking connected clients
	RadiusAccountingEnabled bool `json:"radiusAccountingEnabled,omitempty"`

	// List of RADIUS accounting servers to require connecting devices to authenticate against before granting network access
	RadiusAccountingServers []*GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0 `json:"radiusAccountingServers"`

	// Change of authentication for RADIUS re-authentication and disconnection
	RadiusCoaSupportEnabled bool `json:"radiusCoaSupportEnabled,omitempty"`

	// Acceptable values are `""` for None, or `"11"` for Group Policies ACL
	RadiusGroupAttribute string `json:"radiusGroupAttribute,omitempty"`

	// List of RADIUS servers to require connecting devices to authenticate against before granting network access
	RadiusServers []*GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0 `json:"radiusServers"`

	// If enabled, Meraki devices will periodically send access-request messages to these RADIUS servers
	RadiusTestingEnabled bool `json:"radiusTestingEnabled,omitempty"`

	// Enable to restrict access for clients to a response_objectific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenEnabled bool `json:"urlRedirectWalledGardenEnabled,omitempty"`

	// IP address ranges, in CIDR notation, to restrict access for clients to a specific set of IP addresses or hostnames prior to authentication
	URLRedirectWalledGardenRanges []string `json:"urlRedirectWalledGardenRanges"`

	// CDP/LLDP capable voice clients will be able to use this VLAN. Automatically true when hostMode is 'Multi-Domain'.
	VoiceVlanClients bool `json:"voiceVlanClients,omitempty"`
}

// Validate validates this get network switch access policy o k body
func (o *GetNetworkSwitchAccessPolicyOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessPolicyType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDot1x(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateHostMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadius(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusAccountingServers(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRadiusServers(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkSwitchAccessPolicyOKBodyTypeAccessPolicyTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["802.1x","Hybrid authentication","MAC authentication bypass"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkSwitchAccessPolicyOKBodyTypeAccessPolicyTypePropEnum = append(getNetworkSwitchAccessPolicyOKBodyTypeAccessPolicyTypePropEnum, v)
	}
}

const (

	// GetNetworkSwitchAccessPolicyOKBodyAccessPolicyTypeNr802Dot1x captures enum value "802.1x"
	GetNetworkSwitchAccessPolicyOKBodyAccessPolicyTypeNr802Dot1x string = "802.1x"

	// GetNetworkSwitchAccessPolicyOKBodyAccessPolicyTypeHybridAuthentication captures enum value "Hybrid authentication"
	GetNetworkSwitchAccessPolicyOKBodyAccessPolicyTypeHybridAuthentication string = "Hybrid authentication"

	// GetNetworkSwitchAccessPolicyOKBodyAccessPolicyTypeMACAuthenticationBypass captures enum value "MAC authentication bypass"
	GetNetworkSwitchAccessPolicyOKBodyAccessPolicyTypeMACAuthenticationBypass string = "MAC authentication bypass"
)

// prop value enum
func (o *GetNetworkSwitchAccessPolicyOKBody) validateAccessPolicyTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkSwitchAccessPolicyOKBodyTypeAccessPolicyTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) validateAccessPolicyType(formats strfmt.Registry) error {
	if swag.IsZero(o.AccessPolicyType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAccessPolicyTypeEnum("getNetworkSwitchAccessPolicyOK"+"."+"accessPolicyType", "body", o.AccessPolicyType); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) validateDot1x(formats strfmt.Registry) error {
	if swag.IsZero(o.Dot1x) { // not required
		return nil
	}

	if o.Dot1x != nil {
		if err := o.Dot1x.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "dot1x")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "dot1x")
			}
			return err
		}
	}

	return nil
}

var getNetworkSwitchAccessPolicyOKBodyTypeHostModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Multi-Auth","Multi-Domain","Multi-Host","Single-Host"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkSwitchAccessPolicyOKBodyTypeHostModePropEnum = append(getNetworkSwitchAccessPolicyOKBodyTypeHostModePropEnum, v)
	}
}

const (

	// GetNetworkSwitchAccessPolicyOKBodyHostModeMultiDashAuth captures enum value "Multi-Auth"
	GetNetworkSwitchAccessPolicyOKBodyHostModeMultiDashAuth string = "Multi-Auth"

	// GetNetworkSwitchAccessPolicyOKBodyHostModeMultiDashDomain captures enum value "Multi-Domain"
	GetNetworkSwitchAccessPolicyOKBodyHostModeMultiDashDomain string = "Multi-Domain"

	// GetNetworkSwitchAccessPolicyOKBodyHostModeMultiDashHost captures enum value "Multi-Host"
	GetNetworkSwitchAccessPolicyOKBodyHostModeMultiDashHost string = "Multi-Host"

	// GetNetworkSwitchAccessPolicyOKBodyHostModeSingleDashHost captures enum value "Single-Host"
	GetNetworkSwitchAccessPolicyOKBodyHostModeSingleDashHost string = "Single-Host"
)

// prop value enum
func (o *GetNetworkSwitchAccessPolicyOKBody) validateHostModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkSwitchAccessPolicyOKBodyTypeHostModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) validateHostMode(formats strfmt.Registry) error {
	if swag.IsZero(o.HostMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateHostModeEnum("getNetworkSwitchAccessPolicyOK"+"."+"hostMode", "body", o.HostMode); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) validateRadius(formats strfmt.Registry) error {
	if swag.IsZero(o.Radius) { // not required
		return nil
	}

	if o.Radius != nil {
		if err := o.Radius.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) validateRadiusAccountingServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusAccountingServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusAccountingServers); i++ {
		if swag.IsZero(o.RadiusAccountingServers[i]) { // not required
			continue
		}

		if o.RadiusAccountingServers[i] != nil {
			if err := o.RadiusAccountingServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) validateRadiusServers(formats strfmt.Registry) error {
	if swag.IsZero(o.RadiusServers) { // not required
		return nil
	}

	for i := 0; i < len(o.RadiusServers); i++ {
		if swag.IsZero(o.RadiusServers[i]) { // not required
			continue
		}

		if o.RadiusServers[i] != nil {
			if err := o.RadiusServers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network switch access policy o k body based on the context it is used
func (o *GetNetworkSwitchAccessPolicyOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDot1x(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadius(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadiusAccountingServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRadiusServers(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) contextValidateDot1x(ctx context.Context, formats strfmt.Registry) error {

	if o.Dot1x != nil {

		if swag.IsZero(o.Dot1x) { // not required
			return nil
		}

		if err := o.Dot1x.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "dot1x")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "dot1x")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) contextValidateRadius(ctx context.Context, formats strfmt.Registry) error {

	if o.Radius != nil {

		if swag.IsZero(o.Radius) { // not required
			return nil
		}

		if err := o.Radius.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) contextValidateRadiusAccountingServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusAccountingServers); i++ {

		if o.RadiusAccountingServers[i] != nil {

			if swag.IsZero(o.RadiusAccountingServers[i]) { // not required
				return nil
			}

			if err := o.RadiusAccountingServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusAccountingServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBody) contextValidateRadiusServers(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RadiusServers); i++ {

		if o.RadiusServers[i] != nil {

			if swag.IsZero(o.RadiusServers[i]) { // not required
				return nil
			}

			if err := o.RadiusServers[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusServers" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radiusServers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchAccessPolicyOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSwitchAccessPolicyOKBodyDot1x 802.1x Settings
swagger:model GetNetworkSwitchAccessPolicyOKBodyDot1x
*/
type GetNetworkSwitchAccessPolicyOKBodyDot1x struct {

	// Supports either 'both' or 'inbound'. Set to 'inbound' to allow unauthorized egress on the switchport. Set to 'both' to control both traffic directions with authorization. Defaults to 'both'
	// Enum: [both inbound]
	ControlDirection string `json:"controlDirection,omitempty"`
}

// Validate validates this get network switch access policy o k body dot1x
func (o *GetNetworkSwitchAccessPolicyOKBodyDot1x) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateControlDirection(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkSwitchAccessPolicyOKBodyDot1xTypeControlDirectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["both","inbound"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkSwitchAccessPolicyOKBodyDot1xTypeControlDirectionPropEnum = append(getNetworkSwitchAccessPolicyOKBodyDot1xTypeControlDirectionPropEnum, v)
	}
}

const (

	// GetNetworkSwitchAccessPolicyOKBodyDot1xControlDirectionBoth captures enum value "both"
	GetNetworkSwitchAccessPolicyOKBodyDot1xControlDirectionBoth string = "both"

	// GetNetworkSwitchAccessPolicyOKBodyDot1xControlDirectionInbound captures enum value "inbound"
	GetNetworkSwitchAccessPolicyOKBodyDot1xControlDirectionInbound string = "inbound"
)

// prop value enum
func (o *GetNetworkSwitchAccessPolicyOKBodyDot1x) validateControlDirectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkSwitchAccessPolicyOKBodyDot1xTypeControlDirectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBodyDot1x) validateControlDirection(formats strfmt.Registry) error {
	if swag.IsZero(o.ControlDirection) { // not required
		return nil
	}

	// value enum
	if err := o.validateControlDirectionEnum("getNetworkSwitchAccessPolicyOK"+"."+"dot1x"+"."+"controlDirection", "body", o.ControlDirection); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network switch access policy o k body dot1x based on context it is used
func (o *GetNetworkSwitchAccessPolicyOKBodyDot1x) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyDot1x) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyDot1x) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchAccessPolicyOKBodyDot1x
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSwitchAccessPolicyOKBodyRadius Object for RADIUS Settings
swagger:model GetNetworkSwitchAccessPolicyOKBodyRadius
*/
type GetNetworkSwitchAccessPolicyOKBodyRadius struct {

	// critical auth
	CriticalAuth *GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth `json:"criticalAuth,omitempty"`

	// VLAN that clients will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	FailedAuthVlanID int64 `json:"failedAuthVlanId,omitempty"`

	// Re-authentication period in seconds. Will be null if hostMode is Multi-Auth
	ReAuthenticationInterval int64 `json:"reAuthenticationInterval,omitempty"`
}

// Validate validates this get network switch access policy o k body radius
func (o *GetNetworkSwitchAccessPolicyOKBodyRadius) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCriticalAuth(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBodyRadius) validateCriticalAuth(formats strfmt.Registry) error {
	if swag.IsZero(o.CriticalAuth) { // not required
		return nil
	}

	if o.CriticalAuth != nil {
		if err := o.CriticalAuth.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius" + "." + "criticalAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius" + "." + "criticalAuth")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network switch access policy o k body radius based on the context it is used
func (o *GetNetworkSwitchAccessPolicyOKBodyRadius) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCriticalAuth(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSwitchAccessPolicyOKBodyRadius) contextValidateCriticalAuth(ctx context.Context, formats strfmt.Registry) error {

	if o.CriticalAuth != nil {

		if swag.IsZero(o.CriticalAuth) { // not required
			return nil
		}

		if err := o.CriticalAuth.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius" + "." + "criticalAuth")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkSwitchAccessPolicyOK" + "." + "radius" + "." + "criticalAuth")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadius) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadius) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchAccessPolicyOKBodyRadius
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0 get network switch access policy o k body radius accounting servers items0
swagger:model GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0
*/
type GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0 struct {

	// Public IP address of the RADIUS accounting server
	Host string `json:"host,omitempty"`

	// UDP port that the RADIUS Accounting server listens on for access requests
	Port int64 `json:"port,omitempty"`
}

// Validate validates this get network switch access policy o k body radius accounting servers items0
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network switch access policy o k body radius accounting servers items0 based on context it is used
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchAccessPolicyOKBodyRadiusAccountingServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth Critical auth settings for when authentication is rejected by the RADIUS server
swagger:model GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth
*/
type GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth struct {

	// VLAN that clients who use data will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	DataVlanID int64 `json:"dataVlanId,omitempty"`

	// Enable to suspend port bounce when RADIUS servers are unreachable
	SuspendPortBounce bool `json:"suspendPortBounce,omitempty"`

	// VLAN that clients who use voice will be placed on when RADIUS authentication fails. Will be null if hostMode is Multi-Auth
	VoiceVlanID int64 `json:"voiceVlanId,omitempty"`
}

// Validate validates this get network switch access policy o k body radius critical auth
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network switch access policy o k body radius critical auth based on context it is used
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchAccessPolicyOKBodyRadiusCriticalAuth
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0 get network switch access policy o k body radius servers items0
swagger:model GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0
*/
type GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0 struct {

	// Public IP address of the RADIUS server
	Host string `json:"host,omitempty"`

	// UDP port that the RADIUS server listens on for access requests
	Port int64 `json:"port,omitempty"`
}

// Validate validates this get network switch access policy o k body radius servers items0
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network switch access policy o k body radius servers items0 based on context it is used
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchAccessPolicyOKBodyRadiusServersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
