// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// CreateNetworkWebhooksWebhookTestReader is a Reader for the CreateNetworkWebhooksWebhookTest structure.
type CreateNetworkWebhooksWebhookTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkWebhooksWebhookTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkWebhooksWebhookTestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkWebhooksWebhookTestCreated creates a CreateNetworkWebhooksWebhookTestCreated with default headers values
func NewCreateNetworkWebhooksWebhookTestCreated() *CreateNetworkWebhooksWebhookTestCreated {
	return &CreateNetworkWebhooksWebhookTestCreated{}
}

/* CreateNetworkWebhooksWebhookTestCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkWebhooksWebhookTestCreated struct {
	Payload *CreateNetworkWebhooksWebhookTestCreatedBody
}

// IsSuccess returns true when this create network webhooks webhook test created response has a 2xx status code
func (o *CreateNetworkWebhooksWebhookTestCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network webhooks webhook test created response has a 3xx status code
func (o *CreateNetworkWebhooksWebhookTestCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network webhooks webhook test created response has a 4xx status code
func (o *CreateNetworkWebhooksWebhookTestCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network webhooks webhook test created response has a 5xx status code
func (o *CreateNetworkWebhooksWebhookTestCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network webhooks webhook test created response a status code equal to that given
func (o *CreateNetworkWebhooksWebhookTestCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateNetworkWebhooksWebhookTestCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/webhooks/webhookTests][%d] createNetworkWebhooksWebhookTestCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkWebhooksWebhookTestCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/webhooks/webhookTests][%d] createNetworkWebhooksWebhookTestCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkWebhooksWebhookTestCreated) GetPayload() *CreateNetworkWebhooksWebhookTestCreatedBody {
	return o.Payload
}

func (o *CreateNetworkWebhooksWebhookTestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateNetworkWebhooksWebhookTestCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateNetworkWebhooksWebhookTestBody create network webhooks webhook test body
// Example: {"alertTypeId":"power_supply_down","payloadTemplateId":"wpt_00001","payloadTemplateName":"Payload Template","sharedSecret":"shhh","url":"https://www.example.com/path"}
swagger:model CreateNetworkWebhooksWebhookTestBody
*/
type CreateNetworkWebhooksWebhookTestBody struct {

	// The type of alert which the test webhook will send. Optional. Defaults to power_supply_down.
	AlertTypeID *string `json:"alertTypeId,omitempty"`

	// The ID of the payload template of the test webhook. Defaults to the HTTP server's template ID if one exists for the given URL, or Generic template ID otherwise
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"`

	// The name of the payload template.
	PayloadTemplateName string `json:"payloadTemplateName,omitempty"`

	// The shared secret the test webhook will send. Optional. Defaults to an empty string.
	SharedSecret string `json:"sharedSecret,omitempty"`

	// The URL where the test webhook will be sent
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this create network webhooks webhook test body
func (o *CreateNetworkWebhooksWebhookTestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkWebhooksWebhookTestBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkWebhooksWebhookTest"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network webhooks webhook test body based on context it is used
func (o *CreateNetworkWebhooksWebhookTestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkWebhooksWebhookTestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkWebhooksWebhookTestBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkWebhooksWebhookTestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*CreateNetworkWebhooksWebhookTestCreatedBody create network webhooks webhook test created body
swagger:model CreateNetworkWebhooksWebhookTestCreatedBody
*/
type CreateNetworkWebhooksWebhookTestCreatedBody struct {

	// Webhook delivery identifier
	ID string `json:"id,omitempty"`

	// Current status of the webhook delivery
	// Enum: [abandoned delivered retrying enqueued processing]
	Status string `json:"status,omitempty"`

	// URL where the webhook was delivered
	URL string `json:"url,omitempty"`
}

// Validate validates this create network webhooks webhook test created body
func (o *CreateNetworkWebhooksWebhookTestCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkWebhooksWebhookTestCreatedBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["abandoned","delivered","retrying","enqueued","processing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkWebhooksWebhookTestCreatedBodyTypeStatusPropEnum = append(createNetworkWebhooksWebhookTestCreatedBodyTypeStatusPropEnum, v)
	}
}

const (

	// CreateNetworkWebhooksWebhookTestCreatedBodyStatusAbandoned captures enum value "abandoned"
	CreateNetworkWebhooksWebhookTestCreatedBodyStatusAbandoned string = "abandoned"

	// CreateNetworkWebhooksWebhookTestCreatedBodyStatusDelivered captures enum value "delivered"
	CreateNetworkWebhooksWebhookTestCreatedBodyStatusDelivered string = "delivered"

	// CreateNetworkWebhooksWebhookTestCreatedBodyStatusRetrying captures enum value "retrying"
	CreateNetworkWebhooksWebhookTestCreatedBodyStatusRetrying string = "retrying"

	// CreateNetworkWebhooksWebhookTestCreatedBodyStatusEnqueued captures enum value "enqueued"
	CreateNetworkWebhooksWebhookTestCreatedBodyStatusEnqueued string = "enqueued"

	// CreateNetworkWebhooksWebhookTestCreatedBodyStatusProcessing captures enum value "processing"
	CreateNetworkWebhooksWebhookTestCreatedBodyStatusProcessing string = "processing"
)

// prop value enum
func (o *CreateNetworkWebhooksWebhookTestCreatedBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkWebhooksWebhookTestCreatedBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkWebhooksWebhookTestCreatedBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("createNetworkWebhooksWebhookTestCreated"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network webhooks webhook test created body based on context it is used
func (o *CreateNetworkWebhooksWebhookTestCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkWebhooksWebhookTestCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkWebhooksWebhookTestCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkWebhooksWebhookTestCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
