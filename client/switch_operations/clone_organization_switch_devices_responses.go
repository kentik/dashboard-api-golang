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

// CloneOrganizationSwitchDevicesReader is a Reader for the CloneOrganizationSwitchDevices structure.
type CloneOrganizationSwitchDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneOrganizationSwitchDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCloneOrganizationSwitchDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /organizations/{organizationId}/switch/devices/clone] cloneOrganizationSwitchDevices", response, response.Code())
	}
}

// NewCloneOrganizationSwitchDevicesOK creates a CloneOrganizationSwitchDevicesOK with default headers values
func NewCloneOrganizationSwitchDevicesOK() *CloneOrganizationSwitchDevicesOK {
	return &CloneOrganizationSwitchDevicesOK{}
}

/*
CloneOrganizationSwitchDevicesOK describes a response with status code 200, with default header values.

Successful operation
*/
type CloneOrganizationSwitchDevicesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this clone organization switch devices o k response has a 2xx status code
func (o *CloneOrganizationSwitchDevicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this clone organization switch devices o k response has a 3xx status code
func (o *CloneOrganizationSwitchDevicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone organization switch devices o k response has a 4xx status code
func (o *CloneOrganizationSwitchDevicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone organization switch devices o k response has a 5xx status code
func (o *CloneOrganizationSwitchDevicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this clone organization switch devices o k response a status code equal to that given
func (o *CloneOrganizationSwitchDevicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the clone organization switch devices o k response
func (o *CloneOrganizationSwitchDevicesOK) Code() int {
	return 200
}

func (o *CloneOrganizationSwitchDevicesOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/switch/devices/clone][%d] cloneOrganizationSwitchDevicesOK  %+v", 200, o.Payload)
}

func (o *CloneOrganizationSwitchDevicesOK) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/switch/devices/clone][%d] cloneOrganizationSwitchDevicesOK  %+v", 200, o.Payload)
}

func (o *CloneOrganizationSwitchDevicesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *CloneOrganizationSwitchDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CloneOrganizationSwitchDevicesBody clone organization switch devices body
// Example: {"sourceSerial":"Q234-ABCD-5678","targetSerials":["Q234-ABCD-0001","Q234-ABCD-0002","Q234-ABCD-0003"]}
swagger:model CloneOrganizationSwitchDevicesBody
*/
type CloneOrganizationSwitchDevicesBody struct {

	// Serial number of the source switch (must be on a network not bound to a template)
	// Required: true
	SourceSerial *string `json:"sourceSerial"`

	// Array of serial numbers of one or more target switches (must be on a network not bound to a template)
	// Required: true
	TargetSerials []string `json:"targetSerials"`
}

// Validate validates this clone organization switch devices body
func (o *CloneOrganizationSwitchDevicesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSourceSerial(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTargetSerials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CloneOrganizationSwitchDevicesBody) validateSourceSerial(formats strfmt.Registry) error {

	if err := validate.Required("cloneOrganizationSwitchDevices"+"."+"sourceSerial", "body", o.SourceSerial); err != nil {
		return err
	}

	return nil
}

func (o *CloneOrganizationSwitchDevicesBody) validateTargetSerials(formats strfmt.Registry) error {

	if err := validate.Required("cloneOrganizationSwitchDevices"+"."+"targetSerials", "body", o.TargetSerials); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this clone organization switch devices body based on context it is used
func (o *CloneOrganizationSwitchDevicesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationSwitchDevicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationSwitchDevicesBody) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationSwitchDevicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
