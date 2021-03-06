// Code generated by go-swagger; DO NOT EDIT.

package auth_token_controller

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

	"github.com/piper-validation/fortify-client-go/models"
)

// NewUpdateAuthTokenParams creates a new UpdateAuthTokenParams object
// with the default values initialized.
func NewUpdateAuthTokenParams() *UpdateAuthTokenParams {
	var ()
	return &UpdateAuthTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateAuthTokenParamsWithTimeout creates a new UpdateAuthTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateAuthTokenParamsWithTimeout(timeout time.Duration) *UpdateAuthTokenParams {
	var ()
	return &UpdateAuthTokenParams{

		timeout: timeout,
	}
}

// NewUpdateAuthTokenParamsWithContext creates a new UpdateAuthTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateAuthTokenParamsWithContext(ctx context.Context) *UpdateAuthTokenParams {
	var ()
	return &UpdateAuthTokenParams{

		Context: ctx,
	}
}

// NewUpdateAuthTokenParamsWithHTTPClient creates a new UpdateAuthTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateAuthTokenParamsWithHTTPClient(client *http.Client) *UpdateAuthTokenParams {
	var ()
	return &UpdateAuthTokenParams{
		HTTPClient: client,
	}
}

/*UpdateAuthTokenParams contains all the parameters to send to the API endpoint
for the update auth token operation typically these are written to a http.Request
*/
type UpdateAuthTokenParams struct {

	/*ID
	  id

	*/
	ID int64
	/*Resource
	  resource

	*/
	Resource *models.AuthenticationToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update auth token params
func (o *UpdateAuthTokenParams) WithTimeout(timeout time.Duration) *UpdateAuthTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update auth token params
func (o *UpdateAuthTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update auth token params
func (o *UpdateAuthTokenParams) WithContext(ctx context.Context) *UpdateAuthTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update auth token params
func (o *UpdateAuthTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update auth token params
func (o *UpdateAuthTokenParams) WithHTTPClient(client *http.Client) *UpdateAuthTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update auth token params
func (o *UpdateAuthTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update auth token params
func (o *UpdateAuthTokenParams) WithID(id int64) *UpdateAuthTokenParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update auth token params
func (o *UpdateAuthTokenParams) SetID(id int64) {
	o.ID = id
}

// WithResource adds the resource to the update auth token params
func (o *UpdateAuthTokenParams) WithResource(resource *models.AuthenticationToken) *UpdateAuthTokenParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update auth token params
func (o *UpdateAuthTokenParams) SetResource(resource *models.AuthenticationToken) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateAuthTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if o.Resource != nil {
		if err := r.SetBodyParam(o.Resource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
