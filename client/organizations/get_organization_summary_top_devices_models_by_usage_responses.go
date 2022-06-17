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

// GetOrganizationSummaryTopDevicesModelsByUsageReader is a Reader for the GetOrganizationSummaryTopDevicesModelsByUsage structure.
type GetOrganizationSummaryTopDevicesModelsByUsageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetOrganizationSummaryTopDevicesModelsByUsageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetOrganizationSummaryTopDevicesModelsByUsageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetOrganizationSummaryTopDevicesModelsByUsageOK creates a GetOrganizationSummaryTopDevicesModelsByUsageOK with default headers values
func NewGetOrganizationSummaryTopDevicesModelsByUsageOK() *GetOrganizationSummaryTopDevicesModelsByUsageOK {
	return &GetOrganizationSummaryTopDevicesModelsByUsageOK{}
}

/* GetOrganizationSummaryTopDevicesModelsByUsageOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetOrganizationSummaryTopDevicesModelsByUsageOK struct {
	Payload []*GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0
}

// IsSuccess returns true when this get organization summary top devices models by usage o k response has a 2xx status code
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get organization summary top devices models by usage o k response has a 3xx status code
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get organization summary top devices models by usage o k response has a 4xx status code
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get organization summary top devices models by usage o k response has a 5xx status code
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get organization summary top devices models by usage o k response a status code equal to that given
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) Error() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/summary/top/devices/models/byUsage][%d] getOrganizationSummaryTopDevicesModelsByUsageOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) String() string {
	return fmt.Sprintf("[GET /organizations/{organizationId}/summary/top/devices/models/byUsage][%d] getOrganizationSummaryTopDevicesModelsByUsageOK  %+v", 200, o.Payload)
}

func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) GetPayload() []*GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0 {
	return o.Payload
}

func (o *GetOrganizationSummaryTopDevicesModelsByUsageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0 get organization summary top devices models by usage o k body items0
swagger:model GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0
*/
type GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0 struct {

	// Total number of devices per model
	Count int64 `json:"count,omitempty"`

	// The device model
	Model string `json:"model,omitempty"`

	// usage
	Usage *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage `json:"usage,omitempty"`
}

// Validate validates this get organization summary top devices models by usage o k body items0
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0) validateUsage(formats strfmt.Registry) error {
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

// ContextValidate validate this get organization summary top devices models by usage o k body items0 based on the context it is used
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateUsage(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0) contextValidateUsage(ctx context.Context, formats strfmt.Registry) error {

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
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage Usage info in megabytes
swagger:model GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage
*/
type GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage struct {

	// Average usage in megabytes
	Average float32 `json:"average,omitempty"`

	// Total usage in megabytes
	Total float32 `json:"total,omitempty"`
}

// Validate validates this get organization summary top devices models by usage o k body items0 usage
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get organization summary top devices models by usage o k body items0 usage based on context it is used
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage) UnmarshalBinary(b []byte) error {
	var res GetOrganizationSummaryTopDevicesModelsByUsageOKBodyItems0Usage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
