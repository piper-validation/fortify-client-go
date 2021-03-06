// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Folder folder
// swagger:model Folder
type Folder struct {

	// Color that should be used to represent folder on UI
	// Read Only: true
	Color string `json:"color,omitempty"`

	// Folder description
	// Read Only: true
	Description string `json:"description,omitempty"`

	// Is folder editable
	// Read Only: true
	Editable *bool `json:"editable,omitempty"`

	// Folder GUID. Guid is unique across all the folders defined for a application version
	// Read Only: true
	GUID string `json:"guid,omitempty"`

	// Folder id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Folder name
	// Read Only: true
	Name string `json:"name,omitempty"`

	// Numeric value used for ordering folder from the most important to the least important
	// Read Only: true
	OrderIndex int32 `json:"orderIndex,omitempty"`

	// ID of the application version for which folder is defined
	// Read Only: true
	ProjectVersionID int64 `json:"projectVersionId,omitempty"`
}

// Validate validates this folder
func (m *Folder) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Folder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Folder) UnmarshalBinary(b []byte) error {
	var res Folder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
