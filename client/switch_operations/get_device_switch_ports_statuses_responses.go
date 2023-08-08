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

// GetDeviceSwitchPortsStatusesReader is a Reader for the GetDeviceSwitchPortsStatuses structure.
type GetDeviceSwitchPortsStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceSwitchPortsStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceSwitchPortsStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}/switch/ports/statuses] getDeviceSwitchPortsStatuses", response, response.Code())
	}
}

// NewGetDeviceSwitchPortsStatusesOK creates a GetDeviceSwitchPortsStatusesOK with default headers values
func NewGetDeviceSwitchPortsStatusesOK() *GetDeviceSwitchPortsStatusesOK {
	return &GetDeviceSwitchPortsStatusesOK{}
}

/*
GetDeviceSwitchPortsStatusesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceSwitchPortsStatusesOK struct {
	Payload []*GetDeviceSwitchPortsStatusesOKBodyItems0
}

// IsSuccess returns true when this get device switch ports statuses o k response has a 2xx status code
func (o *GetDeviceSwitchPortsStatusesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device switch ports statuses o k response has a 3xx status code
func (o *GetDeviceSwitchPortsStatusesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device switch ports statuses o k response has a 4xx status code
func (o *GetDeviceSwitchPortsStatusesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device switch ports statuses o k response has a 5xx status code
func (o *GetDeviceSwitchPortsStatusesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device switch ports statuses o k response a status code equal to that given
func (o *GetDeviceSwitchPortsStatusesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device switch ports statuses o k response
func (o *GetDeviceSwitchPortsStatusesOK) Code() int {
	return 200
}

func (o *GetDeviceSwitchPortsStatusesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/ports/statuses][%d] getDeviceSwitchPortsStatusesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchPortsStatusesOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/ports/statuses][%d] getDeviceSwitchPortsStatusesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchPortsStatusesOK) GetPayload() []*GetDeviceSwitchPortsStatusesOKBodyItems0 {
	return o.Payload
}

func (o *GetDeviceSwitchPortsStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetDeviceSwitchPortsStatusesOKBodyItems0 get device switch ports statuses o k body items0
swagger:model GetDeviceSwitchPortsStatusesOKBodyItems0
*/
type GetDeviceSwitchPortsStatusesOKBodyItems0 struct {

	// cdp
	Cdp *GetDeviceSwitchPortsStatusesOKBodyItems0Cdp `json:"cdp,omitempty"`

	// The number of clients connected through this port.
	ClientCount int64 `json:"clientCount,omitempty"`

	// The current duplex of a connected port.
	// Enum: [ full half]
	Duplex string `json:"duplex,omitempty"`

	// Whether the port is configured to be enabled.
	Enabled bool `json:"enabled,omitempty"`

	// All errors present on the port.
	Errors []string `json:"errors"`

	// Whether the port is the switch's uplink.
	IsUplink bool `json:"isUplink,omitempty"`

	// lldp
	Lldp *GetDeviceSwitchPortsStatusesOKBodyItems0Lldp `json:"lldp,omitempty"`

	// The string identifier of this port on the switch. This is commonly just the port number but may contain additional identifying information such as the slot and module-type if the port is located on a port module.
	PortID string `json:"portId,omitempty"`

	// How much power (in watt-hours) has been delivered by this port during the timespan.
	PowerUsageInWh float32 `json:"powerUsageInWh,omitempty"`

	// secure port
	SecurePort *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort `json:"securePort,omitempty"`

	// The current data transfer rate which the port is operating at.
	// Enum: [ 1 Gbps 10 Gbps 10 Mbps 100 Gbps 100 Mbps 2.5 Gbps 20 Gbps 40 Gbps 5 Gbps]
	Speed string `json:"speed,omitempty"`

	// The current connection status of the port.
	// Enum: [Connected Disabled Disconnected]
	Status string `json:"status,omitempty"`

	// traffic in kbps
	TrafficInKbps *GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps `json:"trafficInKbps,omitempty"`

	// usage in kb
	UsageInKb *GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb `json:"usageInKb,omitempty"`

	// All warnings present on the port.
	Warnings []string `json:"warnings"`
}

