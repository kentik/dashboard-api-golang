// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateNetworkSwitchStackRoutingStaticRouteReader is a Reader for the CreateNetworkSwitchStackRoutingStaticRoute structure.
type CreateNetworkSwitchStackRoutingStaticRouteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchStackRoutingStaticRouteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchStackRoutingStaticRouteCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes] createNetworkSwitchStackRoutingStaticRoute", response, response.Code())
	}
}

// NewCreateNetworkSwitchStackRoutingStaticRouteCreated creates a CreateNetworkSwitchStackRoutingStaticRouteCreated with default headers values
func NewCreateNetworkSwitchStackRoutingStaticRouteCreated() *CreateNetworkSwitchStackRoutingStaticRouteCreated {
	return &CreateNetworkSwitchStackRoutingStaticRouteCreated{}
}

/*
CreateNetworkSwitchStackRoutingStaticRouteCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchStackRoutingStaticRouteCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network switch stack routing static route created response has a 2xx status code
func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network switch stack routing static route created response has a 3xx status code
func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network switch stack routing static route created response has a 4xx status code
func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network switch stack routing static route created response has a 5xx status code
func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network switch stack routing static route created response a status code equal to that given
func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network switch stack routing static route created response
func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) Code() int {
	return 201
}

func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes][%d] createNetworkSwitchStackRoutingStaticRouteCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes][%d] createNetworkSwitchStackRoutingStaticRouteCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchStackRoutingStaticRouteCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkSwitchStackRoutingStaticRouteBody create network switch stack routing static route body
// Example: {"advertiseViaOspfEnabled":false,"name":"My route","nextHopIp":"1.2.3.4","preferOverOspfRoutesEnabled":false,"subnet":"192.168.1.0/24"}
swagger:model CreateNetworkSwitchStackRoutingStaticRouteBody
*/
type CreateNetworkSwitchStackRoutingStaticRouteBody struct {

	// Option to advertise static route via OSPF
	AdvertiseViaOspfEnabled bool `json:"advertiseViaOspfEnabled,omitempty"`

	// Name or description for layer 3 static route
	Name string `json:"name,omitempty"`

	// IP address of the next hop device to which the device sends its traffic for the subnet
	// Required: true
	NextHopIP *string `json:"nextHopIp"`

	// Option to prefer static route over OSPF routes
	PreferOverOspfRoutesEnabled bool `json:"preferOverOspfRoutesEnabled,omitempty"`

	// The subnet which is routed via this static route and should be specified in CIDR notation (ex. 1.2.3.0/24)
	// Required: true
	Subnet *string `json:"subnet"`
}

// Validate validates this create network switch stack routing static route body
func (o *CreateNetworkSwitchStackRoutingStaticRouteBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNextHopIP(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSubnet(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchStackRoutingStaticRouteBody) validateNextHopIP(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchStackRoutingStaticRoute"+"."+"nextHopIp", "body", o.NextHopIP); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchStackRoutingStaticRouteBody) validateSubnet(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchStackRoutingStaticRoute"+"."+"subnet", "body", o.Subnet); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network switch stack routing static route body based on context it is used
func (o *CreateNetworkSwitchStackRoutingStaticRouteBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingStaticRouteBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchStackRoutingStaticRouteBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchStackRoutingStaticRouteBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
