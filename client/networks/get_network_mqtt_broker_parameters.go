// Code generated by go-swagger; DO NOT EDIT.

package networks

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

// NewGetNetworkMqttBrokerParams creates a new GetNetworkMqttBrokerParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetNetworkMqttBrokerParams() *GetNetworkMqttBrokerParams {
	return &GetNetworkMqttBrokerParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetNetworkMqttBrokerParamsWithTimeout creates a new GetNetworkMqttBrokerParams object
// with the ability to set a timeout on a request.
func NewGetNetworkMqttBrokerParamsWithTimeout(timeout time.Duration) *GetNetworkMqttBrokerParams {
	return &GetNetworkMqttBrokerParams{
		timeout: timeout,
	}
}

// NewGetNetworkMqttBrokerParamsWithContext creates a new GetNetworkMqttBrokerParams object
// with the ability to set a context for a request.
func NewGetNetworkMqttBrokerParamsWithContext(ctx context.Context) *GetNetworkMqttBrokerParams {
	return &GetNetworkMqttBrokerParams{
		Context: ctx,
	}
}

// NewGetNetworkMqttBrokerParamsWithHTTPClient creates a new GetNetworkMqttBrokerParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetNetworkMqttBrokerParamsWithHTTPClient(client *http.Client) *GetNetworkMqttBrokerParams {
	return &GetNetworkMqttBrokerParams{
		HTTPClient: client,
	}
}

/*
GetNetworkMqttBrokerParams contains all the parameters to send to the API endpoint

	for the get network mqtt broker operation.

	Typically these are written to a http.Request.
*/
type GetNetworkMqttBrokerParams struct {

	/* MqttBrokerID.

	   Mqtt broker ID
	*/
	MqttBrokerID string

	/* NetworkID.

	   Network ID
	*/
	NetworkID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get network mqtt broker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkMqttBrokerParams) WithDefaults() *GetNetworkMqttBrokerParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get network mqtt broker params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetNetworkMqttBrokerParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) WithTimeout(timeout time.Duration) *GetNetworkMqttBrokerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) WithContext(ctx context.Context) *GetNetworkMqttBrokerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) WithHTTPClient(client *http.Client) *GetNetworkMqttBrokerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithMqttBrokerID adds the mqttBrokerID to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) WithMqttBrokerID(mqttBrokerID string) *GetNetworkMqttBrokerParams {
	o.SetMqttBrokerID(mqttBrokerID)
	return o
}

// SetMqttBrokerID adds the mqttBrokerId to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) SetMqttBrokerID(mqttBrokerID string) {
	o.MqttBrokerID = mqttBrokerID
}

// WithNetworkID adds the networkID to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) WithNetworkID(networkID string) *GetNetworkMqttBrokerParams {
	o.SetNetworkID(networkID)
	return o
}

// SetNetworkID adds the networkId to the get network mqtt broker params
func (o *GetNetworkMqttBrokerParams) SetNetworkID(networkID string) {
	o.NetworkID = networkID
}

// WriteToRequest writes these params to a swagger request
func (o *GetNetworkMqttBrokerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
