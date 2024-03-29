// Code generated by go-swagger; DO NOT EDIT.

package licensing

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

// GetOrganizationLicensingCotermLicensesReader is a Reader for the GetOrganizationLicensingCotermLicenses structure.
type GetOrganizationLicensingCotermLicensesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationLicensingCotermLicensesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationLicensingCotermLicensesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/licensing/coterm/licenses] getOrganizationLicensingCotermLicenses", response, response.Code())
	}
}

// NewGetOrganizationLicensingCotermLicensesOK creates a GetOrganizationLicensingCotermLicensesOK with default headers values
func NewGetOrganizationLicensingCotermLicensesOK() *GetOrganizationLicensingCotermLicensesOK {
	return &GetOrganizationLicensingCotermLicensesOK{}
}

/*
GetOrganizationLicensingCotermLicensesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationLicensingCotermLicensesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetOrganizationLicensingCotermLicensesOKBodyItems0
}

// IsSuccess returns true when this get organization licensing coterm licenses o k response has a 2xx status code
func (o *GetOrganizationLicensingCotermLicensesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization licensing coterm licenses o k response has a 3xx status code
func (o *GetOrganizationLicensingCotermLicensesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization licensing coterm licenses o k response has a 4xx status code
func (o *GetOrganizationLicensingCotermLicensesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization licensing coterm licenses o k response has a 5xx status code
func (o *GetOrganizationLicensingCotermLicensesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization licensing coterm licenses o k response a status code equal to that given
func (o *GetOrganizationLicensingCotermLicensesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization licensing coterm licenses o k response
func (o *GetOrganizationLicensingCotermLicensesOK) Code() int {
	return 200
}

func (o *GetOrganizationLicensingCotermLicensesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/licensing/coterm/licenses][%d] getOrganizationLicensingCotermLicensesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLicensingCotermLicensesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/licensing/coterm/licenses][%d] getOrganizationLicensingCotermLicensesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLicensingCotermLicensesOK) GetPayload() []*GetOrganizationLicensingCotermLicensesOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationLicensingCotermLicensesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*
GetOrganizationLicensingCotermLicensesOKBodyItems0 get organization licensing coterm licenses o k body items0
swagger:model GetOrganizationLicensingCotermLicensesOKBodyItems0
*/
type GetOrganizationLicensingCotermLicensesOKBodyItems0 struct {

	// When the license was claimed into the organization
	// Format: date-time
	ClaimedAt strfmt.DateTime `json:"claimedAt,omitempty"`

	// The counts of the license by model type
	Counts []*GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0 `json:"counts"`

	// The duration (term length) of the license, measured in days
	Duration int64 `json:"duration,omitempty"`

	// The editions of the license for each relevant product type
	Editions []*GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0 `json:"editions"`

	// Flag to indicate if the license is expired
	Expired bool `json:"expired,omitempty"`

	// Flag to indicated that the license is invalidated
	Invalidated bool `json:"invalidated,omitempty"`

	// When the license was invalidated. Will be null for active licenses
	// Format: date-time
	InvalidatedAt strfmt.DateTime `json:"invalidatedAt,omitempty"`

	// The key of the license
	Key string `json:"key,omitempty"`

	// The operation mode of the license when it was claimed
	// Enum: [addDevices renew]
	Mode string `json:"mode,omitempty"`

	// The ID of the organization that the license is claimed in
	OrganizationID string `json:"organizationId,omitempty"`

	// When the license's term began (approximately the date when the license was created)
	// Format: date-time
	StartedAt strfmt.DateTime `json:"startedAt,omitempty"`
}

// Validate validates this get organization licensing coterm licenses o k body items0
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClaimedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateCounts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEditions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateInvalidatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) validateClaimedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.ClaimedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("claimedAt", "body", "date-time", o.ClaimedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) validateCounts(formats strfmt.Registry) error {
	if swag.IsZero(o.Counts) { // not required
		return nil
	}

	for i := 0; i < len(o.Counts); i++ {
		if swag.IsZero(o.Counts[i]) { // not required
			continue
		}

		if o.Counts[i] != nil {
			if err := o.Counts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("counts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) validateEditions(formats strfmt.Registry) error {
	if swag.IsZero(o.Editions) { // not required
		return nil
	}

	for i := 0; i < len(o.Editions); i++ {
		if swag.IsZero(o.Editions[i]) { // not required
			continue
		}

		if o.Editions[i] != nil {
			if err := o.Editions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("editions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("editions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) validateInvalidatedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.InvalidatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("invalidatedAt", "body", "date-time", o.InvalidatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var getOrganizationLicensingCotermLicensesOKBodyItems0TypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["addDevices","renew"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationLicensingCotermLicensesOKBodyItems0TypeModePropEnum = append(getOrganizationLicensingCotermLicensesOKBodyItems0TypeModePropEnum, v)
	}
}

const (

	// GetOrganizationLicensingCotermLicensesOKBodyItems0ModeAddDevices captures enum value "addDevices"
	GetOrganizationLicensingCotermLicensesOKBodyItems0ModeAddDevices string = "addDevices"

	// GetOrganizationLicensingCotermLicensesOKBodyItems0ModeRenew captures enum value "renew"
	GetOrganizationLicensingCotermLicensesOKBodyItems0ModeRenew string = "renew"
)

// prop value enum
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationLicensingCotermLicensesOKBodyItems0TypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("mode", "body", o.Mode); err != nil {
		return err
	}

	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) validateStartedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.StartedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("startedAt", "body", "date-time", o.StartedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organization licensing coterm licenses o k body items0 based on the context it is used
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateEditions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) contextValidateCounts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Counts); i++ {

		if o.Counts[i] != nil {

			if swag.IsZero(o.Counts[i]) { // not required
				return nil
			}

			if err := o.Counts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("counts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("counts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) contextValidateEditions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Editions); i++ {

		if o.Editions[i] != nil {

			if swag.IsZero(o.Editions[i]) { // not required
				return nil
			}

			if err := o.Editions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("editions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("editions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLicensingCotermLicensesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0 get organization licensing coterm licenses o k body items0 counts items0
swagger:model GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0
*/
type GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0 struct {

	// The number of counts the license contains of this model
	Count int64 `json:"count,omitempty"`

	// The license model type
	Model string `json:"model,omitempty"`
}

// Validate validates this get organization licensing coterm licenses o k body items0 counts items0
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization licensing coterm licenses o k body items0 counts items0 based on context it is used
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLicensingCotermLicensesOKBodyItems0CountsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0 get organization licensing coterm licenses o k body items0 editions items0
swagger:model GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0
*/
type GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0 struct {

	// The name of the license edition
	Edition string `json:"edition,omitempty"`

	// The product type of the license edition
	ProductType string `json:"productType,omitempty"`
}

// Validate validates this get organization licensing coterm licenses o k body items0 editions items0
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization licensing coterm licenses o k body items0 editions items0 based on context it is used
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLicensingCotermLicensesOKBodyItems0EditionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
