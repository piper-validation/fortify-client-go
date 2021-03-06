// Code generated by go-swagger; DO NOT EDIT.

package performance_indicator_history_of_project_version_controller

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

// NewReadPerformanceIndicatorHistoryOfProjectVersionParams creates a new ReadPerformanceIndicatorHistoryOfProjectVersionParams object
// with the default values initialized.
func NewReadPerformanceIndicatorHistoryOfProjectVersionParams() *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	var ()
	return &ReadPerformanceIndicatorHistoryOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionParamsWithTimeout creates a new ReadPerformanceIndicatorHistoryOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadPerformanceIndicatorHistoryOfProjectVersionParamsWithTimeout(timeout time.Duration) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	var ()
	return &ReadPerformanceIndicatorHistoryOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionParamsWithContext creates a new ReadPerformanceIndicatorHistoryOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadPerformanceIndicatorHistoryOfProjectVersionParamsWithContext(ctx context.Context) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	var ()
	return &ReadPerformanceIndicatorHistoryOfProjectVersionParams{

		Context: ctx,
	}
}

// NewReadPerformanceIndicatorHistoryOfProjectVersionParamsWithHTTPClient creates a new ReadPerformanceIndicatorHistoryOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadPerformanceIndicatorHistoryOfProjectVersionParamsWithHTTPClient(client *http.Client) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	var ()
	return &ReadPerformanceIndicatorHistoryOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ReadPerformanceIndicatorHistoryOfProjectVersionParams contains all the parameters to send to the API endpoint
for the read performance indicator history of project version operation typically these are written to a http.Request
*/
type ReadPerformanceIndicatorHistoryOfProjectVersionParams struct {

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

// WithTimeout adds the timeout to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) WithTimeout(timeout time.Duration) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) WithContext(ctx context.Context) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) WithHTTPClient(client *http.Client) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) WithFields(fields *string) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) WithID(id string) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) SetID(id string) {
	o.ID = id
}

// WithParentID adds the parentID to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) WithParentID(parentID int64) *ReadPerformanceIndicatorHistoryOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the read performance indicator history of project version params
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ReadPerformanceIndicatorHistoryOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
