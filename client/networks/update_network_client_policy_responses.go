// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// UpdateNetworkClientPolicyReader is a Reader for the UpdateNetworkClientPolicy structure.
type UpdateNetworkClientPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkClientPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkClientPolicyOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/clients/{clientId}/policy] updateNetworkClientPolicy", response, response.Code())
	}
}

// NewUpdateNetworkClientPolicyOK creates a UpdateNetworkClientPolicyOK with default headers values
func NewUpdateNetworkClientPolicyOK() *UpdateNetworkClientPolicyOK {
	return &UpdateNetworkClientPolicyOK{}
}

/*
UpdateNetworkClientPolicyOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkClientPolicyOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network client policy o k response has a 2xx status code
func (o *UpdateNetworkClientPolicyOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network client policy o k response has a 3xx status code
func (o *UpdateNetworkClientPolicyOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network client policy o k response has a 4xx status code
func (o *UpdateNetworkClientPolicyOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network client policy o k response has a 5xx status code
func (o *UpdateNetworkClientPolicyOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network client policy o k response a status code equal to that given
func (o *UpdateNetworkClientPolicyOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network client policy o k response
func (o *UpdateNetworkClientPolicyOK) Code() int {
	return 200
}

func (o *UpdateNetworkClientPolicyOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/clients/{clientId}/policy][%d] updateNetworkClientPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkClientPolicyOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/clients/{clientId}/policy][%d] updateNetworkClientPolicyOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkClientPolicyOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkClientPolicyOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkClientPolicyBody update network client policy body
// Example: {"devicePolicy":"Group policy","groupPolicyId":"101","mac":"00:11:22:33:44:55"}
swagger:model UpdateNetworkClientPolicyBody
*/
type UpdateNetworkClientPolicyBody struct {

	// The policy to assign. Can be 'Whitelisted', 'Blocked', 'Normal' or 'Group policy'. Required.
	// Required: true
	DevicePolicy *string `json:"devicePolicy"`

	// [optional] If 'devicePolicy' is set to 'Group policy' this param is used to specify the group policy ID.
	GroupPolicyID string `json:"groupPolicyId,omitempty"`
}

// Validate validates this update network client policy body
func (o *UpdateNetworkClientPolicyBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevicePolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkClientPolicyBody) validateDevicePolicy(formats strfmt.Registry) error {

	if err := validate.Required("updateNetworkClientPolicy"+"."+"devicePolicy", "body", o.DevicePolicy); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network client policy body based on context it is used
func (o *UpdateNetworkClientPolicyBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkClientPolicyBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkClientPolicyBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkClientPolicyBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
