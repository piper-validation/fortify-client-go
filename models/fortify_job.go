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

// FortifyJob Fortify job DTO object, carries less information than the main Job DTO. Used primarily for the snapshot refresh which be initiated by a non-admin user
// swagger:model Fortify job
type FortifyJob struct {

	// id
	// Required: true
	// Read Only: true
	ID string `json:"id"`

	// SSC username which initiated the job
	// Required: true
	// Read Only: true
	InvokingUserName string `json:"invokingUserName"`

	// job state
	// Required: true
	// Read Only: true
	JobState int32 `json:"jobState"`

	// job type
	// Required: true
	// Read Only: true
	JobType int32 `json:"jobType"`
}

// Validate validates this fortify job
func (m *FortifyJob) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvokingUserName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateJobType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FortifyJob) validateID(formats strfmt.Registry) error {

	if err := validate.RequiredString("id", "body", string(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *FortifyJob) validateInvokingUserName(formats strfmt.Registry) error {

	if err := validate.RequiredString("invokingUserName", "body", string(m.InvokingUserName)); err != nil {
		return err
	}

	return nil
}

func (m *FortifyJob) validateJobState(formats strfmt.Registry) error {

	if err := validate.Required("jobState", "body", int32(m.JobState)); err != nil {
		return err
	}

	return nil
}

func (m *FortifyJob) validateJobType(formats strfmt.Registry) error {

	if err := validate.Required("jobType", "body", int32(m.JobType)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FortifyJob) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FortifyJob) UnmarshalBinary(b []byte) error {
	var res FortifyJob
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
