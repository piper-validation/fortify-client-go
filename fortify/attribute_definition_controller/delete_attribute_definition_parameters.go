// Code generated by go-swagger; DO NOT EDIT.

package attribute_definition_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDeleteAttributeDefinitionParams creates a new DeleteAttributeDefinitionParams object
// with the default values initialized.
func NewDeleteAttributeDefinitionParams() *DeleteAttributeDefinitionParams {
	var ()
	return &DeleteAttributeDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteAttributeDefinitionParamsWithTimeout creates a new DeleteAttributeDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteAttributeDefinitionParamsWithTimeout(timeout time.Duration) *DeleteAttributeDefinitionParams {
	var ()
	return &DeleteAttributeDefinitionParams{

		timeout: timeout,
	}
}

// NewDeleteAttributeDefinitionParamsWithContext creates a new DeleteAttributeDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteAttributeDefinitionParamsWithContext(ctx context.Context) *DeleteAttributeDefinitionParams {
	var ()
	return &DeleteAttributeDefinitionParams{

		Context: ctx,
	}
}

// NewDeleteAttributeDefinitionParamsWithHTTPClient creates a new DeleteAttributeDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteAttributeDefinitionParamsWithHTTPClient(client *http.Client) *DeleteAttributeDefinitionParams {
	var ()
	return &DeleteAttributeDefinitionParams{
		HTTPClient: client,
	}
}

/*DeleteAttributeDefinitionParams contains all the parameters to send to the API endpoint
for the delete attribute definition operation typically these are written to a http.Request
*/
type DeleteAttributeDefinitionParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) WithTimeout(timeout time.Duration) *DeleteAttributeDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) WithContext(ctx context.Context) *DeleteAttributeDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) WithHTTPClient(client *http.Client) *DeleteAttributeDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) WithID(id int64) *DeleteAttributeDefinitionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete attribute definition params
func (o *DeleteAttributeDefinitionParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteAttributeDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}