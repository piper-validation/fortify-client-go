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

// NewAssignCloudWorkerOfCloudPoolParams creates a new AssignCloudWorkerOfCloudPoolParams object
// with the default values initialized.
func NewAssignCloudWorkerOfCloudPoolParams() *AssignCloudWorkerOfCloudPoolParams {
	var ()
	return &AssignCloudWorkerOfCloudPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewAssignCloudWorkerOfCloudPoolParamsWithTimeout creates a new AssignCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewAssignCloudWorkerOfCloudPoolParamsWithTimeout(timeout time.Duration) *AssignCloudWorkerOfCloudPoolParams {
	var ()
	return &AssignCloudWorkerOfCloudPoolParams{

		timeout: timeout,
	}
}

// NewAssignCloudWorkerOfCloudPoolParamsWithContext creates a new AssignCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewAssignCloudWorkerOfCloudPoolParamsWithContext(ctx context.Context) *AssignCloudWorkerOfCloudPoolParams {
	var ()
	return &AssignCloudWorkerOfCloudPoolParams{

		Context: ctx,
	}
}

// NewAssignCloudWorkerOfCloudPoolParamsWithHTTPClient creates a new AssignCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewAssignCloudWorkerOfCloudPoolParamsWithHTTPClient(client *http.Client) *AssignCloudWorkerOfCloudPoolParams {
	var ()
	return &AssignCloudWorkerOfCloudPoolParams{
		HTTPClient: client,
	}
}

/*AssignCloudWorkerOfCloudPoolParams contains all the parameters to send to the API endpoint
for the assign cloud worker of cloud pool operation typically these are written to a http.Request
*/
type AssignCloudWorkerOfCloudPoolParams struct {

	/*ParentID
	  parentId

	*/
	ParentID string
	/*Resource
	  resource

	*/
	Resource *models.CloudPoolWorkerAssignRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) WithTimeout(timeout time.Duration) *AssignCloudWorkerOfCloudPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) WithContext(ctx context.Context) *AssignCloudWorkerOfCloudPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) WithHTTPClient(client *http.Client) *AssignCloudWorkerOfCloudPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) WithParentID(parentID string) *AssignCloudWorkerOfCloudPoolParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) SetParentID(parentID string) {
	o.ParentID = parentID
}

// WithResource adds the resource to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) WithResource(resource *models.CloudPoolWorkerAssignRequest) *AssignCloudWorkerOfCloudPoolParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the assign cloud worker of cloud pool params
func (o *AssignCloudWorkerOfCloudPoolParams) SetResource(resource *models.CloudPoolWorkerAssignRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *AssignCloudWorkerOfCloudPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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