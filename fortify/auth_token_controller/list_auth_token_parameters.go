// Code generated by go-swagger; DO NOT EDIT.

package auth_token_controller

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

// NewListAuthTokenParams creates a new ListAuthTokenParams object
// with the default values initialized.
func NewListAuthTokenParams() *ListAuthTokenParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListAuthTokenParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListAuthTokenParamsWithTimeout creates a new ListAuthTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAuthTokenParamsWithTimeout(timeout time.Duration) *ListAuthTokenParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListAuthTokenParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: timeout,
	}
}

// NewListAuthTokenParamsWithContext creates a new ListAuthTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAuthTokenParamsWithContext(ctx context.Context) *ListAuthTokenParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListAuthTokenParams{
		Limit: &limitDefault,
		Start: &startDefault,

		Context: ctx,
	}
}

// NewListAuthTokenParamsWithHTTPClient creates a new ListAuthTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAuthTokenParamsWithHTTPClient(client *http.Client) *ListAuthTokenParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListAuthTokenParams{
		Limit:      &limitDefault,
		Start:      &startDefault,
		HTTPClient: client,
	}
}

/*ListAuthTokenParams contains all the parameters to send to the API endpoint
for the list auth token operation typically these are written to a http.Request
*/
type ListAuthTokenParams struct {

	/*Fields
	  Output fields

	*/
	Fields *string
	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
	/*Orderby
	  Fields to order by

	*/
	Orderby *string
	/*Q
	  A search query

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

// WithTimeout adds the timeout to the list auth token params
func (o *ListAuthTokenParams) WithTimeout(timeout time.Duration) *ListAuthTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list auth token params
func (o *ListAuthTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list auth token params
func (o *ListAuthTokenParams) WithContext(ctx context.Context) *ListAuthTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list auth token params
func (o *ListAuthTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list auth token params
func (o *ListAuthTokenParams) WithHTTPClient(client *http.Client) *ListAuthTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list auth token params
func (o *ListAuthTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list auth token params
func (o *ListAuthTokenParams) WithFields(fields *string) *ListAuthTokenParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list auth token params
func (o *ListAuthTokenParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the list auth token params
func (o *ListAuthTokenParams) WithLimit(limit *int32) *ListAuthTokenParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list auth token params
func (o *ListAuthTokenParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOrderby adds the orderby to the list auth token params
func (o *ListAuthTokenParams) WithOrderby(orderby *string) *ListAuthTokenParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the list auth token params
func (o *ListAuthTokenParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithQ adds the q to the list auth token params
func (o *ListAuthTokenParams) WithQ(q *string) *ListAuthTokenParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list auth token params
func (o *ListAuthTokenParams) SetQ(q *string) {
	o.Q = q
}

// WithStart adds the start to the list auth token params
func (o *ListAuthTokenParams) WithStart(start *int32) *ListAuthTokenParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list auth token params
func (o *ListAuthTokenParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListAuthTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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