// Code generated by go-swagger; DO NOT EDIT.

package appliance

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

// CreateNetworkApplianceTrafficShapingCustomPerformanceClassReader is a Reader for the CreateNetworkApplianceTrafficShapingCustomPerformanceClass structure.
type CreateNetworkApplianceTrafficShapingCustomPerformanceClassReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated creates a CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated with default headers values
func NewCreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated() *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated {
	return &CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated{}
}

/* CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network appliance traffic shaping custom performance class created response has a 2xx status code
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network appliance traffic shaping custom performance class created response has a 3xx status code
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network appliance traffic shaping custom performance class created response has a 4xx status code
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network appliance traffic shaping custom performance class created response has a 5xx status code
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network appliance traffic shaping custom performance class created response a status code equal to that given
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) IsCode(code int) bool {
	return code == 201
}

func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses][%d] createNetworkApplianceTrafficShapingCustomPerformanceClassCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/appliance/trafficShaping/customPerformanceClasses][%d] createNetworkApplianceTrafficShapingCustomPerformanceClassCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody create network appliance traffic shaping custom performance class body
// Example: {"maxJitter":100,"maxLatency":100,"maxLossPercentage":5,"name":"myCustomPerformanceClass"}
swagger:model CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody
*/
type CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody struct {

	// Maximum jitter in milliseconds
	MaxJitter int64 `json:"maxJitter,omitempty"`

	// Maximum latency in milliseconds
	MaxLatency int64 `json:"maxLatency,omitempty"`

	// Maximum percentage of packet loss
	MaxLossPercentage int64 `json:"maxLossPercentage,omitempty"`

	// Name of the custom performance class
	// Required: true
	Name *string `json:"name"`
}

// Validate validates this create network appliance traffic shaping custom performance class body
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody) validateName(formats strfmt.Registry) error {

	if err := validate.Required("createNetworkApplianceTrafficShapingCustomPerformanceClass"+"."+"name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network appliance traffic shaping custom performance class body based on context it is used
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkApplianceTrafficShapingCustomPerformanceClassBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
