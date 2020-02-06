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

// ProjectVersionIssue Application version issue DTO object
// swagger:model Project version issue
type ProjectVersionIssue struct {

	// Analyzer
	// Required: true
	Analyzer *string `json:"analyzer"`

	// Flag is set for issues that has been audited and primary tag value was set for this issue.
	Audited bool `json:"audited,omitempty"`

	// Bug url
	// Required: true
	BugURL *string `json:"bugURL"`

	// Issue confidence
	// Required: true
	Confidence *float32 `json:"confidence"`

	// Display name for engine type
	// Required: true
	DisplayEngineType *string `json:"displayEngineType"`

	// Engine category
	// Required: true
	// Enum: [STATIC DYNAMIC]
	EngineCategory *string `json:"engineCategory"`

	// Engine type
	// Required: true
	EngineType *string `json:"engineType"`

	// Identifying information for the bug in the external bug tracking system. The actual format depends on the bug tracker plugin which filed the bug.
	// Required: true
	ExternalBugID *string `json:"externalBugId"`

	// Issue folder globally unique identifier
	// Required: true
	FolderGUID *string `json:"folderGuid"`

	// Deprecated - Issue folder identifier.  This may be incorrect or invalid.  Please use folderGuid instead.
	// Required: true
	FolderID *int64 `json:"folderId"`

	// Date when issue found
	// Required: true
	// Format: date-time
	FoundDate *strfmt.DateTime `json:"foundDate"`

	// Friority
	// Required: true
	Friority *string `json:"friority"`

	// Full file name where issue found
	// Required: true
	FullFileName *string `json:"fullFileName"`

	// Set to true if issue has attachments
	// Required: true
	HasAttachments *bool `json:"hasAttachments"`

	// Set to true if issue has other correlated issues
	// Required: true
	HasCorrelatedIssues *bool `json:"hasCorrelatedIssues"`

	// Set to true if issue is hidden
	// Required: true
	Hidden *bool `json:"hidden"`

	// Application version issue identifier
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Issue impact
	// Required: true
	Impact *float32 `json:"impact"`

	// Issue instance identifier
	// Required: true
	IssueInstanceID *string `json:"issueInstanceId"`

	// Issue name
	// Required: true
	IssueName *string `json:"issueName"`

	// Flag represents issue review status and can have 3 types of values: Unreviewed, Under Review, Reviewed.
	IssueStatus string `json:"issueStatus,omitempty"`

	// Kingdom
	// Required: true
	Kingdom *string `json:"kingdom"`

	// ID of the latest scan that found the issue
	LastScanID int64 `json:"lastScanId,omitempty"`

	// Likelihood of issue
	// Required: true
	Likelihood *float32 `json:"likelihood"`

	// Line number where issue found
	// Required: true
	LineNumber *int32 `json:"lineNumber"`

	// Issue primary location
	// Required: true
	PrimaryLocation *string `json:"primaryLocation"`

	// Primary rule global unique identifier
	// Required: true
	PrimaryRuleGUID *string `json:"primaryRuleGuid"`

	// Issue primary tag
	// Required: true
	PrimaryTag *string `json:"primaryTag"`

	// Flag equals true if value of custom tag was applied automatically
	PrimaryTagValueAutoApplied bool `json:"primaryTagValueAutoApplied,omitempty"`

	// Application name
	// Required: true
	ProjectName *string `json:"projectName"`

	// Application version identifier
	// Required: true
	ProjectVersionID *int64 `json:"projectVersionId"`

	// Application version name
	// Required: true
	ProjectVersionName *string `json:"projectVersionName"`

	// Set to true if issue is suppressed
	// Required: true
	Removed *bool `json:"removed"`

	// Date when issue removed
	// Required: true
	// Format: date-time
	RemovedDate *strfmt.DateTime `json:"removedDate"`

	// Issue reviewer
	// Required: true
	Reviewed *string `json:"reviewed"`

	// Application version revision
	// Required: true
	Revision *int32 `json:"revision"`

	// Scan status
	// Required: true
	ScanStatus *string `json:"scanStatus"`

	// Issue severity
	// Required: true
	Severity *float32 `json:"severity"`

	// Set to true if issue is suppressed
	// Required: true
	Suppressed *bool `json:"suppressed"`
}

