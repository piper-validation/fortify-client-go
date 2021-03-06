// Code generated by go-swagger; DO NOT EDIT.

package issue_controller

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

// NewListIssueParams creates a new ListIssueParams object
// with the default values initialized.
func NewListIssueParams() *ListIssueParams {
	var (
		fulltextsearchDefault = string("true")
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListIssueParams{
		Fulltextsearch: fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListIssueParamsWithTimeout creates a new ListIssueParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIssueParamsWithTimeout(timeout time.Duration) *ListIssueParams {
	var (
		fulltextsearchDefault = string("true")
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListIssueParams{
		Fulltextsearch: fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,

		timeout: timeout,
	}
}

// NewListIssueParamsWithContext creates a new ListIssueParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIssueParamsWithContext(ctx context.Context) *ListIssueParams {
	var (
		fulltextsearchDefault = string("true")
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListIssueParams{
		Fulltextsearch: fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,

		Context: ctx,
	}
}

// NewListIssueParamsWithHTTPClient creates a new ListIssueParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIssueParamsWithHTTPClient(client *http.Client) *ListIssueParams {
	var (
		fulltextsearchDefault = string("true")
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListIssueParams{
		Fulltextsearch: fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,
		HTTPClient:     client,
	}
}

/*ListIssueParams contains all the parameters to send to the API endpoint
for the list issue operation typically these are written to a http.Request
*/
type ListIssueParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*Fulltextsearch
	  Only 'true' is supported

	*/
	Fulltextsearch string
	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
	/*Q
	  A full text search query

	*/
	Q string
	/*Start
	  A start offset in object listing

	*/
	Start *int32

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list issue params
func (o *ListIssueParams) WithTimeout(timeout time.Duration) *ListIssueParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list issue params
func (o *ListIssueParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list issue params
func (o *ListIssueParams) WithContext(ctx context.Context) *ListIssueParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list issue params
func (o *ListIssueParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list issue params
func (o *ListIssueParams) WithHTTPClient(client *http.Client) *ListIssueParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list issue params
func (o *ListIssueParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list issue params
func (o *ListIssueParams) WithFields(fields *string) *ListIssueParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list issue params
func (o *ListIssueParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFulltextsearch adds the fulltextsearch to the list issue params
func (o *ListIssueParams) WithFulltextsearch(fulltextsearch string) *ListIssueParams {
	o.SetFulltextsearch(fulltextsearch)
	return o
}

// SetFulltextsearch adds the fulltextsearch to the list issue params
func (o *ListIssueParams) SetFulltextsearch(fulltextsearch string) {
	o.Fulltextsearch = fulltextsearch
}

// WithLimit adds the limit to the list issue params
func (o *ListIssueParams) WithLimit(limit *int32) *ListIssueParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list issue params
func (o *ListIssueParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithQ adds the q to the list issue params
func (o *ListIssueParams) WithQ(q string) *ListIssueParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list issue params
func (o *ListIssueParams) SetQ(q string) {
	o.Q = q
}

// WithStart adds the start to the list issue params
func (o *ListIssueParams) WithStart(start *int32) *ListIssueParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list issue params
func (o *ListIssueParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListIssueParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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

	// query param fulltextsearch
	qrFulltextsearch := o.Fulltextsearch
	qFulltextsearch := qrFulltextsearch
	if qFulltextsearch != "" {
		if err := r.SetQueryParam("fulltextsearch", qFulltextsearch); err != nil {
			return err
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

	// query param q
	qrQ := o.Q
	qQ := qrQ
	if qQ != "" {
		if err := r.SetQueryParam("q", qQ); err != nil {
			return err
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
