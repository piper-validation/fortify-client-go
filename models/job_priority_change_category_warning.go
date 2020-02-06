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

// JobPriorityChangeCategoryWarning Job priority change warning
// swagger:model JobPriorityChangeCategoryWarning
type JobPriorityChangeCategoryWarning struct {

	// Job priority change warning category.
	// Required: true
	// Read Only: true
	WarningCategory string `json:"warningCategory"`

	// Detailed information about job priority change.
	// Required: true
	// Read Only: true
	WarningList []*JobPriorityChangeInfo `json:"warningList"`
}

// Validate validates this job priority change category warning
func (m *JobPriorityChangeCategoryWarning) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateWarningCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWarningList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *JobPriorityChangeCategoryWarning) validateWarningCategory(formats strfmt.Registry) error {

	if err := validate.RequiredString("warningCategory", "body", string(m.WarningCategory)); err != nil {
		return err
	}

	return nil
}

func (m *JobPriorityChangeCategoryWarning) validateWarningList(formats strfmt.Registry) error {

	if err := validate.Required("warningList", "body", m.WarningList); err != nil {
		return err
	}

	for i := 0; i < len(m.WarningList); i++ {
		if swag.IsZero(m.WarningList[i]) { // not required
			continue
		}

		if m.WarningList[i] != nil {
			if err := m.WarningList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("warningList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *JobPriorityChangeCategoryWarning) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *JobPriorityChangeCategoryWarning) UnmarshalBinary(b []byte) error {
	var res JobPriorityChangeCategoryWarning
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}