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

// LDAPEntity An LDAP entity like user, group or organization
// swagger:model LDAP entity
type LDAPEntity struct {

	// Distinguished name of LDAP entity
	// Required: true
	DistinguishedName *string `json:"distinguishedName"`

	// Email of LDAP entity
	// Required: true
	Email *string `json:"email"`

	// First name of LDAP entity
	// Required: true
	FirstName *string `json:"firstName"`

	// LDAP entity identifier
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last name of LDAP entity
	// Required: true
	LastName *string `json:"lastName"`

	// Type of LDAP entity.
	// Required: true
	// Enum: [NONE GROUP USER ORG_UNIT]
	LdapType *string `json:"ldapType"`

	// LDAP entity name
	// Required: true
	Name *string `json:"name"`

	// List of roles pertaining to this LDAP entity
	// Required: true
	Roles []*Role `json:"roles"`

	// Photo object if LDAP entity is user
	UserPhoto *UserPhoto `json:"userPhoto,omitempty"`
}

// Validate validates this l d a p entity
func (m *LDAPEntity) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDistinguishedName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEmail(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFirstName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLdapType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoles(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUserPhoto(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *LDAPEntity) validateDistinguishedName(formats strfmt.Registry) error {

	if err := validate.Required("distinguishedName", "body", m.DistinguishedName); err != nil {
		return err
	}

	return nil
}

func (m *LDAPEntity) validateEmail(formats strfmt.Registry) error {

	if err := validate.Required("email", "body", m.Email); err != nil {
		return err
	}

	return nil
}

func (m *LDAPEntity) validateFirstName(formats strfmt.Registry) error {

	if err := validate.Required("firstName", "body", m.FirstName); err != nil {
		return err
	}

	return nil
}

func (m *LDAPEntity) validateLastName(formats strfmt.Registry) error {

	if err := validate.Required("lastName", "body", m.LastName); err != nil {
		return err
	}

	return nil
}

var lDAPEntityTypeLdapTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","GROUP","USER","ORG_UNIT"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		lDAPEntityTypeLdapTypePropEnum = append(lDAPEntityTypeLdapTypePropEnum, v)
	}
}

const (

	// LDAPEntityLdapTypeNONE captures enum value "NONE"
	LDAPEntityLdapTypeNONE string = "NONE"

	// LDAPEntityLdapTypeGROUP captures enum value "GROUP"
	LDAPEntityLdapTypeGROUP string = "GROUP"

	// LDAPEntityLdapTypeUSER captures enum value "USER"
	LDAPEntityLdapTypeUSER string = "USER"

	// LDAPEntityLdapTypeORGUNIT captures enum value "ORG_UNIT"
	LDAPEntityLdapTypeORGUNIT string = "ORG_UNIT"
)

// prop value enum
func (m *LDAPEntity) validateLdapTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, lDAPEntityTypeLdapTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *LDAPEntity) validateLdapType(formats strfmt.Registry) error {

	if err := validate.Required("ldapType", "body", m.LdapType); err != nil {
		return err
	}

	// value enum
	if err := m.validateLdapTypeEnum("ldapType", "body", *m.LdapType); err != nil {
		return err
	}

	return nil
}

func (m *LDAPEntity) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *LDAPEntity) validateRoles(formats strfmt.Registry) error {

	if err := validate.Required("roles", "body", m.Roles); err != nil {
		return err
	}

	for i := 0; i < len(m.Roles); i++ {
		if swag.IsZero(m.Roles[i]) { // not required
			continue
		}

		if m.Roles[i] != nil {
			if err := m.Roles[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("roles" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *LDAPEntity) validateUserPhoto(formats strfmt.Registry) error {

	if swag.IsZero(m.UserPhoto) { // not required
		return nil
	}

	if m.UserPhoto != nil {
		if err := m.UserPhoto.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("userPhoto")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *LDAPEntity) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *LDAPEntity) UnmarshalBinary(b []byte) error {
	var res LDAPEntity
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
