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

// GetDeviceCameraAnalyticsLiveReader is a Reader for the GetDeviceCameraAnalyticsLive structure.
type GetDeviceCameraAnalyticsLiveReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceCameraAnalyticsLiveReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceCameraAnalyticsLiveOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /devices/{serial}/camera/analytics/live] getDeviceCameraAnalyticsLive", response, response.Code())
	}
}

// NewGetDeviceCameraAnalyticsLiveOK creates a GetDeviceCameraAnalyticsLiveOK with default headers values
func NewGetDeviceCameraAnalyticsLiveOK() *GetDeviceCameraAnalyticsLiveOK {
	return &GetDeviceCameraAnalyticsLiveOK{}
}

/*
GetDeviceCameraAnalyticsLiveOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceCameraAnalyticsLiveOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device camera analytics live o k response has a 2xx status code
func (o *GetDeviceCameraAnalyticsLiveOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device camera analytics live o k response has a 3xx status code
func (o *GetDeviceCameraAnalyticsLiveOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device camera analytics live o k response has a 4xx status code
func (o *GetDeviceCameraAnalyticsLiveOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device camera analytics live o k response has a 5xx status code
func (o *GetDeviceCameraAnalyticsLiveOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device camera analytics live o k response a status code equal to that given
func (o *GetDeviceCameraAnalyticsLiveOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get device camera analytics live o k response
func (o *GetDeviceCameraAnalyticsLiveOK) Code() int {
	return 200
}

func (o *GetDeviceCameraAnalyticsLiveOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/analytics/live][%d] getDeviceCameraAnalyticsLiveOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraAnalyticsLiveOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/camera/analytics/live][%d] getDeviceCameraAnalyticsLiveOK  %+v", 200, o.Payload)
}

func (o *GetDeviceCameraAnalyticsLiveOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceCameraAnalyticsLiveOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
