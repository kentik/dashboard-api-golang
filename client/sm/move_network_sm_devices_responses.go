// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// MoveNetworkSmDevicesReader is a Reader for the MoveNetworkSmDevices structure.
type MoveNetworkSmDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MoveNetworkSmDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMoveNetworkSmDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/sm/devices/move] moveNetworkSmDevices", response, response.Code())
	}
}

// NewMoveNetworkSmDevicesOK creates a MoveNetworkSmDevicesOK with default headers values
func NewMoveNetworkSmDevicesOK() *MoveNetworkSmDevicesOK {
	return &MoveNetworkSmDevicesOK{}
}

/*
MoveNetworkSmDevicesOK describes a response with status code 200, with default header values.

Successful operation
*/
type MoveNetworkSmDevicesOK struct {
	Payload *MoveNetworkSmDevicesOKBody
}

// IsSuccess returns true when this move network sm devices o k response has a 2xx status code
func (o *MoveNetworkSmDevicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this move network sm devices o k response has a 3xx status code
func (o *MoveNetworkSmDevicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this move network sm devices o k response has a 4xx status code
func (o *MoveNetworkSmDevicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this move network sm devices o k response has a 5xx status code
func (o *MoveNetworkSmDevicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this move network sm devices o k response a status code equal to that given
func (o *MoveNetworkSmDevicesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the move network sm devices o k response
func (o *MoveNetworkSmDevicesOK) Code() int {
	return 200
}

func (o *MoveNetworkSmDevicesOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/devices/move][%d] moveNetworkSmDevicesOK  %+v", 200, o.Payload)
}

func (o *MoveNetworkSmDevicesOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/devices/move][%d] moveNetworkSmDevicesOK  %+v", 200, o.Payload)
}

func (o *MoveNetworkSmDevicesOK) GetPayload() *MoveNetworkSmDevicesOKBody {
	return o.Payload
}

func (o *MoveNetworkSmDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(MoveNetworkSmDevicesOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
MoveNetworkSmDevicesBody move network sm devices body
// Example: {"ids":["1284392014819","2983092129865"],"newNetwork":"1284392014819","scope":["withAny","tag1","tag2"],"serials":["Q234-ABCD-0001","Q234-ABCD-0002","Q234-ABCD-0003"],"wifiMacs":["00:11:22:33:44:55"]}
swagger:model MoveNetworkSmDevicesBody
*/
type MoveNetworkSmDevicesBody struct {

	// The ids of the devices to be moved.
	Ids []string `json:"ids"`

	// The new network to which the devices will be moved.
	// Required: true
	NewNetwork *string `json:"newNetwork"`

	// The scope (one of all, none, withAny, withAll, withoutAny, or withoutAll) and a set of tags of the devices to be moved.
	Scope []string `json:"scope"`

	// The serials of the devices to be moved.
	Serials []string `json:"serials"`

	// The wifiMacs of the devices to be moved.
	WifiMacs []string `json:"wifiMacs"`
}

// Validate validates this move network sm devices body
func (o *MoveNetworkSmDevicesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNewNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *MoveNetworkSmDevicesBody) validateNewNetwork(formats strfmt.Registry) error {

	if err := validate.Required("moveNetworkSmDevices"+"."+"newNetwork", "body", o.NewNetwork); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this move network sm devices body based on context it is used
func (o *MoveNetworkSmDevicesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveNetworkSmDevicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveNetworkSmDevicesBody) UnmarshalBinary(b []byte) error {
	var res MoveNetworkSmDevicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
MoveNetworkSmDevicesOKBody move network sm devices o k body
swagger:model MoveNetworkSmDevicesOKBody
*/
type MoveNetworkSmDevicesOKBody struct {

	// The Meraki Ids of the set of devices.
	Ids []string `json:"ids"`

	// The network to which the devices was moved.
	NewNetwork string `json:"newNetwork,omitempty"`
}

// Validate validates this move network sm devices o k body
func (o *MoveNetworkSmDevicesOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this move network sm devices o k body based on context it is used
func (o *MoveNetworkSmDevicesOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *MoveNetworkSmDevicesOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *MoveNetworkSmDevicesOKBody) UnmarshalBinary(b []byte) error {
	var res MoveNetworkSmDevicesOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
