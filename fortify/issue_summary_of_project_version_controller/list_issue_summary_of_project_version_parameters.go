// Code generated by go-swagger; DO NOT EDIT.

package issue_summary_of_project_version_controller

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

// NewListIssueSummaryOfProjectVersionParams creates a new ListIssueSummaryOfProjectVersionParams object
// with the default values initialized.
func NewListIssueSummaryOfProjectVersionParams() *ListIssueSummaryOfProjectVersionParams {
	var ()
	return &ListIssueSummaryOfProjectVersionParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewListIssueSummaryOfProjectVersionParamsWithTimeout creates a new ListIssueSummaryOfProjectVersionParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewListIssueSummaryOfProjectVersionParamsWithTimeout(timeout time.Duration) *ListIssueSummaryOfProjectVersionParams {
	var ()
	return &ListIssueSummaryOfProjectVersionParams{

		timeout: timeout,
	}
}

// NewListIssueSummaryOfProjectVersionParamsWithContext creates a new ListIssueSummaryOfProjectVersionParams object
// with the default values initialized, and the ability to set a context for a request
func NewListIssueSummaryOfProjectVersionParamsWithContext(ctx context.Context) *ListIssueSummaryOfProjectVersionParams {
	var ()
	return &ListIssueSummaryOfProjectVersionParams{

		Context: ctx,
	}
}

// NewListIssueSummaryOfProjectVersionParamsWithHTTPClient creates a new ListIssueSummaryOfProjectVersionParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewListIssueSummaryOfProjectVersionParamsWithHTTPClient(client *http.Client) *ListIssueSummaryOfProjectVersionParams {
	var ()
	return &ListIssueSummaryOfProjectVersionParams{
		HTTPClient: client,
	}
}

/*ListIssueSummaryOfProjectVersionParams contains all the parameters to send to the API endpoint
for the list issue summary of project version operation typically these are written to a http.Request
*/
type ListIssueSummaryOfProjectVersionParams struct {

	/*Audited
	  audited

	*/
	Audited *string
	/*Filter
	  filter

	*/
	Filter *string
	/*Groupaxistype
	  List of allowed patterns: APP_NAME, SCAN_DATE, SCAN_PRODUCT, ISSUE_FOLDER, ISSUE_CATEGORY, ISSUE_KINGDOM, ISSUE_FILENAME, ISSUE_FRIORITY, ISSUE_AUDITED, ISSUE_PACKAGE_NAME, ISSUE_CLASS_NAME, ISSUE_FUNCTION_NAME, ISSUE_MAPPED_CATEGORY, FOLDER_FOLDER, ISSUE_{name of issue attribute}, EXTERNALLIST_{external category name}, CUSTOMTAG_{custom tag name}.

	*/
	Groupaxistype string
	/*ParentID
	  parentId

	*/
	ParentID int64
	/*Seriestype
	  seriestype

	*/
	Seriestype string
	/*Showhidden
	  showhidden

	*/
	Showhidden *bool
	/*Showremoved
	  showremoved

	*/
	Showremoved *bool
	/*Showsuppressed
	  showsuppressed

	*/
	Showsuppressed *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithTimeout(timeout time.Duration) *ListIssueSummaryOfProjectVersionParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithContext(ctx context.Context) *ListIssueSummaryOfProjectVersionParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithHTTPClient(client *http.Client) *ListIssueSummaryOfProjectVersionParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAudited adds the audited to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithAudited(audited *string) *ListIssueSummaryOfProjectVersionParams {
	o.SetAudited(audited)
	return o
}

// SetAudited adds the audited to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetAudited(audited *string) {
	o.Audited = audited
}

// WithFilter adds the filter to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithFilter(filter *string) *ListIssueSummaryOfProjectVersionParams {
	o.SetFilter(filter)
	return o
}

// SetFilter adds the filter to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetFilter(filter *string) {
	o.Filter = filter
}

// WithGroupaxistype adds the groupaxistype to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithGroupaxistype(groupaxistype string) *ListIssueSummaryOfProjectVersionParams {
	o.SetGroupaxistype(groupaxistype)
	return o
}

// SetGroupaxistype adds the groupaxistype to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetGroupaxistype(groupaxistype string) {
	o.Groupaxistype = groupaxistype
}

