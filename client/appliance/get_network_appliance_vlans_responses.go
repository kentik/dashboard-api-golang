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

// GetNetworkApplianceVlansReader is a Reader for the GetNetworkApplianceVlans structure.
type GetNetworkApplianceVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/appliance/vlans] getNetworkApplianceVlans", response, response.Code())
	}
}

// NewGetNetworkApplianceVlansOK creates a GetNetworkApplianceVlansOK with default headers values
func NewGetNetworkApplianceVlansOK() *GetNetworkApplianceVlansOK {
	return &GetNetworkApplianceVlansOK{}
}

/*
GetNetworkApplianceVlansOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceVlansOK struct {
	Payload []*GetNetworkApplianceVlansOKBodyItems0
}

// IsSuccess returns true when this get network appliance vlans o k response has a 2xx status code
func (o *GetNetworkApplianceVlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance vlans o k response has a 3xx status code
func (o *GetNetworkApplianceVlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance vlans o k response has a 4xx status code
func (o *GetNetworkApplianceVlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance vlans o k response has a 5xx status code
func (o *GetNetworkApplianceVlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance vlans o k response a status code equal to that given
func (o *GetNetworkApplianceVlansOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network appliance vlans o k response
func (o *GetNetworkApplianceVlansOK) Code() int {
	return 200
}

func (o *GetNetworkApplianceVlansOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/vlans][%d] getNetworkApplianceVlansOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceVlansOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/vlans][%d] getNetworkApplianceVlansOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceVlansOK) GetPayload() []*GetNetworkApplianceVlansOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkApplianceVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkApplianceVlansOKBodyItems0 get network appliance vlans o k body items0
swagger:model GetNetworkApplianceVlansOKBodyItems0
*/
type GetNetworkApplianceVlansOKBodyItems0 struct {

	// The local IP of the appliance on the VLAN
	ApplianceIP string `json:"applianceIp,omitempty"`

	// CIDR of the pool of subnets. Applicable only for template network. Each network bound to the template will automatically pick a subnet from this pool to build its own VLAN.
	Cidr string `json:"cidr,omitempty"`

	// DHCP boot option for boot filename
	DhcpBootFilename string `json:"dhcpBootFilename,omitempty"`

	// DHCP boot option to direct boot clients to the server to load the boot file from
	DhcpBootNextServer string `json:"dhcpBootNextServer,omitempty"`

	// Use DHCP boot options specified in other properties
	DhcpBootOptionsEnabled bool `json:"dhcpBootOptionsEnabled,omitempty"`

	// The appliance's handling of DHCP requests on this VLAN. One of: 'Run a DHCP server', 'Relay DHCP to another server' or 'Do not respond to DHCP requests'
	// Enum: [Do not respond to DHCP requests Relay DHCP to another server Run a DHCP server]
	DhcpHandling string `json:"dhcpHandling,omitempty"`

	// The term of DHCP leases if the appliance is running a DHCP server on this VLAN. One of: '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'
	// Enum: [1 day 1 hour 1 week 12 hours 30 minutes 4 hours]
	DhcpLeaseTime string `json:"dhcpLeaseTime,omitempty"`

	// The list of DHCP options that will be included in DHCP responses. Each object in the list should have "code", "type", and "value" properties.
	DhcpOptions []*GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0 `json:"dhcpOptions"`

	// The IPs of the DHCP servers that DHCP requests should be relayed to
	DhcpRelayServerIps []string `json:"dhcpRelayServerIps"`

	// The DNS nameservers used for DHCP responses, either "upstream_dns", "google_dns", "opendns", or a newline seperated string of IP addresses or domain names
	DNSNameservers string `json:"dnsNameservers,omitempty"`

	// The DHCP fixed IP assignments on the VLAN. This should be an object that contains mappings from MAC addresses to objects that themselves each contain "ip" and "name" string fields. See the sample request/response for more details.
	FixedIPAssignments interface{} `json:"fixedIpAssignments,omitempty"`

	// The id of the desired group policy to apply to the VLAN
	GroupPolicyID string `json:"groupPolicyId,omitempty"`

	// The VLAN ID of the VLAN
	ID string `json:"id,omitempty"`

	// The interface ID of the VLAN
	InterfaceID string `json:"interfaceId,omitempty"`

	// ipv6
	IPV6 *GetNetworkApplianceVlansOKBodyItems0IPV6 `json:"ipv6,omitempty"`

	// mandatory dhcp
	MandatoryDhcp *GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp `json:"mandatoryDhcp,omitempty"`

	// Mask used for the subnet of all bound to the template networks. Applicable only for template network.
	Mask int64 `json:"mask,omitempty"`

	// The name of the VLAN
	Name string `json:"name,omitempty"`

	// The DHCP reserved IP ranges on the VLAN
	ReservedIPRanges []*GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0 `json:"reservedIpRanges"`

	// The subnet of the VLAN
	Subnet string `json:"subnet,omitempty"`

	// Type of subnetting of the VLAN. Applicable only for template network.
	// Enum: [same unique]
	TemplateVlanType *string `json:"templateVlanType,omitempty"`

	// The translated VPN subnet if VPN and VPN subnet translation are enabled on the VLAN
	VpnNatSubnet string `json:"vpnNatSubnet,omitempty"`
}

