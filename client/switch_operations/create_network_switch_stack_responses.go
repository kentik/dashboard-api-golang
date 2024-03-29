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

// CreateNetworkSwitchStackReader is a Reader for the CreateNetworkSwitchStack structure.
type CreateNetworkSwitchStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchStackCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/switch/stacks] createNetworkSwitchStack", response, response.Code())
	}
}

// NewCreateNetworkSwitchStackCreated creates a CreateNetworkSwitchStackCreated with default headers values
func NewCreateNetworkSwitchStackCreated() *CreateNetworkSwitchStackCreated {
	return &CreateNetworkSwitchStackCreated{}
}

/*
CreateNetworkSwitchStackCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchStackCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network switch stack created response has a 2xx status code
func (o *CreateNetworkSwitchStackCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network switch stack created response has a 3xx status code
func (o *CreateNetworkSwitchStackCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network switch stack created response has a 4xx status code
func (o *CreateNetworkSwitchStackCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network switch stack created response has a 5xx status code
func (o *CreateNetworkSwitchStackCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network switch stack created response a status code equal to that given
func (o *CreateNetworkSwitchStackCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network switch stack created response
func (o *CreateNetworkSwitchStackCreated) Code() int {
	return 201
}

func (o *CreateNetworkSwitchStackCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/stacks][%d] createNetworkSwitchStackCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchStackCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/stacks][%d] createNetworkSwitchStackCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchStackCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchStackCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkSwitchStackBody create network switch stack body
// Example: {"name":"A cool stack","serials":["QBZY-XWVU-TSRQ","QBAB-CDEF-GHIJ"]}
swagger:model CreateNetworkSwitchStackBody
*/
type CreateNetworkSwitchStackBody struct {

	// The name of the new stack
	// Required: true
	Name *string `json:"name"`

	// An array of switch serials to be added into the new stack
	// Required: true
	Serials []string `json:"serials"`
}

// Validate validates this create network switch stack body
func (o *CreateNetworkSwitchStackBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSerials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchStackBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchStack"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchStackBody) validateSerials(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchStack"+"."+"serials", "body", o.Serials); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch stack body based on context it is used
func (o *CreateNetworkSwitchStackBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchStackBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchStackBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchStackBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
