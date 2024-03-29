// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkMqttBrokerReader is a Reader for the GetNetworkMqttBroker structure.
type GetNetworkMqttBrokerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkMqttBrokerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkMqttBrokerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/mqttBrokers/{mqttBrokerId}] getNetworkMqttBroker", response, response.Code())
	}
}

// NewGetNetworkMqttBrokerOK creates a GetNetworkMqttBrokerOK with default headers values
func NewGetNetworkMqttBrokerOK() *GetNetworkMqttBrokerOK {
	return &GetNetworkMqttBrokerOK{}
}

/*
GetNetworkMqttBrokerOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkMqttBrokerOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get network mqtt broker o k response has a 2xx status code
func (o *GetNetworkMqttBrokerOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network mqtt broker o k response has a 3xx status code
func (o *GetNetworkMqttBrokerOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network mqtt broker o k response has a 4xx status code
func (o *GetNetworkMqttBrokerOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network mqtt broker o k response has a 5xx status code
func (o *GetNetworkMqttBrokerOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network mqtt broker o k response a status code equal to that given
func (o *GetNetworkMqttBrokerOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network mqtt broker o k response
func (o *GetNetworkMqttBrokerOK) Code() int {
	return 200
}

func (o *GetNetworkMqttBrokerOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/mqttBrokers/{mqttBrokerId}][%d] getNetworkMqttBrokerOK  %+v", 200, o.Payload)
}

func (o *GetNetworkMqttBrokerOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/mqttBrokers/{mqttBrokerId}][%d] getNetworkMqttBrokerOK  %+v", 200, o.Payload)
}

func (o *GetNetworkMqttBrokerOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetNetworkMqttBrokerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
