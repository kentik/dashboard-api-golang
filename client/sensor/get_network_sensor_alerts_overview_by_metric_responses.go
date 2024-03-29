// Code generated by go-swagger; DO NOT EDIT.

package sensor

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

// GetNetworkSensorAlertsOverviewByMetricReader is a Reader for the GetNetworkSensorAlertsOverviewByMetric structure.
type GetNetworkSensorAlertsOverviewByMetricReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSensorAlertsOverviewByMetricReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSensorAlertsOverviewByMetricOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/sensor/alerts/overview/byMetric] getNetworkSensorAlertsOverviewByMetric", response, response.Code())
	}
}

// NewGetNetworkSensorAlertsOverviewByMetricOK creates a GetNetworkSensorAlertsOverviewByMetricOK with default headers values
func NewGetNetworkSensorAlertsOverviewByMetricOK() *GetNetworkSensorAlertsOverviewByMetricOK {
	return &GetNetworkSensorAlertsOverviewByMetricOK{}
}

/*
GetNetworkSensorAlertsOverviewByMetricOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSensorAlertsOverviewByMetricOK struct {
	Payload []*GetNetworkSensorAlertsOverviewByMetricOKBodyItems0
}

// IsSuccess returns true when this get network sensor alerts overview by metric o k response has a 2xx status code
func (o *GetNetworkSensorAlertsOverviewByMetricOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sensor alerts overview by metric o k response has a 3xx status code
func (o *GetNetworkSensorAlertsOverviewByMetricOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sensor alerts overview by metric o k response has a 4xx status code
func (o *GetNetworkSensorAlertsOverviewByMetricOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sensor alerts overview by metric o k response has a 5xx status code
func (o *GetNetworkSensorAlertsOverviewByMetricOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sensor alerts overview by metric o k response a status code equal to that given
func (o *GetNetworkSensorAlertsOverviewByMetricOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sensor alerts overview by metric o k response
func (o *GetNetworkSensorAlertsOverviewByMetricOK) Code() int {
	return 200
}

func (o *GetNetworkSensorAlertsOverviewByMetricOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sensor/alerts/overview/byMetric][%d] getNetworkSensorAlertsOverviewByMetricOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSensorAlertsOverviewByMetricOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sensor/alerts/overview/byMetric][%d] getNetworkSensorAlertsOverviewByMetricOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSensorAlertsOverviewByMetricOK) GetPayload() []*GetNetworkSensorAlertsOverviewByMetricOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSensorAlertsOverviewByMetricOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSensorAlertsOverviewByMetricOKBodyItems0 get network sensor alerts overview by metric o k body items0
swagger:model GetNetworkSensorAlertsOverviewByMetricOKBodyItems0
*/
type GetNetworkSensorAlertsOverviewByMetricOKBodyItems0 struct {

	// counts
	Counts *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts `json:"counts,omitempty"`

	// End of the timespan over which sensor alerts are counted
	// Format: date-time
	EndTs strfmt.DateTime `json:"endTs,omitempty"`

	// Start of the timespan over which sensor alerts are counted
	// Format: date-time
	StartTs strfmt.DateTime `json:"startTs,omitempty"`
}

// Validate validates this get network sensor alerts overview by metric o k body items0
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCounts(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateEndTs(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateStartTs(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) validateCounts(formats strfmt.Registry) error {
	if swag.IsZero(o.Counts) { // not required
		return nil
	}

	if o.Counts != nil {
		if err := o.Counts.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counts")
			}
			return err
		}
	}

	return nil
}

func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) validateEndTs(formats strfmt.Registry) error {
	if swag.IsZero(o.EndTs) { // not required
		return nil
	}

	if err := validate.FormatOf("endTs", "body", "date-time", o.EndTs.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) validateStartTs(formats strfmt.Registry) error {
	if swag.IsZero(o.StartTs) { // not required
		return nil
	}

	if err := validate.FormatOf("startTs", "body", "date-time", o.StartTs.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this get network sensor alerts overview by metric o k body items0 based on the context it is used
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCounts(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) contextValidateCounts(ctx context.Context, formats strfmt.Registry) error {

	if o.Counts != nil {

		if swag.IsZero(o.Counts) { // not required
			return nil
		}

		if err := o.Counts.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counts")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counts")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorAlertsOverviewByMetricOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts Counts of sensor alerts over the timespan, by reading metric
swagger:model GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts
*/
type GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts struct {

	// Number of sensor alerts that occurred due to an open door
	Door int64 `json:"door,omitempty"`

	// Number of sensor alerts that occurred due to humidity readings
	Humidity int64 `json:"humidity,omitempty"`

	// Number of sensor alerts that occurred due to indoor air quality readings
	IndoorAirQuality int64 `json:"indoorAirQuality,omitempty"`

	// noise
	Noise *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise `json:"noise,omitempty"`

	// Number of sensor alerts that occurred due to PM2.5 readings
	Pm25 int64 `json:"pm25,omitempty"`

	// Number of sensor alerts that occurred due to temperature readings
	Temperature int64 `json:"temperature,omitempty"`

	// Number of sensor alerts that occurred due to TVOC readings
	Tvoc int64 `json:"tvoc,omitempty"`

	// Number of sensor alerts that occurred due to the presence of water
	Water int64 `json:"water,omitempty"`
}

// Validate validates this get network sensor alerts overview by metric o k body items0 counts
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateNoise(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts) validateNoise(formats strfmt.Registry) error {
	if swag.IsZero(o.Noise) { // not required
		return nil
	}

	if o.Noise != nil {
		if err := o.Noise.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counts" + "." + "noise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counts" + "." + "noise")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network sensor alerts overview by metric o k body items0 counts based on the context it is used
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateNoise(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts) contextValidateNoise(ctx context.Context, formats strfmt.Registry) error {

	if o.Noise != nil {

		if swag.IsZero(o.Noise) { // not required
			return nil
		}

		if err := o.Noise.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("counts" + "." + "noise")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("counts" + "." + "noise")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorAlertsOverviewByMetricOKBodyItems0Counts
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise Object containing the number of sensor alerts that occurred due to noise readings
swagger:model GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise
*/
type GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise struct {

	// Number of sensor alerts that occurred due to ambient noise readings
	Ambient int64 `json:"ambient,omitempty"`
}

// Validate validates this get network sensor alerts overview by metric o k body items0 counts noise
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network sensor alerts overview by metric o k body items0 counts noise based on context it is used
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise) UnmarshalBinary(b []byte) error {
	var res GetNetworkSensorAlertsOverviewByMetricOKBodyItems0CountsNoise
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
