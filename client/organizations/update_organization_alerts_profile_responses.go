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

// UpdateOrganizationAlertsProfileReader is a Reader for the UpdateOrganizationAlertsProfile structure.
type UpdateOrganizationAlertsProfileReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateOrganizationAlertsProfileReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateOrganizationAlertsProfileOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[PUT /organizations/{organizationId}/alerts/profiles/{alertConfigId}] updateOrganizationAlertsProfile", response, response.Code())
	}
}

// NewUpdateOrganizationAlertsProfileOK creates a UpdateOrganizationAlertsProfileOK with default headers values
func NewUpdateOrganizationAlertsProfileOK() *UpdateOrganizationAlertsProfileOK {
	return &UpdateOrganizationAlertsProfileOK{}
}

/*
UpdateOrganizationAlertsProfileOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateOrganizationAlertsProfileOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update organization alerts profile o k response has a 2xx status code
func (o *UpdateOrganizationAlertsProfileOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update organization alerts profile o k response has a 3xx status code
func (o *UpdateOrganizationAlertsProfileOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update organization alerts profile o k response has a 4xx status code
func (o *UpdateOrganizationAlertsProfileOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update organization alerts profile o k response has a 5xx status code
func (o *UpdateOrganizationAlertsProfileOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update organization alerts profile o k response a status code equal to that given
func (o *UpdateOrganizationAlertsProfileOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the update organization alerts profile o k response
func (o *UpdateOrganizationAlertsProfileOK) Code() int {
	return 200
}

func (o *UpdateOrganizationAlertsProfileOK) Error() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/alerts/profiles/{alertConfigId}][%d] updateOrganizationAlertsProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationAlertsProfileOK) String() string {
	return fmt.Sprintf("[PUT /organizations/{organizationId}/alerts/profiles/{alertConfigId}][%d] updateOrganizationAlertsProfileOK  %+v", 200, o.Payload)
}

func (o *UpdateOrganizationAlertsProfileOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateOrganizationAlertsProfileOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
UpdateOrganizationAlertsProfileBody update organization alerts profile body
// Example: {"enabled":true}
swagger:model UpdateOrganizationAlertsProfileBody
*/
type UpdateOrganizationAlertsProfileBody struct {

	// alert condition
	AlertCondition *UpdateOrganizationAlertsProfileParamsBodyAlertCondition `json:"alertCondition,omitempty"`

	// User supplied description of the alert
	Description string `json:"description,omitempty"`

	// Is the alert config enabled
	Enabled bool `json:"enabled,omitempty"`

	// Networks with these tags will be monitored for the alert
	NetworkTags []string `json:"networkTags"`

	// recipients
	Recipients *UpdateOrganizationAlertsProfileParamsBodyRecipients `json:"recipients,omitempty"`

	// The alert type
	// Enum: [appOutage voipJitter voipMos voipPacketLoss wanLatency wanPacketLoss wanStatus wanUtilization]
	Type string `json:"type,omitempty"`
}

