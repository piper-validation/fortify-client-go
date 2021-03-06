// Code generated by go-swagger; DO NOT EDIT.

package variable_history_of_project_version_controller

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

// NewReadVariableHistoryOfProjectVersionParams creates a new ReadVariableHistoryOfProjectVersionParams object
// with the default values initialized.
func NewReadVariableHistoryOfProjectVersionParams() *ReadVariableHistoryOfProjectVersionParams {
	var ()
	return &ReadVariableHistoryOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadVariableHistoryOfProjectVersionParamsWithTimeout creates a new ReadVariableHistoryOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadVariableHistoryOfProjectVersionParamsWithTimeout(timeout time.Duration) *ReadVariableHistoryOfProjectVersionParams {
	var ()
	return &ReadVariableHistoryOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewReadVariableHistoryOfProjectVersionParamsWithContext creates a new ReadVariableHistoryOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadVariableHistoryOfProjectVersionParamsWithContext(ctx context.Context) *ReadVariableHistoryOfProjectVersionParams {
	var ()
	return &ReadVariableHistoryOfProjectVersionParams{

		Context: ctx,
	}
}

// NewReadVariableHistoryOfProjectVersionParamsWithHTTPClient creates a new ReadVariableHistoryOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadVariableHistoryOfProjectVersionParamsWithHTTPClient(client *http.Client) *ReadVariableHistoryOfProjectVersionParams {
	var ()
	return &ReadVariableHistoryOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ReadVariableHistoryOfProjectVersionParams contains all the parameters to send to the API endpoint
for the read variable history of project version operation typically these are written to a http.Request
*/
type ReadVariableHistoryOfProjectVersionParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ID
	  id

	*/
	ID string
	/*ParentID
	  parentId

	*/
	ParentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) WithTimeout(timeout time.Duration) *ReadVariableHistoryOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) WithContext(ctx context.Context) *ReadVariableHistoryOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) WithHTTPClient(client *http.Client) *ReadVariableHistoryOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) WithFields(fields *string) *ReadVariableHistoryOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) WithID(id string) *ReadVariableHistoryOfProjectVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) SetID(id string) {
	o.ID = id
}

// WithParentID adds the parentID to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) WithParentID(parentID int64) *ReadVariableHistoryOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the read variable history of project version params
func (o *ReadVariableHistoryOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ReadVariableHistoryOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
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
