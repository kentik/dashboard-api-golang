// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// CombineOrganizationNetworksReader is a Reader for the CombineOrganizationNetworks structure.
type CombineOrganizationNetworksReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CombineOrganizationNetworksReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCombineOrganizationNetworksOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /organizations/{organizationId}/networks/combine] combineOrganizationNetworks", response, response.Code())
	}
}

// NewCombineOrganizationNetworksOK creates a CombineOrganizationNetworksOK with default headers values
func NewCombineOrganizationNetworksOK() *CombineOrganizationNetworksOK {
	return &CombineOrganizationNetworksOK{}
}

/*
CombineOrganizationNetworksOK describes a response with status code 200, with default header values.

Successful operation
*/
type CombineOrganizationNetworksOK struct {
	Payload *CombineOrganizationNetworksOKBody
}

// IsSuccess returns true when this combine organization networks o k response has a 2xx status code
func (o *CombineOrganizationNetworksOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this combine organization networks o k response has a 3xx status code
func (o *CombineOrganizationNetworksOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this combine organization networks o k response has a 4xx status code
func (o *CombineOrganizationNetworksOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this combine organization networks o k response has a 5xx status code
func (o *CombineOrganizationNetworksOK) IsServerError() bool {
	return false
}

// IsCode returns true when this combine organization networks o k response a status code equal to that given
func (o *CombineOrganizationNetworksOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the combine organization networks o k response
func (o *CombineOrganizationNetworksOK) Code() int {
	return 200
}

func (o *CombineOrganizationNetworksOK) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/networks/combine][%d] combineOrganizationNetworksOK  %+v", 200, o.Payload)
}

func (o *CombineOrganizationNetworksOK) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/networks/combine][%d] combineOrganizationNetworksOK  %+v", 200, o.Payload)
}

func (o *CombineOrganizationNetworksOK) GetPayload() *CombineOrganizationNetworksOKBody {
	return o.Payload
}

func (o *CombineOrganizationNetworksOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CombineOrganizationNetworksOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CombineOrganizationNetworksBody combine organization networks body
// Example: {"enrollmentString":"my-enrollment-string","name":"Long Island Office","networkIds":["N_1234","N_5678"]}
swagger:model CombineOrganizationNetworksBody
*/
type CombineOrganizationNetworksBody struct {

	// A unique identifier which can be used for device enrollment or easy access through the Meraki SM Registration page or the Self Service Portal. Please note that changing this field may cause existing bookmarks to break. All networks that are part of this combined network will have their enrollment string appended by '-network_type'. If left empty, all exisitng enrollment strings will be deleted.
	EnrollmentString string `json:"enrollmentString,omitempty"`

	// The name of the combined network
	// Required: true
	Name *string `json:"name"`

	// A list of the network IDs that will be combined. If an ID of a combined network is included in this list, the other networks in the list will be grouped into that network
	// Required: true
	NetworkIds []string `json:"networkIds"`
}

// Validate validates this combine organization networks body
func (o *CombineOrganizationNetworksBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworkIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CombineOrganizationNetworksBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("combineOrganizationNetworks"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CombineOrganizationNetworksBody) validateNetworkIds(formats strfmt.Registry) error {

	if err := validate.Required("combineOrganizationNetworks"+"."+"networkIds", "body", o.NetworkIds); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this combine organization networks body based on context it is used
func (o *CombineOrganizationNetworksBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CombineOrganizationNetworksBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CombineOrganizationNetworksBody) UnmarshalBinary(b []byte) error {
	var res CombineOrganizationNetworksBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CombineOrganizationNetworksOKBody combine organization networks o k body
swagger:model CombineOrganizationNetworksOKBody
*/
type CombineOrganizationNetworksOKBody struct {

	// resulting network
	ResultingNetwork *CombineOrganizationNetworksOKBodyResultingNetwork `json:"resultingNetwork,omitempty"`
}

// Validate validates this combine organization networks o k body
func (o *CombineOrganizationNetworksOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateResultingNetwork(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CombineOrganizationNetworksOKBody) validateResultingNetwork(formats strfmt.Registry) error {
	if swag.IsZero(o.ResultingNetwork) { // not required
		return nil
	}

	if o.ResultingNetwork != nil {
		if err := o.ResultingNetwork.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("combineOrganizationNetworksOK" + "." + "resultingNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("combineOrganizationNetworksOK" + "." + "resultingNetwork")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this combine organization networks o k body based on the context it is used
func (o *CombineOrganizationNetworksOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateResultingNetwork(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CombineOrganizationNetworksOKBody) contextValidateResultingNetwork(ctx context.Context, formats strfmt.Registry) error {

	if o.ResultingNetwork != nil {

		if swag.IsZero(o.ResultingNetwork) { // not required
			return nil
		}

		if err := o.ResultingNetwork.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("combineOrganizationNetworksOK" + "." + "resultingNetwork")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("combineOrganizationNetworksOK" + "." + "resultingNetwork")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CombineOrganizationNetworksOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CombineOrganizationNetworksOKBody) UnmarshalBinary(b []byte) error {
	var res CombineOrganizationNetworksOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CombineOrganizationNetworksOKBodyResultingNetwork Network after the combination
swagger:model CombineOrganizationNetworksOKBodyResultingNetwork
*/
type CombineOrganizationNetworksOKBodyResultingNetwork struct {

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

// Validate validates this combine organization networks o k body resulting network
func (o *CombineOrganizationNetworksOKBodyResultingNetwork) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this combine organization networks o k body resulting network based on context it is used
func (o *CombineOrganizationNetworksOKBodyResultingNetwork) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CombineOrganizationNetworksOKBodyResultingNetwork) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CombineOrganizationNetworksOKBodyResultingNetwork) UnmarshalBinary(b []byte) error {
	var res CombineOrganizationNetworksOKBodyResultingNetwork
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
