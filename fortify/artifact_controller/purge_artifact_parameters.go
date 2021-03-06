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

// NewPurgeArtifactParams creates a new PurgeArtifactParams object
// with the default values initialized.
func NewPurgeArtifactParams() *PurgeArtifactParams {
	var ()
	return &PurgeArtifactParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPurgeArtifactParamsWithTimeout creates a new PurgeArtifactParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPurgeArtifactParamsWithTimeout(timeout time.Duration) *PurgeArtifactParams {
	var ()
	return &PurgeArtifactParams{

		timeout: timeout,
	}
}

// NewPurgeArtifactParamsWithContext creates a new PurgeArtifactParams object
// with the default values initialized, and the ability to set a context for a request
func NewPurgeArtifactParamsWithContext(ctx context.Context) *PurgeArtifactParams {
	var ()
	return &PurgeArtifactParams{

		Context: ctx,
	}
}

// NewPurgeArtifactParamsWithHTTPClient creates a new PurgeArtifactParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPurgeArtifactParamsWithHTTPClient(client *http.Client) *PurgeArtifactParams {
	var ()
	return &PurgeArtifactParams{
		HTTPClient: client,
	}
}

/*PurgeArtifactParams contains all the parameters to send to the API endpoint
for the purge artifact operation typically these are written to a http.Request
*/
type PurgeArtifactParams struct {

	/*Resource
	  resource

	*/
	Resource *models.ArtifactPurgeRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the purge artifact params
func (o *PurgeArtifactParams) WithTimeout(timeout time.Duration) *PurgeArtifactParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the purge artifact params
func (o *PurgeArtifactParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the purge artifact params
func (o *PurgeArtifactParams) WithContext(ctx context.Context) *PurgeArtifactParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the purge artifact params
func (o *PurgeArtifactParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the purge artifact params
func (o *PurgeArtifactParams) WithHTTPClient(client *http.Client) *PurgeArtifactParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the purge artifact params
func (o *PurgeArtifactParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the purge artifact params
func (o *PurgeArtifactParams) WithResource(resource *models.ArtifactPurgeRequest) *PurgeArtifactParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the purge artifact params
func (o *PurgeArtifactParams) SetResource(resource *models.ArtifactPurgeRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *PurgeArtifactParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
