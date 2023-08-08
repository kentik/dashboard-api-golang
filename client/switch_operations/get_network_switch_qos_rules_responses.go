// Code generated by go-swagger; DO NOT EDIT.

package switch_operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetNetworkSwitchQosRulesReader is a Reader for the GetNetworkSwitchQosRules structure.
type GetNetworkSwitchQosRulesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSwitchQosRulesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSwitchQosRulesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/switch/qosRules] getNetworkSwitchQosRules", response, response.Code())
	}
}

// NewGetNetworkSwitchQosRulesOK creates a GetNetworkSwitchQosRulesOK with default headers values
func NewGetNetworkSwitchQosRulesOK() *GetNetworkSwitchQosRulesOK {
	return &GetNetworkSwitchQosRulesOK{}
}

/*
GetNetworkSwitchQosRulesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSwitchQosRulesOK struct {
	Payload []interface{}
}

// IsSuccess returns true when this get network switch qos rules o k response has a 2xx status code
func (o *GetNetworkSwitchQosRulesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network switch qos rules o k response has a 3xx status code
func (o *GetNetworkSwitchQosRulesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network switch qos rules o k response has a 4xx status code
func (o *GetNetworkSwitchQosRulesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network switch qos rules o k response has a 5xx status code
func (o *GetNetworkSwitchQosRulesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network switch qos rules o k response a status code equal to that given
func (o *GetNetworkSwitchQosRulesOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network switch qos rules o k response
func (o *GetNetworkSwitchQosRulesOK) Code() int {
	return 200
}

func (o *GetNetworkSwitchQosRulesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/qosRules][%d] getNetworkSwitchQosRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchQosRulesOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/switch/qosRules][%d] getNetworkSwitchQosRulesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSwitchQosRulesOK) GetPayload() []interface{} {
	return o.Payload
}

func (o *GetNetworkSwitchQosRulesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
