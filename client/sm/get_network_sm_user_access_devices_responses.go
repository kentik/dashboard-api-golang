// Code generated by go-swagger; DO NOT EDIT.

package sm

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"fmt"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetNetworkSmUserAccessDevicesReader is a Reader for the GetNetworkSmUserAccessDevices structure.
type GetNetworkSmUserAccessDevicesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkSmUserAccessDevicesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkSmUserAccessDevicesOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGetNetworkSmUserAccessDevicesOK creates a GetNetworkSmUserAccessDevicesOK with default headers values
func NewGetNetworkSmUserAccessDevicesOK() *GetNetworkSmUserAccessDevicesOK {
	return &GetNetworkSmUserAccessDevicesOK{}
}

/* GetNetworkSmUserAccessDevicesOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkSmUserAccessDevicesOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload []*GetNetworkSmUserAccessDevicesOKBodyItems0
}

// IsSuccess returns true when this get network sm user access devices o k response has a 2xx status code
func (o *GetNetworkSmUserAccessDevicesOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network sm user access devices o k response has a 3xx status code
func (o *GetNetworkSmUserAccessDevicesOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network sm user access devices o k response has a 4xx status code
func (o *GetNetworkSmUserAccessDevicesOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network sm user access devices o k response has a 5xx status code
func (o *GetNetworkSmUserAccessDevicesOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network sm user access devices o k response a status code equal to that given
func (o *GetNetworkSmUserAccessDevicesOK) IsCode(code int) bool {
	return code == 200
}

func (o *GetNetworkSmUserAccessDevicesOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/userAccessDevices][%d] getNetworkSmUserAccessDevicesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmUserAccessDevicesOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/sm/userAccessDevices][%d] getNetworkSmUserAccessDevicesOK  %+v", 200, o.Payload)
}

func (o *GetNetworkSmUserAccessDevicesOK) GetPayload() []*GetNetworkSmUserAccessDevicesOKBodyItems0 {
	return o.Payload
}

func (o *GetNetworkSmUserAccessDevicesOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GetNetworkSmUserAccessDevicesOKBodyItems0 get network sm user access devices o k body items0
swagger:model GetNetworkSmUserAccessDevicesOKBodyItems0
*/
type GetNetworkSmUserAccessDevicesOKBodyItems0 struct {

	// user email
	Email string `json:"email,omitempty"`

	// device ID
	ID string `json:"id,omitempty"`

	// mac address
	Mac string `json:"mac,omitempty"`

	// device name
	Name string `json:"name,omitempty"`

	// system type
	SystemType string `json:"systemType,omitempty"`

	// device tags
	Tags []string `json:"tags"`

	// Array of trusted access configs
	TrustedAccessConnections []*GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0 `json:"trustedAccessConnections"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this get network sm user access devices o k body items0
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTrustedAccessConnections(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSmUserAccessDevicesOKBodyItems0) validateTrustedAccessConnections(formats strfmt.Registry) error {
	if swag.IsZero(o.TrustedAccessConnections) { // not required
		return nil
	}

	for i := 0; i < len(o.TrustedAccessConnections); i++ {
		if swag.IsZero(o.TrustedAccessConnections[i]) { // not required
			continue
		}

		if o.TrustedAccessConnections[i] != nil {
			if err := o.TrustedAccessConnections[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trustedAccessConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trustedAccessConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network sm user access devices o k body items0 based on the context it is used
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateTrustedAccessConnections(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSmUserAccessDevicesOKBodyItems0) contextValidateTrustedAccessConnections(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.TrustedAccessConnections); i++ {

		if o.TrustedAccessConnections[i] != nil {
			if err := o.TrustedAccessConnections[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("trustedAccessConnections" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("trustedAccessConnections" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmUserAccessDevicesOKBodyItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0 get network sm user access devices o k body items0 trusted access connections items0
swagger:model GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0
*/
type GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0 struct {

	// time that config was downloaded
	DownloadedAt string `json:"downloadedAt,omitempty"`

	// time of last connection
	// Format: date-time
	LastConnectedAt strfmt.DateTime `json:"lastConnectedAt,omitempty"`

	// time that SCEP completed
	// Format: date-time
	ScepCompletedAt strfmt.DateTime `json:"scepCompletedAt,omitempty"`

	// config id
	TrustedAccessConfigID string `json:"trustedAccessConfigId,omitempty"`
}

// Validate validates this get network sm user access devices o k body items0 trusted access connections items0
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateLastConnectedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateScepCompletedAt(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0) validateLastConnectedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.LastConnectedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("lastConnectedAt", "body", "date-time", o.LastConnectedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (o *GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0) validateScepCompletedAt(formats strfmt.Registry) error {
	if swag.IsZero(o.ScepCompletedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("scepCompletedAt", "body", "date-time", o.ScepCompletedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this get network sm user access devices o k body items0 trusted access connections items0 based on context it is used
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkSmUserAccessDevicesOKBodyItems0TrustedAccessConnectionsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
