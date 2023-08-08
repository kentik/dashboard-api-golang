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

// GetOrganizationPolicyObjectReader is a Reader for the GetOrganizationPolicyObject structure.
type GetOrganizationPolicyObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationPolicyObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationPolicyObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/policyObjects/{policyObjectId}] getOrganizationPolicyObject", response, response.Code())
	}
}

// NewGetOrganizationPolicyObjectOK creates a GetOrganizationPolicyObjectOK with default headers values
func NewGetOrganizationPolicyObjectOK() *GetOrganizationPolicyObjectOK {
	return &GetOrganizationPolicyObjectOK{}
}

/*
GetOrganizationPolicyObjectOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationPolicyObjectOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization policy object o k response has a 2xx status code
func (o *GetOrganizationPolicyObjectOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization policy object o k response has a 3xx status code
func (o *GetOrganizationPolicyObjectOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization policy object o k response has a 4xx status code
func (o *GetOrganizationPolicyObjectOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization policy object o k response has a 5xx status code
func (o *GetOrganizationPolicyObjectOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization policy object o k response a status code equal to that given
func (o *GetOrganizationPolicyObjectOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization policy object o k response
func (o *GetOrganizationPolicyObjectOK) Code() int {
	return 200
}

func (o *GetOrganizationPolicyObjectOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/policyObjects/{policyObjectId}][%d] getOrganizationPolicyObjectOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationPolicyObjectOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/policyObjects/{policyObjectId}][%d] getOrganizationPolicyObjectOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationPolicyObjectOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationPolicyObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
