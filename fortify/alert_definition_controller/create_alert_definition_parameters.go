// Code generated by go-swagger; DO NOT EDIT.

package alert_definition_controller

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

// NewCreateAlertDefinitionParams creates a new CreateAlertDefinitionParams object
// with the default values initialized.
func NewCreateAlertDefinitionParams() *CreateAlertDefinitionParams {
	var ()
	return &CreateAlertDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAlertDefinitionParamsWithTimeout creates a new CreateAlertDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAlertDefinitionParamsWithTimeout(timeout time.Duration) *CreateAlertDefinitionParams {
	var ()
	return &CreateAlertDefinitionParams{

		timeout: timeout,
	}
}

// NewCreateAlertDefinitionParamsWithContext creates a new CreateAlertDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAlertDefinitionParamsWithContext(ctx context.Context) *CreateAlertDefinitionParams {
	var ()
	return &CreateAlertDefinitionParams{

		Context: ctx,
	}
}

// NewCreateAlertDefinitionParamsWithHTTPClient creates a new CreateAlertDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAlertDefinitionParamsWithHTTPClient(client *http.Client) *CreateAlertDefinitionParams {
	var ()
	return &CreateAlertDefinitionParams{
		HTTPClient: client,
	}
}

/*CreateAlertDefinitionParams contains all the parameters to send to the API endpoint
for the create alert definition operation typically these are written to a http.Request
*/
type CreateAlertDefinitionParams struct {

	/*Resource
	  resource

	*/
	Resource *models.AlertDefinitionDto

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create alert definition params
func (o *CreateAlertDefinitionParams) WithTimeout(timeout time.Duration) *CreateAlertDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create alert definition params
func (o *CreateAlertDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create alert definition params
func (o *CreateAlertDefinitionParams) WithContext(ctx context.Context) *CreateAlertDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create alert definition params
func (o *CreateAlertDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create alert definition params
func (o *CreateAlertDefinitionParams) WithHTTPClient(client *http.Client) *CreateAlertDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create alert definition params
func (o *CreateAlertDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the create alert definition params
func (o *CreateAlertDefinitionParams) WithResource(resource *models.AlertDefinitionDto) *CreateAlertDefinitionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the create alert definition params
func (o *CreateAlertDefinitionParams) SetResource(resource *models.AlertDefinitionDto) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAlertDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
