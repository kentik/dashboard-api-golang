// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkWirelessSsidVpnReader is a Reader for the GetNetworkWirelessSsidVpn structure.
type GetNetworkWirelessSsidVpnReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSsidVpnReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSsidVpnOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/wireless/ssids/{number}/vpn] getNetworkWirelessSsidVpn", response, response.Code())
	}
}

// NewGetNetworkWirelessSsidVpnOK creates a GetNetworkWirelessSsidVpnOK with default headers values
func NewGetNetworkWirelessSsidVpnOK() *GetNetworkWirelessSsidVpnOK {
	return &GetNetworkWirelessSsidVpnOK{}
}

/*
GetNetworkWirelessSsidVpnOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSsidVpnOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network wireless ssid vpn o k response has a 2xx status code
func (o *GetNetworkWirelessSsidVpnOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless ssid vpn o k response has a 3xx status code
func (o *GetNetworkWirelessSsidVpnOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless ssid vpn o k response has a 4xx status code
func (o *GetNetworkWirelessSsidVpnOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless ssid vpn o k response has a 5xx status code
func (o *GetNetworkWirelessSsidVpnOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless ssid vpn o k response a status code equal to that given
func (o *GetNetworkWirelessSsidVpnOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless ssid vpn o k response
func (o *GetNetworkWirelessSsidVpnOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessSsidVpnOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/vpn][%d] getNetworkWirelessSsidVpnOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidVpnOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/vpn][%d] getNetworkWirelessSsidVpnOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidVpnOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessSsidVpnOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
