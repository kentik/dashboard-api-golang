// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSwitchStackRoutingStaticRoutesReader is a Reader for the GetNetworkSwitchStackRoutingStaticRoutes structure.
type GetNetworkSwitchStackRoutingStaticRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchStackRoutingStaticRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchStackRoutingStaticRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSwitchStackRoutingStaticRoutesOK creates a GetNetworkSwitchStackRoutingStaticRoutesOK with default headers values
func NewGetNetworkSwitchStackRoutingStaticRoutesOK() *GetNetworkSwitchStackRoutingStaticRoutesOK {
	return &GetNetworkSwitchStackRoutingStaticRoutesOK{}
}

/* GetNetworkSwitchStackRoutingStaticRoutesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchStackRoutingStaticRoutesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network switch stack routing static routes o k response has a 2xx status code
func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch stack routing static routes o k response has a 3xx status code
func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch stack routing static routes o k response has a 4xx status code
func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch stack routing static routes o k response has a 5xx status code
func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch stack routing static routes o k response a status code equal to that given
func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes][%d] getNetworkSwitchStackRoutingStaticRoutesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/stacks/{switchStackId}/routing/staticRoutes][%d] getNetworkSwitchStackRoutingStaticRoutesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkSwitchStackRoutingStaticRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
