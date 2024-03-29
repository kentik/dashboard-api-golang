// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkAlertsSettingsReader is a Reader for the UpdateNetworkAlertsSettings structure.
type UpdateNetworkAlertsSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkAlertsSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkAlertsSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /networks/{networkId}/alerts/settings] updateNetworkAlertsSettings", response, response.Code())
	}
}

// NewUpdateNetworkAlertsSettingsOK creates a UpdateNetworkAlertsSettingsOK with default headers values
func NewUpdateNetworkAlertsSettingsOK() *UpdateNetworkAlertsSettingsOK {
	return &UpdateNetworkAlertsSettingsOK{}
}

/*
UpdateNetworkAlertsSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkAlertsSettingsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network alerts settings o k response has a 2xx status code
func (o *UpdateNetworkAlertsSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network alerts settings o k response has a 3xx status code
func (o *UpdateNetworkAlertsSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network alerts settings o k response has a 4xx status code
func (o *UpdateNetworkAlertsSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network alerts settings o k response has a 5xx status code
func (o *UpdateNetworkAlertsSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network alerts settings o k response a status code equal to that given
func (o *UpdateNetworkAlertsSettingsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update network alerts settings o k response
func (o *UpdateNetworkAlertsSettingsOK) Code() int {
	return 200
}

func (o *UpdateNetworkAlertsSettingsOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/alerts/settings][%d] updateNetworkAlertsSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkAlertsSettingsOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/alerts/settings][%d] updateNetworkAlertsSettingsOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkAlertsSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkAlertsSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateNetworkAlertsSettingsBody update network alerts settings body
// Example: {"alerts":[{"alertDestinations":{"allAdmins":false,"emails":["miles@meraki.com"],"httpServerIds":["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="],"snmp":false},"enabled":true,"filters":{"timeout":60},"type":"gatewayDown"}],"defaultDestinations":{"allAdmins":true,"emails":["miles@meraki.com"],"httpServerIds":["aHR0cHM6Ly93d3cuZXhhbXBsZS5jb20vd2ViaG9va3M="],"snmp":true}}
swagger:model UpdateNetworkAlertsSettingsBody
*/
type UpdateNetworkAlertsSettingsBody struct {

	// Alert-specific configuration for each type. Only alerts that pertain to the network can be updated.
	Alerts []*UpdateNetworkAlertsSettingsParamsBodyAlertsItems0 `json:"alerts"`

	// default destinations
	DefaultDestinations *UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations `json:"defaultDestinations,omitempty"`

	// muting
	Muting *UpdateNetworkAlertsSettingsParamsBodyMuting `json:"muting,omitempty"`
}

