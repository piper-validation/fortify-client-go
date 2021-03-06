// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// BugParam bug param
// swagger:model BugParam
type BugParam struct {

	// bug param type
	// Enum: [BUGPARAM_TEXT BUGPARAM_TEXTAREA BUGPARAM_CHOICE]
	BugParamType string `json:"bugParamType,omitempty"`

	// choice list
	ChoiceList []string `json:"choiceList"`

	// description
	Description string `json:"description,omitempty"`

	// display label
	DisplayLabel string `json:"displayLabel,omitempty"`

	// has dependent params
	HasDependentParams bool `json:"hasDependentParams,omitempty"`

	// identifier
	Identifier string `json:"identifier,omitempty"`

	// max length
	MaxLength int32 `json:"maxLength,omitempty"`

	// required
	Required bool `json:"required,omitempty"`

	// value
	Value string `json:"value,omitempty"`
}

// Validate validates this bug param
func (m *BugParam) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBugParamType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var bugParamTypeBugParamTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["BUGPARAM_TEXT","BUGPARAM_TEXTAREA","BUGPARAM_CHOICE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		bugParamTypeBugParamTypePropEnum = append(bugParamTypeBugParamTypePropEnum, v)
	}
}

const (

	// BugParamBugParamTypeBUGPARAMTEXT captures enum value "BUGPARAM_TEXT"
	BugParamBugParamTypeBUGPARAMTEXT string = "BUGPARAM_TEXT"

	// BugParamBugParamTypeBUGPARAMTEXTAREA captures enum value "BUGPARAM_TEXTAREA"
	BugParamBugParamTypeBUGPARAMTEXTAREA string = "BUGPARAM_TEXTAREA"

	// BugParamBugParamTypeBUGPARAMCHOICE captures enum value "BUGPARAM_CHOICE"
	BugParamBugParamTypeBUGPARAMCHOICE string = "BUGPARAM_CHOICE"
)

// prop value enum
func (m *BugParam) validateBugParamTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, bugParamTypeBugParamTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *BugParam) validateBugParamType(formats strfmt.Registry) error {

	if swag.IsZero(m.BugParamType) { // not required
		return nil
	}

	// value enum
	if err := m.validateBugParamTypeEnum("bugParamType", "body", m.BugParamType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *BugParam) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BugParam) UnmarshalBinary(b []byte) error {
	var res BugParam
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
