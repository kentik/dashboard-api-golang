// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetNetworkSensorRelationshipsReader is a Reader for the GetNetworkSensorRelationships structure.
type GetNetworkSensorRelationshipsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSensorRelationshipsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSensorRelationshipsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/sensor/relationships] getNetworkSensorRelationships", response, response.Code())
	}
}

// NewGetNetworkSensorRelationshipsOK creates a GetNetworkSensorRelationshipsOK with default headers values
func NewGetNetworkSensorRelationshipsOK() *GetNetworkSensorRelationshipsOK {
	return &GetNetworkSensorRelationshipsOK{}
}

/*
GetNetworkSensorRelationshipsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSensorRelationshipsOK struct {
	Payload []*GetNetworkSensorRelationshipsOKBodyItems0
}

// IsSuccess returns true when this get network sensor relationships o k response has a 2xx status code
func (o *GetNetworkSensorRelationshipsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sensor relationships o k response has a 3xx status code
func (o *GetNetworkSensorRelationshipsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sensor relationships o k response has a 4xx status code
func (o *GetNetworkSensorRelationshipsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sensor relationships o k response has a 5xx status code
func (o *GetNetworkSensorRelationshipsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sensor relationships o k response a status code equal to that given
func (o *GetNetworkSensorRelationshipsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sensor relationships o k response
func (o *GetNetworkSensorRelationshipsOK) Code() int {
	return 200
}

func (o *GetNetworkSensorRelationshipsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sensor/relationships][%d] getNetworkSensorRelationshipsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSensorRelationshipsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sensor/relationships][%d] getNetworkSensorRelationshipsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSensorRelationshipsOK) GetPayload() []*GetNetworkSensorRelationshipsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSensorRelationshipsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSensorRelationshipsOKBodyItems0 get network sensor relationships o k body items0
swagger:model GetNetworkSensorRelationshipsOKBodyItems0
*/
type GetNetworkSensorRelationshipsOKBodyItems0 struct {

	// device
	Device *GetNetworkSensorRelationshipsOKBodyItems0Device `json:"device,omitempty"`

	// relationships
	Relationships *GetNetworkSensorRelationshipsOKBodyItems0Relationships `json:"relationships,omitempty"`
}

