// Code generated by go-swagger; DO NOT EDIT.

package cloud_system_poll_status_controller

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

// NewGetCloudSystemPollStatusParams creates a new GetCloudSystemPollStatusParams object
// with the default values initialized.
func NewGetCloudSystemPollStatusParams() *GetCloudSystemPollStatusParams {

	return &GetCloudSystemPollStatusParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetCloudSystemPollStatusParamsWithTimeout creates a new GetCloudSystemPollStatusParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetCloudSystemPollStatusParamsWithTimeout(timeout time.Duration) *GetCloudSystemPollStatusParams {

	return &GetCloudSystemPollStatusParams{

		timeout: timeout,
	}
}

// NewGetCloudSystemPollStatusParamsWithContext creates a new GetCloudSystemPollStatusParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetCloudSystemPollStatusParamsWithContext(ctx context.Context) *GetCloudSystemPollStatusParams {

	return &GetCloudSystemPollStatusParams{

		Context: ctx,
	}
}

// NewGetCloudSystemPollStatusParamsWithHTTPClient creates a new GetCloudSystemPollStatusParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetCloudSystemPollStatusParamsWithHTTPClient(client *http.Client) *GetCloudSystemPollStatusParams {

	return &GetCloudSystemPollStatusParams{
		HTTPClient: client,
	}
}

/*GetCloudSystemPollStatusParams contains all the parameters to send to the API endpoint
for the get cloud system poll status operation typically these are written to a http.Request
*/
type GetCloudSystemPollStatusParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get cloud system poll status params
func (o *GetCloudSystemPollStatusParams) WithTimeout(timeout time.Duration) *GetCloudSystemPollStatusParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get cloud system poll status params
func (o *GetCloudSystemPollStatusParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get cloud system poll status params
func (o *GetCloudSystemPollStatusParams) WithContext(ctx context.Context) *GetCloudSystemPollStatusParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get cloud system poll status params
func (o *GetCloudSystemPollStatusParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get cloud system poll status params
func (o *GetCloudSystemPollStatusParams) WithHTTPClient(client *http.Client) *GetCloudSystemPollStatusParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get cloud system poll status params
func (o *GetCloudSystemPollStatusParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetCloudSystemPollStatusParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}