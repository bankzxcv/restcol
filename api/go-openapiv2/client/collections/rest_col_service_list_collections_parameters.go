// Code generated by go-swagger; DO NOT EDIT.

package collections

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

// NewRestColServiceListCollectionsParams creates a new RestColServiceListCollectionsParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestColServiceListCollectionsParams() *RestColServiceListCollectionsParams {
	return &RestColServiceListCollectionsParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestColServiceListCollectionsParamsWithTimeout creates a new RestColServiceListCollectionsParams object
// with the ability to set a timeout on a request.
func NewRestColServiceListCollectionsParamsWithTimeout(timeout time.Duration) *RestColServiceListCollectionsParams {
	return &RestColServiceListCollectionsParams{
		timeout: timeout,
	}
}

// NewRestColServiceListCollectionsParamsWithContext creates a new RestColServiceListCollectionsParams object
// with the ability to set a context for a request.
func NewRestColServiceListCollectionsParamsWithContext(ctx context.Context) *RestColServiceListCollectionsParams {
	return &RestColServiceListCollectionsParams{
		Context: ctx,
	}
}

// NewRestColServiceListCollectionsParamsWithHTTPClient creates a new RestColServiceListCollectionsParams object
// with the ability to set a custom HTTPClient for a request.
func NewRestColServiceListCollectionsParamsWithHTTPClient(client *http.Client) *RestColServiceListCollectionsParams {
	return &RestColServiceListCollectionsParams{
		HTTPClient: client,
	}
}

/*
RestColServiceListCollectionsParams contains all the parameters to send to the API endpoint

	for the rest col service list collections operation.

	Typically these are written to a http.Request.
*/
type RestColServiceListCollectionsParams struct {

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rest col service list collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceListCollectionsParams) WithDefaults() *RestColServiceListCollectionsParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rest col service list collections params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceListCollectionsParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) WithTimeout(timeout time.Duration) *RestColServiceListCollectionsParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) WithContext(ctx context.Context) *RestColServiceListCollectionsParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) WithHTTPClient(client *http.Client) *RestColServiceListCollectionsParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectID adds the projectID to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) WithProjectID(projectID string) *RestColServiceListCollectionsParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the rest col service list collections params
func (o *RestColServiceListCollectionsParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *RestColServiceListCollectionsParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param projectId
	if err := r.SetPathParam("projectId", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
