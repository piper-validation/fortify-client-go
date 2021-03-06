// Code generated by go-swagger; DO NOT EDIT.

package dashboard_version_controller

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

// NewListDashboardVersionParams creates a new ListDashboardVersionParams object
// with the default values initialized.
func NewListDashboardVersionParams() *ListDashboardVersionParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListDashboardVersionParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: cr.DefaultTimeout,
	}
}

// NewListDashboardVersionParamsWithTimeout creates a new ListDashboardVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListDashboardVersionParamsWithTimeout(timeout time.Duration) *ListDashboardVersionParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListDashboardVersionParams{
		Limit: &limitDefault,
		Start: &startDefault,

		timeout: timeout,
	}
}

// NewListDashboardVersionParamsWithContext creates a new ListDashboardVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListDashboardVersionParamsWithContext(ctx context.Context) *ListDashboardVersionParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListDashboardVersionParams{
		Limit: &limitDefault,
		Start: &startDefault,

		Context: ctx,
	}
}

// NewListDashboardVersionParamsWithHTTPClient creates a new ListDashboardVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListDashboardVersionParamsWithHTTPClient(client *http.Client) *ListDashboardVersionParams {
	var (
		limitDefault = int32(200)
		startDefault = int32(0)
	)
	return &ListDashboardVersionParams{
		Limit:      &limitDefault,
		Start:      &startDefault,
		HTTPClient: client,
	}
}

