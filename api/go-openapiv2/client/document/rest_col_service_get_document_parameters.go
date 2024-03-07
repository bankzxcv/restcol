// Code generated by go-swagger; DO NOT EDIT.

package document

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

// NewRestColServiceGetDocumentParams creates a new RestColServiceGetDocumentParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestColServiceGetDocumentParams() *RestColServiceGetDocumentParams {
	return &RestColServiceGetDocumentParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestColServiceGetDocumentParamsWithTimeout creates a new RestColServiceGetDocumentParams object
// with the ability to set a timeout on a request.
func NewRestColServiceGetDocumentParamsWithTimeout(timeout time.Duration) *RestColServiceGetDocumentParams {
	return &RestColServiceGetDocumentParams{
		timeout: timeout,
	}
}

// NewRestColServiceGetDocumentParamsWithContext creates a new RestColServiceGetDocumentParams object
// with the ability to set a context for a request.
func NewRestColServiceGetDocumentParamsWithContext(ctx context.Context) *RestColServiceGetDocumentParams {
	return &RestColServiceGetDocumentParams{
		Context: ctx,
	}
}

// NewRestColServiceGetDocumentParamsWithHTTPClient creates a new RestColServiceGetDocumentParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestColServiceGetDocumentParamsWithHTTPClient(client *http.Client) *RestColServiceGetDocumentParams {
	return &RestColServiceGetDocumentParams{
		HTTPClient: client,
	}
}

/*
RestColServiceGetDocumentParams contains all the parameters to send to the API endpoint

	for the rest col service get document operation.

	Typically these are written to a http.Request.
*/
type RestColServiceGetDocumentParams struct {

	// Cid.
	Cid string

	// Did.
	Did string

	// Pid.
	Pid *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rest col service get document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceGetDocumentParams) WithDefaults() *RestColServiceGetDocumentParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rest col service get document params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceGetDocumentParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rest col service get document params
func (o *RestColServiceGetDocumentParams) WithTimeout(timeout time.Duration) *RestColServiceGetDocumentParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rest col service get document params
func (o *RestColServiceGetDocumentParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rest col service get document params
func (o *RestColServiceGetDocumentParams) WithContext(ctx context.Context) *RestColServiceGetDocumentParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rest col service get document params
func (o *RestColServiceGetDocumentParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rest col service get document params
func (o *RestColServiceGetDocumentParams) WithHTTPClient(client *http.Client) *RestColServiceGetDocumentParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rest col service get document params
func (o *RestColServiceGetDocumentParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCid adds the cid to the rest col service get document params
func (o *RestColServiceGetDocumentParams) WithCid(cid string) *RestColServiceGetDocumentParams {
	o.SetCid(cid)
	return o
}

// SetCid adds the cid to the rest col service get document params
func (o *RestColServiceGetDocumentParams) SetCid(cid string) {
	o.Cid = cid
}

// WithDid adds the did to the rest col service get document params
func (o *RestColServiceGetDocumentParams) WithDid(did string) *RestColServiceGetDocumentParams {
	o.SetDid(did)
	return o
}

// SetDid adds the did to the rest col service get document params
func (o *RestColServiceGetDocumentParams) SetDid(did string) {
	o.Did = did
}

// WithPid adds the pid to the rest col service get document params
func (o *RestColServiceGetDocumentParams) WithPid(pid *string) *RestColServiceGetDocumentParams {
	o.SetPid(pid)
	return o
}

// SetPid adds the pid to the rest col service get document params
func (o *RestColServiceGetDocumentParams) SetPid(pid *string) {
	o.Pid = pid
}

// WriteToRequest writes these params to a swagger request
func (o *RestColServiceGetDocumentParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Pid != nil {

		// query param pid
		var qrPid string

		if o.Pid != nil {
			qrPid = *o.Pid
		}
		qPid := qrPid
		if qPid != "" {

			if err := r.SetQueryParam("pid", qPid); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
