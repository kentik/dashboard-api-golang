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

// GetNetworkCameraQualityRetentionProfileReader is a Reader for the GetNetworkCameraQualityRetentionProfile structure.
type GetNetworkCameraQualityRetentionProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkCameraQualityRetentionProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkCameraQualityRetentionProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}] getNetworkCameraQualityRetentionProfile", response, response.Code())
	}
}

// NewGetNetworkCameraQualityRetentionProfileOK creates a GetNetworkCameraQualityRetentionProfileOK with default headers values
func NewGetNetworkCameraQualityRetentionProfileOK() *GetNetworkCameraQualityRetentionProfileOK {
	return &GetNetworkCameraQualityRetentionProfileOK{}
}

/*
GetNetworkCameraQualityRetentionProfileOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkCameraQualityRetentionProfileOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network camera quality retention profile o k response has a 2xx status code
func (o *GetNetworkCameraQualityRetentionProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network camera quality retention profile o k response has a 3xx status code
func (o *GetNetworkCameraQualityRetentionProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network camera quality retention profile o k response has a 4xx status code
func (o *GetNetworkCameraQualityRetentionProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network camera quality retention profile o k response has a 5xx status code
func (o *GetNetworkCameraQualityRetentionProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network camera quality retention profile o k response a status code equal to that given
func (o *GetNetworkCameraQualityRetentionProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network camera quality retention profile o k response
func (o *GetNetworkCameraQualityRetentionProfileOK) Code() int {
	return 200
}

func (o *GetNetworkCameraQualityRetentionProfileOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}][%d] getNetworkCameraQualityRetentionProfileOK  %+v", 200, o.Payload)
}

func (o *GetNetworkCameraQualityRetentionProfileOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/camera/qualityRetentionProfiles/{qualityRetentionProfileId}][%d] getNetworkCameraQualityRetentionProfileOK  %+v", 200, o.Payload)
}

func (o *GetNetworkCameraQualityRetentionProfileOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkCameraQualityRetentionProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
