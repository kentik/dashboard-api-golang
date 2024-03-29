// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkSwitchMtuReader is a Reader for the UpdateNetworkSwitchMtu structure.
type UpdateNetworkSwitchMtuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkSwitchMtuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkSwitchMtuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/switch/mtu] updateNetworkSwitchMtu", response, response.Code())
	}
}

// NewUpdateNetworkSwitchMtuOK creates a UpdateNetworkSwitchMtuOK with default headers values
func NewUpdateNetworkSwitchMtuOK() *UpdateNetworkSwitchMtuOK {
	return &UpdateNetworkSwitchMtuOK{}
}

/*
UpdateNetworkSwitchMtuOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkSwitchMtuOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network switch mtu o k response has a 2xx status code
func (o *UpdateNetworkSwitchMtuOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network switch mtu o k response has a 3xx status code
func (o *UpdateNetworkSwitchMtuOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network switch mtu o k response has a 4xx status code
func (o *UpdateNetworkSwitchMtuOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network switch mtu o k response has a 5xx status code
func (o *UpdateNetworkSwitchMtuOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network switch mtu o k response a status code equal to that given
func (o *UpdateNetworkSwitchMtuOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network switch mtu o k response
func (o *UpdateNetworkSwitchMtuOK) Code() int {
	return 200
}

func (o *UpdateNetworkSwitchMtuOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/mtu][%d] updateNetworkSwitchMtuOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchMtuOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/switch/mtu][%d] updateNetworkSwitchMtuOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkSwitchMtuOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkSwitchMtuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkSwitchMtuBody update network switch mtu body
// Example: {"defaultMtuSize":9578,"overrides":[{"mtuSize":1500,"switches":["Q234-ABCD-0001","Q234-ABCD-0002","Q234-ABCD-0003"]},{"mtuSize":1600,"switchProfiles":["1284392014819","2983092129865"]}]}
swagger:model UpdateNetworkSwitchMtuBody
*/
type UpdateNetworkSwitchMtuBody struct {

	// MTU size for the entire network. Default value is 9578.
	DefaultMtuSize int64 `json:"defaultMtuSize,omitempty"`

	// Override MTU size for individual switches or switch templates. An empty array will clear overrides.
	Overrides []*UpdateNetworkSwitchMtuParamsBodyOverridesItems0 `json:"overrides"`
}

// Validate validates this update network switch mtu body
func (o *UpdateNetworkSwitchMtuBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOverrides(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchMtuBody) validateOverrides(formats strfmt.Registry) error {
	if swag.IsZero(o.Overrides) { // not required
		return nil
	}

	for i := 0; i < len(o.Overrides); i++ {
		if swag.IsZero(o.Overrides[i]) { // not required
			continue
		}

		if o.Overrides[i] != nil {
			if err := o.Overrides[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchMtu" + "." + "overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchMtu" + "." + "overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network switch mtu body based on the context it is used
func (o *UpdateNetworkSwitchMtuBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOverrides(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchMtuBody) contextValidateOverrides(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Overrides); i++ {

		if o.Overrides[i] != nil {

			if swag.IsZero(o.Overrides[i]) { // not required
				return nil
			}

			if err := o.Overrides[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkSwitchMtu" + "." + "overrides" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkSwitchMtu" + "." + "overrides" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchMtuBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchMtuBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchMtuBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkSwitchMtuParamsBodyOverridesItems0 update network switch mtu params body overrides items0
swagger:model UpdateNetworkSwitchMtuParamsBodyOverridesItems0
*/
type UpdateNetworkSwitchMtuParamsBodyOverridesItems0 struct {

	// MTU size for the switches or switch templates.
	// Required: true
	MtuSize *int64 `json:"mtuSize"`

	// List of switch template IDs. Applicable only for template network.
	SwitchProfiles []string `json:"switchProfiles"`

	// List of switch serials. Applicable only for switch network.
	Switches []string `json:"switches"`
}

// Validate validates this update network switch mtu params body overrides items0
func (o *UpdateNetworkSwitchMtuParamsBodyOverridesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMtuSize(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkSwitchMtuParamsBodyOverridesItems0) validateMtuSize(formats strfmt.Registry) error {

	if err := validate.Required("mtuSize", "body", o.MtuSize); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network switch mtu params body overrides items0 based on context it is used
func (o *UpdateNetworkSwitchMtuParamsBodyOverridesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkSwitchMtuParamsBodyOverridesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkSwitchMtuParamsBodyOverridesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkSwitchMtuParamsBodyOverridesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
