// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// CreateNetworkMerakiAuthUserReader is a Reader for the CreateNetworkMerakiAuthUser structure.
type CreateNetworkMerakiAuthUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkMerakiAuthUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkMerakiAuthUserCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/merakiAuthUsers] createNetworkMerakiAuthUser", response, response.Code())
	}
}

// NewCreateNetworkMerakiAuthUserCreated creates a CreateNetworkMerakiAuthUserCreated with default headers values
func NewCreateNetworkMerakiAuthUserCreated() *CreateNetworkMerakiAuthUserCreated {
	return &CreateNetworkMerakiAuthUserCreated{}
}

/*
CreateNetworkMerakiAuthUserCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkMerakiAuthUserCreated struct {
	Payload *CreateNetworkMerakiAuthUserCreatedBody
}

// IsSuccess returns true when this create network meraki auth user created response has a 2xx status code
func (o *CreateNetworkMerakiAuthUserCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network meraki auth user created response has a 3xx status code
func (o *CreateNetworkMerakiAuthUserCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network meraki auth user created response has a 4xx status code
func (o *CreateNetworkMerakiAuthUserCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network meraki auth user created response has a 5xx status code
func (o *CreateNetworkMerakiAuthUserCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network meraki auth user created response a status code equal to that given
func (o *CreateNetworkMerakiAuthUserCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network meraki auth user created response
func (o *CreateNetworkMerakiAuthUserCreated) Code() int {
	return 201
}

func (o *CreateNetworkMerakiAuthUserCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/merakiAuthUsers][%d] createNetworkMerakiAuthUserCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkMerakiAuthUserCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/merakiAuthUsers][%d] createNetworkMerakiAuthUserCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkMerakiAuthUserCreated) GetPayload() *CreateNetworkMerakiAuthUserCreatedBody {
	return o.Payload
}

func (o *CreateNetworkMerakiAuthUserCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateNetworkMerakiAuthUserCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkMerakiAuthUserBody create network meraki auth user body
// Example: {"accountType":"802.1X","authorizations":[{"expiresAt":"2018-03-13T00:00:00.090210Z","ssidNumber":1}],"email":"miles@meraki.com","emailPasswordToUser":false,"isAdmin":false,"name":"Miles Meraki","password":"secret"}
swagger:model CreateNetworkMerakiAuthUserBody
*/
type CreateNetworkMerakiAuthUserBody struct {

	// Authorization type for user. Can be 'Guest' or '802.1X' for wireless networks, or 'Client VPN' for MX networks. Defaults to '802.1X'.
	// Enum: [802.1X Client VPN Guest]
	AccountType *string `json:"accountType,omitempty"`

	// Authorization zones and expiration dates for the user.
	// Required: true
	Authorizations []*CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0 `json:"authorizations"`

	// Email address of the user
	// Required: true
	Email *string `json:"email"`

	// Whether or not Meraki should email the password to user. Default is false.
	EmailPasswordToUser bool `json:"emailPasswordToUser,omitempty"`

	// Whether or not the user is a Dashboard administrator.
	IsAdmin bool `json:"isAdmin,omitempty"`

	// Name of the user. Only required If the user is not a Dashboard administrator.
	Name string `json:"name,omitempty"`

	// The password for this user account. Only required If the user is not a Dashboard administrator.
	Password string `json:"password,omitempty"`
}

// Validate validates this create network meraki auth user body
func (o *CreateNetworkMerakiAuthUserBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccountType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkMerakiAuthUserBodyTypeAccountTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["802.1X","Client VPN","Guest"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkMerakiAuthUserBodyTypeAccountTypePropEnum = append(createNetworkMerakiAuthUserBodyTypeAccountTypePropEnum, v)
	}
}

const (

	// CreateNetworkMerakiAuthUserBodyAccountTypeNr802Dot1X captures enum value "802.1X"
	CreateNetworkMerakiAuthUserBodyAccountTypeNr802Dot1X string = "802.1X"

	// CreateNetworkMerakiAuthUserBodyAccountTypeClientVPN captures enum value "Client VPN"
	CreateNetworkMerakiAuthUserBodyAccountTypeClientVPN string = "Client VPN"

	// CreateNetworkMerakiAuthUserBodyAccountTypeGuest captures enum value "Guest"
	CreateNetworkMerakiAuthUserBodyAccountTypeGuest string = "Guest"
)

