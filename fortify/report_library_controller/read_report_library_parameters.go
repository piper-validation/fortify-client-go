// Code generated by go-swagger; DO NOT EDIT.

package report_library_controller

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

// NewReadReportLibraryParams creates a new ReadReportLibraryParams object
// with the default values initialized.
func NewReadReportLibraryParams() *ReadReportLibraryParams {
	var ()
	return &ReadReportLibraryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadReportLibraryParamsWithTimeout creates a new ReadReportLibraryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadReportLibraryParamsWithTimeout(timeout time.Duration) *ReadReportLibraryParams {
	var ()
	return &ReadReportLibraryParams{

		timeout: timeout,
	}
}

// NewReadReportLibraryParamsWithContext creates a new ReadReportLibraryParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadReportLibraryParamsWithContext(ctx context.Context) *ReadReportLibraryParams {
	var ()
	return &ReadReportLibraryParams{

		Context: ctx,
	}
}

// NewReadReportLibraryParamsWithHTTPClient creates a new ReadReportLibraryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadReportLibraryParamsWithHTTPClient(client *http.Client) *ReadReportLibraryParams {
	var ()
	return &ReadReportLibraryParams{
		HTTPClient: client,
	}
}

/*ReadReportLibraryParams contains all the parameters to send to the API endpoint
for the read report library operation typically these are written to a http.Request
*/
type ReadReportLibraryParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ID
	  id

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read report library params
func (o *ReadReportLibraryParams) WithTimeout(timeout time.Duration) *ReadReportLibraryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read report library params
func (o *ReadReportLibraryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read report library params
func (o *ReadReportLibraryParams) WithContext(ctx context.Context) *ReadReportLibraryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read report library params
func (o *ReadReportLibraryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read report library params
func (o *ReadReportLibraryParams) WithHTTPClient(client *http.Client) *ReadReportLibraryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read report library params
func (o *ReadReportLibraryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read report library params
func (o *ReadReportLibraryParams) WithFields(fields *string) *ReadReportLibraryParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read report library params
func (o *ReadReportLibraryParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read report library params
func (o *ReadReportLibraryParams) WithID(id int64) *ReadReportLibraryParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read report library params
func (o *ReadReportLibraryParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadReportLibraryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}