// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ObsoletePolicy Audit Assistant obsolete prediction policy
// swagger:model Obsolete policy
type ObsoletePolicy struct {

	// True if this obsolete policy is also configured as system default policy
	DefaultPolicy bool `json:"defaultPolicy,omitempty"`

	// Count of application versions where this obsolete policy has been found
	// Read Only: true
	ObsoleteAvCount int64 `json:"obsoleteAvCount,omitempty"`

	// Obsolete prediction policy name found in application version
	// Read Only: true
	PolicyName string `json:"policyName,omitempty"`
}

// Validate validates this obsolete policy
func (m *ObsoletePolicy) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ObsoletePolicy) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ObsoletePolicy) UnmarshalBinary(b []byte) error {
	var res ObsoletePolicy
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}