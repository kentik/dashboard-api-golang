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

// UpdateDeviceCellularGatewayLanReader is a Reader for the UpdateDeviceCellularGatewayLan structure.
type UpdateDeviceCellularGatewayLanReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceCellularGatewayLanReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceCellularGatewayLanOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}/cellularGateway/lan] updateDeviceCellularGatewayLan", response, response.Code())
	}
}

// NewUpdateDeviceCellularGatewayLanOK creates a UpdateDeviceCellularGatewayLanOK with default headers values
func NewUpdateDeviceCellularGatewayLanOK() *UpdateDeviceCellularGatewayLanOK {
	return &UpdateDeviceCellularGatewayLanOK{}
}

/*
UpdateDeviceCellularGatewayLanOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceCellularGatewayLanOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device cellular gateway lan o k response has a 2xx status code
func (o *UpdateDeviceCellularGatewayLanOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device cellular gateway lan o k response has a 3xx status code
func (o *UpdateDeviceCellularGatewayLanOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device cellular gateway lan o k response has a 4xx status code
func (o *UpdateDeviceCellularGatewayLanOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device cellular gateway lan o k response has a 5xx status code
func (o *UpdateDeviceCellularGatewayLanOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device cellular gateway lan o k response a status code equal to that given
func (o *UpdateDeviceCellularGatewayLanOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device cellular gateway lan o k response
func (o *UpdateDeviceCellularGatewayLanOK) Code() int {
	return 200
}

func (o *UpdateDeviceCellularGatewayLanOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/cellularGateway/lan][%d] updateDeviceCellularGatewayLanOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCellularGatewayLanOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/cellularGateway/lan][%d] updateDeviceCellularGatewayLanOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceCellularGatewayLanOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceCellularGatewayLanOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceCellularGatewayLanBody update device cellular gateway lan body
// Example: {"deviceLanIp":"192.168.0.33","deviceName":"name of the MG","deviceSubnet":"192.168.0.32/27","fixedIpAssignments":[{"ip":"192.168.0.10","mac":"0b:00:00:00:00:ac","name":"server 1"},{"ip":"192.168.0.20","mac":"0b:00:00:00:00:ab","name":"server 2"}],"reservedIpRanges":[{"comment":"A reserved IP range","end":"192.168.1.1","start":"192.168.1.0"}]}
swagger:model UpdateDeviceCellularGatewayLanBody
*/
type UpdateDeviceCellularGatewayLanBody struct {

	// list of all fixed IP assignments for a single MG
	FixedIPAssignments []*UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0 `json:"fixedIpAssignments"`

	// list of all reserved IP ranges for a single MG
	ReservedIPRanges []*UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0 `json:"reservedIpRanges"`
}

// Validate validates this update device cellular gateway lan body
func (o *UpdateDeviceCellularGatewayLanBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFixedIPAssignments(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateReservedIPRanges(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularGatewayLanBody) validateFixedIPAssignments(formats strfmt.Registry) error {
	if swag.IsZero(o.FixedIPAssignments) { // not required
		return nil
	}

	for i := 0; i < len(o.FixedIPAssignments); i++ {
		if swag.IsZero(o.FixedIPAssignments[i]) { // not required
			continue
		}

		if o.FixedIPAssignments[i] != nil {
			if err := o.FixedIPAssignments[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularGatewayLan" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularGatewayLan" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateDeviceCellularGatewayLanBody) validateReservedIPRanges(formats strfmt.Registry) error {
	if swag.IsZero(o.ReservedIPRanges) { // not required
		return nil
	}

	for i := 0; i < len(o.ReservedIPRanges); i++ {
		if swag.IsZero(o.ReservedIPRanges[i]) { // not required
			continue
		}

		if o.ReservedIPRanges[i] != nil {
			if err := o.ReservedIPRanges[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularGatewayLan" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularGatewayLan" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update device cellular gateway lan body based on the context it is used
func (o *UpdateDeviceCellularGatewayLanBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFixedIPAssignments(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateReservedIPRanges(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularGatewayLanBody) contextValidateFixedIPAssignments(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.FixedIPAssignments); i++ {

		if o.FixedIPAssignments[i] != nil {

			if swag.IsZero(o.FixedIPAssignments[i]) { // not required
				return nil
			}

			if err := o.FixedIPAssignments[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularGatewayLan" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularGatewayLan" + "." + "fixedIpAssignments" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateDeviceCellularGatewayLanBody) contextValidateReservedIPRanges(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ReservedIPRanges); i++ {

		if o.ReservedIPRanges[i] != nil {

			if swag.IsZero(o.ReservedIPRanges[i]) { // not required
				return nil
			}

			if err := o.ReservedIPRanges[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateDeviceCellularGatewayLan" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateDeviceCellularGatewayLan" + "." + "reservedIpRanges" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayLanBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayLanBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularGatewayLanBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0 update device cellular gateway lan params body fixed IP assignments items0
swagger:model UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0
*/
type UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0 struct {

	// The IP address you want to assign to a specific server or device
	// Required: true
	IP *string `json:"ip"`

	// The MAC address of the server or device that hosts the internal resource that you wish to receive the specified IP address
	// Required: true
	Mac *string `json:"mac"`

	// A descriptive name of the assignment
	Name string `json:"name,omitempty"`
}

// Validate validates this update device cellular gateway lan params body fixed IP assignments items0
func (o *UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMac(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", o.IP); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0) validateMac(formats strfmt.Registry) error {

	if err := validate.Required("mac", "body", o.Mac); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device cellular gateway lan params body fixed IP assignments items0 based on context it is used
func (o *UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularGatewayLanParamsBodyFixedIPAssignmentsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0 update device cellular gateway lan params body reserved IP ranges items0
swagger:model UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0
*/
type UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0 struct {

	// Comment explaining the reserved IP range
	// Required: true
	Comment *string `json:"comment"`

	// Ending IP included in the reserved range of IPs
	// Required: true
	End *string `json:"end"`

	// Starting IP included in the reserved range of IPs
	// Required: true
	Start *string `json:"start"`
}

// Validate validates this update device cellular gateway lan params body reserved IP ranges items0
func (o *UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateComment(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEnd(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStart(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0) validateComment(formats strfmt.Registry) error {

	if err := validate.Required("comment", "body", o.Comment); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0) validateEnd(formats strfmt.Registry) error {

	if err := validate.Required("end", "body", o.End); err != nil {
		return err
	}

	return nil
}

func (o *UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0) validateStart(formats strfmt.Registry) error {

	if err := validate.Required("start", "body", o.Start); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update device cellular gateway lan params body reserved IP ranges items0 based on context it is used
func (o *UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceCellularGatewayLanParamsBodyReservedIPRangesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
