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

// CreateNetworkPiiRequestReader is a Reader for the CreateNetworkPiiRequest structure.
type CreateNetworkPiiRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateNetworkPiiRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateNetworkPiiRequestCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /networks/{networkId}/pii/requests] createNetworkPiiRequest", response, response.Code())
	}
}

// NewCreateNetworkPiiRequestCreated creates a CreateNetworkPiiRequestCreated with default headers values
func NewCreateNetworkPiiRequestCreated() *CreateNetworkPiiRequestCreated {
	return &CreateNetworkPiiRequestCreated{}
}

/*
CreateNetworkPiiRequestCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateNetworkPiiRequestCreated struct {
	Payload interface{}
}

// IsSuccess returns true when this create network pii request created response has a 2xx status code
func (o *CreateNetworkPiiRequestCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create network pii request created response has a 3xx status code
func (o *CreateNetworkPiiRequestCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create network pii request created response has a 4xx status code
func (o *CreateNetworkPiiRequestCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create network pii request created response has a 5xx status code
func (o *CreateNetworkPiiRequestCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create network pii request created response a status code equal to that given
func (o *CreateNetworkPiiRequestCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create network pii request created response
func (o *CreateNetworkPiiRequestCreated) Code() int {
	return 201
}

func (o *CreateNetworkPiiRequestCreated) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/pii/requests][%d] createNetworkPiiRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkPiiRequestCreated) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/pii/requests][%d] createNetworkPiiRequestCreated  %+v", 201, o.Payload)
}

func (o *CreateNetworkPiiRequestCreated) GetPayload() interface{} {
	return o.Payload
}

func (o *CreateNetworkPiiRequestCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateNetworkPiiRequestBody create network pii request body
// Example: {"datasets":["usage","events"],"mac":"00:77:00:77:00:77","type":"delete"}
swagger:model CreateNetworkPiiRequestBody
*/
type CreateNetworkPiiRequestBody struct {

	// The datasets related to the provided key that should be deleted. Only applies to "delete" requests. The value "all" will be expanded to all datasets applicable to this type. The datasets by applicable to each type are: mac (usage, events, traffic), email (users, loginAttempts), username (users, loginAttempts), bluetoothMac (client, connectivity), smDeviceId (device), smUserId (user)
	Datasets []string `json:"datasets"`

	// The email of a network user account. Only applies to "delete" requests.
	Email string `json:"email,omitempty"`

	// The MAC of a network client device. Applies to both "restrict processing" and "delete" requests.
	Mac string `json:"mac,omitempty"`

	// The sm_device_id of a Systems Manager device. The only way to "restrict processing" or "delete" a Systems Manager device. Must include "device" in the dataset for a "delete" request to destroy the device.
	SmDeviceID string `json:"smDeviceId,omitempty"`

	// The sm_user_id of a Systems Manager user. The only way to "restrict processing" or "delete" a Systems Manager user. Must include "user" in the dataset for a "delete" request to destroy the user.
	SmUserID string `json:"smUserId,omitempty"`

	// One of "delete" or "restrict processing"
	// Enum: [delete restrict processing]
	Type string `json:"type,omitempty"`

	// The username of a network log in. Only applies to "delete" requests.
	Username string `json:"username,omitempty"`
}

// Validate validates this create network pii request body
func (o *CreateNetworkPiiRequestBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var createNetworkPiiRequestBodyTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["delete","restrict processing"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		createNetworkPiiRequestBodyTypeTypePropEnum = append(createNetworkPiiRequestBodyTypeTypePropEnum, v)
	}
}

const (

	// CreateNetworkPiiRequestBodyTypeDelete captures enum value "delete"
	CreateNetworkPiiRequestBodyTypeDelete string = "delete"

	// CreateNetworkPiiRequestBodyTypeRestrictProcessing captures enum value "restrict processing"
	CreateNetworkPiiRequestBodyTypeRestrictProcessing string = "restrict processing"
)

// prop value enum
func (o *CreateNetworkPiiRequestBody) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, createNetworkPiiRequestBodyTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *CreateNetworkPiiRequestBody) validateType(formats strfmt.Registry) error {
	if swag.IsZero(o.Type) { // not required
		return nil
	}

	// value enum
	if err := o.validateTypeEnum("createNetworkPiiRequest"+"."+"type", "body", o.Type); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create network pii request body based on context it is used
func (o *CreateNetworkPiiRequestBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateNetworkPiiRequestBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateNetworkPiiRequestBody) UnmarshalBinary(b []byte) error {
	var res CreateNetworkPiiRequestBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
