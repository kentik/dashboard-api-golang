// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationAdaptivePolicyPolicyReader is a Reader for the DeleteOrganizationAdaptivePolicyPolicy structure.
type DeleteOrganizationAdaptivePolicyPolicyReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationAdaptivePolicyPolicyReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationAdaptivePolicyPolicyNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /organizations/{organizationId}/adaptivePolicy/policies/{id}] deleteOrganizationAdaptivePolicyPolicy", response, response.Code())
	}
}

// NewDeleteOrganizationAdaptivePolicyPolicyNoContent creates a DeleteOrganizationAdaptivePolicyPolicyNoContent with default headers values
func NewDeleteOrganizationAdaptivePolicyPolicyNoContent() *DeleteOrganizationAdaptivePolicyPolicyNoContent {
	return &DeleteOrganizationAdaptivePolicyPolicyNoContent{}
}

/*
DeleteOrganizationAdaptivePolicyPolicyNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteOrganizationAdaptivePolicyPolicyNoContent struct {
}

// IsSuccess returns true when this delete organization adaptive policy policy no content response has a 2xx status code
func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization adaptive policy policy no content response has a 3xx status code
func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization adaptive policy policy no content response has a 4xx status code
func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization adaptive policy policy no content response has a 5xx status code
func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization adaptive policy policy no content response a status code equal to that given
func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization adaptive policy policy no content response
func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/adaptivePolicy/policies/{id}][%d] deleteOrganizationAdaptivePolicyPolicyNoContent ", 204)
}

func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/adaptivePolicy/policies/{id}][%d] deleteOrganizationAdaptivePolicyPolicyNoContent ", 204)
}

func (o *DeleteOrganizationAdaptivePolicyPolicyNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
