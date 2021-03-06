// Code generated by go-swagger; DO NOT EDIT.

package variable_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new variable controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for variable controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateVariable(params *CreateVariableParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVariableCreated, error)

	DeleteVariable(params *DeleteVariableParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteVariableOK, error)

	ListVariable(params *ListVariableParams, authInfo runtime.ClientAuthInfoWriter) (*ListVariableOK, error)

	MultiDeleteVariable(params *MultiDeleteVariableParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteVariableOK, error)

	UpdateVariable(params *UpdateVariableParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateVariableOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateVariable creates
*/
func (a *Client) CreateVariable(params *CreateVariableParams, authInfo runtime.ClientAuthInfoWriter) (*CreateVariableCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateVariableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createVariable",
		Method:             "POST",
		PathPattern:        "/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateVariableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateVariableCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createVariable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteVariable deletes
*/
func (a *Client) DeleteVariable(params *DeleteVariableParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteVariableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteVariableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteVariable",
		Method:             "DELETE",
		PathPattern:        "/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteVariableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteVariable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListVariable lists
*/
func (a *Client) ListVariable(params *ListVariableParams, authInfo runtime.ClientAuthInfoWriter) (*ListVariableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListVariableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listVariable",
		Method:             "GET",
		PathPattern:        "/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListVariableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listVariable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MultiDeleteVariable multis delete
*/
func (a *Client) MultiDeleteVariable(params *MultiDeleteVariableParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteVariableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMultiDeleteVariableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "multiDeleteVariable",
		Method:             "DELETE",
		PathPattern:        "/variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MultiDeleteVariableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MultiDeleteVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for multiDeleteVariable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateVariable updates
*/
func (a *Client) UpdateVariable(params *UpdateVariableParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateVariableOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateVariableParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateVariable",
		Method:             "PUT",
		PathPattern:        "/variables/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateVariableReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateVariableOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateVariable: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
