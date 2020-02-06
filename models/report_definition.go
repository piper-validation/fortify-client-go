// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ReportDefinition Report Definition DTO object
// swagger:model Report Definition
type ReportDefinition struct {

	// True if report applies to multiple application versions
	CrossApp bool `json:"crossApp,omitempty"`

	// Report description
	Description string `json:"description,omitempty"`

	// The name of the report definition file
	FileName string `json:"fileName,omitempty"`

	// Report definition GUID
	GUID string `json:"guid,omitempty"`

	// Report definition identifier
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Indicates whether the report definition is in use
	InUse bool `json:"inUse,omitempty"`

	// Report name
	// Required: true
	Name *string `json:"name"`

	// Object version
	ObjectVersion int32 `json:"objectVersion,omitempty"`

	// List of report parameters
	Parameters []*ReportParameter `json:"parameters"`

	// Publish version
	PublishVersion int32 `json:"publishVersion,omitempty"`

	// The engine used to render the report, e.g. BIRT
	RenderingEngine string `json:"renderingEngine,omitempty"`

	// Template doc identifier
	TemplateDocID int64 `json:"templateDocId,omitempty"`

	// Type of report
	// Required: true
	// Enum: [PROJECT SSA_PROJECT SSA_PORTFOLIO PORTFOLIO COMPLIANCE ISSUE RUNTIME_REPORTS]
	Type *string `json:"type"`

	// Report type default text
	TypeDefaultText string `json:"typeDefaultText,omitempty"`
}

// Validate validates this report definition
func (m *ReportDefinition) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParameters(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ReportDefinition) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ReportDefinition) validateParameters(formats strfmt.Registry) error {

	if swag.IsZero(m.Parameters) { // not required
		return nil
	}

	for i := 0; i < len(m.Parameters); i++ {
		if swag.IsZero(m.Parameters[i]) { // not required
			continue
		}

		if m.Parameters[i] != nil {
			if err := m.Parameters[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("parameters" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var reportDefinitionTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["PROJECT","SSA_PROJECT","SSA_PORTFOLIO","PORTFOLIO","COMPLIANCE","ISSUE","RUNTIME_REPORTS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportDefinitionTypeTypePropEnum = append(reportDefinitionTypeTypePropEnum, v)
	}
}

const (

	// ReportDefinitionTypePROJECT captures enum value "PROJECT"
	ReportDefinitionTypePROJECT string = "PROJECT"

	// ReportDefinitionTypeSSAPROJECT captures enum value "SSA_PROJECT"
	ReportDefinitionTypeSSAPROJECT string = "SSA_PROJECT"

	// ReportDefinitionTypeSSAPORTFOLIO captures enum value "SSA_PORTFOLIO"
	ReportDefinitionTypeSSAPORTFOLIO string = "SSA_PORTFOLIO"

	// ReportDefinitionTypePORTFOLIO captures enum value "PORTFOLIO"
	ReportDefinitionTypePORTFOLIO string = "PORTFOLIO"

	// ReportDefinitionTypeCOMPLIANCE captures enum value "COMPLIANCE"
	ReportDefinitionTypeCOMPLIANCE string = "COMPLIANCE"

	// ReportDefinitionTypeISSUE captures enum value "ISSUE"
	ReportDefinitionTypeISSUE string = "ISSUE"

	// ReportDefinitionTypeRUNTIMEREPORTS captures enum value "RUNTIME_REPORTS"
	ReportDefinitionTypeRUNTIMEREPORTS string = "RUNTIME_REPORTS"
)

// prop value enum
func (m *ReportDefinition) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, reportDefinitionTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReportDefinition) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ReportDefinition) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportDefinition) UnmarshalBinary(b []byte) error {
	var res ReportDefinition
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}