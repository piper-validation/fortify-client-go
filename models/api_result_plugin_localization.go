// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIResultPluginLocalization Api result plugin localization
// swagger:model ApiResult«Plugin Localization»
type APIResultPluginLocalization struct {

	// count
	Count int64 `json:"count,omitempty"`

	// data
	Data map[string]interface{} `json:"data,omitempty"`

	// error code
	ErrorCode int32 `json:"errorCode,omitempty"`

	// links
	Links interface{} `json:"links,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// response code
	ResponseCode int32 `json:"responseCode,omitempty"`

	// stack trace
	StackTrace string `json:"stackTrace,omitempty"`

	// success count
	SuccessCount int64 `json:"successCount,omitempty"`
}

// Validate validates this Api result plugin localization
func (m *APIResultPluginLocalization) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *APIResultPluginLocalization) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResultPluginLocalization) UnmarshalBinary(b []byte) error {
	var res APIResultPluginLocalization
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
