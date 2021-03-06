// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CloudPoolProjectVersionActionResponse Result of cloud pool actions relating to application versions
// swagger:model CloudPoolProjectVersionActionResponse
type CloudPoolProjectVersionActionResponse struct {

	// List of application version ids
	// Read Only: true
	ProjectVersionIds []int64 `json:"projectVersionIds"`
}

// Validate validates this cloud pool project version action response
func (m *CloudPoolProjectVersionActionResponse) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *CloudPoolProjectVersionActionResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CloudPoolProjectVersionActionResponse) UnmarshalBinary(b []byte) error {
	var res CloudPoolProjectVersionActionResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
