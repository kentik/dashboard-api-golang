// Code generated by go-swagger; DO NOT EDIT.

package appliance

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationApplianceSecurityEventsReader is a Reader for the GetOrganizationApplianceSecurityEvents structure.
type GetOrganizationApplianceSecurityEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationApplianceSecurityEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationApplianceSecurityEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/appliance/security/events] getOrganizationApplianceSecurityEvents", response, response.Code())
	}
}

// NewGetOrganizationApplianceSecurityEventsOK creates a GetOrganizationApplianceSecurityEventsOK with default headers values
func NewGetOrganizationApplianceSecurityEventsOK() *GetOrganizationApplianceSecurityEventsOK {
	return &GetOrganizationApplianceSecurityEventsOK{}
}

/*
GetOrganizationApplianceSecurityEventsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationApplianceSecurityEventsOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

// IsSuccess returns true when this get organization appliance security events o k response has a 2xx status code
func (o *GetOrganizationApplianceSecurityEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization appliance security events o k response has a 3xx status code
func (o *GetOrganizationApplianceSecurityEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization appliance security events o k response has a 4xx status code
func (o *GetOrganizationApplianceSecurityEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization appliance security events o k response has a 5xx status code
func (o *GetOrganizationApplianceSecurityEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization appliance security events o k response a status code equal to that given
func (o *GetOrganizationApplianceSecurityEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization appliance security events o k response
func (o *GetOrganizationApplianceSecurityEventsOK) Code() int {
	return 200
}

func (o *GetOrganizationApplianceSecurityEventsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/appliance/security/events][%d] getOrganizationApplianceSecurityEventsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationApplianceSecurityEventsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/appliance/security/events][%d] getOrganizationApplianceSecurityEventsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationApplianceSecurityEventsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationApplianceSecurityEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
