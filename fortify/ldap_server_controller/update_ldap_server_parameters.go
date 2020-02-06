// Code generated by go-swagger; DO NOT EDIT.

package ldap_server_controller

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

// NewUpdateLdapServerParams creates a new UpdateLdapServerParams object
// with the default values initialized.
func NewUpdateLdapServerParams() *UpdateLdapServerParams {
	var ()
	return &UpdateLdapServerParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateLdapServerParamsWithTimeout creates a new UpdateLdapServerParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateLdapServerParamsWithTimeout(timeout time.Duration) *UpdateLdapServerParams {
	var ()
	return &UpdateLdapServerParams{

		timeout: timeout,
	}
}

// NewUpdateLdapServerParamsWithContext creates a new UpdateLdapServerParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateLdapServerParamsWithContext(ctx context.Context) *UpdateLdapServerParams {
	var ()
	return &UpdateLdapServerParams{

		Context: ctx,
	}
}

// NewUpdateLdapServerParamsWithHTTPClient creates a new UpdateLdapServerParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateLdapServerParamsWithHTTPClient(client *http.Client) *UpdateLdapServerParams {
	var ()
	return &UpdateLdapServerParams{
		HTTPClient: client,
	}
}

/*UpdateLdapServerParams contains all the parameters to send to the API endpoint
for the update ldap server operation typically these are written to a http.Request
*/
type UpdateLdapServerParams struct {

	/*ID
	  id

	*/
	ID int64
	/*Resource
	  resource

	*/
	Resource *models.LdapServerDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update ldap server params
func (o *UpdateLdapServerParams) WithTimeout(timeout time.Duration) *UpdateLdapServerParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update ldap server params
func (o *UpdateLdapServerParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update ldap server params
func (o *UpdateLdapServerParams) WithContext(ctx context.Context) *UpdateLdapServerParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update ldap server params
func (o *UpdateLdapServerParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update ldap server params
func (o *UpdateLdapServerParams) WithHTTPClient(client *http.Client) *UpdateLdapServerParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update ldap server params
func (o *UpdateLdapServerParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update ldap server params
func (o *UpdateLdapServerParams) WithID(id int64) *UpdateLdapServerParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update ldap server params
func (o *UpdateLdapServerParams) SetID(id int64) {
	o.ID = id
}

// WithResource adds the resource to the update ldap server params
func (o *UpdateLdapServerParams) WithResource(resource *models.LdapServerDto) *UpdateLdapServerParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update ldap server params
func (o *UpdateLdapServerParams) SetResource(resource *models.LdapServerDto) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateLdapServerParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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