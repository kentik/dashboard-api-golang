// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CloneOrganizationReader is a Reader for the CloneOrganization structure.
type CloneOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CloneOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCloneOrganizationCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCloneOrganizationCreated creates a CloneOrganizationCreated with default headers values
func NewCloneOrganizationCreated() *CloneOrganizationCreated {
	return &CloneOrganizationCreated{}
}

/* CloneOrganizationCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CloneOrganizationCreated struct {
	Payload *CloneOrganizationCreatedBody
}

// IsSuccess returns true when this clone organization created response has a 2xx status code
func (o *CloneOrganizationCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this clone organization created response has a 3xx status code
func (o *CloneOrganizationCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this clone organization created response has a 4xx status code
func (o *CloneOrganizationCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this clone organization created response has a 5xx status code
func (o *CloneOrganizationCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this clone organization created response a status code equal to that given
func (o *CloneOrganizationCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CloneOrganizationCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/clone][%d] cloneOrganizationCreated  %+v", 201, o.Payload)
}

func (o *CloneOrganizationCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/clone][%d] cloneOrganizationCreated  %+v", 201, o.Payload)
}

func (o *CloneOrganizationCreated) GetPayload() *CloneOrganizationCreatedBody {
	return o.Payload
}

func (o *CloneOrganizationCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CloneOrganizationCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CloneOrganizationBody clone organization body
// Example: {"name":"My organization"}
swagger:model CloneOrganizationBody
*/
type CloneOrganizationBody struct {

	// The name of the new organization
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this clone organization body
func (o *CloneOrganizationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CloneOrganizationBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("cloneOrganization"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this clone organization body based on context it is used
func (o *CloneOrganizationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationBody) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CloneOrganizationCreatedBody clone organization created body
swagger:model CloneOrganizationCreatedBody
*/
type CloneOrganizationCreatedBody struct {

	// api
	API *CloneOrganizationCreatedBodyAPI `json:"api,omitempty"`

	// cloud
	Cloud *CloneOrganizationCreatedBodyCloud `json:"cloud,omitempty"`

	// Organization ID
	ID string `json:"id,omitempty"`

	// licensing
	Licensing *CloneOrganizationCreatedBodyLicensing `json:"licensing,omitempty"`

	// Organization name
	Name string `json:"name,omitempty"`

	// Organization URL
	URL string `json:"url,omitempty"`
}

// Validate validates this clone organization created body
func (o *CloneOrganizationCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCloud(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLicensing(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CloneOrganizationCreatedBody) validateAPI(formats strfmt.Registry) error {
	if swag.IsZero(o.API) { // not required
		return nil
	}

	if o.API != nil {
		if err := o.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "api")
			}
			return err
		}
	}

	return nil
}

func (o *CloneOrganizationCreatedBody) validateCloud(formats strfmt.Registry) error {
	if swag.IsZero(o.Cloud) { // not required
		return nil
	}

	if o.Cloud != nil {
		if err := o.Cloud.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "cloud")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "cloud")
			}
			return err
		}
	}

	return nil
}

func (o *CloneOrganizationCreatedBody) validateLicensing(formats strfmt.Registry) error {
	if swag.IsZero(o.Licensing) { // not required
		return nil
	}

	if o.Licensing != nil {
		if err := o.Licensing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "licensing")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clone organization created body based on the context it is used
func (o *CloneOrganizationCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateCloud(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateLicensing(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CloneOrganizationCreatedBody) contextValidateAPI(ctx context.Context, formats strfmt.Registry) error {

	if o.API != nil {
		if err := o.API.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "api")
			}
			return err
		}
	}

	return nil
}

func (o *CloneOrganizationCreatedBody) contextValidateCloud(ctx context.Context, formats strfmt.Registry) error {

	if o.Cloud != nil {
		if err := o.Cloud.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "cloud")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "cloud")
			}
			return err
		}
	}

	return nil
}

