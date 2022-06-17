// Code generated by go-swagger; DO NOT EDIT.

package wireless

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// GetDeviceWirelessBluetoothSettingsReader is a Reader for the GetDeviceWirelessBluetoothSettings structure.
type GetDeviceWirelessBluetoothSettingsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetDeviceWirelessBluetoothSettingsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetDeviceWirelessBluetoothSettingsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetDeviceWirelessBluetoothSettingsOK creates a GetDeviceWirelessBluetoothSettingsOK with default headers values
func NewGetDeviceWirelessBluetoothSettingsOK() *GetDeviceWirelessBluetoothSettingsOK {
	return &GetDeviceWirelessBluetoothSettingsOK{}
}

/* GetDeviceWirelessBluetoothSettingsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetDeviceWirelessBluetoothSettingsOK struct {
	Payload interface{}
}

// IsSuccess returns true when this get device wireless bluetooth settings o k response has a 2xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get device wireless bluetooth settings o k response has a 3xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get device wireless bluetooth settings o k response has a 4xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get device wireless bluetooth settings o k response has a 5xx status code
func (o *GetDeviceWirelessBluetoothSettingsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get device wireless bluetooth settings o k response a status code equal to that given
func (o *GetDeviceWirelessBluetoothSettingsOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetDeviceWirelessBluetoothSettingsOK) Error() string {
	return fmt.Sprintf("[GET /devices/{serial}/wireless/bluetooth/settings][%d] getDeviceWirelessBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceWirelessBluetoothSettingsOK) String() string {
	return fmt.Sprintf("[GET /devices/{serial}/wireless/bluetooth/settings][%d] getDeviceWirelessBluetoothSettingsOK  %+v", 200, o.Payload)
}

func (o *GetDeviceWirelessBluetoothSettingsOK) GetPayload() interface{} {
	return o.Payload
}

func (o *GetDeviceWirelessBluetoothSettingsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}