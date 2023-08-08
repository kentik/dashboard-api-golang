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

// GetDeviceReader is a Reader for the GetDevice structure.
type GetDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}] getDevice", response, response.Code())
	}
}

// NewGetDeviceOK creates a GetDeviceOK with default headers values
func NewGetDeviceOK() *GetDeviceOK {
	return &GetDeviceOK{}
}

/*
GetDeviceOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device o k response has a 2xx status code
func (o *GetDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device o k response has a 3xx status code
func (o *GetDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device o k response has a 4xx status code
func (o *GetDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device o k response has a 5xx status code
func (o *GetDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device o k response a status code equal to that given
func (o *GetDeviceOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device o k response
func (o *GetDeviceOK) Code() int {
	return 200
}

func (o *GetDeviceOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}][%d] getDeviceOK  %+v", 200, o.Payload)
}

func (o *GetDeviceOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}][%d] getDeviceOK  %+v", 200, o.Payload)
}

func (o *GetDeviceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
