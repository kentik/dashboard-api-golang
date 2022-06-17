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
)

// UpdateOrganizationLoginSecurityReader is a Reader for the UpdateOrganizationLoginSecurity structure.
type UpdateOrganizationLoginSecurityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationLoginSecurityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationLoginSecurityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateOrganizationLoginSecurityOK creates a UpdateOrganizationLoginSecurityOK with default headers values
func NewUpdateOrganizationLoginSecurityOK() *UpdateOrganizationLoginSecurityOK {
	return &UpdateOrganizationLoginSecurityOK{}
}

/* UpdateOrganizationLoginSecurityOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationLoginSecurityOK struct {
	Payload *UpdateOrganizationLoginSecurityOKBody
}

// IsSuccess returns true when this update organization login security o k response has a 2xx status code
func (o *UpdateOrganizationLoginSecurityOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization login security o k response has a 3xx status code
func (o *UpdateOrganizationLoginSecurityOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization login security o k response has a 4xx status code
func (o *UpdateOrganizationLoginSecurityOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization login security o k response has a 5xx status code
func (o *UpdateOrganizationLoginSecurityOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization login security o k response a status code equal to that given
func (o *UpdateOrganizationLoginSecurityOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateOrganizationLoginSecurityOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/loginSecurity][%d] updateOrganizationLoginSecurityOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationLoginSecurityOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/loginSecurity][%d] updateOrganizationLoginSecurityOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationLoginSecurityOK) GetPayload() *UpdateOrganizationLoginSecurityOKBody {
	return o.Payload
}

func (o *UpdateOrganizationLoginSecurityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(UpdateOrganizationLoginSecurityOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateOrganizationLoginSecurityBody update organization login security body
// Example: {"accountLockoutAttempts":3,"apiAuthentication":{"ipRestrictionsForKeys":{"enabled":true,"ranges":["192.195.83.1","192.168.33.33"]}},"enforceAccountLockout":true,"enforceDifferentPasswords":true,"enforceIdleTimeout":true,"enforceLoginIpRanges":true,"enforcePasswordExpiration":true,"enforceStrongPasswords":true,"enforceTwoFactorAuth":true,"idleTimeoutMinutes":30,"loginIpRanges":["192.195.83.1","192.195.83.255"],"numDifferentPasswords":3,"passwordExpirationDays":90}
swagger:model UpdateOrganizationLoginSecurityBody
*/
type UpdateOrganizationLoginSecurityBody struct {

	// Number of consecutive failed login attempts after which users' accounts will be locked.
	AccountLockoutAttempts int64 `json:"accountLockoutAttempts,omitempty"`

	// api authentication
	APIAuthentication *UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication `json:"apiAuthentication,omitempty"`

	// Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.
	EnforceAccountLockout bool `json:"enforceAccountLockout,omitempty"`

	// Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.
	EnforceDifferentPasswords bool `json:"enforceDifferentPasswords,omitempty"`

	// Boolean indicating whether users will be logged out after being idle for the specified number of minutes.
	EnforceIdleTimeout bool `json:"enforceIdleTimeout,omitempty"`

	// Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.
	EnforceLoginIPRanges bool `json:"enforceLoginIpRanges,omitempty"`

	// Boolean indicating whether users are forced to change their password every X number of days.
	EnforcePasswordExpiration bool `json:"enforcePasswordExpiration,omitempty"`

	// Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol
	EnforceStrongPasswords bool `json:"enforceStrongPasswords,omitempty"`

	// Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the Google Authenticator application.
	EnforceTwoFactorAuth bool `json:"enforceTwoFactorAuth,omitempty"`

	// Number of minutes users can remain idle before being logged out of their accounts.
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes,omitempty"`

	// List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	LoginIPRanges []string `json:"loginIpRanges"`

	// Number of recent passwords that new password must be distinct from.
	NumDifferentPasswords int64 `json:"numDifferentPasswords,omitempty"`

	// Number of days after which users will be forced to change their password.
	PasswordExpirationDays int64 `json:"passwordExpirationDays,omitempty"`
}

