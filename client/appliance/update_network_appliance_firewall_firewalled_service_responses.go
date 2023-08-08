// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// UpdateNetworkApplianceFirewallFirewalledServiceReader is a Reader for the UpdateNetworkApplianceFirewallFirewalledService structure.
type UpdateNetworkApplianceFirewallFirewalledServiceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceFirewallFirewalledServiceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceFirewallFirewalledServiceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/appliance/firewall/firewalledServices/{service}] updateNetworkApplianceFirewallFirewalledService", response, response.Code())
	}
}

// NewUpdateNetworkApplianceFirewallFirewalledServiceOK creates a UpdateNetworkApplianceFirewallFirewalledServiceOK with default headers values
func NewUpdateNetworkApplianceFirewallFirewalledServiceOK() *UpdateNetworkApplianceFirewallFirewalledServiceOK {
	return &UpdateNetworkApplianceFirewallFirewalledServiceOK{}
}

/*
UpdateNetworkApplianceFirewallFirewalledServiceOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceFirewallFirewalledServiceOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance firewall firewalled service o k response has a 2xx status code
func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance firewall firewalled service o k response has a 3xx status code
func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance firewall firewalled service o k response has a 4xx status code
func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance firewall firewalled service o k response has a 5xx status code
func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance firewall firewalled service o k response a status code equal to that given
func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network appliance firewall firewalled service o k response
func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) Code() int {
	return 200
}

func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/firewalledServices/{service}][%d] updateNetworkApplianceFirewallFirewalledServiceOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/firewall/firewalledServices/{service}][%d] updateNetworkApplianceFirewallFirewalledServiceOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceFirewallFirewalledServiceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkApplianceFirewallFirewalledServiceBody update network appliance firewall firewalled service body
// Example: {"access":"restricted","allowedIps":["123.123.123.1"]}
swagger:model UpdateNetworkApplianceFirewallFirewalledServiceBody
*/
type UpdateNetworkApplianceFirewallFirewalledServiceBody struct {

	// A string indicating the rule for which IPs are allowed to use the specified service. Acceptable values are "blocked" (no remote IPs can access the service), "restricted" (only allowed IPs can access the service), and "unrestriced" (any remote IP can access the service). This field is required
	// Required: true
	// Enum: [blocked restricted unrestricted]
	Access *string `json:"access"`

	// An array of allowed IPs that can access the service. This field is required if "access" is set to "restricted". Otherwise this field is ignored
	AllowedIps []string `json:"allowedIps"`
}

// Validate validates this update network appliance firewall firewalled service body
func (o *UpdateNetworkApplianceFirewallFirewalledServiceBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccess(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateNetworkApplianceFirewallFirewalledServiceBodyTypeAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["blocked","restricted","unrestricted"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkApplianceFirewallFirewalledServiceBodyTypeAccessPropEnum = append(updateNetworkApplianceFirewallFirewalledServiceBodyTypeAccessPropEnum, v)
	}
}

const (

	// UpdateNetworkApplianceFirewallFirewalledServiceBodyAccessBlocked captures enum value "blocked"
	UpdateNetworkApplianceFirewallFirewalledServiceBodyAccessBlocked string = "blocked"

	// UpdateNetworkApplianceFirewallFirewalledServiceBodyAccessRestricted captures enum value "restricted"
	UpdateNetworkApplianceFirewallFirewalledServiceBodyAccessRestricted string = "restricted"

	// UpdateNetworkApplianceFirewallFirewalledServiceBodyAccessUnrestricted captures enum value "unrestricted"
	UpdateNetworkApplianceFirewallFirewalledServiceBodyAccessUnrestricted string = "unrestricted"
)

// prop value enum
func (o *UpdateNetworkApplianceFirewallFirewalledServiceBody) validateAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkApplianceFirewallFirewalledServiceBodyTypeAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkApplianceFirewallFirewalledServiceBody) validateAccess(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkApplianceFirewallFirewalledService"+"."+"access", "body", o.Access); err != nil {
		return err
	}

	// value enum
	if err := o.validateAccessEnum("updateNetworkApplianceFirewallFirewalledService"+"."+"access", "body", *o.Access); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance firewall firewalled service body based on context it is used
func (o *UpdateNetworkApplianceFirewallFirewalledServiceBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallFirewalledServiceBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceFirewallFirewalledServiceBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceFirewallFirewalledServiceBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
