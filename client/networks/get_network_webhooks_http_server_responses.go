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
)

// GetNetworkWebhooksHTTPServerReader is a Reader for the GetNetworkWebhooksHTTPServer structure.
type GetNetworkWebhooksHTTPServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWebhooksHTTPServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWebhooksHTTPServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/webhooks/httpServers/{httpServerId}] getNetworkWebhooksHttpServer", response, response.Code())
	}
}

// NewGetNetworkWebhooksHTTPServerOK creates a GetNetworkWebhooksHTTPServerOK with default headers values
func NewGetNetworkWebhooksHTTPServerOK() *GetNetworkWebhooksHTTPServerOK {
	return &GetNetworkWebhooksHTTPServerOK{}
}

/*
GetNetworkWebhooksHTTPServerOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWebhooksHTTPServerOK struct {
	Payload *GetNetworkWebhooksHTTPServerOKBody
}

// IsSuccess returns true when this get network webhooks Http server o k response has a 2xx status code
func (o *GetNetworkWebhooksHTTPServerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network webhooks Http server o k response has a 3xx status code
func (o *GetNetworkWebhooksHTTPServerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network webhooks Http server o k response has a 4xx status code
func (o *GetNetworkWebhooksHTTPServerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network webhooks Http server o k response has a 5xx status code
func (o *GetNetworkWebhooksHTTPServerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network webhooks Http server o k response a status code equal to that given
func (o *GetNetworkWebhooksHTTPServerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network webhooks Http server o k response
func (o *GetNetworkWebhooksHTTPServerOK) Code() int {
	return 200
}

func (o *GetNetworkWebhooksHTTPServerOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/webhooks/httpServers/{httpServerId}][%d] getNetworkWebhooksHttpServerOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWebhooksHTTPServerOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/webhooks/httpServers/{httpServerId}][%d] getNetworkWebhooksHttpServerOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWebhooksHTTPServerOK) GetPayload() *GetNetworkWebhooksHTTPServerOKBody {
	return o.Payload
}

func (o *GetNetworkWebhooksHTTPServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkWebhooksHTTPServerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWebhooksHTTPServerOKBody get network webhooks HTTP server o k body
swagger:model GetNetworkWebhooksHTTPServerOKBody
*/
type GetNetworkWebhooksHTTPServerOKBody struct {

	// A Base64 encoded ID.
	// Format: byte
	ID strfmt.Base64 `json:"id,omitempty"`

	// A name for easy reference to the HTTP server
	Name string `json:"name,omitempty"`

	// A Meraki network ID.
	NetworkID string `json:"networkId,omitempty"`

	// payload template
	PayloadTemplate *GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate `json:"payloadTemplate,omitempty"`

	// The URL of the HTTP server.
	URL string `json:"url,omitempty"`
}

// Validate validates this get network webhooks HTTP server o k body
func (o *GetNetworkWebhooksHTTPServerOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validatePayloadTemplate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWebhooksHTTPServerOKBody) validatePayloadTemplate(formats strfmt.Registry) error {
	if swag.IsZero(o.PayloadTemplate) { // not required
		return nil
	}

	if o.PayloadTemplate != nil {
		if err := o.PayloadTemplate.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWebhooksHttpServerOK" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWebhooksHttpServerOK" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network webhooks HTTP server o k body based on the context it is used
func (o *GetNetworkWebhooksHTTPServerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidatePayloadTemplate(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkWebhooksHTTPServerOKBody) contextValidatePayloadTemplate(ctx context.Context, formats strfmt.Registry) error {

	if o.PayloadTemplate != nil {

		if swag.IsZero(o.PayloadTemplate) { // not required
			return nil
		}

		if err := o.PayloadTemplate.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getNetworkWebhooksHttpServerOK" + "." + "payloadTemplate")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getNetworkWebhooksHttpServerOK" + "." + "payloadTemplate")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWebhooksHTTPServerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWebhooksHTTPServerOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkWebhooksHTTPServerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate The payload template to use when posting data to the HTTP server.
swagger:model GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate
*/
type GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate struct {

	// The name of the payload template.
	Name string `json:"name,omitempty"`

	// The ID of the payload template.
	PayloadTemplateID string `json:"payloadTemplateId,omitempty"`
}

// Validate validates this get network webhooks HTTP server o k body payload template
func (o *GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network webhooks HTTP server o k body payload template based on context it is used
func (o *GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate) UnmarshalBinary(b []byte) error {
	var res GetNetworkWebhooksHTTPServerOKBodyPayloadTemplate
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
