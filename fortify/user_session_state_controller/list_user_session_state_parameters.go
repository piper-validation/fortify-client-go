// Code generated by go-swagger; DO NOT EDIT.

package user_session_state_controller

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

// NewListUserSessionStateParams creates a new ListUserSessionStateParams object
// with the default values initialized.
func NewListUserSessionStateParams() *ListUserSessionStateParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListUserSessionStateParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListUserSessionStateParamsWithTimeout creates a new ListUserSessionStateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListUserSessionStateParamsWithTimeout(timeout time.Duration) *ListUserSessionStateParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListUserSessionStateParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: timeout,
	}
}

// NewListUserSessionStateParamsWithContext creates a new ListUserSessionStateParams object
// with the default values initialized, and the ability to set a context for a request
func NewListUserSessionStateParamsWithContext(ctx context.Context) *ListUserSessionStateParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListUserSessionStateParams{
		Limit: &limitDefault,
		Start: &startDefault,

		Context: ctx,
	}
}

// NewListUserSessionStateParamsWithHTTPClient creates a new ListUserSessionStateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListUserSessionStateParamsWithHTTPClient(client *http.Client) *ListUserSessionStateParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListUserSessionStateParams{
		Limit:      &limitDefault,
		Start:      &startDefault,
		HTTPClient: client,
	}
}

/*ListUserSessionStateParams contains all the parameters to send to the API endpoint
for the list user session state operation typically these are written to a http.Request
*/
type ListUserSessionStateParams struct {

	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
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

// WithTimeout adds the timeout to the list user session state params
func (o *ListUserSessionStateParams) WithTimeout(timeout time.Duration) *ListUserSessionStateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list user session state params
func (o *ListUserSessionStateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list user session state params
func (o *ListUserSessionStateParams) WithContext(ctx context.Context) *ListUserSessionStateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list user session state params
func (o *ListUserSessionStateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list user session state params
func (o *ListUserSessionStateParams) WithHTTPClient(client *http.Client) *ListUserSessionStateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list user session state params
func (o *ListUserSessionStateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithLimit adds the limit to the list user session state params
func (o *ListUserSessionStateParams) WithLimit(limit *int32) *ListUserSessionStateParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list user session state params
func (o *ListUserSessionStateParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithQ adds the q to the list user session state params
func (o *ListUserSessionStateParams) WithQ(q *string) *ListUserSessionStateParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list user session state params
func (o *ListUserSessionStateParams) SetQ(q *string) {
	o.Q = q
}

// WithStart adds the start to the list user session state params
func (o *ListUserSessionStateParams) WithStart(start *int32) *ListUserSessionStateParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list user session state params
func (o *ListUserSessionStateParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListUserSessionStateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
