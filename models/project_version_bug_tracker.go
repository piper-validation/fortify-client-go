// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ProjectVersionBugTracker Object containing application version and bug tracker related data
// swagger:model Project version bug tracker
type ProjectVersionBugTracker struct {

	// the short display name of the bug tracker. (will be null if no bugtracker is assigned or if there is no currently installed plugin matching the assigned pluginId.)
	// Required: true
	AssignedBugtrackerShortNameIfKnown *string `json:"assignedBugtrackerShortNameIfKnown"`

	// identifier of the bug tracker plugin assigned to the application version. (Bug tracker integration will be active only if the plugin is also enabled in the system.)
	// Required: true
	AssignedPluginID *string `json:"assignedPluginId"`

	// Bug state management configuration
	// Required: true
	BugStateManagementConfig *BugStateManagementCfg `json:"bugStateManagementConfig"`

	// Bug tracker (note that this field will be null if the assigned bug tracker is not enabled in the SSC plugin framework.)
	// Required: true
	BugTracker *BugTracker `json:"bugTracker"`

	// indicates whether a bug tracker is assigned to the application version
	// Required: true
	BugtrackerIsAssigned *bool `json:"bugtrackerIsAssigned"`

	// clear app version bugs
	ClearAppVersionBugs bool `json:"clearAppVersionBugs,omitempty"`
}

// Validate validates this project version bug tracker
func (m *ProjectVersionBugTracker) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAssignedBugtrackerShortNameIfKnown(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssignedPluginID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBugStateManagementConfig(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBugTracker(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBugtrackerIsAssigned(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVersionBugTracker) validateAssignedBugtrackerShortNameIfKnown(formats strfmt.Registry) error {

	if err := validate.Required("assignedBugtrackerShortNameIfKnown", "body", m.AssignedBugtrackerShortNameIfKnown); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionBugTracker) validateAssignedPluginID(formats strfmt.Registry) error {

	if err := validate.Required("assignedPluginId", "body", m.AssignedPluginID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionBugTracker) validateBugStateManagementConfig(formats strfmt.Registry) error {

	if err := validate.Required("bugStateManagementConfig", "body", m.BugStateManagementConfig); err != nil {
		return err
	}

	if m.BugStateManagementConfig != nil {
		if err := m.BugStateManagementConfig.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bugStateManagementConfig")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectVersionBugTracker) validateBugTracker(formats strfmt.Registry) error {

	if err := validate.Required("bugTracker", "body", m.BugTracker); err != nil {
		return err
	}

	if m.BugTracker != nil {
		if err := m.BugTracker.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bugTracker")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectVersionBugTracker) validateBugtrackerIsAssigned(formats strfmt.Registry) error {

	if err := validate.Required("bugtrackerIsAssigned", "body", m.BugtrackerIsAssigned); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectVersionBugTracker) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVersionBugTracker) UnmarshalBinary(b []byte) error {
	var res ProjectVersionBugTracker
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}