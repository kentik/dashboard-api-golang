// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetOrganizationPolicyObjectsGroupsReader is a Reader for the GetOrganizationPolicyObjectsGroups structure.
type GetOrganizationPolicyObjectsGroupsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationPolicyObjectsGroupsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationPolicyObjectsGroupsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/policyObjects/groups] getOrganizationPolicyObjectsGroups", response, response.Code())
	}
}

// NewGetOrganizationPolicyObjectsGroupsOK creates a GetOrganizationPolicyObjectsGroupsOK with default headers values
func NewGetOrganizationPolicyObjectsGroupsOK() *GetOrganizationPolicyObjectsGroupsOK {
	return &GetOrganizationPolicyObjectsGroupsOK{}
}

/*
GetOrganizationPolicyObjectsGroupsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationPolicyObjectsGroupsOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []interface{}
}

// IsSuccess returns true when this get organization policy objects groups o k response has a 2xx status code
func (o *GetOrganizationPolicyObjectsGroupsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization policy objects groups o k response has a 3xx status code
func (o *GetOrganizationPolicyObjectsGroupsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization policy objects groups o k response has a 4xx status code
func (o *GetOrganizationPolicyObjectsGroupsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization policy objects groups o k response has a 5xx status code
func (o *GetOrganizationPolicyObjectsGroupsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization policy objects groups o k response a status code equal to that given
func (o *GetOrganizationPolicyObjectsGroupsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization policy objects groups o k response
func (o *GetOrganizationPolicyObjectsGroupsOK) Code() int {
	return 200
}

func (o *GetOrganizationPolicyObjectsGroupsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/policyObjects/groups][%d] getOrganizationPolicyObjectsGroupsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationPolicyObjectsGroupsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/policyObjects/groups][%d] getOrganizationPolicyObjectsGroupsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationPolicyObjectsGroupsOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationPolicyObjectsGroupsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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
