// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// GetNetworkReader is a Reader for the GetNetwork structure.
type GetNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}] getNetwork", response, response.Code())
	}
}

// NewGetNetworkOK creates a GetNetworkOK with default headers values
func NewGetNetworkOK() *GetNetworkOK {
	return &GetNetworkOK{}
}

/*
GetNetworkOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkOK struct {
	Payload *GetNetworkOKBody
}

// IsSuccess returns true when this get network o k response has a 2xx status code
func (o *GetNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network o k response has a 3xx status code
func (o *GetNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network o k response has a 4xx status code
func (o *GetNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network o k response has a 5xx status code
func (o *GetNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network o k response a status code equal to that given
func (o *GetNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network o k response
func (o *GetNetworkOK) Code() int {
	return 200
}

func (o *GetNetworkOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}][%d] getNetworkOK  %+v", 200, o.Payload)
}

func (o *GetNetworkOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}][%d] getNetworkOK  %+v", 200, o.Payload)
}

func (o *GetNetworkOK) GetPayload() *GetNetworkOKBody {
	return o.Payload
}

func (o *GetNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkOKBody get network o k body
swagger:model GetNetworkOKBody
*/
type GetNetworkOKBody struct {

	// Enrollment string for the network
	EnrollmentString string `json:"enrollmentString,omitempty"`

	// Network ID
	ID string `json:"id,omitempty"`

	// If the network is bound to a config template
	IsBoundToConfigTemplate bool `json:"isBoundToConfigTemplate,omitempty"`

	// Network name
	Name string `json:"name,omitempty"`

	// Notes for the network
	Notes string `json:"notes,omitempty"`

	// Organization ID
	OrganizationID string `json:"organizationId,omitempty"`

	// List of the product types that the network supports
	ProductTypes []string `json:"productTypes"`

	// Network tags
	Tags []string `json:"tags"`

	// Timezone of the network
	TimeZone string `json:"timeZone,omitempty"`

	// URL to the network Dashboard UI
	URL string `json:"url,omitempty"`
}

// Validate validates this get network o k body
func (o *GetNetworkOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network o k body based on context it is used
func (o *GetNetworkOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