// Validate validates this update network alerts settings body
func (o *UpdateNetworkAlertsSettingsBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlerts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateDefaultDestinations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMuting(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkAlertsSettingsBody) validateAlerts(formats strfmt.Registry) error {
	if swag.IsZero(o.Alerts) { // not required
		return nil
	}

	for i := 0; i < len(o.Alerts); i++ {
		if swag.IsZero(o.Alerts[i]) { // not required
			continue
		}

		if o.Alerts[i] != nil {
			if err := o.Alerts[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkAlertsSettings" + "." + "alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkAlertsSettings" + "." + "alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkAlertsSettingsBody) validateDefaultDestinations(formats strfmt.Registry) error {
	if swag.IsZero(o.DefaultDestinations) { // not required
		return nil
	}

	if o.DefaultDestinations != nil {
		if err := o.DefaultDestinations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkAlertsSettings" + "." + "defaultDestinations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkAlertsSettings" + "." + "defaultDestinations")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkAlertsSettingsBody) validateMuting(formats strfmt.Registry) error {
	if swag.IsZero(o.Muting) { // not required
		return nil
	}

	if o.Muting != nil {
		if err := o.Muting.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkAlertsSettings" + "." + "muting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkAlertsSettings" + "." + "muting")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network alerts settings body based on the context it is used
func (o *UpdateNetworkAlertsSettingsBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAlerts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateDefaultDestinations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateMuting(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkAlertsSettingsBody) contextValidateAlerts(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Alerts); i++ {

		if o.Alerts[i] != nil {

			if swag.IsZero(o.Alerts[i]) { // not required
				return nil
			}

			if err := o.Alerts[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkAlertsSettings" + "." + "alerts" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkAlertsSettings" + "." + "alerts" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *UpdateNetworkAlertsSettingsBody) contextValidateDefaultDestinations(ctx context.Context, formats strfmt.Registry) error {

	if o.DefaultDestinations != nil {

		if swag.IsZero(o.DefaultDestinations) { // not required
			return nil
		}

		if err := o.DefaultDestinations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkAlertsSettings" + "." + "defaultDestinations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkAlertsSettings" + "." + "defaultDestinations")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkAlertsSettingsBody) contextValidateMuting(ctx context.Context, formats strfmt.Registry) error {

	if o.Muting != nil {

		if swag.IsZero(o.Muting) { // not required
			return nil
		}

		if err := o.Muting.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkAlertsSettings" + "." + "muting")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkAlertsSettings" + "." + "muting")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkAlertsSettingsBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkAlertsSettingsParamsBodyAlertsItems0 update network alerts settings params body alerts items0
swagger:model UpdateNetworkAlertsSettingsParamsBodyAlertsItems0
*/
type UpdateNetworkAlertsSettingsParamsBodyAlertsItems0 struct {

	// alert destinations
	AlertDestinations *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations `json:"alertDestinations,omitempty"`

	// A boolean depicting if the alert is turned on or off
	Enabled bool `json:"enabled,omitempty"`

	// A hash of specific configuration data for the alert. Only filters specific to the alert will be updated.
	Filters interface{} `json:"filters,omitempty"`

	// The type of alert
	// Required: true
	Type *string `json:"type"`
}

// Validate validates this update network alerts settings params body alerts items0
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlertDestinations(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0) validateAlertDestinations(formats strfmt.Registry) error {
	if swag.IsZero(o.AlertDestinations) { // not required
		return nil
	}

	if o.AlertDestinations != nil {
		if err := o.AlertDestinations.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertDestinations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertDestinations")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update network alerts settings params body alerts items0 based on the context it is used
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAlertDestinations(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0) contextValidateAlertDestinations(ctx context.Context, formats strfmt.Registry) error {

	if o.AlertDestinations != nil {

		if swag.IsZero(o.AlertDestinations) { // not required
			return nil
		}

		if err := o.AlertDestinations.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("alertDestinations")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("alertDestinations")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkAlertsSettingsParamsBodyAlertsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations A hash of destinations for this specific alert
swagger:model UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations
*/
type UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations struct {

	// If true, then all network admins will receive emails for this alert
	AllAdmins bool `json:"allAdmins,omitempty"`

	// A list of emails that will receive information about the alert
	Emails []string `json:"emails"`

	// A list of HTTP server IDs to send a Webhook to for this alert
	HTTPServerIds []string `json:"httpServerIds"`

	// If true, then an SNMP trap will be sent for this alert if there is an SNMP trap server configured for this network
	Snmp bool `json:"snmp,omitempty"`
}

// Validate validates this update network alerts settings params body alerts items0 alert destinations
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network alerts settings params body alerts items0 alert destinations based on context it is used
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkAlertsSettingsParamsBodyAlertsItems0AlertDestinations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations The network-wide destinations for all alerts on the network.
swagger:model UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations
*/
type UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations struct {

	// If true, then all network admins will receive emails.
	AllAdmins bool `json:"allAdmins,omitempty"`

	// A list of emails that will recieve the alert(s).
	Emails []string `json:"emails"`

	// A list of HTTP server IDs to send a Webhook to
	HTTPServerIds []string `json:"httpServerIds"`

	// If true, then an SNMP trap will be sent if there is an SNMP trap server configured for this network.
	Snmp bool `json:"snmp,omitempty"`
}

// Validate validates this update network alerts settings params body default destinations
func (o *UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network alerts settings params body default destinations based on context it is used
func (o *UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkAlertsSettingsParamsBodyDefaultDestinations
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkAlertsSettingsParamsBodyMuting Mute alerts under certain conditions
swagger:model UpdateNetworkAlertsSettingsParamsBodyMuting
*/
type UpdateNetworkAlertsSettingsParamsBodyMuting struct {

	// by port schedules
	ByPortSchedules *UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules `json:"byPortSchedules,omitempty"`
}

// Validate validates this update network alerts settings params body muting
func (o *UpdateNetworkAlertsSettingsParamsBodyMuting) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateByPortSchedules(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkAlertsSettingsParamsBodyMuting) validateByPortSchedules(formats strfmt.Registry) error {
	if swag.IsZero(o.ByPortSchedules) { // not required
		return nil
	}

	if o.ByPortSchedules != nil {
		if err := o.ByPortSchedules.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkAlertsSettings" + "." + "muting" + "." + "byPortSchedules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkAlertsSettings" + "." + "muting" + "." + "byPortSchedules")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this update network alerts settings params body muting based on the context it is used
func (o *UpdateNetworkAlertsSettingsParamsBodyMuting) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateByPortSchedules(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkAlertsSettingsParamsBodyMuting) contextValidateByPortSchedules(ctx context.Context, formats strfmt.Registry) error {

	if o.ByPortSchedules != nil {

		if swag.IsZero(o.ByPortSchedules) { // not required
			return nil
		}

		if err := o.ByPortSchedules.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateNetworkAlertsSettings" + "." + "muting" + "." + "byPortSchedules")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateNetworkAlertsSettings" + "." + "muting" + "." + "byPortSchedules")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyMuting) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyMuting) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkAlertsSettingsParamsBodyMuting
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules Mute wireless unreachable alerts based on switch port schedules
swagger:model UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules
*/
type UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules struct {

	// If true, then wireless unreachable alerts will be muted when caused by a port schedule
	Enabled bool `json:"enabled,omitempty"`
}

// Validate validates this update network alerts settings params body muting by port schedules
func (o *UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network alerts settings params body muting by port schedules based on context it is used
func (o *UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkAlertsSettingsParamsBodyMutingByPortSchedules
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
