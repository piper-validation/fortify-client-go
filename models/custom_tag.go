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

// CustomTag custom tag
// swagger:model Custom tag
type CustomTag struct {

	// Custom tag type.
	// Required: true
	// Enum: [UNKNOWN CUSTOM HYBRID METAGROUP SYSTEM AUDITASSISTANT]
	CustomTagType *string `json:"customTagType"`

	// Default value of the custom tag. Actual string value is presented here.
	// Read Only: true
	DefaultValue string `json:"defaultValue,omitempty"`

	// Index of default value of the custom tag. This is ordinal number of the item in CustomTagLookup collection.
	DefaultValueIndex int32 `json:"defaultValueIndex,omitempty"`

	// Flag that says if custom tag can be deleted. Custom tag which values are currently in use cannot be deleted.
	// Read Only: true
	Deletable *bool `json:"deletable,omitempty"`

	// Custom tag description.
	Description string `json:"description,omitempty"`

	// Flag that says if custom tag is extensible or not.
	// Required: true
	Extensible *bool `json:"extensible"`

	// Custom tag GUID.
	// Required: true
	// Read Only: true
	GUID string `json:"guid"`

	// Is custom tag hidden or not.
	// Required: true
	Hidden *bool `json:"hidden"`

	// Custom tag id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Is custom tag values are selected for any issues in the system.
	// Read Only: true
	InUse *bool `json:"inUse,omitempty"`

	// Custom tag unique name.
	// Required: true
	Name *string `json:"name"`

	// Custom tag version stored on the server. This value is used for optimistic locking of the custom tag object to prevent concurrent modification of the custom tag by different users at the same time.
	// Required: true
	// Read Only: true
	ObjectVersion int32 `json:"objectVersion"`

	// If this custom tag is set as primary tag for a specific application version. This value is initialized only if custom tags for specific application version are requested.
	// Read Only: true
	PrimaryTag *bool `json:"primaryTag,omitempty"`

	// Flag is set to true if special permission is required to set values of this custom tag.
	Restriction bool `json:"restriction,omitempty"`

	// Special permission type if restriction is set to TRUE.
	// Enum: [NONE RESTRICTED READONLY]
	RestrictionType string `json:"restrictionType,omitempty"`

	// Collection of all possible custom tag values.
	ValueList []*CustomTagLookup `json:"valueList"`

	// Custom tag value type.
	// Required: true
	// Enum: [LIST DECIMAL DATE TEXT]
	ValueType *string `json:"valueType"`
}

// Validate validates this custom tag
func (m *CustomTag) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCustomTagType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExtensible(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHidden(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRestrictionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var customTagTypeCustomTagTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNKNOWN","CUSTOM","HYBRID","METAGROUP","SYSTEM","AUDITASSISTANT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customTagTypeCustomTagTypePropEnum = append(customTagTypeCustomTagTypePropEnum, v)
	}
}

const (

	// CustomTagCustomTagTypeUNKNOWN captures enum value "UNKNOWN"
	CustomTagCustomTagTypeUNKNOWN string = "UNKNOWN"

	// CustomTagCustomTagTypeCUSTOM captures enum value "CUSTOM"
	CustomTagCustomTagTypeCUSTOM string = "CUSTOM"

	// CustomTagCustomTagTypeHYBRID captures enum value "HYBRID"
	CustomTagCustomTagTypeHYBRID string = "HYBRID"

	// CustomTagCustomTagTypeMETAGROUP captures enum value "METAGROUP"
	CustomTagCustomTagTypeMETAGROUP string = "METAGROUP"

	// CustomTagCustomTagTypeSYSTEM captures enum value "SYSTEM"
	CustomTagCustomTagTypeSYSTEM string = "SYSTEM"

	// CustomTagCustomTagTypeAUDITASSISTANT captures enum value "AUDITASSISTANT"
	CustomTagCustomTagTypeAUDITASSISTANT string = "AUDITASSISTANT"
)

