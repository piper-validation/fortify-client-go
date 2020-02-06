// Code generated by go-swagger; DO NOT EDIT.

package plugin_controller

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

// NewReadPluginParams creates a new ReadPluginParams object
// with the default values initialized.
func NewReadPluginParams() *ReadPluginParams {
	var ()
	return &ReadPluginParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadPluginParamsWithTimeout creates a new ReadPluginParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadPluginParamsWithTimeout(timeout time.Duration) *ReadPluginParams {
	var ()
	return &ReadPluginParams{

		timeout: timeout,
	}
}

// NewReadPluginParamsWithContext creates a new ReadPluginParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadPluginParamsWithContext(ctx context.Context) *ReadPluginParams {
	var ()
	return &ReadPluginParams{

		Context: ctx,
	}
}

// NewReadPluginParamsWithHTTPClient creates a new ReadPluginParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadPluginParamsWithHTTPClient(client *http.Client) *ReadPluginParams {
	var ()
	return &ReadPluginParams{
		HTTPClient: client,
	}
}

/*ReadPluginParams contains all the parameters to send to the API endpoint
for the read plugin operation typically these are written to a http.Request
*/
type ReadPluginParams struct {

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

// WithTimeout adds the timeout to the read plugin params
func (o *ReadPluginParams) WithTimeout(timeout time.Duration) *ReadPluginParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read plugin params
func (o *ReadPluginParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read plugin params
func (o *ReadPluginParams) WithContext(ctx context.Context) *ReadPluginParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read plugin params
func (o *ReadPluginParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read plugin params
func (o *ReadPluginParams) WithHTTPClient(client *http.Client) *ReadPluginParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read plugin params
func (o *ReadPluginParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read plugin params
func (o *ReadPluginParams) WithFields(fields *string) *ReadPluginParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read plugin params
func (o *ReadPluginParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read plugin params
func (o *ReadPluginParams) WithID(id int64) *ReadPluginParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read plugin params
func (o *ReadPluginParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadPluginParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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