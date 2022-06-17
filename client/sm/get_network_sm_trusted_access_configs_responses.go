// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetNetworkSmTrustedAccessConfigsReader is a Reader for the GetNetworkSmTrustedAccessConfigs structure.
type GetNetworkSmTrustedAccessConfigsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmTrustedAccessConfigsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmTrustedAccessConfigsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmTrustedAccessConfigsOK creates a GetNetworkSmTrustedAccessConfigsOK with default headers values
func NewGetNetworkSmTrustedAccessConfigsOK() *GetNetworkSmTrustedAccessConfigsOK {
	return &GetNetworkSmTrustedAccessConfigsOK{}
}

/* GetNetworkSmTrustedAccessConfigsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmTrustedAccessConfigsOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetNetworkSmTrustedAccessConfigsOKBodyItems0
}

// IsSuccess returns true when this get network sm trusted access configs o k response has a 2xx status code
func (o *GetNetworkSmTrustedAccessConfigsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm trusted access configs o k response has a 3xx status code
func (o *GetNetworkSmTrustedAccessConfigsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm trusted access configs o k response has a 4xx status code
func (o *GetNetworkSmTrustedAccessConfigsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm trusted access configs o k response has a 5xx status code
func (o *GetNetworkSmTrustedAccessConfigsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm trusted access configs o k response a status code equal to that given
func (o *GetNetworkSmTrustedAccessConfigsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkSmTrustedAccessConfigsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/trustedAccessConfigs][%d] getNetworkSmTrustedAccessConfigsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmTrustedAccessConfigsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/trustedAccessConfigs][%d] getNetworkSmTrustedAccessConfigsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmTrustedAccessConfigsOK) GetPayload() []*GetNetworkSmTrustedAccessConfigsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSmTrustedAccessConfigsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetNetworkSmTrustedAccessConfigsOKBodyItems0 get network sm trusted access configs o k body items0
swagger:model GetNetworkSmTrustedAccessConfigsOKBodyItems0
*/
type GetNetworkSmTrustedAccessConfigsOKBodyItems0 struct {

	// time that access ends
	// Format: date-time
	AccessEndAt strfmt.DateTime `json:"accessEndAt,omitempty"`

	// time that access starts
	// Format: date-time
	AccessStartAt strfmt.DateTime `json:"accessStartAt,omitempty"`

	// device ID
	ID string `json:"id,omitempty"`

	// device name
	Name string `json:"name,omitempty"`

	// scope
	Scope string `json:"scope,omitempty"`

	// SSID name
	SsidName string `json:"ssidName,omitempty"`

	// device tags
	Tags []string `json:"tags"`
}

// Validate validates this get network sm trusted access configs o k body items0
func (o *GetNetworkSmTrustedAccessConfigsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessEndAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAccessStartAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSmTrustedAccessConfigsOKBodyItems0) validateAccessEndAt(formats strfmt.Registry) error {
	if swag.IsZero(o.AccessEndAt) { // not required
		return nil
	}

	if err := validate.FormatOf("accessEndAt", "body", "date-time", o.AccessEndAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkSmTrustedAccessConfigsOKBodyItems0) validateAccessStartAt(formats strfmt.Registry) error {
	if swag.IsZero(o.AccessStartAt) { // not required
		return nil
	}

	if err := validate.FormatOf("accessStartAt", "body", "date-time", o.AccessStartAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network sm trusted access configs o k body items0 based on context it is used
func (o *GetNetworkSmTrustedAccessConfigsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmTrustedAccessConfigsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmTrustedAccessConfigsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmTrustedAccessConfigsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
