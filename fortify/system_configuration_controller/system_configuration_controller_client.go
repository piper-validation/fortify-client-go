// Code generated by go-swagger; DO NOT EDIT.

package system_configuration_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new system configuration controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for system configuration controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListSystemConfiguration(params *ListSystemConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ListSystemConfigurationOK, error)

	ReadSystemConfiguration(params *ReadSystemConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ReadSystemConfigurationOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListSystemConfiguration lists
*/
func (a *Client) ListSystemConfiguration(params *ListSystemConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ListSystemConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListSystemConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listSystemConfiguration",
		Method:             "GET",
		PathPattern:        "/systemConfiguration",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListSystemConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListSystemConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listSystemConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadSystemConfiguration reads
*/
func (a *Client) ReadSystemConfiguration(params *ReadSystemConfigurationParams, authInfo runtime.ClientAuthInfoWriter) (*ReadSystemConfigurationOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadSystemConfigurationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readSystemConfiguration",
		Method:             "GET",
		PathPattern:        "/systemConfiguration/{name}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadSystemConfigurationReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadSystemConfigurationOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readSystemConfiguration: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}