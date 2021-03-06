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

// NewDisableCloudWorkerOfCloudPoolParams creates a new DisableCloudWorkerOfCloudPoolParams object
// with the default values initialized.
func NewDisableCloudWorkerOfCloudPoolParams() *DisableCloudWorkerOfCloudPoolParams {
	var ()
	return &DisableCloudWorkerOfCloudPoolParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDisableCloudWorkerOfCloudPoolParamsWithTimeout creates a new DisableCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDisableCloudWorkerOfCloudPoolParamsWithTimeout(timeout time.Duration) *DisableCloudWorkerOfCloudPoolParams {
	var ()
	return &DisableCloudWorkerOfCloudPoolParams{

		timeout: timeout,
	}
}

// NewDisableCloudWorkerOfCloudPoolParamsWithContext creates a new DisableCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a context for a request
func NewDisableCloudWorkerOfCloudPoolParamsWithContext(ctx context.Context) *DisableCloudWorkerOfCloudPoolParams {
	var ()
	return &DisableCloudWorkerOfCloudPoolParams{

		Context: ctx,
	}
}

// NewDisableCloudWorkerOfCloudPoolParamsWithHTTPClient creates a new DisableCloudWorkerOfCloudPoolParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDisableCloudWorkerOfCloudPoolParamsWithHTTPClient(client *http.Client) *DisableCloudWorkerOfCloudPoolParams {
	var ()
	return &DisableCloudWorkerOfCloudPoolParams{
		HTTPClient: client,
	}
}

/*DisableCloudWorkerOfCloudPoolParams contains all the parameters to send to the API endpoint
for the disable cloud worker of cloud pool operation typically these are written to a http.Request
*/
type DisableCloudWorkerOfCloudPoolParams struct {

	/*ParentID
	  parentId

	*/
	ParentID string
	/*Resource
	  resource

	*/
	Resource *models.CloudPoolWorkerDisableRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) WithTimeout(timeout time.Duration) *DisableCloudWorkerOfCloudPoolParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) WithContext(ctx context.Context) *DisableCloudWorkerOfCloudPoolParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) WithHTTPClient(client *http.Client) *DisableCloudWorkerOfCloudPoolParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) WithParentID(parentID string) *DisableCloudWorkerOfCloudPoolParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) SetParentID(parentID string) {
	o.ParentID = parentID
}

// WithResource adds the resource to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) WithResource(resource *models.CloudPoolWorkerDisableRequest) *DisableCloudWorkerOfCloudPoolParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the disable cloud worker of cloud pool params
func (o *DisableCloudWorkerOfCloudPoolParams) SetResource(resource *models.CloudPoolWorkerDisableRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *DisableCloudWorkerOfCloudPoolParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
