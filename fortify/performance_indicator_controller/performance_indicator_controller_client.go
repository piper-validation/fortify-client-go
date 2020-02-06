// Code generated by go-swagger; DO NOT EDIT.

package performance_indicator_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new performance indicator controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for performance indicator controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreatePerformanceIndicator(params *CreatePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePerformanceIndicatorCreated, error)

	DeletePerformanceIndicator(params *DeletePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePerformanceIndicatorOK, error)

	ListPerformanceIndicator(params *ListPerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*ListPerformanceIndicatorOK, error)

	MultiDeletePerformanceIndicator(params *MultiDeletePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeletePerformanceIndicatorOK, error)

	UpdatePerformanceIndicator(params *UpdatePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePerformanceIndicatorOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreatePerformanceIndicator creates
*/
func (a *Client) CreatePerformanceIndicator(params *CreatePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*CreatePerformanceIndicatorCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreatePerformanceIndicatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createPerformanceIndicator",
		Method:             "POST",
		PathPattern:        "/performanceIndicators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreatePerformanceIndicatorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreatePerformanceIndicatorCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createPerformanceIndicator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeletePerformanceIndicator deletes
*/
func (a *Client) DeletePerformanceIndicator(params *DeletePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*DeletePerformanceIndicatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeletePerformanceIndicatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deletePerformanceIndicator",
		Method:             "DELETE",
		PathPattern:        "/performanceIndicators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeletePerformanceIndicatorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeletePerformanceIndicatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deletePerformanceIndicator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListPerformanceIndicator lists
*/
func (a *Client) ListPerformanceIndicator(params *ListPerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*ListPerformanceIndicatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListPerformanceIndicatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listPerformanceIndicator",
		Method:             "GET",
		PathPattern:        "/performanceIndicators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListPerformanceIndicatorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListPerformanceIndicatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listPerformanceIndicator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MultiDeletePerformanceIndicator multis delete
*/
func (a *Client) MultiDeletePerformanceIndicator(params *MultiDeletePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeletePerformanceIndicatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMultiDeletePerformanceIndicatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "multiDeletePerformanceIndicator",
		Method:             "DELETE",
		PathPattern:        "/performanceIndicators",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MultiDeletePerformanceIndicatorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MultiDeletePerformanceIndicatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for multiDeletePerformanceIndicator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdatePerformanceIndicator updates
*/
func (a *Client) UpdatePerformanceIndicator(params *UpdatePerformanceIndicatorParams, authInfo runtime.ClientAuthInfoWriter) (*UpdatePerformanceIndicatorOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdatePerformanceIndicatorParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updatePerformanceIndicator",
		Method:             "PUT",
		PathPattern:        "/performanceIndicators/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdatePerformanceIndicatorReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdatePerformanceIndicatorOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updatePerformanceIndicator: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}