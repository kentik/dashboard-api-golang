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

// GetOrganizationAdaptivePolicyPoliciesReader is a Reader for the GetOrganizationAdaptivePolicyPolicies structure.
type GetOrganizationAdaptivePolicyPoliciesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAdaptivePolicyPoliciesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAdaptivePolicyPoliciesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/adaptivePolicy/policies] getOrganizationAdaptivePolicyPolicies", response, response.Code())
	}
}

// NewGetOrganizationAdaptivePolicyPoliciesOK creates a GetOrganizationAdaptivePolicyPoliciesOK with default headers values
func NewGetOrganizationAdaptivePolicyPoliciesOK() *GetOrganizationAdaptivePolicyPoliciesOK {
	return &GetOrganizationAdaptivePolicyPoliciesOK{}
}

/*
GetOrganizationAdaptivePolicyPoliciesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationAdaptivePolicyPoliciesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get organization adaptive policy policies o k response has a 2xx status code
func (o *GetOrganizationAdaptivePolicyPoliciesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization adaptive policy policies o k response has a 3xx status code
func (o *GetOrganizationAdaptivePolicyPoliciesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization adaptive policy policies o k response has a 4xx status code
func (o *GetOrganizationAdaptivePolicyPoliciesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization adaptive policy policies o k response has a 5xx status code
func (o *GetOrganizationAdaptivePolicyPoliciesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization adaptive policy policies o k response a status code equal to that given
func (o *GetOrganizationAdaptivePolicyPoliciesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization adaptive policy policies o k response
func (o *GetOrganizationAdaptivePolicyPoliciesOK) Code() int {
	return 200
}

func (o *GetOrganizationAdaptivePolicyPoliciesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/policies][%d] getOrganizationAdaptivePolicyPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicyPoliciesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/adaptivePolicy/policies][%d] getOrganizationAdaptivePolicyPoliciesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdaptivePolicyPoliciesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetOrganizationAdaptivePolicyPoliciesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
