// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

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

// CreateNetworkSwitchPortScheduleReader is a Reader for the CreateNetworkSwitchPortSchedule structure.
type CreateNetworkSwitchPortScheduleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSwitchPortScheduleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSwitchPortScheduleCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkSwitchPortScheduleCreated creates a CreateNetworkSwitchPortScheduleCreated with default headers values
func NewCreateNetworkSwitchPortScheduleCreated() *CreateNetworkSwitchPortScheduleCreated {
	return &CreateNetworkSwitchPortScheduleCreated{}
}

/* CreateNetworkSwitchPortScheduleCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSwitchPortScheduleCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network switch port schedule created response has a 2xx status code
func (o *CreateNetworkSwitchPortScheduleCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network switch port schedule created response has a 3xx status code
func (o *CreateNetworkSwitchPortScheduleCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network switch port schedule created response has a 4xx status code
func (o *CreateNetworkSwitchPortScheduleCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network switch port schedule created response has a 5xx status code
func (o *CreateNetworkSwitchPortScheduleCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network switch port schedule created response a status code equal to that given
func (o *CreateNetworkSwitchPortScheduleCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateNetworkSwitchPortScheduleCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/portSchedules][%d] createNetworkSwitchPortScheduleCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchPortScheduleCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/switch/portSchedules][%d] createNetworkSwitchPortScheduleCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSwitchPortScheduleCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSwitchPortScheduleCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateNetworkSwitchPortScheduleBody create network switch port schedule body
// Example: {"name":"Weekdays schedule","portSchedule":{"friday":{"active":true,"from":"9:00","to":"17:00"},"monday":{"active":true,"from":"9:00","to":"17:00"},"saturday":{"active":false,"from":"0:00","to":"24:00"},"sunday":{"active":false,"from":"0:00","to":"24:00"},"thursday":{"active":true,"from":"9:00","to":"17:00"},"tuesday":{"active":true,"from":"9:00","to":"17:00"},"wednesday":{"active":true,"from":"9:00","to":"17:00"}}}
swagger:model CreateNetworkSwitchPortScheduleBody
*/
type CreateNetworkSwitchPortScheduleBody struct {

	// The name for your port schedule. Required
	// Required: true
	Name *string `json:"name"`

	// port schedule
	PortSchedule *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule `json:"portSchedule,omitempty"`
}

