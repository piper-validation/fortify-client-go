// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IssuePrimaryTag Store the information about issue primary custom tag
// swagger:model Issue Primary Tag
type IssuePrimaryTag struct {

	// tag Guid
	TagGUID string `json:"tagGuid,omitempty"`

	// tag Id
	TagID int64 `json:"tagId,omitempty"`

	// tag name
	TagName string `json:"tagName,omitempty"`

	// tag value
	TagValue string `json:"tagValue,omitempty"`
}

// Validate validates this issue primary tag
func (m *IssuePrimaryTag) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IssuePrimaryTag) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssuePrimaryTag) UnmarshalBinary(b []byte) error {
	var res IssuePrimaryTag
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
