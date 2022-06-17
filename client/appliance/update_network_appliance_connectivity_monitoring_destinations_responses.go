// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// UpdateNetworkApplianceConnectivityMonitoringDestinationsReader is a Reader for the UpdateNetworkApplianceConnectivityMonitoringDestinations structure.
type UpdateNetworkApplianceConnectivityMonitoringDestinationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceConnectivityMonitoringDestinationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceConnectivityMonitoringDestinationsOK creates a UpdateNetworkApplianceConnectivityMonitoringDestinationsOK with default headers values
func NewUpdateNetworkApplianceConnectivityMonitoringDestinationsOK() *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK {
	return &UpdateNetworkApplianceConnectivityMonitoringDestinationsOK{}
}

/* UpdateNetworkApplianceConnectivityMonitoringDestinationsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceConnectivityMonitoringDestinationsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance connectivity monitoring destinations o k response has a 2xx status code
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance connectivity monitoring destinations o k response has a 3xx status code
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance connectivity monitoring destinations o k response has a 4xx status code
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance connectivity monitoring destinations o k response has a 5xx status code
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance connectivity monitoring destinations o k response a status code equal to that given
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/connectivityMonitoringDestinations][%d] updateNetworkApplianceConnectivityMonitoringDestinationsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/connectivityMonitoringDestinations][%d] updateNetworkApplianceConnectivityMonitoringDestinationsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkApplianceConnectivityMonitoringDestinationsBody update network appliance connectivity monitoring destinations body
// Example: {"destinations":[{"default":false,"description":"Google","ip":"8.8.8.8"},{"default":true,"description":"test description","ip":"1.23.45.67"},{"ip":"9.8.7.6"}]}
swagger:model UpdateNetworkApplianceConnectivityMonitoringDestinationsBody
*/
type UpdateNetworkApplianceConnectivityMonitoringDestinationsBody struct {

	// The list of connectivity monitoring destinations
	Destinations []*UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0 `json:"destinations"`
}

// Validate validates this update network appliance connectivity monitoring destinations body
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDestinations(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsBody) validateDestinations(formats strfmt.Registry) error {
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
					return ve.ValidateName("updateNetworkApplianceConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network appliance connectivity monitoring destinations body based on the context it is used
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDestinations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsBody) contextValidateDestinations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Destinations); i++ {

		if o.Destinations[i] != nil {
			if err := o.Destinations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkApplianceConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkApplianceConnectivityMonitoringDestinations" + "." + "destinations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceConnectivityMonitoringDestinationsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0 update network appliance connectivity monitoring destinations params body destinations items0
swagger:model UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0
*/
type UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0 struct {

	// Boolean indicating whether this is the default testing destination (true) or not (false). Defaults to false. Only one default is allowed
	Default bool `json:"default,omitempty"`

	// Description of the testing destination. Optional, defaults to null
	Description string `json:"description,omitempty"`

	// The IP address to test connectivity with
	// Required: true
	IP *string `json:"ip"`
}

// Validate validates this update network appliance connectivity monitoring destinations params body destinations items0
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", o.IP); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network appliance connectivity monitoring destinations params body destinations items0 based on context it is used
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceConnectivityMonitoringDestinationsParamsBodyDestinationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
