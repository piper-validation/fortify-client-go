// Code generated by go-swagger; DO NOT EDIT.

package bug_filing_requirements_of_project_version_controller

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

// NewLoginBugFilingRequirementsOfProjectVersionParams creates a new LoginBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized.
func NewLoginBugFilingRequirementsOfProjectVersionParams() *LoginBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &LoginBugFilingRequirementsOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLoginBugFilingRequirementsOfProjectVersionParamsWithTimeout creates a new LoginBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLoginBugFilingRequirementsOfProjectVersionParamsWithTimeout(timeout time.Duration) *LoginBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &LoginBugFilingRequirementsOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewLoginBugFilingRequirementsOfProjectVersionParamsWithContext creates a new LoginBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewLoginBugFilingRequirementsOfProjectVersionParamsWithContext(ctx context.Context) *LoginBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &LoginBugFilingRequirementsOfProjectVersionParams{

		Context: ctx,
	}
}

// NewLoginBugFilingRequirementsOfProjectVersionParamsWithHTTPClient creates a new LoginBugFilingRequirementsOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLoginBugFilingRequirementsOfProjectVersionParamsWithHTTPClient(client *http.Client) *LoginBugFilingRequirementsOfProjectVersionParams {
	var ()
	return &LoginBugFilingRequirementsOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*LoginBugFilingRequirementsOfProjectVersionParams contains all the parameters to send to the API endpoint
for the login bug filing requirements of project version operation typically these are written to a http.Request
*/
type LoginBugFilingRequirementsOfProjectVersionParams struct {

	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Resource
	  resource

	*/
	Resource *models.BugFilingRequirementsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) WithTimeout(timeout time.Duration) *LoginBugFilingRequirementsOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) WithContext(ctx context.Context) *LoginBugFilingRequirementsOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) WithHTTPClient(client *http.Client) *LoginBugFilingRequirementsOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithParentID adds the parentID to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) WithParentID(parentID int64) *LoginBugFilingRequirementsOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithResource adds the resource to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) WithResource(resource *models.BugFilingRequirementsRequest) *LoginBugFilingRequirementsOfProjectVersionParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the login bug filing requirements of project version params
func (o *LoginBugFilingRequirementsOfProjectVersionParams) SetResource(resource *models.BugFilingRequirementsRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *LoginBugFilingRequirementsOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
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