// Validate validates this get network appliance vlans o k body items0
func (o *GetNetworkApplianceVlansOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDhcpHandling(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDhcpLeaseTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDhcpOptions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMandatoryDhcp(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReservedIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTemplateVlanType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkApplianceVlansOKBodyItems0TypeDhcpHandlingPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Do not respond to DHCP requests","Relay DHCP to another server","Run a DHCP server"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkApplianceVlansOKBodyItems0TypeDhcpHandlingPropEnum = append(getNetworkApplianceVlansOKBodyItems0TypeDhcpHandlingPropEnum, v)
	}
}

const (

	// GetNetworkApplianceVlansOKBodyItems0DhcpHandlingDoNotRespondToDHCPRequests captures enum value "Do not respond to DHCP requests"
	GetNetworkApplianceVlansOKBodyItems0DhcpHandlingDoNotRespondToDHCPRequests string = "Do not respond to DHCP requests"

	// GetNetworkApplianceVlansOKBodyItems0DhcpHandlingRelayDHCPToAnotherServer captures enum value "Relay DHCP to another server"
	GetNetworkApplianceVlansOKBodyItems0DhcpHandlingRelayDHCPToAnotherServer string = "Relay DHCP to another server"

	// GetNetworkApplianceVlansOKBodyItems0DhcpHandlingRunaDHCPServer captures enum value "Run a DHCP server"
	GetNetworkApplianceVlansOKBodyItems0DhcpHandlingRunaDHCPServer string = "Run a DHCP server"
)

// prop value enum
func (o *GetNetworkApplianceVlansOKBodyItems0) validateDhcpHandlingEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkApplianceVlansOKBodyItems0TypeDhcpHandlingPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) validateDhcpHandling(formats strfmt.Registry) error {
	if swag.IsZero(o.DhcpHandling) { // not required
		return nil
	}

	// value enum
	if err := o.validateDhcpHandlingEnum("dhcpHandling", "body", o.DhcpHandling); err != nil {
		return err
	}

	return nil
}

var getNetworkApplianceVlansOKBodyItems0TypeDhcpLeaseTimePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1 day","1 hour","1 week","12 hours","30 minutes","4 hours"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkApplianceVlansOKBodyItems0TypeDhcpLeaseTimePropEnum = append(getNetworkApplianceVlansOKBodyItems0TypeDhcpLeaseTimePropEnum, v)
	}
}

const (

	// GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr1Day captures enum value "1 day"
	GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr1Day string = "1 day"

	// GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr1Hour captures enum value "1 hour"
	GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr1Hour string = "1 hour"

	// GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr1Week captures enum value "1 week"
	GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr1Week string = "1 week"

	// GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr12Hours captures enum value "12 hours"
	GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr12Hours string = "12 hours"

	// GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr30Minutes captures enum value "30 minutes"
	GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr30Minutes string = "30 minutes"

	// GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr4Hours captures enum value "4 hours"
	GetNetworkApplianceVlansOKBodyItems0DhcpLeaseTimeNr4Hours string = "4 hours"
)

