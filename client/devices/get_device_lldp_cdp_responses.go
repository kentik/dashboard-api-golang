// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceLldpCdpReader is a Reader for the GetDeviceLldpCdp structure.
type GetDeviceLldpCdpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceLldpCdpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceLldpCdpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceLldpCdpOK creates a GetDeviceLldpCdpOK with default headers values
func NewGetDeviceLldpCdpOK() *GetDeviceLldpCdpOK {
	return &GetDeviceLldpCdpOK{}
}

/* GetDeviceLldpCdpOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceLldpCdpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device lldp cdp o k response has a 2xx status code
func (o *GetDeviceLldpCdpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device lldp cdp o k response has a 3xx status code
func (o *GetDeviceLldpCdpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device lldp cdp o k response has a 4xx status code
func (o *GetDeviceLldpCdpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device lldp cdp o k response has a 5xx status code
func (o *GetDeviceLldpCdpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device lldp cdp o k response a status code equal to that given
func (o *GetDeviceLldpCdpOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceLldpCdpOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/lldpCdp][%d] getDeviceLldpCdpOK  %+v", 200, o.Payload)
}

func (o *GetDeviceLldpCdpOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/lldpCdp][%d] getDeviceLldpCdpOK  %+v", 200, o.Payload)
}

func (o *GetDeviceLldpCdpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceLldpCdpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
