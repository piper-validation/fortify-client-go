// Code generated by go-swagger; DO NOT EDIT.

package bug_tracker_of_project_version_controller

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

// NewTestBugTrackerOfProjectVersionParams creates a new TestBugTrackerOfProjectVersionParams object
// with the default values initialized.
func NewTestBugTrackerOfProjectVersionParams() *TestBugTrackerOfProjectVersionParams {
	var ()
	return &TestBugTrackerOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewTestBugTrackerOfProjectVersionParamsWithTimeout creates a new TestBugTrackerOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewTestBugTrackerOfProjectVersionParamsWithTimeout(timeout time.Duration) *TestBugTrackerOfProjectVersionParams {
	var ()
	return &TestBugTrackerOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewTestBugTrackerOfProjectVersionParamsWithContext creates a new TestBugTrackerOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewTestBugTrackerOfProjectVersionParamsWithContext(ctx context.Context) *TestBugTrackerOfProjectVersionParams {
	var ()
	return &TestBugTrackerOfProjectVersionParams{

		Context: ctx,
	}
}

// NewTestBugTrackerOfProjectVersionParamsWithHTTPClient creates a new TestBugTrackerOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewTestBugTrackerOfProjectVersionParamsWithHTTPClient(client *http.Client) *TestBugTrackerOfProjectVersionParams {
	var ()
	return &TestBugTrackerOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*TestBugTrackerOfProjectVersionParams contains all the parameters to send to the API endpoint
for the test bug tracker of project version operation typically these are written to a http.Request
*/
type TestBugTrackerOfProjectVersionParams struct {

	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Resource
	  resource

	*/
	Resource *models.BugTrackerRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) WithTimeout(timeout time.Duration) *TestBugTrackerOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) WithContext(ctx context.Context) *TestBugTrackerOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) WithHTTPClient(client *http.Client) *TestBugTrackerOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) WithParentID(parentID int64) *TestBugTrackerOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithResource adds the resource to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) WithResource(resource *models.BugTrackerRequest) *TestBugTrackerOfProjectVersionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the test bug tracker of project version params
func (o *TestBugTrackerOfProjectVersionParams) SetResource(resource *models.BugTrackerRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *TestBugTrackerOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
