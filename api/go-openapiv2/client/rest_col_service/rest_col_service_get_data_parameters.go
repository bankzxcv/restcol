// Code generated by go-swagger; DO NOT EDIT.

package rest_col_service

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

// NewRestColServiceGetDataParams creates a new RestColServiceGetDataParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestColServiceGetDataParams() *RestColServiceGetDataParams {
	return &RestColServiceGetDataParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestColServiceGetDataParamsWithTimeout creates a new RestColServiceGetDataParams object
// with the ability to set a timeout on a request.
func NewRestColServiceGetDataParamsWithTimeout(timeout time.Duration) *RestColServiceGetDataParams {
	return &RestColServiceGetDataParams{
		timeout: timeout,
	}
}

// NewRestColServiceGetDataParamsWithContext creates a new RestColServiceGetDataParams object
// with the ability to set a context for a request.
func NewRestColServiceGetDataParamsWithContext(ctx context.Context) *RestColServiceGetDataParams {
	return &RestColServiceGetDataParams{
		Context: ctx,
	}
}

// NewRestColServiceGetDataParamsWithHTTPClient creates a new RestColServiceGetDataParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestColServiceGetDataParamsWithHTTPClient(client *http.Client) *RestColServiceGetDataParams {
	return &RestColServiceGetDataParams{
		HTTPClient: client,
	}
}

/*
RestColServiceGetDataParams contains all the parameters to send to the API endpoint

	for the rest col service get data operation.

	Typically these are written to a http.Request.
*/
type RestColServiceGetDataParams struct {

	// Cid.
	Cid string

	// Did.
	Did string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rest col service get data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceGetDataParams) WithDefaults() *RestColServiceGetDataParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rest col service get data params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceGetDataParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rest col service get data params
func (o *RestColServiceGetDataParams) WithTimeout(timeout time.Duration) *RestColServiceGetDataParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rest col service get data params
func (o *RestColServiceGetDataParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rest col service get data params
func (o *RestColServiceGetDataParams) WithContext(ctx context.Context) *RestColServiceGetDataParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rest col service get data params
func (o *RestColServiceGetDataParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rest col service get data params
func (o *RestColServiceGetDataParams) WithHTTPClient(client *http.Client) *RestColServiceGetDataParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rest col service get data params
func (o *RestColServiceGetDataParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the rest col service get data params
func (o *RestColServiceGetDataParams) WithCid(cid string) *RestColServiceGetDataParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the rest col service get data params
func (o *RestColServiceGetDataParams) SetCid(cid string) {
	o.Cid = cid
}

// WithDid adds the did to the rest col service get data params
func (o *RestColServiceGetDataParams) WithDid(did string) *RestColServiceGetDataParams {
	o.SetDid(did)
	return o
}

// SetDid adds the did to the rest col service get data params
func (o *RestColServiceGetDataParams) SetDid(did string) {
	o.Did = did
}

// WriteToRequest writes these params to a swagger request
func (o *RestColServiceGetDataParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cid
	if err := r.SetPathParam("cid", o.Cid); err != nil {
		return err
	}

	// path param did
	if err := r.SetPathParam("did", o.Did); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}