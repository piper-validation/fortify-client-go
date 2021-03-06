// Code generated by go-swagger; DO NOT EDIT.

package role_controller

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
)

// NewReadRoleParams creates a new ReadRoleParams object
// with the default values initialized.
func NewReadRoleParams() *ReadRoleParams {
	var ()
	return &ReadRoleParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewReadRoleParamsWithTimeout creates a new ReadRoleParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewReadRoleParamsWithTimeout(timeout time.Duration) *ReadRoleParams {
	var ()
	return &ReadRoleParams{

		timeout: timeout,
	}
}

// NewReadRoleParamsWithContext creates a new ReadRoleParams object
// with the default values initialized, and the ability to set a context for a request
func NewReadRoleParamsWithContext(ctx context.Context) *ReadRoleParams {
	var ()
	return &ReadRoleParams{

		Context: ctx,
	}
}

// NewReadRoleParamsWithHTTPClient creates a new ReadRoleParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewReadRoleParamsWithHTTPClient(client *http.Client) *ReadRoleParams {
	var ()
	return &ReadRoleParams{
		HTTPClient: client,
	}
}

/*ReadRoleParams contains all the parameters to send to the API endpoint
for the read role operation typically these are written to a http.Request
*/
type ReadRoleParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ID
	  id

	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the read role params
func (o *ReadRoleParams) WithTimeout(timeout time.Duration) *ReadRoleParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the read role params
func (o *ReadRoleParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the read role params
func (o *ReadRoleParams) WithContext(ctx context.Context) *ReadRoleParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the read role params
func (o *ReadRoleParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the read role params
func (o *ReadRoleParams) WithHTTPClient(client *http.Client) *ReadRoleParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the read role params
func (o *ReadRoleParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the read role params
func (o *ReadRoleParams) WithFields(fields *string) *ReadRoleParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the read role params
func (o *ReadRoleParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithID adds the id to the read role params
func (o *ReadRoleParams) WithID(id string) *ReadRoleParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the read role params
func (o *ReadRoleParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *ReadRoleParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
