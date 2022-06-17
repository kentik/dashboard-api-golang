// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkApplianceConnectivityMonitoringDestinationsReader is a Reader for the GetNetworkApplianceConnectivityMonitoringDestinations structure.
type GetNetworkApplianceConnectivityMonitoringDestinationsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceConnectivityMonitoringDestinationsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceConnectivityMonitoringDestinationsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceConnectivityMonitoringDestinationsOK creates a GetNetworkApplianceConnectivityMonitoringDestinationsOK with default headers values
func NewGetNetworkApplianceConnectivityMonitoringDestinationsOK() *GetNetworkApplianceConnectivityMonitoringDestinationsOK {
	return &GetNetworkApplianceConnectivityMonitoringDestinationsOK{}
}

/* GetNetworkApplianceConnectivityMonitoringDestinationsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceConnectivityMonitoringDestinationsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network appliance connectivity monitoring destinations o k response has a 2xx status code
func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance connectivity monitoring destinations o k response has a 3xx status code
func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance connectivity monitoring destinations o k response has a 4xx status code
func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance connectivity monitoring destinations o k response has a 5xx status code
func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance connectivity monitoring destinations o k response a status code equal to that given
func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/connectivityMonitoringDestinations][%d] getNetworkApplianceConnectivityMonitoringDestinationsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/connectivityMonitoringDestinations][%d] getNetworkApplianceConnectivityMonitoringDestinationsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceConnectivityMonitoringDestinationsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
