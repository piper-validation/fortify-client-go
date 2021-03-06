// Code generated by go-swagger; DO NOT EDIT.

package bugfield_template_group_controller

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

// NewUpdateBugfieldTemplateGroupParams creates a new UpdateBugfieldTemplateGroupParams object
// with the default values initialized.
func NewUpdateBugfieldTemplateGroupParams() *UpdateBugfieldTemplateGroupParams {
	var ()
	return &UpdateBugfieldTemplateGroupParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateBugfieldTemplateGroupParamsWithTimeout creates a new UpdateBugfieldTemplateGroupParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateBugfieldTemplateGroupParamsWithTimeout(timeout time.Duration) *UpdateBugfieldTemplateGroupParams {
	var ()
	return &UpdateBugfieldTemplateGroupParams{

		timeout: timeout,
	}
}

// NewUpdateBugfieldTemplateGroupParamsWithContext creates a new UpdateBugfieldTemplateGroupParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateBugfieldTemplateGroupParamsWithContext(ctx context.Context) *UpdateBugfieldTemplateGroupParams {
	var ()
	return &UpdateBugfieldTemplateGroupParams{

		Context: ctx,
	}
}

// NewUpdateBugfieldTemplateGroupParamsWithHTTPClient creates a new UpdateBugfieldTemplateGroupParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateBugfieldTemplateGroupParamsWithHTTPClient(client *http.Client) *UpdateBugfieldTemplateGroupParams {
	var ()
	return &UpdateBugfieldTemplateGroupParams{
		HTTPClient: client,
	}
}

/*UpdateBugfieldTemplateGroupParams contains all the parameters to send to the API endpoint
for the update bugfield template group operation typically these are written to a http.Request
*/
type UpdateBugfieldTemplateGroupParams struct {

	/*ID
	  id

	*/
	ID int64
	/*Resource
	  resource

	*/
	Resource *models.BugfieldTemplateGroupDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) WithTimeout(timeout time.Duration) *UpdateBugfieldTemplateGroupParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) WithContext(ctx context.Context) *UpdateBugfieldTemplateGroupParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) WithHTTPClient(client *http.Client) *UpdateBugfieldTemplateGroupParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) WithID(id int64) *UpdateBugfieldTemplateGroupParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) SetID(id int64) {
	o.ID = id
}

// WithResource adds the resource to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) WithResource(resource *models.BugfieldTemplateGroupDto) *UpdateBugfieldTemplateGroupParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update bugfield template group params
func (o *UpdateBugfieldTemplateGroupParams) SetResource(resource *models.BugfieldTemplateGroupDto) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateBugfieldTemplateGroupParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
