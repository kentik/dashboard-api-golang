// Code generated by go-swagger; DO NOT EDIT.

package sm

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

// CreateNetworkSmTargetGroupReader is a Reader for the CreateNetworkSmTargetGroup structure.
type CreateNetworkSmTargetGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkSmTargetGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkSmTargetGroupCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/sm/targetGroups] createNetworkSmTargetGroup", response, response.Code())
	}
}

// NewCreateNetworkSmTargetGroupCreated creates a CreateNetworkSmTargetGroupCreated with default headers values
func NewCreateNetworkSmTargetGroupCreated() *CreateNetworkSmTargetGroupCreated {
	return &CreateNetworkSmTargetGroupCreated{}
}

/*
CreateNetworkSmTargetGroupCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkSmTargetGroupCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network sm target group created response has a 2xx status code
func (o *CreateNetworkSmTargetGroupCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network sm target group created response has a 3xx status code
func (o *CreateNetworkSmTargetGroupCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network sm target group created response has a 4xx status code
func (o *CreateNetworkSmTargetGroupCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network sm target group created response has a 5xx status code
func (o *CreateNetworkSmTargetGroupCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network sm target group created response a status code equal to that given
func (o *CreateNetworkSmTargetGroupCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network sm target group created response
func (o *CreateNetworkSmTargetGroupCreated) Code() int {
	return 201
}

func (o *CreateNetworkSmTargetGroupCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/targetGroups][%d] createNetworkSmTargetGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSmTargetGroupCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/sm/targetGroups][%d] createNetworkSmTargetGroupCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkSmTargetGroupCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkSmTargetGroupCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkSmTargetGroupBody create network sm target group body
// Example: {"name":"My target group","scope":"none","tags":"[]","type":"devices"}
swagger:model CreateNetworkSmTargetGroupBody
*/
type CreateNetworkSmTargetGroupBody struct {

	// The name of this target group
	Name string `json:"name,omitempty"`

	// The scope and tag options of the target group. Comma separated values beginning with one of withAny, withAll, withoutAny, withoutAll, all, none, followed by tags. Default to none if empty.
	Scope string `json:"scope,omitempty"`
}

// Validate validates this create network sm target group body
func (o *CreateNetworkSmTargetGroupBody) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create network sm target group body based on context it is used
func (o *CreateNetworkSmTargetGroupBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkSmTargetGroupBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkSmTargetGroupBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkSmTargetGroupBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
