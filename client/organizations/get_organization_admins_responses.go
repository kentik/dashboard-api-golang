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

// GetOrganizationAdminsReader is a Reader for the GetOrganizationAdmins structure.
type GetOrganizationAdminsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationAdminsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationAdminsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/admins] getOrganizationAdmins", response, response.Code())
	}
}

// NewGetOrganizationAdminsOK creates a GetOrganizationAdminsOK with default headers values
func NewGetOrganizationAdminsOK() *GetOrganizationAdminsOK {
	return &GetOrganizationAdminsOK{}
}

/*
GetOrganizationAdminsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationAdminsOK struct {
	Payload []*GetOrganizationAdminsOKBodyItems0
}

// IsSuccess returns true when this get organization admins o k response has a 2xx status code
func (o *GetOrganizationAdminsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization admins o k response has a 3xx status code
func (o *GetOrganizationAdminsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization admins o k response has a 4xx status code
func (o *GetOrganizationAdminsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization admins o k response has a 5xx status code
func (o *GetOrganizationAdminsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization admins o k response a status code equal to that given
func (o *GetOrganizationAdminsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization admins o k response
func (o *GetOrganizationAdminsOK) Code() int {
	return 200
}

func (o *GetOrganizationAdminsOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/admins][%d] getOrganizationAdminsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdminsOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/admins][%d] getOrganizationAdminsOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationAdminsOK) GetPayload() []*GetOrganizationAdminsOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationAdminsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationAdminsOKBodyItems0 get organization admins o k body items0
swagger:model GetOrganizationAdminsOKBodyItems0
*/
type GetOrganizationAdminsOKBodyItems0 struct {

	// Status of the admin's account
	// Enum: [locked ok pending unverified]
	AccountStatus string `json:"accountStatus,omitempty"`

	// Admin's authentication method
	// Enum: [Cisco SecureX Sign-On Email]
	AuthenticationMethod string `json:"authenticationMethod,omitempty"`

	// Admin's email address
	Email string `json:"email,omitempty"`

	// Indicates whether the admin has an API key
	HasAPIKey bool `json:"hasApiKey,omitempty"`

	// Admin's ID
	ID string `json:"id,omitempty"`

	// Time when the admin was last active
	// Format: date-time
	LastActive strfmt.DateTime `json:"lastActive,omitempty"`

	// Admin's username
	Name string `json:"name,omitempty"`

	// Admin network access information
	Networks []*GetOrganizationAdminsOKBodyItems0NetworksItems0 `json:"networks"`

	// Admin's level of access to the organization
	// Enum: [enterprise full none read-only]
	OrgAccess string `json:"orgAccess,omitempty"`

	// Admin tag information
	Tags []*GetOrganizationAdminsOKBodyItems0TagsItems0 `json:"tags"`

	// Indicates whether two-factor authentication is enabled
	TwoFactorAuthEnabled bool `json:"twoFactorAuthEnabled,omitempty"`
}

// Validate validates this get organization admins o k body items0
func (o *GetOrganizationAdminsOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccountStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthenticationMethod(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateLastActive(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateNetworks(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateOrgAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getOrganizationAdminsOKBodyItems0TypeAccountStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["locked","ok","pending","unverified"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationAdminsOKBodyItems0TypeAccountStatusPropEnum = append(getOrganizationAdminsOKBodyItems0TypeAccountStatusPropEnum, v)
	}
}

const (

	// GetOrganizationAdminsOKBodyItems0AccountStatusLocked captures enum value "locked"
	GetOrganizationAdminsOKBodyItems0AccountStatusLocked string = "locked"

	// GetOrganizationAdminsOKBodyItems0AccountStatusOk captures enum value "ok"
	GetOrganizationAdminsOKBodyItems0AccountStatusOk string = "ok"

	// GetOrganizationAdminsOKBodyItems0AccountStatusPending captures enum value "pending"
	GetOrganizationAdminsOKBodyItems0AccountStatusPending string = "pending"

	// GetOrganizationAdminsOKBodyItems0AccountStatusUnverified captures enum value "unverified"
	GetOrganizationAdminsOKBodyItems0AccountStatusUnverified string = "unverified"
)

// prop value enum
func (o *GetOrganizationAdminsOKBodyItems0) validateAccountStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationAdminsOKBodyItems0TypeAccountStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) validateAccountStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.AccountStatus) { // not required
		return nil
	}

	// value enum
	if err := o.validateAccountStatusEnum("accountStatus", "body", o.AccountStatus); err != nil {
		return err
	}

	return nil
}

var getOrganizationAdminsOKBodyItems0TypeAuthenticationMethodPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["Cisco SecureX Sign-On","Email"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationAdminsOKBodyItems0TypeAuthenticationMethodPropEnum = append(getOrganizationAdminsOKBodyItems0TypeAuthenticationMethodPropEnum, v)
	}
}

