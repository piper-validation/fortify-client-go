// Code generated by go-swagger; DO NOT EDIT.

package project_version_of_project_controller

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

// NewListProjectVersionOfProjectParams creates a new ListProjectVersionOfProjectParams object
// with the default values initialized.
func NewListProjectVersionOfProjectParams() *ListProjectVersionOfProjectParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionOfProjectParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		Start:            &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListProjectVersionOfProjectParamsWithTimeout creates a new ListProjectVersionOfProjectParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListProjectVersionOfProjectParamsWithTimeout(timeout time.Duration) *ListProjectVersionOfProjectParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionOfProjectParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		Start:            &startDefault,

		timeout: timeout,
	}
}

// NewListProjectVersionOfProjectParamsWithContext creates a new ListProjectVersionOfProjectParams object
// with the default values initialized, and the ability to set a context for a request
func NewListProjectVersionOfProjectParamsWithContext(ctx context.Context) *ListProjectVersionOfProjectParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionOfProjectParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		Start:            &startDefault,

		Context: ctx,
	}
}

// NewListProjectVersionOfProjectParamsWithHTTPClient creates a new ListProjectVersionOfProjectParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListProjectVersionOfProjectParamsWithHTTPClient(client *http.Client) *ListProjectVersionOfProjectParams {
	var (
		fulltextsearchDefault   = bool(false)
		includeInactiveDefault  = bool(false)
		limitDefault            = int32(200)
		myAssignedIssuesDefault = bool(false)
		startDefault            = int32(0)
	)
	return &ListProjectVersionOfProjectParams{
		Fulltextsearch:   &fulltextsearchDefault,
		IncludeInactive:  &includeInactiveDefault,
		Limit:            &limitDefault,
		MyAssignedIssues: &myAssignedIssuesDefault,
		Start:            &startDefault,
		HTTPClient:       client,
	}
}

/*ListProjectVersionOfProjectParams contains all the parameters to send to the API endpoint
for the list project version of project operation typically these are written to a http.Request
*/
type ListProjectVersionOfProjectParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*Fulltextsearch
	  If 'true', interpret 'q' parameter as full text search query, defaults to 'false'

	*/
	Fulltextsearch *bool
	/*IncludeInactive
	  includeInactive

	*/
	IncludeInactive *bool
	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
	/*MyAssignedIssues
	  myAssignedIssues

	*/
	MyAssignedIssues *bool
	/*Orderby
	  Fields to order by

	*/
	Orderby *string
	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Q
	  A search-spec of full text search query (see fulltextsearch parameter)

	*/
	Q *string
	/*Start
	  A start offset in object listing

	*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithTimeout(timeout time.Duration) *ListProjectVersionOfProjectParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithContext(ctx context.Context) *ListProjectVersionOfProjectParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithHTTPClient(client *http.Client) *ListProjectVersionOfProjectParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithFields(fields *string) *ListProjectVersionOfProjectParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFulltextsearch adds the fulltextsearch to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithFulltextsearch(fulltextsearch *bool) *ListProjectVersionOfProjectParams {
	o.SetFulltextsearch(fulltextsearch)
	return o
}

// SetFulltextsearch adds the fulltextsearch to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetFulltextsearch(fulltextsearch *bool) {
	o.Fulltextsearch = fulltextsearch
}

// WithIncludeInactive adds the includeInactive to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithIncludeInactive(includeInactive *bool) *ListProjectVersionOfProjectParams {
	o.SetIncludeInactive(includeInactive)
	return o
}

// SetIncludeInactive adds the includeInactive to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetIncludeInactive(includeInactive *bool) {
	o.IncludeInactive = includeInactive
}

// WithLimit adds the limit to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithLimit(limit *int32) *ListProjectVersionOfProjectParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithMyAssignedIssues adds the myAssignedIssues to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithMyAssignedIssues(myAssignedIssues *bool) *ListProjectVersionOfProjectParams {
	o.SetMyAssignedIssues(myAssignedIssues)
	return o
}

// SetMyAssignedIssues adds the myAssignedIssues to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetMyAssignedIssues(myAssignedIssues *bool) {
	o.MyAssignedIssues = myAssignedIssues
}

// WithOrderby adds the orderby to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithOrderby(orderby *string) *ListProjectVersionOfProjectParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithParentID adds the parentID to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithParentID(parentID int64) *ListProjectVersionOfProjectParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithQ adds the q to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithQ(q *string) *ListProjectVersionOfProjectParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetQ(q *string) {
	o.Q = q
}

// WithStart adds the start to the list project version of project params
func (o *ListProjectVersionOfProjectParams) WithStart(start *int32) *ListProjectVersionOfProjectParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list project version of project params
func (o *ListProjectVersionOfProjectParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListProjectVersionOfProjectParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	if o.Fulltextsearch != nil {

		// query param fulltextsearch
		var qrFulltextsearch bool
		if o.Fulltextsearch != nil {
			qrFulltextsearch = *o.Fulltextsearch
		}
		qFulltextsearch := swag.FormatBool(qrFulltextsearch)
		if qFulltextsearch != "" {
			if err := r.SetQueryParam("fulltextsearch", qFulltextsearch); err != nil {
				return err
			}
		}

	}

	if o.IncludeInactive != nil {

		// query param includeInactive
		var qrIncludeInactive bool
		if o.IncludeInactive != nil {
			qrIncludeInactive = *o.IncludeInactive
		}
		qIncludeInactive := swag.FormatBool(qrIncludeInactive)
		if qIncludeInactive != "" {
			if err := r.SetQueryParam("includeInactive", qIncludeInactive); err != nil {
				return err
			}
		}

	}

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

	if o.MyAssignedIssues != nil {

		// query param myAssignedIssues
		var qrMyAssignedIssues bool
		if o.MyAssignedIssues != nil {
			qrMyAssignedIssues = *o.MyAssignedIssues
		}
		qMyAssignedIssues := swag.FormatBool(qrMyAssignedIssues)
		if qMyAssignedIssues != "" {
			if err := r.SetQueryParam("myAssignedIssues", qMyAssignedIssues); err != nil {
				return err
			}
		}

	}

	if o.Orderby != nil {

		// query param orderby
		var qrOrderby string
		if o.Orderby != nil {
			qrOrderby = *o.Orderby
		}
		qOrderby := qrOrderby
		if qOrderby != "" {
			if err := r.SetQueryParam("orderby", qOrderby); err != nil {
				return err
			}
		}

	}

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
		return err
	}

	if o.Q != nil {

		// query param q
		var qrQ string
		if o.Q != nil {
			qrQ = *o.Q
		}
		qQ := qrQ
		if qQ != "" {
			if err := r.SetQueryParam("q", qQ); err != nil {
				return err
			}
		}

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