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

// GetNetworkApplianceFirewallL7FirewallRulesReader is a Reader for the GetNetworkApplianceFirewallL7FirewallRules structure.
type GetNetworkApplianceFirewallL7FirewallRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkApplianceFirewallL7FirewallRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkApplianceFirewallL7FirewallRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkApplianceFirewallL7FirewallRulesOK creates a GetNetworkApplianceFirewallL7FirewallRulesOK with default headers values
func NewGetNetworkApplianceFirewallL7FirewallRulesOK() *GetNetworkApplianceFirewallL7FirewallRulesOK {
	return &GetNetworkApplianceFirewallL7FirewallRulesOK{}
}

/* GetNetworkApplianceFirewallL7FirewallRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkApplianceFirewallL7FirewallRulesOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network appliance firewall l7 firewall rules o k response has a 2xx status code
func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance firewall l7 firewall rules o k response has a 3xx status code
func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance firewall l7 firewall rules o k response has a 4xx status code
func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance firewall l7 firewall rules o k response has a 5xx status code
func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance firewall l7 firewall rules o k response a status code equal to that given
func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/firewall/l7FirewallRules][%d] getNetworkApplianceFirewallL7FirewallRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/firewall/l7FirewallRules][%d] getNetworkApplianceFirewallL7FirewallRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkApplianceFirewallL7FirewallRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
