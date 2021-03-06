// Code generated by go-swagger; DO NOT EDIT.

package bugfield_template_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"
)

// New creates a new bugfield template group controller API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) ClientService {
	return &Client{transport: transport, formats: formats}
}

/*
Client for bugfield template group controller API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

// ClientService is the interface for Client methods
type ClientService interface {
	CreateBugfieldTemplateGroup(params *CreateBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBugfieldTemplateGroupCreated, error)

	DeleteBugfieldTemplateGroup(params *DeleteBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBugfieldTemplateGroupOK, error)

	ListBugfieldTemplateGroup(params *ListBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ListBugfieldTemplateGroupOK, error)

	MultiDeleteBugfieldTemplateGroup(params *MultiDeleteBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteBugfieldTemplateGroupOK, error)

	ReadBugfieldTemplateGroup(params *ReadBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ReadBugfieldTemplateGroupOK, error)

	UpdateBugfieldTemplateGroup(params *UpdateBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBugfieldTemplateGroupOK, error)

	SetTransport(transport runtime.ClientTransport)
}

/*
  CreateBugfieldTemplateGroup creates
*/
func (a *Client) CreateBugfieldTemplateGroup(params *CreateBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*CreateBugfieldTemplateGroupCreated, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewCreateBugfieldTemplateGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "createBugfieldTemplateGroup",
		Method:             "POST",
		PathPattern:        "/bugfieldTemplateGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &CreateBugfieldTemplateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*CreateBugfieldTemplateGroupCreated)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for createBugfieldTemplateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  DeleteBugfieldTemplateGroup deletes
*/
func (a *Client) DeleteBugfieldTemplateGroup(params *DeleteBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*DeleteBugfieldTemplateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDeleteBugfieldTemplateGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "deleteBugfieldTemplateGroup",
		Method:             "DELETE",
		PathPattern:        "/bugfieldTemplateGroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DeleteBugfieldTemplateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*DeleteBugfieldTemplateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for deleteBugfieldTemplateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ListBugfieldTemplateGroup lists
*/
func (a *Client) ListBugfieldTemplateGroup(params *ListBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ListBugfieldTemplateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewListBugfieldTemplateGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "listBugfieldTemplateGroup",
		Method:             "GET",
		PathPattern:        "/bugfieldTemplateGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ListBugfieldTemplateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ListBugfieldTemplateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for listBugfieldTemplateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  MultiDeleteBugfieldTemplateGroup multis delete
*/
func (a *Client) MultiDeleteBugfieldTemplateGroup(params *MultiDeleteBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*MultiDeleteBugfieldTemplateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewMultiDeleteBugfieldTemplateGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "multiDeleteBugfieldTemplateGroup",
		Method:             "DELETE",
		PathPattern:        "/bugfieldTemplateGroups",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &MultiDeleteBugfieldTemplateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*MultiDeleteBugfieldTemplateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for multiDeleteBugfieldTemplateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  ReadBugfieldTemplateGroup reads
*/
func (a *Client) ReadBugfieldTemplateGroup(params *ReadBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*ReadBugfieldTemplateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewReadBugfieldTemplateGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "readBugfieldTemplateGroup",
		Method:             "GET",
		PathPattern:        "/bugfieldTemplateGroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ReadBugfieldTemplateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*ReadBugfieldTemplateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for readBugfieldTemplateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

/*
  UpdateBugfieldTemplateGroup updates
*/
func (a *Client) UpdateBugfieldTemplateGroup(params *UpdateBugfieldTemplateGroupParams, authInfo runtime.ClientAuthInfoWriter) (*UpdateBugfieldTemplateGroupOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewUpdateBugfieldTemplateGroupParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "updateBugfieldTemplateGroup",
		Method:             "PUT",
		PathPattern:        "/bugfieldTemplateGroups/{id}",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &UpdateBugfieldTemplateGroupReader{formats: a.formats},
		AuthInfo:           authInfo,
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	success, ok := result.(*UpdateBugfieldTemplateGroupOK)
	if ok {
		return success, nil
	}
	// unexpected success response
	// safeguard: normally, absent a default response, unknown success responses return an error above: so this is a codegen issue
	msg := fmt.Sprintf("unexpected success response for updateBugfieldTemplateGroup: API contract not enforced by server. Client expected to get an error, but got: %T", result)
	panic(msg)
}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
