// Code generated by go-swagger; DO NOT EDIT.

package alert_history_controller

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

	"github.com/piper-validation/fortify-client-go/models"
)

// NewSetStatusForAlertHistoryParams creates a new SetStatusForAlertHistoryParams object
// with the default values initialized.
func NewSetStatusForAlertHistoryParams() *SetStatusForAlertHistoryParams {
	var ()
	return &SetStatusForAlertHistoryParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSetStatusForAlertHistoryParamsWithTimeout creates a new SetStatusForAlertHistoryParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSetStatusForAlertHistoryParamsWithTimeout(timeout time.Duration) *SetStatusForAlertHistoryParams {
	var ()
	return &SetStatusForAlertHistoryParams{

		timeout: timeout,
	}
}

// NewSetStatusForAlertHistoryParamsWithContext creates a new SetStatusForAlertHistoryParams object
// with the default values initialized, and the ability to set a context for a request
func NewSetStatusForAlertHistoryParamsWithContext(ctx context.Context) *SetStatusForAlertHistoryParams {
	var ()
	return &SetStatusForAlertHistoryParams{

		Context: ctx,
	}
}

// NewSetStatusForAlertHistoryParamsWithHTTPClient creates a new SetStatusForAlertHistoryParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSetStatusForAlertHistoryParamsWithHTTPClient(client *http.Client) *SetStatusForAlertHistoryParams {
	var ()
	return &SetStatusForAlertHistoryParams{
		HTTPClient: client,
	}
}

/*SetStatusForAlertHistoryParams contains all the parameters to send to the API endpoint
for the set status for alert history operation typically these are written to a http.Request
*/
type SetStatusForAlertHistoryParams struct {

	/*Resource
	  resource

	*/
	Resource *models.AlertSetStatusRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) WithTimeout(timeout time.Duration) *SetStatusForAlertHistoryParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) WithContext(ctx context.Context) *SetStatusForAlertHistoryParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) WithHTTPClient(client *http.Client) *SetStatusForAlertHistoryParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) WithResource(resource *models.AlertSetStatusRequest) *SetStatusForAlertHistoryParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the set status for alert history params
func (o *SetStatusForAlertHistoryParams) SetResource(resource *models.AlertSetStatusRequest) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *SetStatusForAlertHistoryParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Resource != nil {
		if err := r.SetBodyParam(o.Resource); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}