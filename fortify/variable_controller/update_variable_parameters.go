// Code generated by go-swagger; DO NOT EDIT.

package variable_controller

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

// NewUpdateVariableParams creates a new UpdateVariableParams object
// with the default values initialized.
func NewUpdateVariableParams() *UpdateVariableParams {
	var ()
	return &UpdateVariableParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateVariableParamsWithTimeout creates a new UpdateVariableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateVariableParamsWithTimeout(timeout time.Duration) *UpdateVariableParams {
	var ()
	return &UpdateVariableParams{

		timeout: timeout,
	}
}

// NewUpdateVariableParamsWithContext creates a new UpdateVariableParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateVariableParamsWithContext(ctx context.Context) *UpdateVariableParams {
	var ()
	return &UpdateVariableParams{

		Context: ctx,
	}
}

// NewUpdateVariableParamsWithHTTPClient creates a new UpdateVariableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateVariableParamsWithHTTPClient(client *http.Client) *UpdateVariableParams {
	var ()
	return &UpdateVariableParams{
		HTTPClient: client,
	}
}

/*UpdateVariableParams contains all the parameters to send to the API endpoint
for the update variable operation typically these are written to a http.Request
*/
type UpdateVariableParams struct {

	/*ID
	  id

	*/
	ID int64
	/*Resource
	  resource

	*/
	Resource *models.Variable

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update variable params
func (o *UpdateVariableParams) WithTimeout(timeout time.Duration) *UpdateVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update variable params
func (o *UpdateVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update variable params
func (o *UpdateVariableParams) WithContext(ctx context.Context) *UpdateVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update variable params
func (o *UpdateVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update variable params
func (o *UpdateVariableParams) WithHTTPClient(client *http.Client) *UpdateVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update variable params
func (o *UpdateVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update variable params
func (o *UpdateVariableParams) WithID(id int64) *UpdateVariableParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update variable params
func (o *UpdateVariableParams) SetID(id int64) {
	o.ID = id
}

// WithResource adds the resource to the update variable params
func (o *UpdateVariableParams) WithResource(resource *models.Variable) *UpdateVariableParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update variable params
func (o *UpdateVariableParams) SetResource(resource *models.Variable) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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