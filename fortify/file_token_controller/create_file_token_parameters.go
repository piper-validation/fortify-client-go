// Code generated by go-swagger; DO NOT EDIT.

package file_token_controller

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

// NewCreateFileTokenParams creates a new CreateFileTokenParams object
// with the default values initialized.
func NewCreateFileTokenParams() *CreateFileTokenParams {
	var ()
	return &CreateFileTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateFileTokenParamsWithTimeout creates a new CreateFileTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateFileTokenParamsWithTimeout(timeout time.Duration) *CreateFileTokenParams {
	var ()
	return &CreateFileTokenParams{

		timeout: timeout,
	}
}

// NewCreateFileTokenParamsWithContext creates a new CreateFileTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateFileTokenParamsWithContext(ctx context.Context) *CreateFileTokenParams {
	var ()
	return &CreateFileTokenParams{

		Context: ctx,
	}
}

// NewCreateFileTokenParamsWithHTTPClient creates a new CreateFileTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateFileTokenParamsWithHTTPClient(client *http.Client) *CreateFileTokenParams {
	var ()
	return &CreateFileTokenParams{
		HTTPClient: client,
	}
}

/*CreateFileTokenParams contains all the parameters to send to the API endpoint
for the create file token operation typically these are written to a http.Request
*/
type CreateFileTokenParams struct {

	/*Resource
	  resource

	*/
	Resource *models.FileToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create file token params
func (o *CreateFileTokenParams) WithTimeout(timeout time.Duration) *CreateFileTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create file token params
func (o *CreateFileTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create file token params
func (o *CreateFileTokenParams) WithContext(ctx context.Context) *CreateFileTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create file token params
func (o *CreateFileTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create file token params
func (o *CreateFileTokenParams) WithHTTPClient(client *http.Client) *CreateFileTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create file token params
func (o *CreateFileTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the create file token params
func (o *CreateFileTokenParams) WithResource(resource *models.FileToken) *CreateFileTokenParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the create file token params
func (o *CreateFileTokenParams) SetResource(resource *models.FileToken) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateFileTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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