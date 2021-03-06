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

// IssueAssignUserRequest Request to assign user to the issues
// swagger:model IssueAssignUserRequest
type IssueAssignUserRequest struct {

	// Issues to audit
	// Required: true
	Issues []*EntityStateIdentifier `json:"issues"`

	// Username to assign
	// Required: true
	User *string `json:"user"`
}

// Validate validates this issue assign user request
func (m *IssueAssignUserRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUser(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueAssignUserRequest) validateIssues(formats strfmt.Registry) error {

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

func (m *IssueAssignUserRequest) validateUser(formats strfmt.Registry) error {

	if err := validate.Required("user", "body", m.User); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueAssignUserRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAssignUserRequest) UnmarshalBinary(b []byte) error {
	var res IssueAssignUserRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
