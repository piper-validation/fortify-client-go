// Code generated by go-swagger; DO NOT EDIT.

package issue_aging_portlet_controller

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

// NewGetIssueAgingPortletParams creates a new GetIssueAgingPortletParams object
// with the default values initialized.
func NewGetIssueAgingPortletParams() *GetIssueAgingPortletParams {

	return &GetIssueAgingPortletParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetIssueAgingPortletParamsWithTimeout creates a new GetIssueAgingPortletParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetIssueAgingPortletParamsWithTimeout(timeout time.Duration) *GetIssueAgingPortletParams {

	return &GetIssueAgingPortletParams{

		timeout: timeout,
	}
}

// NewGetIssueAgingPortletParamsWithContext creates a new GetIssueAgingPortletParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetIssueAgingPortletParamsWithContext(ctx context.Context) *GetIssueAgingPortletParams {

	return &GetIssueAgingPortletParams{

		Context: ctx,
	}
}

// NewGetIssueAgingPortletParamsWithHTTPClient creates a new GetIssueAgingPortletParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetIssueAgingPortletParamsWithHTTPClient(client *http.Client) *GetIssueAgingPortletParams {

	return &GetIssueAgingPortletParams{
		HTTPClient: client,
	}
}

/*GetIssueAgingPortletParams contains all the parameters to send to the API endpoint
for the get issue aging portlet operation typically these are written to a http.Request
*/
type GetIssueAgingPortletParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get issue aging portlet params
func (o *GetIssueAgingPortletParams) WithTimeout(timeout time.Duration) *GetIssueAgingPortletParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get issue aging portlet params
func (o *GetIssueAgingPortletParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get issue aging portlet params
func (o *GetIssueAgingPortletParams) WithContext(ctx context.Context) *GetIssueAgingPortletParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get issue aging portlet params
func (o *GetIssueAgingPortletParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get issue aging portlet params
func (o *GetIssueAgingPortletParams) WithHTTPClient(client *http.Client) *GetIssueAgingPortletParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get issue aging portlet params
func (o *GetIssueAgingPortletParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *GetIssueAgingPortletParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
