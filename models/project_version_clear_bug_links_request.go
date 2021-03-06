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

// ProjectVersionClearBugLinksRequest Request to clear the specified bug links from the application version
// swagger:model ProjectVersionClearBugLinksRequest
type ProjectVersionClearBugLinksRequest struct {

	// Identifying information for the bug links to be cleared. (The format of an externalBugId depends on the SSC bug tracker plugin which filed the bug)
	// Required: true
	ExternalBugIds []string `json:"externalBugIds"`
}

// Validate validates this project version clear bug links request
func (m *ProjectVersionClearBugLinksRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateExternalBugIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVersionClearBugLinksRequest) validateExternalBugIds(formats strfmt.Registry) error {

	if err := validate.Required("externalBugIds", "body", m.ExternalBugIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectVersionClearBugLinksRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVersionClearBugLinksRequest) UnmarshalBinary(b []byte) error {
	var res ProjectVersionClearBugLinksRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
