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

// UpdateOrganizationReader is a Reader for the UpdateOrganization structure.
type UpdateOrganizationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationOK creates a UpdateOrganizationOK with default headers values
func NewUpdateOrganizationOK() *UpdateOrganizationOK {
	return &UpdateOrganizationOK{}
}

/* UpdateOrganizationOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationOK struct {
	Payload *UpdateOrganizationOKBody
}

// IsSuccess returns true when this update organization o k response has a 2xx status code
func (o *UpdateOrganizationOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization o k response has a 3xx status code
func (o *UpdateOrganizationOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization o k response has a 4xx status code
func (o *UpdateOrganizationOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization o k response has a 5xx status code
func (o *UpdateOrganizationOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization o k response a status code equal to that given
func (o *UpdateOrganizationOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateOrganizationOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}][%d] updateOrganizationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}][%d] updateOrganizationOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationOK) GetPayload() *UpdateOrganizationOKBody {
	return o.Payload
}

func (o *UpdateOrganizationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateOrganizationOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateOrganizationBody update organization body
// Example: {"api":{"enabled":true},"name":"My organization"}
swagger:model UpdateOrganizationBody
*/
type UpdateOrganizationBody struct {

	// api
	API *UpdateOrganizationParamsBodyAPI `json:"api,omitempty"`

	// The name of the organization
	Name string `json:"name,omitempty"`
}