// Validate validates this create network switch port schedule body
func (o *CreateNetworkSwitchPortScheduleBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePortSchedule(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchPortScheduleBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkSwitchPortSchedule"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleBody) validatePortSchedule(formats strfmt.Registry) error {
	if swag.IsZero(o.PortSchedule) { // not required
		return nil
	}

	if o.PortSchedule != nil {
		if err := o.PortSchedule.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network switch port schedule body based on the context it is used
func (o *CreateNetworkSwitchPortScheduleBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePortSchedule(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchPortScheduleBody) contextValidatePortSchedule(ctx context.Context, formats strfmt.Registry) error {

	if o.PortSchedule != nil {
		if err := o.PortSchedule.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortSchedule     The schedule for switch port scheduling. Schedules are applied to days of the week.
//     When it's empty, default schedule with all days of a week are configured.
//     Any unspecified day in the schedule is added as a default schedule configuration of the day.
//
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortSchedule
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortSchedule struct {

	// friday
	Friday *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday `json:"friday,omitempty"`

	// monday
	Monday *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday `json:"monday,omitempty"`

	// saturday
	Saturday *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday `json:"saturday,omitempty"`

	// sunday
	Sunday *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday `json:"sunday,omitempty"`

	// thursday
	Thursday *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday `json:"thursday,omitempty"`

	// tuesday
	Tuesday *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday `json:"tuesday,omitempty"`

	// wednesday
	Wednesday *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday `json:"wednesday,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateFriday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMonday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSaturday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSunday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateThursday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateTuesday(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateWednesday(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateFriday(formats strfmt.Registry) error {
	if swag.IsZero(o.Friday) { // not required
		return nil
	}

	if o.Friday != nil {
		if err := o.Friday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateMonday(formats strfmt.Registry) error {
	if swag.IsZero(o.Monday) { // not required
		return nil
	}

	if o.Monday != nil {
		if err := o.Monday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateSaturday(formats strfmt.Registry) error {
	if swag.IsZero(o.Saturday) { // not required
		return nil
	}

	if o.Saturday != nil {
		if err := o.Saturday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateSunday(formats strfmt.Registry) error {
	if swag.IsZero(o.Sunday) { // not required
		return nil
	}

	if o.Sunday != nil {
		if err := o.Sunday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateThursday(formats strfmt.Registry) error {
	if swag.IsZero(o.Thursday) { // not required
		return nil
	}

	if o.Thursday != nil {
		if err := o.Thursday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateTuesday(formats strfmt.Registry) error {
	if swag.IsZero(o.Tuesday) { // not required
		return nil
	}

	if o.Tuesday != nil {
		if err := o.Tuesday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) validateWednesday(formats strfmt.Registry) error {
	if swag.IsZero(o.Wednesday) { // not required
		return nil
	}

	if o.Wednesday != nil {
		if err := o.Wednesday.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network switch port schedule params body port schedule based on the context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateFriday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMonday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSaturday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateSunday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateThursday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateTuesday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateWednesday(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateFriday(ctx context.Context, formats strfmt.Registry) error {

	if o.Friday != nil {
		if err := o.Friday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "friday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateMonday(ctx context.Context, formats strfmt.Registry) error {

	if o.Monday != nil {
		if err := o.Monday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "monday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateSaturday(ctx context.Context, formats strfmt.Registry) error {

	if o.Saturday != nil {
		if err := o.Saturday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "saturday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateSunday(ctx context.Context, formats strfmt.Registry) error {

	if o.Sunday != nil {
		if err := o.Sunday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "sunday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateThursday(ctx context.Context, formats strfmt.Registry) error {

	if o.Thursday != nil {
		if err := o.Thursday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "thursday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateTuesday(ctx context.Context, formats strfmt.Registry) error {

	if o.Tuesday != nil {
		if err := o.Tuesday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "tuesday")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) contextValidateWednesday(ctx context.Context, formats strfmt.Registry) error {

	if o.Wednesday != nil {
		if err := o.Wednesday.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkSwitchPortSchedule" + "." + "portSchedule" + "." + "wednesday")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortSchedule) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortSchedule
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday The schedule object for Friday.
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule friday
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch port schedule params body port schedule friday based on context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortScheduleFriday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday The schedule object for Monday.
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule monday
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch port schedule params body port schedule monday based on context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortScheduleMonday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday The schedule object for Saturday.
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule saturday
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch port schedule params body port schedule saturday based on context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSaturday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday The schedule object for Sunday.
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule sunday
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch port schedule params body port schedule sunday based on context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortScheduleSunday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday The schedule object for Thursday.
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule thursday
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch port schedule params body port schedule thursday based on context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortScheduleThursday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday The schedule object for Tuesday.
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule tuesday
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch port schedule params body port schedule tuesday based on context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortScheduleTuesday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday The schedule object for Wednesday.
swagger:model CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday
*/
type CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday struct {

	// Whether the schedule is active (true) or inactive (false) during the time specified between 'from' and 'to'. Defaults to true.
	Active bool `json:"active,omitempty"`

	// The time, from '00:00' to '24:00'. Must be less than the time specified in 'to'. Defaults to '00:00'. Only 30 minute increments are allowed.
	From string `json:"from,omitempty"`

	// The time, from '00:00' to '24:00'. Must be greater than the time specified in 'from'. Defaults to '24:00'. Only 30 minute increments are allowed.
	To string `json:"to,omitempty"`
}

// Validate validates this create network switch port schedule params body port schedule wednesday
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network switch port schedule params body port schedule wednesday based on context it is used
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSwitchPortScheduleParamsBodyPortScheduleWednesday
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}