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

// GetNetworkSwitchMtuReader is a Reader for the GetNetworkSwitchMtu structure.
type GetNetworkSwitchMtuReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchMtuReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchMtuOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSwitchMtuOK creates a GetNetworkSwitchMtuOK with default headers values
func NewGetNetworkSwitchMtuOK() *GetNetworkSwitchMtuOK {
	return &GetNetworkSwitchMtuOK{}
}

/* GetNetworkSwitchMtuOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchMtuOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network switch mtu o k response has a 2xx status code
func (o *GetNetworkSwitchMtuOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch mtu o k response has a 3xx status code
func (o *GetNetworkSwitchMtuOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch mtu o k response has a 4xx status code
func (o *GetNetworkSwitchMtuOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch mtu o k response has a 5xx status code
func (o *GetNetworkSwitchMtuOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch mtu o k response a status code equal to that given
func (o *GetNetworkSwitchMtuOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkSwitchMtuOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/mtu][%d] getNetworkSwitchMtuOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchMtuOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/mtu][%d] getNetworkSwitchMtuOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchMtuOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkSwitchMtuOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
