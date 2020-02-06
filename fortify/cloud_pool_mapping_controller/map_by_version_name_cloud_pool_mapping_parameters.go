// Code generated by go-swagger; DO NOT EDIT.

package cloud_pool_mapping_controller

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

// NewMapByVersionNameCloudPoolMappingParams creates a new MapByVersionNameCloudPoolMappingParams object
// with the default values initialized.
func NewMapByVersionNameCloudPoolMappingParams() *MapByVersionNameCloudPoolMappingParams {
	var ()
	return &MapByVersionNameCloudPoolMappingParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewMapByVersionNameCloudPoolMappingParamsWithTimeout creates a new MapByVersionNameCloudPoolMappingParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewMapByVersionNameCloudPoolMappingParamsWithTimeout(timeout time.Duration) *MapByVersionNameCloudPoolMappingParams {
	var ()
	return &MapByVersionNameCloudPoolMappingParams{

		timeout: timeout,
	}
}

// NewMapByVersionNameCloudPoolMappingParamsWithContext creates a new MapByVersionNameCloudPoolMappingParams object
// with the default values initialized, and the ability to set a context for a request
func NewMapByVersionNameCloudPoolMappingParamsWithContext(ctx context.Context) *MapByVersionNameCloudPoolMappingParams {
	var ()
	return &MapByVersionNameCloudPoolMappingParams{

		Context: ctx,
	}
}

// NewMapByVersionNameCloudPoolMappingParamsWithHTTPClient creates a new MapByVersionNameCloudPoolMappingParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewMapByVersionNameCloudPoolMappingParamsWithHTTPClient(client *http.Client) *MapByVersionNameCloudPoolMappingParams {
	var ()
	return &MapByVersionNameCloudPoolMappingParams{
		HTTPClient: client,
	}
}

/*MapByVersionNameCloudPoolMappingParams contains all the parameters to send to the API endpoint
for the map by version name cloud pool mapping operation typically these are written to a http.Request
*/
type MapByVersionNameCloudPoolMappingParams struct {

	/*ProjectName
	  projectName

	*/
	ProjectName string
	/*ProjectVersionName
	  projectVersionName

	*/
	ProjectVersionName string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) WithTimeout(timeout time.Duration) *MapByVersionNameCloudPoolMappingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) WithContext(ctx context.Context) *MapByVersionNameCloudPoolMappingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) WithHTTPClient(client *http.Client) *MapByVersionNameCloudPoolMappingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithProjectName adds the projectName to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) WithProjectName(projectName string) *MapByVersionNameCloudPoolMappingParams {
	o.SetProjectName(projectName)
	return o
}

// SetProjectName adds the projectName to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) SetProjectName(projectName string) {
	o.ProjectName = projectName
}

// WithProjectVersionName adds the projectVersionName to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) WithProjectVersionName(projectVersionName string) *MapByVersionNameCloudPoolMappingParams {
	o.SetProjectVersionName(projectVersionName)
	return o
}

// SetProjectVersionName adds the projectVersionName to the map by version name cloud pool mapping params
func (o *MapByVersionNameCloudPoolMappingParams) SetProjectVersionName(projectVersionName string) {
	o.ProjectVersionName = projectVersionName
}

// WriteToRequest writes these params to a swagger request
func (o *MapByVersionNameCloudPoolMappingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// query param projectName
	qrProjectName := o.ProjectName
	qProjectName := qrProjectName
	if qProjectName != "" {
		if err := r.SetQueryParam("projectName", qProjectName); err != nil {
			return err
		}
	}

	// query param projectVersionName
	qrProjectVersionName := o.ProjectVersionName
	qProjectVersionName := qrProjectVersionName
	if qProjectVersionName != "" {
		if err := r.SetQueryParam("projectVersionName", qProjectVersionName); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}