// Validate validates this project version issue
func (m *ProjectVersionIssue) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnalyzer(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBugURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateConfidence(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDisplayEngineType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEngineCategory(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateEngineType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExternalBugID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolderGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFolderID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFoundDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateFullFileName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasAttachments(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHasCorrelatedIssues(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHidden(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateImpact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueInstanceID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateKingdom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLikelihood(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLineNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryLocation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryRuleGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePrimaryTag(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectVersionID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProjectVersionName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemoved(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRemovedDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateReviewed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRevision(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateScanStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSeverity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSuppressed(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVersionIssue) validateAnalyzer(formats strfmt.Registry) error {

	if err := validate.Required("analyzer", "body", m.Analyzer); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateBugURL(formats strfmt.Registry) error {

	if err := validate.Required("bugURL", "body", m.BugURL); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateConfidence(formats strfmt.Registry) error {

	if err := validate.Required("confidence", "body", m.Confidence); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateDisplayEngineType(formats strfmt.Registry) error {

	if err := validate.Required("displayEngineType", "body", m.DisplayEngineType); err != nil {
		return err
	}

	return nil
}

var projectVersionIssueTypeEngineCategoryPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["STATIC","DYNAMIC"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectVersionIssueTypeEngineCategoryPropEnum = append(projectVersionIssueTypeEngineCategoryPropEnum, v)
	}
}

const (

	// ProjectVersionIssueEngineCategorySTATIC captures enum value "STATIC"
	ProjectVersionIssueEngineCategorySTATIC string = "STATIC"

	// ProjectVersionIssueEngineCategoryDYNAMIC captures enum value "DYNAMIC"
	ProjectVersionIssueEngineCategoryDYNAMIC string = "DYNAMIC"
)

// prop value enum
func (m *ProjectVersionIssue) validateEngineCategoryEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, projectVersionIssueTypeEngineCategoryPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProjectVersionIssue) validateEngineCategory(formats strfmt.Registry) error {

	if err := validate.Required("engineCategory", "body", m.EngineCategory); err != nil {
		return err
	}

	// value enum
	if err := m.validateEngineCategoryEnum("engineCategory", "body", *m.EngineCategory); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateEngineType(formats strfmt.Registry) error {

	if err := validate.Required("engineType", "body", m.EngineType); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateExternalBugID(formats strfmt.Registry) error {

	if err := validate.Required("externalBugId", "body", m.ExternalBugID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateFolderGUID(formats strfmt.Registry) error {

	if err := validate.Required("folderGuid", "body", m.FolderGUID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateFolderID(formats strfmt.Registry) error {

	if err := validate.Required("folderId", "body", m.FolderID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateFoundDate(formats strfmt.Registry) error {

	if err := validate.Required("foundDate", "body", m.FoundDate); err != nil {
		return err
	}

	if err := validate.FormatOf("foundDate", "body", "date-time", m.FoundDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateFriority(formats strfmt.Registry) error {

	if err := validate.Required("friority", "body", m.Friority); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateFullFileName(formats strfmt.Registry) error {

	if err := validate.Required("fullFileName", "body", m.FullFileName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateHasAttachments(formats strfmt.Registry) error {

	if err := validate.Required("hasAttachments", "body", m.HasAttachments); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateHasCorrelatedIssues(formats strfmt.Registry) error {

	if err := validate.Required("hasCorrelatedIssues", "body", m.HasCorrelatedIssues); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateHidden(formats strfmt.Registry) error {

	if err := validate.Required("hidden", "body", m.Hidden); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateImpact(formats strfmt.Registry) error {

	if err := validate.Required("impact", "body", m.Impact); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateIssueInstanceID(formats strfmt.Registry) error {

	if err := validate.Required("issueInstanceId", "body", m.IssueInstanceID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateIssueName(formats strfmt.Registry) error {

	if err := validate.Required("issueName", "body", m.IssueName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateKingdom(formats strfmt.Registry) error {

	if err := validate.Required("kingdom", "body", m.Kingdom); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateLikelihood(formats strfmt.Registry) error {

	if err := validate.Required("likelihood", "body", m.Likelihood); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateLineNumber(formats strfmt.Registry) error {

	if err := validate.Required("lineNumber", "body", m.LineNumber); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validatePrimaryLocation(formats strfmt.Registry) error {

	if err := validate.Required("primaryLocation", "body", m.PrimaryLocation); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validatePrimaryRuleGUID(formats strfmt.Registry) error {

	if err := validate.Required("primaryRuleGuid", "body", m.PrimaryRuleGUID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validatePrimaryTag(formats strfmt.Registry) error {

	if err := validate.Required("primaryTag", "body", m.PrimaryTag); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateProjectName(formats strfmt.Registry) error {

	if err := validate.Required("projectName", "body", m.ProjectName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateProjectVersionID(formats strfmt.Registry) error {

	if err := validate.Required("projectVersionId", "body", m.ProjectVersionID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateProjectVersionName(formats strfmt.Registry) error {

	if err := validate.Required("projectVersionName", "body", m.ProjectVersionName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateRemoved(formats strfmt.Registry) error {

	if err := validate.Required("removed", "body", m.Removed); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateRemovedDate(formats strfmt.Registry) error {

	if err := validate.Required("removedDate", "body", m.RemovedDate); err != nil {
		return err
	}

	if err := validate.FormatOf("removedDate", "body", "date-time", m.RemovedDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateReviewed(formats strfmt.Registry) error {

	if err := validate.Required("reviewed", "body", m.Reviewed); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateRevision(formats strfmt.Registry) error {

	if err := validate.Required("revision", "body", m.Revision); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateScanStatus(formats strfmt.Registry) error {

	if err := validate.Required("scanStatus", "body", m.ScanStatus); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateSeverity(formats strfmt.Registry) error {

	if err := validate.Required("severity", "body", m.Severity); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersionIssue) validateSuppressed(formats strfmt.Registry) error {

	if err := validate.Required("suppressed", "body", m.Suppressed); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ProjectVersionIssue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVersionIssue) UnmarshalBinary(b []byte) error {
	var res ProjectVersionIssue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}