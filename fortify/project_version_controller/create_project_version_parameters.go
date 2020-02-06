// Code generated by go-swagger; DO NOT EDIT.

package project_version_controller

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

	"github.com/piper-validation/fortify-client-go/models"
)

// NewCreateProjectVersionParams creates a new CreateProjectVersionParams object
// with the default values initialized.
func NewCreateProjectVersionParams() *CreateProjectVersionParams {
	var ()
	return &CreateProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateProjectVersionParamsWithTimeout creates a new CreateProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateProjectVersionParamsWithTimeout(timeout time.Duration) *CreateProjectVersionParams {
	var ()
	return &CreateProjectVersionParams{

		timeout: timeout,
	}
}

// NewCreateProjectVersionParamsWithContext creates a new CreateProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateProjectVersionParamsWithContext(ctx context.Context) *CreateProjectVersionParams {
	var ()
	return &CreateProjectVersionParams{

		Context: ctx,
	}
}

// NewCreateProjectVersionParamsWithHTTPClient creates a new CreateProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateProjectVersionParamsWithHTTPClient(client *http.Client) *CreateProjectVersionParams {
	var ()
	return &CreateProjectVersionParams{
		HTTPClient: client,
	}
}

/*CreateProjectVersionParams contains all the parameters to send to the API endpoint
for the create project version operation typically these are written to a http.Request
*/
type CreateProjectVersionParams struct {

	/*Resource
	  resource

	*/
	Resource *models.ProjectVersion

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create project version params
func (o *CreateProjectVersionParams) WithTimeout(timeout time.Duration) *CreateProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create project version params
func (o *CreateProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create project version params
func (o *CreateProjectVersionParams) WithContext(ctx context.Context) *CreateProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create project version params
func (o *CreateProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create project version params
func (o *CreateProjectVersionParams) WithHTTPClient(client *http.Client) *CreateProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create project version params
func (o *CreateProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the create project version params
func (o *CreateProjectVersionParams) WithResource(resource *models.ProjectVersion) *CreateProjectVersionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the create project version params
func (o *CreateProjectVersionParams) SetResource(resource *models.ProjectVersion) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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