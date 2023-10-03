// Code generated by go-swagger; DO NOT EDIT.

package act_device_api_service

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

// NewActDeviceAPIServiceRemoveDeviceV1Params creates a new ActDeviceAPIServiceRemoveDeviceV1Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewActDeviceAPIServiceRemoveDeviceV1Params() *ActDeviceAPIServiceRemoveDeviceV1Params {
	return &ActDeviceAPIServiceRemoveDeviceV1Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewActDeviceAPIServiceRemoveDeviceV1ParamsWithTimeout creates a new ActDeviceAPIServiceRemoveDeviceV1Params object
// with the ability to set a timeout on a request.
func NewActDeviceAPIServiceRemoveDeviceV1ParamsWithTimeout(timeout time.Duration) *ActDeviceAPIServiceRemoveDeviceV1Params {
	return &ActDeviceAPIServiceRemoveDeviceV1Params{
		timeout: timeout,
	}
}

// NewActDeviceAPIServiceRemoveDeviceV1ParamsWithContext creates a new ActDeviceAPIServiceRemoveDeviceV1Params object
// with the ability to set a context for a request.
func NewActDeviceAPIServiceRemoveDeviceV1ParamsWithContext(ctx context.Context) *ActDeviceAPIServiceRemoveDeviceV1Params {
	return &ActDeviceAPIServiceRemoveDeviceV1Params{
		Context: ctx,
	}
}

// NewActDeviceAPIServiceRemoveDeviceV1ParamsWithHTTPClient creates a new ActDeviceAPIServiceRemoveDeviceV1Params object
// with the ability to set a custom HTTPClient for a request.
func NewActDeviceAPIServiceRemoveDeviceV1ParamsWithHTTPClient(client *http.Client) *ActDeviceAPIServiceRemoveDeviceV1Params {
	return &ActDeviceAPIServiceRemoveDeviceV1Params{
		HTTPClient: client,
	}
}

/*
ActDeviceAPIServiceRemoveDeviceV1Params contains all the parameters to send to the API endpoint

	for the act device Api service remove device v1 operation.

	Typically these are written to a http.Request.
*/
type ActDeviceAPIServiceRemoveDeviceV1Params struct {

	// DeviceID.
	//
	// Format: uint64
	DeviceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the act device Api service remove device v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) WithDefaults() *ActDeviceAPIServiceRemoveDeviceV1Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the act device Api service remove device v1 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) WithTimeout(timeout time.Duration) *ActDeviceAPIServiceRemoveDeviceV1Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) WithContext(ctx context.Context) *ActDeviceAPIServiceRemoveDeviceV1Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) WithHTTPClient(client *http.Client) *ActDeviceAPIServiceRemoveDeviceV1Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDeviceID adds the deviceID to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) WithDeviceID(deviceID string) *ActDeviceAPIServiceRemoveDeviceV1Params {
	o.SetDeviceID(deviceID)
	return o
}

// SetDeviceID adds the deviceId to the act device Api service remove device v1 params
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) SetDeviceID(deviceID string) {
	o.DeviceID = deviceID
}

// WriteToRequest writes these params to a swagger request
func (o *ActDeviceAPIServiceRemoveDeviceV1Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param deviceId
	if err := r.SetPathParam("deviceId", o.DeviceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}