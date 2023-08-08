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

// GetDeviceClientsReader is a Reader for the GetDeviceClients structure.
type GetDeviceClientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceClientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceClientsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}/clients] getDeviceClients", response, response.Code())
	}
}

// NewGetDeviceClientsOK creates a GetDeviceClientsOK with default headers values
func NewGetDeviceClientsOK() *GetDeviceClientsOK {
	return &GetDeviceClientsOK{}
}

/*
GetDeviceClientsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceClientsOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get device clients o k response has a 2xx status code
func (o *GetDeviceClientsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device clients o k response has a 3xx status code
func (o *GetDeviceClientsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device clients o k response has a 4xx status code
func (o *GetDeviceClientsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device clients o k response has a 5xx status code
func (o *GetDeviceClientsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device clients o k response a status code equal to that given
func (o *GetDeviceClientsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device clients o k response
func (o *GetDeviceClientsOK) Code() int {
	return 200
}

func (o *GetDeviceClientsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/clients][%d] getDeviceClientsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceClientsOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/clients][%d] getDeviceClientsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceClientsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetDeviceClientsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