// prop value enum
func (o *CreateNetworkMerakiAuthUserBody) validateAccountTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkMerakiAuthUserBodyTypeAccountTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkMerakiAuthUserBody) validateAccountType(formats strfmt.Registry) error {
	if swag.IsZero(o.AccountType) { // not required
		return nil
	}

	// value enum
	if err := o.validateAccountTypeEnum("createNetworkMerakiAuthUser"+"."+"accountType", "body", *o.AccountType); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkMerakiAuthUserBody) validateAuthorizations(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkMerakiAuthUser"+"."+"authorizations", "body", o.Authorizations); err != nil {
		return err
	}

	for i := 0; i < len(o.Authorizations); i++ {
		if swag.IsZero(o.Authorizations[i]) { // not required
			continue
		}

		if o.Authorizations[i] != nil {
			if err := o.Authorizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkMerakiAuthUser" + "." + "authorizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkMerakiAuthUser" + "." + "authorizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkMerakiAuthUserBody) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkMerakiAuthUser"+"."+"email", "body", o.Email); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network meraki auth user body based on the context it is used
func (o *CreateNetworkMerakiAuthUserBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAuthorizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkMerakiAuthUserBody) contextValidateAuthorizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Authorizations); i++ {

		if o.Authorizations[i] != nil {

			if swag.IsZero(o.Authorizations[i]) { // not required
				return nil
			}

			if err := o.Authorizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkMerakiAuthUser" + "." + "authorizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkMerakiAuthUser" + "." + "authorizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkMerakiAuthUserBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkMerakiAuthUserCreatedBody create network meraki auth user created body
swagger:model CreateNetworkMerakiAuthUserCreatedBody
*/
type CreateNetworkMerakiAuthUserCreatedBody struct {

	// Authorization type for user.
	AccountType string `json:"accountType,omitempty"`

	// User authorization info
	Authorizations []*CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0 `json:"authorizations"`

	// Creation time of the user
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Email address of the user
	Email string `json:"email,omitempty"`

	// Meraki auth user id
	ID string `json:"id,omitempty"`

	// Whether or not the user is a Dashboard administrator
	IsAdmin bool `json:"isAdmin,omitempty"`

	// Name of the user
	Name string `json:"name,omitempty"`
}

// Validate validates this create network meraki auth user created body
func (o *CreateNetworkMerakiAuthUserCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAuthorizations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkMerakiAuthUserCreatedBody) validateAuthorizations(formats strfmt.Registry) error {
	if swag.IsZero(o.Authorizations) { // not required
		return nil
	}

	for i := 0; i < len(o.Authorizations); i++ {
		if swag.IsZero(o.Authorizations[i]) { // not required
			continue
		}

		if o.Authorizations[i] != nil {
			if err := o.Authorizations[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkMerakiAuthUserCreated" + "." + "authorizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkMerakiAuthUserCreated" + "." + "authorizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateNetworkMerakiAuthUserCreatedBody) validateCreatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createNetworkMerakiAuthUserCreated"+"."+"createdAt", "body", "date-time", o.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network meraki auth user created body based on the context it is used
func (o *CreateNetworkMerakiAuthUserCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAuthorizations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkMerakiAuthUserCreatedBody) contextValidateAuthorizations(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Authorizations); i++ {

		if o.Authorizations[i] != nil {

			if swag.IsZero(o.Authorizations[i]) { // not required
				return nil
			}

			if err := o.Authorizations[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createNetworkMerakiAuthUserCreated" + "." + "authorizations" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createNetworkMerakiAuthUserCreated" + "." + "authorizations" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkMerakiAuthUserCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0 create network meraki auth user created body authorizations items0
swagger:model CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0
*/
type CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0 struct {

	// User is authorized by the account email address
	AuthorizedByEmail string `json:"authorizedByEmail,omitempty"`

	// User is authorized by the account name
	AuthorizedByName string `json:"authorizedByName,omitempty"`

	// Authorized zone of the user
	AuthorizedZone string `json:"authorizedZone,omitempty"`

	// Authorization expiration time
	// Format: date-time
	ExpiresAt strfmt.DateTime `json:"expiresAt,omitempty"`

	// SSID number
	SsidNumber int64 `json:"ssidNumber,omitempty"`
}

// Validate validates this create network meraki auth user created body authorizations items0
func (o *CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateExpiresAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0) validateExpiresAt(formats strfmt.Registry) error {
	if swag.IsZero(o.ExpiresAt) { // not required
		return nil
	}

	if err := validate.FormatOf("expiresAt", "body", "date-time", o.ExpiresAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network meraki auth user created body authorizations items0 based on context it is used
func (o *CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkMerakiAuthUserCreatedBodyAuthorizationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0 create network meraki auth user params body authorizations items0
swagger:model CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0
*/
type CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0 struct {

	// Date for authorization to expire. Set to 'Never' for the authorization to not expire, which is the default.
	ExpiresAt *string `json:"expiresAt,omitempty"`

	// Required for wireless networks. The SSID for which the user is being authorized, which must be configured for the user's given accountType.
	SsidNumber int64 `json:"ssidNumber,omitempty"`
}

// Validate validates this create network meraki auth user params body authorizations items0
func (o *CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network meraki auth user params body authorizations items0 based on context it is used
func (o *CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0) UnmarshalBinary(b []byte) error {
	var res CreateNetworkMerakiAuthUserParamsBodyAuthorizationsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
