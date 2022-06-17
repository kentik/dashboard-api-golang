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

// GetNetworkSmDeviceCellularUsageHistoryReader is a Reader for the GetNetworkSmDeviceCellularUsageHistory structure.
type GetNetworkSmDeviceCellularUsageHistoryReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceCellularUsageHistoryReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceCellularUsageHistoryOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmDeviceCellularUsageHistoryOK creates a GetNetworkSmDeviceCellularUsageHistoryOK with default headers values
func NewGetNetworkSmDeviceCellularUsageHistoryOK() *GetNetworkSmDeviceCellularUsageHistoryOK {
	return &GetNetworkSmDeviceCellularUsageHistoryOK{}
}

/* GetNetworkSmDeviceCellularUsageHistoryOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceCellularUsageHistoryOK struct {
	Payload []*GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0
}

// IsSuccess returns true when this get network sm device cellular usage history o k response has a 2xx status code
func (o *GetNetworkSmDeviceCellularUsageHistoryOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm device cellular usage history o k response has a 3xx status code
func (o *GetNetworkSmDeviceCellularUsageHistoryOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm device cellular usage history o k response has a 4xx status code
func (o *GetNetworkSmDeviceCellularUsageHistoryOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm device cellular usage history o k response has a 5xx status code
func (o *GetNetworkSmDeviceCellularUsageHistoryOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm device cellular usage history o k response a status code equal to that given
func (o *GetNetworkSmDeviceCellularUsageHistoryOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkSmDeviceCellularUsageHistoryOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory][%d] getNetworkSmDeviceCellularUsageHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceCellularUsageHistoryOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/cellularUsageHistory][%d] getNetworkSmDeviceCellularUsageHistoryOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceCellularUsageHistoryOK) GetPayload() []*GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSmDeviceCellularUsageHistoryOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0 get network sm device cellular usage history o k body items0
swagger:model GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0
*/
type GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0 struct {

	// The amount of cellular data received by the device.
	Received float32 `json:"received,omitempty"`

	// The amount of cellular sent received by the device.
	Sent float32 `json:"sent,omitempty"`

	// When the cellular usage data was collected.
	Ts string `json:"ts,omitempty"`
}

// Validate validates this get network sm device cellular usage history o k body items0
func (o *GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network sm device cellular usage history o k body items0 based on context it is used
func (o *GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmDeviceCellularUsageHistoryOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}