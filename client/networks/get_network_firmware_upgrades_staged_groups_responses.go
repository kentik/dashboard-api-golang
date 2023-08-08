// Code generated by go-swagger; DO NOT EDIT.

package networks

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
)

// GetNetworkFirmwareUpgradesStagedGroupsReader is a Reader for the GetNetworkFirmwareUpgradesStagedGroups structure.
type GetNetworkFirmwareUpgradesStagedGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkFirmwareUpgradesStagedGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkFirmwareUpgradesStagedGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/firmwareUpgrades/staged/groups] getNetworkFirmwareUpgradesStagedGroups", response, response.Code())
	}
}

// NewGetNetworkFirmwareUpgradesStagedGroupsOK creates a GetNetworkFirmwareUpgradesStagedGroupsOK with default headers values
func NewGetNetworkFirmwareUpgradesStagedGroupsOK() *GetNetworkFirmwareUpgradesStagedGroupsOK {
	return &GetNetworkFirmwareUpgradesStagedGroupsOK{}
}

/*
GetNetworkFirmwareUpgradesStagedGroupsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkFirmwareUpgradesStagedGroupsOK struct {
	Payload []*GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0
}

// IsSuccess returns true when this get network firmware upgrades staged groups o k response has a 2xx status code
func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network firmware upgrades staged groups o k response has a 3xx status code
func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network firmware upgrades staged groups o k response has a 4xx status code
func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network firmware upgrades staged groups o k response has a 5xx status code
func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network firmware upgrades staged groups o k response a status code equal to that given
func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network firmware upgrades staged groups o k response
func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) Code() int {
	return 200
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/firmwareUpgrades/staged/groups][%d] getNetworkFirmwareUpgradesStagedGroupsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/firmwareUpgrades/staged/groups][%d] getNetworkFirmwareUpgradesStagedGroupsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) GetPayload() []*GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0 get network firmware upgrades staged groups o k body items0
swagger:model GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0
*/
type GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0 struct {

	// assigned devices
	AssignedDevices *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices `json:"assignedDevices,omitempty"`

	// Description of the Staged Upgrade Group
	Description string `json:"description,omitempty"`

	// Id of staged upgrade group
	GroupID string `json:"groupId,omitempty"`

	// Boolean indicating the default Group. Any device that does not have a group explicitly assigned will upgrade with this group
	IsDefault bool `json:"isDefault,omitempty"`

	// Name of the Staged Upgrade Group
	Name string `json:"name,omitempty"`
}

// Validate validates this get network firmware upgrades staged groups o k body items0
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAssignedDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0) validateAssignedDevices(formats strfmt.Registry) error {
	if swag.IsZero(o.AssignedDevices) { // not required
		return nil
	}

	if o.AssignedDevices != nil {
		if err := o.AssignedDevices.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignedDevices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignedDevices")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network firmware upgrades staged groups o k body items0 based on the context it is used
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAssignedDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0) contextValidateAssignedDevices(ctx context.Context, formats strfmt.Registry) error {

	if o.AssignedDevices != nil {

		if swag.IsZero(o.AssignedDevices) { // not required
			return nil
		}

		if err := o.AssignedDevices.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assignedDevices")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("assignedDevices")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices The devices and Switch Stacks assigned to the Group
swagger:model GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices
*/
type GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices struct {

	// Data Array of Devices containing the name and serial
	Devices []*GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0 `json:"devices"`

	// Data Array of Switch Stacks containing the name and id
	SwitchStacks []*GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0 `json:"switchStacks"`
}

// Validate validates this get network firmware upgrades staged groups o k body items0 assigned devices
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSwitchStacks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) validateDevices(formats strfmt.Registry) error {
	if swag.IsZero(o.Devices) { // not required
		return nil
	}

	for i := 0; i < len(o.Devices); i++ {
		if swag.IsZero(o.Devices[i]) { // not required
			continue
		}

		if o.Devices[i] != nil {
			if err := o.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) validateSwitchStacks(formats strfmt.Registry) error {
	if swag.IsZero(o.SwitchStacks) { // not required
		return nil
	}

	for i := 0; i < len(o.SwitchStacks); i++ {
		if swag.IsZero(o.SwitchStacks[i]) { // not required
			continue
		}

		if o.SwitchStacks[i] != nil {
			if err := o.SwitchStacks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network firmware upgrades staged groups o k body items0 assigned devices based on the context it is used
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSwitchStacks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Devices); i++ {

		if o.Devices[i] != nil {

			if swag.IsZero(o.Devices[i]) { // not required
				return nil
			}

			if err := o.Devices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignedDevices" + "." + "devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) contextValidateSwitchStacks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.SwitchStacks); i++ {

		if o.SwitchStacks[i] != nil {

			if swag.IsZero(o.SwitchStacks[i]) { // not required
				return nil
			}

			if err := o.SwitchStacks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("assignedDevices" + "." + "switchStacks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices) UnmarshalBinary(b []byte) error {
	var res GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevices
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0 get network firmware upgrades staged groups o k body items0 assigned devices devices items0
swagger:model GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0
*/
type GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0 struct {

	// Name of the device
	Name string `json:"name,omitempty"`

	// Serial of the device
	Serial string `json:"serial,omitempty"`
}

// Validate validates this get network firmware upgrades staged groups o k body items0 assigned devices devices items0
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network firmware upgrades staged groups o k body items0 assigned devices devices items0 based on context it is used
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesDevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0 get network firmware upgrades staged groups o k body items0 assigned devices switch stacks items0
swagger:model GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0
*/
type GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0 struct {

	// ID of the Switch Stack
	ID string `json:"id,omitempty"`

	// Name of the Switch Stack
	Name string `json:"name,omitempty"`
}

// Validate validates this get network firmware upgrades staged groups o k body items0 assigned devices switch stacks items0
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network firmware upgrades staged groups o k body items0 assigned devices switch stacks items0 based on context it is used
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkFirmwareUpgradesStagedGroupsOKBodyItems0AssignedDevicesSwitchStacksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
