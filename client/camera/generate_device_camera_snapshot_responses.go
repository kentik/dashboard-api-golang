// Code generated by go-swagger; DO NOT EDIT.

package camera

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

// GenerateDeviceCameraSnapshotReader is a Reader for the GenerateDeviceCameraSnapshot structure.
type GenerateDeviceCameraSnapshotReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GenerateDeviceCameraSnapshotReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 202:
		result := NewGenerateDeviceCameraSnapshotAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewGenerateDeviceCameraSnapshotAccepted creates a GenerateDeviceCameraSnapshotAccepted with default headers values
func NewGenerateDeviceCameraSnapshotAccepted() *GenerateDeviceCameraSnapshotAccepted {
	return &GenerateDeviceCameraSnapshotAccepted{}
}

/* GenerateDeviceCameraSnapshotAccepted describes a response with status code 202, with default header values.

Successful operation
*/
type GenerateDeviceCameraSnapshotAccepted struct {
	Payload interface{}
}

// IsSuccess returns true when this generate device camera snapshot accepted response has a 2xx status code
func (o *GenerateDeviceCameraSnapshotAccepted) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this generate device camera snapshot accepted response has a 3xx status code
func (o *GenerateDeviceCameraSnapshotAccepted) IsRedirect() bool {
	return false
}

// IsClientError returns true when this generate device camera snapshot accepted response has a 4xx status code
func (o *GenerateDeviceCameraSnapshotAccepted) IsClientError() bool {
	return false
}

// IsServerError returns true when this generate device camera snapshot accepted response has a 5xx status code
func (o *GenerateDeviceCameraSnapshotAccepted) IsServerError() bool {
	return false
}

// IsCode returns true when this generate device camera snapshot accepted response a status code equal to that given
func (o *GenerateDeviceCameraSnapshotAccepted) IsCode(code int) bool {
	return code == 202
}

func (o *GenerateDeviceCameraSnapshotAccepted) Error() string {
	return fmt.Sprintf("[POST /devices/{serial}/camera/generateSnapshot][%d] generateDeviceCameraSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *GenerateDeviceCameraSnapshotAccepted) String() string {
	return fmt.Sprintf("[POST /devices/{serial}/camera/generateSnapshot][%d] generateDeviceCameraSnapshotAccepted  %+v", 202, o.Payload)
}

func (o *GenerateDeviceCameraSnapshotAccepted) GetPayload() interface{} {
	return o.Payload
}

func (o *GenerateDeviceCameraSnapshotAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*GenerateDeviceCameraSnapshotBody generate device camera snapshot body
// Example: {"fullframe":false,"timestamp":"2021-04-30T15:18:08Z"}
swagger:model GenerateDeviceCameraSnapshotBody
*/
type GenerateDeviceCameraSnapshotBody struct {

	// [optional] If set to "true" the snapshot will be taken at full sensor resolution. This will error if used with timestamp.
	Fullframe bool `json:"fullframe,omitempty"`

	// [optional] The snapshot will be taken from this time on the camera. The timestamp is expected to be in ISO 8601 format. If no timestamp is specified, we will assume current time.
	// Format: date-time
	Timestamp strfmt.DateTime `json:"timestamp,omitempty"`
}

// Validate validates this generate device camera snapshot body
func (o *GenerateDeviceCameraSnapshotBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GenerateDeviceCameraSnapshotBody) validateTimestamp(formats strfmt.Registry) error {
	if swag.IsZero(o.Timestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("generateDeviceCameraSnapshot"+"."+"timestamp", "body", "date-time", o.Timestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this generate device camera snapshot body based on context it is used
func (o *GenerateDeviceCameraSnapshotBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GenerateDeviceCameraSnapshotBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GenerateDeviceCameraSnapshotBody) UnmarshalBinary(b []byte) error {
	var res GenerateDeviceCameraSnapshotBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