// Validate validates this get device switch ports statuses o k body items0
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCdp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDuplex(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLldp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecurePort(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTrafficInKbps(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsageInKb(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateCdp(formats strfmt.Registry) error {
	if swag.IsZero(o.Cdp) { // not required
		return nil
	}

	if o.Cdp != nil {
		if err := o.Cdp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdp")
			}
			return err
		}
	}

	return nil
}

var getDeviceSwitchPortsStatusesOKBodyItems0TypeDuplexPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","full","half"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDeviceSwitchPortsStatusesOKBodyItems0TypeDuplexPropEnum = append(getDeviceSwitchPortsStatusesOKBodyItems0TypeDuplexPropEnum, v)
	}
}

const (

	// GetDeviceSwitchPortsStatusesOKBodyItems0DuplexEmpty captures enum value ""
	GetDeviceSwitchPortsStatusesOKBodyItems0DuplexEmpty string = ""

	// GetDeviceSwitchPortsStatusesOKBodyItems0DuplexFull captures enum value "full"
	GetDeviceSwitchPortsStatusesOKBodyItems0DuplexFull string = "full"

	// GetDeviceSwitchPortsStatusesOKBodyItems0DuplexHalf captures enum value "half"
	GetDeviceSwitchPortsStatusesOKBodyItems0DuplexHalf string = "half"
)

// prop value enum
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateDuplexEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDeviceSwitchPortsStatusesOKBodyItems0TypeDuplexPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateDuplex(formats strfmt.Registry) error {
	if swag.IsZero(o.Duplex) { // not required
		return nil
	}

	// value enum
	if err := o.validateDuplexEnum("duplex", "body", o.Duplex); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateLldp(formats strfmt.Registry) error {
	if swag.IsZero(o.Lldp) { // not required
		return nil
	}

	if o.Lldp != nil {
		if err := o.Lldp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lldp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lldp")
			}
			return err
		}
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateSecurePort(formats strfmt.Registry) error {
	if swag.IsZero(o.SecurePort) { // not required
		return nil
	}

	if o.SecurePort != nil {
		if err := o.SecurePort.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securePort")
			}
			return err
		}
	}

	return nil
}

var getDeviceSwitchPortsStatusesOKBodyItems0TypeSpeedPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["","1 Gbps","10 Gbps","10 Mbps","100 Gbps","100 Mbps","2.5 Gbps","20 Gbps","40 Gbps","5 Gbps"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDeviceSwitchPortsStatusesOKBodyItems0TypeSpeedPropEnum = append(getDeviceSwitchPortsStatusesOKBodyItems0TypeSpeedPropEnum, v)
	}
}

const (

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedEmpty captures enum value ""
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedEmpty string = ""

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr1Gbps captures enum value "1 Gbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr1Gbps string = "1 Gbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr10Gbps captures enum value "10 Gbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr10Gbps string = "10 Gbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr10Mbps captures enum value "10 Mbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr10Mbps string = "10 Mbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr100Gbps captures enum value "100 Gbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr100Gbps string = "100 Gbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr100Mbps captures enum value "100 Mbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr100Mbps string = "100 Mbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr2Dot5Gbps captures enum value "2.5 Gbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr2Dot5Gbps string = "2.5 Gbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr20Gbps captures enum value "20 Gbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr20Gbps string = "20 Gbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr40Gbps captures enum value "40 Gbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr40Gbps string = "40 Gbps"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr5Gbps captures enum value "5 Gbps"
	GetDeviceSwitchPortsStatusesOKBodyItems0SpeedNr5Gbps string = "5 Gbps"
)

// prop value enum
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateSpeedEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDeviceSwitchPortsStatusesOKBodyItems0TypeSpeedPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateSpeed(formats strfmt.Registry) error {
	if swag.IsZero(o.Speed) { // not required
		return nil
	}

	// value enum
	if err := o.validateSpeedEnum("speed", "body", o.Speed); err != nil {
		return err
	}

	return nil
}

