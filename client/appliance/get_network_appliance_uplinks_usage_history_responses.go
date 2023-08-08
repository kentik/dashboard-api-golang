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

// GetNetworkApplianceUplinksUsageHistoryReader is a Reader for the GetNetworkApplianceUplinksUsageHistory structure.
type GetNetworkApplianceUplinksUsageHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceUplinksUsageHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceUplinksUsageHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/appliance/uplinks/usageHistory] getNetworkApplianceUplinksUsageHistory", response, response.Code())
	}
}

// NewGetNetworkApplianceUplinksUsageHistoryOK creates a GetNetworkApplianceUplinksUsageHistoryOK with default headers values
func NewGetNetworkApplianceUplinksUsageHistoryOK() *GetNetworkApplianceUplinksUsageHistoryOK {
	return &GetNetworkApplianceUplinksUsageHistoryOK{}
}

/*
GetNetworkApplianceUplinksUsageHistoryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceUplinksUsageHistoryOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network appliance uplinks usage history o k response has a 2xx status code
func (o *GetNetworkApplianceUplinksUsageHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance uplinks usage history o k response has a 3xx status code
func (o *GetNetworkApplianceUplinksUsageHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance uplinks usage history o k response has a 4xx status code
func (o *GetNetworkApplianceUplinksUsageHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance uplinks usage history o k response has a 5xx status code
func (o *GetNetworkApplianceUplinksUsageHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance uplinks usage history o k response a status code equal to that given
func (o *GetNetworkApplianceUplinksUsageHistoryOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network appliance uplinks usage history o k response
func (o *GetNetworkApplianceUplinksUsageHistoryOK) Code() int {
	return 200
}

func (o *GetNetworkApplianceUplinksUsageHistoryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/uplinks/usageHistory][%d] getNetworkApplianceUplinksUsageHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceUplinksUsageHistoryOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/uplinks/usageHistory][%d] getNetworkApplianceUplinksUsageHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceUplinksUsageHistoryOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceUplinksUsageHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
