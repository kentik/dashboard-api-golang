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

// GetDeviceLiveToolsPingDeviceReader is a Reader for the GetDeviceLiveToolsPingDevice structure.
type GetDeviceLiveToolsPingDeviceReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceLiveToolsPingDeviceReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceLiveToolsPingDeviceOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceLiveToolsPingDeviceOK creates a GetDeviceLiveToolsPingDeviceOK with default headers values
func NewGetDeviceLiveToolsPingDeviceOK() *GetDeviceLiveToolsPingDeviceOK {
	return &GetDeviceLiveToolsPingDeviceOK{}
}

/* GetDeviceLiveToolsPingDeviceOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceLiveToolsPingDeviceOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device live tools ping device o k response has a 2xx status code
func (o *GetDeviceLiveToolsPingDeviceOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device live tools ping device o k response has a 3xx status code
func (o *GetDeviceLiveToolsPingDeviceOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device live tools ping device o k response has a 4xx status code
func (o *GetDeviceLiveToolsPingDeviceOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device live tools ping device o k response has a 5xx status code
func (o *GetDeviceLiveToolsPingDeviceOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device live tools ping device o k response a status code equal to that given
func (o *GetDeviceLiveToolsPingDeviceOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceLiveToolsPingDeviceOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/liveTools/pingDevice/{id}][%d] getDeviceLiveToolsPingDeviceOK  %+v", 200, o.Payload)
}

func (o *GetDeviceLiveToolsPingDeviceOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/liveTools/pingDevice/{id}][%d] getDeviceLiveToolsPingDeviceOK  %+v", 200, o.Payload)
}

func (o *GetDeviceLiveToolsPingDeviceOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceLiveToolsPingDeviceOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