const (

	// GetOrganizationAdminsOKBodyItems0AuthenticationMethodCiscoSecureXSignDashOn captures enum value "Cisco SecureX Sign-On"
	GetOrganizationAdminsOKBodyItems0AuthenticationMethodCiscoSecureXSignDashOn string = "Cisco SecureX Sign-On"

	// GetOrganizationAdminsOKBodyItems0AuthenticationMethodEmail captures enum value "Email"
	GetOrganizationAdminsOKBodyItems0AuthenticationMethodEmail string = "Email"
)

// prop value enum
func (o *GetOrganizationAdminsOKBodyItems0) validateAuthenticationMethodEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationAdminsOKBodyItems0TypeAuthenticationMethodPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) validateAuthenticationMethod(formats strfmt.Registry) error {
	if swag.IsZero(o.AuthenticationMethod) { // not required
		return nil
	}

	// value enum
	if err := o.validateAuthenticationMethodEnum("authenticationMethod", "body", o.AuthenticationMethod); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) validateLastActive(formats strfmt.Registry) error {
	if swag.IsZero(o.LastActive) { // not required
		return nil
	}

	if err := validate.FormatOf("lastActive", "body", "date-time", o.LastActive.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) validateNetworks(formats strfmt.Registry) error {
	if swag.IsZero(o.Networks) { // not required
		return nil
	}

	for i := 0; i < len(o.Networks); i++ {
		if swag.IsZero(o.Networks[i]) { // not required
			continue
		}

		if o.Networks[i] != nil {
			if err := o.Networks[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getOrganizationAdminsOKBodyItems0TypeOrgAccessPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["enterprise","full","none","read-only"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationAdminsOKBodyItems0TypeOrgAccessPropEnum = append(getOrganizationAdminsOKBodyItems0TypeOrgAccessPropEnum, v)
	}
}

const (

	// GetOrganizationAdminsOKBodyItems0OrgAccessEnterprise captures enum value "enterprise"
	GetOrganizationAdminsOKBodyItems0OrgAccessEnterprise string = "enterprise"

	// GetOrganizationAdminsOKBodyItems0OrgAccessFull captures enum value "full"
	GetOrganizationAdminsOKBodyItems0OrgAccessFull string = "full"

	// GetOrganizationAdminsOKBodyItems0OrgAccessNone captures enum value "none"
	GetOrganizationAdminsOKBodyItems0OrgAccessNone string = "none"

	// GetOrganizationAdminsOKBodyItems0OrgAccessReadDashOnly captures enum value "read-only"
	GetOrganizationAdminsOKBodyItems0OrgAccessReadDashOnly string = "read-only"
)

// prop value enum
func (o *GetOrganizationAdminsOKBodyItems0) validateOrgAccessEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationAdminsOKBodyItems0TypeOrgAccessPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) validateOrgAccess(formats strfmt.Registry) error {
	if swag.IsZero(o.OrgAccess) { // not required
		return nil
	}

	// value enum
	if err := o.validateOrgAccessEnum("orgAccess", "body", o.OrgAccess); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(o.Tags) { // not required
		return nil
	}

	for i := 0; i < len(o.Tags); i++ {
		if swag.IsZero(o.Tags[i]) { // not required
			continue
		}

		if o.Tags[i] != nil {
			if err := o.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get organization admins o k body items0 based on the context it is used
func (o *GetOrganizationAdminsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNetworks(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) contextValidateNetworks(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Networks); i++ {

		if o.Networks[i] != nil {

			if swag.IsZero(o.Networks[i]) { // not required
				return nil
			}

			if err := o.Networks[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("networks" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("networks" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationAdminsOKBodyItems0) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Tags); i++ {

		if o.Tags[i] != nil {

			if swag.IsZero(o.Tags[i]) { // not required
				return nil
			}

			if err := o.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdminsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdminsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdminsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationAdminsOKBodyItems0NetworksItems0 get organization admins o k body items0 networks items0
swagger:model GetOrganizationAdminsOKBodyItems0NetworksItems0
*/
type GetOrganizationAdminsOKBodyItems0NetworksItems0 struct {

	// Admin's level of access to the network
	Access string `json:"access,omitempty"`

	// Network ID
	ID string `json:"id,omitempty"`
}

// Validate validates this get organization admins o k body items0 networks items0
func (o *GetOrganizationAdminsOKBodyItems0NetworksItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization admins o k body items0 networks items0 based on context it is used
func (o *GetOrganizationAdminsOKBodyItems0NetworksItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdminsOKBodyItems0NetworksItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdminsOKBodyItems0NetworksItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdminsOKBodyItems0NetworksItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationAdminsOKBodyItems0TagsItems0 get organization admins o k body items0 tags items0
swagger:model GetOrganizationAdminsOKBodyItems0TagsItems0
*/
type GetOrganizationAdminsOKBodyItems0TagsItems0 struct {

	// Access level for the tag
	Access string `json:"access,omitempty"`

	// Tag value
	Tag string `json:"tag,omitempty"`
}

// Validate validates this get organization admins o k body items0 tags items0
func (o *GetOrganizationAdminsOKBodyItems0TagsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization admins o k body items0 tags items0 based on context it is used
func (o *GetOrganizationAdminsOKBodyItems0TagsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationAdminsOKBodyItems0TagsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationAdminsOKBodyItems0TagsItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationAdminsOKBodyItems0TagsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
