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
)

// UpdateNetworkApplianceTrafficShapingReader is a Reader for the UpdateNetworkApplianceTrafficShaping structure.
type UpdateNetworkApplianceTrafficShapingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceTrafficShapingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceTrafficShapingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/appliance/trafficShaping] updateNetworkApplianceTrafficShaping", response, response.Code())
	}
}

// NewUpdateNetworkApplianceTrafficShapingOK creates a UpdateNetworkApplianceTrafficShapingOK with default headers values
func NewUpdateNetworkApplianceTrafficShapingOK() *UpdateNetworkApplianceTrafficShapingOK {
	return &UpdateNetworkApplianceTrafficShapingOK{}
}

/*
UpdateNetworkApplianceTrafficShapingOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceTrafficShapingOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance traffic shaping o k response has a 2xx status code
func (o *UpdateNetworkApplianceTrafficShapingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance traffic shaping o k response has a 3xx status code
func (o *UpdateNetworkApplianceTrafficShapingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance traffic shaping o k response has a 4xx status code
func (o *UpdateNetworkApplianceTrafficShapingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance traffic shaping o k response has a 5xx status code
func (o *UpdateNetworkApplianceTrafficShapingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance traffic shaping o k response a status code equal to that given
func (o *UpdateNetworkApplianceTrafficShapingOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network appliance traffic shaping o k response
func (o *UpdateNetworkApplianceTrafficShapingOK) Code() int {
	return 200
}

func (o *UpdateNetworkApplianceTrafficShapingOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/trafficShaping][%d] updateNetworkApplianceTrafficShapingOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceTrafficShapingOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/trafficShaping][%d] updateNetworkApplianceTrafficShapingOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceTrafficShapingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceTrafficShapingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkApplianceTrafficShapingBody update network appliance traffic shaping body
// Example: {"globalBandwidthLimits":{"limitDown":5120,"limitUp":2048}}
swagger:model UpdateNetworkApplianceTrafficShapingBody
*/
type UpdateNetworkApplianceTrafficShapingBody struct {

	// global bandwidth limits
	GlobalBandwidthLimits *UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits `json:"globalBandwidthLimits,omitempty"`
}

// Validate validates this update network appliance traffic shaping body
func (o *UpdateNetworkApplianceTrafficShapingBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateGlobalBandwidthLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingBody) validateGlobalBandwidthLimits(formats strfmt.Registry) error {
	if swag.IsZero(o.GlobalBandwidthLimits) { // not required
		return nil
	}

	if o.GlobalBandwidthLimits != nil {
		if err := o.GlobalBandwidthLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShaping" + "." + "globalBandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShaping" + "." + "globalBandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network appliance traffic shaping body based on the context it is used
func (o *UpdateNetworkApplianceTrafficShapingBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateGlobalBandwidthLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceTrafficShapingBody) contextValidateGlobalBandwidthLimits(ctx context.Context, formats strfmt.Registry) error {

	if o.GlobalBandwidthLimits != nil {

		if swag.IsZero(o.GlobalBandwidthLimits) { // not required
			return nil
		}

		if err := o.GlobalBandwidthLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkApplianceTrafficShaping" + "." + "globalBandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkApplianceTrafficShaping" + "." + "globalBandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits Global per-client bandwidth limit
swagger:model UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits
*/
type UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits struct {

	// The download bandwidth limit in Kbps. (0 represents no limit.)
	LimitDown int64 `json:"limitDown,omitempty"`

	// The upload bandwidth limit in Kbps. (0 represents no limit.)
	LimitUp int64 `json:"limitUp,omitempty"`
}

// Validate validates this update network appliance traffic shaping params body global bandwidth limits
func (o *UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance traffic shaping params body global bandwidth limits based on context it is used
func (o *UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingParamsBodyGlobalBandwidthLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