var getDeviceSwitchPortsStatusesOKBodyItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Connected","Disabled","Disconnected"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDeviceSwitchPortsStatusesOKBodyItems0TypeStatusPropEnum = append(getDeviceSwitchPortsStatusesOKBodyItems0TypeStatusPropEnum, v)
	}
}

const (

	// GetDeviceSwitchPortsStatusesOKBodyItems0StatusConnected captures enum value "Connected"
	GetDeviceSwitchPortsStatusesOKBodyItems0StatusConnected string = "Connected"

	// GetDeviceSwitchPortsStatusesOKBodyItems0StatusDisabled captures enum value "Disabled"
	GetDeviceSwitchPortsStatusesOKBodyItems0StatusDisabled string = "Disabled"

	// GetDeviceSwitchPortsStatusesOKBodyItems0StatusDisconnected captures enum value "Disconnected"
	GetDeviceSwitchPortsStatusesOKBodyItems0StatusDisconnected string = "Disconnected"
)

// prop value enum
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDeviceSwitchPortsStatusesOKBodyItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateTrafficInKbps(formats strfmt.Registry) error {
	if swag.IsZero(o.TrafficInKbps) { // not required
		return nil
	}

	if o.TrafficInKbps != nil {
		if err := o.TrafficInKbps.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trafficInKbps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trafficInKbps")
			}
			return err
		}
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) validateUsageInKb(formats strfmt.Registry) error {
	if swag.IsZero(o.UsageInKb) { // not required
		return nil
	}

	if o.UsageInKb != nil {
		if err := o.UsageInKb.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usageInKb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usageInKb")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get device switch ports statuses o k body items0 based on the context it is used
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCdp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLldp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSecurePort(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTrafficInKbps(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUsageInKb(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) contextValidateCdp(ctx context.Context, formats strfmt.Registry) error {

	if o.Cdp != nil {

		if swag.IsZero(o.Cdp) { // not required
			return nil
		}

		if err := o.Cdp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cdp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cdp")
			}
			return err
		}
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) contextValidateLldp(ctx context.Context, formats strfmt.Registry) error {

	if o.Lldp != nil {

		if swag.IsZero(o.Lldp) { // not required
			return nil
		}

		if err := o.Lldp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lldp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("lldp")
			}
			return err
		}
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) contextValidateSecurePort(ctx context.Context, formats strfmt.Registry) error {

	if o.SecurePort != nil {

		if swag.IsZero(o.SecurePort) { // not required
			return nil
		}

		if err := o.SecurePort.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securePort")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securePort")
			}
			return err
		}
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) contextValidateTrafficInKbps(ctx context.Context, formats strfmt.Registry) error {

	if o.TrafficInKbps != nil {

		if swag.IsZero(o.TrafficInKbps) { // not required
			return nil
		}

		if err := o.TrafficInKbps.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("trafficInKbps")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("trafficInKbps")
			}
			return err
		}
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) contextValidateUsageInKb(ctx context.Context, formats strfmt.Registry) error {

	if o.UsageInKb != nil {

		if swag.IsZero(o.UsageInKb) { // not required
			return nil
		}

		if err := o.UsageInKb.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usageInKb")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usageInKb")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetDeviceSwitchPortsStatusesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDeviceSwitchPortsStatusesOKBodyItems0Cdp The Cisco Discovery Protocol (CDP) information of the connected device.
swagger:model GetDeviceSwitchPortsStatusesOKBodyItems0Cdp
*/
type GetDeviceSwitchPortsStatusesOKBodyItems0Cdp struct {

	// Contains network addresses of both receiving and sending devices.
	Address string `json:"address,omitempty"`

	// Identifies the device type, which indicates the functional capabilities of the device.
	Capabilities string `json:"capabilities,omitempty"`

	// Identifies the device name.
	DeviceID string `json:"deviceId,omitempty"`

	// The device's management IP.
	ManagementAddress string `json:"managementAddress,omitempty"`

	// Indicates, per interface, the assumed VLAN for untagged packets on the interface.
	NativeVlan int64 `json:"nativeVlan,omitempty"`

	// Identifies the hardware platform of the device.
	Platform string `json:"platform,omitempty"`

	// Identifies the port from which the CDP packet was sent.
	PortID string `json:"portId,omitempty"`

	// The system name.
	SystemName string `json:"systemName,omitempty"`

	// Contains the device software release information.
	Version string `json:"version,omitempty"`

	// Advertises the configured VLAN Trunking Protocl (VTP)-management-domain name of the system.
	VtpManagementDomain string `json:"vtpManagementDomain,omitempty"`
}

// Validate validates this get device switch ports statuses o k body items0 cdp
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Cdp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get device switch ports statuses o k body items0 cdp based on context it is used
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Cdp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Cdp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Cdp) UnmarshalBinary(b []byte) error {
	var res GetDeviceSwitchPortsStatusesOKBodyItems0Cdp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDeviceSwitchPortsStatusesOKBodyItems0Lldp The Link Layer Discovery Protocol (LLDP) information of the connected device.
swagger:model GetDeviceSwitchPortsStatusesOKBodyItems0Lldp
*/
type GetDeviceSwitchPortsStatusesOKBodyItems0Lldp struct {

	// The device's chassis ID.
	ChassisID string `json:"chassisId,omitempty"`

	// The device's management IP.
	ManagementAddress string `json:"managementAddress,omitempty"`

	// The device's management VLAN.
	ManagementVlan int64 `json:"managementVlan,omitempty"`

	// Description of the port from which the LLDP packet was sent.
	PortDescription string `json:"portDescription,omitempty"`

	// Identifies the port from which the LLDP packet was sent
	PortID string `json:"portId,omitempty"`

	// The port's VLAN.
	PortVlan int64 `json:"portVlan,omitempty"`

	// Identifies the device type, which indicates the functional capabilities of the device.
	SystemCapabilities string `json:"systemCapabilities,omitempty"`

	// The device's system description.
	SystemDescription string `json:"systemDescription,omitempty"`

	// The device's system name.
	SystemName string `json:"systemName,omitempty"`
}

// Validate validates this get device switch ports statuses o k body items0 lldp
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Lldp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get device switch ports statuses o k body items0 lldp based on context it is used
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Lldp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Lldp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0Lldp) UnmarshalBinary(b []byte) error {
	var res GetDeviceSwitchPortsStatusesOKBodyItems0Lldp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort The Secure Port status of the port.
swagger:model GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort
*/
type GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort struct {

	// Whether Secure Port is currently active for this port.
	Active bool `json:"active,omitempty"`

	// The current Secure Port status.
	// Enum: [Authentication failure Authentication in progress Authentication successful Authentication timed out Disabled Enabled]
	AuthenticationStatus string `json:"authenticationStatus,omitempty"`

	// config overrides
	ConfigOverrides *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides `json:"configOverrides,omitempty"`

	// Whether Secure Port is turned on for this port.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this get device switch ports statuses o k body items0 secure port
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthenticationStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateConfigOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getDeviceSwitchPortsStatusesOKBodyItems0SecurePortTypeAuthenticationStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Authentication failure","Authentication in progress","Authentication successful","Authentication timed out","Disabled","Enabled"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDeviceSwitchPortsStatusesOKBodyItems0SecurePortTypeAuthenticationStatusPropEnum = append(getDeviceSwitchPortsStatusesOKBodyItems0SecurePortTypeAuthenticationStatusPropEnum, v)
	}
}

