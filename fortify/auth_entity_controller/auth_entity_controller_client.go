// Code generated by go-swagger; DO NOT EDIT.

package auth_entity_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new auth entity controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for auth entity controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListAuthEntity(params *ListAuthEntityParams, authInfo runtime.ClientAuthInfoWriter) (*ListAuthEntityOK, error)

	MultiDeleteAuthEntity(params *MultiDeleteAuthEntityParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteAuthEntityOK, error)

	ReadAuthEntity(params *ReadAuthEntityParams, authInfo runtime.ClientAuthInfoWriter) (*ReadAuthEntityOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListAuthEntity lists
*/
func (a *Client) ListAuthEntity(params *ListAuthEntityParams, authInfo runtime.ClientAuthInfoWriter) (*ListAuthEntityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAuthEntityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAuthEntity",
		Method:             "GET",
		PathPattern:        "/authEntities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAuthEntityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAuthEntityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAuthEntity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MultiDeleteAuthEntity multis delete
*/
func (a *Client) MultiDeleteAuthEntity(params *MultiDeleteAuthEntityParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteAuthEntityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMultiDeleteAuthEntityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "multiDeleteAuthEntity",
		Method:             "DELETE",
		PathPattern:        "/authEntities",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MultiDeleteAuthEntityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MultiDeleteAuthEntityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for multiDeleteAuthEntity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadAuthEntity reads
*/
func (a *Client) ReadAuthEntity(params *ReadAuthEntityParams, authInfo runtime.ClientAuthInfoWriter) (*ReadAuthEntityOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadAuthEntityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readAuthEntity",
		Method:             "GET",
		PathPattern:        "/authEntities/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadAuthEntityReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadAuthEntityOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readAuthEntity: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}