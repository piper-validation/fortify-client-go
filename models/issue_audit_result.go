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

// IssueAuditResult Issue audit result contains the most rescent issue audit information after performing issue audit.
// swagger:model IssueAuditResult
type IssueAuditResult struct {

	// Custom tag values that are set for the issue.
	// Required: true
	// Read Only: true
	CustomTagValues []*CustomTag `json:"customTagValues"`

	// Audited issue ID.
	// Required: true
	// Read Only: true
	ID int64 `json:"id"`

	// Issue instance ID.
	// Required: true
	// Read Only: true
	IssueInstanceID string `json:"issueInstanceId"`

	// ID of the application version the audited issue belongs to.
	// Required: true
	// Read Only: true
	ProjectVersionID int64 `json:"projectVersionId"`

	// Audited issue revision.
	// Required: true
	// Read Only: true
	Revision int32 `json:"revision"`

	// Is issue suppressed.
	// Required: true
	// Read Only: true
	Suppressed bool `json:"suppressed"`

	// User assigned to the audited issue.
	// Required: true
	// Read Only: true
	UserName string `json:"userName"`
}

// Validate validates this issue audit result
func (m *IssueAuditResult) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomTagValues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuppressed(formats); err != nil {
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

func (m *IssueAuditResult) validateCustomTagValues(formats strfmt.Registry) error {

	if err := validate.Required("customTagValues", "body", m.CustomTagValues); err != nil {
		return err
	}

	for i := 0; i < len(m.CustomTagValues); i++ {
		if swag.IsZero(m.CustomTagValues[i]) { // not required
			continue
		}

		if m.CustomTagValues[i] != nil {
			if err := m.CustomTagValues[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("customTagValues" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *IssueAuditResult) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditResult) validateIssueInstanceID(formats strfmt.Registry) error {

	if err := validate.RequiredString("issueInstanceId", "body", string(m.IssueInstanceID)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditResult) validateProjectVersionID(formats strfmt.Registry) error {

	if err := validate.Required("projectVersionId", "body", int64(m.ProjectVersionID)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditResult) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", int32(m.Revision)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditResult) validateSuppressed(formats strfmt.Registry) error {

	if err := validate.Required("suppressed", "body", bool(m.Suppressed)); err != nil {
		return err
	}

	return nil
}

func (m *IssueAuditResult) validateUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("userName", "body", string(m.UserName)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueAuditResult) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAuditResult) UnmarshalBinary(b []byte) error {
	var res IssueAuditResult
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
