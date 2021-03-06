// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// IssueSuppressRequest Request to suppress issues
// swagger:model IssueSuppressRequest
type IssueSuppressRequest struct {

	// Issues to suppress
	// Required: true
	Issues []*EntityStateIdentifier `json:"issues"`

	// Will suppress the issue
	// Required: true
	Suppressed *bool `json:"suppressed"`
}

// Validate validates this issue suppress request
func (m *IssueSuppressRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuppressed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueSuppressRequest) validateIssues(formats strfmt.Registry) error {

	if err := validate.Required("issues", "body", m.Issues); err != nil {
		return err
	}

	for i := 0; i < len(m.Issues); i++ {
		if swag.IsZero(m.Issues[i]) { // not required
			continue
		}

		if m.Issues[i] != nil {
			if err := m.Issues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("issues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IssueSuppressRequest) validateSuppressed(formats strfmt.Registry) error {

	if err := validate.Required("suppressed", "body", m.Suppressed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueSuppressRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueSuppressRequest) UnmarshalBinary(b []byte) error {
	var res IssueSuppressRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
