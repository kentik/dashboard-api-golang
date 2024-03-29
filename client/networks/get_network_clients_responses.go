// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// GetNetworkClientsReader is a Reader for the GetNetworkClients structure.
type GetNetworkClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/clients] getNetworkClients", response, response.Code())
	}
}

// NewGetNetworkClientsOK creates a GetNetworkClientsOK with default headers values
func NewGetNetworkClientsOK() *GetNetworkClientsOK {
	return &GetNetworkClientsOK{}
}

/*
GetNetworkClientsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkClientsOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload *GetNetworkClientsOKBody
}

// IsSuccess returns true when this get network clients o k response has a 2xx status code
func (o *GetNetworkClientsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network clients o k response has a 3xx status code
func (o *GetNetworkClientsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network clients o k response has a 4xx status code
func (o *GetNetworkClientsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network clients o k response has a 5xx status code
func (o *GetNetworkClientsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network clients o k response a status code equal to that given
func (o *GetNetworkClientsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network clients o k response
func (o *GetNetworkClientsOK) Code() int {
	return 200
}

func (o *GetNetworkClientsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/clients][%d] getNetworkClientsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkClientsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/clients][%d] getNetworkClientsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkClientsOK) GetPayload() *GetNetworkClientsOKBody {
	return o.Payload
}

func (o *GetNetworkClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	o.Payload = new(GetNetworkClientsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkClientsOKBody get network clients o k body
swagger:model GetNetworkClientsOKBody
*/
type GetNetworkClientsOKBody struct {

	// The adaptive policy group of the client
	AdaptivePolicyGroup string `json:"adaptivePolicyGroup,omitempty"`

	// Short description of the client
	Description string `json:"description,omitempty"`

	// Prediction of the client's device type
	DeviceTypePrediction string `json:"deviceTypePrediction,omitempty"`

	// Timestamp client was first seen in the network
	FirstSeen int64 `json:"firstSeen,omitempty"`

	// 802.1x group policy of the client
	GroupPolicy8021x string `json:"groupPolicy8021x,omitempty"`

	// The ID of the client
	ID string `json:"id,omitempty"`

	// The IP address of the client
	IP string `json:"ip,omitempty"`

	// The IPv6 address of the client
	Ip6 string `json:"ip6,omitempty"`

	// Local IPv6 address of the client
	Ip6Local string `json:"ip6Local,omitempty"`

	// Timestamp client was last seen in the network
	LastSeen int64 `json:"lastSeen,omitempty"`

	// The MAC address of the client
	Mac string `json:"mac,omitempty"`

	// Manufacturer of the client
	Manufacturer string `json:"manufacturer,omitempty"`

	// Named VLAN of the client
	NamedVlan string `json:"namedVlan,omitempty"`

	// Notes on the client
	Notes string `json:"notes,omitempty"`

	// The operating system of the client
	Os string `json:"os,omitempty"`

	// iPSK name of the client
	PskGroup string `json:"pskGroup,omitempty"`

	// Client's most recent connection type
	// Enum: [Wired Wireless]
	RecentDeviceConnection string `json:"recentDeviceConnection,omitempty"`

	// The MAC address of the node that the device was last connected to
	RecentDeviceMac string `json:"recentDeviceMac,omitempty"`

	// The name of the node the device was last connected to
	RecentDeviceName string `json:"recentDeviceName,omitempty"`

	// The serial of the node the device was last connected to
	RecentDeviceSerial string `json:"recentDeviceSerial,omitempty"`

	// Status of SM for the client
	SmInstalled bool `json:"smInstalled,omitempty"`

	// The name of the SSID that the client is connected to
	Ssid string `json:"ssid,omitempty"`

	// The connection status of the client
	// Enum: [Offline Online]
	Status string `json:"status,omitempty"`

	// The switch port that the client is connected to
	Switchport string `json:"switchport,omitempty"`

	// usage
	Usage *GetNetworkClientsOKBodyUsage `json:"usage,omitempty"`

	// The username of the user of the client
	User string `json:"user,omitempty"`

	// The name of the VLAN that the client is connected to
	Vlan string `json:"vlan,omitempty"`

	// Wireless capabilities of the client
	WirelessCapabilities string `json:"wirelessCapabilities,omitempty"`
}

// Validate validates this get network clients o k body
func (o *GetNetworkClientsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRecentDeviceConnection(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkClientsOKBodyTypeRecentDeviceConnectionPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Wired","Wireless"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkClientsOKBodyTypeRecentDeviceConnectionPropEnum = append(getNetworkClientsOKBodyTypeRecentDeviceConnectionPropEnum, v)
	}
}

const (

	// GetNetworkClientsOKBodyRecentDeviceConnectionWired captures enum value "Wired"
	GetNetworkClientsOKBodyRecentDeviceConnectionWired string = "Wired"

	// GetNetworkClientsOKBodyRecentDeviceConnectionWireless captures enum value "Wireless"
	GetNetworkClientsOKBodyRecentDeviceConnectionWireless string = "Wireless"
)

// prop value enum
func (o *GetNetworkClientsOKBody) validateRecentDeviceConnectionEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkClientsOKBodyTypeRecentDeviceConnectionPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkClientsOKBody) validateRecentDeviceConnection(formats strfmt.Registry) error {
	if swag.IsZero(o.RecentDeviceConnection) { // not required
		return nil
	}

	// value enum
	if err := o.validateRecentDeviceConnectionEnum("getNetworkClientsOK"+"."+"recentDeviceConnection", "body", o.RecentDeviceConnection); err != nil {
		return err
	}

	return nil
}

var getNetworkClientsOKBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Offline","Online"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkClientsOKBodyTypeStatusPropEnum = append(getNetworkClientsOKBodyTypeStatusPropEnum, v)
	}
}

const (

	// GetNetworkClientsOKBodyStatusOffline captures enum value "Offline"
	GetNetworkClientsOKBodyStatusOffline string = "Offline"

	// GetNetworkClientsOKBodyStatusOnline captures enum value "Online"
	GetNetworkClientsOKBodyStatusOnline string = "Online"
)

// prop value enum
func (o *GetNetworkClientsOKBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkClientsOKBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkClientsOKBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("getNetworkClientsOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkClientsOKBody) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(o.Usage) { // not required
		return nil
	}

	if o.Usage != nil {
		if err := o.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkClientsOK" + "." + "usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkClientsOK" + "." + "usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network clients o k body based on the context it is used
func (o *GetNetworkClientsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkClientsOKBody) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if o.Usage != nil {

		if swag.IsZero(o.Usage) { // not required
			return nil
		}

		if err := o.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkClientsOK" + "." + "usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkClientsOK" + "." + "usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkClientsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkClientsOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkClientsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkClientsOKBodyUsage Usage, sent and received
swagger:model GetNetworkClientsOKBodyUsage
*/
type GetNetworkClientsOKBodyUsage struct {

	// Usage received by the client
	Recv float32 `json:"recv,omitempty"`

	// Usage sent by the client
	Sent float32 `json:"sent,omitempty"`
}

// Validate validates this get network clients o k body usage
func (o *GetNetworkClientsOKBodyUsage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network clients o k body usage based on context it is used
func (o *GetNetworkClientsOKBodyUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkClientsOKBodyUsage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkClientsOKBodyUsage) UnmarshalBinary(b []byte) error {
	var res GetNetworkClientsOKBodyUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
