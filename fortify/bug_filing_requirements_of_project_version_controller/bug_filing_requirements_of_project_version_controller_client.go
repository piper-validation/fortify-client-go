// Code generated by go-swagger; DO NOT EDIT.

package bug_filing_requirements_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bug filing requirements of project version controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bug filing requirements of project version controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	ListBugFilingRequirementsOfProjectVersion(params *ListBugFilingRequirementsOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListBugFilingRequirementsOfProjectVersionOK, error)

	LoginBugFilingRequirementsOfProjectVersion(params *LoginBugFilingRequirementsOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*LoginBugFilingRequirementsOfProjectVersionOK, error)

	UpdateCollectionBugFilingRequirementsOfProjectVersion(params *UpdateCollectionBugFilingRequirementsOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCollectionBugFilingRequirementsOfProjectVersionOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  ListBugFilingRequirementsOfProjectVersion lists
*/
func (a *Client) ListBugFilingRequirementsOfProjectVersion(params *ListBugFilingRequirementsOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*ListBugFilingRequirementsOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBugFilingRequirementsOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBugFilingRequirementsOfProjectVersion",
		Method:             "GET",
		PathPattern:        "/projectVersions/{parentId}/bugfilingrequirements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBugFilingRequirementsOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBugFilingRequirementsOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBugFilingRequirementsOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  LoginBugFilingRequirementsOfProjectVersion authenticates to the bug tracking system and return the initial set of bug filing requirements
*/
func (a *Client) LoginBugFilingRequirementsOfProjectVersion(params *LoginBugFilingRequirementsOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*LoginBugFilingRequirementsOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoginBugFilingRequirementsOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loginBugFilingRequirementsOfProjectVersion",
		Method:             "POST",
		PathPattern:        "/projectVersions/{parentId}/bugfilingrequirements/action/login",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoginBugFilingRequirementsOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*LoginBugFilingRequirementsOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for loginBugFilingRequirementsOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateCollectionBugFilingRequirementsOfProjectVersion updates collection
*/
func (a *Client) UpdateCollectionBugFilingRequirementsOfProjectVersion(params *UpdateCollectionBugFilingRequirementsOfProjectVersionParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateCollectionBugFilingRequirementsOfProjectVersionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateCollectionBugFilingRequirementsOfProjectVersionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateCollectionBugFilingRequirementsOfProjectVersion",
		Method:             "PUT",
		PathPattern:        "/projectVersions/{parentId}/bugfilingrequirements",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateCollectionBugFilingRequirementsOfProjectVersionReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateCollectionBugFilingRequirementsOfProjectVersionOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateCollectionBugFilingRequirementsOfProjectVersion: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}