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

// GetOrganizationSummaryTopClientsManufacturersByUsageReader is a Reader for the GetOrganizationSummaryTopClientsManufacturersByUsage structure.
type GetOrganizationSummaryTopClientsManufacturersByUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSummaryTopClientsManufacturersByUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationSummaryTopClientsManufacturersByUsageOK creates a GetOrganizationSummaryTopClientsManufacturersByUsageOK with default headers values
func NewGetOrganizationSummaryTopClientsManufacturersByUsageOK() *GetOrganizationSummaryTopClientsManufacturersByUsageOK {
	return &GetOrganizationSummaryTopClientsManufacturersByUsageOK{}
}

/* GetOrganizationSummaryTopClientsManufacturersByUsageOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSummaryTopClientsManufacturersByUsageOK struct {
	Payload []*GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0
}

// IsSuccess returns true when this get organization summary top clients manufacturers by usage o k response has a 2xx status code
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization summary top clients manufacturers by usage o k response has a 3xx status code
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization summary top clients manufacturers by usage o k response has a 4xx status code
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization summary top clients manufacturers by usage o k response has a 5xx status code
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization summary top clients manufacturers by usage o k response a status code equal to that given
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/summary/top/clients/manufacturers/byUsage][%d] getOrganizationSummaryTopClientsManufacturersByUsageOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/summary/top/clients/manufacturers/byUsage][%d] getOrganizationSummaryTopClientsManufacturersByUsageOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) GetPayload() []*GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0 get organization summary top clients manufacturers by usage o k body items0
swagger:model GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0
*/
type GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0 struct {

	// clients
	Clients *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients `json:"clients,omitempty"`

	// Name of the manufacturer
	Name string `json:"name,omitempty"`

	// usage
	Usage *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage `json:"usage,omitempty"`
}

// Validate validates this get organization summary top clients manufacturers by usage o k body items0
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateClients(formats); err != nil {
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

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) validateClients(formats strfmt.Registry) error {
	if swag.IsZero(o.Clients) { // not required
		return nil
	}

	if o.Clients != nil {
		if err := o.Clients.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clients")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) validateUsage(formats strfmt.Registry) error {
	if swag.IsZero(o.Usage) { // not required
		return nil
	}

	if o.Usage != nil {
		if err := o.Usage.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization summary top clients manufacturers by usage o k body items0 based on the context it is used
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateClients(ctx, formats); err != nil {
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

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) contextValidateClients(ctx context.Context, formats strfmt.Registry) error {

	if o.Clients != nil {
		if err := o.Clients.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clients")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clients")
			}
			return err
		}
	}

	return nil
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

	if o.Usage != nil {
		if err := o.Usage.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("usage")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("usage")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients Clients info
swagger:model GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients
*/
type GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients struct {

	// counts
	Counts *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts `json:"counts,omitempty"`
}

// Validate validates this get organization summary top clients manufacturers by usage o k body items0 clients
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCounts(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients) validateCounts(formats strfmt.Registry) error {
	if swag.IsZero(o.Counts) { // not required
		return nil
	}

	if o.Counts != nil {
		if err := o.Counts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clients" + "." + "counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clients" + "." + "counts")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get organization summary top clients manufacturers by usage o k body items0 clients based on the context it is used
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients) contextValidateCounts(ctx context.Context, formats strfmt.Registry) error {

	if o.Counts != nil {
		if err := o.Counts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("clients" + "." + "counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("clients" + "." + "counts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Clients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts Counts of clients
swagger:model GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts
*/
type GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts struct {

	// Total counts of clients
	Total int64 `json:"total,omitempty"`
}

// Validate validates this get organization summary top clients manufacturers by usage o k body items0 clients counts
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization summary top clients manufacturers by usage o k body items0 clients counts based on context it is used
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0ClientsCounts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage Clients usage
swagger:model GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage
*/
type GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage struct {

	// Downstream data usage by client
	Downstream float32 `json:"downstream,omitempty"`

	// Total data usage by client
	Total float32 `json:"total,omitempty"`

	// Upstream data usage by client
	Upstream float32 `json:"upstream,omitempty"`
}

// Validate validates this get organization summary top clients manufacturers by usage o k body items0 usage
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization summary top clients manufacturers by usage o k body items0 usage based on context it is used
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopClientsManufacturersByUsageOKBodyItems0Usage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
