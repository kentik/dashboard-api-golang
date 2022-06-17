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

// GetOrganizationCameraCustomAnalyticsArtifactsReader is a Reader for the GetOrganizationCameraCustomAnalyticsArtifacts structure.
type GetOrganizationCameraCustomAnalyticsArtifactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationCameraCustomAnalyticsArtifactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationCameraCustomAnalyticsArtifactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationCameraCustomAnalyticsArtifactsOK creates a GetOrganizationCameraCustomAnalyticsArtifactsOK with default headers values
func NewGetOrganizationCameraCustomAnalyticsArtifactsOK() *GetOrganizationCameraCustomAnalyticsArtifactsOK {
	return &GetOrganizationCameraCustomAnalyticsArtifactsOK{}
}

/* GetOrganizationCameraCustomAnalyticsArtifactsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationCameraCustomAnalyticsArtifactsOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get organization camera custom analytics artifacts o k response has a 2xx status code
func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization camera custom analytics artifacts o k response has a 3xx status code
func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization camera custom analytics artifacts o k response has a 4xx status code
func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization camera custom analytics artifacts o k response has a 5xx status code
func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization camera custom analytics artifacts o k response a status code equal to that given
func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/camera/customAnalytics/artifacts][%d] getOrganizationCameraCustomAnalyticsArtifactsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/camera/customAnalytics/artifacts][%d] getOrganizationCameraCustomAnalyticsArtifactsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationCameraCustomAnalyticsArtifactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
