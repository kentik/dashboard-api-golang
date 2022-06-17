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

// GetNetworkWirelessFailedConnectionsReader is a Reader for the GetNetworkWirelessFailedConnections structure.
type GetNetworkWirelessFailedConnectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessFailedConnectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessFailedConnectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessFailedConnectionsOK creates a GetNetworkWirelessFailedConnectionsOK with default headers values
func NewGetNetworkWirelessFailedConnectionsOK() *GetNetworkWirelessFailedConnectionsOK {
	return &GetNetworkWirelessFailedConnectionsOK{}
}

/* GetNetworkWirelessFailedConnectionsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessFailedConnectionsOK struct {
	Payload []*GetNetworkWirelessFailedConnectionsOKBodyItems0
}

// IsSuccess returns true when this get network wireless failed connections o k response has a 2xx status code
func (o *GetNetworkWirelessFailedConnectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless failed connections o k response has a 3xx status code
func (o *GetNetworkWirelessFailedConnectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless failed connections o k response has a 4xx status code
func (o *GetNetworkWirelessFailedConnectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless failed connections o k response has a 5xx status code
func (o *GetNetworkWirelessFailedConnectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless failed connections o k response a status code equal to that given
func (o *GetNetworkWirelessFailedConnectionsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkWirelessFailedConnectionsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/failedConnections][%d] getNetworkWirelessFailedConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessFailedConnectionsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/failedConnections][%d] getNetworkWirelessFailedConnectionsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessFailedConnectionsOK) GetPayload() []*GetNetworkWirelessFailedConnectionsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkWirelessFailedConnectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNetworkWirelessFailedConnectionsOKBodyItems0 get network wireless failed connections o k body items0
swagger:model GetNetworkWirelessFailedConnectionsOKBodyItems0
*/
type GetNetworkWirelessFailedConnectionsOKBodyItems0 struct {

	// Client Mac
	ClientMac string `json:"clientMac,omitempty"`

	// The failed onboarding step. One of: assoc, auth, dhcp, dns.
	FailureStep string `json:"failureStep,omitempty"`

	// Serial Number
	Serial string `json:"serial,omitempty"`

	// SSID Number
	SsidNumber int64 `json:"ssidNumber,omitempty"`

	// The timestamp when the client mac failed
	// Format: date-time
	Ts strfmt.DateTime `json:"ts,omitempty"`

	// The failure type in the onboarding step
	Type string `json:"type,omitempty"`

	// LAN
	Vlan int64 `json:"vlan,omitempty"`
}

// Validate validates this get network wireless failed connections o k body items0
func (o *GetNetworkWirelessFailedConnectionsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWirelessFailedConnectionsOKBodyItems0) validateTs(formats strfmt.Registry) error {
	if swag.IsZero(o.Ts) { // not required
		return nil
	}

	if err := validate.FormatOf("ts", "body", "date-time", o.Ts.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network wireless failed connections o k body items0 based on context it is used
func (o *GetNetworkWirelessFailedConnectionsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessFailedConnectionsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessFailedConnectionsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessFailedConnectionsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