func (o *CloneOrganizationCreatedBody) contextValidateLicensing(ctx context.Context, formats strfmt.Registry) error {

	if o.Licensing != nil {
		if err := o.Licensing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "licensing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationCreatedBody) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CloneOrganizationCreatedBodyAPI API related settings
swagger:model CloneOrganizationCreatedBodyAPI
*/
type CloneOrganizationCreatedBodyAPI struct {

	// Enable API access
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this clone organization created body API
func (o *CloneOrganizationCreatedBodyAPI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clone organization created body API based on context it is used
func (o *CloneOrganizationCreatedBodyAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyAPI) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyAPI) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationCreatedBodyAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CloneOrganizationCreatedBodyCloud Data for this organization
swagger:model CloneOrganizationCreatedBodyCloud
*/
type CloneOrganizationCreatedBodyCloud struct {

	// region
	Region *CloneOrganizationCreatedBodyCloudRegion `json:"region,omitempty"`
}

// Validate validates this clone organization created body cloud
func (o *CloneOrganizationCreatedBodyCloud) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CloneOrganizationCreatedBodyCloud) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "cloud" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "cloud" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this clone organization created body cloud based on the context it is used
func (o *CloneOrganizationCreatedBodyCloud) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CloneOrganizationCreatedBodyCloud) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {
		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cloneOrganizationCreated" + "." + "cloud" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("cloneOrganizationCreated" + "." + "cloud" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyCloud) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyCloud) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationCreatedBodyCloud
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CloneOrganizationCreatedBodyCloudRegion Region info
swagger:model CloneOrganizationCreatedBodyCloudRegion
*/
type CloneOrganizationCreatedBodyCloudRegion struct {

	// Name of region
	Name string `json:"name,omitempty"`
}

// Validate validates this clone organization created body cloud region
func (o *CloneOrganizationCreatedBodyCloudRegion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this clone organization created body cloud region based on context it is used
func (o *CloneOrganizationCreatedBodyCloudRegion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyCloudRegion) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyCloudRegion) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationCreatedBodyCloudRegion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CloneOrganizationCreatedBodyLicensing Licensing related settings
swagger:model CloneOrganizationCreatedBodyLicensing
*/
type CloneOrganizationCreatedBodyLicensing struct {

	// Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
	// Enum: [co-term per-device subscription]
	Model string `json:"model,omitempty"`
}

// Validate validates this clone organization created body licensing
func (o *CloneOrganizationCreatedBodyLicensing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var cloneOrganizationCreatedBodyLicensingTypeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["co-term","per-device","subscription"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		cloneOrganizationCreatedBodyLicensingTypeModelPropEnum = append(cloneOrganizationCreatedBodyLicensingTypeModelPropEnum, v)
	}
}

const (

	// CloneOrganizationCreatedBodyLicensingModelCoDashTerm captures enum value "co-term"
	CloneOrganizationCreatedBodyLicensingModelCoDashTerm string = "co-term"

	// CloneOrganizationCreatedBodyLicensingModelPerDashDevice captures enum value "per-device"
	CloneOrganizationCreatedBodyLicensingModelPerDashDevice string = "per-device"

	// CloneOrganizationCreatedBodyLicensingModelSubscription captures enum value "subscription"
	CloneOrganizationCreatedBodyLicensingModelSubscription string = "subscription"
)

// prop value enum
func (o *CloneOrganizationCreatedBodyLicensing) validateModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, cloneOrganizationCreatedBodyLicensingTypeModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CloneOrganizationCreatedBodyLicensing) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(o.Model) { // not required
		return nil
	}

	// value enum
	if err := o.validateModelEnum("cloneOrganizationCreated"+"."+"licensing"+"."+"model", "body", o.Model); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this clone organization created body licensing based on context it is used
func (o *CloneOrganizationCreatedBodyLicensing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyLicensing) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CloneOrganizationCreatedBodyLicensing) UnmarshalBinary(b []byte) error {
	var res CloneOrganizationCreatedBodyLicensing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}