// prop value enum
func (m *CustomTag) validateCustomTagTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customTagTypeCustomTagTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomTag) validateCustomTagType(formats strfmt.Registry) error {

	if err := validate.Required("customTagType", "body", m.CustomTagType); err != nil {
		return err
	}

	// value enum
	if err := m.validateCustomTagTypeEnum("customTagType", "body", *m.CustomTagType); err != nil {
		return err
	}

	return nil
}

func (m *CustomTag) validateExtensible(formats strfmt.Registry) error {

	if err := validate.Required("extensible", "body", m.Extensible); err != nil {
		return err
	}

	return nil
}

func (m *CustomTag) validateGUID(formats strfmt.Registry) error {

	if err := validate.RequiredString("guid", "body", string(m.GUID)); err != nil {
		return err
	}

	return nil
}

func (m *CustomTag) validateHidden(formats strfmt.Registry) error {

	if err := validate.Required("hidden", "body", m.Hidden); err != nil {
		return err
	}

	return nil
}

func (m *CustomTag) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *CustomTag) validateObjectVersion(formats strfmt.Registry) error {

	if err := validate.Required("objectVersion", "body", int32(m.ObjectVersion)); err != nil {
		return err
	}

	return nil
}

var customTagTypeRestrictionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","RESTRICTED","READONLY"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customTagTypeRestrictionTypePropEnum = append(customTagTypeRestrictionTypePropEnum, v)
	}
}

const (

	// CustomTagRestrictionTypeNONE captures enum value "NONE"
	CustomTagRestrictionTypeNONE string = "NONE"

	// CustomTagRestrictionTypeRESTRICTED captures enum value "RESTRICTED"
	CustomTagRestrictionTypeRESTRICTED string = "RESTRICTED"

	// CustomTagRestrictionTypeREADONLY captures enum value "READONLY"
	CustomTagRestrictionTypeREADONLY string = "READONLY"
)

// prop value enum
func (m *CustomTag) validateRestrictionTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customTagTypeRestrictionTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomTag) validateRestrictionType(formats strfmt.Registry) error {

	if swag.IsZero(m.RestrictionType) { // not required
		return nil
	}

	// value enum
	if err := m.validateRestrictionTypeEnum("restrictionType", "body", m.RestrictionType); err != nil {
		return err
	}

	return nil
}

func (m *CustomTag) validateValueList(formats strfmt.Registry) error {

	if swag.IsZero(m.ValueList) { // not required
		return nil
	}

	for i := 0; i < len(m.ValueList); i++ {
		if swag.IsZero(m.ValueList[i]) { // not required
			continue
		}

		if m.ValueList[i] != nil {
			if err := m.ValueList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("valueList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var customTagTypeValueTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["LIST","DECIMAL","DATE","TEXT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		customTagTypeValueTypePropEnum = append(customTagTypeValueTypePropEnum, v)
	}
}

const (

	// CustomTagValueTypeLIST captures enum value "LIST"
	CustomTagValueTypeLIST string = "LIST"

	// CustomTagValueTypeDECIMAL captures enum value "DECIMAL"
	CustomTagValueTypeDECIMAL string = "DECIMAL"

	// CustomTagValueTypeDATE captures enum value "DATE"
	CustomTagValueTypeDATE string = "DATE"

	// CustomTagValueTypeTEXT captures enum value "TEXT"
	CustomTagValueTypeTEXT string = "TEXT"
)

// prop value enum
func (m *CustomTag) validateValueTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, customTagTypeValueTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *CustomTag) validateValueType(formats strfmt.Registry) error {

	if err := validate.Required("valueType", "body", m.ValueType); err != nil {
		return err
	}

	// value enum
	if err := m.validateValueTypeEnum("valueType", "body", *m.ValueType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *CustomTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomTag) UnmarshalBinary(b []byte) error {
	var res CustomTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
