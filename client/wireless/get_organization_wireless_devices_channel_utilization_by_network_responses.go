// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// GetOrganizationWirelessDevicesChannelUtilizationByNetworkReader is a Reader for the GetOrganizationWirelessDevicesChannelUtilizationByNetwork structure.
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationWirelessDevicesChannelUtilizationByNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork] getOrganizationWirelessDevicesChannelUtilizationByNetwork", response, response.Code())
	}
}

// NewGetOrganizationWirelessDevicesChannelUtilizationByNetworkOK creates a GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK with default headers values
func NewGetOrganizationWirelessDevicesChannelUtilizationByNetworkOK() *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK {
	return &GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK{}
}

/*
GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0
}

// IsSuccess returns true when this get organization wireless devices channel utilization by network o k response has a 2xx status code
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization wireless devices channel utilization by network o k response has a 3xx status code
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization wireless devices channel utilization by network o k response has a 4xx status code
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization wireless devices channel utilization by network o k response has a 5xx status code
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization wireless devices channel utilization by network o k response a status code equal to that given
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization wireless devices channel utilization by network o k response
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) Code() int {
	return 200
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork][%d] getOrganizationWirelessDevicesChannelUtilizationByNetworkOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/wireless/devices/channelUtilization/byNetwork][%d] getOrganizationWirelessDevicesChannelUtilizationByNetworkOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) GetPayload() []*GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0 get organization wireless devices channel utilization by network o k body items0
swagger:model GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0
*/
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0 struct {

	// Channel utilization broken down by band.
	ByBand []*GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0 `json:"byBand"`

	// network
	Network *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network `json:"network,omitempty"`
}

// Validate validates this get organization wireless devices channel utilization by network o k body items0
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateByBand(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) validateByBand(formats strfmt.Registry) error {
	if swag.IsZero(o.ByBand) { // not required
		return nil
	}

	for i := 0; i < len(o.ByBand); i++ {
		if swag.IsZero(o.ByBand[i]) { // not required
			continue
		}

		if o.ByBand[i] != nil {
			if err := o.ByBand[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("byBand" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("byBand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) validateNetwork(formats strfmt.Registry) error {
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

// ContextValidate validate this get organization wireless devices channel utilization by network o k body items0 based on the context it is used
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateByBand(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) contextValidateByBand(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ByBand); i++ {

		if o.ByBand[i] != nil {

			if swag.IsZero(o.ByBand[i]) { // not required
				return nil
			}

			if err := o.ByBand[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("byBand" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("byBand" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) contextValidateNetwork(ctx context.Context, formats strfmt.Registry) error {

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
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0 get organization wireless devices channel utilization by network o k body items0 by band items0
swagger:model GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0
*/
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0 struct {

	// The band for the given metrics.
	Band string `json:"band,omitempty"`

	// non wifi
	NonWifi *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi `json:"nonWifi,omitempty"`

	// total
	Total *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total `json:"total,omitempty"`

	// wifi
	Wifi *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi `json:"wifi,omitempty"`
}

// Validate validates this get organization wireless devices channel utilization by network o k body items0 by band items0
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNonWifi(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTotal(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWifi(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) validateNonWifi(formats strfmt.Registry) error {
	if swag.IsZero(o.NonWifi) { // not required
		return nil
	}

	if o.NonWifi != nil {
		if err := o.NonWifi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonWifi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonWifi")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) validateTotal(formats strfmt.Registry) error {
	if swag.IsZero(o.Total) { // not required
		return nil
	}

	if o.Total != nil {
		if err := o.Total.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) validateWifi(formats strfmt.Registry) error {
	if swag.IsZero(o.Wifi) { // not required
		return nil
	}

	if o.Wifi != nil {
		if err := o.Wifi.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wifi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wifi")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization wireless devices channel utilization by network o k body items0 by band items0 based on the context it is used
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNonWifi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTotal(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWifi(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) contextValidateNonWifi(ctx context.Context, formats strfmt.Registry) error {

	if o.NonWifi != nil {

		if swag.IsZero(o.NonWifi) { // not required
			return nil
		}

		if err := o.NonWifi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("nonWifi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("nonWifi")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) contextValidateTotal(ctx context.Context, formats strfmt.Registry) error {

	if o.Total != nil {

		if swag.IsZero(o.Total) { // not required
			return nil
		}

		if err := o.Total.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("total")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("total")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) contextValidateWifi(ctx context.Context, formats strfmt.Registry) error {

	if o.Wifi != nil {

		if swag.IsZero(o.Wifi) { // not required
			return nil
		}

		if err := o.Wifi.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("wifi")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("wifi")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi An object containing non-wifi utilization.
swagger:model GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi
*/
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi struct {

	// Percentage of non-wifi channel utiliation for the given band.
	Percentage float32 `json:"percentage,omitempty"`
}

// Validate validates this get organization wireless devices channel utilization by network o k body items0 by band items0 non wifi
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization wireless devices channel utilization by network o k body items0 by band items0 non wifi based on context it is used
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi) UnmarshalBinary(b []byte) error {
	var res GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0NonWifi
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total An object containing total channel utilization.
swagger:model GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total
*/
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total struct {

	// Percentage of total channel utiliation for the given band.
	Percentage float32 `json:"percentage,omitempty"`
}

// Validate validates this get organization wireless devices channel utilization by network o k body items0 by band items0 total
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization wireless devices channel utilization by network o k body items0 by band items0 total based on context it is used
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total) UnmarshalBinary(b []byte) error {
	var res GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Total
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi An object containing wifi utilization.
swagger:model GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi
*/
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi struct {

	// Percentage of wifi channel utiliation for the given band.
	Percentage float32 `json:"percentage,omitempty"`
}

// Validate validates this get organization wireless devices channel utilization by network o k body items0 by band items0 wifi
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization wireless devices channel utilization by network o k body items0 by band items0 wifi based on context it is used
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi) UnmarshalBinary(b []byte) error {
	var res GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0ByBandItems0Wifi
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network Network for the given utilization metrics.
swagger:model GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network
*/
type GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network struct {

	// Network ID of the given utilization metrics.
	ID string `json:"id,omitempty"`
}

// Validate validates this get organization wireless devices channel utilization by network o k body items0 network
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization wireless devices channel utilization by network o k body items0 network based on context it is used
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network) UnmarshalBinary(b []byte) error {
	var res GetOrganizationWirelessDevicesChannelUtilizationByNetworkOKBodyItems0Network
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
