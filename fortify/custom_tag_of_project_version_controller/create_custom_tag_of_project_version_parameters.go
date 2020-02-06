// Code generated by go-swagger; DO NOT EDIT.

package custom_tag_of_project_version_controller

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

// NewCreateCustomTagOfProjectVersionParams creates a new CreateCustomTagOfProjectVersionParams object
// with the default values initialized.
func NewCreateCustomTagOfProjectVersionParams() *CreateCustomTagOfProjectVersionParams {
	var ()
	return &CreateCustomTagOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateCustomTagOfProjectVersionParamsWithTimeout creates a new CreateCustomTagOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateCustomTagOfProjectVersionParamsWithTimeout(timeout time.Duration) *CreateCustomTagOfProjectVersionParams {
	var ()
	return &CreateCustomTagOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewCreateCustomTagOfProjectVersionParamsWithContext creates a new CreateCustomTagOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateCustomTagOfProjectVersionParamsWithContext(ctx context.Context) *CreateCustomTagOfProjectVersionParams {
	var ()
	return &CreateCustomTagOfProjectVersionParams{

		Context: ctx,
	}
}

// NewCreateCustomTagOfProjectVersionParamsWithHTTPClient creates a new CreateCustomTagOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateCustomTagOfProjectVersionParamsWithHTTPClient(client *http.Client) *CreateCustomTagOfProjectVersionParams {
	var ()
	return &CreateCustomTagOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*CreateCustomTagOfProjectVersionParams contains all the parameters to send to the API endpoint
for the create custom tag of project version operation typically these are written to a http.Request
*/
type CreateCustomTagOfProjectVersionParams struct {

	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Resource
	  resource

	*/
	Resource *models.CustomTag

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) WithTimeout(timeout time.Duration) *CreateCustomTagOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) WithContext(ctx context.Context) *CreateCustomTagOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) WithHTTPClient(client *http.Client) *CreateCustomTagOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) WithParentID(parentID int64) *CreateCustomTagOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithResource adds the resource to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) WithResource(resource *models.CustomTag) *CreateCustomTagOfProjectVersionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the create custom tag of project version params
func (o *CreateCustomTagOfProjectVersionParams) SetResource(resource *models.CustomTag) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateCustomTagOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
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