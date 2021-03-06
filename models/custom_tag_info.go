// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// CustomTagInfo Holds the details about existed custom tag that is assigned or linked with some other entity
// swagger:model customTagInfo
type CustomTagInfo struct {

	// Custom tag GUID
	// Read Only: true
	GUID string `json:"guid,omitempty"`

	// Custom tag id
	// Required: true
	ID *int64 `json:"id"`

	// Custom tag unique name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Collection of custom tag values
	ValueList []*CustomTagLookupInfo `json:"valueList"`
}

// Validate validates this custom tag info
func (m *CustomTagInfo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateValueList(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CustomTagInfo) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *CustomTagInfo) validateValueList(formats strfmt.Registry) error {

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

// MarshalBinary interface implementation
func (m *CustomTagInfo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CustomTagInfo) UnmarshalBinary(b []byte) error {
	var res CustomTagInfo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
