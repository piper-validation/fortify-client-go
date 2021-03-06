// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// EngineType Types of the analyzers (engines) that produces issues
// swagger:model Engine type
type EngineType struct {

	// name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// served by plugin
	// Read Only: true
	ServedByPlugin *bool `json:"servedByPlugin,omitempty"`
}

// Validate validates this engine type
func (m *EngineType) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *EngineType) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EngineType) UnmarshalBinary(b []byte) error {
	var res EngineType
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
