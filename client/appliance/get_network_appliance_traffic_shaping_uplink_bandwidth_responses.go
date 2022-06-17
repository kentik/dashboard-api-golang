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

// GetNetworkApplianceTrafficShapingUplinkBandwidthReader is a Reader for the GetNetworkApplianceTrafficShapingUplinkBandwidth structure.
type GetNetworkApplianceTrafficShapingUplinkBandwidthReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceTrafficShapingUplinkBandwidthOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceTrafficShapingUplinkBandwidthOK creates a GetNetworkApplianceTrafficShapingUplinkBandwidthOK with default headers values
func NewGetNetworkApplianceTrafficShapingUplinkBandwidthOK() *GetNetworkApplianceTrafficShapingUplinkBandwidthOK {
	return &GetNetworkApplianceTrafficShapingUplinkBandwidthOK{}
}

/* GetNetworkApplianceTrafficShapingUplinkBandwidthOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceTrafficShapingUplinkBandwidthOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network appliance traffic shaping uplink bandwidth o k response has a 2xx status code
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance traffic shaping uplink bandwidth o k response has a 3xx status code
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance traffic shaping uplink bandwidth o k response has a 4xx status code
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance traffic shaping uplink bandwidth o k response has a 5xx status code
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance traffic shaping uplink bandwidth o k response a status code equal to that given
func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth][%d] getNetworkApplianceTrafficShapingUplinkBandwidthOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/trafficShaping/uplinkBandwidth][%d] getNetworkApplianceTrafficShapingUplinkBandwidthOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceTrafficShapingUplinkBandwidthOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}