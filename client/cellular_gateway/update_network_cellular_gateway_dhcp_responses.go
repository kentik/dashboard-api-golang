// Code generated by go-swagger; DO NOT EDIT.

package cellular_gateway

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

// UpdateNetworkCellularGatewayDhcpReader is a Reader for the UpdateNetworkCellularGatewayDhcp structure.
type UpdateNetworkCellularGatewayDhcpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkCellularGatewayDhcpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkCellularGatewayDhcpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkCellularGatewayDhcpOK creates a UpdateNetworkCellularGatewayDhcpOK with default headers values
func NewUpdateNetworkCellularGatewayDhcpOK() *UpdateNetworkCellularGatewayDhcpOK {
	return &UpdateNetworkCellularGatewayDhcpOK{}
}

/* UpdateNetworkCellularGatewayDhcpOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkCellularGatewayDhcpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network cellular gateway dhcp o k response has a 2xx status code
func (o *UpdateNetworkCellularGatewayDhcpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network cellular gateway dhcp o k response has a 3xx status code
func (o *UpdateNetworkCellularGatewayDhcpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network cellular gateway dhcp o k response has a 4xx status code
func (o *UpdateNetworkCellularGatewayDhcpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network cellular gateway dhcp o k response has a 5xx status code
func (o *UpdateNetworkCellularGatewayDhcpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network cellular gateway dhcp o k response a status code equal to that given
func (o *UpdateNetworkCellularGatewayDhcpOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkCellularGatewayDhcpOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/cellularGateway/dhcp][%d] updateNetworkCellularGatewayDhcpOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkCellularGatewayDhcpOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/cellularGateway/dhcp][%d] updateNetworkCellularGatewayDhcpOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkCellularGatewayDhcpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkCellularGatewayDhcpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkCellularGatewayDhcpBody update network cellular gateway dhcp body
// Example: {"dhcpLeaseTime":"1 hour","dnsCustomNameservers":["172.16.2.111","172.16.2.30"],"dnsNameservers":"custom"}
swagger:model UpdateNetworkCellularGatewayDhcpBody
*/
type UpdateNetworkCellularGatewayDhcpBody struct {

	// DHCP Lease time for all MG of the network. It can be '30 minutes', '1 hour', '4 hours', '12 hours', '1 day' or '1 week'.
	DhcpLeaseTime string `json:"dhcpLeaseTime,omitempty"`

	// list of fixed IP representing the the DNS Name servers when the mode is 'custom'
	DNSCustomNameservers []string `json:"dnsCustomNameservers"`

	// DNS name servers mode for all MG of the network. It can take 4 different values: 'upstream_dns', 'google_dns', 'opendns', 'custom'.
	DNSNameservers string `json:"dnsNameservers,omitempty"`
}

// Validate validates this update network cellular gateway dhcp body
func (o *UpdateNetworkCellularGatewayDhcpBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network cellular gateway dhcp body based on context it is used
func (o *UpdateNetworkCellularGatewayDhcpBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayDhcpBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkCellularGatewayDhcpBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkCellularGatewayDhcpBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
