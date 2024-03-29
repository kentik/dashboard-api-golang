// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// CreateNetworkWebhooksHTTPServerReader is a Reader for the CreateNetworkWebhooksHTTPServer structure.
type CreateNetworkWebhooksHTTPServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkWebhooksHTTPServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkWebhooksHTTPServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/webhooks/httpServers] createNetworkWebhooksHttpServer", response, response.Code())
	}
}

// NewCreateNetworkWebhooksHTTPServerCreated creates a CreateNetworkWebhooksHTTPServerCreated with default headers values
func NewCreateNetworkWebhooksHTTPServerCreated() *CreateNetworkWebhooksHTTPServerCreated {
	return &CreateNetworkWebhooksHTTPServerCreated{}
}

/*
CreateNetworkWebhooksHTTPServerCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkWebhooksHTTPServerCreated struct {
	Payload *CreateNetworkWebhooksHTTPServerCreatedBody
}

// IsSuccess returns true when this create network webhooks Http server created response has a 2xx status code
func (o *CreateNetworkWebhooksHTTPServerCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network webhooks Http server created response has a 3xx status code
func (o *CreateNetworkWebhooksHTTPServerCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network webhooks Http server created response has a 4xx status code
func (o *CreateNetworkWebhooksHTTPServerCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network webhooks Http server created response has a 5xx status code
func (o *CreateNetworkWebhooksHTTPServerCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network webhooks Http server created response a status code equal to that given
func (o *CreateNetworkWebhooksHTTPServerCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network webhooks Http server created response
func (o *CreateNetworkWebhooksHTTPServerCreated) Code() int {
	return 201
}

func (o *CreateNetworkWebhooksHTTPServerCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/webhooks/httpServers][%d] createNetworkWebhooksHttpServerCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkWebhooksHTTPServerCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/webhooks/httpServers][%d] createNetworkWebhooksHttpServerCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkWebhooksHTTPServerCreated) GetPayload() *CreateNetworkWebhooksHTTPServerCreatedBody {
	return o.Payload
}

func (o *CreateNetworkWebhooksHTTPServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateNetworkWebhooksHTTPServerCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkWebhooksHTTPServerBody create network webhooks HTTP server body
// Example: {"name":"Example Webhook Server","payloadTemplate":{"name":"Meraki (included)","payloadTemplateId":"wpt_00001"},"sharedSecret":"shhh","url":"https://example.com"}
swagger:model CreateNetworkWebhooksHTTPServerBody
*/
type CreateNetworkWebhooksHTTPServerBody struct {

	// A name for easy reference to the HTTP server
	// Required: true
	Name *string `json:"name"`

	// payload template
	PayloadTemplate *CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate `json:"payloadTemplate,omitempty"`

	// A shared secret that will be included in POSTs sent to the HTTP server. This secret can be used to verify that the request was sent by Meraki.
	SharedSecret string `json:"sharedSecret,omitempty"`

	// The URL of the HTTP server. Once set, cannot be updated.
	// Required: true
	URL *string `json:"url"`
}

// Validate validates this create network webhooks HTTP server body
func (o *CreateNetworkWebhooksHTTPServerBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validatePayloadTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkWebhooksHTTPServerBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkWebhooksHttpServer"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

func (o *CreateNetworkWebhooksHTTPServerBody) validatePayloadTemplate(formats strfmt.Registry) error {
	if swag.IsZero(o.PayloadTemplate) { // not required
		return nil
	}

	if o.PayloadTemplate != nil {
		if err := o.PayloadTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

func (o *CreateNetworkWebhooksHTTPServerBody) validateURL(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkWebhooksHttpServer"+"."+"url", "body", o.URL); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this create network webhooks HTTP server body based on the context it is used
func (o *CreateNetworkWebhooksHTTPServerBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePayloadTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkWebhooksHTTPServerBody) contextValidatePayloadTemplate(ctx context.Context, formats strfmt.Registry) error {

	if o.PayloadTemplate != nil {

		if swag.IsZero(o.PayloadTemplate) { // not required
			return nil
		}

		if err := o.PayloadTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkWebhooksHttpServer" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkWebhooksHTTPServerBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkWebhooksHTTPServerCreatedBody create network webhooks HTTP server created body
swagger:model CreateNetworkWebhooksHTTPServerCreatedBody
*/
type CreateNetworkWebhooksHTTPServerCreatedBody struct {

	// A Base64 encoded ID.
	// Format: byte
	ID strfmt.Base64 `json:"id,omitempty"`

	// A name for easy reference to the HTTP server
	Name string `json:"name,omitempty"`

	// A Meraki network ID.
	NetworkID string `json:"networkId,omitempty"`

	// payload template
	PayloadTemplate *CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate `json:"payloadTemplate,omitempty"`

	// The URL of the HTTP server.
	URL string `json:"url,omitempty"`
}

// Validate validates this create network webhooks HTTP server created body
func (o *CreateNetworkWebhooksHTTPServerCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePayloadTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkWebhooksHTTPServerCreatedBody) validatePayloadTemplate(formats strfmt.Registry) error {
	if swag.IsZero(o.PayloadTemplate) { // not required
		return nil
	}

	if o.PayloadTemplate != nil {
		if err := o.PayloadTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkWebhooksHttpServerCreated" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkWebhooksHttpServerCreated" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create network webhooks HTTP server created body based on the context it is used
func (o *CreateNetworkWebhooksHTTPServerCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePayloadTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkWebhooksHTTPServerCreatedBody) contextValidatePayloadTemplate(ctx context.Context, formats strfmt.Registry) error {

	if o.PayloadTemplate != nil {

		if swag.IsZero(o.PayloadTemplate) { // not required
			return nil
		}

		if err := o.PayloadTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createNetworkWebhooksHttpServerCreated" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createNetworkWebhooksHttpServerCreated" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkWebhooksHTTPServerCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate The payload template to use when posting data to the HTTP server.
swagger:model CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate
*/
type CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate struct {

	// The name of the payload template.
	Name string `json:"name,omitempty"`

	// The ID of the payload template.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"`
}

// Validate validates this create network webhooks HTTP server created body payload template
func (o *CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network webhooks HTTP server created body payload template based on context it is used
func (o *CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate) UnmarshalBinary(b []byte) error {
	var res CreateNetworkWebhooksHTTPServerCreatedBodyPayloadTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate The payload template to use when posting data to the HTTP server.
swagger:model CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate
*/
type CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate struct {

	// The name of the payload template.
	Name string `json:"name,omitempty"`

	// The ID of the payload template. Defaults to 'wpt_00001' for the Meraki template. For Meraki-included templates: for the Webex (included) template use 'wpt_00002'; for the Slack (included) template use 'wpt_00003'; for the Microsoft Teams (included) template use 'wpt_00004'; for the ServiceNow (included) template use 'wpt_00006'
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"`
}

// Validate validates this create network webhooks HTTP server params body payload template
func (o *CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network webhooks HTTP server params body payload template based on context it is used
func (o *CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate) UnmarshalBinary(b []byte) error {
	var res CreateNetworkWebhooksHTTPServerParamsBodyPayloadTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