// prop value enum
func (o *GetNetworkApplianceVlansOKBodyItems0) validateDhcpLeaseTimeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkApplianceVlansOKBodyItems0TypeDhcpLeaseTimePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) validateDhcpLeaseTime(formats strfmt.Registry) error {
	if swag.IsZero(o.DhcpLeaseTime) { // not required
		return nil
	}

	// value enum
	if err := o.validateDhcpLeaseTimeEnum("dhcpLeaseTime", "body", o.DhcpLeaseTime); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) validateDhcpOptions(formats strfmt.Registry) error {
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
					return ve.ValidateName("dhcpOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcpOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) validateMandatoryDhcp(formats strfmt.Registry) error {
	if swag.IsZero(o.MandatoryDhcp) { // not required
		return nil
	}

	if o.MandatoryDhcp != nil {
		if err := o.MandatoryDhcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) validateReservedIPRanges(formats strfmt.Registry) error {
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
					return ve.ValidateName("reservedIpRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reservedIpRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getNetworkApplianceVlansOKBodyItems0TypeTemplateVlanTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["same","unique"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkApplianceVlansOKBodyItems0TypeTemplateVlanTypePropEnum = append(getNetworkApplianceVlansOKBodyItems0TypeTemplateVlanTypePropEnum, v)
	}
}

const (

	// GetNetworkApplianceVlansOKBodyItems0TemplateVlanTypeSame captures enum value "same"
	GetNetworkApplianceVlansOKBodyItems0TemplateVlanTypeSame string = "same"

	// GetNetworkApplianceVlansOKBodyItems0TemplateVlanTypeUnique captures enum value "unique"
	GetNetworkApplianceVlansOKBodyItems0TemplateVlanTypeUnique string = "unique"
)

// prop value enum
func (o *GetNetworkApplianceVlansOKBodyItems0) validateTemplateVlanTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkApplianceVlansOKBodyItems0TypeTemplateVlanTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) validateTemplateVlanType(formats strfmt.Registry) error {
	if swag.IsZero(o.TemplateVlanType) { // not required
		return nil
	}

	// value enum
	if err := o.validateTemplateVlanTypeEnum("templateVlanType", "body", *o.TemplateVlanType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get network appliance vlans o k body items0 based on the context it is used
func (o *GetNetworkApplianceVlansOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDhcpOptions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMandatoryDhcp(ctx, formats); err != nil {
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

func (o *GetNetworkApplianceVlansOKBodyItems0) contextValidateDhcpOptions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.DhcpOptions); i++ {

		if o.DhcpOptions[i] != nil {

			if swag.IsZero(o.DhcpOptions[i]) { // not required
				return nil
			}

			if err := o.DhcpOptions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("dhcpOptions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("dhcpOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {

		if swag.IsZero(o.IPV6) { // not required
			return nil
		}

		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) contextValidateMandatoryDhcp(ctx context.Context, formats strfmt.Registry) error {

	if o.MandatoryDhcp != nil {

		if swag.IsZero(o.MandatoryDhcp) { // not required
			return nil
		}

		if err := o.MandatoryDhcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0) contextValidateReservedIPRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ReservedIPRanges); i++ {

		if o.ReservedIPRanges[i] != nil {

			if swag.IsZero(o.ReservedIPRanges[i]) { // not required
				return nil
			}

			if err := o.ReservedIPRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reservedIpRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("reservedIpRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceVlansOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0 get network appliance vlans o k body items0 dhcp options items0
swagger:model GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0
*/
type GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0 struct {

	// The code for the DHCP option. This should be an integer between 2 and 254.
	// Required: true
	Code *string `json:"code"`

	// The type for the DHCP option. One of: 'text', 'ip', 'hex' or 'integer'
	// Required: true
	// Enum: [hex integer ip text]
	Type *string `json:"type"`

	// The value for the DHCP option
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this get network appliance vlans o k body items0 dhcp options items0
func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) Validate(formats strfmt.Registry) error {
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

func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", o.Code); err != nil {
		return err
	}

	return nil
}

var getNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["hex","integer","ip","text"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeTypePropEnum = append(getNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeTypePropEnum, v)
	}
}

const (

	// GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeHex captures enum value "hex"
	GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeHex string = "hex"

	// GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeInteger captures enum value "integer"
	GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeInteger string = "integer"

	// GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeIP captures enum value "ip"
	GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeIP string = "ip"

	// GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeText captures enum value "text"
	GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeText string = "text"
)

// prop value enum
func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network appliance vlans o k body items0 dhcp options items0 based on context it is used
func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceVlansOKBodyItems0DhcpOptionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkApplianceVlansOKBodyItems0IPV6 IPv6 configuration on the VLAN
swagger:model GetNetworkApplianceVlansOKBodyItems0IPV6
*/
type GetNetworkApplianceVlansOKBodyItems0IPV6 struct {

	// Enable IPv6 on VLAN
	Enabled bool `json:"enabled,omitempty"`

	// Prefix assignments on the VLAN
	PrefixAssignments []*GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0 `json:"prefixAssignments"`
}

// Validate validates this get network appliance vlans o k body items0 IP v6
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrefixAssignments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0IPV6) validatePrefixAssignments(formats strfmt.Registry) error {
	if swag.IsZero(o.PrefixAssignments) { // not required
		return nil
	}

	for i := 0; i < len(o.PrefixAssignments); i++ {
		if swag.IsZero(o.PrefixAssignments[i]) { // not required
			continue
		}

		if o.PrefixAssignments[i] != nil {
			if err := o.PrefixAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network appliance vlans o k body items0 IP v6 based on the context it is used
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrefixAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0IPV6) contextValidatePrefixAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PrefixAssignments); i++ {

		if o.PrefixAssignments[i] != nil {

			if swag.IsZero(o.PrefixAssignments[i]) { // not required
				return nil
			}

			if err := o.PrefixAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceVlansOKBodyItems0IPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0 get network appliance vlans o k body items0 IP v6 prefix assignments items0
swagger:model GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0
*/
type GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0 struct {

	// Auto assign a /64 prefix from the origin to the VLAN
	Autonomous bool `json:"autonomous,omitempty"`

	// origin
	Origin *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin `json:"origin,omitempty"`

	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 string `json:"staticApplianceIp6,omitempty"`

	// Manual configuration of a /64 prefix on the VLAN
	StaticPrefix string `json:"staticPrefix,omitempty"`
}

// Validate validates this get network appliance vlans o k body items0 IP v6 prefix assignments items0
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(o.Origin) { // not required
		return nil
	}

	if o.Origin != nil {
		if err := o.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network appliance vlans o k body items0 IP v6 prefix assignments items0 based on the context it is used
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if o.Origin != nil {

		if swag.IsZero(o.Origin) { // not required
			return nil
		}

		if err := o.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin The origin of the prefix
swagger:model GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin
*/
type GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin struct {

	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces"`

	// Type of the origin
	// Enum: [independent internet]
	Type string `json:"type,omitempty"`
}

// Validate validates this get network appliance vlans o k body items0 IP v6 prefix assignments items0 origin
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkApplianceVlansOKBodyItems0IpV6PrefixAssignmentsItems0OriginTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["independent","internet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkApplianceVlansOKBodyItems0IpV6PrefixAssignmentsItems0OriginTypeTypePropEnum = append(getNetworkApplianceVlansOKBodyItems0IpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, v)
	}
}

const (

	// GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0OriginTypeIndependent captures enum value "independent"
	GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0OriginTypeIndependent string = "independent"

	// GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0OriginTypeInternet captures enum value "internet"
	GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0OriginTypeInternet string = "internet"
)

// prop value enum
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkApplianceVlansOKBodyItems0IpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("origin"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network appliance vlans o k body items0 IP v6 prefix assignments items0 origin based on context it is used
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceVlansOKBodyItems0IPV6PrefixAssignmentsItems0Origin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp Mandatory DHCP will enforce that clients connecting to this VLAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
swagger:model GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp
*/
type GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp struct {

	// Enable Mandatory DHCP on VLAN.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this get network appliance vlans o k body items0 mandatory dhcp
func (o *GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network appliance vlans o k body items0 mandatory dhcp based on context it is used
func (o *GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceVlansOKBodyItems0MandatoryDhcp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0 get network appliance vlans o k body items0 reserved IP ranges items0
swagger:model GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0
*/
type GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0 struct {

	// A text comment for the reserved range
	Comment string `json:"comment,omitempty"`

	// The last IP in the reserved range
	End string `json:"end,omitempty"`

	// The first IP in the reserved range
	Start string `json:"start,omitempty"`
}

// Validate validates this get network appliance vlans o k body items0 reserved IP ranges items0
func (o *GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network appliance vlans o k body items0 reserved IP ranges items0 based on context it is used
func (o *GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkApplianceVlansOKBodyItems0ReservedIPRangesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
