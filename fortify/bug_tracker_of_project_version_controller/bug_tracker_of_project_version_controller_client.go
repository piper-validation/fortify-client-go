// Code generated by go-swagger; DO NOT EDIT.

package bug_tracker_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bug tracker of project version controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bug tracker of project version controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ClearBugLinksByExternalIdsBugTrackerOfProjectVersion(params *ClearBugLinksByExternalIdsBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ClearBugLinksByExternalIdsBugTrackerOfProjectVersionOK, error)

	ListBugTrackerOfProjectVersion(params *ListBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListBugTrackerOfProjectVersionOK, error)

	TestBugTrackerOfProjectVersion(params *TestBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*TestBugTrackerOfProjectVersionOK, error)

	UpdateCollectionBugTrackerOfProjectVersion(params *UpdateCollectionBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCollectionBugTrackerOfProjectVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ClearBugLinksByExternalIdsBugTrackerOfProjectVersion clears the specified bug references from the application version does not affect the external bug tracking system in any way
*/
func (a *Client) ClearBugLinksByExternalIdsBugTrackerOfProjectVersion(params *ClearBugLinksByExternalIdsBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ClearBugLinksByExternalIdsBugTrackerOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewClearBugLinksByExternalIdsBugTrackerOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "clearBugLinksByExternalIdsBugTrackerOfProjectVersion",
		Method:             "POST",
		PathPattern:        "/projectVersions/{parentId}/bugtracker/action/clearBugLinksByExternalIds",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ClearBugLinksByExternalIdsBugTrackerOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ClearBugLinksByExternalIdsBugTrackerOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for clearBugLinksByExternalIdsBugTrackerOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBugTrackerOfProjectVersion lists
*/
func (a *Client) ListBugTrackerOfProjectVersion(params *ListBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListBugTrackerOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBugTrackerOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBugTrackerOfProjectVersion",
		Method:             "GET",
		PathPattern:        "/projectVersions/{parentId}/bugtracker",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBugTrackerOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBugTrackerOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBugTrackerOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  TestBugTrackerOfProjectVersion validates that the user can authenticate to the bug tracking system using the provided configuration and credentials
*/
func (a *Client) TestBugTrackerOfProjectVersion(params *TestBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*TestBugTrackerOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewTestBugTrackerOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "testBugTrackerOfProjectVersion",
		Method:             "POST",
		PathPattern:        "/projectVersions/{parentId}/bugtracker/action/test",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &TestBugTrackerOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*TestBugTrackerOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for testBugTrackerOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCollectionBugTrackerOfProjectVersion updates collection
*/
func (a *Client) UpdateCollectionBugTrackerOfProjectVersion(params *UpdateCollectionBugTrackerOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCollectionBugTrackerOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCollectionBugTrackerOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCollectionBugTrackerOfProjectVersion",
		Method:             "PUT",
		PathPattern:        "/projectVersions/{parentId}/bugtracker",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCollectionBugTrackerOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCollectionBugTrackerOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCollectionBugTrackerOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
