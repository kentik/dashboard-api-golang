// Code generated by go-swagger; DO NOT EDIT.

package camera

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceCameraAnalyticsZonesReader is a Reader for the GetDeviceCameraAnalyticsZones structure.
type GetDeviceCameraAnalyticsZonesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCameraAnalyticsZonesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCameraAnalyticsZonesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}/camera/analytics/zones] getDeviceCameraAnalyticsZones", response, response.Code())
	}
}

// NewGetDeviceCameraAnalyticsZonesOK creates a GetDeviceCameraAnalyticsZonesOK with default headers values
func NewGetDeviceCameraAnalyticsZonesOK() *GetDeviceCameraAnalyticsZonesOK {
	return &GetDeviceCameraAnalyticsZonesOK{}
}

/*
GetDeviceCameraAnalyticsZonesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceCameraAnalyticsZonesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get device camera analytics zones o k response has a 2xx status code
func (o *GetDeviceCameraAnalyticsZonesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device camera analytics zones o k response has a 3xx status code
func (o *GetDeviceCameraAnalyticsZonesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device camera analytics zones o k response has a 4xx status code
func (o *GetDeviceCameraAnalyticsZonesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device camera analytics zones o k response has a 5xx status code
func (o *GetDeviceCameraAnalyticsZonesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device camera analytics zones o k response a status code equal to that given
func (o *GetDeviceCameraAnalyticsZonesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device camera analytics zones o k response
func (o *GetDeviceCameraAnalyticsZonesOK) Code() int {
	return 200
}

func (o *GetDeviceCameraAnalyticsZonesOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/analytics/zones][%d] getDeviceCameraAnalyticsZonesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraAnalyticsZonesOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/analytics/zones][%d] getDeviceCameraAnalyticsZonesOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraAnalyticsZonesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetDeviceCameraAnalyticsZonesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
