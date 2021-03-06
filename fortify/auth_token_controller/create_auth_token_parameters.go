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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// NewCreateAuthTokenParams creates a new CreateAuthTokenParams object
// with the default values initialized.
func NewCreateAuthTokenParams() *CreateAuthTokenParams {
	var ()
	return &CreateAuthTokenParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCreateAuthTokenParamsWithTimeout creates a new CreateAuthTokenParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCreateAuthTokenParamsWithTimeout(timeout time.Duration) *CreateAuthTokenParams {
	var ()
	return &CreateAuthTokenParams{

		timeout: timeout,
	}
}

// NewCreateAuthTokenParamsWithContext creates a new CreateAuthTokenParams object
// with the default values initialized, and the ability to set a context for a request
func NewCreateAuthTokenParamsWithContext(ctx context.Context) *CreateAuthTokenParams {
	var ()
	return &CreateAuthTokenParams{

		Context: ctx,
	}
}

// NewCreateAuthTokenParamsWithHTTPClient creates a new CreateAuthTokenParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCreateAuthTokenParamsWithHTTPClient(client *http.Client) *CreateAuthTokenParams {
	var ()
	return &CreateAuthTokenParams{
		HTTPClient: client,
	}
}

/*CreateAuthTokenParams contains all the parameters to send to the API endpoint
for the create auth token operation typically these are written to a http.Request
*/
type CreateAuthTokenParams struct {

	/*AuthToken
	  authToken

	*/
	AuthToken *models.AuthenticationToken

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the create auth token params
func (o *CreateAuthTokenParams) WithTimeout(timeout time.Duration) *CreateAuthTokenParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create auth token params
func (o *CreateAuthTokenParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create auth token params
func (o *CreateAuthTokenParams) WithContext(ctx context.Context) *CreateAuthTokenParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create auth token params
func (o *CreateAuthTokenParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create auth token params
func (o *CreateAuthTokenParams) WithHTTPClient(client *http.Client) *CreateAuthTokenParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create auth token params
func (o *CreateAuthTokenParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthToken adds the authToken to the create auth token params
func (o *CreateAuthTokenParams) WithAuthToken(authToken *models.AuthenticationToken) *CreateAuthTokenParams {
	o.SetAuthToken(authToken)
	return o
}

// SetAuthToken adds the authToken to the create auth token params
func (o *CreateAuthTokenParams) SetAuthToken(authToken *models.AuthenticationToken) {
	o.AuthToken = authToken
}

// WriteToRequest writes these params to a swagger request
func (o *CreateAuthTokenParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.AuthToken != nil {
		if err := r.SetBodyParam(o.AuthToken); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
