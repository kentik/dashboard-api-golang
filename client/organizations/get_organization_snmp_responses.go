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

// GetOrganizationSnmpReader is a Reader for the GetOrganizationSnmp structure.
type GetOrganizationSnmpReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSnmpReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSnmpOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/snmp] getOrganizationSnmp", response, response.Code())
	}
}

// NewGetOrganizationSnmpOK creates a GetOrganizationSnmpOK with default headers values
func NewGetOrganizationSnmpOK() *GetOrganizationSnmpOK {
	return &GetOrganizationSnmpOK{}
}

/*
GetOrganizationSnmpOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSnmpOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get organization snmp o k response has a 2xx status code
func (o *GetOrganizationSnmpOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization snmp o k response has a 3xx status code
func (o *GetOrganizationSnmpOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization snmp o k response has a 4xx status code
func (o *GetOrganizationSnmpOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization snmp o k response has a 5xx status code
func (o *GetOrganizationSnmpOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization snmp o k response a status code equal to that given
func (o *GetOrganizationSnmpOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization snmp o k response
func (o *GetOrganizationSnmpOK) Code() int {
	return 200
}

func (o *GetOrganizationSnmpOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/snmp][%d] getOrganizationSnmpOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSnmpOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/snmp][%d] getOrganizationSnmpOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSnmpOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetOrganizationSnmpOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
