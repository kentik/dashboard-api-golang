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

// GetDeviceSwitchRoutingStaticRoutesReader is a Reader for the GetDeviceSwitchRoutingStaticRoutes structure.
type GetDeviceSwitchRoutingStaticRoutesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceSwitchRoutingStaticRoutesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceSwitchRoutingStaticRoutesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceSwitchRoutingStaticRoutesOK creates a GetDeviceSwitchRoutingStaticRoutesOK with default headers values
func NewGetDeviceSwitchRoutingStaticRoutesOK() *GetDeviceSwitchRoutingStaticRoutesOK {
	return &GetDeviceSwitchRoutingStaticRoutesOK{}
}

/* GetDeviceSwitchRoutingStaticRoutesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceSwitchRoutingStaticRoutesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get device switch routing static routes o k response has a 2xx status code
func (o *GetDeviceSwitchRoutingStaticRoutesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device switch routing static routes o k response has a 3xx status code
func (o *GetDeviceSwitchRoutingStaticRoutesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device switch routing static routes o k response has a 4xx status code
func (o *GetDeviceSwitchRoutingStaticRoutesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device switch routing static routes o k response has a 5xx status code
func (o *GetDeviceSwitchRoutingStaticRoutesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device switch routing static routes o k response a status code equal to that given
func (o *GetDeviceSwitchRoutingStaticRoutesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceSwitchRoutingStaticRoutesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/routing/staticRoutes][%d] getDeviceSwitchRoutingStaticRoutesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchRoutingStaticRoutesOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/routing/staticRoutes][%d] getDeviceSwitchRoutingStaticRoutesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchRoutingStaticRoutesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetDeviceSwitchRoutingStaticRoutesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