/*ListDashboardVersionParams contains all the parameters to send to the API endpoint
for the list dashboard version operation typically these are written to a http.Request
*/
type ListDashboardVersionParams struct {

	/*Aggregateby
	  aggregateby

	*/
	Aggregateby *string
	/*Attributefilter
	  attributefilter

	*/
	Attributefilter *string
	/*Attributes
	  attributes

	*/
	Attributes *string
	/*Enddate
	  enddate

	*/
	Enddate *string
	/*Fields
	  Output fields

	*/
	Fields *string
	/*Groupby
	  Fields to group by

	*/
	Groupby *string
	/*Limit
	  A maximum number of returned objects in listing, if '-1' or '0' no limit is applied

	*/
	Limit *int32
	/*Orderby
	  Fields to order by

	*/
	Orderby *string
	/*Performanceindicators
	  performanceindicators

	*/
	Performanceindicators *string
	/*Start
	  A start offset in object listing

	*/
	Start *int32
	/*Startdate
	  startdate

	*/
	Startdate *string
	/*Variables
	  variables

	*/
	Variables *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list dashboard version params
func (o *ListDashboardVersionParams) WithTimeout(timeout time.Duration) *ListDashboardVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list dashboard version params
func (o *ListDashboardVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list dashboard version params
func (o *ListDashboardVersionParams) WithContext(ctx context.Context) *ListDashboardVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list dashboard version params
func (o *ListDashboardVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list dashboard version params
func (o *ListDashboardVersionParams) WithHTTPClient(client *http.Client) *ListDashboardVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list dashboard version params
func (o *ListDashboardVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAggregateby adds the aggregateby to the list dashboard version params
func (o *ListDashboardVersionParams) WithAggregateby(aggregateby *string) *ListDashboardVersionParams {
	o.SetAggregateby(aggregateby)
	return o
}

// SetAggregateby adds the aggregateby to the list dashboard version params
func (o *ListDashboardVersionParams) SetAggregateby(aggregateby *string) {
	o.Aggregateby = aggregateby
}

// WithAttributefilter adds the attributefilter to the list dashboard version params
func (o *ListDashboardVersionParams) WithAttributefilter(attributefilter *string) *ListDashboardVersionParams {
	o.SetAttributefilter(attributefilter)
	return o
}

// SetAttributefilter adds the attributefilter to the list dashboard version params
func (o *ListDashboardVersionParams) SetAttributefilter(attributefilter *string) {
	o.Attributefilter = attributefilter
}

// WithAttributes adds the attributes to the list dashboard version params
func (o *ListDashboardVersionParams) WithAttributes(attributes *string) *ListDashboardVersionParams {
	o.SetAttributes(attributes)
	return o
}

// SetAttributes adds the attributes to the list dashboard version params
func (o *ListDashboardVersionParams) SetAttributes(attributes *string) {
	o.Attributes = attributes
}

// WithEnddate adds the enddate to the list dashboard version params
func (o *ListDashboardVersionParams) WithEnddate(enddate *string) *ListDashboardVersionParams {
	o.SetEnddate(enddate)
	return o
}

// SetEnddate adds the enddate to the list dashboard version params
func (o *ListDashboardVersionParams) SetEnddate(enddate *string) {
	o.Enddate = enddate
}

// WithFields adds the fields to the list dashboard version params
func (o *ListDashboardVersionParams) WithFields(fields *string) *ListDashboardVersionParams {
	o.SetFields(fields)
	return o
}

// SetFields adds the fields to the list dashboard version params
func (o *ListDashboardVersionParams) SetFields(fields *string) {
	o.Fields = fields
}

// WithGroupby adds the groupby to the list dashboard version params
func (o *ListDashboardVersionParams) WithGroupby(groupby *string) *ListDashboardVersionParams {
	o.SetGroupby(groupby)
	return o
}

// SetGroupby adds the groupby to the list dashboard version params
func (o *ListDashboardVersionParams) SetGroupby(groupby *string) {
	o.Groupby = groupby
}

// WithLimit adds the limit to the list dashboard version params
func (o *ListDashboardVersionParams) WithLimit(limit *int32) *ListDashboardVersionParams {
	o.SetLimit(limit)
	return o
}

// SetLimit adds the limit to the list dashboard version params
func (o *ListDashboardVersionParams) SetLimit(limit *int32) {
	o.Limit = limit
}

// WithOrderby adds the orderby to the list dashboard version params
func (o *ListDashboardVersionParams) WithOrderby(orderby *string) *ListDashboardVersionParams {
	o.SetOrderby(orderby)
	return o
}

// SetOrderby adds the orderby to the list dashboard version params
func (o *ListDashboardVersionParams) SetOrderby(orderby *string) {
	o.Orderby = orderby
}

// WithPerformanceindicators adds the performanceindicators to the list dashboard version params
func (o *ListDashboardVersionParams) WithPerformanceindicators(performanceindicators *string) *ListDashboardVersionParams {
	o.SetPerformanceindicators(performanceindicators)
	return o
}

// SetPerformanceindicators adds the performanceindicators to the list dashboard version params
func (o *ListDashboardVersionParams) SetPerformanceindicators(performanceindicators *string) {
	o.Performanceindicators = performanceindicators
}

// WithStart adds the start to the list dashboard version params
func (o *ListDashboardVersionParams) WithStart(start *int32) *ListDashboardVersionParams {
	o.SetStart(start)
	return o
}

// SetStart adds the start to the list dashboard version params
func (o *ListDashboardVersionParams) SetStart(start *int32) {
	o.Start = start
}

// WithStartdate adds the startdate to the list dashboard version params
func (o *ListDashboardVersionParams) WithStartdate(startdate *string) *ListDashboardVersionParams {
	o.SetStartdate(startdate)
	return o
}

// SetStartdate adds the startdate to the list dashboard version params
func (o *ListDashboardVersionParams) SetStartdate(startdate *string) {
	o.Startdate = startdate
}

// WithVariables adds the variables to the list dashboard version params
func (o *ListDashboardVersionParams) WithVariables(variables *string) *ListDashboardVersionParams {
	o.SetVariables(variables)
	return o
}

// SetVariables adds the variables to the list dashboard version params
func (o *ListDashboardVersionParams) SetVariables(variables *string) {
	o.Variables = variables
}

// WriteToRequest writes these params to a swagger request
func (o *ListDashboardVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Aggregateby != nil {

		// query param aggregateby
		var qrAggregateby string
		if o.Aggregateby != nil {
			qrAggregateby = *o.Aggregateby
		}
		qAggregateby := qrAggregateby
		if qAggregateby != "" {
			if err := r.SetQueryParam("aggregateby", qAggregateby); err != nil {
				return err
			}
		}

	}

	if o.Attributefilter != nil {

		// query param attributefilter
		var qrAttributefilter string
		if o.Attributefilter != nil {
			qrAttributefilter = *o.Attributefilter
		}
		qAttributefilter := qrAttributefilter
		if qAttributefilter != "" {
			if err := r.SetQueryParam("attributefilter", qAttributefilter); err != nil {
				return err
			}
		}

	}

	if o.Attributes != nil {

		// query param attributes
		var qrAttributes string
		if o.Attributes != nil {
			qrAttributes = *o.Attributes
		}
		qAttributes := qrAttributes
		if qAttributes != "" {
			if err := r.SetQueryParam("attributes", qAttributes); err != nil {
				return err
			}
		}

	}

	if o.Enddate != nil {

		// query param enddate
		var qrEnddate string
		if o.Enddate != nil {
			qrEnddate = *o.Enddate
		}
		qEnddate := qrEnddate
		if qEnddate != "" {
			if err := r.SetQueryParam("enddate", qEnddate); err != nil {
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

	if o.Groupby != nil {

		// query param groupby
		var qrGroupby string
		if o.Groupby != nil {
			qrGroupby = *o.Groupby
		}
		qGroupby := qrGroupby
		if qGroupby != "" {
			if err := r.SetQueryParam("groupby", qGroupby); err != nil {
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

	if o.Performanceindicators != nil {

		// query param performanceindicators
		var qrPerformanceindicators string
		if o.Performanceindicators != nil {
			qrPerformanceindicators = *o.Performanceindicators
		}
		qPerformanceindicators := qrPerformanceindicators
		if qPerformanceindicators != "" {
			if err := r.SetQueryParam("performanceindicators", qPerformanceindicators); err != nil {
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

	if o.Startdate != nil {

		// query param startdate
		var qrStartdate string
		if o.Startdate != nil {
			qrStartdate = *o.Startdate
		}
		qStartdate := qrStartdate
		if qStartdate != "" {
			if err := r.SetQueryParam("startdate", qStartdate); err != nil {
				return err
			}
		}

	}

	if o.Variables != nil {

		// query param variables
		var qrVariables string
		if o.Variables != nil {
			qrVariables = *o.Variables
		}
		qVariables := qrVariables
		if qVariables != "" {
			if err := r.SetQueryParam("variables", qVariables); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
