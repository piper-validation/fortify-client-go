// Code generated by go-swagger; DO NOT EDIT.

package rulepack_core_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new rulepack core controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for rulepack core controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	DeleteRulepackCore(params *DeleteRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRulepackCoreOK, error)

	ListRulepackCore(params *ListRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*ListRulepackCoreOK, error)

	MultiDeleteRulepackCore(params *MultiDeleteRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteRulepackCoreOK, error)

	UploadRulepackCore(params *UploadRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*UploadRulepackCoreOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  DeleteRulepackCore deletes
*/
func (a *Client) DeleteRulepackCore(params *DeleteRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteRulepackCoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteRulepackCoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteRulepackCore",
		Method:             "DELETE",
		PathPattern:        "/coreRulepacks/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteRulepackCoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteRulepackCoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteRulepackCore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListRulepackCore lists
*/
func (a *Client) ListRulepackCore(params *ListRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*ListRulepackCoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListRulepackCoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listRulepackCore",
		Method:             "GET",
		PathPattern:        "/coreRulepacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListRulepackCoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListRulepackCoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listRulepackCore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MultiDeleteRulepackCore multis delete
*/
func (a *Client) MultiDeleteRulepackCore(params *MultiDeleteRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteRulepackCoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMultiDeleteRulepackCoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "multiDeleteRulepackCore",
		Method:             "DELETE",
		PathPattern:        "/coreRulepacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MultiDeleteRulepackCoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MultiDeleteRulepackCoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for multiDeleteRulepackCore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UploadRulepackCore uploads
*/
func (a *Client) UploadRulepackCore(params *UploadRulepackCoreParams, authInfo runtime.ClientAuthInfoWriter) (*UploadRulepackCoreOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUploadRulepackCoreParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "uploadRulepackCore",
		Method:             "POST",
		PathPattern:        "/coreRulepacks",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"multipart/form-data"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UploadRulepackCoreReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UploadRulepackCoreOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for uploadRulepackCore: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