// Validate validates this update organization login security body
func (o *UpdateOrganizationLoginSecurityBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAPIAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityBody) validateAPIAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(o.APIAuthentication) { // not required
		return nil
	}

	if o.APIAuthentication != nil {
		if err := o.APIAuthentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization login security body based on the context it is used
func (o *UpdateOrganizationLoginSecurityBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAPIAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityBody) contextValidateAPIAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if o.APIAuthentication != nil {
		if err := o.APIAuthentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationLoginSecurityBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationLoginSecurityOKBody update organization login security o k body
swagger:model UpdateOrganizationLoginSecurityOKBody
*/
type UpdateOrganizationLoginSecurityOKBody struct {

	// Number of consecutive failed login attempts after which users' accounts will be locked.
	AccountLockoutAttempts int64 `json:"accountLockoutAttempts,omitempty"`

	// api authentication
	APIAuthentication *UpdateOrganizationLoginSecurityOKBodyAPIAuthentication `json:"apiAuthentication,omitempty"`

	// Boolean indicating whether users' Dashboard accounts will be locked out after a specified number of consecutive failed login attempts.
	EnforceAccountLockout bool `json:"enforceAccountLockout,omitempty"`

	// Boolean indicating whether users, when setting a new password, are forced to choose a new password that is different from any past passwords.
	EnforceDifferentPasswords bool `json:"enforceDifferentPasswords,omitempty"`

	// Boolean indicating whether users will be logged out after being idle for the specified number of minutes.
	EnforceIdleTimeout bool `json:"enforceIdleTimeout,omitempty"`

	// Boolean indicating whether organization will restrict access to Dashboard (including the API) from certain IP addresses.
	EnforceLoginIPRanges bool `json:"enforceLoginIpRanges,omitempty"`

	// Boolean indicating whether users are forced to change their password every X number of days.
	EnforcePasswordExpiration bool `json:"enforcePasswordExpiration,omitempty"`

	// Boolean indicating whether users will be forced to choose strong passwords for their accounts. Strong passwords are at least 8 characters that contain 3 of the following: number, uppercase letter, lowercase letter, and symbol
	EnforceStrongPasswords bool `json:"enforceStrongPasswords,omitempty"`

	// Boolean indicating whether users in this organization will be required to use an extra verification code when logging in to Dashboard. This code will be sent to their mobile phone via SMS, or can be generated by the Google Authenticator application.
	EnforceTwoFactorAuth bool `json:"enforceTwoFactorAuth,omitempty"`

	// Number of minutes users can remain idle before being logged out of their accounts.
	IdleTimeoutMinutes int64 `json:"idleTimeoutMinutes,omitempty"`

	// List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	LoginIPRanges []string `json:"loginIpRanges"`

	// Number of recent passwords that new password must be distinct from.
	NumDifferentPasswords int64 `json:"numDifferentPasswords,omitempty"`

	// Number of days after which users will be forced to change their password.
	PasswordExpirationDays int64 `json:"passwordExpirationDays,omitempty"`
}

// Validate validates this update organization login security o k body
func (o *UpdateOrganizationLoginSecurityOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAPIAuthentication(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityOKBody) validateAPIAuthentication(formats strfmt.Registry) error {
	if swag.IsZero(o.APIAuthentication) { // not required
		return nil
	}

	if o.APIAuthentication != nil {
		if err := o.APIAuthentication.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization login security o k body based on the context it is used
func (o *UpdateOrganizationLoginSecurityOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAPIAuthentication(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityOKBody) contextValidateAPIAuthentication(ctx context.Context, formats strfmt.Registry) error {

	if o.APIAuthentication != nil {
		if err := o.APIAuthentication.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityOKBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationLoginSecurityOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationLoginSecurityOKBodyAPIAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
swagger:model UpdateOrganizationLoginSecurityOKBodyAPIAuthentication
*/
type UpdateOrganizationLoginSecurityOKBodyAPIAuthentication struct {

	// ip restrictions for keys
	IPRestrictionsForKeys *UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// Validate validates this update organization login security o k body API authentication
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPRestrictionsForKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthentication) validateIPRestrictionsForKeys(formats strfmt.Registry) error {
	if swag.IsZero(o.IPRestrictionsForKeys) { // not required
		return nil
	}

	if o.IPRestrictionsForKeys != nil {
		if err := o.IPRestrictionsForKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization login security o k body API authentication based on the context it is used
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPRestrictionsForKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthentication) contextValidateIPRestrictionsForKeys(ctx context.Context, formats strfmt.Registry) error {

	if o.IPRestrictionsForKeys != nil {
		if err := o.IPRestrictionsForKeys.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurityOK" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthentication) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthentication) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationLoginSecurityOKBodyAPIAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys Details for API-only IP restrictions.
swagger:model UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys
*/
type UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys struct {

	// Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
	Enabled bool `json:"enabled,omitempty"`

	// List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	Ranges []string `json:"ranges"`
}

// Validate validates this update organization login security o k body API authentication IP restrictions for keys
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization login security o k body API authentication IP restrictions for keys based on context it is used
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationLoginSecurityOKBodyAPIAuthenticationIPRestrictionsForKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication Details for indicating whether organization will restrict access to API (but not Dashboard) to certain IP addresses.
swagger:model UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication
*/
type UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication struct {

	// ip restrictions for keys
	IPRestrictionsForKeys *UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys `json:"ipRestrictionsForKeys,omitempty"`
}

// Validate validates this update organization login security params body API authentication
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateIPRestrictionsForKeys(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication) validateIPRestrictionsForKeys(formats strfmt.Registry) error {
	if swag.IsZero(o.IPRestrictionsForKeys) { // not required
		return nil
	}

	if o.IPRestrictionsForKeys != nil {
		if err := o.IPRestrictionsForKeys.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update organization login security params body API authentication based on the context it is used
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateIPRestrictionsForKeys(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication) contextValidateIPRestrictionsForKeys(ctx context.Context, formats strfmt.Registry) error {

	if o.IPRestrictionsForKeys != nil {
		if err := o.IPRestrictionsForKeys.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationLoginSecurity" + "." + "apiAuthentication" + "." + "ipRestrictionsForKeys")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationLoginSecurityParamsBodyAPIAuthentication
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys Details for API-only IP restrictions.
swagger:model UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys
*/
type UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys struct {

	// Boolean indicating whether the organization will restrict API key (not Dashboard GUI) usage to a specific list of IP addresses or CIDR ranges.
	Enabled bool `json:"enabled,omitempty"`

	// List of acceptable IP ranges. Entries can be single IP addresses, IP address ranges, and CIDR subnets.
	Ranges []string `json:"ranges"`
}

// Validate validates this update organization login security params body API authentication IP restrictions for keys
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization login security params body API authentication IP restrictions for keys based on context it is used
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationLoginSecurityParamsBodyAPIAuthenticationIPRestrictionsForKeys
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
