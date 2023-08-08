// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// DeleteOrganizationAlertsProfileReader is a Reader for the DeleteOrganizationAlertsProfile structure.
type DeleteOrganizationAlertsProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteOrganizationAlertsProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteOrganizationAlertsProfileNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[DELETE /organizations/{organizationId}/alerts/profiles/{alertConfigId}] deleteOrganizationAlertsProfile", response, response.Code())
	}
}

// NewDeleteOrganizationAlertsProfileNoContent creates a DeleteOrganizationAlertsProfileNoContent with default headers values
func NewDeleteOrganizationAlertsProfileNoContent() *DeleteOrganizationAlertsProfileNoContent {
	return &DeleteOrganizationAlertsProfileNoContent{}
}

/*
DeleteOrganizationAlertsProfileNoContent describes a response with status code 204, with default header values.

Successful operation
*/
type DeleteOrganizationAlertsProfileNoContent struct {
}

// IsSuccess returns true when this delete organization alerts profile no content response has a 2xx status code
func (o *DeleteOrganizationAlertsProfileNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this delete organization alerts profile no content response has a 3xx status code
func (o *DeleteOrganizationAlertsProfileNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this delete organization alerts profile no content response has a 4xx status code
func (o *DeleteOrganizationAlertsProfileNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this delete organization alerts profile no content response has a 5xx status code
func (o *DeleteOrganizationAlertsProfileNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this delete organization alerts profile no content response a status code equal to that given
func (o *DeleteOrganizationAlertsProfileNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the delete organization alerts profile no content response
func (o *DeleteOrganizationAlertsProfileNoContent) Code() int {
	return 204
}

func (o *DeleteOrganizationAlertsProfileNoContent) Error() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/alerts/profiles/{alertConfigId}][%d] deleteOrganizationAlertsProfileNoContent ", 204)
}

func (o *DeleteOrganizationAlertsProfileNoContent) String() string {
	return fmt.Sprintf("[DELETE /organizations/{organizationId}/alerts/profiles/{alertConfigId}][%d] deleteOrganizationAlertsProfileNoContent ", 204)
}

func (o *DeleteOrganizationAlertsProfileNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}
