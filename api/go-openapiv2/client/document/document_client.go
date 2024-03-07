// Code generated by go-swagger; DO NOT EDIT.

package document

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new document API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for document API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientOption is the option for Client methods
type ClientOption func(*runtime.ClientOperation)

// ClientService is the interface for Client methods
type ClientService interface {
	RestColServiceCreateDocument(params *RestColServiceCreateDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceCreateDocumentOK, error)

	RestColServiceCreateDocument2(params *RestColServiceCreateDocument2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceCreateDocument2OK, error)

	RestColServiceDeleteDocument(params *RestColServiceDeleteDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceDeleteDocumentOK, error)

	RestColServiceDeleteDocument2(params *RestColServiceDeleteDocument2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceDeleteDocument2OK, error)

	RestColServiceGetDocument(params *RestColServiceGetDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceGetDocumentOK, error)

	RestColServiceGetDocument2(params *RestColServiceGetDocument2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceGetDocument2OK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
RestColServiceCreateDocument create a document to the collection
*/
func (a *Client) RestColServiceCreateDocument(params *RestColServiceCreateDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceCreateDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestColServiceCreateDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestColService_CreateDocument",
		Method:             "POST",
		PathPattern:        "/v1/collections/{cid}:add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestColServiceCreateDocumentReader{formats: a.formats},
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
	success, ok := result.(*RestColServiceCreateDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestColServiceCreateDocumentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestColServiceCreateDocument2 create a document to the collection
*/
func (a *Client) RestColServiceCreateDocument2(params *RestColServiceCreateDocument2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceCreateDocument2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestColServiceCreateDocument2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestColService_CreateDocument2",
		Method:             "POST",
		PathPattern:        "/v1/projects/{pid}/collections/{cid}:add",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestColServiceCreateDocument2Reader{formats: a.formats},
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
	success, ok := result.(*RestColServiceCreateDocument2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestColServiceCreateDocument2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestColServiceDeleteDocument deletes document endpoint is a generic endpoint for deleting a specific data

Remove the specific document from the collection
*/
func (a *Client) RestColServiceDeleteDocument(params *RestColServiceDeleteDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceDeleteDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestColServiceDeleteDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestColService_DeleteDocument",
		Method:             "DELETE",
		PathPattern:        "/v1/collections/{cid}/{did}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestColServiceDeleteDocumentReader{formats: a.formats},
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
	success, ok := result.(*RestColServiceDeleteDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestColServiceDeleteDocumentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestColServiceDeleteDocument2 deletes document endpoint is a generic endpoint for deleting a specific data

Remove the specific document from the collection
*/
func (a *Client) RestColServiceDeleteDocument2(params *RestColServiceDeleteDocument2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceDeleteDocument2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestColServiceDeleteDocument2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestColService_DeleteDocument2",
		Method:             "DELETE",
		PathPattern:        "/v1/projects/{pid}/collections/{cid}/{did}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestColServiceDeleteDocument2Reader{formats: a.formats},
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
	success, ok := result.(*RestColServiceDeleteDocument2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestColServiceDeleteDocument2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestColServiceGetDocument gets document endpoint is a generic endpoint for retrieving data across multiple collections

retrieve a document information from the collection.
*/
func (a *Client) RestColServiceGetDocument(params *RestColServiceGetDocumentParams, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceGetDocumentOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestColServiceGetDocumentParams()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestColService_GetDocument",
		Method:             "GET",
		PathPattern:        "/v1/collections/{cid}/{did}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestColServiceGetDocumentReader{formats: a.formats},
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
	success, ok := result.(*RestColServiceGetDocumentOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestColServiceGetDocumentDefault)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

/*
RestColServiceGetDocument2 gets document endpoint is a generic endpoint for retrieving data across multiple collections

retrieve a document information from the collection.
*/
func (a *Client) RestColServiceGetDocument2(params *RestColServiceGetDocument2Params, authInfo runtime.ClientAuthInfoWriter, opts ...ClientOption) (*RestColServiceGetDocument2OK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRestColServiceGetDocument2Params()
	}
	op := &runtime.ClientOperation{
		ID:                 "RestColService_GetDocument2",
		Method:             "GET",
		PathPattern:        "/v1/projects/{pid}/collections/{cid}/{did}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &RestColServiceGetDocument2Reader{formats: a.formats},
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
	success, ok := result.(*RestColServiceGetDocument2OK)
	if ok {
		return success, nil
	}
	// unexpected success response
	unexpectedSuccess := result.(*RestColServiceGetDocument2Default)
	return nil, runtime.NewAPIError("unexpected success response: content available as default response in error", unexpectedSuccess, unexpectedSuccess.Code())
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
