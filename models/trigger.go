// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Trigger Trigger in a request or response that identifies a vulnerability
// swagger:model Trigger
type Trigger struct {

	// length
	Length int32 `json:"length,omitempty"`

	// location
	Location int32 `json:"location,omitempty"`

	// trigger string
	TriggerString string `json:"triggerString,omitempty"`
}

// Validate validates this trigger
func (m *Trigger) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Trigger) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Trigger) UnmarshalBinary(b []byte) error {
	var res Trigger
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}