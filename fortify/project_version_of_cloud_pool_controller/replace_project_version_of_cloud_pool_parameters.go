// Code generated by go-swagger; DO NOT EDIT.

package project_version_of_cloud_pool_controller

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

// NewReplaceProjectVersionOfCloudPoolParams creates a new ReplaceProjectVersionOfCloudPoolParams object
// with the default values initialized.
func NewReplaceProjectVersionOfCloudPoolParams() *ReplaceProjectVersionOfCloudPoolParams {
	var ()
	return &ReplaceProjectVersionOfCloudPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceProjectVersionOfCloudPoolParamsWithTimeout creates a new ReplaceProjectVersionOfCloudPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceProjectVersionOfCloudPoolParamsWithTimeout(timeout time.Duration) *ReplaceProjectVersionOfCloudPoolParams {
	var ()
	return &ReplaceProjectVersionOfCloudPoolParams{

		timeout: timeout,
	}
}

// NewReplaceProjectVersionOfCloudPoolParamsWithContext creates a new ReplaceProjectVersionOfCloudPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceProjectVersionOfCloudPoolParamsWithContext(ctx context.Context) *ReplaceProjectVersionOfCloudPoolParams {
	var ()
	return &ReplaceProjectVersionOfCloudPoolParams{

		Context: ctx,
	}
}

// NewReplaceProjectVersionOfCloudPoolParamsWithHTTPClient creates a new ReplaceProjectVersionOfCloudPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceProjectVersionOfCloudPoolParamsWithHTTPClient(client *http.Client) *ReplaceProjectVersionOfCloudPoolParams {
	var ()
	return &ReplaceProjectVersionOfCloudPoolParams{
		HTTPClient: client,
	}
}

/*ReplaceProjectVersionOfCloudPoolParams contains all the parameters to send to the API endpoint
for the replace project version of cloud pool operation typically these are written to a http.Request
*/
type ReplaceProjectVersionOfCloudPoolParams struct {

	/*ParentID
	  parentId

	*/
	ParentID string
	/*Resource
	  resource

	*/
	Resource *models.CloudPoolProjectVersionReplaceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) WithTimeout(timeout time.Duration) *ReplaceProjectVersionOfCloudPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) WithContext(ctx context.Context) *ReplaceProjectVersionOfCloudPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) WithHTTPClient(client *http.Client) *ReplaceProjectVersionOfCloudPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) WithParentID(parentID string) *ReplaceProjectVersionOfCloudPoolParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) SetParentID(parentID string) {
	o.ParentID = parentID
}

// WithResource adds the resource to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) WithResource(resource *models.CloudPoolProjectVersionReplaceRequest) *ReplaceProjectVersionOfCloudPoolParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the replace project version of cloud pool params
func (o *ReplaceProjectVersionOfCloudPoolParams) SetResource(resource *models.CloudPoolProjectVersionReplaceRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceProjectVersionOfCloudPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param parentId
	if err := r.SetPathParam("parentId", o.ParentID); err != nil {
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
