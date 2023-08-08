// Code generated by go-swagger; DO NOT EDIT.

package sensor

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// NewUpdateNetworkSensorMqttBrokerParams creates a new UpdateNetworkSensorMqttBrokerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewUpdateNetworkSensorMqttBrokerParams() *UpdateNetworkSensorMqttBrokerParams {
	return &UpdateNetworkSensorMqttBrokerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateNetworkSensorMqttBrokerParamsWithTimeout creates a new UpdateNetworkSensorMqttBrokerParams object
// with the ability to set a timeout on a request.
func NewUpdateNetworkSensorMqttBrokerParamsWithTimeout(timeout time.Duration) *UpdateNetworkSensorMqttBrokerParams {
	return &UpdateNetworkSensorMqttBrokerParams{
		timeout: timeout,
	}
}

// NewUpdateNetworkSensorMqttBrokerParamsWithContext creates a new UpdateNetworkSensorMqttBrokerParams object
// with the ability to set a context for a request.
func NewUpdateNetworkSensorMqttBrokerParamsWithContext(ctx context.Context) *UpdateNetworkSensorMqttBrokerParams {
	return &UpdateNetworkSensorMqttBrokerParams{
		Context: ctx,
	}
}

// NewUpdateNetworkSensorMqttBrokerParamsWithHTTPClient creates a new UpdateNetworkSensorMqttBrokerParams object
// with the ability to set a custom HTTPClient for a request.
func NewUpdateNetworkSensorMqttBrokerParamsWithHTTPClient(client *http.Client) *UpdateNetworkSensorMqttBrokerParams {
	return &UpdateNetworkSensorMqttBrokerParams{
		HTTPClient: client,
	}
}

/*
UpdateNetworkSensorMqttBrokerParams contains all the parameters to send to the API endpoint

	for the update network sensor mqtt broker operation.

	Typically these are written to a http.Request.
*/
type UpdateNetworkSensorMqttBrokerParams struct {

	/* MqttBrokerID.

	   Mqtt broker ID
	*/
	MqttBrokerID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	// UpdateNetworkSensorMqttBroker.
	UpdateNetworkSensorMqttBroker UpdateNetworkSensorMqttBrokerBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the update network sensor mqtt broker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSensorMqttBrokerParams) WithDefaults() *UpdateNetworkSensorMqttBrokerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the update network sensor mqtt broker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *UpdateNetworkSensorMqttBrokerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) WithTimeout(timeout time.Duration) *UpdateNetworkSensorMqttBrokerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) WithContext(ctx context.Context) *UpdateNetworkSensorMqttBrokerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) WithHTTPClient(client *http.Client) *UpdateNetworkSensorMqttBrokerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMqttBrokerID adds the mqttBrokerID to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) WithMqttBrokerID(mqttBrokerID string) *UpdateNetworkSensorMqttBrokerParams {
	o.SetMqttBrokerID(mqttBrokerID)
	return o
}

// SetMqttBrokerID adds the mqttBrokerId to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) SetMqttBrokerID(mqttBrokerID string) {
	o.MqttBrokerID = mqttBrokerID
}

// WithNetworkID adds the networkID to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) WithNetworkID(networkID string) *UpdateNetworkSensorMqttBrokerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WithUpdateNetworkSensorMqttBroker adds the updateNetworkSensorMqttBroker to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) WithUpdateNetworkSensorMqttBroker(updateNetworkSensorMqttBroker UpdateNetworkSensorMqttBrokerBody) *UpdateNetworkSensorMqttBrokerParams {
	o.SetUpdateNetworkSensorMqttBroker(updateNetworkSensorMqttBroker)
	return o
}

// SetUpdateNetworkSensorMqttBroker adds the updateNetworkSensorMqttBroker to the update network sensor mqtt broker params
func (o *UpdateNetworkSensorMqttBrokerParams) SetUpdateNetworkSensorMqttBroker(updateNetworkSensorMqttBroker UpdateNetworkSensorMqttBrokerBody) {
	o.UpdateNetworkSensorMqttBroker = updateNetworkSensorMqttBroker
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateNetworkSensorMqttBrokerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param mqttBrokerId
	if err := r.SetPathParam("mqttBrokerId", o.MqttBrokerID); err != nil {
		return err
	}

	// path param networkId
	if err := r.SetPathParam("networkId", o.NetworkID); err != nil {
		return err
	}
	if err := r.SetBodyParam(o.UpdateNetworkSensorMqttBroker); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
