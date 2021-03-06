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

// AuditAssistantStatus AuditAssistantStatus DTO object
// swagger:model AuditAssistantStatus
type AuditAssistantStatus struct {

	// fpr file path
	FprFilePath string `json:"fprFilePath,omitempty"`

	// message
	Message string `json:"message,omitempty"`

	// project version Id
	ProjectVersionID int64 `json:"projectVersionId,omitempty"`

	// server Id
	ServerID int64 `json:"serverId,omitempty"`

	// server status
	ServerStatus int32 `json:"serverStatus,omitempty"`

	// status
	// Enum: [NONE PRE_PREDICT_FAILURE PREDICT_QUEUED PREDICT_PREP PREDICT_SUBMITTED PREDICT_COMPLETE PREDICT_FAILED PREDICTIONS_RETRIEVING PREDICTIONS_RETRIEVED PREDICTIONS_RETRIEVE_FAILED UPLOAD_FPR_FAILED]
	Status string `json:"status,omitempty"`

	// user name
	UserName string `json:"userName,omitempty"`
}

// Validate validates this audit assistant status
func (m *AuditAssistantStatus) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var auditAssistantStatusTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","PRE_PREDICT_FAILURE","PREDICT_QUEUED","PREDICT_PREP","PREDICT_SUBMITTED","PREDICT_COMPLETE","PREDICT_FAILED","PREDICTIONS_RETRIEVING","PREDICTIONS_RETRIEVED","PREDICTIONS_RETRIEVE_FAILED","UPLOAD_FPR_FAILED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		auditAssistantStatusTypeStatusPropEnum = append(auditAssistantStatusTypeStatusPropEnum, v)
	}
}

const (

	// AuditAssistantStatusStatusNONE captures enum value "NONE"
	AuditAssistantStatusStatusNONE string = "NONE"

	// AuditAssistantStatusStatusPREPREDICTFAILURE captures enum value "PRE_PREDICT_FAILURE"
	AuditAssistantStatusStatusPREPREDICTFAILURE string = "PRE_PREDICT_FAILURE"

	// AuditAssistantStatusStatusPREDICTQUEUED captures enum value "PREDICT_QUEUED"
	AuditAssistantStatusStatusPREDICTQUEUED string = "PREDICT_QUEUED"

	// AuditAssistantStatusStatusPREDICTPREP captures enum value "PREDICT_PREP"
	AuditAssistantStatusStatusPREDICTPREP string = "PREDICT_PREP"

	// AuditAssistantStatusStatusPREDICTSUBMITTED captures enum value "PREDICT_SUBMITTED"
	AuditAssistantStatusStatusPREDICTSUBMITTED string = "PREDICT_SUBMITTED"

	// AuditAssistantStatusStatusPREDICTCOMPLETE captures enum value "PREDICT_COMPLETE"
	AuditAssistantStatusStatusPREDICTCOMPLETE string = "PREDICT_COMPLETE"

	// AuditAssistantStatusStatusPREDICTFAILED captures enum value "PREDICT_FAILED"
	AuditAssistantStatusStatusPREDICTFAILED string = "PREDICT_FAILED"

	// AuditAssistantStatusStatusPREDICTIONSRETRIEVING captures enum value "PREDICTIONS_RETRIEVING"
	AuditAssistantStatusStatusPREDICTIONSRETRIEVING string = "PREDICTIONS_RETRIEVING"

	// AuditAssistantStatusStatusPREDICTIONSRETRIEVED captures enum value "PREDICTIONS_RETRIEVED"
	AuditAssistantStatusStatusPREDICTIONSRETRIEVED string = "PREDICTIONS_RETRIEVED"

	// AuditAssistantStatusStatusPREDICTIONSRETRIEVEFAILED captures enum value "PREDICTIONS_RETRIEVE_FAILED"
	AuditAssistantStatusStatusPREDICTIONSRETRIEVEFAILED string = "PREDICTIONS_RETRIEVE_FAILED"

	// AuditAssistantStatusStatusUPLOADFPRFAILED captures enum value "UPLOAD_FPR_FAILED"
	AuditAssistantStatusStatusUPLOADFPRFAILED string = "UPLOAD_FPR_FAILED"
)

// prop value enum
func (m *AuditAssistantStatus) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, auditAssistantStatusTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *AuditAssistantStatus) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *AuditAssistantStatus) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *AuditAssistantStatus) UnmarshalBinary(b []byte) error {
	var res AuditAssistantStatus
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
