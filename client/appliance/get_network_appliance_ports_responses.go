// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// GetNetworkAppliancePortsReader is a Reader for the GetNetworkAppliancePorts structure.
type GetNetworkAppliancePortsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkAppliancePortsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkAppliancePortsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/appliance/ports] getNetworkAppliancePorts", response, response.Code())
	}
}

// NewGetNetworkAppliancePortsOK creates a GetNetworkAppliancePortsOK with default headers values
func NewGetNetworkAppliancePortsOK() *GetNetworkAppliancePortsOK {
	return &GetNetworkAppliancePortsOK{}
}

/*
GetNetworkAppliancePortsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkAppliancePortsOK struct {
	Payload []*GetNetworkAppliancePortsOKBodyItems0
}

// IsSuccess returns true when this get network appliance ports o k response has a 2xx status code
func (o *GetNetworkAppliancePortsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network appliance ports o k response has a 3xx status code
func (o *GetNetworkAppliancePortsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network appliance ports o k response has a 4xx status code
func (o *GetNetworkAppliancePortsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network appliance ports o k response has a 5xx status code
func (o *GetNetworkAppliancePortsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network appliance ports o k response a status code equal to that given
func (o *GetNetworkAppliancePortsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network appliance ports o k response
func (o *GetNetworkAppliancePortsOK) Code() int {
	return 200
}

func (o *GetNetworkAppliancePortsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/ports][%d] getNetworkAppliancePortsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAppliancePortsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/appliance/ports][%d] getNetworkAppliancePortsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkAppliancePortsOK) GetPayload() []*GetNetworkAppliancePortsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkAppliancePortsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkAppliancePortsOKBodyItems0 get network appliance ports o k body items0
swagger:model GetNetworkAppliancePortsOKBodyItems0
*/
type GetNetworkAppliancePortsOKBodyItems0 struct {

	// The name of the policy. Only applicable to Access ports.
	AccessPolicy string `json:"accessPolicy,omitempty"`

	// Comma-delimited list of the VLAN ID's allowed on the port, or 'all' to permit all VLAN's on the port.
	AllowedVlans string `json:"allowedVlans,omitempty"`

	// Whether the trunk port can drop all untagged traffic.
	DropUntaggedTraffic bool `json:"dropUntaggedTraffic,omitempty"`

	// The status of the port
	Enabled bool `json:"enabled,omitempty"`

	// Number of the port
	Number int64 `json:"number,omitempty"`

	// The type of the port: 'access' or 'trunk'.
	Type string `json:"type,omitempty"`

	// Native VLAN when the port is in Trunk mode. Access VLAN when the port is in Access mode.
	Vlan int64 `json:"vlan,omitempty"`
}

// Validate validates this get network appliance ports o k body items0
func (o *GetNetworkAppliancePortsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network appliance ports o k body items0 based on context it is used
func (o *GetNetworkAppliancePortsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkAppliancePortsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkAppliancePortsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkAppliancePortsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
