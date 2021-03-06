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

// UpdateNetworkReader is a Reader for the UpdateNetwork structure.
type UpdateNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkOK creates a UpdateNetworkOK with default headers values
func NewUpdateNetworkOK() *UpdateNetworkOK {
	return &UpdateNetworkOK{}
}

/* UpdateNetworkOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network o k response has a 2xx status code
func (o *UpdateNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network o k response has a 3xx status code
func (o *UpdateNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network o k response has a 4xx status code
func (o *UpdateNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network o k response has a 5xx status code
func (o *UpdateNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network o k response a status code equal to that given
func (o *UpdateNetworkOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}][%d] updateNetworkOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}][%d] updateNetworkOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkBody update network body
// Example: {"name":"Long Island Office","notes":"Combined network for Long Island Office","tags":["tag1","tag2"],"timeZone":"America/Los_Angeles"}
swagger:model UpdateNetworkBody
*/
type UpdateNetworkBody struct {

	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break.
	EnrollmentString string `json:"enrollmentString,omitempty"`

	// The name of the network
	Name string `json:"name,omitempty"`

	// Add any notes or additional information about this network here.
	Notes string `json:"notes,omitempty"`

	// A list of tags to be applied to the network
	Tags []string `json:"tags"`

	// The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this update network body
func (o *UpdateNetworkBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network body based on context it is used
func (o *UpdateNetworkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
