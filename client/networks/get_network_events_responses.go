// Code generated by go-swagger; DO NOT EDIT.

package networks

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
)

// GetNetworkEventsReader is a Reader for the GetNetworkEvents structure.
type GetNetworkEventsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetNetworkEventsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetNetworkEventsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("[GET /networks/{networkId}/events] getNetworkEvents", response, response.Code())
	}
}

// NewGetNetworkEventsOK creates a GetNetworkEventsOK with default headers values
func NewGetNetworkEventsOK() *GetNetworkEventsOK {
	return &GetNetworkEventsOK{}
}

/*
GetNetworkEventsOK describes a response with status code 200, with default header values.

Successful operation
*/
type GetNetworkEventsOK struct {

	/* A comma-separated list of first, last, prev, and next relative links used for subsequent paginated requests.
	 */
	Link string

	Payload *GetNetworkEventsOKBody
}

// IsSuccess returns true when this get network events o k response has a 2xx status code
func (o *GetNetworkEventsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this get network events o k response has a 3xx status code
func (o *GetNetworkEventsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this get network events o k response has a 4xx status code
func (o *GetNetworkEventsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this get network events o k response has a 5xx status code
func (o *GetNetworkEventsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this get network events o k response a status code equal to that given
func (o *GetNetworkEventsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the get network events o k response
func (o *GetNetworkEventsOK) Code() int {
	return 200
}

func (o *GetNetworkEventsOK) Error() string {
	return fmt.Sprintf("[GET /networks/{networkId}/events][%d] getNetworkEventsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkEventsOK) String() string {
	return fmt.Sprintf("[GET /networks/{networkId}/events][%d] getNetworkEventsOK  %+v", 200, o.Payload)
}

func (o *GetNetworkEventsOK) GetPayload() *GetNetworkEventsOKBody {
	return o.Payload
}

func (o *GetNetworkEventsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// hydrates response header Link
	hdrLink := response.GetHeader("Link")

	if hdrLink != "" {
		o.Link = hdrLink
	}

	o.Payload = new(GetNetworkEventsOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*
GetNetworkEventsOKBody get network events o k body
swagger:model GetNetworkEventsOKBody
*/
type GetNetworkEventsOKBody struct {

	// An array of events that took place in the network.
	Events []*GetNetworkEventsOKBodyEventsItems0 `json:"events"`

	// A message regarding the events sent. Usually 'null' unless there are no events
	Message string `json:"message,omitempty"`

	// An UTC ISO8601 string of the latest occured at time of the listed events of the page.
	PageEndAt string `json:"pageEndAt,omitempty"`

	// An UTC ISO8601 string of the earliest occured at time of the listed events of the page.
	PageStartAt string `json:"pageStartAt,omitempty"`
}

// Validate validates this get network events o k body
func (o *GetNetworkEventsOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEvents(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkEventsOKBody) validateEvents(formats strfmt.Registry) error {
	if swag.IsZero(o.Events) { // not required
		return nil
	}

	for i := 0; i < len(o.Events); i++ {
		if swag.IsZero(o.Events[i]) { // not required
			continue
		}

		if o.Events[i] != nil {
			if err := o.Events[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkEventsOK" + "." + "events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkEventsOK" + "." + "events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// ContextValidate validate this get network events o k body based on the context it is used
func (o *GetNetworkEventsOKBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEvents(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkEventsOKBody) contextValidateEvents(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(o.Events); i++ {

		if o.Events[i] != nil {

			if swag.IsZero(o.Events[i]) { // not required
				return nil
			}

			if err := o.Events[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("getNetworkEventsOK" + "." + "events" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("getNetworkEventsOK" + "." + "events" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkEventsOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkEventsOKBody) UnmarshalBinary(b []byte) error {
	var res GetNetworkEventsOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkEventsOKBodyEventsItems0 get network events o k body events items0
swagger:model GetNetworkEventsOKBodyEventsItems0
*/
type GetNetworkEventsOKBodyEventsItems0 struct {

	// The category that the event type belongs to
	Category string `json:"category,omitempty"`

	// A description of the client. This is usually the client's device name.
	ClientDescription string `json:"clientDescription,omitempty"`

	// A string identifying the client. This could be a client's MAC or IP address
	ClientID string `json:"clientId,omitempty"`

	// The client's MAC address.
	ClientMac string `json:"clientMac,omitempty"`

	// A description of the event the happened.
	Description string `json:"description,omitempty"`

	// The name of the device. Only shown if the device is an access point.
	DeviceName string `json:"deviceName,omitempty"`

	// The serial number of the device. Only shown if the device is an access point.
	DeviceSerial string `json:"deviceSerial,omitempty"`

	// event data
	EventData *GetNetworkEventsOKBodyEventsItems0EventData `json:"eventData,omitempty"`

	// The ID of the network.
	NetworkID string `json:"networkId,omitempty"`

	// An UTC ISO8601 string of the time the event occurred at.
	OccurredAt string `json:"occurredAt,omitempty"`

	// The SSID number of the device.
	SsidNumber int64 `json:"ssidNumber,omitempty"`

	// The type of event being listed.
	Type string `json:"type,omitempty"`
}

// Validate validates this get network events o k body events items0
func (o *GetNetworkEventsOKBodyEventsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateEventData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkEventsOKBodyEventsItems0) validateEventData(formats strfmt.Registry) error {
	if swag.IsZero(o.EventData) { // not required
		return nil
	}

	if o.EventData != nil {
		if err := o.EventData.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventData")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this get network events o k body events items0 based on the context it is used
func (o *GetNetworkEventsOKBodyEventsItems0) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := o.contextValidateEventData(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *GetNetworkEventsOKBodyEventsItems0) contextValidateEventData(ctx context.Context, formats strfmt.Registry) error {

	if o.EventData != nil {

		if swag.IsZero(o.EventData) { // not required
			return nil
		}

		if err := o.EventData.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("eventData")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("eventData")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkEventsOKBodyEventsItems0) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkEventsOKBodyEventsItems0) UnmarshalBinary(b []byte) error {
	var res GetNetworkEventsOKBodyEventsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*
GetNetworkEventsOKBodyEventsItems0EventData An object containing more data related to the event.
swagger:model GetNetworkEventsOKBodyEventsItems0EventData
*/
type GetNetworkEventsOKBodyEventsItems0EventData struct {

	// The association ID of the client.
	Aid string `json:"aid,omitempty"`

	// The radio channel the client is connecting to.
	Channel string `json:"channel,omitempty"`

	// The client's IP address
	ClientIP string `json:"client_ip,omitempty"`

	// The client's MAC address
	ClientMac string `json:"client_mac,omitempty"`

	// The radio band number the client is trying to connect to.
	Radio string `json:"radio,omitempty"`

	// The current received signal strength indication (RSSI) of the client connected to an AP.
	Rssi string `json:"rssi,omitempty"`

	// The virtual access point (VAP) number the client is connecting to.
	Vap string `json:"vap,omitempty"`
}

// Validate validates this get network events o k body events items0 event data
func (o *GetNetworkEventsOKBodyEventsItems0EventData) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this get network events o k body events items0 event data based on context it is used
func (o *GetNetworkEventsOKBodyEventsItems0EventData) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *GetNetworkEventsOKBodyEventsItems0EventData) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *GetNetworkEventsOKBodyEventsItems0EventData) UnmarshalBinary(b []byte) error {
	var res GetNetworkEventsOKBodyEventsItems0EventData
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
