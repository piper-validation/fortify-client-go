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

// AssignedUser user assigned to issue
// swagger:model AssignedUser
type AssignedUser struct {

	// whether the assigned user currently has access to the application version in current context
	// Required: true
	// Read Only: true
	HasAccess bool `json:"hasAccess"`

	// username assigned to issue
	// Required: true
	// Read Only: true
	UserName string `json:"userName"`
}

// Validate validates this assigned user
func (m *AssignedUser) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHasAccess(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *AssignedUser) validateHasAccess(formats strfmt.Registry) error {

	if err := validate.Required("hasAccess", "body", bool(m.HasAccess)); err != nil {
		return err
	}

	return nil
}

func (m *AssignedUser) validateUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("userName", "body", string(m.UserName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AssignedUser) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AssignedUser) UnmarshalBinary(b []byte) error {
	var res AssignedUser
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
