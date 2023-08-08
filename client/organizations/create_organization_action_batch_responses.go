// Code generated by go-swagger; DO NOT EDIT.

package organizations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CreateOrganizationActionBatchReader is a Reader for the CreateOrganizationActionBatch structure.
type CreateOrganizationActionBatchReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateOrganizationActionBatchReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateOrganizationActionBatchCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[POST /organizations/{organizationId}/actionBatches] createOrganizationActionBatch", response, response.Code())
	}
}

// NewCreateOrganizationActionBatchCreated creates a CreateOrganizationActionBatchCreated with default headers values
func NewCreateOrganizationActionBatchCreated() *CreateOrganizationActionBatchCreated {
	return &CreateOrganizationActionBatchCreated{}
}

/*
CreateOrganizationActionBatchCreated describes a response with status code 201, with default header values.

Successful operation
*/
type CreateOrganizationActionBatchCreated struct {
	Payload *CreateOrganizationActionBatchCreatedBody
}

// IsSuccess returns true when this create organization action batch created response has a 2xx status code
func (o *CreateOrganizationActionBatchCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this create organization action batch created response has a 3xx status code
func (o *CreateOrganizationActionBatchCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this create organization action batch created response has a 4xx status code
func (o *CreateOrganizationActionBatchCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this create organization action batch created response has a 5xx status code
func (o *CreateOrganizationActionBatchCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this create organization action batch created response a status code equal to that given
func (o *CreateOrganizationActionBatchCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the create organization action batch created response
func (o *CreateOrganizationActionBatchCreated) Code() int {
	return 201
}

func (o *CreateOrganizationActionBatchCreated) Error() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/actionBatches][%d] createOrganizationActionBatchCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationActionBatchCreated) String() string {
	return fmt.Sprintf("[POST /organizations/{organizationId}/actionBatches][%d] createOrganizationActionBatchCreated  %+v", 201, o.Payload)
}

func (o *CreateOrganizationActionBatchCreated) GetPayload() *CreateOrganizationActionBatchCreatedBody {
	return o.Payload
}

func (o *CreateOrganizationActionBatchCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(CreateOrganizationActionBatchCreatedBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
CreateOrganizationActionBatchBody create organization action batch body
// Example: {"actions":[{"operation":"create","resource":"/devices/QXXX-XXXX-XXXX/switch/ports/3"}],"confirmed":true,"synchronous":true}
swagger:model CreateOrganizationActionBatchBody
*/
type CreateOrganizationActionBatchBody struct {

	// A set of changes to make as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	// Required: true
	Actions []*CreateOrganizationActionBatchParamsBodyActionsItems0 `json:"actions"`

	// Set to true for immediate execution. Set to false if the action should be previewed before executing. This property cannot be unset once it is true. Defaults to false.
	Confirmed bool `json:"confirmed,omitempty"`

	// Set to true to force the batch to run synchronous. There can be at most 20 actions in synchronous batch. Defaults to false.
	Synchronous bool `json:"synchronous,omitempty"`
}

// Validate validates this create organization action batch body
func (o *CreateOrganizationActionBatchBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchBody) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationActionBatch"+"."+"actions", "body", o.Actions); err != nil {
		return err
	}

	for i := 0; i < len(o.Actions); i++ {
		if swag.IsZero(o.Actions[i]) { // not required
			continue
		}

		if o.Actions[i] != nil {
			if err := o.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create organization action batch body based on the context it is used
func (o *CreateOrganizationActionBatchBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchBody) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Actions); i++ {

		if o.Actions[i] != nil {

			if swag.IsZero(o.Actions[i]) { // not required
				return nil
			}

			if err := o.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatch" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationActionBatchCreatedBody create organization action batch created body
swagger:model CreateOrganizationActionBatchCreatedBody
*/
type CreateOrganizationActionBatchCreatedBody struct {

	// A set of changes made as part of this action (<a href='https://developer.cisco.com/meraki/api/#/rest/guides/action-batches/'>more details</a>)
	// Required: true
	Actions []*CreateOrganizationActionBatchCreatedBodyActionsItems0 `json:"actions"`

	// Flag describing whether the action should be previewed before executing or not
	Confirmed bool `json:"confirmed,omitempty"`

	// ID of the action batch. Can be used to check the status of the action batch at /organizations/{organizationId}/actionBatches/{actionBatchId}
	ID string `json:"id,omitempty"`

	// ID of the organization this action batch belongs to
	OrganizationID string `json:"organizationId,omitempty"`

	// status
	Status *CreateOrganizationActionBatchCreatedBodyStatus `json:"status,omitempty"`

	// Flag describing whether actions should run synchronously or asynchronously
	Synchronous bool `json:"synchronous,omitempty"`
}

// Validate validates this create organization action batch created body
func (o *CreateOrganizationActionBatchCreatedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateActions(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchCreatedBody) validateActions(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationActionBatchCreated"+"."+"actions", "body", o.Actions); err != nil {
		return err
	}

	for i := 0; i < len(o.Actions); i++ {
		if swag.IsZero(o.Actions[i]) { // not required
			continue
		}

		if o.Actions[i] != nil {
			if err := o.Actions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatchCreated" + "." + "actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatchCreated" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateOrganizationActionBatchCreatedBody) validateStatus(formats strfmt.Registry) error {
	if swag.IsZero(o.Status) { // not required
		return nil
	}

	if o.Status != nil {
		if err := o.Status.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationActionBatchCreated" + "." + "status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationActionBatchCreated" + "." + "status")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this create organization action batch created body based on the context it is used
func (o *CreateOrganizationActionBatchCreatedBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateActions(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateStatus(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchCreatedBody) contextValidateActions(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Actions); i++ {

		if o.Actions[i] != nil {

			if swag.IsZero(o.Actions[i]) { // not required
				return nil
			}

			if err := o.Actions[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatchCreated" + "." + "actions" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatchCreated" + "." + "actions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (o *CreateOrganizationActionBatchCreatedBody) contextValidateStatus(ctx context.Context, formats strfmt.Registry) error {

	if o.Status != nil {

		if swag.IsZero(o.Status) { // not required
			return nil
		}

		if err := o.Status.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createOrganizationActionBatchCreated" + "." + "status")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("createOrganizationActionBatchCreated" + "." + "status")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBody) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchCreatedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationActionBatchCreatedBodyActionsItems0 create organization action batch created body actions items0
swagger:model CreateOrganizationActionBatchCreatedBodyActionsItems0
*/
type CreateOrganizationActionBatchCreatedBodyActionsItems0 struct {

	// The operation to be used by this action
	// Required: true
	Operation *string `json:"operation"`

	// Unique identifier for the resource to be acted on
	// Required: true
	Resource *string `json:"resource"`
}

// Validate validates this create organization action batch created body actions items0
func (o *CreateOrganizationActionBatchCreatedBodyActionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchCreatedBodyActionsItems0) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", o.Operation); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationActionBatchCreatedBodyActionsItems0) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", o.Resource); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization action batch created body actions items0 based on context it is used
func (o *CreateOrganizationActionBatchCreatedBodyActionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBodyActionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBodyActionsItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchCreatedBodyActionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationActionBatchCreatedBodyStatus Status of action batch
swagger:model CreateOrganizationActionBatchCreatedBodyStatus
*/
type CreateOrganizationActionBatchCreatedBodyStatus struct {

	// Flag describing whether all actions in the action batch have completed
	Completed bool `json:"completed,omitempty"`

	// Resources created as a result of this action batch
	// Required: true
	CreatedResources []*CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0 `json:"createdResources"`

	// List of errors encountered when running actions in the action batch
	Errors []string `json:"errors"`

	// Flag describing whether any actions in the action batch failed
	Failed bool `json:"failed,omitempty"`
}

// Validate validates this create organization action batch created body status
func (o *CreateOrganizationActionBatchCreatedBodyStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCreatedResources(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchCreatedBodyStatus) validateCreatedResources(formats strfmt.Registry) error {

	if err := validate.Required("createOrganizationActionBatchCreated"+"."+"status"+"."+"createdResources", "body", o.CreatedResources); err != nil {
		return err
	}

	for i := 0; i < len(o.CreatedResources); i++ {
		if swag.IsZero(o.CreatedResources[i]) { // not required
			continue
		}

		if o.CreatedResources[i] != nil {
			if err := o.CreatedResources[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatchCreated" + "." + "status" + "." + "createdResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatchCreated" + "." + "status" + "." + "createdResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this create organization action batch created body status based on the context it is used
func (o *CreateOrganizationActionBatchCreatedBodyStatus) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCreatedResources(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchCreatedBodyStatus) contextValidateCreatedResources(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.CreatedResources); i++ {

		if o.CreatedResources[i] != nil {

			if swag.IsZero(o.CreatedResources[i]) { // not required
				return nil
			}

			if err := o.CreatedResources[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("createOrganizationActionBatchCreated" + "." + "status" + "." + "createdResources" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("createOrganizationActionBatchCreated" + "." + "status" + "." + "createdResources" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBodyStatus) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBodyStatus) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchCreatedBodyStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0 create organization action batch created body status created resources items0
swagger:model CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0
*/
type CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0 struct {

	// ID of the created resource
	ID string `json:"id,omitempty"`

	// URI, not including base, of the created resource
	URI string `json:"uri,omitempty"`
}

// Validate validates this create organization action batch created body status created resources items0
func (o *CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this create organization action batch created body status created resources items0 based on context it is used
func (o *CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchCreatedBodyStatusCreatedResourcesItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
CreateOrganizationActionBatchParamsBodyActionsItems0 create organization action batch params body actions items0
swagger:model CreateOrganizationActionBatchParamsBodyActionsItems0
*/
type CreateOrganizationActionBatchParamsBodyActionsItems0 struct {

	// The body of the action
	Body interface{} `json:"body,omitempty"`

	// The operation to be used
	// Required: true
	Operation *string `json:"operation"`

	// Unique identifier for the resource to be acted on
	// Required: true
	Resource *string `json:"resource"`
}

// Validate validates this create organization action batch params body actions items0
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateResource(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) validateOperation(formats strfmt.Registry) error {

	if err := validate.Required("operation", "body", o.Operation); err != nil {
		return err
	}

	return nil
}

func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) validateResource(formats strfmt.Registry) error {

	if err := validate.Required("resource", "body", o.Resource); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this create organization action batch params body actions items0 based on context it is used
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *CreateOrganizationActionBatchParamsBodyActionsItems0) UnmarshalBinary(b []byte) error {
	var res CreateOrganizationActionBatchParamsBodyActionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
