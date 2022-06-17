// Code generated by go-swagger; DO NOT EDIT.

package wireless

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
)

// GetNetworkWirelessDevicesConnectionStatsReader is a Reader for the GetNetworkWirelessDevicesConnectionStats structure.
type GetNetworkWirelessDevicesConnectionStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessDevicesConnectionStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessDevicesConnectionStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessDevicesConnectionStatsOK creates a GetNetworkWirelessDevicesConnectionStatsOK with default headers values
func NewGetNetworkWirelessDevicesConnectionStatsOK() *GetNetworkWirelessDevicesConnectionStatsOK {
	return &GetNetworkWirelessDevicesConnectionStatsOK{}
}

/* GetNetworkWirelessDevicesConnectionStatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessDevicesConnectionStatsOK struct {
	Payload []*GetNetworkWirelessDevicesConnectionStatsOKBodyItems0
}

// IsSuccess returns true when this get network wireless devices connection stats o k response has a 2xx status code
func (o *GetNetworkWirelessDevicesConnectionStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless devices connection stats o k response has a 3xx status code
func (o *GetNetworkWirelessDevicesConnectionStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless devices connection stats o k response has a 4xx status code
func (o *GetNetworkWirelessDevicesConnectionStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless devices connection stats o k response has a 5xx status code
func (o *GetNetworkWirelessDevicesConnectionStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless devices connection stats o k response a status code equal to that given
func (o *GetNetworkWirelessDevicesConnectionStatsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkWirelessDevicesConnectionStatsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/devices/connectionStats][%d] getNetworkWirelessDevicesConnectionStatsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessDevicesConnectionStatsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/devices/connectionStats][%d] getNetworkWirelessDevicesConnectionStatsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessDevicesConnectionStatsOK) GetPayload() []*GetNetworkWirelessDevicesConnectionStatsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkWirelessDevicesConnectionStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNetworkWirelessDevicesConnectionStatsOKBodyItems0 get network wireless devices connection stats o k body items0
swagger:model GetNetworkWirelessDevicesConnectionStatsOKBodyItems0
*/
type GetNetworkWirelessDevicesConnectionStatsOKBodyItems0 struct {

	// connection stats
	ConnectionStats *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats `json:"connectionStats,omitempty"`

	// The serial number for the device
	Serial string `json:"serial,omitempty"`
}

// Validate validates this get network wireless devices connection stats o k body items0
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateConnectionStats(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0) validateConnectionStats(formats strfmt.Registry) error {
	if swag.IsZero(o.ConnectionStats) { // not required
		return nil
	}

	if o.ConnectionStats != nil {
		if err := o.ConnectionStats.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionStats")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network wireless devices connection stats o k body items0 based on the context it is used
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateConnectionStats(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0) contextValidateConnectionStats(ctx context.Context, formats strfmt.Registry) error {

	if o.ConnectionStats != nil {
		if err := o.ConnectionStats.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("connectionStats")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("connectionStats")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessDevicesConnectionStatsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats The connection stats of the device
swagger:model GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats
*/
type GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats struct {

	// The number of failed association attempts
	Assoc int64 `json:"assoc,omitempty"`

	// The number of failed authentication attempts
	Auth int64 `json:"auth,omitempty"`

	// The number of failed DHCP attempts
	Dhcp int64 `json:"dhcp,omitempty"`

	// The number of failed DNS attempts
	DNS int64 `json:"dns,omitempty"`

	// The number of successful connection attempts
	Success int64 `json:"success,omitempty"`
}

// Validate validates this get network wireless devices connection stats o k body items0 connection stats
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network wireless devices connection stats o k body items0 connection stats based on context it is used
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessDevicesConnectionStatsOKBodyItems0ConnectionStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}