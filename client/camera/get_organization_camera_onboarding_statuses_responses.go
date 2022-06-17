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

// GetOrganizationCameraOnboardingStatusesReader is a Reader for the GetOrganizationCameraOnboardingStatuses structure.
type GetOrganizationCameraOnboardingStatusesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationCameraOnboardingStatusesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationCameraOnboardingStatusesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationCameraOnboardingStatusesOK creates a GetOrganizationCameraOnboardingStatusesOK with default headers values
func NewGetOrganizationCameraOnboardingStatusesOK() *GetOrganizationCameraOnboardingStatusesOK {
	return &GetOrganizationCameraOnboardingStatusesOK{}
}

/* GetOrganizationCameraOnboardingStatusesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationCameraOnboardingStatusesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get organization camera onboarding statuses o k response has a 2xx status code
func (o *GetOrganizationCameraOnboardingStatusesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization camera onboarding statuses o k response has a 3xx status code
func (o *GetOrganizationCameraOnboardingStatusesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization camera onboarding statuses o k response has a 4xx status code
func (o *GetOrganizationCameraOnboardingStatusesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization camera onboarding statuses o k response has a 5xx status code
func (o *GetOrganizationCameraOnboardingStatusesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization camera onboarding statuses o k response a status code equal to that given
func (o *GetOrganizationCameraOnboardingStatusesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationCameraOnboardingStatusesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/camera/onboarding/statuses][%d] getOrganizationCameraOnboardingStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationCameraOnboardingStatusesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/camera/onboarding/statuses][%d] getOrganizationCameraOnboardingStatusesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationCameraOnboardingStatusesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationCameraOnboardingStatusesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
