// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// BugFilingRequirementsResponse Response of validating connection credentials and returns bug filing requirements
// swagger:model BugFilingRequirementsResponse
type BugFilingRequirementsResponse struct {

	// Requirements to file the bugs
	// Read Only: true
	BugFilingRequirements *BugFilingRequirements `json:"bugFilingRequirements,omitempty"`
}

// Validate validates this bug filing requirements response
func (m *BugFilingRequirementsResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBugFilingRequirements(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *BugFilingRequirementsResponse) validateBugFilingRequirements(formats strfmt.Registry) error {

	if swag.IsZero(m.BugFilingRequirements) { // not required
		return nil
	}

	if m.BugFilingRequirements != nil {
		if err := m.BugFilingRequirements.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("bugFilingRequirements")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BugFilingRequirementsResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BugFilingRequirementsResponse) UnmarshalBinary(b []byte) error {
	var res BugFilingRequirementsResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
