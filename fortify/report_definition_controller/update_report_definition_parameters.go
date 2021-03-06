// Code generated by go-swagger; DO NOT EDIT.

package report_definition_controller

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

	"github.com/piper-validation/fortify-client-go/models"
)

// NewUpdateReportDefinitionParams creates a new UpdateReportDefinitionParams object
// with the default values initialized.
func NewUpdateReportDefinitionParams() *UpdateReportDefinitionParams {
	var ()
	return &UpdateReportDefinitionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateReportDefinitionParamsWithTimeout creates a new UpdateReportDefinitionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateReportDefinitionParamsWithTimeout(timeout time.Duration) *UpdateReportDefinitionParams {
	var ()
	return &UpdateReportDefinitionParams{

		timeout: timeout,
	}
}

// NewUpdateReportDefinitionParamsWithContext creates a new UpdateReportDefinitionParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateReportDefinitionParamsWithContext(ctx context.Context) *UpdateReportDefinitionParams {
	var ()
	return &UpdateReportDefinitionParams{

		Context: ctx,
	}
}

// NewUpdateReportDefinitionParamsWithHTTPClient creates a new UpdateReportDefinitionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateReportDefinitionParamsWithHTTPClient(client *http.Client) *UpdateReportDefinitionParams {
	var ()
	return &UpdateReportDefinitionParams{
		HTTPClient: client,
	}
}

/*UpdateReportDefinitionParams contains all the parameters to send to the API endpoint
for the update report definition operation typically these are written to a http.Request
*/
type UpdateReportDefinitionParams struct {

	/*ID
	  id

	*/
	ID int64
	/*Resource
	  resource

	*/
	Resource *models.ReportDefinition

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update report definition params
func (o *UpdateReportDefinitionParams) WithTimeout(timeout time.Duration) *UpdateReportDefinitionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update report definition params
func (o *UpdateReportDefinitionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update report definition params
func (o *UpdateReportDefinitionParams) WithContext(ctx context.Context) *UpdateReportDefinitionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update report definition params
func (o *UpdateReportDefinitionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update report definition params
func (o *UpdateReportDefinitionParams) WithHTTPClient(client *http.Client) *UpdateReportDefinitionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update report definition params
func (o *UpdateReportDefinitionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithID adds the id to the update report definition params
func (o *UpdateReportDefinitionParams) WithID(id int64) *UpdateReportDefinitionParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the update report definition params
func (o *UpdateReportDefinitionParams) SetID(id int64) {
	o.ID = id
}

// WithResource adds the resource to the update report definition params
func (o *UpdateReportDefinitionParams) WithResource(resource *models.ReportDefinition) *UpdateReportDefinitionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update report definition params
func (o *UpdateReportDefinitionParams) SetResource(resource *models.ReportDefinition) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateReportDefinitionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

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
