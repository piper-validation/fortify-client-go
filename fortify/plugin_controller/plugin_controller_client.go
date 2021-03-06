// Code generated by go-swagger; DO NOT EDIT.

package plugin_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new plugin controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for plugin controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeletePlugin(params *DeletePluginParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePluginOK, error)

	DisablePlugin(params *DisablePluginParams, authInfo runtime.ClientAuthInfoWriter) (*DisablePluginOK, error)

	EnablePlugin(params *EnablePluginParams, authInfo runtime.ClientAuthInfoWriter) (*EnablePluginOK, error)

	ListPlugin(params *ListPluginParams, authInfo runtime.ClientAuthInfoWriter) (*ListPluginOK, error)

	ReadPlugin(params *ReadPluginParams, authInfo runtime.ClientAuthInfoWriter) (*ReadPluginOK, error)

	UploadPlugin(params *UploadPluginParams, authInfo runtime.ClientAuthInfoWriter) (*UploadPluginCreated, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeletePlugin deletes
*/
func (a *Client) DeletePlugin(params *DeletePluginParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePlugin",
		Method:             "DELETE",
		PathPattern:        "/plugins/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePlugin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DisablePlugin disables a plugin
*/
func (a *Client) DisablePlugin(params *DisablePluginParams, authInfo runtime.ClientAuthInfoWriter) (*DisablePluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDisablePluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "disablePlugin",
		Method:             "POST",
		PathPattern:        "/plugins/action/disable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DisablePluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DisablePluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for disablePlugin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  EnablePlugin enables a plugin
*/
func (a *Client) EnablePlugin(params *EnablePluginParams, authInfo runtime.ClientAuthInfoWriter) (*EnablePluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewEnablePluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "enablePlugin",
		Method:             "POST",
		PathPattern:        "/plugins/action/enable",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &EnablePluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*EnablePluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for enablePlugin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPlugin lists
*/
func (a *Client) ListPlugin(params *ListPluginParams, authInfo runtime.ClientAuthInfoWriter) (*ListPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPlugin",
		Method:             "GET",
		PathPattern:        "/plugins",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPlugin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadPlugin reads
*/
func (a *Client) ReadPlugin(params *ReadPluginParams, authInfo runtime.ClientAuthInfoWriter) (*ReadPluginOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadPluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readPlugin",
		Method:             "GET",
		PathPattern:        "/plugins/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadPluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadPluginOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readPlugin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadPlugin uploads
*/
func (a *Client) UploadPlugin(params *UploadPluginParams, authInfo runtime.ClientAuthInfoWriter) (*UploadPluginCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadPluginParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadPlugin",
		Method:             "POST",
		PathPattern:        "/plugins",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadPluginReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadPluginCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadPlugin: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
