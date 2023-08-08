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

// GetNetworkSmDeviceWlanListsReader is a Reader for the GetNetworkSmDeviceWlanLists structure.
type GetNetworkSmDeviceWlanListsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceWlanListsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceWlanListsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/sm/devices/{deviceId}/wlanLists] getNetworkSmDeviceWlanLists", response, response.Code())
	}
}

// NewGetNetworkSmDeviceWlanListsOK creates a GetNetworkSmDeviceWlanListsOK with default headers values
func NewGetNetworkSmDeviceWlanListsOK() *GetNetworkSmDeviceWlanListsOK {
	return &GetNetworkSmDeviceWlanListsOK{}
}

/*
GetNetworkSmDeviceWlanListsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceWlanListsOK struct {
	Payload []*GetNetworkSmDeviceWlanListsOKBodyItems0
}

// IsSuccess returns true when this get network sm device wlan lists o k response has a 2xx status code
func (o *GetNetworkSmDeviceWlanListsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm device wlan lists o k response has a 3xx status code
func (o *GetNetworkSmDeviceWlanListsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm device wlan lists o k response has a 4xx status code
func (o *GetNetworkSmDeviceWlanListsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm device wlan lists o k response has a 5xx status code
func (o *GetNetworkSmDeviceWlanListsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm device wlan lists o k response a status code equal to that given
func (o *GetNetworkSmDeviceWlanListsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sm device wlan lists o k response
func (o *GetNetworkSmDeviceWlanListsOK) Code() int {
	return 200
}

func (o *GetNetworkSmDeviceWlanListsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/wlanLists][%d] getNetworkSmDeviceWlanListsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceWlanListsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/wlanLists][%d] getNetworkSmDeviceWlanListsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceWlanListsOK) GetPayload() []*GetNetworkSmDeviceWlanListsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSmDeviceWlanListsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSmDeviceWlanListsOKBodyItems0 get network sm device wlan lists o k body items0
swagger:model GetNetworkSmDeviceWlanListsOKBodyItems0
*/
type GetNetworkSmDeviceWlanListsOKBodyItems0 struct {

	// When the Meraki record for the wlanList was created.
	CreatedAt string `json:"createdAt,omitempty"`

	// The Meraki managed Id of the wlanList record.
	ID string `json:"id,omitempty"`

	// An XML string containing the WLAN List for the device.
	XML string `json:"xml,omitempty"`
}

// Validate validates this get network sm device wlan lists o k body items0
func (o *GetNetworkSmDeviceWlanListsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network sm device wlan lists o k body items0 based on context it is used
func (o *GetNetworkSmDeviceWlanListsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmDeviceWlanListsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmDeviceWlanListsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmDeviceWlanListsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
