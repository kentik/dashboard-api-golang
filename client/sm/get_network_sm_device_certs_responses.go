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

// GetNetworkSmDeviceCertsReader is a Reader for the GetNetworkSmDeviceCerts structure.
type GetNetworkSmDeviceCertsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmDeviceCertsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmDeviceCertsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/sm/devices/{deviceId}/certs] getNetworkSmDeviceCerts", response, response.Code())
	}
}

// NewGetNetworkSmDeviceCertsOK creates a GetNetworkSmDeviceCertsOK with default headers values
func NewGetNetworkSmDeviceCertsOK() *GetNetworkSmDeviceCertsOK {
	return &GetNetworkSmDeviceCertsOK{}
}

/*
GetNetworkSmDeviceCertsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmDeviceCertsOK struct {
	Payload []*GetNetworkSmDeviceCertsOKBodyItems0
}

// IsSuccess returns true when this get network sm device certs o k response has a 2xx status code
func (o *GetNetworkSmDeviceCertsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm device certs o k response has a 3xx status code
func (o *GetNetworkSmDeviceCertsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm device certs o k response has a 4xx status code
func (o *GetNetworkSmDeviceCertsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm device certs o k response has a 5xx status code
func (o *GetNetworkSmDeviceCertsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm device certs o k response a status code equal to that given
func (o *GetNetworkSmDeviceCertsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network sm device certs o k response
func (o *GetNetworkSmDeviceCertsOK) Code() int {
	return 200
}

func (o *GetNetworkSmDeviceCertsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/certs][%d] getNetworkSmDeviceCertsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceCertsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/devices/{deviceId}/certs][%d] getNetworkSmDeviceCertsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmDeviceCertsOK) GetPayload() []*GetNetworkSmDeviceCertsOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSmDeviceCertsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkSmDeviceCertsOKBodyItems0 get network sm device certs o k body items0
swagger:model GetNetworkSmDeviceCertsOKBodyItems0
*/
type GetNetworkSmDeviceCertsOKBodyItems0 struct {

	// The PEM of the certificate.
	CertPem string `json:"certPem,omitempty"`

	// The Meraki managed device Id.
	DeviceID string `json:"deviceId,omitempty"`

	// The Meraki Id of the certificate record.
	ID string `json:"id,omitempty"`

	// The certificate issuer.
	Issuer string `json:"issuer,omitempty"`

	// The name of the certificate.
	Name string `json:"name,omitempty"`

	// The date after which the certificate is no longer valid.
	NotValidAfter string `json:"notValidAfter,omitempty"`

	// The date before which the certificate is not valid.
	NotValidBefore string `json:"notValidBefore,omitempty"`

	// The subject of the certificate.
	Subject string `json:"subject,omitempty"`
}

// Validate validates this get network sm device certs o k body items0
func (o *GetNetworkSmDeviceCertsOKBodyItems0) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network sm device certs o k body items0 based on context it is used
func (o *GetNetworkSmDeviceCertsOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmDeviceCertsOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmDeviceCertsOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmDeviceCertsOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
