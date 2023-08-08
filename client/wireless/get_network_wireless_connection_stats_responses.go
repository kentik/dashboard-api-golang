// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetNetworkWirelessConnectionStatsReader is a Reader for the GetNetworkWirelessConnectionStats structure.
type GetNetworkWirelessConnectionStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessConnectionStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessConnectionStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/wireless/connectionStats] getNetworkWirelessConnectionStats", response, response.Code())
	}
}

// NewGetNetworkWirelessConnectionStatsOK creates a GetNetworkWirelessConnectionStatsOK with default headers values
func NewGetNetworkWirelessConnectionStatsOK() *GetNetworkWirelessConnectionStatsOK {
	return &GetNetworkWirelessConnectionStatsOK{}
}

/*
GetNetworkWirelessConnectionStatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessConnectionStatsOK struct {
	Payload *GetNetworkWirelessConnectionStatsOKBody
}

// IsSuccess returns true when this get network wireless connection stats o k response has a 2xx status code
func (o *GetNetworkWirelessConnectionStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless connection stats o k response has a 3xx status code
func (o *GetNetworkWirelessConnectionStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless connection stats o k response has a 4xx status code
func (o *GetNetworkWirelessConnectionStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless connection stats o k response has a 5xx status code
func (o *GetNetworkWirelessConnectionStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless connection stats o k response a status code equal to that given
func (o *GetNetworkWirelessConnectionStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless connection stats o k response
func (o *GetNetworkWirelessConnectionStatsOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessConnectionStatsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/connectionStats][%d] getNetworkWirelessConnectionStatsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessConnectionStatsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/connectionStats][%d] getNetworkWirelessConnectionStatsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessConnectionStatsOK) GetPayload() *GetNetworkWirelessConnectionStatsOKBody {
	return o.Payload
}

func (o *GetNetworkWirelessConnectionStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkWirelessConnectionStatsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWirelessConnectionStatsOKBody get network wireless connection stats o k body
swagger:model GetNetworkWirelessConnectionStatsOKBody
*/
type GetNetworkWirelessConnectionStatsOKBody struct {

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

// Validate validates this get network wireless connection stats o k body
func (o *GetNetworkWirelessConnectionStatsOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network wireless connection stats o k body based on context it is used
func (o *GetNetworkWirelessConnectionStatsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWirelessConnectionStatsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWirelessConnectionStatsOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkWirelessConnectionStatsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