// Validate validates this update organization alerts profile body
func (o *UpdateOrganizationAlertsProfileBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAlertCondition(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateRecipients(formats); err != nil {
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

func (o *UpdateOrganizationAlertsProfileBody) validateAlertCondition(formats strfmt.Registry) error {
	if swag.IsZero(o.AlertCondition) { // not required
		return nil
	}

	if o.AlertCondition != nil {
		if err := o.AlertCondition.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAlertsProfile" + "." + "alertCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAlertsProfile" + "." + "alertCondition")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateOrganizationAlertsProfileBody) validateRecipients(formats strfmt.Registry) error {
	if swag.IsZero(o.Recipients) { // not required
		return nil
	}

	if o.Recipients != nil {
		if err := o.Recipients.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAlertsProfile" + "." + "recipients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAlertsProfile" + "." + "recipients")
			}
			return err
		}
	}

	return nil
}

var updateOrganizationAlertsProfileBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["appOutage","voipJitter","voipMos","voipPacketLoss","wanLatency","wanPacketLoss","wanStatus","wanUtilization"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateOrganizationAlertsProfileBodyTypeTypePropEnum = append(updateOrganizationAlertsProfileBodyTypeTypePropEnum, v)
	}
}

const (

	// UpdateOrganizationAlertsProfileBodyTypeAppOutage captures enum value "appOutage"
	UpdateOrganizationAlertsProfileBodyTypeAppOutage string = "appOutage"

	// UpdateOrganizationAlertsProfileBodyTypeVoipJitter captures enum value "voipJitter"
	UpdateOrganizationAlertsProfileBodyTypeVoipJitter string = "voipJitter"

	// UpdateOrganizationAlertsProfileBodyTypeVoipMos captures enum value "voipMos"
	UpdateOrganizationAlertsProfileBodyTypeVoipMos string = "voipMos"

	// UpdateOrganizationAlertsProfileBodyTypeVoipPacketLoss captures enum value "voipPacketLoss"
	UpdateOrganizationAlertsProfileBodyTypeVoipPacketLoss string = "voipPacketLoss"

	// UpdateOrganizationAlertsProfileBodyTypeWanLatency captures enum value "wanLatency"
	UpdateOrganizationAlertsProfileBodyTypeWanLatency string = "wanLatency"

	// UpdateOrganizationAlertsProfileBodyTypeWanPacketLoss captures enum value "wanPacketLoss"
	UpdateOrganizationAlertsProfileBodyTypeWanPacketLoss string = "wanPacketLoss"

	// UpdateOrganizationAlertsProfileBodyTypeWanStatus captures enum value "wanStatus"
	UpdateOrganizationAlertsProfileBodyTypeWanStatus string = "wanStatus"

	// UpdateOrganizationAlertsProfileBodyTypeWanUtilization captures enum value "wanUtilization"
	UpdateOrganizationAlertsProfileBodyTypeWanUtilization string = "wanUtilization"
)

// prop value enum
func (o *UpdateOrganizationAlertsProfileBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateOrganizationAlertsProfileBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateOrganizationAlertsProfileBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("updateOrganizationAlertsProfile"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update organization alerts profile body based on the context it is used
func (o *UpdateOrganizationAlertsProfileBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateAlertCondition(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateRecipients(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateOrganizationAlertsProfileBody) contextValidateAlertCondition(ctx context.Context, formats strfmt.Registry) error {

	if o.AlertCondition != nil {

		if swag.IsZero(o.AlertCondition) { // not required
			return nil
		}

		if err := o.AlertCondition.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAlertsProfile" + "." + "alertCondition")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAlertsProfile" + "." + "alertCondition")
			}
			return err
		}
	}

	return nil
}

func (o *UpdateOrganizationAlertsProfileBody) contextValidateRecipients(ctx context.Context, formats strfmt.Registry) error {

	if o.Recipients != nil {

		if swag.IsZero(o.Recipients) { // not required
			return nil
		}

		if err := o.Recipients.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updateOrganizationAlertsProfile" + "." + "recipients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("updateOrganizationAlertsProfile" + "." + "recipients")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAlertsProfileBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAlertsProfileBody) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAlertsProfileBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateOrganizationAlertsProfileParamsBodyAlertCondition The conditions that determine if the alert triggers
swagger:model UpdateOrganizationAlertsProfileParamsBodyAlertCondition
*/
type UpdateOrganizationAlertsProfileParamsBodyAlertCondition struct {

	// The threshold the metric must cross to be valid for alerting. Used only for WAN Utilization alerts.
	BitRateBps int64 `json:"bit_rate_bps,omitempty"`

	// The total duration in seconds that the threshold should be crossed before alerting
	Duration int64 `json:"duration,omitempty"`

	// The uplink observed for the alert.  interface must be one of the following: wan1, wan2, wan3, cellular
	// Enum: [cellular wan1 wan2 wan3]
	Interface string `json:"interface,omitempty"`

	// The threshold the metric must cross to be valid for alerting. Used only for VoIP Jitter alerts.
	JitterMs int64 `json:"jitter_ms,omitempty"`

	// The threshold the metric must cross to be valid for alerting. Used only for WAN Latency alerts.
	LatencyMs int64 `json:"latency_ms,omitempty"`

	// The threshold the metric must cross to be valid for alerting. Used only for Packet Loss alerts.
	LossRatio float32 `json:"loss_ratio,omitempty"`

	// The threshold the metric must drop below to be valid for alerting. Used only for VoIP MOS alerts.
	Mos float32 `json:"mos,omitempty"`

	// The look back period in seconds for sensing the alert
	Window int64 `json:"window,omitempty"`
}

// Validate validates this update organization alerts profile params body alert condition
func (o *UpdateOrganizationAlertsProfileParamsBodyAlertCondition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateInterface(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var updateOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["cellular","wan1","wan2","wan3"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum = append(updateOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum, v)
	}
}

const (

	// UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceCellular captures enum value "cellular"
	UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceCellular string = "cellular"

	// UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan1 captures enum value "wan1"
	UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan1 string = "wan1"

	// UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan2 captures enum value "wan2"
	UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan2 string = "wan2"

	// UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan3 captures enum value "wan3"
	UpdateOrganizationAlertsProfileParamsBodyAlertConditionInterfaceWan3 string = "wan3"
)

// prop value enum
func (o *UpdateOrganizationAlertsProfileParamsBodyAlertCondition) validateInterfaceEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateOrganizationAlertsProfileParamsBodyAlertConditionTypeInterfacePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateOrganizationAlertsProfileParamsBodyAlertCondition) validateInterface(formats strfmt.Registry) error {
	if swag.IsZero(o.Interface) { // not required
		return nil
	}

	// value enum
	if err := o.validateInterfaceEnum("updateOrganizationAlertsProfile"+"."+"alertCondition"+"."+"interface", "body", o.Interface); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update organization alerts profile params body alert condition based on context it is used
func (o *UpdateOrganizationAlertsProfileParamsBodyAlertCondition) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAlertsProfileParamsBodyAlertCondition) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAlertsProfileParamsBodyAlertCondition) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAlertsProfileParamsBodyAlertCondition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
UpdateOrganizationAlertsProfileParamsBodyRecipients List of recipients that will recieve the alert.
swagger:model UpdateOrganizationAlertsProfileParamsBodyRecipients
*/
type UpdateOrganizationAlertsProfileParamsBodyRecipients struct {

	// A list of emails that will receive information about the alert
	Emails []string `json:"emails"`

	// A list base64 encoded urls of webhook endpoints that will receive information about the alert
	HTTPServerIds []string `json:"httpServerIds"`
}

// Validate validates this update organization alerts profile params body recipients
func (o *UpdateOrganizationAlertsProfileParamsBodyRecipients) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update organization alerts profile params body recipients based on context it is used
func (o *UpdateOrganizationAlertsProfileParamsBodyRecipients) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateOrganizationAlertsProfileParamsBodyRecipients) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateOrganizationAlertsProfileParamsBodyRecipients) UnmarshalBinary(b []byte) error {
	var res UpdateOrganizationAlertsProfileParamsBodyRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
