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

// ReportParameter Report Parameter DTO object
// swagger:model Report Parameter
type ReportParameter struct {

	// description
	Description string `json:"description,omitempty"`

	// id
	ID int64 `json:"id,omitempty"`

	// Report parameter identifier
	// Required: true
	Identifier *string `json:"identifier"`

	// name
	Name string `json:"name,omitempty"`

	// Order in which the parameter should appear
	ParamOrder int32 `json:"paramOrder,omitempty"`

	// Report definition identifier
	// Required: true
	ReportDefinitionID *int64 `json:"reportDefinitionId"`

	// Report parameter options
	ReportParameterOptions []*ReportParameterOption `json:"reportParameterOptions"`

	// Report parameter type
	// Required: true
	// Enum: [SINGLE_PROJECT SINGLE_RUNTIME_APP SINGLE_SSA_PROJECT MULTI_PROJECT MULTI_RUNTIME_APP MULTI_SSA_PROJECT PROJECT_ATTRIBUTE STRING BOOLEAN DATE SINGLE_SELECT_DEFAULT METADEF_GUID]
	Type *string `json:"type"`
}

// Validate validates this report parameter
func (m *ReportParameter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportDefinitionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReportParameterOptions(formats); err != nil {
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

func (m *ReportParameter) validateIdentifier(formats strfmt.Registry) error {

	if err := validate.Required("identifier", "body", m.Identifier); err != nil {
		return err
	}

	return nil
}

func (m *ReportParameter) validateReportDefinitionID(formats strfmt.Registry) error {

	if err := validate.Required("reportDefinitionId", "body", m.ReportDefinitionID); err != nil {
		return err
	}

	return nil
}

func (m *ReportParameter) validateReportParameterOptions(formats strfmt.Registry) error {

	if swag.IsZero(m.ReportParameterOptions) { // not required
		return nil
	}

	for i := 0; i < len(m.ReportParameterOptions); i++ {
		if swag.IsZero(m.ReportParameterOptions[i]) { // not required
			continue
		}

		if m.ReportParameterOptions[i] != nil {
			if err := m.ReportParameterOptions[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("reportParameterOptions" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var reportParameterTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SINGLE_PROJECT","SINGLE_RUNTIME_APP","SINGLE_SSA_PROJECT","MULTI_PROJECT","MULTI_RUNTIME_APP","MULTI_SSA_PROJECT","PROJECT_ATTRIBUTE","STRING","BOOLEAN","DATE","SINGLE_SELECT_DEFAULT","METADEF_GUID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		reportParameterTypeTypePropEnum = append(reportParameterTypeTypePropEnum, v)
	}
}

const (

	// ReportParameterTypeSINGLEPROJECT captures enum value "SINGLE_PROJECT"
	ReportParameterTypeSINGLEPROJECT string = "SINGLE_PROJECT"

	// ReportParameterTypeSINGLERUNTIMEAPP captures enum value "SINGLE_RUNTIME_APP"
	ReportParameterTypeSINGLERUNTIMEAPP string = "SINGLE_RUNTIME_APP"

	// ReportParameterTypeSINGLESSAPROJECT captures enum value "SINGLE_SSA_PROJECT"
	ReportParameterTypeSINGLESSAPROJECT string = "SINGLE_SSA_PROJECT"

	// ReportParameterTypeMULTIPROJECT captures enum value "MULTI_PROJECT"
	ReportParameterTypeMULTIPROJECT string = "MULTI_PROJECT"

	// ReportParameterTypeMULTIRUNTIMEAPP captures enum value "MULTI_RUNTIME_APP"
	ReportParameterTypeMULTIRUNTIMEAPP string = "MULTI_RUNTIME_APP"

	// ReportParameterTypeMULTISSAPROJECT captures enum value "MULTI_SSA_PROJECT"
	ReportParameterTypeMULTISSAPROJECT string = "MULTI_SSA_PROJECT"

	// ReportParameterTypePROJECTATTRIBUTE captures enum value "PROJECT_ATTRIBUTE"
	ReportParameterTypePROJECTATTRIBUTE string = "PROJECT_ATTRIBUTE"

	// ReportParameterTypeSTRING captures enum value "STRING"
	ReportParameterTypeSTRING string = "STRING"

	// ReportParameterTypeBOOLEAN captures enum value "BOOLEAN"
	ReportParameterTypeBOOLEAN string = "BOOLEAN"

	// ReportParameterTypeDATE captures enum value "DATE"
	ReportParameterTypeDATE string = "DATE"

	// ReportParameterTypeSINGLESELECTDEFAULT captures enum value "SINGLE_SELECT_DEFAULT"
	ReportParameterTypeSINGLESELECTDEFAULT string = "SINGLE_SELECT_DEFAULT"

	// ReportParameterTypeMETADEFGUID captures enum value "METADEF_GUID"
	ReportParameterTypeMETADEFGUID string = "METADEF_GUID"
)

// prop value enum
func (m *ReportParameter) validateTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, reportParameterTypeTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ReportParameter) validateType(formats strfmt.Registry) error {

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
func (m *ReportParameter) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ReportParameter) UnmarshalBinary(b []byte) error {
	var res ReportParameter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}