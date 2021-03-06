// Code generated by go-swagger; DO NOT EDIT.

package permission_required_by_permission_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new permission required by permission controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for permission required by permission controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListPermissionRequiredByPermission(params *ListPermissionRequiredByPermissionParams, authInfo runtime.ClientAuthInfoWriter) (*ListPermissionRequiredByPermissionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListPermissionRequiredByPermission lists
*/
func (a *Client) ListPermissionRequiredByPermission(params *ListPermissionRequiredByPermissionParams, authInfo runtime.ClientAuthInfoWriter) (*ListPermissionRequiredByPermissionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPermissionRequiredByPermissionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPermissionRequiredByPermission",
		Method:             "GET",
		PathPattern:        "/permissions/{parentId}/dependsOn",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPermissionRequiredByPermissionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPermissionRequiredByPermissionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPermissionRequiredByPermission: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
