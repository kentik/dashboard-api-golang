// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationPolicyObjectReader is a Reader for the DeleteOrganizationPolicyObject structure.
type DeleteOrganizationPolicyObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationPolicyObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationPolicyObjectNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /organizations/{organizationId}/policyObjects/{policyObjectId}] deleteOrganizationPolicyObject", response, response.Code())
	}
}

// NewDeleteOrganizationPolicyObjectNoContent creates a DeleteOrganizationPolicyObjectNoContent with default headers values
func NewDeleteOrganizationPolicyObjectNoContent() *DeleteOrganizationPolicyObjectNoContent {
	return &DeleteOrganizationPolicyObjectNoContent{}
}

/*
DeleteOrganizationPolicyObjectNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteOrganizationPolicyObjectNoContent struct {
}

// IsSuccess returns true when this delete organization policy object no content response has a 2xx status code
func (o *DeleteOrganizationPolicyObjectNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization policy object no content response has a 3xx status code
func (o *DeleteOrganizationPolicyObjectNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization policy object no content response has a 4xx status code
func (o *DeleteOrganizationPolicyObjectNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization policy object no content response has a 5xx status code
func (o *DeleteOrganizationPolicyObjectNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization policy object no content response a status code equal to that given
func (o *DeleteOrganizationPolicyObjectNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization policy object no content response
func (o *DeleteOrganizationPolicyObjectNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationPolicyObjectNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/policyObjects/{policyObjectId}][%d] deleteOrganizationPolicyObjectNoContent ", 204)
}

func (o *DeleteOrganizationPolicyObjectNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/policyObjects/{policyObjectId}][%d] deleteOrganizationPolicyObjectNoContent ", 204)
}

func (o *DeleteOrganizationPolicyObjectNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
