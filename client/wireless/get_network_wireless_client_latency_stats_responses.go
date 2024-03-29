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

// GetNetworkWirelessClientLatencyStatsReader is a Reader for the GetNetworkWirelessClientLatencyStats structure.
type GetNetworkWirelessClientLatencyStatsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWirelessClientLatencyStatsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWirelessClientLatencyStatsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/wireless/clients/{clientId}/latencyStats] getNetworkWirelessClientLatencyStats", response, response.Code())
	}
}

// NewGetNetworkWirelessClientLatencyStatsOK creates a GetNetworkWirelessClientLatencyStatsOK with default headers values
func NewGetNetworkWirelessClientLatencyStatsOK() *GetNetworkWirelessClientLatencyStatsOK {
	return &GetNetworkWirelessClientLatencyStatsOK{}
}

/*
GetNetworkWirelessClientLatencyStatsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWirelessClientLatencyStatsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network wireless client latency stats o k response has a 2xx status code
func (o *GetNetworkWirelessClientLatencyStatsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network wireless client latency stats o k response has a 3xx status code
func (o *GetNetworkWirelessClientLatencyStatsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network wireless client latency stats o k response has a 4xx status code
func (o *GetNetworkWirelessClientLatencyStatsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network wireless client latency stats o k response has a 5xx status code
func (o *GetNetworkWirelessClientLatencyStatsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network wireless client latency stats o k response a status code equal to that given
func (o *GetNetworkWirelessClientLatencyStatsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network wireless client latency stats o k response
func (o *GetNetworkWirelessClientLatencyStatsOK) Code() int {
	return 200
}

func (o *GetNetworkWirelessClientLatencyStatsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/clients/{clientId}/latencyStats][%d] getNetworkWirelessClientLatencyStatsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessClientLatencyStatsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/wireless/clients/{clientId}/latencyStats][%d] getNetworkWirelessClientLatencyStatsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWirelessClientLatencyStatsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkWirelessClientLatencyStatsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
