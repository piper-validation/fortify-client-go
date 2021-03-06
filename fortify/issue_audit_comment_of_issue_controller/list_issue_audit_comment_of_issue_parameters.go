// Code generated by go-swagger; DO NOT EDIT.

package issue_audit_comment_of_issue_controller

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

// NewListIssueAuditCommentOfIssueParams creates a new ListIssueAuditCommentOfIssueParams object
// with the default values initialized.
func NewListIssueAuditCommentOfIssueParams() *ListIssueAuditCommentOfIssueParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListIssueAuditCommentOfIssueParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListIssueAuditCommentOfIssueParamsWithTimeout creates a new ListIssueAuditCommentOfIssueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIssueAuditCommentOfIssueParamsWithTimeout(timeout time.Duration) *ListIssueAuditCommentOfIssueParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListIssueAuditCommentOfIssueParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: timeout,
	}
}

// NewListIssueAuditCommentOfIssueParamsWithContext creates a new ListIssueAuditCommentOfIssueParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIssueAuditCommentOfIssueParamsWithContext(ctx context.Context) *ListIssueAuditCommentOfIssueParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListIssueAuditCommentOfIssueParams{
		Limit: &limitDefault,
		Start: &startDefault,

		Context: ctx,
	}
}

// NewListIssueAuditCommentOfIssueParamsWithHTTPClient creates a new ListIssueAuditCommentOfIssueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIssueAuditCommentOfIssueParamsWithHTTPClient(client *http.Client) *ListIssueAuditCommentOfIssueParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListIssueAuditCommentOfIssueParams{
		Limit:      &limitDefault,
		Start:      &startDefault,
		HTTPClient: client,
	}
}

/*ListIssueAuditCommentOfIssueParams contains all the parameters to send to the API endpoint
for the list issue audit comment of issue operation typically these are written to a http.Request
*/
type ListIssueAuditCommentOfIssueParams struct {

	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Start
	  A start offset in object listing

	*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) WithTimeout(timeout time.Duration) *ListIssueAuditCommentOfIssueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) WithContext(ctx context.Context) *ListIssueAuditCommentOfIssueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) WithHTTPClient(client *http.Client) *ListIssueAuditCommentOfIssueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) WithLimit(limit *int32) *ListIssueAuditCommentOfIssueParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithParentID adds the parentID to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) WithParentID(parentID int64) *ListIssueAuditCommentOfIssueParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithStart adds the start to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) WithStart(start *int32) *ListIssueAuditCommentOfIssueParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list issue audit comment of issue params
func (o *ListIssueAuditCommentOfIssueParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListIssueAuditCommentOfIssueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Limit != nil {

		// query param limit
		var qrLimit int32
		if o.Limit != nil {
			qrLimit = *o.Limit
		}
		qLimit := swag.FormatInt32(qrLimit)
		if qLimit != "" {
			if err := r.SetQueryParam("limit", qLimit); err != nil {
				return err
			}
		}

	}

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
		return err
	}

	if o.Start != nil {

		// query param start
		var qrStart int32
		if o.Start != nil {
			qrStart = *o.Start
		}
		qStart := swag.FormatInt32(qrStart)
		if qStart != "" {
			if err := r.SetQueryParam("start", qStart); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
