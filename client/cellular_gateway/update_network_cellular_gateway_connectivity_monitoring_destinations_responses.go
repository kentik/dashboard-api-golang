// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

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

// UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsReader is a Reader for the UpdateNetworkCellularGatewayConnectivityMonitoringDestinations structure.
type UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations] updateNetworkCellularGatewayConnectivityMonitoringDestinations", response, response.Code())
	}
}

// NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK creates a UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK with default headers values
func NewUpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK() *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK {
	return &UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK{}
}

/*
UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network cellular gateway connectivity monitoring destinations o k response has a 2xx status code
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network cellular gateway connectivity monitoring destinations o k response has a 3xx status code
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network cellular gateway connectivity monitoring destinations o k response has a 4xx status code
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network cellular gateway connectivity monitoring destinations o k response has a 5xx status code
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network cellular gateway connectivity monitoring destinations o k response a status code equal to that given
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network cellular gateway connectivity monitoring destinations o k response
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) Code() int {
	return 200
}

func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations][%d] updateNetworkCellularGatewayConnectivityMonitoringDestinationsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/cellularGateway/connectivityMonitoringDestinations][%d] updateNetworkCellularGatewayConnectivityMonitoringDestinationsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody update network cellular gateway connectivity monitoring destinations body
// Example: {"destinations":[{"default":false,"description":"Google","ip":"8.8.8.8"},{"default":true,"description":"test description","ip":"1.23.45.67"},{"ip":"9.8.7.6"}]}
swagger:model UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody
*/
type UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody struct {

	// The list of connectivity monitoring destinations
	Destinations []*UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0 `json:"destinations"`
}

// Validate validates this update network cellular gateway connectivity monitoring destinations body
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDestinations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody) validateDestinations(formats strfmt.Registry) error {
	if swag.IsZero(o.Destinations) { // not required
		return nil
	}

	for i := 0; i < len(o.Destinations); i++ {
		if swag.IsZero(o.Destinations[i]) { // not required
			continue
		}

		if o.Destinations[i] != nil {
			if err := o.Destinations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkCellularGatewayConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkCellularGatewayConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network cellular gateway connectivity monitoring destinations body based on the context it is used
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDestinations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody) contextValidateDestinations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Destinations); i++ {

		if o.Destinations[i] != nil {

			if swag.IsZero(o.Destinations[i]) { // not required
				return nil
			}

			if err := o.Destinations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkCellularGatewayConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkCellularGatewayConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0 update network cellular gateway connectivity monitoring destinations params body destinations items0
swagger:model UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0
*/
type UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0 struct {

	// Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Default *bool `json:"default,omitempty"`

	// Description of the testing destination. Optional, defaults to an empty string
	Description string `json:"description,omitempty"`

	// The IP address to test connectivity with
	// Required: true
	IP *string `json:"ip"`
}

// Validate validates this update network cellular gateway connectivity monitoring destinations params body destinations items0
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", o.IP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network cellular gateway connectivity monitoring destinations params body destinations items0 based on context it is used
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkCellularGatewayConnectivityMonitoringDestinationsParamsBodyDestinationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
