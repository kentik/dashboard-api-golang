// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// UpdateDeviceCameraCustomAnalyticsReader is a Reader for the UpdateDeviceCameraCustomAnalytics structure.
type UpdateDeviceCameraCustomAnalyticsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceCameraCustomAnalyticsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceCameraCustomAnalyticsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}/camera/customAnalytics] updateDeviceCameraCustomAnalytics", response, response.Code())
	}
}

// NewUpdateDeviceCameraCustomAnalyticsOK creates a UpdateDeviceCameraCustomAnalyticsOK with default headers values
func NewUpdateDeviceCameraCustomAnalyticsOK() *UpdateDeviceCameraCustomAnalyticsOK {
	return &UpdateDeviceCameraCustomAnalyticsOK{}
}

/*
UpdateDeviceCameraCustomAnalyticsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceCameraCustomAnalyticsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device camera custom analytics o k response has a 2xx status code
func (o *UpdateDeviceCameraCustomAnalyticsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device camera custom analytics o k response has a 3xx status code
func (o *UpdateDeviceCameraCustomAnalyticsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device camera custom analytics o k response has a 4xx status code
func (o *UpdateDeviceCameraCustomAnalyticsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device camera custom analytics o k response has a 5xx status code
func (o *UpdateDeviceCameraCustomAnalyticsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device camera custom analytics o k response a status code equal to that given
func (o *UpdateDeviceCameraCustomAnalyticsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device camera custom analytics o k response
func (o *UpdateDeviceCameraCustomAnalyticsOK) Code() int {
	return 200
}

func (o *UpdateDeviceCameraCustomAnalyticsOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/camera/customAnalytics][%d] updateDeviceCameraCustomAnalyticsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCameraCustomAnalyticsOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/camera/customAnalytics][%d] updateDeviceCameraCustomAnalyticsOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCameraCustomAnalyticsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceCameraCustomAnalyticsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceCameraCustomAnalyticsBody update device camera custom analytics body
// Example: {"artifactId":"1","enabled":true,"parameters":[{"name":"detection_threshold","value":"0.5"}]}
swagger:model UpdateDeviceCameraCustomAnalyticsBody
*/
type UpdateDeviceCameraCustomAnalyticsBody struct {

	// The ID of the custom analytics artifact
	ArtifactID string `json:"artifactId,omitempty"`

	// Enable custom analytics
	Enabled bool `json:"enabled,omitempty"`

	// Parameters for the custom analytics workload
	Parameters []*UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0 `json:"parameters"`
}

// Validate validates this update device camera custom analytics body
func (o *UpdateDeviceCameraCustomAnalyticsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCameraCustomAnalyticsBody) validateParameters(formats strfmt.Registry) error {
	if swag.IsZero(o.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(o.Parameters); i++ {
		if swag.IsZero(o.Parameters[i]) { // not required
			continue
		}

		if o.Parameters[i] != nil {
			if err := o.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCameraCustomAnalytics" + "." + "parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCameraCustomAnalytics" + "." + "parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update device camera custom analytics body based on the context it is used
func (o *UpdateDeviceCameraCustomAnalyticsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateParameters(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCameraCustomAnalyticsBody) contextValidateParameters(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Parameters); i++ {

		if o.Parameters[i] != nil {

			if swag.IsZero(o.Parameters[i]) { // not required
				return nil
			}

			if err := o.Parameters[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCameraCustomAnalytics" + "." + "parameters" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCameraCustomAnalytics" + "." + "parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCameraCustomAnalyticsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCameraCustomAnalyticsBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCameraCustomAnalyticsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0 update device camera custom analytics params body parameters items0
swagger:model UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0
*/
type UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0 struct {

	// Name of the parameter
	// Required: true
	Name *string `json:"name"`

	// Value of the parameter
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update device camera custom analytics params body parameters items0
func (o *UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
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

func (o *UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device camera custom analytics params body parameters items0 based on context it is used
func (o *UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCameraCustomAnalyticsParamsBodyParametersItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
