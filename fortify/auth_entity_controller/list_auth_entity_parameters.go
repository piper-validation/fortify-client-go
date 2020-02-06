// Code generated by go-swagger; DO NOT EDIT.

package auth_entity_controller

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

// NewListAuthEntityParams creates a new ListAuthEntityParams object
// with the default values initialized.
func NewListAuthEntityParams() *ListAuthEntityParams {
	var (
		fulltextsearchDefault = bool(false)
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListAuthEntityParams{
		Fulltextsearch: &fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListAuthEntityParamsWithTimeout creates a new ListAuthEntityParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListAuthEntityParamsWithTimeout(timeout time.Duration) *ListAuthEntityParams {
	var (
		fulltextsearchDefault = bool(false)
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListAuthEntityParams{
		Fulltextsearch: &fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,

		timeout: timeout,
	}
}

// NewListAuthEntityParamsWithContext creates a new ListAuthEntityParams object
// with the default values initialized, and the ability to set a context for a request
func NewListAuthEntityParamsWithContext(ctx context.Context) *ListAuthEntityParams {
	var (
		fulltextsearchDefault = bool(false)
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListAuthEntityParams{
		Fulltextsearch: &fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,

		Context: ctx,
	}
}

// NewListAuthEntityParamsWithHTTPClient creates a new ListAuthEntityParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListAuthEntityParamsWithHTTPClient(client *http.Client) *ListAuthEntityParams {
	var (
		fulltextsearchDefault = bool(false)
		limitDefault          = int32(200)
		startDefault          = int32(0)
	)
	return &ListAuthEntityParams{
		Fulltextsearch: &fulltextsearchDefault,
		Limit:          &limitDefault,
		Start:          &startDefault,
		HTTPClient:     client,
	}
}

/*ListAuthEntityParams contains all the parameters to send to the API endpoint
for the list auth entity operation typically these are written to a http.Request
*/
type ListAuthEntityParams struct {

	/*Embed
	  Fields to embed

	*/
	Embed *string
	/*Entityname
	  entityname

	*/
	Entityname *string
	/*Fields
	  Output fields

	*/
	Fields *string
	/*Fulltextsearch
	  If 'true', interpret 'q' parameter as full text search query, defaults to 'false'

	*/
	Fulltextsearch *bool
	/*Ldaptype
	  ldaptype

	*/
	Ldaptype *string
	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
	/*Orderby
	  Fields to order by

	*/
	Orderby *string
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

// WithTimeout adds the timeout to the list auth entity params
func (o *ListAuthEntityParams) WithTimeout(timeout time.Duration) *ListAuthEntityParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list auth entity params
func (o *ListAuthEntityParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list auth entity params
func (o *ListAuthEntityParams) WithContext(ctx context.Context) *ListAuthEntityParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list auth entity params
func (o *ListAuthEntityParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list auth entity params
func (o *ListAuthEntityParams) WithHTTPClient(client *http.Client) *ListAuthEntityParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list auth entity params
func (o *ListAuthEntityParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmbed adds the embed to the list auth entity params
func (o *ListAuthEntityParams) WithEmbed(embed *string) *ListAuthEntityParams {
	o.SetEmbed(embed)
	return o
}

// SetEmbed adds the embed to the list auth entity params
func (o *ListAuthEntityParams) SetEmbed(embed *string) {
	o.Embed = embed
}

// WithEntityname adds the entityname to the list auth entity params
func (o *ListAuthEntityParams) WithEntityname(entityname *string) *ListAuthEntityParams {
	o.SetEntityname(entityname)
	return o
}

// SetEntityname adds the entityname to the list auth entity params
func (o *ListAuthEntityParams) SetEntityname(entityname *string) {
	o.Entityname = entityname
}

// WithFields adds the fields to the list auth entity params
func (o *ListAuthEntityParams) WithFields(fields *string) *ListAuthEntityParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list auth entity params
func (o *ListAuthEntityParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithFulltextsearch adds the fulltextsearch to the list auth entity params
func (o *ListAuthEntityParams) WithFulltextsearch(fulltextsearch *bool) *ListAuthEntityParams {
	o.SetFulltextsearch(fulltextsearch)
	return o
}

// SetFulltextsearch adds the fulltextsearch to the list auth entity params
func (o *ListAuthEntityParams) SetFulltextsearch(fulltextsearch *bool) {
	o.Fulltextsearch = fulltextsearch
}

// WithLdaptype adds the ldaptype to the list auth entity params
func (o *ListAuthEntityParams) WithLdaptype(ldaptype *string) *ListAuthEntityParams {
	o.SetLdaptype(ldaptype)
	return o
}

// SetLdaptype adds the ldaptype to the list auth entity params
func (o *ListAuthEntityParams) SetLdaptype(ldaptype *string) {
	o.Ldaptype = ldaptype
}

// WithLimit adds the limit to the list auth entity params
func (o *ListAuthEntityParams) WithLimit(limit *int32) *ListAuthEntityParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list auth entity params
func (o *ListAuthEntityParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOrderby adds the orderby to the list auth entity params
func (o *ListAuthEntityParams) WithOrderby(orderby *string) *ListAuthEntityParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the list auth entity params
func (o *ListAuthEntityParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithQ adds the q to the list auth entity params
func (o *ListAuthEntityParams) WithQ(q *string) *ListAuthEntityParams {
	o.SetQ(q)
	return o
}

// SetQ adds the q to the list auth entity params
func (o *ListAuthEntityParams) SetQ(q *string) {
	o.Q = q
}

// WithStart adds the start to the list auth entity params
func (o *ListAuthEntityParams) WithStart(start *int32) *ListAuthEntityParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list auth entity params
func (o *ListAuthEntityParams) SetStart(start *int32) {
	o.Start = start
}

// WriteToRequest writes these params to a swagger request
func (o *ListAuthEntityParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Embed != nil {

		// query param embed
		var qrEmbed string
		if o.Embed != nil {
			qrEmbed = *o.Embed
		}
		qEmbed := qrEmbed
		if qEmbed != "" {
			if err := r.SetQueryParam("embed", qEmbed); err != nil {
				return err
			}
		}

	}

	if o.Entityname != nil {

		// query param entityname
		var qrEntityname string
		if o.Entityname != nil {
			qrEntityname = *o.Entityname
		}
		qEntityname := qrEntityname
		if qEntityname != "" {
			if err := r.SetQueryParam("entityname", qEntityname); err != nil {
				return err
			}
		}

	}

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

	if o.Ldaptype != nil {

		// query param ldaptype
		var qrLdaptype string
		if o.Ldaptype != nil {
			qrLdaptype = *o.Ldaptype
		}
		qLdaptype := qrLdaptype
		if qLdaptype != "" {
			if err := r.SetQueryParam("ldaptype", qLdaptype); err != nil {
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