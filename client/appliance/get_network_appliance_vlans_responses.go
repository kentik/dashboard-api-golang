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

// GetNetworkApplianceVlansReader is a Reader for the GetNetworkApplianceVlans structure.
type GetNetworkApplianceVlansReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceVlansReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceVlansOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceVlansOK creates a GetNetworkApplianceVlansOK with default headers values
func NewGetNetworkApplianceVlansOK() *GetNetworkApplianceVlansOK {
	return &GetNetworkApplianceVlansOK{}
}

/* GetNetworkApplianceVlansOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceVlansOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network appliance vlans o k response has a 2xx status code
func (o *GetNetworkApplianceVlansOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance vlans o k response has a 3xx status code
func (o *GetNetworkApplianceVlansOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance vlans o k response has a 4xx status code
func (o *GetNetworkApplianceVlansOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance vlans o k response has a 5xx status code
func (o *GetNetworkApplianceVlansOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance vlans o k response a status code equal to that given
func (o *GetNetworkApplianceVlansOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkApplianceVlansOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/vlans][%d] getNetworkApplianceVlansOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceVlansOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/vlans][%d] getNetworkApplianceVlansOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceVlansOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceVlansOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
