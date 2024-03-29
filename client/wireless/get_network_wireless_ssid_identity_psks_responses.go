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
	"github.com/go-openapi/validate"
)

// GetNetworkWirelessSsidIdentityPsksReader is a Reader for the GetNetworkWirelessSsidIdentityPsks structure.
type GetNetworkWirelessSsidIdentityPsksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSsidIdentityPsksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSsidIdentityPsksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/wireless/ssids/{number}/identityPsks] getNetworkWirelessSsidIdentityPsks", response, response.Code())
	}
}

// NewGetNetworkWirelessSsidIdentityPsksOK creates a GetNetworkWirelessSsidIdentityPsksOK with default headers values
func NewGetNetworkWirelessSsidIdentityPsksOK() *GetNetworkWirelessSsidIdentityPsksOK {
	return &GetNetworkWirelessSsidIdentityPsksOK{}
}

/*
GetNetworkWirelessSsidIdentityPsksOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSsidIdentityPsksOK struct {
	Payload []*GetNetworkWirelessSsidIdentityPsksOKBodyItems0
}

// IsSuccess returns true when this get network wireless ssid identity psks o k response has a 2xx status code
func (o *GetNetworkWirelessSsidIdentityPsksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless ssid identity psks o k response has a 3xx status code
func (o *GetNetworkWirelessSsidIdentityPsksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless ssid identity psks o k response has a 4xx status code
func (o *GetNetworkWirelessSsidIdentityPsksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless ssid identity psks o k response has a 5xx status code
func (o *GetNetworkWirelessSsidIdentityPsksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless ssid identity psks o k response a status code equal to that given
func (o *GetNetworkWirelessSsidIdentityPsksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless ssid identity psks o k response
func (o *GetNetworkWirelessSsidIdentityPsksOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessSsidIdentityPsksOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/identityPsks][%d] getNetworkWirelessSsidIdentityPsksOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidIdentityPsksOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/identityPsks][%d] getNetworkWirelessSsidIdentityPsksOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidIdentityPsksOK) GetPayload() []*GetNetworkWirelessSsidIdentityPsksOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkWirelessSsidIdentityPsksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessSsidIdentityPsksOKBodyItems0 get network wireless ssid identity psks o k body items0
swagger:model GetNetworkWirelessSsidIdentityPsksOKBodyItems0
*/
type GetNetworkWirelessSsidIdentityPsksOKBodyItems0 struct {

	// The email associated with the System's Manager User
	Email string `json:"email,omitempty"`

	// Timestamp for when the Identity PSK expires, or 'null' to never expire
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expiresAt,omitempty"`

	// The group policy to be applied to clients
	GroupPolicyID string `json:"groupPolicyId,omitempty"`

	// The unique identifier of the Identity PSK
	ID string `json:"id,omitempty"`

	// The name of the Identity PSK
	Name string `json:"name,omitempty"`

	// The passphrase for client authentication
	Passphrase string `json:"passphrase,omitempty"`

	// The WiFi Personal Network unique identifier
	WifiPersonalNetworkID string `json:"wifiPersonalNetworkId,omitempty"`
}

// Validate validates this get network wireless ssid identity psks o k body items0
func (o *GetNetworkWirelessSsidIdentityPsksOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessSsidIdentityPsksOKBodyItems0) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(o.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", o.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network wireless ssid identity psks o k body items0 based on context it is used
func (o *GetNetworkWirelessSsidIdentityPsksOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessSsidIdentityPsksOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessSsidIdentityPsksOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessSsidIdentityPsksOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
