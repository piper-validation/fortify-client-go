// Code generated by go-swagger; DO NOT EDIT.

package cloud_worker_of_cloud_pool_controller

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

// NewReplaceCloudWorkerOfCloudPoolParams creates a new ReplaceCloudWorkerOfCloudPoolParams object
// with the default values initialized.
func NewReplaceCloudWorkerOfCloudPoolParams() *ReplaceCloudWorkerOfCloudPoolParams {
	var ()
	return &ReplaceCloudWorkerOfCloudPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReplaceCloudWorkerOfCloudPoolParamsWithTimeout creates a new ReplaceCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReplaceCloudWorkerOfCloudPoolParamsWithTimeout(timeout time.Duration) *ReplaceCloudWorkerOfCloudPoolParams {
	var ()
	return &ReplaceCloudWorkerOfCloudPoolParams{

		timeout: timeout,
	}
}

// NewReplaceCloudWorkerOfCloudPoolParamsWithContext creates a new ReplaceCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewReplaceCloudWorkerOfCloudPoolParamsWithContext(ctx context.Context) *ReplaceCloudWorkerOfCloudPoolParams {
	var ()
	return &ReplaceCloudWorkerOfCloudPoolParams{

		Context: ctx,
	}
}

// NewReplaceCloudWorkerOfCloudPoolParamsWithHTTPClient creates a new ReplaceCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReplaceCloudWorkerOfCloudPoolParamsWithHTTPClient(client *http.Client) *ReplaceCloudWorkerOfCloudPoolParams {
	var ()
	return &ReplaceCloudWorkerOfCloudPoolParams{
		HTTPClient: client,
	}
}

/*ReplaceCloudWorkerOfCloudPoolParams contains all the parameters to send to the API endpoint
for the replace cloud worker of cloud pool operation typically these are written to a http.Request
*/
type ReplaceCloudWorkerOfCloudPoolParams struct {

	/*ParentID
	  parentId

	*/
	ParentID string
	/*Resource
	  resource

	*/
	Resource *models.CloudPoolWorkerReplaceRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) WithTimeout(timeout time.Duration) *ReplaceCloudWorkerOfCloudPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) WithContext(ctx context.Context) *ReplaceCloudWorkerOfCloudPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) WithHTTPClient(client *http.Client) *ReplaceCloudWorkerOfCloudPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) WithParentID(parentID string) *ReplaceCloudWorkerOfCloudPoolParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) SetParentID(parentID string) {
	o.ParentID = parentID
}

// WithResource adds the resource to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) WithResource(resource *models.CloudPoolWorkerReplaceRequest) *ReplaceCloudWorkerOfCloudPoolParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the replace cloud worker of cloud pool params
func (o *ReplaceCloudWorkerOfCloudPoolParams) SetResource(resource *models.CloudPoolWorkerReplaceRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *ReplaceCloudWorkerOfCloudPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
