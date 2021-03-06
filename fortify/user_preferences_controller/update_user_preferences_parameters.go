// Code generated by go-swagger; DO NOT EDIT.

package user_preferences_controller

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

// NewUpdateUserPreferencesParams creates a new UpdateUserPreferencesParams object
// with the default values initialized.
func NewUpdateUserPreferencesParams() *UpdateUserPreferencesParams {
	var ()
	return &UpdateUserPreferencesParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateUserPreferencesParamsWithTimeout creates a new UpdateUserPreferencesParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateUserPreferencesParamsWithTimeout(timeout time.Duration) *UpdateUserPreferencesParams {
	var ()
	return &UpdateUserPreferencesParams{

		timeout: timeout,
	}
}

// NewUpdateUserPreferencesParamsWithContext creates a new UpdateUserPreferencesParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateUserPreferencesParamsWithContext(ctx context.Context) *UpdateUserPreferencesParams {
	var ()
	return &UpdateUserPreferencesParams{

		Context: ctx,
	}
}

// NewUpdateUserPreferencesParamsWithHTTPClient creates a new UpdateUserPreferencesParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateUserPreferencesParamsWithHTTPClient(client *http.Client) *UpdateUserPreferencesParams {
	var ()
	return &UpdateUserPreferencesParams{
		HTTPClient: client,
	}
}

/*UpdateUserPreferencesParams contains all the parameters to send to the API endpoint
for the update user preferences operation typically these are written to a http.Request
*/
type UpdateUserPreferencesParams struct {

	/*Resource
	  resource

	*/
	Resource *models.UserPreferences

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update user preferences params
func (o *UpdateUserPreferencesParams) WithTimeout(timeout time.Duration) *UpdateUserPreferencesParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update user preferences params
func (o *UpdateUserPreferencesParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update user preferences params
func (o *UpdateUserPreferencesParams) WithContext(ctx context.Context) *UpdateUserPreferencesParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update user preferences params
func (o *UpdateUserPreferencesParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update user preferences params
func (o *UpdateUserPreferencesParams) WithHTTPClient(client *http.Client) *UpdateUserPreferencesParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update user preferences params
func (o *UpdateUserPreferencesParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithResource adds the resource to the update user preferences params
func (o *UpdateUserPreferencesParams) WithResource(resource *models.UserPreferences) *UpdateUserPreferencesParams {
	o.SetResource(resource)
	return o
}

// SetResource adds the resource to the update user preferences params
func (o *UpdateUserPreferencesParams) SetResource(resource *models.UserPreferences) {
	o.Resource = resource
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateUserPreferencesParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
