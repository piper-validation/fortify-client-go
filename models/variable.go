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

// Variable Variables count issues, conditions, and other categories of numeric data.
// swagger:model Variable
type Variable struct {

	// attribute for operation
	AttributeForOperation string `json:"attributeForOperation,omitempty"`

	// description
	Description string `json:"description,omitempty"`

	// Select a folder from the default filter set to associate with the variable.
	// Required: true
	FolderName *string `json:"folderName"`

	// guid
	GUID string `json:"guid,omitempty"`

	// id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// in use
	InUse bool `json:"inUse,omitempty"`

	// name
	// Required: true
	Name *string `json:"name"`

	// object version
	ObjectVersion int32 `json:"objectVersion,omitempty"`

	// operation
	// Enum: [COUNT SUM]
	Operation string `json:"operation,omitempty"`

	// publish version
	PublishVersion int32 `json:"publishVersion,omitempty"`

	// search string
	// Required: true
	SearchString *string `json:"searchString"`

	// variable type
	// Enum: [SYSTEM_DEFINED USER_DEFINED]
	VariableType string `json:"variableType,omitempty"`
}

// Validate validates this variable
func (m *Variable) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateFolderName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOperation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSearchString(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVariableType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Variable) validateFolderName(formats strfmt.Registry) error {

	if err := validate.Required("folderName", "body", m.FolderName); err != nil {
		return err
	}

	return nil
}

func (m *Variable) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

var variableTypeOperationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["COUNT","SUM"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		variableTypeOperationPropEnum = append(variableTypeOperationPropEnum, v)
	}
}

const (

	// VariableOperationCOUNT captures enum value "COUNT"
	VariableOperationCOUNT string = "COUNT"

	// VariableOperationSUM captures enum value "SUM"
	VariableOperationSUM string = "SUM"
)

// prop value enum
func (m *Variable) validateOperationEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, variableTypeOperationPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Variable) validateOperation(formats strfmt.Registry) error {

	if swag.IsZero(m.Operation) { // not required
		return nil
	}

	// value enum
	if err := m.validateOperationEnum("operation", "body", m.Operation); err != nil {
		return err
	}

	return nil
}

func (m *Variable) validateSearchString(formats strfmt.Registry) error {

	if err := validate.Required("searchString", "body", m.SearchString); err != nil {
		return err
	}

	return nil
}

var variableTypeVariableTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SYSTEM_DEFINED","USER_DEFINED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		variableTypeVariableTypePropEnum = append(variableTypeVariableTypePropEnum, v)
	}
}

const (

	// VariableVariableTypeSYSTEMDEFINED captures enum value "SYSTEM_DEFINED"
	VariableVariableTypeSYSTEMDEFINED string = "SYSTEM_DEFINED"

	// VariableVariableTypeUSERDEFINED captures enum value "USER_DEFINED"
	VariableVariableTypeUSERDEFINED string = "USER_DEFINED"
)

// prop value enum
func (m *Variable) validateVariableTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, variableTypeVariableTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *Variable) validateVariableType(formats strfmt.Registry) error {

	if swag.IsZero(m.VariableType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVariableTypeEnum("variableType", "body", m.VariableType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Variable) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Variable) UnmarshalBinary(b []byte) error {
	var res Variable
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
