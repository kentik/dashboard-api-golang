// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// CreateOrganizationInventoryOnboardingCloudMonitoringImportReader is a Reader for the CreateOrganizationInventoryOnboardingCloudMonitoringImport structure.
type CreateOrganizationInventoryOnboardingCloudMonitoringImportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationInventoryOnboardingCloudMonitoringImportCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports] createOrganizationInventoryOnboardingCloudMonitoringImport", response, response.Code())
	}
}

// NewCreateOrganizationInventoryOnboardingCloudMonitoringImportCreated creates a CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated with default headers values
func NewCreateOrganizationInventoryOnboardingCloudMonitoringImportCreated() *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated {
	return &CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated{}
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated struct {
	Payload []*CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0
}

// IsSuccess returns true when this create organization inventory onboarding cloud monitoring import created response has a 2xx status code
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization inventory onboarding cloud monitoring import created response has a 3xx status code
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization inventory onboarding cloud monitoring import created response has a 4xx status code
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization inventory onboarding cloud monitoring import created response has a 5xx status code
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization inventory onboarding cloud monitoring import created response a status code equal to that given
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create organization inventory onboarding cloud monitoring import created response
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) Code() int {
	return 201
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports][%d] createOrganizationInventoryOnboardingCloudMonitoringImportCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/inventory/onboarding/cloudMonitoring/imports][%d] createOrganizationInventoryOnboardingCloudMonitoringImportCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) GetPayload() []*CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0 {
	return o.Payload
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImportBody create organization inventory onboarding cloud monitoring import body
// Example: {"devices":[{"deviceId":"161b2602-a713-4aac-b1eb-d9b55205353d","networkId":"1338481","udi":"PID:C9200L-24P-4G SN:JAE25220R2K"}]}
swagger:model CreateOrganizationInventoryOnboardingCloudMonitoringImportBody
*/
type CreateOrganizationInventoryOnboardingCloudMonitoringImportBody struct {

	// A set of device imports to commit
	// Required: true
	Devices []*CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0 `json:"devices"`
}

// Validate validates this create organization inventory onboarding cloud monitoring import body
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) validateDevices(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationInventoryOnboardingCloudMonitoringImport"+"."+"devices", "body", o.Devices); err != nil {
		return err
	}

	for i := 0; i < len(o.Devices); i++ {
		if swag.IsZero(o.Devices[i]) { // not required
			continue
		}

		if o.Devices[i] != nil {
			if err := o.Devices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationInventoryOnboardingCloudMonitoringImport" + "." + "devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationInventoryOnboardingCloudMonitoringImport" + "." + "devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create organization inventory onboarding cloud monitoring import body based on the context it is used
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) contextValidateDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Devices); i++ {

		if o.Devices[i] != nil {

			if swag.IsZero(o.Devices[i]) { // not required
				return nil
			}

			if err := o.Devices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationInventoryOnboardingCloudMonitoringImport" + "." + "devices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationInventoryOnboardingCloudMonitoringImport" + "." + "devices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationInventoryOnboardingCloudMonitoringImportBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0 create organization inventory onboarding cloud monitoring import created body items0
swagger:model CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0
*/
type CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0 struct {

	// Unique id associated with the import of the device
	ImportID string `json:"importId,omitempty"`

	// Response method
	Message string `json:"message,omitempty"`

	// Cloud monitor import status
	Status string `json:"status,omitempty"`
}

// Validate validates this create organization inventory onboarding cloud monitoring import created body items0
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create organization inventory onboarding cloud monitoring import created body items0 based on context it is used
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationInventoryOnboardingCloudMonitoringImportCreatedBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0 create organization inventory onboarding cloud monitoring import params body devices items0
swagger:model CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0
*/
type CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0 struct {

	// Import ID from the Import operation
	// Required: true
	DeviceID *string `json:"deviceId"`

	// Network Id
	// Required: true
	NetworkID *string `json:"networkId"`

	// Device UDI certificate
	// Required: true
	Udi *string `json:"udi"`
}

// Validate validates this create organization inventory onboarding cloud monitoring import params body devices items0
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDeviceID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkID(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUdi(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0) validateDeviceID(formats strfmt.Registry) error {

	if err := validate.Required("deviceId", "body", o.DeviceID); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0) validateNetworkID(formats strfmt.Registry) error {

	if err := validate.Required("networkId", "body", o.NetworkID); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0) validateUdi(formats strfmt.Registry) error {

	if err := validate.Required("udi", "body", o.Udi); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization inventory onboarding cloud monitoring import params body devices items0 based on context it is used
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationInventoryOnboardingCloudMonitoringImportParamsBodyDevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
