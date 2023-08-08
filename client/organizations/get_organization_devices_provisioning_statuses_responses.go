// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetOrganizationDevicesProvisioningStatusesReader is a Reader for the GetOrganizationDevicesProvisioningStatuses structure.
type GetOrganizationDevicesProvisioningStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationDevicesProvisioningStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationDevicesProvisioningStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/devices/provisioning/statuses] getOrganizationDevicesProvisioningStatuses", response, response.Code())
	}
}

// NewGetOrganizationDevicesProvisioningStatusesOK creates a GetOrganizationDevicesProvisioningStatusesOK with default headers values
func NewGetOrganizationDevicesProvisioningStatusesOK() *GetOrganizationDevicesProvisioningStatusesOK {
	return &GetOrganizationDevicesProvisioningStatusesOK{}
}

/*
GetOrganizationDevicesProvisioningStatusesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationDevicesProvisioningStatusesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetOrganizationDevicesProvisioningStatusesOKBodyItems0
}

// IsSuccess returns true when this get organization devices provisioning statuses o k response has a 2xx status code
func (o *GetOrganizationDevicesProvisioningStatusesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization devices provisioning statuses o k response has a 3xx status code
func (o *GetOrganizationDevicesProvisioningStatusesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization devices provisioning statuses o k response has a 4xx status code
func (o *GetOrganizationDevicesProvisioningStatusesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization devices provisioning statuses o k response has a 5xx status code
func (o *GetOrganizationDevicesProvisioningStatusesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization devices provisioning statuses o k response a status code equal to that given
func (o *GetOrganizationDevicesProvisioningStatusesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization devices provisioning statuses o k response
func (o *GetOrganizationDevicesProvisioningStatusesOK) Code() int {
	return 200
}

func (o *GetOrganizationDevicesProvisioningStatusesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/devices/provisioning/statuses][%d] getOrganizationDevicesProvisioningStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationDevicesProvisioningStatusesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/devices/provisioning/statuses][%d] getOrganizationDevicesProvisioningStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationDevicesProvisioningStatusesOK) GetPayload() []*GetOrganizationDevicesProvisioningStatusesOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationDevicesProvisioningStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationDevicesProvisioningStatusesOKBodyItems0 get organization devices provisioning statuses o k body items0
swagger:model GetOrganizationDevicesProvisioningStatusesOKBodyItems0
*/
type GetOrganizationDevicesProvisioningStatusesOKBodyItems0 struct {

	// The device MAC address.
	Mac string `json:"mac,omitempty"`

	// The device name.
	Name string `json:"name,omitempty"`

	// network
	Network *GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network `json:"network,omitempty"`

	// Device product type.
	// Enum: [appliance camera cellularGateway cloudGateway sensor switch systemsManager wireless]
	ProductType string `json:"productType,omitempty"`

	// The device serial number.
	Serial string `json:"serial,omitempty"`

	// The device provisioning status. Possible statuses: unprovisioned, incomplete, complete.
	// Enum: [complete incomplete unprovisioned]
	Status string `json:"status,omitempty"`

	// List of custom tags for the device.
	Tags []string `json:"tags"`
}

// Validate validates this get organization devices provisioning statuses o k body items0
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) validateNetwork(formats strfmt.Registry) error {
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

var getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeProductTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["appliance","camera","cellularGateway","cloudGateway","sensor","switch","systemsManager","wireless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeProductTypePropEnum = append(getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeProductTypePropEnum, v)
	}
}

const (

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeAppliance captures enum value "appliance"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeAppliance string = "appliance"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeCamera captures enum value "camera"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeCamera string = "camera"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeCellularGateway captures enum value "cellularGateway"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeCellularGateway string = "cellularGateway"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeCloudGateway captures enum value "cloudGateway"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeCloudGateway string = "cloudGateway"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeSensor captures enum value "sensor"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeSensor string = "sensor"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeSwitch captures enum value "switch"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeSwitch string = "switch"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeSystemsManager captures enum value "systemsManager"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeSystemsManager string = "systemsManager"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeWireless captures enum value "wireless"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0ProductTypeWireless string = "wireless"
)

// prop value enum
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) validateProductTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeProductTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) validateProductType(formats strfmt.Registry) error {
	if swag.IsZero(o.ProductType) { // not required
		return nil
	}

	// value enum
	if err := o.validateProductTypeEnum("productType", "body", o.ProductType); err != nil {
		return err
	}

	return nil
}

var getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["complete","incomplete","unprovisioned"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeStatusPropEnum = append(getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeStatusPropEnum, v)
	}
}

const (

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0StatusComplete captures enum value "complete"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0StatusComplete string = "complete"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0StatusIncomplete captures enum value "incomplete"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0StatusIncomplete string = "incomplete"

	// GetOrganizationDevicesProvisioningStatusesOKBodyItems0StatusUnprovisioned captures enum value "unprovisioned"
	GetOrganizationDevicesProvisioningStatusesOKBodyItems0StatusUnprovisioned string = "unprovisioned"
)

// prop value enum
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationDevicesProvisioningStatusesOKBodyItems0TypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organization devices provisioning statuses o k body items0 based on the context it is used
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationDevicesProvisioningStatusesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network Network info.
swagger:model GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network
*/
type GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network struct {

	// ID for the network containing the device.
	ID string `json:"id,omitempty"`
}

// Validate validates this get organization devices provisioning statuses o k body items0 network
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization devices provisioning statuses o k body items0 network based on context it is used
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network) UnmarshalBinary(b []byte) error {
	var res GetOrganizationDevicesProvisioningStatusesOKBodyItems0Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}