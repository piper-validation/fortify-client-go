// Code generated by go-swagger; DO NOT EDIT.

package alert_history_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new alert history controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for alert history controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListAlertHistory(params *ListAlertHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*ListAlertHistoryOK, error)

	SetStatusForAlertHistory(params *SetStatusForAlertHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*SetStatusForAlertHistoryOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListAlertHistory lists
*/
func (a *Client) ListAlertHistory(params *ListAlertHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*ListAlertHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListAlertHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listAlertHistory",
		Method:             "GET",
		PathPattern:        "/alerts",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListAlertHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListAlertHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listAlertHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  SetStatusForAlertHistory marks a triggered alert event as read or unread
*/
func (a *Client) SetStatusForAlertHistory(params *SetStatusForAlertHistoryParams, authInfo runtime.ClientAuthInfoWriter) (*SetStatusForAlertHistoryOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetStatusForAlertHistoryParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setStatusForAlertHistory",
		Method:             "POST",
		PathPattern:        "/alerts/action/setStatus",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetStatusForAlertHistoryReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*SetStatusForAlertHistoryOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for setStatusForAlertHistory: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
