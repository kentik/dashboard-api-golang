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

// GetOrganizationLicensesReader is a Reader for the GetOrganizationLicenses structure.
type GetOrganizationLicensesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationLicensesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationLicensesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationLicensesOK creates a GetOrganizationLicensesOK with default headers values
func NewGetOrganizationLicensesOK() *GetOrganizationLicensesOK {
	return &GetOrganizationLicensesOK{}
}

/* GetOrganizationLicensesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationLicensesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetOrganizationLicensesOKBodyItems0
}

// IsSuccess returns true when this get organization licenses o k response has a 2xx status code
func (o *GetOrganizationLicensesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization licenses o k response has a 3xx status code
func (o *GetOrganizationLicensesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization licenses o k response has a 4xx status code
func (o *GetOrganizationLicensesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization licenses o k response has a 5xx status code
func (o *GetOrganizationLicensesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization licenses o k response a status code equal to that given
func (o *GetOrganizationLicensesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationLicensesOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/licenses][%d] getOrganizationLicensesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLicensesOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/licenses][%d] getOrganizationLicensesOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationLicensesOK) GetPayload() []*GetOrganizationLicensesOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationLicensesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

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

/*GetOrganizationLicensesOKBodyItems0 get organization licenses o k body items0
swagger:model GetOrganizationLicensesOKBodyItems0
*/
type GetOrganizationLicensesOKBodyItems0 struct {

	// The date the license started burning
	ActivationDate string `json:"activationDate,omitempty"`

	// The date the license was claimed into the organization
	ClaimDate string `json:"claimDate,omitempty"`

	// Serial number of the device the license is assigned to
	DeviceSerial string `json:"deviceSerial,omitempty"`

	// The duration of the individual license
	DurationInDays int64 `json:"durationInDays,omitempty"`

	// The date the license will expire
	ExpirationDate string `json:"expirationDate,omitempty"`

	// License ID
	ID string `json:"id,omitempty"`

	// License key
	LicenseKey string `json:"licenseKey,omitempty"`

	// License type
	LicenseType string `json:"licenseType,omitempty"`

	// ID of the network the license is assigned to
	NetworkID string `json:"networkId,omitempty"`

	// Order number
	OrderNumber string `json:"orderNumber,omitempty"`

	// List of permanently queued licenses attached to the license
	PermanentlyQueuedLicenses []*GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0 `json:"permanentlyQueuedLicenses"`

	// The number of seats of the license. Only applicable to SM licenses.
	SeatCount int64 `json:"seatCount,omitempty"`

	// The state of the license
	// Enum: [active expired expiring unused unusedActive recentlyQueued]
	State string `json:"state,omitempty"`

	// The duration of the license plus all permanently queued licenses associated with it
	TotalDurationInDays int64 `json:"totalDurationInDays,omitempty"`
}

// Validate validates this get organization licenses o k body items0
func (o *GetOrganizationLicensesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePermanentlyQueuedLicenses(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLicensesOKBodyItems0) validatePermanentlyQueuedLicenses(formats strfmt.Registry) error {
	if swag.IsZero(o.PermanentlyQueuedLicenses) { // not required
		return nil
	}

	for i := 0; i < len(o.PermanentlyQueuedLicenses); i++ {
		if swag.IsZero(o.PermanentlyQueuedLicenses[i]) { // not required
			continue
		}

		if o.PermanentlyQueuedLicenses[i] != nil {
			if err := o.PermanentlyQueuedLicenses[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var getOrganizationLicensesOKBodyItems0TypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","expired","expiring","unused","unusedActive","recentlyQueued"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getOrganizationLicensesOKBodyItems0TypeStatePropEnum = append(getOrganizationLicensesOKBodyItems0TypeStatePropEnum, v)
	}
}

const (

	// GetOrganizationLicensesOKBodyItems0StateActive captures enum value "active"
	GetOrganizationLicensesOKBodyItems0StateActive string = "active"

	// GetOrganizationLicensesOKBodyItems0StateExpired captures enum value "expired"
	GetOrganizationLicensesOKBodyItems0StateExpired string = "expired"

	// GetOrganizationLicensesOKBodyItems0StateExpiring captures enum value "expiring"
	GetOrganizationLicensesOKBodyItems0StateExpiring string = "expiring"

	// GetOrganizationLicensesOKBodyItems0StateUnused captures enum value "unused"
	GetOrganizationLicensesOKBodyItems0StateUnused string = "unused"

	// GetOrganizationLicensesOKBodyItems0StateUnusedActive captures enum value "unusedActive"
	GetOrganizationLicensesOKBodyItems0StateUnusedActive string = "unusedActive"

	// GetOrganizationLicensesOKBodyItems0StateRecentlyQueued captures enum value "recentlyQueued"
	GetOrganizationLicensesOKBodyItems0StateRecentlyQueued string = "recentlyQueued"
)

// prop value enum
func (o *GetOrganizationLicensesOKBodyItems0) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getOrganizationLicensesOKBodyItems0TypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetOrganizationLicensesOKBodyItems0) validateState(formats strfmt.Registry) error {
	if swag.IsZero(o.State) { // not required
		return nil
	}

	// value enum
	if err := o.validateStateEnum("state", "body", o.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get organization licenses o k body items0 based on the context it is used
func (o *GetOrganizationLicensesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePermanentlyQueuedLicenses(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationLicensesOKBodyItems0) contextValidatePermanentlyQueuedLicenses(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.PermanentlyQueuedLicenses); i++ {

		if o.PermanentlyQueuedLicenses[i] != nil {
			if err := o.PermanentlyQueuedLicenses[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("permanentlyQueuedLicenses" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLicensesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLicensesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLicensesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0 get organization licenses o k body items0 permanently queued licenses items0
swagger:model GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0
*/
type GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0 struct {

	// The duration of the individual license
	DurationInDays int64 `json:"durationInDays,omitempty"`

	// Permanently queued license ID
	ID string `json:"id,omitempty"`

	// License key
	LicenseKey string `json:"licenseKey,omitempty"`

	// License type
	LicenseType string `json:"licenseType,omitempty"`

	// Order number
	OrderNumber string `json:"orderNumber,omitempty"`
}

// Validate validates this get organization licenses o k body items0 permanently queued licenses items0
func (o *GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization licenses o k body items0 permanently queued licenses items0 based on context it is used
func (o *GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationLicensesOKBodyItems0PermanentlyQueuedLicensesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
