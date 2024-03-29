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

// UpdateNetworkSwitchStackRoutingInterfaceDhcpReader is a Reader for the UpdateNetworkSwitchStackRoutingInterfaceDhcp structure.
type UpdateNetworkSwitchStackRoutingInterfaceDhcpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchStackRoutingInterfaceDhcpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp] updateNetworkSwitchStackRoutingInterfaceDhcp", response, response.Code())
	}
}

// NewUpdateNetworkSwitchStackRoutingInterfaceDhcpOK creates a UpdateNetworkSwitchStackRoutingInterfaceDhcpOK with default headers values
func NewUpdateNetworkSwitchStackRoutingInterfaceDhcpOK() *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK {
	return &UpdateNetworkSwitchStackRoutingInterfaceDhcpOK{}
}

/*
UpdateNetworkSwitchStackRoutingInterfaceDhcpOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchStackRoutingInterfaceDhcpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network switch stack routing interface dhcp o k response has a 2xx status code
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network switch stack routing interface dhcp o k response has a 3xx status code
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network switch stack routing interface dhcp o k response has a 4xx status code
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network switch stack routing interface dhcp o k response has a 5xx status code
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network switch stack routing interface dhcp o k response a status code equal to that given
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network switch stack routing interface dhcp o k response
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) Code() int {
	return 200
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp][%d] updateNetworkSwitchStackRoutingInterfaceDhcpOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/stacks/{switchStackId}/routing/interfaces/{interfaceId}/dhcp][%d] updateNetworkSwitchStackRoutingInterfaceDhcpOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSwitchStackRoutingInterfaceDhcpBody update network switch stack routing interface dhcp body
// Example: {"bootFileName":"home_boot_file","bootNextServer":"1.2.3.4","bootOptionsEnabled":true,"dhcpLeaseTime":"1 day","dhcpMode":"dhcpServer","dhcpOptions":[{"code":"5","type":"text","value":"five"}],"dnsCustomNameservers":["8.8.8.8, 8.8.4.4"],"dnsNameserversOption":"custom","fixedIpAssignments":[{"ip":"192.168.1.12","mac":"22:33:44:55:66:77","name":"Cisco Meraki valued client"}],"reservedIpRanges":[{"comment":"A reserved IP range","end":"192.168.1.10","start":"192.168.1.1"}]}
swagger:model UpdateNetworkSwitchStackRoutingInterfaceDhcpBody
*/
type UpdateNetworkSwitchStackRoutingInterfaceDhcpBody struct {

	// The PXE boot server file name for the DHCP server running on the switch stack interface
	BootFileName string `json:"bootFileName,omitempty"`

	// The PXE boot server IP for the DHCP server running on the switch stack interface
	BootNextServer string `json:"bootNextServer,omitempty"`

	// Enable DHCP boot options to provide PXE boot options configs for the dhcp server running on the switch
	//         stack interface
	BootOptionsEnabled bool `json:"bootOptionsEnabled,omitempty"`

	// The DHCP lease time config for the dhcp server running on switch stack interface
	//         ('30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week')
	// Enum: [1 day 1 hour 1 week 12 hours 30 minutes 4 hours]
	DhcpLeaseTime string `json:"dhcpLeaseTime,omitempty"`

	// The DHCP mode options for the switch stack interface
	//         ('dhcpDisabled', 'dhcpRelay' or 'dhcpServer')
	// Enum: [dhcpDisabled dhcpRelay dhcpServer]
	DhcpMode string `json:"dhcpMode,omitempty"`

	// Array of DHCP options consisting of code, type and value for the DHCP server running on the
	//         switch stack interface
	DhcpOptions []*UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0 `json:"dhcpOptions"`

	// The DHCP relay server IPs to which DHCP packets would get relayed for the switch stack interface
	DhcpRelayServerIps []string `json:"dhcpRelayServerIps"`

	// The DHCP name server IPs when DHCP name server option is '
	//         custom'
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`

	// The DHCP name server option for the dhcp server running on the switch stack interface
	//         ('googlePublicDns', 'openDns' or 'custom')
	// Enum: [custom googlePublicDns openDns]
	DNSNameserversOption string `json:"dnsNameserversOption,omitempty"`

	// Array of DHCP fixed IP assignments for the DHCP server running on the switch stack interface
	FixedIPAssignments []*UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0 `json:"fixedIpAssignments"`

	// Array of DHCP reserved IP assignments for the DHCP server running on the switch stack interface
	ReservedIPRanges []*UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0 `json:"reservedIpRanges"`
}

// Validate validates this update network switch stack routing interface dhcp body
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDhcpLeaseTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDhcpMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDhcpOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDNSNameserversOption(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateFixedIPAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReservedIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpLeaseTimePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1 day","1 hour","1 week","12 hours","30 minutes","4 hours"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpLeaseTimePropEnum = append(updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpLeaseTimePropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr1Day captures enum value "1 day"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr1Day string = "1 day"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr1Hour captures enum value "1 hour"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr1Hour string = "1 hour"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr1Week captures enum value "1 week"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr1Week string = "1 week"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr12Hours captures enum value "12 hours"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr12Hours string = "12 hours"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr30Minutes captures enum value "30 minutes"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr30Minutes string = "30 minutes"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr4Hours captures enum value "4 hours"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpLeaseTimeNr4Hours string = "4 hours"
)

// prop value enum
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateDhcpLeaseTimeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpLeaseTimePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateDhcpLeaseTime(formats strfmt.Registry) error {
	if swag.IsZero(o.DhcpLeaseTime) { // not required
		return nil
	}

	// value enum
	if err := o.validateDhcpLeaseTimeEnum("updateNetworkSwitchStackRoutingInterfaceDhcp"+"."+"dhcpLeaseTime", "body", o.DhcpLeaseTime); err != nil {
		return err
	}

	return nil
}

var updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["dhcpDisabled","dhcpRelay","dhcpServer"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpModePropEnum = append(updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpModePropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpModeDhcpDisabled captures enum value "dhcpDisabled"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpModeDhcpDisabled string = "dhcpDisabled"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpModeDhcpRelay captures enum value "dhcpRelay"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpModeDhcpRelay string = "dhcpRelay"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpModeDhcpServer captures enum value "dhcpServer"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDhcpModeDhcpServer string = "dhcpServer"
)

// prop value enum
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateDhcpModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDhcpModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateDhcpMode(formats strfmt.Registry) error {
	if swag.IsZero(o.DhcpMode) { // not required
		return nil
	}

	// value enum
	if err := o.validateDhcpModeEnum("updateNetworkSwitchStackRoutingInterfaceDhcp"+"."+"dhcpMode", "body", o.DhcpMode); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateDhcpOptions(formats strfmt.Registry) error {
	if swag.IsZero(o.DhcpOptions) { // not required
		return nil
	}

	for i := 0; i < len(o.DhcpOptions); i++ {
		if swag.IsZero(o.DhcpOptions[i]) { // not required
			continue
		}

		if o.DhcpOptions[i] != nil {
			if err := o.DhcpOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "dhcpOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "dhcpOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDNSNameserversOptionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["custom","googlePublicDns","openDns"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDNSNameserversOptionPropEnum = append(updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDNSNameserversOptionPropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDNSNameserversOptionCustom captures enum value "custom"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDNSNameserversOptionCustom string = "custom"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDNSNameserversOptionGooglePublicDNS captures enum value "googlePublicDns"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDNSNameserversOptionGooglePublicDNS string = "googlePublicDns"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDNSNameserversOptionOpenDNS captures enum value "openDns"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpBodyDNSNameserversOptionOpenDNS string = "openDns"
)

// prop value enum
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateDNSNameserversOptionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchStackRoutingInterfaceDhcpBodyTypeDNSNameserversOptionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateDNSNameserversOption(formats strfmt.Registry) error {
	if swag.IsZero(o.DNSNameserversOption) { // not required
		return nil
	}

	// value enum
	if err := o.validateDNSNameserversOptionEnum("updateNetworkSwitchStackRoutingInterfaceDhcp"+"."+"dnsNameserversOption", "body", o.DNSNameserversOption); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateFixedIPAssignments(formats strfmt.Registry) error {
	if swag.IsZero(o.FixedIPAssignments) { // not required
		return nil
	}

	for i := 0; i < len(o.FixedIPAssignments); i++ {
		if swag.IsZero(o.FixedIPAssignments[i]) { // not required
			continue
		}

		if o.FixedIPAssignments[i] != nil {
			if err := o.FixedIPAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) validateReservedIPRanges(formats strfmt.Registry) error {
	if swag.IsZero(o.ReservedIPRanges) { // not required
		return nil
	}

	for i := 0; i < len(o.ReservedIPRanges); i++ {
		if swag.IsZero(o.ReservedIPRanges[i]) { // not required
			continue
		}

		if o.ReservedIPRanges[i] != nil {
			if err := o.ReservedIPRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch stack routing interface dhcp body based on the context it is used
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDhcpOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateFixedIPAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateReservedIPRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) contextValidateDhcpOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DhcpOptions); i++ {

		if o.DhcpOptions[i] != nil {

			if swag.IsZero(o.DhcpOptions[i]) { // not required
				return nil
			}

			if err := o.DhcpOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "dhcpOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "dhcpOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) contextValidateFixedIPAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FixedIPAssignments); i++ {

		if o.FixedIPAssignments[i] != nil {

			if swag.IsZero(o.FixedIPAssignments[i]) { // not required
				return nil
			}

			if err := o.FixedIPAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) contextValidateReservedIPRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ReservedIPRanges); i++ {

		if o.ReservedIPRanges[i] != nil {

			if swag.IsZero(o.ReservedIPRanges[i]) { // not required
				return nil
			}

			if err := o.ReservedIPRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchStackRoutingInterfaceDhcp" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchStackRoutingInterfaceDhcpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0 update network switch stack routing interface dhcp params body dhcp options items0
swagger:model UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0
*/
type UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0 struct {

	// The code for DHCP option which should be from 2 to 254
	// Required: true
	Code *string `json:"code"`

	// The type of the DHCP option which should be one of
	//           ('text', 'ip', 'integer' or 'hex')
	// Required: true
	// Enum: [hex integer ip text]
	Type *string `json:"type"`

	// The value of the DHCP option
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update network switch stack routing interface dhcp params body dhcp options items0
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", o.Code); err != nil {
		return err
	}

	return nil
}

var updateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hex","integer","ip","text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeTypePropEnum = append(updateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeTypePropEnum, v)
	}
}

const (

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeHex captures enum value "hex"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeHex string = "hex"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeInteger captures enum value "integer"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeInteger string = "integer"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeIP captures enum value "ip"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeIP string = "ip"

	// UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeText captures enum value "text"
	UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeText string = "text"
)

// prop value enum
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch stack routing interface dhcp params body dhcp options items0 based on context it is used
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyDhcpOptionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0 update network switch stack routing interface dhcp params body fixed IP assignments items0
swagger:model UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0
*/
type UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0 struct {

	// The IP address of the client which has fixed IP address assigned to it
	// Required: true
	IP *string `json:"ip"`

	// The MAC address of the client which has fixed IP address
	// Required: true
	Mac *string `json:"mac"`

	// The name of the client which has fixed IP address
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this update network switch stack routing interface dhcp params body fixed IP assignments items0
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", o.IP); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0) validateMac(formats strfmt.Registry) error {

	if err := validate.Required("mac", "body", o.Mac); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch stack routing interface dhcp params body fixed IP assignments items0 based on context it is used
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyFixedIPAssignmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0 update network switch stack routing interface dhcp params body reserved IP ranges items0
swagger:model UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0
*/
type UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0 struct {

	// The comment for the reserved IP range
	Comment string `json:"comment,omitempty"`

	// The ending IP address of the reserved IP range
	// Required: true
	End *string `json:"end"`

	// The starting IP address of the reserved IP range
	// Required: true
	Start *string `json:"start"`
}

// Validate validates this update network switch stack routing interface dhcp params body reserved IP ranges items0
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", o.End); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", o.Start); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch stack routing interface dhcp params body reserved IP ranges items0 based on context it is used
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchStackRoutingInterfaceDhcpParamsBodyReservedIPRangesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
