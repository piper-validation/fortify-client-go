// Code generated by go-swagger; DO NOT EDIT.

package variable_controller

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

// NewListVariableParams creates a new ListVariableParams object
// with the default values initialized.
func NewListVariableParams() *ListVariableParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListVariableParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListVariableParamsWithTimeout creates a new ListVariableParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListVariableParamsWithTimeout(timeout time.Duration) *ListVariableParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListVariableParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: timeout,
	}
}

// NewListVariableParamsWithContext creates a new ListVariableParams object
// with the default values initialized, and the ability to set a context for a request
func NewListVariableParamsWithContext(ctx context.Context) *ListVariableParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListVariableParams{
		Limit: &limitDefault,
		Start: &startDefault,

		Context: ctx,
	}
}

// NewListVariableParamsWithHTTPClient creates a new ListVariableParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListVariableParamsWithHTTPClient(client *http.Client) *ListVariableParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListVariableParams{
		Limit:      &limitDefault,
		Start:      &startDefault,
		HTTPClient: client,
	}
}

/*ListVariableParams contains all the parameters to send to the API endpoint
for the list variable operation typically these are written to a http.Request
*/
type ListVariableParams struct {

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

// WithTimeout adds the timeout to the list variable params
func (o *ListVariableParams) WithTimeout(timeout time.Duration) *ListVariableParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list variable params
func (o *ListVariableParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list variable params
func (o *ListVariableParams) WithContext(ctx context.Context) *ListVariableParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list variable params
func (o *ListVariableParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list variable params
func (o *ListVariableParams) WithHTTPClient(client *http.Client) *ListVariableParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list variable params
func (o *ListVariableParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithFields adds the fields to the list variable params
func (o *ListVariableParams) WithFields(fields *string) *ListVariableParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list variable params
func (o *ListVariableParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithLimit adds the limit to the list variable params
func (o *ListVariableParams) WithLimit(limit *int32) *ListVariableParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list variable params
func (o *ListVariableParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOrderby adds the orderby to the list variable params
func (o *ListVariableParams) WithOrderby(orderby *string) *ListVariableParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the list variable params
func (o *ListVariableParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithQ adds the q to the list variable params
func (o *ListVariableParams) WithQ(q *string) *ListVariableParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list variable params
func (o *ListVariableParams) SetQ(q *string) {
	o.Q = q
}

// WithStart adds the start to the list variable params
func (o *ListVariableParams) WithStart(start *int32) *ListVariableParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list variable params
func (o *ListVariableParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListVariableParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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