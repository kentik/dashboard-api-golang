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

// GetDeviceSwitchWarmSpareReader is a Reader for the GetDeviceSwitchWarmSpare structure.
type GetDeviceSwitchWarmSpareReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceSwitchWarmSpareReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceSwitchWarmSpareOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}/switch/warmSpare] getDeviceSwitchWarmSpare", response, response.Code())
	}
}

// NewGetDeviceSwitchWarmSpareOK creates a GetDeviceSwitchWarmSpareOK with default headers values
func NewGetDeviceSwitchWarmSpareOK() *GetDeviceSwitchWarmSpareOK {
	return &GetDeviceSwitchWarmSpareOK{}
}

/*
GetDeviceSwitchWarmSpareOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceSwitchWarmSpareOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device switch warm spare o k response has a 2xx status code
func (o *GetDeviceSwitchWarmSpareOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device switch warm spare o k response has a 3xx status code
func (o *GetDeviceSwitchWarmSpareOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device switch warm spare o k response has a 4xx status code
func (o *GetDeviceSwitchWarmSpareOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device switch warm spare o k response has a 5xx status code
func (o *GetDeviceSwitchWarmSpareOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device switch warm spare o k response a status code equal to that given
func (o *GetDeviceSwitchWarmSpareOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device switch warm spare o k response
func (o *GetDeviceSwitchWarmSpareOK) Code() int {
	return 200
}

func (o *GetDeviceSwitchWarmSpareOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/warmSpare][%d] getDeviceSwitchWarmSpareOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchWarmSpareOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/switch/warmSpare][%d] getDeviceSwitchWarmSpareOK  %+v", 200, o.Payload)
}

func (o *GetDeviceSwitchWarmSpareOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceSwitchWarmSpareOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
