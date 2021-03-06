// Code generated by go-swagger; DO NOT EDIT.

package configuration_controller

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

// NewRefreshAuditAssistantPoliciesConfigurationParams creates a new RefreshAuditAssistantPoliciesConfigurationParams object
// with the default values initialized.
func NewRefreshAuditAssistantPoliciesConfigurationParams() *RefreshAuditAssistantPoliciesConfigurationParams {
	var ()
	return &RefreshAuditAssistantPoliciesConfigurationParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewRefreshAuditAssistantPoliciesConfigurationParamsWithTimeout creates a new RefreshAuditAssistantPoliciesConfigurationParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewRefreshAuditAssistantPoliciesConfigurationParamsWithTimeout(timeout time.Duration) *RefreshAuditAssistantPoliciesConfigurationParams {
	var ()
	return &RefreshAuditAssistantPoliciesConfigurationParams{

		timeout: timeout,
	}
}

// NewRefreshAuditAssistantPoliciesConfigurationParamsWithContext creates a new RefreshAuditAssistantPoliciesConfigurationParams object
// with the default values initialized, and the ability to set a context for a request
func NewRefreshAuditAssistantPoliciesConfigurationParamsWithContext(ctx context.Context) *RefreshAuditAssistantPoliciesConfigurationParams {
	var ()
	return &RefreshAuditAssistantPoliciesConfigurationParams{

		Context: ctx,
	}
}

// NewRefreshAuditAssistantPoliciesConfigurationParamsWithHTTPClient creates a new RefreshAuditAssistantPoliciesConfigurationParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewRefreshAuditAssistantPoliciesConfigurationParamsWithHTTPClient(client *http.Client) *RefreshAuditAssistantPoliciesConfigurationParams {
	var ()
	return &RefreshAuditAssistantPoliciesConfigurationParams{
		HTTPClient: client,
	}
}

/*RefreshAuditAssistantPoliciesConfigurationParams contains all the parameters to send to the API endpoint
for the refresh audit assistant policies configuration operation typically these are written to a http.Request
*/
type RefreshAuditAssistantPoliciesConfigurationParams struct {

	/*Resource
	  resource

	*/
	Resource *models.RefreshAuditAssistantPoliciesRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) WithTimeout(timeout time.Duration) *RefreshAuditAssistantPoliciesConfigurationParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) WithContext(ctx context.Context) *RefreshAuditAssistantPoliciesConfigurationParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) WithHTTPClient(client *http.Client) *RefreshAuditAssistantPoliciesConfigurationParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) WithResource(resource *models.RefreshAuditAssistantPoliciesRequest) *RefreshAuditAssistantPoliciesConfigurationParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the refresh audit assistant policies configuration params
func (o *RefreshAuditAssistantPoliciesConfigurationParams) SetResource(resource *models.RefreshAuditAssistantPoliciesRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *RefreshAuditAssistantPoliciesConfigurationParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
