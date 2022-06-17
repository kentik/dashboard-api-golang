// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateOrganizationNetworkReader is a Reader for the CreateOrganizationNetwork structure.
type CreateOrganizationNetworkReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationNetworkReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationNetworkCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateOrganizationNetworkCreated creates a CreateOrganizationNetworkCreated with default headers values
func NewCreateOrganizationNetworkCreated() *CreateOrganizationNetworkCreated {
	return &CreateOrganizationNetworkCreated{}
}

/* CreateOrganizationNetworkCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationNetworkCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create organization network created response has a 2xx status code
func (o *CreateOrganizationNetworkCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization network created response has a 3xx status code
func (o *CreateOrganizationNetworkCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization network created response has a 4xx status code
func (o *CreateOrganizationNetworkCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization network created response has a 5xx status code
func (o *CreateOrganizationNetworkCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization network created response a status code equal to that given
func (o *CreateOrganizationNetworkCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateOrganizationNetworkCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/networks][%d] createOrganizationNetworkCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationNetworkCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/networks][%d] createOrganizationNetworkCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationNetworkCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateOrganizationNetworkCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateOrganizationNetworkBody create organization network body
// Example: {"name":"Long Island Office","notes":"Combined network for Long Island Office","productTypes":["appliance","switch","camera"],"tags":["tag1","tag2"],"timeZone":"America/Los_Angeles"}
swagger:model CreateOrganizationNetworkBody
*/
type CreateOrganizationNetworkBody struct {

	// The ID of the network to copy configuration from. Other provided parameters will override the copied configuration, except type which must match this network's type exactly.
	CopyFromNetworkID string `json:"copyFromNetworkId,omitempty"`

	// The name of the new network
	// Required: true
	Name *string `json:"name"`

	// Add any notes or additional information about this network here.
	Notes string `json:"notes,omitempty"`

	// The product type(s) of the new network. Valid types are wireless, appliance, switch, systemsManager, camera, cellularGateway, sensor, environmental. If more than one type is included, the network will be a combined network.
	// Required: true
	ProductTypes []string `json:"productTypes"`

	// A list of tags to be applied to the network
	Tags []string `json:"tags"`

	// The timezone of the network. For a list of allowed timezones, please see the 'TZ' column in the table in <a target='_blank' href='https://en.wikipedia.org/wiki/List_of_tz_database_time_zones'>this article.</a>
	TimeZone string `json:"timeZone,omitempty"`
}

// Validate validates this create organization network body
func (o *CreateOrganizationNetworkBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateProductTypes(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationNetworkBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationNetwork"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

var createOrganizationNetworkBodyProductTypesItemsEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["wireless","appliance","switch","systemsManager","camera","cellularGateway","sensor","environmental"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createOrganizationNetworkBodyProductTypesItemsEnum = append(createOrganizationNetworkBodyProductTypesItemsEnum, v)
	}
}

func (o *CreateOrganizationNetworkBody) validateProductTypesItemsEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createOrganizationNetworkBodyProductTypesItemsEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateOrganizationNetworkBody) validateProductTypes(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationNetwork"+"."+"productTypes", "body", o.ProductTypes); err != nil {
		return err
	}

	for i := 0; i < len(o.ProductTypes); i++ {

		// value enum
		if err := o.validateProductTypesItemsEnum("createOrganizationNetwork"+"."+"productTypes"+"."+strconv.Itoa(i), "body", o.ProductTypes[i]); err != nil {
			return err
		}

	}

	return nil
}

// ContextValidate validates this create organization network body based on context it is used
func (o *CreateOrganizationNetworkBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationNetworkBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationNetworkBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationNetworkBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}