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

	strfmt "github.com/go-openapi/strfmt"
)

// NewMultiDeleteAttributeDefinitionParams creates a new MultiDeleteAttributeDefinitionParams object
// with the default values initialized.
func NewMultiDeleteAttributeDefinitionParams() *MultiDeleteAttributeDefinitionParams {
	var ()
	return &MultiDeleteAttributeDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMultiDeleteAttributeDefinitionParamsWithTimeout creates a new MultiDeleteAttributeDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMultiDeleteAttributeDefinitionParamsWithTimeout(timeout time.Duration) *MultiDeleteAttributeDefinitionParams {
	var ()
	return &MultiDeleteAttributeDefinitionParams{

		timeout: timeout,
	}
}

// NewMultiDeleteAttributeDefinitionParamsWithContext creates a new MultiDeleteAttributeDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewMultiDeleteAttributeDefinitionParamsWithContext(ctx context.Context) *MultiDeleteAttributeDefinitionParams {
	var ()
	return &MultiDeleteAttributeDefinitionParams{

		Context: ctx,
	}
}

// NewMultiDeleteAttributeDefinitionParamsWithHTTPClient creates a new MultiDeleteAttributeDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMultiDeleteAttributeDefinitionParamsWithHTTPClient(client *http.Client) *MultiDeleteAttributeDefinitionParams {
	var ()
	return &MultiDeleteAttributeDefinitionParams{
		HTTPClient: client,
	}
}

/*MultiDeleteAttributeDefinitionParams contains all the parameters to send to the API endpoint
for the multi delete attribute definition operation typically these are written to a http.Request
*/
type MultiDeleteAttributeDefinitionParams struct {

	/*Ids
	  A comma-separated list of resource identifiers

	*/
	Ids string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) WithTimeout(timeout time.Duration) *MultiDeleteAttributeDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) WithContext(ctx context.Context) *MultiDeleteAttributeDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) WithHTTPClient(client *http.Client) *MultiDeleteAttributeDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithIds adds the ids to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) WithIds(ids string) *MultiDeleteAttributeDefinitionParams {
	o.SetIds(ids)
	return o
}

// SetIds adds the ids to the multi delete attribute definition params
func (o *MultiDeleteAttributeDefinitionParams) SetIds(ids string) {
	o.Ids = ids
}

// WriteToRequest writes these params to a swagger request
func (o *MultiDeleteAttributeDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param ids
	qrIds := o.Ids
	qIds := qrIds
	if qIds != "" {
		if err := r.SetQueryParam("ids", qIds); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
