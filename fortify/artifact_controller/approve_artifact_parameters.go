// Code generated by go-swagger; DO NOT EDIT.

package artifact_controller

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

// NewApproveArtifactParams creates a new ApproveArtifactParams object
// with the default values initialized.
func NewApproveArtifactParams() *ApproveArtifactParams {
	var ()
	return &ApproveArtifactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewApproveArtifactParamsWithTimeout creates a new ApproveArtifactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewApproveArtifactParamsWithTimeout(timeout time.Duration) *ApproveArtifactParams {
	var ()
	return &ApproveArtifactParams{

		timeout: timeout,
	}
}

// NewApproveArtifactParamsWithContext creates a new ApproveArtifactParams object
// with the default values initialized, and the ability to set a context for a request
func NewApproveArtifactParamsWithContext(ctx context.Context) *ApproveArtifactParams {
	var ()
	return &ApproveArtifactParams{

		Context: ctx,
	}
}

// NewApproveArtifactParamsWithHTTPClient creates a new ApproveArtifactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewApproveArtifactParamsWithHTTPClient(client *http.Client) *ApproveArtifactParams {
	var ()
	return &ApproveArtifactParams{
		HTTPClient: client,
	}
}

/*ApproveArtifactParams contains all the parameters to send to the API endpoint
for the approve artifact operation typically these are written to a http.Request
*/
type ApproveArtifactParams struct {

	/*Resource
	  resource

	*/
	Resource *models.ArtifactApproveRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the approve artifact params
func (o *ApproveArtifactParams) WithTimeout(timeout time.Duration) *ApproveArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the approve artifact params
func (o *ApproveArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the approve artifact params
func (o *ApproveArtifactParams) WithContext(ctx context.Context) *ApproveArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the approve artifact params
func (o *ApproveArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the approve artifact params
func (o *ApproveArtifactParams) WithHTTPClient(client *http.Client) *ApproveArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the approve artifact params
func (o *ApproveArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the approve artifact params
func (o *ApproveArtifactParams) WithResource(resource *models.ArtifactApproveRequest) *ApproveArtifactParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the approve artifact params
func (o *ApproveArtifactParams) SetResource(resource *models.ArtifactApproveRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *ApproveArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
