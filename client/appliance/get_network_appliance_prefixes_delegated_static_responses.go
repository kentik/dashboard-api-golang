// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// GetNetworkAppliancePrefixesDelegatedStaticReader is a Reader for the GetNetworkAppliancePrefixesDelegatedStatic structure.
type GetNetworkAppliancePrefixesDelegatedStaticReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAppliancePrefixesDelegatedStaticReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkAppliancePrefixesDelegatedStaticOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}] getNetworkAppliancePrefixesDelegatedStatic", response, response.Code())
	}
}

// NewGetNetworkAppliancePrefixesDelegatedStaticOK creates a GetNetworkAppliancePrefixesDelegatedStaticOK with default headers values
func NewGetNetworkAppliancePrefixesDelegatedStaticOK() *GetNetworkAppliancePrefixesDelegatedStaticOK {
	return &GetNetworkAppliancePrefixesDelegatedStaticOK{}
}

/*
GetNetworkAppliancePrefixesDelegatedStaticOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkAppliancePrefixesDelegatedStaticOK struct {
	Payload *GetNetworkAppliancePrefixesDelegatedStaticOKBody
}

// IsSuccess returns true when this get network appliance prefixes delegated static o k response has a 2xx status code
func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance prefixes delegated static o k response has a 3xx status code
func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance prefixes delegated static o k response has a 4xx status code
func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance prefixes delegated static o k response has a 5xx status code
func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance prefixes delegated static o k response a status code equal to that given
func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network appliance prefixes delegated static o k response
func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) Code() int {
	return 200
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}][%d] getNetworkAppliancePrefixesDelegatedStaticOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/prefixes/delegated/statics/{staticDelegatedPrefixId}][%d] getNetworkAppliancePrefixesDelegatedStaticOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) GetPayload() *GetNetworkAppliancePrefixesDelegatedStaticOKBody {
	return o.Payload
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkAppliancePrefixesDelegatedStaticOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkAppliancePrefixesDelegatedStaticOKBody get network appliance prefixes delegated static o k body
swagger:model GetNetworkAppliancePrefixesDelegatedStaticOKBody
*/
type GetNetworkAppliancePrefixesDelegatedStaticOKBody struct {

	// Prefix creation time.
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Identifying description for the prefix.
	Description string `json:"description,omitempty"`

	// origin
	Origin *GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin `json:"origin,omitempty"`

	// IPv6 prefix/prefix length.
	Prefix string `json:"prefix,omitempty"`

	// Static delegated prefix id.
	StaticDelegatedPrefixID string `json:"staticDelegatedPrefixId,omitempty"`

	// Prefix Updated time.
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`
}

// Validate validates this get network appliance prefixes delegated static o k body
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("getNetworkAppliancePrefixesDelegatedStaticOK"+"."+"createdAt", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) validateOrigin(formats strfmt.Registry) error {
	if swag.IsZero(o.Origin) { // not required
		return nil
	}

	if o.Origin != nil {
		if err := o.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkAppliancePrefixesDelegatedStaticOK" + "." + "origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkAppliancePrefixesDelegatedStaticOK" + "." + "origin")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) validateUpdatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("getNetworkAppliancePrefixesDelegatedStaticOK"+"."+"updatedAt", "body", "date-time", o.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get network appliance prefixes delegated static o k body based on the context it is used
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOrigin(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) contextValidateOrigin(ctx context.Context, formats strfmt.Registry) error {

	if o.Origin != nil {

		if swag.IsZero(o.Origin) { // not required
			return nil
		}

		if err := o.Origin.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkAppliancePrefixesDelegatedStaticOK" + "." + "origin")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkAppliancePrefixesDelegatedStaticOK" + "." + "origin")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkAppliancePrefixesDelegatedStaticOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin WAN1/WAN2/Independent prefix.
swagger:model GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin
*/
type GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin struct {

	// Uplink provided or independent
	Interfaces []string `json:"interfaces"`

	// Origin type
	Type string `json:"type,omitempty"`
}

// Validate validates this get network appliance prefixes delegated static o k body origin
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network appliance prefixes delegated static o k body origin based on context it is used
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin) UnmarshalBinary(b []byte) error {
	var res GetNetworkAppliancePrefixesDelegatedStaticOKBodyOrigin
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
