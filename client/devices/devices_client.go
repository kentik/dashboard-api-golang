// Code generated by go-swagger; DO NOT EDIT.

package devices

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new devices API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for devices API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	BlinkDeviceLeds(params *BlinkDeviceLedsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BlinkDeviceLedsAccepted, error)

	CreateDeviceLiveToolsPing(params *CreateDeviceLiveToolsPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeviceLiveToolsPingCreated, error)

	CreateDeviceLiveToolsPingDevice(params *CreateDeviceLiveToolsPingDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeviceLiveToolsPingDeviceCreated, error)

	GetDevice(params *GetDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceOK, error)

	GetDeviceCellularSims(params *GetDeviceCellularSimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceCellularSimsOK, error)

	GetDeviceClients(params *GetDeviceClientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceClientsOK, error)

	GetDeviceLiveToolsPing(params *GetDeviceLiveToolsPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLiveToolsPingOK, error)

	GetDeviceLiveToolsPingDevice(params *GetDeviceLiveToolsPingDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLiveToolsPingDeviceOK, error)

	GetDeviceLldpCdp(params *GetDeviceLldpCdpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLldpCdpOK, error)

	GetDeviceLossAndLatencyHistory(params *GetDeviceLossAndLatencyHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLossAndLatencyHistoryOK, error)

	GetDeviceManagementInterface(params *GetDeviceManagementInterfaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceManagementInterfaceOK, error)

	RebootDevice(params *RebootDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RebootDeviceAccepted, error)

	UpdateDevice(params *UpdateDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeviceOK, error)

	UpdateDeviceCellularSims(params *UpdateDeviceCellularSimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeviceCellularSimsOK, error)

	UpdateDeviceManagementInterface(params *UpdateDeviceManagementInterfaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeviceManagementInterfaceOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
BlinkDeviceLeds blinks the l e ds on a device

Blink the LEDs on a device
*/
func (a *Client) BlinkDeviceLeds(params *BlinkDeviceLedsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*BlinkDeviceLedsAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewBlinkDeviceLedsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "blinkDeviceLeds",
		Method:             "POST",
		PathPattern:        "/devices/{serial}/blinkLeds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &BlinkDeviceLedsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*BlinkDeviceLedsAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for blinkDeviceLeds: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateDeviceLiveToolsPing enqueues a job to ping a target host from the device

Enqueue a job to ping a target host from the device
*/
func (a *Client) CreateDeviceLiveToolsPing(params *CreateDeviceLiveToolsPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeviceLiveToolsPingCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceLiveToolsPingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeviceLiveToolsPing",
		Method:             "POST",
		PathPattern:        "/devices/{serial}/liveTools/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeviceLiveToolsPingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeviceLiveToolsPingCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeviceLiveToolsPing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
CreateDeviceLiveToolsPingDevice enqueues a job to check connectivity status to the device

Enqueue a job to check connectivity status to the device
*/
func (a *Client) CreateDeviceLiveToolsPingDevice(params *CreateDeviceLiveToolsPingDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*CreateDeviceLiveToolsPingDeviceCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateDeviceLiveToolsPingDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "createDeviceLiveToolsPingDevice",
		Method:             "POST",
		PathPattern:        "/devices/{serial}/liveTools/pingDevice",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &CreateDeviceLiveToolsPingDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateDeviceLiveToolsPingDeviceCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createDeviceLiveToolsPingDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDevice returns a single device

Return a single device
*/
func (a *Client) GetDevice(params *GetDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDevice",
		Method:             "GET",
		PathPattern:        "/devices/{serial}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceCellularSims returns the s i m and a p n configurations for a cellular device

Return the SIM and APN configurations for a cellular device.
*/
func (a *Client) GetDeviceCellularSims(params *GetDeviceCellularSimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceCellularSimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceCellularSimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceCellularSims",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/cellular/sims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceCellularSimsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceCellularSimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceCellularSims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceClients lists the clients of a device up to a maximum of a month ago

List the clients of a device, up to a maximum of a month ago. The usage of each client is returned in kilobytes. If the device is a switch, the switchport is returned; otherwise the switchport field is null.
*/
func (a *Client) GetDeviceClients(params *GetDeviceClientsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceClientsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceClientsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceClients",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/clients",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceClientsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceClientsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceClients: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceLiveToolsPing returns a ping job

Return a ping job. Latency unit in response is in milliseconds. Size is in bytes.
*/
func (a *Client) GetDeviceLiveToolsPing(params *GetDeviceLiveToolsPingParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLiveToolsPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceLiveToolsPingParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceLiveToolsPing",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/liveTools/ping/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceLiveToolsPingReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceLiveToolsPingOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceLiveToolsPing: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceLiveToolsPingDevice returns a ping device job

Return a ping device job. Latency unit in response is in milliseconds. Size is in bytes.
*/
func (a *Client) GetDeviceLiveToolsPingDevice(params *GetDeviceLiveToolsPingDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLiveToolsPingDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceLiveToolsPingDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceLiveToolsPingDevice",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/liveTools/pingDevice/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceLiveToolsPingDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceLiveToolsPingDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceLiveToolsPingDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceLldpCdp lists l l d p and c d p information for a device

List LLDP and CDP information for a device
*/
func (a *Client) GetDeviceLldpCdp(params *GetDeviceLldpCdpParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLldpCdpOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceLldpCdpParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceLldpCdp",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/lldpCdp",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceLldpCdpReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceLldpCdpOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceLldpCdp: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceLossAndLatencyHistory gets the uplink loss percentage and latency in milliseconds and goodput in kilobits per second for m x m g and z devices

Get the uplink loss percentage and latency in milliseconds, and goodput in kilobits per second for MX, MG and Z devices.
*/
func (a *Client) GetDeviceLossAndLatencyHistory(params *GetDeviceLossAndLatencyHistoryParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceLossAndLatencyHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceLossAndLatencyHistoryParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceLossAndLatencyHistory",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/lossAndLatencyHistory",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceLossAndLatencyHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceLossAndLatencyHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceLossAndLatencyHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
GetDeviceManagementInterface returns the management interface settings for a device

Return the management interface settings for a device
*/
func (a *Client) GetDeviceManagementInterface(params *GetDeviceManagementInterfaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*GetDeviceManagementInterfaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetDeviceManagementInterfaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "getDeviceManagementInterface",
		Method:             "GET",
		PathPattern:        "/devices/{serial}/managementInterface",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &GetDeviceManagementInterfaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetDeviceManagementInterfaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getDeviceManagementInterface: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
RebootDevice reboots a device

Reboot a device
*/
func (a *Client) RebootDevice(params *RebootDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RebootDeviceAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRebootDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "rebootDevice",
		Method:             "POST",
		PathPattern:        "/devices/{serial}/reboot",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RebootDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RebootDeviceAccepted)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for rebootDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDevice updates the attributes of a device

Update the attributes of a device
*/
func (a *Client) UpdateDevice(params *UpdateDeviceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeviceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDevice",
		Method:             "PUT",
		PathPattern:        "/devices/{serial}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeviceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDevice: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDeviceCellularSims updates the s i m and a p n configurations for a cellular device

Updates the SIM and APN configurations for a cellular device.
*/
func (a *Client) UpdateDeviceCellularSims(params *UpdateDeviceCellularSimsParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeviceCellularSimsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceCellularSimsParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDeviceCellularSims",
		Method:             "PUT",
		PathPattern:        "/devices/{serial}/cellular/sims",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceCellularSimsReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeviceCellularSimsOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDeviceCellularSims: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
UpdateDeviceManagementInterface updates the management interface settings for a device

Update the management interface settings for a device
*/
func (a *Client) UpdateDeviceManagementInterface(params *UpdateDeviceManagementInterfaceParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*UpdateDeviceManagementInterfaceOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateDeviceManagementInterfaceParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "updateDeviceManagementInterface",
		Method:             "PUT",
		PathPattern:        "/devices/{serial}/managementInterface",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &UpdateDeviceManagementInterfaceReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	}
	for _, opt := range opts {
		opt(op)
	}

	result, err := a.transport.Submit(op)
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateDeviceManagementInterfaceOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateDeviceManagementInterface: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