// Validate validates this update organization body
func (o *UpdateOrganizationBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAPI(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationBody) validateAPI(formats strfmt.Registry) error {
	if swag.IsZero(o.API) { // not required
		return nil
	}

	if o.API != nil {
		if err := o.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganization" + "." + "api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganization" + "." + "api")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization body based on the context it is used
func (o *UpdateOrganizationBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAPI(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationBody) contextValidateAPI(ctx context.Context, formats strfmt.Registry) error {

	if o.API != nil {
		if err := o.API.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganization" + "." + "api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganization" + "." + "api")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationOKBody update organization o k body
swagger:model UpdateOrganizationOKBody
*/
type UpdateOrganizationOKBody struct {

	// api
	API *UpdateOrganizationOKBodyAPI `json:"api,omitempty"`

	// cloud
	Cloud *UpdateOrganizationOKBodyCloud `json:"cloud,omitempty"`

	// Organization ID
	ID string `json:"id,omitempty"`

	// licensing
	Licensing *UpdateOrganizationOKBodyLicensing `json:"licensing,omitempty"`

	// Organization name
	Name string `json:"name,omitempty"`

	// Organization URL
	URL string `json:"url,omitempty"`
}

// Validate validates this update organization o k body
func (o *UpdateOrganizationOKBody) Validate(formats strfmt.Registry) error {
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

func (o *UpdateOrganizationOKBody) validateAPI(formats strfmt.Registry) error {
	if swag.IsZero(o.API) { // not required
		return nil
	}

	if o.API != nil {
		if err := o.API.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "api")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateOrganizationOKBody) validateCloud(formats strfmt.Registry) error {
	if swag.IsZero(o.Cloud) { // not required
		return nil
	}

	if o.Cloud != nil {
		if err := o.Cloud.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "cloud")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "cloud")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateOrganizationOKBody) validateLicensing(formats strfmt.Registry) error {
	if swag.IsZero(o.Licensing) { // not required
		return nil
	}

	if o.Licensing != nil {
		if err := o.Licensing.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "licensing")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization o k body based on the context it is used
func (o *UpdateOrganizationOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

func (o *UpdateOrganizationOKBody) contextValidateAPI(ctx context.Context, formats strfmt.Registry) error {

	if o.API != nil {
		if err := o.API.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "api")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "api")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateOrganizationOKBody) contextValidateCloud(ctx context.Context, formats strfmt.Registry) error {

	if o.Cloud != nil {
		if err := o.Cloud.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "cloud")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "cloud")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateOrganizationOKBody) contextValidateLicensing(ctx context.Context, formats strfmt.Registry) error {

	if o.Licensing != nil {
		if err := o.Licensing.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "licensing")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "licensing")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationOKBodyAPI API related settings
swagger:model UpdateOrganizationOKBodyAPI
*/
type UpdateOrganizationOKBodyAPI struct {

	// Enable API access
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update organization o k body API
func (o *UpdateOrganizationOKBodyAPI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization o k body API based on context it is used
func (o *UpdateOrganizationOKBodyAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyAPI) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyAPI) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationOKBodyAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationOKBodyCloud Data for this organization
swagger:model UpdateOrganizationOKBodyCloud
*/
type UpdateOrganizationOKBodyCloud struct {

	// region
	Region *UpdateOrganizationOKBodyCloudRegion `json:"region,omitempty"`
}

// Validate validates this update organization o k body cloud
func (o *UpdateOrganizationOKBodyCloud) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateRegion(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationOKBodyCloud) validateRegion(formats strfmt.Registry) error {
	if swag.IsZero(o.Region) { // not required
		return nil
	}

	if o.Region != nil {
		if err := o.Region.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "cloud" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "cloud" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization o k body cloud based on the context it is used
func (o *UpdateOrganizationOKBodyCloud) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateRegion(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationOKBodyCloud) contextValidateRegion(ctx context.Context, formats strfmt.Registry) error {

	if o.Region != nil {
		if err := o.Region.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationOK" + "." + "cloud" + "." + "region")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationOK" + "." + "cloud" + "." + "region")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyCloud) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyCloud) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationOKBodyCloud
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationOKBodyCloudRegion Region info
swagger:model UpdateOrganizationOKBodyCloudRegion
*/
type UpdateOrganizationOKBodyCloudRegion struct {

	// Name of region
	Name string `json:"name,omitempty"`
}

// Validate validates this update organization o k body cloud region
func (o *UpdateOrganizationOKBodyCloudRegion) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization o k body cloud region based on context it is used
func (o *UpdateOrganizationOKBodyCloudRegion) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyCloudRegion) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyCloudRegion) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationOKBodyCloudRegion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationOKBodyLicensing Licensing related settings
swagger:model UpdateOrganizationOKBodyLicensing
*/
type UpdateOrganizationOKBodyLicensing struct {

	// Organization licensing model. Can be 'co-term', 'per-device', or 'subscription'.
	// Enum: [co-term per-device subscription]
	Model string `json:"model,omitempty"`
}

// Validate validates this update organization o k body licensing
func (o *UpdateOrganizationOKBodyLicensing) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateModel(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateOrganizationOKBodyLicensingTypeModelPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["co-term","per-device","subscription"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateOrganizationOKBodyLicensingTypeModelPropEnum = append(updateOrganizationOKBodyLicensingTypeModelPropEnum, v)
	}
}

const (

	// UpdateOrganizationOKBodyLicensingModelCoDashTerm captures enum value "co-term"
	UpdateOrganizationOKBodyLicensingModelCoDashTerm string = "co-term"

	// UpdateOrganizationOKBodyLicensingModelPerDashDevice captures enum value "per-device"
	UpdateOrganizationOKBodyLicensingModelPerDashDevice string = "per-device"

	// UpdateOrganizationOKBodyLicensingModelSubscription captures enum value "subscription"
	UpdateOrganizationOKBodyLicensingModelSubscription string = "subscription"
)

// prop value enum
func (o *UpdateOrganizationOKBodyLicensing) validateModelEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateOrganizationOKBodyLicensingTypeModelPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateOrganizationOKBodyLicensing) validateModel(formats strfmt.Registry) error {
	if swag.IsZero(o.Model) { // not required
		return nil
	}

	// value enum
	if err := o.validateModelEnum("updateOrganizationOK"+"."+"licensing"+"."+"model", "body", o.Model); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update organization o k body licensing based on context it is used
func (o *UpdateOrganizationOKBodyLicensing) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyLicensing) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationOKBodyLicensing) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationOKBodyLicensing
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationParamsBodyAPI API-specific settings
swagger:model UpdateOrganizationParamsBodyAPI
*/
type UpdateOrganizationParamsBodyAPI struct {

	// If true, enable the access to the Cisco Meraki Dashboard API
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update organization params body API
func (o *UpdateOrganizationParamsBodyAPI) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization params body API based on context it is used
func (o *UpdateOrganizationParamsBodyAPI) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationParamsBodyAPI) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationParamsBodyAPI) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationParamsBodyAPI
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}