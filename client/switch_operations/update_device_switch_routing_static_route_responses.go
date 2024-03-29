// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// UpdateDeviceSwitchRoutingStaticRouteReader is a Reader for the UpdateDeviceSwitchRoutingStaticRoute structure.
type UpdateDeviceSwitchRoutingStaticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateDeviceSwitchRoutingStaticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateDeviceSwitchRoutingStaticRouteOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /devices/{serial}/switch/routing/staticRoutes/{staticRouteId}] updateDeviceSwitchRoutingStaticRoute", response, response.Code())
	}
}

// NewUpdateDeviceSwitchRoutingStaticRouteOK creates a UpdateDeviceSwitchRoutingStaticRouteOK with default headers values
func NewUpdateDeviceSwitchRoutingStaticRouteOK() *UpdateDeviceSwitchRoutingStaticRouteOK {
	return &UpdateDeviceSwitchRoutingStaticRouteOK{}
}

/*
UpdateDeviceSwitchRoutingStaticRouteOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateDeviceSwitchRoutingStaticRouteOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update device switch routing static route o k response has a 2xx status code
func (o *UpdateDeviceSwitchRoutingStaticRouteOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update device switch routing static route o k response has a 3xx status code
func (o *UpdateDeviceSwitchRoutingStaticRouteOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update device switch routing static route o k response has a 4xx status code
func (o *UpdateDeviceSwitchRoutingStaticRouteOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update device switch routing static route o k response has a 5xx status code
func (o *UpdateDeviceSwitchRoutingStaticRouteOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update device switch routing static route o k response a status code equal to that given
func (o *UpdateDeviceSwitchRoutingStaticRouteOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update device switch routing static route o k response
func (o *UpdateDeviceSwitchRoutingStaticRouteOK) Code() int {
	return 200
}

func (o *UpdateDeviceSwitchRoutingStaticRouteOK) Error() string {
	return fmt.Sprintf("[PUT /devices/{serial}/switch/routing/staticRoutes/{staticRouteId}][%d] updateDeviceSwitchRoutingStaticRouteOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceSwitchRoutingStaticRouteOK) String() string {
	return fmt.Sprintf("[PUT /devices/{serial}/switch/routing/staticRoutes/{staticRouteId}][%d] updateDeviceSwitchRoutingStaticRouteOK  %+v", 200, o.Payload)
}

func (o *UpdateDeviceSwitchRoutingStaticRouteOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateDeviceSwitchRoutingStaticRouteOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateDeviceSwitchRoutingStaticRouteBody update device switch routing static route body
// Example: {"advertiseViaOspfEnabled":false,"name":"My route","nextHopIp":"1.2.3.4","preferOverOspfRoutesEnabled":false,"subnet":"192.168.1.0/24"}
swagger:model UpdateDeviceSwitchRoutingStaticRouteBody
*/
type UpdateDeviceSwitchRoutingStaticRouteBody struct {

	// Option to advertise static route via OSPF
	AdvertiseViaOspfEnabled bool `json:"advertiseViaOspfEnabled,omitempty"`

	// Name or description for layer 3 static route
	Name string `json:"name,omitempty"`

	// IP address of the next hop device to which the device sends its traffic for the subnet
	NextHopIP string `json:"nextHopIp,omitempty"`

	// Option to prefer static route over OSPF routes
	PreferOverOspfRoutesEnabled bool `json:"preferOverOspfRoutesEnabled,omitempty"`

	// The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
	Subnet string `json:"subnet,omitempty"`
}

// Validate validates this update device switch routing static route body
func (o *UpdateDeviceSwitchRoutingStaticRouteBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update device switch routing static route body based on context it is used
func (o *UpdateDeviceSwitchRoutingStaticRouteBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingStaticRouteBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateDeviceSwitchRoutingStaticRouteBody) UnmarshalBinary(b []byte) error {
	var res UpdateDeviceSwitchRoutingStaticRouteBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
