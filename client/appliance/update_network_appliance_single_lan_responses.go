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

// UpdateNetworkApplianceSingleLanReader is a Reader for the UpdateNetworkApplianceSingleLan structure.
type UpdateNetworkApplianceSingleLanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceSingleLanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceSingleLanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/appliance/singleLan] updateNetworkApplianceSingleLan", response, response.Code())
	}
}

// NewUpdateNetworkApplianceSingleLanOK creates a UpdateNetworkApplianceSingleLanOK with default headers values
func NewUpdateNetworkApplianceSingleLanOK() *UpdateNetworkApplianceSingleLanOK {
	return &UpdateNetworkApplianceSingleLanOK{}
}

/*
UpdateNetworkApplianceSingleLanOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceSingleLanOK struct {
	Payload *UpdateNetworkApplianceSingleLanOKBody
}

// IsSuccess returns true when this update network appliance single lan o k response has a 2xx status code
func (o *UpdateNetworkApplianceSingleLanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance single lan o k response has a 3xx status code
func (o *UpdateNetworkApplianceSingleLanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance single lan o k response has a 4xx status code
func (o *UpdateNetworkApplianceSingleLanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance single lan o k response has a 5xx status code
func (o *UpdateNetworkApplianceSingleLanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance single lan o k response a status code equal to that given
func (o *UpdateNetworkApplianceSingleLanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network appliance single lan o k response
func (o *UpdateNetworkApplianceSingleLanOK) Code() int {
	return 200
}

func (o *UpdateNetworkApplianceSingleLanOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/singleLan][%d] updateNetworkApplianceSingleLanOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceSingleLanOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/singleLan][%d] updateNetworkApplianceSingleLanOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceSingleLanOK) GetPayload() *UpdateNetworkApplianceSingleLanOKBody {
	return o.Payload
}

func (o *UpdateNetworkApplianceSingleLanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateNetworkApplianceSingleLanOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkApplianceSingleLanBody update network appliance single lan body
// Example: {}
swagger:model UpdateNetworkApplianceSingleLanBody
*/
type UpdateNetworkApplianceSingleLanBody struct {

	// The appliance IP address of the single LAN
	ApplianceIP string `json:"applianceIp,omitempty"`

	// ipv6
	IPV6 *UpdateNetworkApplianceSingleLanParamsBodyIPV6 `json:"ipv6,omitempty"`

	// mandatory dhcp
	MandatoryDhcp *UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp `json:"mandatoryDhcp,omitempty"`

	// The subnet of the single LAN configuration
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this update network appliance single lan body
func (o *UpdateNetworkApplianceSingleLanBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMandatoryDhcp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanBody) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceSingleLanBody) validateMandatoryDhcp(formats strfmt.Registry) error {
	if swag.IsZero(o.MandatoryDhcp) { // not required
		return nil
	}

	if o.MandatoryDhcp != nil {
		if err := o.MandatoryDhcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLan" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLan" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance single lan body based on the context it is used
func (o *UpdateNetworkApplianceSingleLanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMandatoryDhcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanBody) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {

		if swag.IsZero(o.IPV6) { // not required
			return nil
		}

		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceSingleLanBody) contextValidateMandatoryDhcp(ctx context.Context, formats strfmt.Registry) error {

	if o.MandatoryDhcp != nil {

		if swag.IsZero(o.MandatoryDhcp) { // not required
			return nil
		}

		if err := o.MandatoryDhcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLan" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLan" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanOKBody update network appliance single lan o k body
swagger:model UpdateNetworkApplianceSingleLanOKBody
*/
type UpdateNetworkApplianceSingleLanOKBody struct {

	// The local IP of the appliance on the single LAN
	ApplianceIP string `json:"applianceIp,omitempty"`

	// ipv6
	IPV6 *UpdateNetworkApplianceSingleLanOKBodyIPV6 `json:"ipv6,omitempty"`

	// mandatory dhcp
	MandatoryDhcp *UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp `json:"mandatoryDhcp,omitempty"`

	// The subnet of the single LAN
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this update network appliance single lan o k body
func (o *UpdateNetworkApplianceSingleLanOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV6(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMandatoryDhcp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBody) validateIPV6(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV6) { // not required
		return nil
	}

	if o.IPV6 != nil {
		if err := o.IPV6.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBody) validateMandatoryDhcp(formats strfmt.Registry) error {
	if swag.IsZero(o.MandatoryDhcp) { // not required
		return nil
	}

	if o.MandatoryDhcp != nil {
		if err := o.MandatoryDhcp.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance single lan o k body based on the context it is used
func (o *UpdateNetworkApplianceSingleLanOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV6(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMandatoryDhcp(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBody) contextValidateIPV6(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV6 != nil {

		if swag.IsZero(o.IPV6) { // not required
			return nil
		}

		if err := o.IPV6.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBody) contextValidateMandatoryDhcp(ctx context.Context, formats strfmt.Registry) error {

	if o.MandatoryDhcp != nil {

		if swag.IsZero(o.MandatoryDhcp) { // not required
			return nil
		}

		if err := o.MandatoryDhcp.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "mandatoryDhcp")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "mandatoryDhcp")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanOKBodyIPV6 IPv6 configuration on the single LAN
swagger:model UpdateNetworkApplianceSingleLanOKBodyIPV6
*/
type UpdateNetworkApplianceSingleLanOKBodyIPV6 struct {

	// Enable IPv6 on single LAN
	Enabled bool `json:"enabled,omitempty"`

	// Prefix assignments on the single LAN
	PrefixAssignments []*UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0 `json:"prefixAssignments"`
}

// Validate validates this update network appliance single lan o k body IP v6
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrefixAssignments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6) validatePrefixAssignments(formats strfmt.Registry) error {
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
					return ve.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance single lan o k body IP v6 based on the context it is used
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrefixAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6) contextValidatePrefixAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PrefixAssignments); i++ {

		if o.PrefixAssignments[i] != nil {

			if swag.IsZero(o.PrefixAssignments[i]) { // not required
				return nil
			}

			if err := o.PrefixAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceSingleLanOK" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanOKBodyIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0 update network appliance single lan o k body IP v6 prefix assignments items0
swagger:model UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0
*/
type UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0 struct {

	// Auto assign a /64 prefix from the origin to the single LAN
	Autonomous bool `json:"autonomous,omitempty"`

	// origin
	Origin *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin `json:"origin,omitempty"`

	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 string `json:"staticApplianceIp6,omitempty"`

	// Manual configuration of a /64 prefix on the single LAN
	StaticPrefix string `json:"staticPrefix,omitempty"`
}

// Validate validates this update network appliance single lan o k body IP v6 prefix assignments items0
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0) validateOrigin(formats strfmt.Registry) error {
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

// ContextValidate validate this update network appliance single lan o k body IP v6 prefix assignments items0 based on the context it is used
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin The origin of the prefix
swagger:model UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin
*/
type UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin struct {

	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces"`

	// Type of the origin
	// Enum: [independent internet]
	Type string `json:"type,omitempty"`
}

// Validate validates this update network appliance single lan o k body IP v6 prefix assignments items0 origin
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkApplianceSingleLanOKBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["independent","internet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceSingleLanOKBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum = append(updateNetworkApplianceSingleLanOKBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent captures enum value "independent"
	UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent string = "independent"

	// UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0OriginTypeInternet captures enum value "internet"
	UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0OriginTypeInternet string = "internet"
)

// prop value enum
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceSingleLanOKBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("origin"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance single lan o k body IP v6 prefix assignments items0 origin based on context it is used
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanOKBodyIPV6PrefixAssignmentsItems0Origin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp Mandatory DHCP will enforce that clients connecting to this single LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
swagger:model UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp
*/
type UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp struct {

	// Enable Mandatory DHCP on single LAN.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network appliance single lan o k body mandatory dhcp
func (o *UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance single lan o k body mandatory dhcp based on context it is used
func (o *UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanOKBodyMandatoryDhcp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanParamsBodyIPV6 IPv6 configuration on the VLAN
swagger:model UpdateNetworkApplianceSingleLanParamsBodyIPV6
*/
type UpdateNetworkApplianceSingleLanParamsBodyIPV6 struct {

	// Enable IPv6 on VLAN.
	Enabled bool `json:"enabled,omitempty"`

	// Prefix assignments on the VLAN
	PrefixAssignments []*UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0 `json:"prefixAssignments"`
}

// Validate validates this update network appliance single lan params body IP v6
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePrefixAssignments(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6) validatePrefixAssignments(formats strfmt.Registry) error {
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
					return ve.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance single lan params body IP v6 based on the context it is used
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePrefixAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6) contextValidatePrefixAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PrefixAssignments); i++ {

		if o.PrefixAssignments[i] != nil {

			if swag.IsZero(o.PrefixAssignments[i]) { // not required
				return nil
			}

			if err := o.PrefixAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceSingleLan" + "." + "ipv6" + "." + "prefixAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanParamsBodyIPV6
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0 update network appliance single lan params body IP v6 prefix assignments items0
swagger:model UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0
*/
type UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0 struct {

	// Auto assign a /64 prefix from the origin to the VLAN
	Autonomous bool `json:"autonomous,omitempty"`

	// origin
	Origin *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin `json:"origin,omitempty"`

	// Manual configuration of the IPv6 Appliance IP
	StaticApplianceIp6 string `json:"staticApplianceIp6,omitempty"`

	// Manual configuration of a /64 prefix on the VLAN
	StaticPrefix string `json:"staticPrefix,omitempty"`
}

// Validate validates this update network appliance single lan params body IP v6 prefix assignments items0
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0) validateOrigin(formats strfmt.Registry) error {
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

// ContextValidate validate this update network appliance single lan params body IP v6 prefix assignments items0 based on the context it is used
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

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
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin The origin of the prefix
swagger:model UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin
*/
type UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin struct {

	// Interfaces associated with the prefix
	Interfaces []string `json:"interfaces"`

	// Type of the origin
	// Required: true
	// Enum: [independent internet]
	Type *string `json:"type"`
}

// Validate validates this update network appliance single lan params body IP v6 prefix assignments items0 origin
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkApplianceSingleLanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["independent","internet"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceSingleLanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum = append(updateNetworkApplianceSingleLanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent captures enum value "independent"
	UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeIndependent string = "independent"

	// UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeInternet captures enum value "internet"
	UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0OriginTypeInternet string = "internet"
)

// prop value enum
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceSingleLanParamsBodyIpV6PrefixAssignmentsItems0OriginTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin) validateType(formats strfmt.Registry) error {

	if err := validate.Required("origin"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("origin"+"."+"type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance single lan params body IP v6 prefix assignments items0 origin based on context it is used
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanParamsBodyIPV6PrefixAssignmentsItems0Origin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp Mandatory DHCP will enforce that clients connecting to this LAN must use the IP address assigned by the DHCP server. Clients who use a static IP address won't be able to associate. Only available on firmware versions 17.0 and above
swagger:model UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp
*/
type UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp struct {

	// Enable Mandatory DHCP on LAN.
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network appliance single lan params body mandatory dhcp
func (o *UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance single lan params body mandatory dhcp based on context it is used
func (o *UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceSingleLanParamsBodyMandatoryDhcp
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
