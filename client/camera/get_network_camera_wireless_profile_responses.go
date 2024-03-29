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

// GetNetworkCameraWirelessProfileReader is a Reader for the GetNetworkCameraWirelessProfile structure.
type GetNetworkCameraWirelessProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkCameraWirelessProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkCameraWirelessProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}] getNetworkCameraWirelessProfile", response, response.Code())
	}
}

// NewGetNetworkCameraWirelessProfileOK creates a GetNetworkCameraWirelessProfileOK with default headers values
func NewGetNetworkCameraWirelessProfileOK() *GetNetworkCameraWirelessProfileOK {
	return &GetNetworkCameraWirelessProfileOK{}
}

/*
GetNetworkCameraWirelessProfileOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkCameraWirelessProfileOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network camera wireless profile o k response has a 2xx status code
func (o *GetNetworkCameraWirelessProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network camera wireless profile o k response has a 3xx status code
func (o *GetNetworkCameraWirelessProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network camera wireless profile o k response has a 4xx status code
func (o *GetNetworkCameraWirelessProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network camera wireless profile o k response has a 5xx status code
func (o *GetNetworkCameraWirelessProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network camera wireless profile o k response a status code equal to that given
func (o *GetNetworkCameraWirelessProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network camera wireless profile o k response
func (o *GetNetworkCameraWirelessProfileOK) Code() int {
	return 200
}

func (o *GetNetworkCameraWirelessProfileOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}][%d] getNetworkCameraWirelessProfileOK  %+v", 200, o.Payload)
}

func (o *GetNetworkCameraWirelessProfileOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/camera/wirelessProfiles/{wirelessProfileId}][%d] getNetworkCameraWirelessProfileOK  %+v", 200, o.Payload)
}

func (o *GetNetworkCameraWirelessProfileOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkCameraWirelessProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
