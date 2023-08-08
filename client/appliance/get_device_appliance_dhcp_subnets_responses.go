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

// GetDeviceApplianceDhcpSubnetsReader is a Reader for the GetDeviceApplianceDhcpSubnets structure.
type GetDeviceApplianceDhcpSubnetsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceApplianceDhcpSubnetsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceApplianceDhcpSubnetsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}/appliance/dhcp/subnets] getDeviceApplianceDhcpSubnets", response, response.Code())
	}
}

// NewGetDeviceApplianceDhcpSubnetsOK creates a GetDeviceApplianceDhcpSubnetsOK with default headers values
func NewGetDeviceApplianceDhcpSubnetsOK() *GetDeviceApplianceDhcpSubnetsOK {
	return &GetDeviceApplianceDhcpSubnetsOK{}
}

/*
GetDeviceApplianceDhcpSubnetsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceApplianceDhcpSubnetsOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get device appliance dhcp subnets o k response has a 2xx status code
func (o *GetDeviceApplianceDhcpSubnetsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device appliance dhcp subnets o k response has a 3xx status code
func (o *GetDeviceApplianceDhcpSubnetsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device appliance dhcp subnets o k response has a 4xx status code
func (o *GetDeviceApplianceDhcpSubnetsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device appliance dhcp subnets o k response has a 5xx status code
func (o *GetDeviceApplianceDhcpSubnetsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device appliance dhcp subnets o k response a status code equal to that given
func (o *GetDeviceApplianceDhcpSubnetsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device appliance dhcp subnets o k response
func (o *GetDeviceApplianceDhcpSubnetsOK) Code() int {
	return 200
}

func (o *GetDeviceApplianceDhcpSubnetsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/appliance/dhcp/subnets][%d] getDeviceApplianceDhcpSubnetsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceApplianceDhcpSubnetsOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/appliance/dhcp/subnets][%d] getDeviceApplianceDhcpSubnetsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceApplianceDhcpSubnetsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetDeviceApplianceDhcpSubnetsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
