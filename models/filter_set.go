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

// FilterSet basic info on a filter set (does not include filters)
// swagger:model Filter Set
type FilterSet struct {

	// whether this filter set is the default within its issue template
	// Required: true
	// Read Only: true
	DefaultFilterSet bool `json:"defaultFilterSet"`

	// description
	// Required: true
	// Read Only: true
	Description string `json:"description"`

	// List of folders defined in filter set
	// Required: true
	// Read Only: true
	Folders []*FolderDto `json:"folders"`

	// GUID of filter set
	// Required: true
	// Read Only: true
	GUID string `json:"guid"`

	// name of filter set
	// Required: true
	// Read Only: true
	Title string `json:"title"`
}

// Validate validates this filter set
func (m *FilterSet) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDefaultFilterSet(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolders(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTitle(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FilterSet) validateDefaultFilterSet(formats strfmt.Registry) error {

	if err := validate.Required("defaultFilterSet", "body", bool(m.DefaultFilterSet)); err != nil {
		return err
	}

	return nil
}

func (m *FilterSet) validateDescription(formats strfmt.Registry) error {

	if err := validate.RequiredString("description", "body", string(m.Description)); err != nil {
		return err
	}

	return nil
}

func (m *FilterSet) validateFolders(formats strfmt.Registry) error {

	if err := validate.Required("folders", "body", m.Folders); err != nil {
		return err
	}

	for i := 0; i < len(m.Folders); i++ {
		if swag.IsZero(m.Folders[i]) { // not required
			continue
		}

		if m.Folders[i] != nil {
			if err := m.Folders[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("folders" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FilterSet) validateGUID(formats strfmt.Registry) error {

	if err := validate.RequiredString("guid", "body", string(m.GUID)); err != nil {
		return err
	}

	return nil
}

func (m *FilterSet) validateTitle(formats strfmt.Registry) error {

	if err := validate.RequiredString("title", "body", string(m.Title)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FilterSet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FilterSet) UnmarshalBinary(b []byte) error {
	var res FilterSet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
