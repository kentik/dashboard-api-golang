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

// GetOrganizationOpenapiSpecReader is a Reader for the GetOrganizationOpenapiSpec structure.
type GetOrganizationOpenapiSpecReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationOpenapiSpecReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationOpenapiSpecOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationOpenapiSpecOK creates a GetOrganizationOpenapiSpecOK with default headers values
func NewGetOrganizationOpenapiSpecOK() *GetOrganizationOpenapiSpecOK {
	return &GetOrganizationOpenapiSpecOK{}
}

/* GetOrganizationOpenapiSpecOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationOpenapiSpecOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization openapi spec o k response has a 2xx status code
func (o *GetOrganizationOpenapiSpecOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization openapi spec o k response has a 3xx status code
func (o *GetOrganizationOpenapiSpecOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization openapi spec o k response has a 4xx status code
func (o *GetOrganizationOpenapiSpecOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization openapi spec o k response has a 5xx status code
func (o *GetOrganizationOpenapiSpecOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization openapi spec o k response a status code equal to that given
func (o *GetOrganizationOpenapiSpecOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationOpenapiSpecOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/openapiSpec][%d] getOrganizationOpenapiSpecOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationOpenapiSpecOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/openapiSpec][%d] getOrganizationOpenapiSpecOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationOpenapiSpecOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationOpenapiSpecOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
