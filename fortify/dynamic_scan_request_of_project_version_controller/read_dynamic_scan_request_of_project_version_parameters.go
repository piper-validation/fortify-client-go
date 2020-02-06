// Code generated by go-swagger; DO NOT EDIT.

package dynamic_scan_request_of_project_version_controller

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

// NewReadDynamicScanRequestOfProjectVersionParams creates a new ReadDynamicScanRequestOfProjectVersionParams object
// with the default values initialized.
func NewReadDynamicScanRequestOfProjectVersionParams() *ReadDynamicScanRequestOfProjectVersionParams {
	var ()
	return &ReadDynamicScanRequestOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadDynamicScanRequestOfProjectVersionParamsWithTimeout creates a new ReadDynamicScanRequestOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadDynamicScanRequestOfProjectVersionParamsWithTimeout(timeout time.Duration) *ReadDynamicScanRequestOfProjectVersionParams {
	var ()
	return &ReadDynamicScanRequestOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewReadDynamicScanRequestOfProjectVersionParamsWithContext creates a new ReadDynamicScanRequestOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadDynamicScanRequestOfProjectVersionParamsWithContext(ctx context.Context) *ReadDynamicScanRequestOfProjectVersionParams {
	var ()
	return &ReadDynamicScanRequestOfProjectVersionParams{

		Context: ctx,
	}
}

// NewReadDynamicScanRequestOfProjectVersionParamsWithHTTPClient creates a new ReadDynamicScanRequestOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadDynamicScanRequestOfProjectVersionParamsWithHTTPClient(client *http.Client) *ReadDynamicScanRequestOfProjectVersionParams {
	var ()
	return &ReadDynamicScanRequestOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ReadDynamicScanRequestOfProjectVersionParams contains all the parameters to send to the API endpoint
for the read dynamic scan request of project version operation typically these are written to a http.Request
*/
type ReadDynamicScanRequestOfProjectVersionParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ID
	  id

	*/
	ID int64
	/*ParentID
	  parentId

	*/
	ParentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) WithTimeout(timeout time.Duration) *ReadDynamicScanRequestOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) WithContext(ctx context.Context) *ReadDynamicScanRequestOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) WithHTTPClient(client *http.Client) *ReadDynamicScanRequestOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) WithFields(fields *string) *ReadDynamicScanRequestOfProjectVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) WithID(id int64) *ReadDynamicScanRequestOfProjectVersionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) SetID(id int64) {
	o.ID = id
}

// WithParentID adds the parentID to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) WithParentID(parentID int64) *ReadDynamicScanRequestOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the read dynamic scan request of project version params
func (o *ReadDynamicScanRequestOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ReadDynamicScanRequestOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
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