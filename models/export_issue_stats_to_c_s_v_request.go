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

// ExportIssueStatsToCSVRequest Request to export issue stats to CSV file
// swagger:model ExportIssueStatsToCSVRequest
type ExportIssueStatsToCSVRequest struct {

	// Dataset name (Issue Stat)
	// Required: true
	DatasetName *string `json:"datasetName"`

	// Csv file name
	// Required: true
	FileName *string `json:"fileName"`

	// Limit
	Limit int32 `json:"limit,omitempty"`

	// Note for csv export
	Note string `json:"note,omitempty"`

	// Start
	Start int32 `json:"start,omitempty"`
}

// Validate validates this export issue stats to c s v request
func (m *ExportIssueStatsToCSVRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDatasetName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFileName(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ExportIssueStatsToCSVRequest) validateDatasetName(formats strfmt.Registry) error {

	if err := validate.Required("datasetName", "body", m.DatasetName); err != nil {
		return err
	}

	return nil
}

func (m *ExportIssueStatsToCSVRequest) validateFileName(formats strfmt.Registry) error {

	if err := validate.Required("fileName", "body", m.FileName); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ExportIssueStatsToCSVRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ExportIssueStatsToCSVRequest) UnmarshalBinary(b []byte) error {
	var res ExportIssueStatsToCSVRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
