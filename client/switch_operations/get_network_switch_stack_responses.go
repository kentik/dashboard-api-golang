// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetNetworkSwitchStackReader is a Reader for the GetNetworkSwitchStack structure.
type GetNetworkSwitchStackReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchStackReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchStackOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/switch/stacks/{switchStackId}] getNetworkSwitchStack", response, response.Code())
	}
}

// NewGetNetworkSwitchStackOK creates a GetNetworkSwitchStackOK with default headers values
func NewGetNetworkSwitchStackOK() *GetNetworkSwitchStackOK {
	return &GetNetworkSwitchStackOK{}
}

/*
GetNetworkSwitchStackOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchStackOK struct {
	Payload *GetNetworkSwitchStackOKBody
}

// IsSuccess returns true when this get network switch stack o k response has a 2xx status code
func (o *GetNetworkSwitchStackOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch stack o k response has a 3xx status code
func (o *GetNetworkSwitchStackOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch stack o k response has a 4xx status code
func (o *GetNetworkSwitchStackOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch stack o k response has a 5xx status code
func (o *GetNetworkSwitchStackOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch stack o k response a status code equal to that given
func (o *GetNetworkSwitchStackOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch stack o k response
func (o *GetNetworkSwitchStackOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchStackOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/stacks/{switchStackId}][%d] getNetworkSwitchStackOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchStackOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/stacks/{switchStackId}][%d] getNetworkSwitchStackOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchStackOK) GetPayload() *GetNetworkSwitchStackOKBody {
	return o.Payload
}

func (o *GetNetworkSwitchStackOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkSwitchStackOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSwitchStackOKBody get network switch stack o k body
swagger:model GetNetworkSwitchStackOKBody
*/
type GetNetworkSwitchStackOKBody struct {

	// Switch stacks id
	ID string `json:"id,omitempty"`

	// Switch stacks name
	Name string `json:"name,omitempty"`

	// Serials of the switches in the switch stack
	Serials []string `json:"serials"`
}

// Validate validates this get network switch stack o k body
func (o *GetNetworkSwitchStackOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network switch stack o k body based on context it is used
func (o *GetNetworkSwitchStackOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSwitchStackOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSwitchStackOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkSwitchStackOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