const (

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationFailure captures enum value "Authentication failure"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationFailure string = "Authentication failure"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationInProgress captures enum value "Authentication in progress"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationInProgress string = "Authentication in progress"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationSuccessful captures enum value "Authentication successful"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationSuccessful string = "Authentication successful"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationTimedOut captures enum value "Authentication timed out"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusAuthenticationTimedOut string = "Authentication timed out"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusDisabled captures enum value "Disabled"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusDisabled string = "Disabled"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusEnabled captures enum value "Enabled"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortAuthenticationStatusEnabled string = "Enabled"
)

// prop value enum
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) validateAuthenticationStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDeviceSwitchPortsStatusesOKBodyItems0SecurePortTypeAuthenticationStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) validateAuthenticationStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.AuthenticationStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateAuthenticationStatusEnum("securePort"+"."+"authenticationStatus", "body", o.AuthenticationStatus); err != nil {
		return err
	}

	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) validateConfigOverrides(formats strfmt.Registry) error {
	if swag.IsZero(o.ConfigOverrides) { // not required
		return nil
	}

	if o.ConfigOverrides != nil {
		if err := o.ConfigOverrides.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securePort" + "." + "configOverrides")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securePort" + "." + "configOverrides")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get device switch ports statuses o k body items0 secure port based on the context it is used
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConfigOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) contextValidateConfigOverrides(ctx context.Context, formats strfmt.Registry) error {

	if o.ConfigOverrides != nil {

		if swag.IsZero(o.ConfigOverrides) { // not required
			return nil
		}

		if err := o.ConfigOverrides.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("securePort" + "." + "configOverrides")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("securePort" + "." + "configOverrides")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort) UnmarshalBinary(b []byte) error {
	var res GetDeviceSwitchPortsStatusesOKBodyItems0SecurePort
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides The configuration overrides applied to this port when Secure Port is active.
swagger:model GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides
*/
type GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides struct {

	// The VLANs allowed on the . Only applicable to trunk ports.
	AllowedVlans string `json:"allowedVlans,omitempty"`

	// The type of the  ('trunk' or 'access').
	// Enum: [access trunk]
	Type string `json:"type,omitempty"`

	// The VLAN of the . A null value will clear the value set for trunk ports.
	Vlan int64 `json:"vlan,omitempty"`

	// The voice VLAN of the . Only applicable to access ports.
	VoiceVlan int64 `json:"voiceVlan,omitempty"`
}

// Validate validates this get device switch ports statuses o k body items0 secure port config overrides
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["access","trunk"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeTypePropEnum = append(getDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeTypePropEnum, v)
	}
}

const (

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeAccess captures enum value "access"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeAccess string = "access"

	// GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeTrunk captures enum value "trunk"
	GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeTrunk string = "trunk"
)

// prop value enum
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverridesTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("securePort"+"."+"configOverrides"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get device switch ports statuses o k body items0 secure port config overrides based on context it is used
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides) UnmarshalBinary(b []byte) error {
	var res GetDeviceSwitchPortsStatusesOKBodyItems0SecurePortConfigOverrides
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps A breakdown of the average speed of data that has passed through this port during the timespan.
swagger:model GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps
*/
type GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps struct {

	// The average speed of the data received (in kilobits-per-second).
	Recv float32 `json:"recv,omitempty"`

	// The average speed of the data sent (in kilobits-per-second).
	Sent float32 `json:"sent,omitempty"`

	// The average speed of the data sent and received (in kilobits-per-second).
	Total float32 `json:"total,omitempty"`
}

// Validate validates this get device switch ports statuses o k body items0 traffic in kbps
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get device switch ports statuses o k body items0 traffic in kbps based on context it is used
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps) UnmarshalBinary(b []byte) error {
	var res GetDeviceSwitchPortsStatusesOKBodyItems0TrafficInKbps
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb A breakdown of how many kilobytes have passed through this port during the timespan.
swagger:model GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb
*/
type GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb struct {

	// The amount of data received (in kilobytes).
	Recv int64 `json:"recv,omitempty"`

	// The amount of data sent (in kilobytes).
	Sent int64 `json:"sent,omitempty"`

	// The total amount of data sent and received (in kilobytes).
	Total int64 `json:"total,omitempty"`
}

// Validate validates this get device switch ports statuses o k body items0 usage in kb
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get device switch ports statuses o k body items0 usage in kb based on context it is used
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb) UnmarshalBinary(b []byte) error {
	var res GetDeviceSwitchPortsStatusesOKBodyItems0UsageInKb
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
