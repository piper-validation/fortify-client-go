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

// ProjectVersionTrainAuditAssistantRequest Request to send the application version's data to Audit Assistant for training
// swagger:model ProjectVersionTrainAuditAssistantRequest
type ProjectVersionTrainAuditAssistantRequest struct {

	// List containing single application version ID to send training for audit assistant
	// Required: true
	ProjectVersionIds []int64 `json:"projectVersionIds"`
}

// Validate validates this project version train audit assistant request
func (m *ProjectVersionTrainAuditAssistantRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateProjectVersionIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVersionTrainAuditAssistantRequest) validateProjectVersionIds(formats strfmt.Registry) error {

	if err := validate.Required("projectVersionIds", "body", m.ProjectVersionIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectVersionTrainAuditAssistantRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVersionTrainAuditAssistantRequest) UnmarshalBinary(b []byte) error {
	var res ProjectVersionTrainAuditAssistantRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
