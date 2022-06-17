// Code generated by go-swagger; DO NOT EDIT.

package networks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ClaimNetworkDevicesReader is a Reader for the ClaimNetworkDevices structure.
type ClaimNetworkDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ClaimNetworkDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewClaimNetworkDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewClaimNetworkDevicesOK creates a ClaimNetworkDevicesOK with default headers values
func NewClaimNetworkDevicesOK() *ClaimNetworkDevicesOK {
	return &ClaimNetworkDevicesOK{}
}

/* ClaimNetworkDevicesOK describes a response with status code 200, with default header values.

Successful operation
*/
type ClaimNetworkDevicesOK struct {
}

// IsSuccess returns true when this claim network devices o k response has a 2xx status code
func (o *ClaimNetworkDevicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this claim network devices o k response has a 3xx status code
func (o *ClaimNetworkDevicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this claim network devices o k response has a 4xx status code
func (o *ClaimNetworkDevicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this claim network devices o k response has a 5xx status code
func (o *ClaimNetworkDevicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this claim network devices o k response a status code equal to that given
func (o *ClaimNetworkDevicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *ClaimNetworkDevicesOK) Error() string {
	return fmt.Sprintf("[POST /networks/{networkId}/devices/claim][%d] claimNetworkDevicesOK ", 200)
}

func (o *ClaimNetworkDevicesOK) String() string {
	return fmt.Sprintf("[POST /networks/{networkId}/devices/claim][%d] claimNetworkDevicesOK ", 200)
}

func (o *ClaimNetworkDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

/*ClaimNetworkDevicesBody claim network devices body
// Example: {"serials":["Q234-ABCD-0001","Q234-ABCD-0002","Q234-ABCD-0003"]}
swagger:model ClaimNetworkDevicesBody
*/
type ClaimNetworkDevicesBody struct {

	// A list of serials of devices to claim
	// Required: true
	Serials []string `json:"serials"`
}

// Validate validates this claim network devices body
func (o *ClaimNetworkDevicesBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateSerials(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ClaimNetworkDevicesBody) validateSerials(formats strfmt.Registry) error {

	if err := validate.Required("claimNetworkDevices"+"."+"serials", "body", o.Serials); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this claim network devices body based on context it is used
func (o *ClaimNetworkDevicesBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ClaimNetworkDevicesBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ClaimNetworkDevicesBody) UnmarshalBinary(b []byte) error {
	var res ClaimNetworkDevicesBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
