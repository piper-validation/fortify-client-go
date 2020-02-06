// Code generated by go-swagger; DO NOT EDIT.

package local_user_controller

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

// NewCheckLocalUserParams creates a new CheckLocalUserParams object
// with the default values initialized.
func NewCheckLocalUserParams() *CheckLocalUserParams {
	var ()
	return &CheckLocalUserParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCheckLocalUserParamsWithTimeout creates a new CheckLocalUserParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCheckLocalUserParamsWithTimeout(timeout time.Duration) *CheckLocalUserParams {
	var ()
	return &CheckLocalUserParams{

		timeout: timeout,
	}
}

// NewCheckLocalUserParamsWithContext creates a new CheckLocalUserParams object
// with the default values initialized, and the ability to set a context for a request
func NewCheckLocalUserParamsWithContext(ctx context.Context) *CheckLocalUserParams {
	var ()
	return &CheckLocalUserParams{

		Context: ctx,
	}
}

// NewCheckLocalUserParamsWithHTTPClient creates a new CheckLocalUserParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCheckLocalUserParamsWithHTTPClient(client *http.Client) *CheckLocalUserParams {
	var ()
	return &CheckLocalUserParams{
		HTTPClient: client,
	}
}

/*CheckLocalUserParams contains all the parameters to send to the API endpoint
for the check local user operation typically these are written to a http.Request
*/
type CheckLocalUserParams struct {

	/*PasswordStrengthCheckRequest
	  passwordStrengthCheckRequest

	*/
	PasswordStrengthCheckRequest *models.PasswordStrengthCheckRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the check local user params
func (o *CheckLocalUserParams) WithTimeout(timeout time.Duration) *CheckLocalUserParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the check local user params
func (o *CheckLocalUserParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the check local user params
func (o *CheckLocalUserParams) WithContext(ctx context.Context) *CheckLocalUserParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the check local user params
func (o *CheckLocalUserParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the check local user params
func (o *CheckLocalUserParams) WithHTTPClient(client *http.Client) *CheckLocalUserParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the check local user params
func (o *CheckLocalUserParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithPasswordStrengthCheckRequest adds the passwordStrengthCheckRequest to the check local user params
func (o *CheckLocalUserParams) WithPasswordStrengthCheckRequest(passwordStrengthCheckRequest *models.PasswordStrengthCheckRequest) *CheckLocalUserParams {
	o.SetPasswordStrengthCheckRequest(passwordStrengthCheckRequest)
	return o
}

// SetPasswordStrengthCheckRequest adds the passwordStrengthCheckRequest to the check local user params
func (o *CheckLocalUserParams) SetPasswordStrengthCheckRequest(passwordStrengthCheckRequest *models.PasswordStrengthCheckRequest) {
	o.PasswordStrengthCheckRequest = passwordStrengthCheckRequest
}

// WriteToRequest writes these params to a swagger request
func (o *CheckLocalUserParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.PasswordStrengthCheckRequest != nil {
		if err := r.SetBodyParam(o.PasswordStrengthCheckRequest); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}