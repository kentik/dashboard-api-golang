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

// GetNetworkWebhooksWebhookTestReader is a Reader for the GetNetworkWebhooksWebhookTest structure.
type GetNetworkWebhooksWebhookTestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkWebhooksWebhookTestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkWebhooksWebhookTestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/webhooks/webhookTests/{webhookTestId}] getNetworkWebhooksWebhookTest", response, response.Code())
	}
}

// NewGetNetworkWebhooksWebhookTestOK creates a GetNetworkWebhooksWebhookTestOK with default headers values
func NewGetNetworkWebhooksWebhookTestOK() *GetNetworkWebhooksWebhookTestOK {
	return &GetNetworkWebhooksWebhookTestOK{}
}

/*
GetNetworkWebhooksWebhookTestOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkWebhooksWebhookTestOK struct {
	Payload *GetNetworkWebhooksWebhookTestOKBody
}

// IsSuccess returns true when this get network webhooks webhook test o k response has a 2xx status code
func (o *GetNetworkWebhooksWebhookTestOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network webhooks webhook test o k response has a 3xx status code
func (o *GetNetworkWebhooksWebhookTestOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network webhooks webhook test o k response has a 4xx status code
func (o *GetNetworkWebhooksWebhookTestOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network webhooks webhook test o k response has a 5xx status code
func (o *GetNetworkWebhooksWebhookTestOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network webhooks webhook test o k response a status code equal to that given
func (o *GetNetworkWebhooksWebhookTestOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network webhooks webhook test o k response
func (o *GetNetworkWebhooksWebhookTestOK) Code() int {
	return 200
}

func (o *GetNetworkWebhooksWebhookTestOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/webhooks/webhookTests/{webhookTestId}][%d] getNetworkWebhooksWebhookTestOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWebhooksWebhookTestOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/webhooks/webhookTests/{webhookTestId}][%d] getNetworkWebhooksWebhookTestOK  %+v", 200, o.Payload)
}

func (o *GetNetworkWebhooksWebhookTestOK) GetPayload() *GetNetworkWebhooksWebhookTestOKBody {
	return o.Payload
}

func (o *GetNetworkWebhooksWebhookTestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkWebhooksWebhookTestOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkWebhooksWebhookTestOKBody get network webhooks webhook test o k body
swagger:model GetNetworkWebhooksWebhookTestOKBody
*/
type GetNetworkWebhooksWebhookTestOKBody struct {

	// Webhook delivery identifier
	ID string `json:"id,omitempty"`

	// Current status of the webhook delivery
	// Enum: [abandoned delivered enqueued processing retrying]
	Status string `json:"status,omitempty"`

	// URL where the webhook was delivered
	URL string `json:"url,omitempty"`
}

// Validate validates this get network webhooks webhook test o k body
func (o *GetNetworkWebhooksWebhookTestOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var getNetworkWebhooksWebhookTestOKBodyTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["abandoned","delivered","enqueued","processing","retrying"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		getNetworkWebhooksWebhookTestOKBodyTypeStatusPropEnum = append(getNetworkWebhooksWebhookTestOKBodyTypeStatusPropEnum, v)
	}
}

const (

	// GetNetworkWebhooksWebhookTestOKBodyStatusAbandoned captures enum value "abandoned"
	GetNetworkWebhooksWebhookTestOKBodyStatusAbandoned string = "abandoned"

	// GetNetworkWebhooksWebhookTestOKBodyStatusDelivered captures enum value "delivered"
	GetNetworkWebhooksWebhookTestOKBodyStatusDelivered string = "delivered"

	// GetNetworkWebhooksWebhookTestOKBodyStatusEnqueued captures enum value "enqueued"
	GetNetworkWebhooksWebhookTestOKBodyStatusEnqueued string = "enqueued"

	// GetNetworkWebhooksWebhookTestOKBodyStatusProcessing captures enum value "processing"
	GetNetworkWebhooksWebhookTestOKBodyStatusProcessing string = "processing"

	// GetNetworkWebhooksWebhookTestOKBodyStatusRetrying captures enum value "retrying"
	GetNetworkWebhooksWebhookTestOKBodyStatusRetrying string = "retrying"
)

// prop value enum
func (o *GetNetworkWebhooksWebhookTestOKBody) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, getNetworkWebhooksWebhookTestOKBodyTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *GetNetworkWebhooksWebhookTestOKBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("getNetworkWebhooksWebhookTestOK"+"."+"status", "body", o.Status); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network webhooks webhook test o k body based on context it is used
func (o *GetNetworkWebhooksWebhookTestOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkWebhooksWebhookTestOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkWebhooksWebhookTestOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkWebhooksWebhookTestOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
