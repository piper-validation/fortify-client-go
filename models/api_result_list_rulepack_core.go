// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// APIResultListRulepackCore Api result list rulepack core
// swagger:model ApiResult«List«Rulepack Core»»
type APIResultListRulepackCore struct {

	// count
	Count int64 `json:"count,omitempty"`

	// data
	Data []*RulepackCore `json:"data"`

	// error code
	ErrorCode int32 `json:"errorCode,omitempty"`

	// links
	Links interface{} `json:"links,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// response code
	ResponseCode int32 `json:"responseCode,omitempty"`

	// stack trace
	StackTrace string `json:"stackTrace,omitempty"`

	// success count
	SuccessCount int64 `json:"successCount,omitempty"`
}

// Validate validates this Api result list rulepack core
func (m *APIResultListRulepackCore) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *APIResultListRulepackCore) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	for i := 0; i < len(m.Data); i++ {
		if swag.IsZero(m.Data[i]) { // not required
			continue
		}

		if m.Data[i] != nil {
			if err := m.Data[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("data" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *APIResultListRulepackCore) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *APIResultListRulepackCore) UnmarshalBinary(b []byte) error {
	var res APIResultListRulepackCore
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
