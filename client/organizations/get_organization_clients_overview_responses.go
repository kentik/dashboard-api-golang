// Code generated by go-swagger; DO NOT EDIT.

package organizations

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

// GetOrganizationClientsOverviewReader is a Reader for the GetOrganizationClientsOverview structure.
type GetOrganizationClientsOverviewReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationClientsOverviewReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationClientsOverviewOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /organizations/{organizationId}/clients/overview] getOrganizationClientsOverview", response, response.Code())
	}
}

// NewGetOrganizationClientsOverviewOK creates a GetOrganizationClientsOverviewOK with default headers values
func NewGetOrganizationClientsOverviewOK() *GetOrganizationClientsOverviewOK {
	return &GetOrganizationClientsOverviewOK{}
}

/*
GetOrganizationClientsOverviewOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationClientsOverviewOK struct {
	Payload *GetOrganizationClientsOverviewOKBody
}

// IsSuccess returns true when this get organization clients overview o k response has a 2xx status code
func (o *GetOrganizationClientsOverviewOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization clients overview o k response has a 3xx status code
func (o *GetOrganizationClientsOverviewOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization clients overview o k response has a 4xx status code
func (o *GetOrganizationClientsOverviewOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization clients overview o k response has a 5xx status code
func (o *GetOrganizationClientsOverviewOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization clients overview o k response a status code equal to that given
func (o *GetOrganizationClientsOverviewOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get organization clients overview o k response
func (o *GetOrganizationClientsOverviewOK) Code() int {
	return 200
}

func (o *GetOrganizationClientsOverviewOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/clients/overview][%d] getOrganizationClientsOverviewOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationClientsOverviewOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/clients/overview][%d] getOrganizationClientsOverviewOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationClientsOverviewOK) GetPayload() *GetOrganizationClientsOverviewOKBody {
	return o.Payload
}

func (o *GetOrganizationClientsOverviewOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(GetOrganizationClientsOverviewOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetOrganizationClientsOverviewOKBody get organization clients overview o k body
swagger:model GetOrganizationClientsOverviewOKBody
*/
type GetOrganizationClientsOverviewOKBody struct {

	// counts
	Counts *GetOrganizationClientsOverviewOKBodyCounts `json:"counts,omitempty"`

	// usage
	Usage *GetOrganizationClientsOverviewOKBodyUsage `json:"usage,omitempty"`
}

// Validate validates this get organization clients overview o k body
func (o *GetOrganizationClientsOverviewOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCounts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationClientsOverviewOKBody) validateCounts(formats strfmt.Registry) error {
	if swag.IsZero(o.Counts) { // not required
		return nil
	}

	if o.Counts != nil {
		if err := o.Counts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationClientsOverviewOK" + "." + "counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationClientsOverviewOK" + "." + "counts")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationClientsOverviewOKBody) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(o.Usage) { // not required
		return nil
	}

	if o.Usage != nil {
		if err := o.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization clients overview o k body based on the context it is used
func (o *GetOrganizationClientsOverviewOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := o.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationClientsOverviewOKBody) contextValidateCounts(ctx context.Context, formats strfmt.Registry) error {

	if o.Counts != nil {

		if swag.IsZero(o.Counts) { // not required
			return nil
		}

		if err := o.Counts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationClientsOverviewOK" + "." + "counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationClientsOverviewOK" + "." + "counts")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationClientsOverviewOKBody) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if o.Usage != nil {

		if swag.IsZero(o.Usage) { // not required
			return nil
		}

		if err := o.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBody) UnmarshalBinary(b []byte) error {
	var res GetOrganizationClientsOverviewOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationClientsOverviewOKBodyCounts Client count information
swagger:model GetOrganizationClientsOverviewOKBodyCounts
*/
type GetOrganizationClientsOverviewOKBodyCounts struct {

	// Total number of clients with data usage in organization
	Total int64 `json:"total,omitempty"`
}

// Validate validates this get organization clients overview o k body counts
func (o *GetOrganizationClientsOverviewOKBodyCounts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization clients overview o k body counts based on context it is used
func (o *GetOrganizationClientsOverviewOKBodyCounts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBodyCounts) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBodyCounts) UnmarshalBinary(b []byte) error {
	var res GetOrganizationClientsOverviewOKBodyCounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationClientsOverviewOKBodyUsage Usage information of all clients across organization
swagger:model GetOrganizationClientsOverviewOKBodyUsage
*/
type GetOrganizationClientsOverviewOKBodyUsage struct {

	// Average data usage (in kb) of each client in organization
	Average float32 `json:"average,omitempty"`

	// overall
	Overall *GetOrganizationClientsOverviewOKBodyUsageOverall `json:"overall,omitempty"`
}

// Validate validates this get organization clients overview o k body usage
func (o *GetOrganizationClientsOverviewOKBodyUsage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateOverall(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationClientsOverviewOKBodyUsage) validateOverall(formats strfmt.Registry) error {
	if swag.IsZero(o.Overall) { // not required
		return nil
	}

	if o.Overall != nil {
		if err := o.Overall.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage" + "." + "overall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage" + "." + "overall")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization clients overview o k body usage based on the context it is used
func (o *GetOrganizationClientsOverviewOKBodyUsage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateOverall(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationClientsOverviewOKBodyUsage) contextValidateOverall(ctx context.Context, formats strfmt.Registry) error {

	if o.Overall != nil {

		if swag.IsZero(o.Overall) { // not required
			return nil
		}

		if err := o.Overall.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage" + "." + "overall")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("getOrganizationClientsOverviewOK" + "." + "usage" + "." + "overall")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBodyUsage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBodyUsage) UnmarshalBinary(b []byte) error {
	var res GetOrganizationClientsOverviewOKBodyUsage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetOrganizationClientsOverviewOKBodyUsageOverall Overall data usage of all clients across organization
swagger:model GetOrganizationClientsOverviewOKBodyUsageOverall
*/
type GetOrganizationClientsOverviewOKBodyUsageOverall struct {

	// Downstream data usage (in kb) of all clients across organization
	Downstream float32 `json:"downstream,omitempty"`

	// Total data usage (in kb) of all clients across organization
	Total float32 `json:"total,omitempty"`

	// Upstream data usage (in kb) of all clients across organization
	Upstream float32 `json:"upstream,omitempty"`
}

// Validate validates this get organization clients overview o k body usage overall
func (o *GetOrganizationClientsOverviewOKBodyUsageOverall) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization clients overview o k body usage overall based on context it is used
func (o *GetOrganizationClientsOverviewOKBodyUsageOverall) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBodyUsageOverall) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationClientsOverviewOKBodyUsageOverall) UnmarshalBinary(b []byte) error {
	var res GetOrganizationClientsOverviewOKBodyUsageOverall
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