// WithParentID adds the parentID to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithParentID(parentID int64) *ListIssueSummaryOfProjectVersionParams {
	o.SetParentID(parentID)
	return o
}

// SetParentID adds the parentId to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetParentID(parentID int64) {
	o.ParentID = parentID
}

// WithSeriestype adds the seriestype to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithSeriestype(seriestype string) *ListIssueSummaryOfProjectVersionParams {
	o.SetSeriestype(seriestype)
	return o
}

// SetSeriestype adds the seriestype to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetSeriestype(seriestype string) {
	o.Seriestype = seriestype
}

// WithShowhidden adds the showhidden to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithShowhidden(showhidden *bool) *ListIssueSummaryOfProjectVersionParams {
	o.SetShowhidden(showhidden)
	return o
}

// SetShowhidden adds the showhidden to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetShowhidden(showhidden *bool) {
	o.Showhidden = showhidden
}

// WithShowremoved adds the showremoved to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithShowremoved(showremoved *bool) *ListIssueSummaryOfProjectVersionParams {
	o.SetShowremoved(showremoved)
	return o
}

// SetShowremoved adds the showremoved to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetShowremoved(showremoved *bool) {
	o.Showremoved = showremoved
}

// WithShowsuppressed adds the showsuppressed to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) WithShowsuppressed(showsuppressed *bool) *ListIssueSummaryOfProjectVersionParams {
	o.SetShowsuppressed(showsuppressed)
	return o
}

// SetShowsuppressed adds the showsuppressed to the list issue summary of project version params
func (o *ListIssueSummaryOfProjectVersionParams) SetShowsuppressed(showsuppressed *bool) {
	o.Showsuppressed = showsuppressed
}

// WriteToRequest writes these params to a swagger request
func (o *ListIssueSummaryOfProjectVersionParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Audited != nil {

		// query param audited
		var qrAudited string
		if o.Audited != nil {
			qrAudited = *o.Audited
		}
		qAudited := qrAudited
		if qAudited != "" {
			if err := r.SetQueryParam("audited", qAudited); err != nil {
				return err
			}
		}

	}

	if o.Filter != nil {

		// query param filter
		var qrFilter string
		if o.Filter != nil {
			qrFilter = *o.Filter
		}
		qFilter := qrFilter
		if qFilter != "" {
			if err := r.SetQueryParam("filter", qFilter); err != nil {
				return err
			}
		}

	}

	// query param groupaxistype
	qrGroupaxistype := o.Groupaxistype
	qGroupaxistype := qrGroupaxistype
	if qGroupaxistype != "" {
		if err := r.SetQueryParam("groupaxistype", qGroupaxistype); err != nil {
			return err
		}
	}

	// path param parentId
	if err := r.SetPathParam("parentId", swag.FormatInt64(o.ParentID)); err != nil {
		return err
	}

	// query param seriestype
	qrSeriestype := o.Seriestype
	qSeriestype := qrSeriestype
	if qSeriestype != "" {
		if err := r.SetQueryParam("seriestype", qSeriestype); err != nil {
			return err
		}
	}

	if o.Showhidden != nil {

		// query param showhidden
		var qrShowhidden bool
		if o.Showhidden != nil {
			qrShowhidden = *o.Showhidden
		}
		qShowhidden := swag.FormatBool(qrShowhidden)
		if qShowhidden != "" {
			if err := r.SetQueryParam("showhidden", qShowhidden); err != nil {
				return err
			}
		}

	}

	if o.Showremoved != nil {

		// query param showremoved
		var qrShowremoved bool
		if o.Showremoved != nil {
			qrShowremoved = *o.Showremoved
		}
		qShowremoved := swag.FormatBool(qrShowremoved)
		if qShowremoved != "" {
			if err := r.SetQueryParam("showremoved", qShowremoved); err != nil {
				return err
			}
		}

	}

	if o.Showsuppressed != nil {

		// query param showsuppressed
		var qrShowsuppressed bool
		if o.Showsuppressed != nil {
			qrShowsuppressed = *o.Showsuppressed
		}
		qShowsuppressed := swag.FormatBool(qrShowsuppressed)
		if qShowsuppressed != "" {
			if err := r.SetQueryParam("showsuppressed", qShowsuppressed); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
