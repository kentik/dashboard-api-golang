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

// GetOrganizationSamlReader is a Reader for the GetOrganizationSaml structure.
type GetOrganizationSamlReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSamlReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSamlOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationSamlOK creates a GetOrganizationSamlOK with default headers values
func NewGetOrganizationSamlOK() *GetOrganizationSamlOK {
	return &GetOrganizationSamlOK{}
}

/* GetOrganizationSamlOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSamlOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization saml o k response has a 2xx status code
func (o *GetOrganizationSamlOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization saml o k response has a 3xx status code
func (o *GetOrganizationSamlOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization saml o k response has a 4xx status code
func (o *GetOrganizationSamlOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization saml o k response has a 5xx status code
func (o *GetOrganizationSamlOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization saml o k response a status code equal to that given
func (o *GetOrganizationSamlOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationSamlOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/saml][%d] getOrganizationSamlOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSamlOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/saml][%d] getOrganizationSamlOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSamlOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSamlOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
