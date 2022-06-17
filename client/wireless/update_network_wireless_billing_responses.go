// Code generated by go-swagger; DO NOT EDIT.

package wireless

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

// UpdateNetworkWirelessBillingReader is a Reader for the UpdateNetworkWirelessBilling structure.
type UpdateNetworkWirelessBillingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkWirelessBillingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkWirelessBillingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkWirelessBillingOK creates a UpdateNetworkWirelessBillingOK with default headers values
func NewUpdateNetworkWirelessBillingOK() *UpdateNetworkWirelessBillingOK {
	return &UpdateNetworkWirelessBillingOK{}
}

/* UpdateNetworkWirelessBillingOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkWirelessBillingOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network wireless billing o k response has a 2xx status code
func (o *UpdateNetworkWirelessBillingOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network wireless billing o k response has a 3xx status code
func (o *UpdateNetworkWirelessBillingOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network wireless billing o k response has a 4xx status code
func (o *UpdateNetworkWirelessBillingOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network wireless billing o k response has a 5xx status code
func (o *UpdateNetworkWirelessBillingOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network wireless billing o k response a status code equal to that given
func (o *UpdateNetworkWirelessBillingOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkWirelessBillingOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/billing][%d] updateNetworkWirelessBillingOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessBillingOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/wireless/billing][%d] updateNetworkWirelessBillingOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkWirelessBillingOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkWirelessBillingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkWirelessBillingBody update network wireless billing body
// Example: {"currency":"USD","plans":[{"bandwidthLimits":{"limitDown":1000,"limitUp":1000},"id":"1","price":5,"timeLimit":"1 hour"}]}
swagger:model UpdateNetworkWirelessBillingBody
*/
type UpdateNetworkWirelessBillingBody struct {

	// The currency code of this node group's billing plans
	Currency string `json:"currency,omitempty"`

	// Array of billing plans in the node group. (Can configure a maximum of 5)
	Plans []*UpdateNetworkWirelessBillingParamsBodyPlansItems0 `json:"plans"`
}

// Validate validates this update network wireless billing body
func (o *UpdateNetworkWirelessBillingBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePlans(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessBillingBody) validatePlans(formats strfmt.Registry) error {
	if swag.IsZero(o.Plans) { // not required
		return nil
	}

	for i := 0; i < len(o.Plans); i++ {
		if swag.IsZero(o.Plans[i]) { // not required
			continue
		}

		if o.Plans[i] != nil {
			if err := o.Plans[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessBilling" + "." + "plans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessBilling" + "." + "plans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this update network wireless billing body based on the context it is used
func (o *UpdateNetworkWirelessBillingBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePlans(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessBillingBody) contextValidatePlans(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Plans); i++ {

		if o.Plans[i] != nil {
			if err := o.Plans[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkWirelessBilling" + "." + "plans" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkWirelessBilling" + "." + "plans" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessBillingBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessBillingBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessBillingBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkWirelessBillingParamsBodyPlansItems0 update network wireless billing params body plans items0
swagger:model UpdateNetworkWirelessBillingParamsBodyPlansItems0
*/
type UpdateNetworkWirelessBillingParamsBodyPlansItems0 struct {

	// bandwidth limits
	// Required: true
	BandwidthLimits *UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits `json:"bandwidthLimits"`

	// The id of the pricing plan to update.
	ID string `json:"id,omitempty"`

	// The price of the billing plan.
	// Required: true
	Price *float32 `json:"price"`

	// The time limit of the pricing plan in minutes. Can be '1 hour', '1 day', '1 week', or '30 days'.
	// Required: true
	// Enum: [1 hour 1 day 1 week 30 days]
	TimeLimit *string `json:"timeLimit"`
}

// Validate validates this update network wireless billing params body plans items0
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateBandwidthLimits(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePrice(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTimeLimit(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) validateBandwidthLimits(formats strfmt.Registry) error {

	if err := validate.Required("bandwidthLimits", "body", o.BandwidthLimits); err != nil {
		return err
	}

	if o.BandwidthLimits != nil {
		if err := o.BandwidthLimits.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) validatePrice(formats strfmt.Registry) error {

	if err := validate.Required("price", "body", o.Price); err != nil {
		return err
	}

	return nil
}

var updateNetworkWirelessBillingParamsBodyPlansItems0TypeTimeLimitPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["1 hour","1 day","1 week","30 days"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkWirelessBillingParamsBodyPlansItems0TypeTimeLimitPropEnum = append(updateNetworkWirelessBillingParamsBodyPlansItems0TypeTimeLimitPropEnum, v)
	}
}

const (

	// UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr1Hour captures enum value "1 hour"
	UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr1Hour string = "1 hour"

	// UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr1Day captures enum value "1 day"
	UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr1Day string = "1 day"

	// UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr1Week captures enum value "1 week"
	UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr1Week string = "1 week"

	// UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr30Days captures enum value "30 days"
	UpdateNetworkWirelessBillingParamsBodyPlansItems0TimeLimitNr30Days string = "30 days"
)

// prop value enum
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) validateTimeLimitEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkWirelessBillingParamsBodyPlansItems0TypeTimeLimitPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) validateTimeLimit(formats strfmt.Registry) error {

	if err := validate.Required("timeLimit", "body", o.TimeLimit); err != nil {
		return err
	}

	// value enum
	if err := o.validateTimeLimitEnum("timeLimit", "body", *o.TimeLimit); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update network wireless billing params body plans items0 based on the context it is used
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateBandwidthLimits(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) contextValidateBandwidthLimits(ctx context.Context, formats strfmt.Registry) error {

	if o.BandwidthLimits != nil {
		if err := o.BandwidthLimits.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bandwidthLimits")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("bandwidthLimits")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessBillingParamsBodyPlansItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits The uplink bandwidth settings for the pricing plan.
swagger:model UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits
*/
type UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits struct {

	// The maximum download limit (integer, in Kbps). null indicates no limit
	LimitDown int64 `json:"limitDown,omitempty"`

	// The maximum upload limit (integer, in Kbps). null indicates no limit
	LimitUp int64 `json:"limitUp,omitempty"`
}

// Validate validates this update network wireless billing params body plans items0 bandwidth limits
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network wireless billing params body plans items0 bandwidth limits based on context it is used
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkWirelessBillingParamsBodyPlansItems0BandwidthLimits
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
