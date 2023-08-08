// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// GetOrganizationFirmwareUpgradesReader is a Reader for the GetOrganizationFirmwareUpgrades structure.
type GetOrganizationFirmwareUpgradesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationFirmwareUpgradesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationFirmwareUpgradesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/firmware/upgrades] getOrganizationFirmwareUpgrades", response, response.Code())
	}
}

// NewGetOrganizationFirmwareUpgradesOK creates a GetOrganizationFirmwareUpgradesOK with default headers values
func NewGetOrganizationFirmwareUpgradesOK() *GetOrganizationFirmwareUpgradesOK {
	return &GetOrganizationFirmwareUpgradesOK{}
}

/*
GetOrganizationFirmwareUpgradesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationFirmwareUpgradesOK struct {
	Payload []*GetOrganizationFirmwareUpgradesOKBodyItems0
}

// IsSuccess returns true when this get organization firmware upgrades o k response has a 2xx status code
func (o *GetOrganizationFirmwareUpgradesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization firmware upgrades o k response has a 3xx status code
func (o *GetOrganizationFirmwareUpgradesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization firmware upgrades o k response has a 4xx status code
func (o *GetOrganizationFirmwareUpgradesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization firmware upgrades o k response has a 5xx status code
func (o *GetOrganizationFirmwareUpgradesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization firmware upgrades o k response a status code equal to that given
func (o *GetOrganizationFirmwareUpgradesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization firmware upgrades o k response
func (o *GetOrganizationFirmwareUpgradesOK) Code() int {
	return 200
}

func (o *GetOrganizationFirmwareUpgradesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/firmware/upgrades][%d] getOrganizationFirmwareUpgradesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationFirmwareUpgradesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/firmware/upgrades][%d] getOrganizationFirmwareUpgradesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationFirmwareUpgradesOK) GetPayload() []*GetOrganizationFirmwareUpgradesOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationFirmwareUpgradesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationFirmwareUpgradesOKBodyItems0 get organization firmware upgrades o k body items0
swagger:model GetOrganizationFirmwareUpgradesOKBodyItems0
*/
type GetOrganizationFirmwareUpgradesOKBodyItems0 struct {

	// Timestamp when upgrade completed. Null if status pending.
	CompletedAt string `json:"completedAt,omitempty"`

	// from version
	FromVersion *GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion `json:"fromVersion,omitempty"`

	// network
	Network *GetOrganizationFirmwareUpgradesOKBodyItems0Network `json:"network,omitempty"`

	// product upgraded [wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor]
	ProductTypes string `json:"productTypes,omitempty"`

	// Status of upgrade event: [Cancelled, Completed]
	Status string `json:"status,omitempty"`

	// Scheduled start time
	// Format: date-time
	Time strfmt.DateTime `json:"time,omitempty"`

	// to version
	ToVersion *GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion `json:"toVersion,omitempty"`

	// The upgrade batch
	UpgradeBatchID string `json:"upgradeBatchId,omitempty"`

	// The upgrade
	UpgradeID string `json:"upgradeId,omitempty"`
}

// Validate validates this get organization firmware upgrades o k body items0
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFromVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTime(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateToVersion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) validateFromVersion(formats strfmt.Registry) error {
	if swag.IsZero(o.FromVersion) { // not required
		return nil
	}

	if o.FromVersion != nil {
		if err := o.FromVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fromVersion")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) validateNetwork(formats strfmt.Registry) error {
	if swag.IsZero(o.Network) { // not required
		return nil
	}

	if o.Network != nil {
		if err := o.Network.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) validateTime(formats strfmt.Registry) error {
	if swag.IsZero(o.Time) { // not required
		return nil
	}

	if err := validate.FormatOf("time", "body", "date-time", o.Time.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) validateToVersion(formats strfmt.Registry) error {
	if swag.IsZero(o.ToVersion) { // not required
		return nil
	}

	if o.ToVersion != nil {
		if err := o.ToVersion.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("toVersion")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization firmware upgrades o k body items0 based on the context it is used
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFromVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateToVersion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) contextValidateFromVersion(ctx context.Context, formats strfmt.Registry) error {

	if o.FromVersion != nil {

		if swag.IsZero(o.FromVersion) { // not required
			return nil
		}

		if err := o.FromVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("fromVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("fromVersion")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

	if o.Network != nil {

		if swag.IsZero(o.Network) { // not required
			return nil
		}

		if err := o.Network.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("network")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("network")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) contextValidateToVersion(ctx context.Context, formats strfmt.Registry) error {

	if o.ToVersion != nil {

		if swag.IsZero(o.ToVersion) { // not required
			return nil
		}

		if err := o.ToVersion.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("toVersion")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("toVersion")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationFirmwareUpgradesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion ID of the upgrade's starting version
swagger:model GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion
*/
type GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion struct {

	// Firmware version ID
	ID string `json:"id,omitempty"`

	// Release date of the firmware version
	// Format: date-time
	ReleaseDate strfmt.DateTime `json:"releaseDate,omitempty"`

	// Release type of the firmware version
	ReleaseType string `json:"releaseType,omitempty"`

	// Firmware version short name
	ShortName string `json:"shortName,omitempty"`
}

// Validate validates this get organization firmware upgrades o k body items0 from version
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion) validateReleaseDate(formats strfmt.Registry) error {
	if swag.IsZero(o.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("fromVersion"+"."+"releaseDate", "body", "date-time", o.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organization firmware upgrades o k body items0 from version based on context it is used
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion) UnmarshalBinary(b []byte) error {
	var res GetOrganizationFirmwareUpgradesOKBodyItems0FromVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationFirmwareUpgradesOKBodyItems0Network Network of the upgrade
swagger:model GetOrganizationFirmwareUpgradesOKBodyItems0Network
*/
type GetOrganizationFirmwareUpgradesOKBodyItems0Network struct {

	// ID of network
	ID string `json:"id,omitempty"`

	// The network
	Name string `json:"name,omitempty"`
}

// Validate validates this get organization firmware upgrades o k body items0 network
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0Network) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization firmware upgrades o k body items0 network based on context it is used
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0Network) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0Network) UnmarshalBinary(b []byte) error {
	var res GetOrganizationFirmwareUpgradesOKBodyItems0Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion ID of the upgrade's target version
swagger:model GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion
*/
type GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion struct {

	// Firmware version ID
	ID string `json:"id,omitempty"`

	// Release date of the firmware version
	// Format: date-time
	ReleaseDate strfmt.DateTime `json:"releaseDate,omitempty"`

	// Release type of the firmware version
	ReleaseType string `json:"releaseType,omitempty"`

	// Firmware version short name
	ShortName string `json:"shortName,omitempty"`
}

// Validate validates this get organization firmware upgrades o k body items0 to version
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateReleaseDate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion) validateReleaseDate(formats strfmt.Registry) error {
	if swag.IsZero(o.ReleaseDate) { // not required
		return nil
	}

	if err := validate.FormatOf("toVersion"+"."+"releaseDate", "body", "date-time", o.ReleaseDate.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get organization firmware upgrades o k body items0 to version based on context it is used
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion) UnmarshalBinary(b []byte) error {
	var res GetOrganizationFirmwareUpgradesOKBodyItems0ToVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}