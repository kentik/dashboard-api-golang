// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetOrganizationSmApnsCertReader is a Reader for the GetOrganizationSmApnsCert structure.
type GetOrganizationSmApnsCertReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSmApnsCertReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSmApnsCertOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/sm/apnsCert] getOrganizationSmApnsCert", response, response.Code())
	}
}

// NewGetOrganizationSmApnsCertOK creates a GetOrganizationSmApnsCertOK with default headers values
func NewGetOrganizationSmApnsCertOK() *GetOrganizationSmApnsCertOK {
	return &GetOrganizationSmApnsCertOK{}
}

/*
GetOrganizationSmApnsCertOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSmApnsCertOK struct {
	Payload *GetOrganizationSmApnsCertOKBody
}

// IsSuccess returns true when this get organization sm apns cert o k response has a 2xx status code
func (o *GetOrganizationSmApnsCertOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization sm apns cert o k response has a 3xx status code
func (o *GetOrganizationSmApnsCertOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization sm apns cert o k response has a 4xx status code
func (o *GetOrganizationSmApnsCertOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization sm apns cert o k response has a 5xx status code
func (o *GetOrganizationSmApnsCertOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization sm apns cert o k response a status code equal to that given
func (o *GetOrganizationSmApnsCertOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization sm apns cert o k response
func (o *GetOrganizationSmApnsCertOK) Code() int {
	return 200
}

func (o *GetOrganizationSmApnsCertOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/sm/apnsCert][%d] getOrganizationSmApnsCertOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSmApnsCertOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/sm/apnsCert][%d] getOrganizationSmApnsCertOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSmApnsCertOK) GetPayload() *GetOrganizationSmApnsCertOKBody {
	return o.Payload
}

func (o *GetOrganizationSmApnsCertOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationSmApnsCertOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationSmApnsCertOKBody get organization sm apns cert o k body
swagger:model GetOrganizationSmApnsCertOKBody
*/
type GetOrganizationSmApnsCertOKBody struct {

	// Organization APNS Certificate used by devices to communication with Apple
	Certificate string `json:"certificate,omitempty"`
}

// Validate validates this get organization sm apns cert o k body
func (o *GetOrganizationSmApnsCertOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization sm apns cert o k body based on context it is used
func (o *GetOrganizationSmApnsCertOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSmApnsCertOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSmApnsCertOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSmApnsCertOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
