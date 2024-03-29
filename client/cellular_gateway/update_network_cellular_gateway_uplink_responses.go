// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

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

// UpdateNetworkCellularGatewayUplinkReader is a Reader for the UpdateNetworkCellularGatewayUplink structure.
type UpdateNetworkCellularGatewayUplinkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkCellularGatewayUplinkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkCellularGatewayUplinkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/cellularGateway/uplink] updateNetworkCellularGatewayUplink", response, response.Code())
	}
}

// NewUpdateNetworkCellularGatewayUplinkOK creates a UpdateNetworkCellularGatewayUplinkOK with default headers values
func NewUpdateNetworkCellularGatewayUplinkOK() *UpdateNetworkCellularGatewayUplinkOK {
	return &UpdateNetworkCellularGatewayUplinkOK{}
}

/*
UpdateNetworkCellularGatewayUplinkOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkCellularGatewayUplinkOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network cellular gateway uplink o k response has a 2xx status code
func (o *UpdateNetworkCellularGatewayUplinkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network cellular gateway uplink o k response has a 3xx status code
func (o *UpdateNetworkCellularGatewayUplinkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network cellular gateway uplink o k response has a 4xx status code
func (o *UpdateNetworkCellularGatewayUplinkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network cellular gateway uplink o k response has a 5xx status code
func (o *UpdateNetworkCellularGatewayUplinkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network cellular gateway uplink o k response a status code equal to that given
func (o *UpdateNetworkCellularGatewayUplinkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network cellular gateway uplink o k response
func (o *UpdateNetworkCellularGatewayUplinkOK) Code() int {
	return 200
}

func (o *UpdateNetworkCellularGatewayUplinkOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/cellularGateway/uplink][%d] updateNetworkCellularGatewayUplinkOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkCellularGatewayUplinkOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/cellularGateway/uplink][%d] updateNetworkCellularGatewayUplinkOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkCellularGatewayUplinkOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkCellularGatewayUplinkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkCellularGatewayUplinkBody update network cellular gateway uplink body
// Example: {"bandwidthLimits":{"limitDown":1000,"limitUp":1000}}
swagger:model UpdateNetworkCellularGatewayUplinkBody
*/
type UpdateNetworkCellularGatewayUplinkBody struct {

	// bandwidth limits
	BandwidthLimits *UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits `json:"bandwidthLimits,omitempty"`
}

// Validate validates this update network cellular gateway uplink body
func (o *UpdateNetworkCellularGatewayUplinkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBandwidthLimits(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkCellularGatewayUplinkBody) validateBandwidthLimits(formats strfmt.Registry) error {
	if swag.IsZero(o.BandwidthLimits) { // not required
		return nil
	}

	if o.BandwidthLimits != nil {
		if err := o.BandwidthLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkCellularGatewayUplink" + "." + "bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkCellularGatewayUplink" + "." + "bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network cellular gateway uplink body based on the context it is used
func (o *UpdateNetworkCellularGatewayUplinkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBandwidthLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkCellularGatewayUplinkBody) contextValidateBandwidthLimits(ctx context.Context, formats strfmt.Registry) error {

	if o.BandwidthLimits != nil {

		if swag.IsZero(o.BandwidthLimits) { // not required
			return nil
		}

		if err := o.BandwidthLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkCellularGatewayUplink" + "." + "bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkCellularGatewayUplink" + "." + "bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayUplinkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayUplinkBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkCellularGatewayUplinkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits The bandwidth settings for the 'cellular' uplink
swagger:model UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits
*/
type UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits struct {

	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown int64 `json:"limitDown,omitempty"`

	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp int64 `json:"limitUp,omitempty"`
}

// Validate validates this update network cellular gateway uplink params body bandwidth limits
func (o *UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network cellular gateway uplink params body bandwidth limits based on context it is used
func (o *UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkCellularGatewayUplinkParamsBodyBandwidthLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
