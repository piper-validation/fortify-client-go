// Code generated by go-swagger; DO NOT EDIT.

package issue_selector_set_of_project_version_controller

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

// NewGetIssueSelectorSetOfProjectVersionParams creates a new GetIssueSelectorSetOfProjectVersionParams object
// with the default values initialized.
func NewGetIssueSelectorSetOfProjectVersionParams() *GetIssueSelectorSetOfProjectVersionParams {
	var ()
	return &GetIssueSelectorSetOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIssueSelectorSetOfProjectVersionParamsWithTimeout creates a new GetIssueSelectorSetOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIssueSelectorSetOfProjectVersionParamsWithTimeout(timeout time.Duration) *GetIssueSelectorSetOfProjectVersionParams {
	var ()
	return &GetIssueSelectorSetOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewGetIssueSelectorSetOfProjectVersionParamsWithContext creates a new GetIssueSelectorSetOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIssueSelectorSetOfProjectVersionParamsWithContext(ctx context.Context) *GetIssueSelectorSetOfProjectVersionParams {
	var ()
	return &GetIssueSelectorSetOfProjectVersionParams{

		Context: ctx,
	}
}

// NewGetIssueSelectorSetOfProjectVersionParamsWithHTTPClient creates a new GetIssueSelectorSetOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIssueSelectorSetOfProjectVersionParamsWithHTTPClient(client *http.Client) *GetIssueSelectorSetOfProjectVersionParams {
	var ()
	return &GetIssueSelectorSetOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*GetIssueSelectorSetOfProjectVersionParams contains all the parameters to send to the API endpoint
for the get issue selector set of project version operation typically these are written to a http.Request
*/
type GetIssueSelectorSetOfProjectVersionParams struct {

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

// WithTimeout adds the timeout to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) WithTimeout(timeout time.Duration) *GetIssueSelectorSetOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) WithContext(ctx context.Context) *GetIssueSelectorSetOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) WithHTTPClient(client *http.Client) *GetIssueSelectorSetOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) WithFields(fields *string) *GetIssueSelectorSetOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithParentID adds the parentID to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) WithParentID(parentID int64) *GetIssueSelectorSetOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the get issue selector set of project version params
func (o *GetIssueSelectorSetOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *GetIssueSelectorSetOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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