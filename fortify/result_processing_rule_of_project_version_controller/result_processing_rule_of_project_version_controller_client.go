// Code generated by go-swagger; DO NOT EDIT.

package result_processing_rule_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new result processing rule of project version controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for result processing rule of project version controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListResultProcessingRuleOfProjectVersion(params *ListResultProcessingRuleOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListResultProcessingRuleOfProjectVersionOK, error)

	UpdateCollectionResultProcessingRuleOfProjectVersion(params *UpdateCollectionResultProcessingRuleOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCollectionResultProcessingRuleOfProjectVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListResultProcessingRuleOfProjectVersion lists
*/
func (a *Client) ListResultProcessingRuleOfProjectVersion(params *ListResultProcessingRuleOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListResultProcessingRuleOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListResultProcessingRuleOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listResultProcessingRuleOfProjectVersion",
		Method:             "GET",
		PathPattern:        "/projectVersions/{parentId}/resultProcessingRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListResultProcessingRuleOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListResultProcessingRuleOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listResultProcessingRuleOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCollectionResultProcessingRuleOfProjectVersion updates collection
*/
func (a *Client) UpdateCollectionResultProcessingRuleOfProjectVersion(params *UpdateCollectionResultProcessingRuleOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCollectionResultProcessingRuleOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCollectionResultProcessingRuleOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCollectionResultProcessingRuleOfProjectVersion",
		Method:             "PUT",
		PathPattern:        "/projectVersions/{parentId}/resultProcessingRules",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCollectionResultProcessingRuleOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCollectionResultProcessingRuleOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCollectionResultProcessingRuleOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
