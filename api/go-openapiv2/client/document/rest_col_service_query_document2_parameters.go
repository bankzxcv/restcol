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
	"github.com/go-openapi/swag"
)

// NewRestColServiceQueryDocument2Params creates a new RestColServiceQueryDocument2Params object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewRestColServiceQueryDocument2Params() *RestColServiceQueryDocument2Params {
	return &RestColServiceQueryDocument2Params{
		timeout: cr.DefaultTimeout,
	}
}

// NewRestColServiceQueryDocument2ParamsWithTimeout creates a new RestColServiceQueryDocument2Params object
// with the ability to set a timeout on a request.
func NewRestColServiceQueryDocument2ParamsWithTimeout(timeout time.Duration) *RestColServiceQueryDocument2Params {
	return &RestColServiceQueryDocument2Params{
		timeout: timeout,
	}
}

// NewRestColServiceQueryDocument2ParamsWithContext creates a new RestColServiceQueryDocument2Params object
// with the ability to set a context for a request.
func NewRestColServiceQueryDocument2ParamsWithContext(ctx context.Context) *RestColServiceQueryDocument2Params {
	return &RestColServiceQueryDocument2Params{
		Context: ctx,
	}
}

// NewRestColServiceQueryDocument2ParamsWithHTTPClient creates a new RestColServiceQueryDocument2Params object
// with the ability to set a custom HTTPClient for a request.
func NewRestColServiceQueryDocument2ParamsWithHTTPClient(client *http.Client) *RestColServiceQueryDocument2Params {
	return &RestColServiceQueryDocument2Params{
		HTTPClient: client,
	}
}

/*
RestColServiceQueryDocument2Params contains all the parameters to send to the API endpoint

	for the rest col service query document2 operation.

	Typically these are written to a http.Request.
*/
type RestColServiceQueryDocument2Params struct {

	// CollectionID.
	CollectionID string

	/* EndedAt.

	   endedAt specifies when is the ended timeframe of the query

	   Format: date-time
	*/
	EndedAt *strfmt.DateTime

	// LimitCount.
	//
	// Format: int32
	LimitCount *int32

	// ProjectID.
	ProjectID *string

	/* Queryfields.

	   dot-concatenated fields, ex: fielda.fieldb.fieldc
	*/
	Queryfields *string

	// SinceTs.
	//
	// Format: date-time
	SinceTs *strfmt.DateTime

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the rest col service query document2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceQueryDocument2Params) WithDefaults() *RestColServiceQueryDocument2Params {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the rest col service query document2 params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *RestColServiceQueryDocument2Params) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithTimeout(timeout time.Duration) *RestColServiceQueryDocument2Params {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithContext(ctx context.Context) *RestColServiceQueryDocument2Params {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithHTTPClient(client *http.Client) *RestColServiceQueryDocument2Params {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCollectionID adds the collectionID to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithCollectionID(collectionID string) *RestColServiceQueryDocument2Params {
	o.SetCollectionID(collectionID)
	return o
}

// SetCollectionID adds the collectionId to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetCollectionID(collectionID string) {
	o.CollectionID = collectionID
}

// WithEndedAt adds the endedAt to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithEndedAt(endedAt *strfmt.DateTime) *RestColServiceQueryDocument2Params {
	o.SetEndedAt(endedAt)
	return o
}

// SetEndedAt adds the endedAt to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetEndedAt(endedAt *strfmt.DateTime) {
	o.EndedAt = endedAt
}

// WithLimitCount adds the limitCount to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithLimitCount(limitCount *int32) *RestColServiceQueryDocument2Params {
	o.SetLimitCount(limitCount)
	return o
}

// SetLimitCount adds the limitCount to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetLimitCount(limitCount *int32) {
	o.LimitCount = limitCount
}

// WithProjectID adds the projectID to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithProjectID(projectID *string) *RestColServiceQueryDocument2Params {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetProjectID(projectID *string) {
	o.ProjectID = projectID
}

// WithQueryfields adds the queryfields to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithQueryfields(queryfields *string) *RestColServiceQueryDocument2Params {
	o.SetQueryfields(queryfields)
	return o
}

// SetQueryfields adds the queryfields to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetQueryfields(queryfields *string) {
	o.Queryfields = queryfields
}

// WithSinceTs adds the sinceTs to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) WithSinceTs(sinceTs *strfmt.DateTime) *RestColServiceQueryDocument2Params {
	o.SetSinceTs(sinceTs)
	return o
}

// SetSinceTs adds the sinceTs to the rest col service query document2 params
func (o *RestColServiceQueryDocument2Params) SetSinceTs(sinceTs *strfmt.DateTime) {
	o.SinceTs = sinceTs
}

// WriteToRequest writes these params to a swagger request
func (o *RestColServiceQueryDocument2Params) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param collectionId
	if err := r.SetPathParam("collectionId", o.CollectionID); err != nil {
		return err
	}

	if o.EndedAt != nil {

		// query param endedAt
		var qrEndedAt strfmt.DateTime

		if o.EndedAt != nil {
			qrEndedAt = *o.EndedAt
		}
		qEndedAt := qrEndedAt.String()
		if qEndedAt != "" {

			if err := r.SetQueryParam("endedAt", qEndedAt); err != nil {
				return err
			}
		}
	}

	if o.LimitCount != nil {

		// query param limitCount
		var qrLimitCount int32

		if o.LimitCount != nil {
			qrLimitCount = *o.LimitCount
		}
		qLimitCount := swag.FormatInt32(qrLimitCount)
		if qLimitCount != "" {

			if err := r.SetQueryParam("limitCount", qLimitCount); err != nil {
				return err
			}
		}
	}

	if o.ProjectID != nil {

		// query param projectId
		var qrProjectID string

		if o.ProjectID != nil {
			qrProjectID = *o.ProjectID
		}
		qProjectID := qrProjectID
		if qProjectID != "" {

			if err := r.SetQueryParam("projectId", qProjectID); err != nil {
				return err
			}
		}
	}

	if o.Queryfields != nil {

		// query param queryfields
		var qrQueryfields string

		if o.Queryfields != nil {
			qrQueryfields = *o.Queryfields
		}
		qQueryfields := qrQueryfields
		if qQueryfields != "" {

			if err := r.SetQueryParam("queryfields", qQueryfields); err != nil {
				return err
			}
		}
	}

	if o.SinceTs != nil {

		// query param sinceTs
		var qrSinceTs strfmt.DateTime

		if o.SinceTs != nil {
			qrSinceTs = *o.SinceTs
		}
		qSinceTs := qrSinceTs.String()
		if qSinceTs != "" {

			if err := r.SetQueryParam("sinceTs", qSinceTs); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}