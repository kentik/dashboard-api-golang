// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UpdateNetworkTrafficAnalysisReader is a Reader for the UpdateNetworkTrafficAnalysis structure.
type UpdateNetworkTrafficAnalysisReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateNetworkTrafficAnalysisReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateNetworkTrafficAnalysisOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateNetworkTrafficAnalysisOK creates a UpdateNetworkTrafficAnalysisOK with default headers values
func NewUpdateNetworkTrafficAnalysisOK() *UpdateNetworkTrafficAnalysisOK {
	return &UpdateNetworkTrafficAnalysisOK{}
}

/* UpdateNetworkTrafficAnalysisOK describes a response with status code 200, with default header values.

Successful operation
*/
type UpdateNetworkTrafficAnalysisOK struct {
	Payload interface{}
}

// IsSuccess returns true when this update network traffic analysis o k response has a 2xx status code
func (o *UpdateNetworkTrafficAnalysisOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this update network traffic analysis o k response has a 3xx status code
func (o *UpdateNetworkTrafficAnalysisOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this update network traffic analysis o k response has a 4xx status code
func (o *UpdateNetworkTrafficAnalysisOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this update network traffic analysis o k response has a 5xx status code
func (o *UpdateNetworkTrafficAnalysisOK) IsServerError() bool {
	return false
}

// IsCode returns true when this update network traffic analysis o k response a status code equal to that given
func (o *UpdateNetworkTrafficAnalysisOK) IsCode(code int) bool {
	return code == 200
}

func (o *UpdateNetworkTrafficAnalysisOK) Error() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/trafficAnalysis][%d] updateNetworkTrafficAnalysisOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkTrafficAnalysisOK) String() string {
	return fmt.Sprintf("[PUT /networks/{networkId}/trafficAnalysis][%d] updateNetworkTrafficAnalysisOK  %+v", 200, o.Payload)
}

func (o *UpdateNetworkTrafficAnalysisOK) GetPayload() interface{} {
	return o.Payload
}

func (o *UpdateNetworkTrafficAnalysisOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*UpdateNetworkTrafficAnalysisBody update network traffic analysis body
// Example: {"customPieChartItems":[{"name":"Item from hostname","type":"host","value":"example.com"},{"name":"Item from port","type":"port","value":"440"},{"name":"Item from IP","type":"ipRange","value":"192.1.0.0"},{"name":"Item from IP range (CIDR)","type":"ipRange","value":"192.2.0.0/16"},{"name":"Item from IP range with port","type":"ipRange","value":"192.3.0.0/16:80"}],"mode":"detailed"}
swagger:model UpdateNetworkTrafficAnalysisBody
*/
type UpdateNetworkTrafficAnalysisBody struct {

	// The list of items that make up the custom pie chart for traffic reporting.
	CustomPieChartItems []*UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0 `json:"customPieChartItems"`

	//     The traffic analysis mode for the network. Can be one of 'disabled' (do not collect traffic types),
	//     'basic' (collect generic traffic categories), or 'detailed' (collect destination hostnames).
	//
	// Enum: [disabled basic detailed]
	Mode string `json:"mode,omitempty"`
}

// Validate validates this update network traffic analysis body
func (o *UpdateNetworkTrafficAnalysisBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateCustomPieChartItems(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkTrafficAnalysisBody) validateCustomPieChartItems(formats strfmt.Registry) error {
	if swag.IsZero(o.CustomPieChartItems) { // not required
		return nil
	}

	for i := 0; i < len(o.CustomPieChartItems); i++ {
		if swag.IsZero(o.CustomPieChartItems[i]) { // not required
			continue
		}

		if o.CustomPieChartItems[i] != nil {
			if err := o.CustomPieChartItems[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkTrafficAnalysis" + "." + "customPieChartItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkTrafficAnalysis" + "." + "customPieChartItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var updateNetworkTrafficAnalysisBodyTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["disabled","basic","detailed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkTrafficAnalysisBodyTypeModePropEnum = append(updateNetworkTrafficAnalysisBodyTypeModePropEnum, v)
	}
}

const (

	// UpdateNetworkTrafficAnalysisBodyModeDisabled captures enum value "disabled"
	UpdateNetworkTrafficAnalysisBodyModeDisabled string = "disabled"

	// UpdateNetworkTrafficAnalysisBodyModeBasic captures enum value "basic"
	UpdateNetworkTrafficAnalysisBodyModeBasic string = "basic"

	// UpdateNetworkTrafficAnalysisBodyModeDetailed captures enum value "detailed"
	UpdateNetworkTrafficAnalysisBodyModeDetailed string = "detailed"
)

// prop value enum
func (o *UpdateNetworkTrafficAnalysisBody) validateModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkTrafficAnalysisBodyTypeModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkTrafficAnalysisBody) validateMode(formats strfmt.Registry) error {
	if swag.IsZero(o.Mode) { // not required
		return nil
	}

	// value enum
	if err := o.validateModeEnum("updateNetworkTrafficAnalysis"+"."+"mode", "body", o.Mode); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this update network traffic analysis body based on the context it is used
func (o *UpdateNetworkTrafficAnalysisBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateCustomPieChartItems(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkTrafficAnalysisBody) contextValidateCustomPieChartItems(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.CustomPieChartItems); i++ {

		if o.CustomPieChartItems[i] != nil {
			if err := o.CustomPieChartItems[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("updateNetworkTrafficAnalysis" + "." + "customPieChartItems" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("updateNetworkTrafficAnalysis" + "." + "customPieChartItems" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkTrafficAnalysisBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkTrafficAnalysisBody) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkTrafficAnalysisBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0 update network traffic analysis params body custom pie chart items items0
swagger:model UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0
*/
type UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0 struct {

	// The name of the custom pie chart item.
	// Required: true
	Name *string `json:"name"`

	//     The signature type for the custom pie chart item. Can be one of 'host', 'port' or 'ipRange'.
	//
	// Required: true
	// Enum: [host port ipRange]
	Type *string `json:"type"`

	//     The value of the custom pie chart item. Valid syntax depends on the signature type of the chart item
	//     (see sample request/response for more details).
	//
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this update network traffic analysis params body custom pie chart items items0
func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", o.Name); err != nil {
		return err
	}

	return nil
}

var updateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["host","port","ipRange"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeTypePropEnum = append(updateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeTypePropEnum, v)
	}
}

const (

	// UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeHost captures enum value "host"
	UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeHost string = "host"

	// UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypePort captures enum value "port"
	UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypePort string = "port"

	// UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeIPRange captures enum value "ipRange"
	UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeIPRange string = "ipRange"
)

// prop value enum
func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0TypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", o.Type); err != nil {
		return err
	}

	// value enum
	if err := o.validateTypeEnum("type", "body", *o.Type); err != nil {
		return err
	}

	return nil
}

func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", o.Value); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this update network traffic analysis params body custom pie chart items items0 based on context it is used
func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0) UnmarshalBinary(b []byte) error {
	var res UpdateNetworkTrafficAnalysisParamsBodyCustomPieChartItemsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
