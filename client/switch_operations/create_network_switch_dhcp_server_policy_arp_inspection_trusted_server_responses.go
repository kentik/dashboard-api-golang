// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerReader is a Reader for the CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer structure.
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated creates a CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated with default headers values
func NewCreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated() *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated {
	return &CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated{}
}

/* CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated struct {
	Payload *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody
}

// IsSuccess returns true when this create network switch dhcp server policy arp inspection trusted server created response has a 2xx status code
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network switch dhcp server policy arp inspection trusted server created response has a 3xx status code
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network switch dhcp server policy arp inspection trusted server created response has a 4xx status code
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network switch dhcp server policy arp inspection trusted server created response has a 5xx status code
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network switch dhcp server policy arp inspection trusted server created response a status code equal to that given
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers][%d] createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/dhcpServerPolicy/arpInspection/trustedServers][%d] createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) GetPayload() *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody {
	return o.Payload
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody create network switch dhcp server policy arp inspection trusted server body
// Example: {"ipv4":{"address":"1.2.3.4"},"mac":"00:11:22:33:44:55","vlan":100}
swagger:model CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody
*/
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody struct {

	// ipv4
	// Required: true
	IPV4 *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4 `json:"ipv4"`

	// The mac address of the trusted server being added
	// Required: true
	Mac *string `json:"mac"`

	// The VLAN of the trusted server being added. It must be between 1 and 4094
	// Required: true
	Vlan *int64 `json:"vlan"`
}

// Validate validates this create network switch dhcp server policy arp inspection trusted server body
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV4(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateVlan(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) validateIPV4(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer"+"."+"ipv4", "body", o.IPV4); err != nil {
		return err
	}

	if o.IPV4 != nil {
		if err := o.IPV4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer" + "." + "ipv4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer" + "." + "ipv4")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) validateMac(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer"+"."+"mac", "body", o.Mac); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) validateVlan(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer"+"."+"vlan", "body", o.Vlan); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network switch dhcp server policy arp inspection trusted server body based on the context it is used
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) contextValidateIPV4(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV4 != nil {
		if err := o.IPV4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer" + "." + "ipv4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServer" + "." + "ipv4")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody create network switch dhcp server policy arp inspection trusted server created body
swagger:model CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody
*/
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody struct {

	// ipv4
	IPV4 *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4 `json:"ipv4,omitempty"`

	// Mac address of the trusted server.
	Mac string `json:"mac,omitempty"`

	// ID of the trusted server.
	TrustedServerID string `json:"trustedServerId,omitempty"`

	// Vlan ID of the trusted server.
	Vlan int64 `json:"vlan,omitempty"`
}

// Validate validates this create network switch dhcp server policy arp inspection trusted server created body
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPV4(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody) validateIPV4(formats strfmt.Registry) error {
	if swag.IsZero(o.IPV4) { // not required
		return nil
	}

	if o.IPV4 != nil {
		if err := o.IPV4.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated" + "." + "ipv4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated" + "." + "ipv4")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network switch dhcp server policy arp inspection trusted server created body based on the context it is used
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPV4(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody) contextValidateIPV4(ctx context.Context, formats strfmt.Registry) error {

	if o.IPV4 != nil {
		if err := o.IPV4.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated" + "." + "ipv4")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreated" + "." + "ipv4")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4 IPv4 attributes of the trusted server.
swagger:model CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4
*/
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4 struct {

	// IPv4 address of the trusted server.
	Address string `json:"address,omitempty"`
}

// Validate validates this create network switch dhcp server policy arp inspection trusted server created body IP v4
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch dhcp server policy arp inspection trusted server created body IP v4 based on context it is used
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerCreatedBodyIPV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4 The IPv4 attributes of the trusted server being added
swagger:model CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4
*/
type CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4 struct {

	// The IPv4 address of the trusted server being added
	Address string `json:"address,omitempty"`
}

// Validate validates this create network switch dhcp server policy arp inspection trusted server params body IP v4
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch dhcp server policy arp inspection trusted server params body IP v4 based on context it is used
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchDhcpServerPolicyArpInspectionTrustedServerParamsBodyIPV4
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
