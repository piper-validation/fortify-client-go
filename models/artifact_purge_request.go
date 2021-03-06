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

// ArtifactPurgeRequest Request to purge artifacts
// swagger:model ArtifactPurgeRequest
type ArtifactPurgeRequest struct {

	// List containing single artifact ID to purge
	// Required: true
	ArtifactIds []int64 `json:"artifactIds"`
}

// Validate validates this artifact purge request
func (m *ArtifactPurgeRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateArtifactIds(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ArtifactPurgeRequest) validateArtifactIds(formats strfmt.Registry) error {

	if err := validate.Required("artifactIds", "body", m.ArtifactIds); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ArtifactPurgeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ArtifactPurgeRequest) UnmarshalBinary(b []byte) error {
	var res ArtifactPurgeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
