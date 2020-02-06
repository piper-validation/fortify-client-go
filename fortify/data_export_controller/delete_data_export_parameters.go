// Code generated by go-swagger; DO NOT EDIT.

package data_export_controller

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

// NewDeleteDataExportParams creates a new DeleteDataExportParams object
// with the default values initialized.
func NewDeleteDataExportParams() *DeleteDataExportParams {
	var ()
	return &DeleteDataExportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDeleteDataExportParamsWithTimeout creates a new DeleteDataExportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDeleteDataExportParamsWithTimeout(timeout time.Duration) *DeleteDataExportParams {
	var ()
	return &DeleteDataExportParams{

		timeout: timeout,
	}
}

// NewDeleteDataExportParamsWithContext creates a new DeleteDataExportParams object
// with the default values initialized, and the ability to set a context for a request
func NewDeleteDataExportParamsWithContext(ctx context.Context) *DeleteDataExportParams {
	var ()
	return &DeleteDataExportParams{

		Context: ctx,
	}
}

// NewDeleteDataExportParamsWithHTTPClient creates a new DeleteDataExportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDeleteDataExportParamsWithHTTPClient(client *http.Client) *DeleteDataExportParams {
	var ()
	return &DeleteDataExportParams{
		HTTPClient: client,
	}
}

/*DeleteDataExportParams contains all the parameters to send to the API endpoint
for the delete data export operation typically these are written to a http.Request
*/
type DeleteDataExportParams struct {

	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the delete data export params
func (o *DeleteDataExportParams) WithTimeout(timeout time.Duration) *DeleteDataExportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the delete data export params
func (o *DeleteDataExportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the delete data export params
func (o *DeleteDataExportParams) WithContext(ctx context.Context) *DeleteDataExportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the delete data export params
func (o *DeleteDataExportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the delete data export params
func (o *DeleteDataExportParams) WithHTTPClient(client *http.Client) *DeleteDataExportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the delete data export params
func (o *DeleteDataExportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the delete data export params
func (o *DeleteDataExportParams) WithID(id int64) *DeleteDataExportParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the delete data export params
func (o *DeleteDataExportParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DeleteDataExportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}