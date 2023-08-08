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

// GetDeviceCellularSimsReader is a Reader for the GetDeviceCellularSims structure.
type GetDeviceCellularSimsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCellularSimsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCellularSimsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}/cellular/sims] getDeviceCellularSims", response, response.Code())
	}
}

// NewGetDeviceCellularSimsOK creates a GetDeviceCellularSimsOK with default headers values
func NewGetDeviceCellularSimsOK() *GetDeviceCellularSimsOK {
	return &GetDeviceCellularSimsOK{}
}

/*
GetDeviceCellularSimsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceCellularSimsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device cellular sims o k response has a 2xx status code
func (o *GetDeviceCellularSimsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device cellular sims o k response has a 3xx status code
func (o *GetDeviceCellularSimsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device cellular sims o k response has a 4xx status code
func (o *GetDeviceCellularSimsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device cellular sims o k response has a 5xx status code
func (o *GetDeviceCellularSimsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device cellular sims o k response a status code equal to that given
func (o *GetDeviceCellularSimsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device cellular sims o k response
func (o *GetDeviceCellularSimsOK) Code() int {
	return 200
}

func (o *GetDeviceCellularSimsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/cellular/sims][%d] getDeviceCellularSimsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCellularSimsOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/cellular/sims][%d] getDeviceCellularSimsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCellularSimsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceCellularSimsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
