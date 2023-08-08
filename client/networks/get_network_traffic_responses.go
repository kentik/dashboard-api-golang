// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkTrafficReader is a Reader for the GetNetworkTraffic structure.
type GetNetworkTrafficReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkTrafficReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkTrafficOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/traffic] getNetworkTraffic", response, response.Code())
	}
}

// NewGetNetworkTrafficOK creates a GetNetworkTrafficOK with default headers values
func NewGetNetworkTrafficOK() *GetNetworkTrafficOK {
	return &GetNetworkTrafficOK{}
}

/*
GetNetworkTrafficOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkTrafficOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network traffic o k response has a 2xx status code
func (o *GetNetworkTrafficOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network traffic o k response has a 3xx status code
func (o *GetNetworkTrafficOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network traffic o k response has a 4xx status code
func (o *GetNetworkTrafficOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network traffic o k response has a 5xx status code
func (o *GetNetworkTrafficOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network traffic o k response a status code equal to that given
func (o *GetNetworkTrafficOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network traffic o k response
func (o *GetNetworkTrafficOK) Code() int {
	return 200
}

func (o *GetNetworkTrafficOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/traffic][%d] getNetworkTrafficOK  %+v", 200, o.Payload)
}

func (o *GetNetworkTrafficOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/traffic][%d] getNetworkTrafficOK  %+v", 200, o.Payload)
}

func (o *GetNetworkTrafficOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkTrafficOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
