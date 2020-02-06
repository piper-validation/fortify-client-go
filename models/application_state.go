// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// ApplicationState Application state DTO object
// swagger:model Application State
type ApplicationState struct {

	// Indicates whether SSC application is in auto configuration mode
	// Read Only: true
	AutoConfigurationMode *bool `json:"autoConfigurationMode,omitempty"`

	// Indicates whether a config visit is needed
	// Read Only: true
	ConfigVisitRequired *bool `json:"configVisitRequired,omitempty"`

	// Indicates whether SSC application is in maintenance mode
	MaintenanceMode bool `json:"maintenanceMode,omitempty"`

	// Indicates whether a restart is needed
	// Read Only: true
	RestartRequired *bool `json:"restartRequired,omitempty"`
}

// Validate validates this application state
func (m *ApplicationState) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationState) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationState) UnmarshalBinary(b []byte) error {
	var res ApplicationState
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}