// Validate validates this get network sensor relationships o k body items0
func (o *GetNetworkSensorRelationshipsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateDevice(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRelationships(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0) validateDevice(formats strfmt.Registry) error {
	if swag.IsZero(o.Device) { // not required
		return nil
	}

	if o.Device != nil {
		if err := o.Device.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0) validateRelationships(formats strfmt.Registry) error {
	if swag.IsZero(o.Relationships) { // not required
		return nil
	}

	if o.Relationships != nil {
		if err := o.Relationships.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network sensor relationships o k body items0 based on the context it is used
func (o *GetNetworkSensorRelationshipsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateDevice(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRelationships(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0) contextValidateDevice(ctx context.Context, formats strfmt.Registry) error {

	if o.Device != nil {

		if swag.IsZero(o.Device) { // not required
			return nil
		}

		if err := o.Device.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("device")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("device")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0) contextValidateRelationships(ctx context.Context, formats strfmt.Registry) error {

	if o.Relationships != nil {

		if swag.IsZero(o.Relationships) { // not required
			return nil
		}

		if err := o.Relationships.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relationships")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorRelationshipsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSensorRelationshipsOKBodyItems0Device A sensor or gateway device in the network
swagger:model GetNetworkSensorRelationshipsOKBodyItems0Device
*/
type GetNetworkSensorRelationshipsOKBodyItems0Device struct {

	// The name of the device
	Name string `json:"name,omitempty"`

	// The product type of the device
	// Enum: [camera sensor]
	ProductType string `json:"productType,omitempty"`

	// The serial of the device
	Serial string `json:"serial,omitempty"`
}

// Validate validates this get network sensor relationships o k body items0 device
func (o *GetNetworkSensorRelationshipsOKBodyItems0Device) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkSensorRelationshipsOKBodyItems0DeviceTypeProductTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["camera","sensor"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkSensorRelationshipsOKBodyItems0DeviceTypeProductTypePropEnum = append(getNetworkSensorRelationshipsOKBodyItems0DeviceTypeProductTypePropEnum, v)
	}
}

const (

	// GetNetworkSensorRelationshipsOKBodyItems0DeviceProductTypeCamera captures enum value "camera"
	GetNetworkSensorRelationshipsOKBodyItems0DeviceProductTypeCamera string = "camera"

	// GetNetworkSensorRelationshipsOKBodyItems0DeviceProductTypeSensor captures enum value "sensor"
	GetNetworkSensorRelationshipsOKBodyItems0DeviceProductTypeSensor string = "sensor"
)

// prop value enum
func (o *GetNetworkSensorRelationshipsOKBodyItems0Device) validateProductTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkSensorRelationshipsOKBodyItems0DeviceTypeProductTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0Device) validateProductType(formats strfmt.Registry) error {
	if swag.IsZero(o.ProductType) { // not required
		return nil
	}

	// value enum
	if err := o.validateProductTypeEnum("device"+"."+"productType", "body", o.ProductType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network sensor relationships o k body items0 device based on context it is used
func (o *GetNetworkSensorRelationshipsOKBodyItems0Device) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0Device) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0Device) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorRelationshipsOKBodyItems0Device
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSensorRelationshipsOKBodyItems0Relationships An object describing the relationships defined between the device and other devices
swagger:model GetNetworkSensorRelationshipsOKBodyItems0Relationships
*/
type GetNetworkSensorRelationshipsOKBodyItems0Relationships struct {

	// livestream
	Livestream *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream `json:"livestream,omitempty"`
}

// Validate validates this get network sensor relationships o k body items0 relationships
func (o *GetNetworkSensorRelationshipsOKBodyItems0Relationships) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLivestream(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0Relationships) validateLivestream(formats strfmt.Registry) error {
	if swag.IsZero(o.Livestream) { // not required
		return nil
	}

	if o.Livestream != nil {
		if err := o.Livestream.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "livestream")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relationships" + "." + "livestream")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network sensor relationships o k body items0 relationships based on the context it is used
func (o *GetNetworkSensorRelationshipsOKBodyItems0Relationships) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateLivestream(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0Relationships) contextValidateLivestream(ctx context.Context, formats strfmt.Registry) error {

	if o.Livestream != nil {

		if swag.IsZero(o.Livestream) { // not required
			return nil
		}

		if err := o.Livestream.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("relationships" + "." + "livestream")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("relationships" + "." + "livestream")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0Relationships) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0Relationships) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorRelationshipsOKBodyItems0Relationships
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream A role defined between an MT sensor and an MV camera that adds the camera's livestream to the sensor's details page. Snapshots from the camera will also appear in alert notifications that the sensor triggers.
swagger:model GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream
*/
type GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream struct {

	// An array of the related devices for the role
	RelatedDevices []*GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0 `json:"relatedDevices"`
}

// Validate validates this get network sensor relationships o k body items0 relationships livestream
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRelatedDevices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream) validateRelatedDevices(formats strfmt.Registry) error {
	if swag.IsZero(o.RelatedDevices) { // not required
		return nil
	}

	for i := 0; i < len(o.RelatedDevices); i++ {
		if swag.IsZero(o.RelatedDevices[i]) { // not required
			continue
		}

		if o.RelatedDevices[i] != nil {
			if err := o.RelatedDevices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "livestream" + "." + "relatedDevices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relationships" + "." + "livestream" + "." + "relatedDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network sensor relationships o k body items0 relationships livestream based on the context it is used
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRelatedDevices(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream) contextValidateRelatedDevices(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.RelatedDevices); i++ {

		if o.RelatedDevices[i] != nil {

			if swag.IsZero(o.RelatedDevices[i]) { // not required
				return nil
			}

			if err := o.RelatedDevices[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("relationships" + "." + "livestream" + "." + "relatedDevices" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("relationships" + "." + "livestream" + "." + "relatedDevices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestream
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0 get network sensor relationships o k body items0 relationships livestream related devices items0
swagger:model GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0
*/
type GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0 struct {

	// The product type of the related device
	// Enum: [camera sensor]
	ProductType string `json:"productType,omitempty"`

	// The serial of the related device
	Serial string `json:"serial,omitempty"`
}

// Validate validates this get network sensor relationships o k body items0 relationships livestream related devices items0
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateProductType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0TypeProductTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["camera","sensor"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0TypeProductTypePropEnum = append(getNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0TypeProductTypePropEnum, v)
	}
}

const (

	// GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0ProductTypeCamera captures enum value "camera"
	GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0ProductTypeCamera string = "camera"

	// GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0ProductTypeSensor captures enum value "sensor"
	GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0ProductTypeSensor string = "sensor"
)

// prop value enum
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0) validateProductTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0TypeProductTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0) validateProductType(formats strfmt.Registry) error {
	if swag.IsZero(o.ProductType) { // not required
		return nil
	}

	// value enum
	if err := o.validateProductTypeEnum("productType", "body", o.ProductType); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network sensor relationships o k body items0 relationships livestream related devices items0 based on context it is used
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorRelationshipsOKBodyItems0RelationshipsLivestreamRelatedDevicesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}