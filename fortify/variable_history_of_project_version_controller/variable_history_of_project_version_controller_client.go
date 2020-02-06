// Code generated by go-swagger; DO NOT EDIT.

package variable_history_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new variable history of project version controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for variable history of project version controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListVariableHistoryOfProjectVersion(params *ListVariableHistoryOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListVariableHistoryOfProjectVersionOK, error)

	ReadVariableHistoryOfProjectVersion(params *ReadVariableHistoryOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ReadVariableHistoryOfProjectVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListVariableHistoryOfProjectVersion lists
*/
func (a *Client) ListVariableHistoryOfProjectVersion(params *ListVariableHistoryOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListVariableHistoryOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVariableHistoryOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVariableHistoryOfProjectVersion",
		Method:             "GET",
		PathPattern:        "/projectVersions/{parentId}/variableHistories",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVariableHistoryOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVariableHistoryOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listVariableHistoryOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadVariableHistoryOfProjectVersion reads
*/
func (a *Client) ReadVariableHistoryOfProjectVersion(params *ReadVariableHistoryOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ReadVariableHistoryOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadVariableHistoryOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readVariableHistoryOfProjectVersion",
		Method:             "GET",
		PathPattern:        "/projectVersions/{parentId}/variableHistories/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadVariableHistoryOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadVariableHistoryOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readVariableHistoryOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}