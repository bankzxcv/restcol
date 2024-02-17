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

// NewRestColServiceGetCollectionParams creates a new RestColServiceGetCollectionParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestColServiceGetCollectionParams() *RestColServiceGetCollectionParams {
	return &RestColServiceGetCollectionParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestColServiceGetCollectionParamsWithTimeout creates a new RestColServiceGetCollectionParams object
// with the ability to set a timeout on a request.
func NewRestColServiceGetCollectionParamsWithTimeout(timeout time.Duration) *RestColServiceGetCollectionParams {
	return &RestColServiceGetCollectionParams{
		timeout: timeout,
	}
}

// NewRestColServiceGetCollectionParamsWithContext creates a new RestColServiceGetCollectionParams object
// with the ability to set a context for a request.
func NewRestColServiceGetCollectionParamsWithContext(ctx context.Context) *RestColServiceGetCollectionParams {
	return &RestColServiceGetCollectionParams{
		Context: ctx,
	}
}

// NewRestColServiceGetCollectionParamsWithHTTPClient creates a new RestColServiceGetCollectionParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestColServiceGetCollectionParamsWithHTTPClient(client *http.Client) *RestColServiceGetCollectionParams {
	return &RestColServiceGetCollectionParams{
		HTTPClient: client,
	}
}

/*
RestColServiceGetCollectionParams contains all the parameters to send to the API endpoint

	for the rest col service get collection operation.

	Typically these are written to a http.Request.
*/
type RestColServiceGetCollectionParams struct {

	// Cid.
	Cid string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rest col service get collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceGetCollectionParams) WithDefaults() *RestColServiceGetCollectionParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rest col service get collection params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceGetCollectionParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) WithTimeout(timeout time.Duration) *RestColServiceGetCollectionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) WithContext(ctx context.Context) *RestColServiceGetCollectionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) WithHTTPClient(client *http.Client) *RestColServiceGetCollectionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) WithCid(cid string) *RestColServiceGetCollectionParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the rest col service get collection params
func (o *RestColServiceGetCollectionParams) SetCid(cid string) {
	o.Cid = cid
}

// WriteToRequest writes these params to a swagger request
func (o *RestColServiceGetCollectionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cid
	if err := r.SetPathParam("cid", o.Cid); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
