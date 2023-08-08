// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// SplitNetworkReader is a Reader for the SplitNetwork structure.
type SplitNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *SplitNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewSplitNetworkOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/split] splitNetwork", response, response.Code())
	}
}

// NewSplitNetworkOK creates a SplitNetworkOK with default headers values
func NewSplitNetworkOK() *SplitNetworkOK {
	return &SplitNetworkOK{}
}

/*
SplitNetworkOK describes a response with status code 200, with default header values.

Successful operation
*/
type SplitNetworkOK struct {
	Payload *SplitNetworkOKBody
}

// IsSuccess returns true when this split network o k response has a 2xx status code
func (o *SplitNetworkOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this split network o k response has a 3xx status code
func (o *SplitNetworkOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this split network o k response has a 4xx status code
func (o *SplitNetworkOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this split network o k response has a 5xx status code
func (o *SplitNetworkOK) IsServerError() bool {
	return false
}

// IsCode returns true when this split network o k response a status code equal to that given
func (o *SplitNetworkOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the split network o k response
func (o *SplitNetworkOK) Code() int {
	return 200
}

func (o *SplitNetworkOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/split][%d] splitNetworkOK  %+v", 200, o.Payload)
}

func (o *SplitNetworkOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/split][%d] splitNetworkOK  %+v", 200, o.Payload)
}

func (o *SplitNetworkOK) GetPayload() *SplitNetworkOKBody {
	return o.Payload
}

func (o *SplitNetworkOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(SplitNetworkOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
SplitNetworkOKBody split network o k body
swagger:model SplitNetworkOKBody
*/
type SplitNetworkOKBody struct {

	// Networks after the split
	ResultingNetworks []*SplitNetworkOKBodyResultingNetworksItems0 `json:"resultingNetworks"`
}

// Validate validates this split network o k body
func (o *SplitNetworkOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResultingNetworks(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SplitNetworkOKBody) validateResultingNetworks(formats strfmt.Registry) error {
	if swag.IsZero(o.ResultingNetworks) { // not required
		return nil
	}

	for i := 0; i < len(o.ResultingNetworks); i++ {
		if swag.IsZero(o.ResultingNetworks[i]) { // not required
			continue
		}

		if o.ResultingNetworks[i] != nil {
			if err := o.ResultingNetworks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("splitNetworkOK" + "." + "resultingNetworks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("splitNetworkOK" + "." + "resultingNetworks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this split network o k body based on the context it is used
func (o *SplitNetworkOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResultingNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *SplitNetworkOKBody) contextValidateResultingNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.ResultingNetworks); i++ {

		if o.ResultingNetworks[i] != nil {

			if swag.IsZero(o.ResultingNetworks[i]) { // not required
				return nil
			}

			if err := o.ResultingNetworks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("splitNetworkOK" + "." + "resultingNetworks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("splitNetworkOK" + "." + "resultingNetworks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *SplitNetworkOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SplitNetworkOKBody) UnmarshalBinary(b []byte) error {
	var res SplitNetworkOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
SplitNetworkOKBodyResultingNetworksItems0 split network o k body resulting networks items0
swagger:model SplitNetworkOKBodyResultingNetworksItems0
*/
type SplitNetworkOKBodyResultingNetworksItems0 struct {

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

// Validate validates this split network o k body resulting networks items0
func (o *SplitNetworkOKBodyResultingNetworksItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this split network o k body resulting networks items0 based on context it is used
func (o *SplitNetworkOKBodyResultingNetworksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *SplitNetworkOKBodyResultingNetworksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *SplitNetworkOKBodyResultingNetworksItems0) UnmarshalBinary(b []byte) error {
	var res SplitNetworkOKBodyResultingNetworksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
