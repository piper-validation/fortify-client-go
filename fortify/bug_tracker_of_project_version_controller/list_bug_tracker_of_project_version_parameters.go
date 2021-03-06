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
)

// NewListBugTrackerOfProjectVersionParams creates a new ListBugTrackerOfProjectVersionParams object
// with the default values initialized.
func NewListBugTrackerOfProjectVersionParams() *ListBugTrackerOfProjectVersionParams {
	var ()
	return &ListBugTrackerOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListBugTrackerOfProjectVersionParamsWithTimeout creates a new ListBugTrackerOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListBugTrackerOfProjectVersionParamsWithTimeout(timeout time.Duration) *ListBugTrackerOfProjectVersionParams {
	var ()
	return &ListBugTrackerOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewListBugTrackerOfProjectVersionParamsWithContext creates a new ListBugTrackerOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListBugTrackerOfProjectVersionParamsWithContext(ctx context.Context) *ListBugTrackerOfProjectVersionParams {
	var ()
	return &ListBugTrackerOfProjectVersionParams{

		Context: ctx,
	}
}

// NewListBugTrackerOfProjectVersionParamsWithHTTPClient creates a new ListBugTrackerOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListBugTrackerOfProjectVersionParamsWithHTTPClient(client *http.Client) *ListBugTrackerOfProjectVersionParams {
	var ()
	return &ListBugTrackerOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ListBugTrackerOfProjectVersionParams contains all the parameters to send to the API endpoint
for the list bug tracker of project version operation typically these are written to a http.Request
*/
type ListBugTrackerOfProjectVersionParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ParentID
	  parentId

	*/
	ParentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) WithTimeout(timeout time.Duration) *ListBugTrackerOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) WithContext(ctx context.Context) *ListBugTrackerOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) WithHTTPClient(client *http.Client) *ListBugTrackerOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) WithFields(fields *string) *ListBugTrackerOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithParentID adds the parentID to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) WithParentID(parentID int64) *ListBugTrackerOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list bug tracker of project version params
func (o *ListBugTrackerOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ListBugTrackerOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
