// Code generated by go-swagger; DO NOT EDIT.

package application_state_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new application state controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for application state controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	GetApplicationState(params *GetApplicationStateParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationStateOK, error)

	PutApplicationState(params *PutApplicationStateParams, authInfo runtime.ClientAuthInfoWriter) (*PutApplicationStateOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  GetApplicationState gets
*/
func (a *Client) GetApplicationState(params *GetApplicationStateParams, authInfo runtime.ClientAuthInfoWriter) (*GetApplicationStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetApplicationStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getApplicationState",
		Method:             "GET",
		PathPattern:        "/applicationState",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetApplicationStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*GetApplicationStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for getApplicationState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  PutApplicationState puts
*/
func (a *Client) PutApplicationState(params *PutApplicationStateParams, authInfo runtime.ClientAuthInfoWriter) (*PutApplicationStateOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPutApplicationStateParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "putApplicationState",
		Method:             "PUT",
		PathPattern:        "/applicationState",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PutApplicationStateReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*PutApplicationStateOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for putApplicationState: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}