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

// GetOrganizationClientsSearchReader is a Reader for the GetOrganizationClientsSearch structure.
type GetOrganizationClientsSearchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationClientsSearchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationClientsSearchOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationClientsSearchOK creates a GetOrganizationClientsSearchOK with default headers values
func NewGetOrganizationClientsSearchOK() *GetOrganizationClientsSearchOK {
	return &GetOrganizationClientsSearchOK{}
}

/* GetOrganizationClientsSearchOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationClientsSearchOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload interface{}
}

// IsSuccess returns true when this get organization clients search o k response has a 2xx status code
func (o *GetOrganizationClientsSearchOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization clients search o k response has a 3xx status code
func (o *GetOrganizationClientsSearchOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization clients search o k response has a 4xx status code
func (o *GetOrganizationClientsSearchOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization clients search o k response has a 5xx status code
func (o *GetOrganizationClientsSearchOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization clients search o k response a status code equal to that given
func (o *GetOrganizationClientsSearchOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationClientsSearchOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/clients/search][%d] getOrganizationClientsSearchOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationClientsSearchOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/clients/search][%d] getOrganizationClientsSearchOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationClientsSearchOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationClientsSearchOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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