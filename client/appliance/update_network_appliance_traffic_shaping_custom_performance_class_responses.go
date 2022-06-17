// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// UpdateNetworkApplianceTrafficShapingCustomPerformanceClassReader is a Reader for the UpdateNetworkApplianceTrafficShapingCustomPerformanceClass structure.
type UpdateNetworkApplianceTrafficShapingCustomPerformanceClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK creates a UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK with default headers values
func NewUpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK() *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK {
	return &UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK{}
}

/* UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network appliance traffic shaping custom performance class o k response has a 2xx status code
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network appliance traffic shaping custom performance class o k response has a 3xx status code
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network appliance traffic shaping custom performance class o k response has a 4xx status code
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network appliance traffic shaping custom performance class o k response has a 5xx status code
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network appliance traffic shaping custom performance class o k response a status code equal to that given
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}][%d] updateNetworkApplianceTrafficShapingCustomPerformanceClassOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses/{customPerformanceClassId}][%d] updateNetworkApplianceTrafficShapingCustomPerformanceClassOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody update network appliance traffic shaping custom performance class body
// Example: {"maxJitter":100,"maxLatency":100,"maxLossPercentage":5,"name":"myCustomPerformanceClass"}
swagger:model UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody
*/
type UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody struct {

	// Maximum jitter in milliseconds
	MaxJitter int64 `json:"maxJitter,omitempty"`

	// Maximum latency in milliseconds
	MaxLatency int64 `json:"maxLatency,omitempty"`

	// Maximum percentage of packet loss
	MaxLossPercentage int64 `json:"maxLossPercentage,omitempty"`

	// Name of the custom performance class
	Name string `json:"name,omitempty"`
}

// Validate validates this update network appliance traffic shaping custom performance class body
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this update network appliance traffic shaping custom performance class body based on context it is used
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkApplianceTrafficShapingCustomPerformanceClassBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
