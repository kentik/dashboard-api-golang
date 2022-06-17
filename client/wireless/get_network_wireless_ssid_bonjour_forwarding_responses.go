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

// GetNetworkWirelessSsidBonjourForwardingReader is a Reader for the GetNetworkWirelessSsidBonjourForwarding structure.
type GetNetworkWirelessSsidBonjourForwardingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessSsidBonjourForwardingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessSsidBonjourForwardingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkWirelessSsidBonjourForwardingOK creates a GetNetworkWirelessSsidBonjourForwardingOK with default headers values
func NewGetNetworkWirelessSsidBonjourForwardingOK() *GetNetworkWirelessSsidBonjourForwardingOK {
	return &GetNetworkWirelessSsidBonjourForwardingOK{}
}

/* GetNetworkWirelessSsidBonjourForwardingOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessSsidBonjourForwardingOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network wireless ssid bonjour forwarding o k response has a 2xx status code
func (o *GetNetworkWirelessSsidBonjourForwardingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless ssid bonjour forwarding o k response has a 3xx status code
func (o *GetNetworkWirelessSsidBonjourForwardingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless ssid bonjour forwarding o k response has a 4xx status code
func (o *GetNetworkWirelessSsidBonjourForwardingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless ssid bonjour forwarding o k response has a 5xx status code
func (o *GetNetworkWirelessSsidBonjourForwardingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless ssid bonjour forwarding o k response a status code equal to that given
func (o *GetNetworkWirelessSsidBonjourForwardingOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkWirelessSsidBonjourForwardingOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding][%d] getNetworkWirelessSsidBonjourForwardingOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidBonjourForwardingOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/ssids/{number}/bonjourForwarding][%d] getNetworkWirelessSsidBonjourForwardingOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessSsidBonjourForwardingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessSsidBonjourForwardingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}