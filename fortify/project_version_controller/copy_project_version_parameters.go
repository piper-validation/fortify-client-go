// Code generated by go-swagger; DO NOT EDIT.

package project_version_controller

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

	"github.com/piper-validation/fortify-client-go/models"
)

// NewCopyProjectVersionParams creates a new CopyProjectVersionParams object
// with the default values initialized.
func NewCopyProjectVersionParams() *CopyProjectVersionParams {
	var ()
	return &CopyProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCopyProjectVersionParamsWithTimeout creates a new CopyProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCopyProjectVersionParamsWithTimeout(timeout time.Duration) *CopyProjectVersionParams {
	var ()
	return &CopyProjectVersionParams{

		timeout: timeout,
	}
}

// NewCopyProjectVersionParamsWithContext creates a new CopyProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCopyProjectVersionParamsWithContext(ctx context.Context) *CopyProjectVersionParams {
	var ()
	return &CopyProjectVersionParams{

		Context: ctx,
	}
}

// NewCopyProjectVersionParamsWithHTTPClient creates a new CopyProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCopyProjectVersionParamsWithHTTPClient(client *http.Client) *CopyProjectVersionParams {
	var ()
	return &CopyProjectVersionParams{
		HTTPClient: client,
	}
}

/*CopyProjectVersionParams contains all the parameters to send to the API endpoint
for the copy project version operation typically these are written to a http.Request
*/
type CopyProjectVersionParams struct {

	/*Resource
	  resource

	*/
	Resource *models.ProjectVersionCopyPartialRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the copy project version params
func (o *CopyProjectVersionParams) WithTimeout(timeout time.Duration) *CopyProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the copy project version params
func (o *CopyProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the copy project version params
func (o *CopyProjectVersionParams) WithContext(ctx context.Context) *CopyProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the copy project version params
func (o *CopyProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the copy project version params
func (o *CopyProjectVersionParams) WithHTTPClient(client *http.Client) *CopyProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the copy project version params
func (o *CopyProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the copy project version params
func (o *CopyProjectVersionParams) WithResource(resource *models.ProjectVersionCopyPartialRequest) *CopyProjectVersionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the copy project version params
func (o *CopyProjectVersionParams) SetResource(resource *models.ProjectVersionCopyPartialRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CopyProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

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
