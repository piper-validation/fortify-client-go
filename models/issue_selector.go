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

// IssueSelector Issue selector DTO object
// swagger:model Issue selector
type IssueSelector struct {

	// Description
	// Required: true
	Description *string `json:"description"`

	// Display name for issue selector
	// Required: true
	DisplayName *string `json:"displayName"`

	// Issue selector entity type
	// Required: true
	// Enum: [UNDEFINED ISSUE EXTERNALLIST CUSTOMTAG HYBRIDTAG FOLDER]
	EntityType *string `json:"entityType"`

	// Issue selector global unique identifier
	// Required: true
	GUID *string `json:"guid"`

	// Issue selector value
	// Required: true
	Value *string `json:"value"`
}

// Validate validates this issue selector
func (m *IssueSelector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEntityType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValue(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *IssueSelector) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *IssueSelector) validateDisplayName(formats strfmt.Registry) error {

	if err := validate.Required("displayName", "body", m.DisplayName); err != nil {
		return err
	}

	return nil
}

var issueSelectorTypeEntityTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNDEFINED","ISSUE","EXTERNALLIST","CUSTOMTAG","HYBRIDTAG","FOLDER"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		issueSelectorTypeEntityTypePropEnum = append(issueSelectorTypeEntityTypePropEnum, v)
	}
}

const (

	// IssueSelectorEntityTypeUNDEFINED captures enum value "UNDEFINED"
	IssueSelectorEntityTypeUNDEFINED string = "UNDEFINED"

	// IssueSelectorEntityTypeISSUE captures enum value "ISSUE"
	IssueSelectorEntityTypeISSUE string = "ISSUE"

	// IssueSelectorEntityTypeEXTERNALLIST captures enum value "EXTERNALLIST"
	IssueSelectorEntityTypeEXTERNALLIST string = "EXTERNALLIST"

	// IssueSelectorEntityTypeCUSTOMTAG captures enum value "CUSTOMTAG"
	IssueSelectorEntityTypeCUSTOMTAG string = "CUSTOMTAG"

	// IssueSelectorEntityTypeHYBRIDTAG captures enum value "HYBRIDTAG"
	IssueSelectorEntityTypeHYBRIDTAG string = "HYBRIDTAG"

	// IssueSelectorEntityTypeFOLDER captures enum value "FOLDER"
	IssueSelectorEntityTypeFOLDER string = "FOLDER"
)

// prop value enum
func (m *IssueSelector) validateEntityTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, issueSelectorTypeEntityTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *IssueSelector) validateEntityType(formats strfmt.Registry) error {

	if err := validate.Required("entityType", "body", m.EntityType); err != nil {
		return err
	}

	// value enum
	if err := m.validateEntityTypeEnum("entityType", "body", *m.EntityType); err != nil {
		return err
	}

	return nil
}

func (m *IssueSelector) validateGUID(formats strfmt.Registry) error {

	if err := validate.Required("guid", "body", m.GUID); err != nil {
		return err
	}

	return nil
}

func (m *IssueSelector) validateValue(formats strfmt.Registry) error {

	if err := validate.Required("value", "body", m.Value); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *IssueSelector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueSelector) UnmarshalBinary(b []byte) error {
	var res IssueSelector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}