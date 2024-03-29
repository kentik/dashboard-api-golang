// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetNetworkSensorMqttBrokerReader is a Reader for the GetNetworkSensorMqttBroker structure.
type GetNetworkSensorMqttBrokerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSensorMqttBrokerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSensorMqttBrokerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId}] getNetworkSensorMqttBroker", response, response.Code())
	}
}

// NewGetNetworkSensorMqttBrokerOK creates a GetNetworkSensorMqttBrokerOK with default headers values
func NewGetNetworkSensorMqttBrokerOK() *GetNetworkSensorMqttBrokerOK {
	return &GetNetworkSensorMqttBrokerOK{}
}

/*
GetNetworkSensorMqttBrokerOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSensorMqttBrokerOK struct {
	Payload *GetNetworkSensorMqttBrokerOKBody
}

// IsSuccess returns true when this get network sensor mqtt broker o k response has a 2xx status code
func (o *GetNetworkSensorMqttBrokerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sensor mqtt broker o k response has a 3xx status code
func (o *GetNetworkSensorMqttBrokerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sensor mqtt broker o k response has a 4xx status code
func (o *GetNetworkSensorMqttBrokerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sensor mqtt broker o k response has a 5xx status code
func (o *GetNetworkSensorMqttBrokerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sensor mqtt broker o k response a status code equal to that given
func (o *GetNetworkSensorMqttBrokerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sensor mqtt broker o k response
func (o *GetNetworkSensorMqttBrokerOK) Code() int {
	return 200
}

func (o *GetNetworkSensorMqttBrokerOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId}][%d] getNetworkSensorMqttBrokerOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSensorMqttBrokerOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sensor/mqttBrokers/{mqttBrokerId}][%d] getNetworkSensorMqttBrokerOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSensorMqttBrokerOK) GetPayload() *GetNetworkSensorMqttBrokerOKBody {
	return o.Payload
}

func (o *GetNetworkSensorMqttBrokerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetNetworkSensorMqttBrokerOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSensorMqttBrokerOKBody get network sensor mqtt broker o k body
swagger:model GetNetworkSensorMqttBrokerOKBody
*/
type GetNetworkSensorMqttBrokerOKBody struct {

	// Specifies whether the broker is enabled for sensor data. Currently, only a single broker may be enabled for sensor data.
	Enabled bool `json:"enabled,omitempty"`

	// ID of the MQTT Broker.
	MqttBrokerID string `json:"mqttBrokerId,omitempty"`
}

// Validate validates this get network sensor mqtt broker o k body
func (o *GetNetworkSensorMqttBrokerOKBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network sensor mqtt broker o k body based on context it is used
func (o *GetNetworkSensorMqttBrokerOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorMqttBrokerOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorMqttBrokerOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorMqttBrokerOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
