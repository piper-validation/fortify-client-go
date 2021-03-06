// Code generated by go-swagger; DO NOT EDIT.

package audit_history_of_issue_controller

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
)

// NewListAuditHistoryOfIssueParams creates a new ListAuditHistoryOfIssueParams object
// with the default values initialized.
func NewListAuditHistoryOfIssueParams() *ListAuditHistoryOfIssueParams {
	var ()
	return &ListAuditHistoryOfIssueParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListAuditHistoryOfIssueParamsWithTimeout creates a new ListAuditHistoryOfIssueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAuditHistoryOfIssueParamsWithTimeout(timeout time.Duration) *ListAuditHistoryOfIssueParams {
	var ()
	return &ListAuditHistoryOfIssueParams{

		timeout: timeout,
	}
}

// NewListAuditHistoryOfIssueParamsWithContext creates a new ListAuditHistoryOfIssueParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAuditHistoryOfIssueParamsWithContext(ctx context.Context) *ListAuditHistoryOfIssueParams {
	var ()
	return &ListAuditHistoryOfIssueParams{

		Context: ctx,
	}
}

// NewListAuditHistoryOfIssueParamsWithHTTPClient creates a new ListAuditHistoryOfIssueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAuditHistoryOfIssueParamsWithHTTPClient(client *http.Client) *ListAuditHistoryOfIssueParams {
	var ()
	return &ListAuditHistoryOfIssueParams{
		HTTPClient: client,
	}
}

/*ListAuditHistoryOfIssueParams contains all the parameters to send to the API endpoint
for the list audit history of issue operation typically these are written to a http.Request
*/
type ListAuditHistoryOfIssueParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*ParentID
	  parentId

	*/
	ParentID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) WithTimeout(timeout time.Duration) *ListAuditHistoryOfIssueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) WithContext(ctx context.Context) *ListAuditHistoryOfIssueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) WithHTTPClient(client *http.Client) *ListAuditHistoryOfIssueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) WithFields(fields *string) *ListAuditHistoryOfIssueParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithParentID adds the parentID to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) WithParentID(parentID int64) *ListAuditHistoryOfIssueParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list audit history of issue params
func (o *ListAuditHistoryOfIssueParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WriteToRequest writes these params to a swagger request
func (o *ListAuditHistoryOfIssueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Fields != nil {

		// query param fields
		var qrFields string
		if o.Fields != nil {
			qrFields = *o.Fields
		}
		qFields := qrFields
		if qFields != "" {
			if err := r.SetQueryParam("fields", qFields); err != nil {
				return err
			}
		}

	}

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
