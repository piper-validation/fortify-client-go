// Code generated by go-swagger; DO NOT EDIT.

package scan_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new scan controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for scan controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ReadScan(params *ReadScanParams, authInfo runtime.ClientAuthInfoWriter) (*ReadScanOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ReadScan reads
*/
func (a *Client) ReadScan(params *ReadScanParams, authInfo runtime.ClientAuthInfoWriter) (*ReadScanOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadScanParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readScan",
		Method:             "GET",
		PathPattern:        "/scans/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadScanReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadScanOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readScan: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
