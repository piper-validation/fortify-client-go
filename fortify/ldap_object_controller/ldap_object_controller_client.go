// Code generated by go-swagger; DO NOT EDIT.

package ldap_object_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new ldap object controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for ldap object controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateLdapObject(params *CreateLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLdapObjectCreated, error)

	DeleteLdapObject(params *DeleteLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLdapObjectOK, error)

	ListLdapObject(params *ListLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*ListLdapObjectOK, error)

	MultiDeleteLdapObject(params *MultiDeleteLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteLdapObjectOK, error)

	ReadLdapObject(params *ReadLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*ReadLdapObjectOK, error)

	RefreshLdapObject(params *RefreshLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*RefreshLdapObjectOK, error)

	UpdateLdapObject(params *UpdateLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLdapObjectOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateLdapObject creates
*/
func (a *Client) CreateLdapObject(params *CreateLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*CreateLdapObjectCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateLdapObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createLdapObject",
		Method:             "POST",
		PathPattern:        "/ldapObjects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateLdapObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateLdapObjectCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createLdapObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteLdapObject deletes
*/
func (a *Client) DeleteLdapObject(params *DeleteLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteLdapObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteLdapObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteLdapObject",
		Method:             "DELETE",
		PathPattern:        "/ldapObjects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteLdapObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteLdapObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteLdapObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListLdapObject lists
*/
func (a *Client) ListLdapObject(params *ListLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*ListLdapObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListLdapObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listLdapObject",
		Method:             "GET",
		PathPattern:        "/ldapObjects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListLdapObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListLdapObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listLdapObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MultiDeleteLdapObject multis delete
*/
func (a *Client) MultiDeleteLdapObject(params *MultiDeleteLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteLdapObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMultiDeleteLdapObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "multiDeleteLdapObject",
		Method:             "DELETE",
		PathPattern:        "/ldapObjects",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MultiDeleteLdapObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MultiDeleteLdapObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for multiDeleteLdapObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadLdapObject reads
*/
func (a *Client) ReadLdapObject(params *ReadLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*ReadLdapObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadLdapObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readLdapObject",
		Method:             "GET",
		PathPattern:        "/ldapObjects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadLdapObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadLdapObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readLdapObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  RefreshLdapObject refreshes the ldap cache
*/
func (a *Client) RefreshLdapObject(params *RefreshLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*RefreshLdapObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewRefreshLdapObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "refreshLdapObject",
		Method:             "POST",
		PathPattern:        "/ldapObjects/action/refresh",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &RefreshLdapObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*RefreshLdapObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for refreshLdapObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateLdapObject updates
*/
func (a *Client) UpdateLdapObject(params *UpdateLdapObjectParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateLdapObjectOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateLdapObjectParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateLdapObject",
		Method:             "PUT",
		PathPattern:        "/ldapObjects/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateLdapObjectReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateLdapObjectOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateLdapObject: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}