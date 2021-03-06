// Code generated by go-swagger; DO NOT EDIT.

package rulepack_update_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDoRulepackUpdateParams creates a new DoRulepackUpdateParams object
// with the default values initialized.
func NewDoRulepackUpdateParams() *DoRulepackUpdateParams {

	return &DoRulepackUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDoRulepackUpdateParamsWithTimeout creates a new DoRulepackUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDoRulepackUpdateParamsWithTimeout(timeout time.Duration) *DoRulepackUpdateParams {

	return &DoRulepackUpdateParams{

		timeout: timeout,
	}
}

// NewDoRulepackUpdateParamsWithContext creates a new DoRulepackUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDoRulepackUpdateParamsWithContext(ctx context.Context) *DoRulepackUpdateParams {

	return &DoRulepackUpdateParams{

		Context: ctx,
	}
}

// NewDoRulepackUpdateParamsWithHTTPClient creates a new DoRulepackUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDoRulepackUpdateParamsWithHTTPClient(client *http.Client) *DoRulepackUpdateParams {

	return &DoRulepackUpdateParams{
		HTTPClient: client,
	}
}

/*DoRulepackUpdateParams contains all the parameters to send to the API endpoint
for the do rulepack update operation typically these are written to a http.Request
*/
type DoRulepackUpdateParams struct {
	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the do rulepack update params
func (o *DoRulepackUpdateParams) WithTimeout(timeout time.Duration) *DoRulepackUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the do rulepack update params
func (o *DoRulepackUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the do rulepack update params
func (o *DoRulepackUpdateParams) WithContext(ctx context.Context) *DoRulepackUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the do rulepack update params
func (o *DoRulepackUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the do rulepack update params
func (o *DoRulepackUpdateParams) WithHTTPClient(client *http.Client) *DoRulepackUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the do rulepack update params
func (o *DoRulepackUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WriteToRequest writes these params to a swagger request
func (o *DoRulepackUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
