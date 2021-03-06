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

// ProjectVersion A version of an application on the server
// swagger:model Project Version
type ProjectVersion struct {

	// True if this version is active
	// Required: true
	Active *bool `json:"active"`

	// assigned issues count
	AssignedIssuesCount int64 `json:"assignedIssuesCount,omitempty"`

	// attachments out of date
	AttachmentsOutOfDate bool `json:"attachmentsOutOfDate,omitempty"`

	// true if auto-prediction is enabled for this application version; false otherwise. This property modification is protected by AUDITASSISTANT_MANAGE permission.
	AutoPredict bool `json:"autoPredict,omitempty"`

	// true if the bug tracker plugin assigned to the application version is currently enabled in the system
	// Required: true
	BugTrackerEnabled *bool `json:"bugTrackerEnabled"`

	// identifier of the bug tracker plugin (if any) assigned to this application version
	// Required: true
	BugTrackerPluginID *string `json:"bugTrackerPluginId"`

	// True if this version is committed and ready to be used
	// Required: true
	Committed *bool `json:"committed"`

	// User that created this version
	// Required: true
	CreatedBy *string `json:"createdBy"`

	// Date this version was created
	// Required: true
	// Format: date-time
	CreationDate *Iso8601MilliDateTime `json:"creationDate"`

	// current state
	CurrentState *ProjectVersionState `json:"currentState,omitempty"`

	// true if custom tag values auto-apply is enabled for this application version; false otherwise. This property modification is protected by AUDITASSISTANT_MANAGE permission.
	CustomTagValuesAutoApply bool `json:"customTagValuesAutoApply,omitempty"`

	// Description
	// Required: true
	Description *string `json:"description"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Id of the Issue Template used by this version
	// Required: true
	IssueTemplateID *string `json:"issueTemplateId"`

	// Last time the Issue Template was modified
	// Required: true
	IssueTemplateModifiedTime *int64 `json:"issueTemplateModifiedTime"`

	// Name of the Issue Template used by this version
	// Required: true
	IssueTemplateName *string `json:"issueTemplateName"`

	// id of the latest scan uploaded to application version
	// Required: true
	LatestScanID *int64 `json:"latestScanId"`

	// load properties
	LoadProperties string `json:"loadProperties,omitempty"`

	// Guid of the primary custom tag for this version
	// Required: true
	MasterAttrGUID *string `json:"masterAttrGuid"`

	// migration version
	MigrationVersion float32 `json:"migrationVersion,omitempty"`

	// mode
	// Enum: [NONE ASSESSMENT BASIC FULL]
	Mode string `json:"mode,omitempty"`

	// Name
	// Required: true
	Name *string `json:"name"`

	// obfuscated Id
	ObfuscatedID string `json:"obfuscatedId,omitempty"`

	// Owner of this version
	// Required: true
	Owner *string `json:"owner"`

	// Name of the policy to be used for predictions for this application version. Modification of this property is protected by AUDITASSISTANT_MANAGE permission.
	PredictionPolicy string `json:"predictionPolicy,omitempty"`

	// project
	Project *Project `json:"project,omitempty"`

	// refresh required
	RefreshRequired bool `json:"refreshRequired,omitempty"`

	// security group
	SecurityGroup string `json:"securityGroup,omitempty"`

	// Server version this version's data supports
	// Required: true
	ServerVersion *float32 `json:"serverVersion"`

	// site Id
	SiteID string `json:"siteId,omitempty"`

	// True if the most recent snapshot is not accurate
	// Required: true
	SnapshotOutOfDate *bool `json:"snapshotOutOfDate"`

	// source base path
	SourceBasePath string `json:"sourceBasePath,omitempty"`

	// True if this version's Issue Template has changed or been modified
	// Required: true
	StaleIssueTemplate *bool `json:"staleIssueTemplate"`

	// status
	// Enum: [ACTIVE DELETING ARCHIVED COPYING_ISSUES]
	Status string `json:"status,omitempty"`

	// traces out of date
	TracesOutOfDate bool `json:"tracesOutOfDate,omitempty"`
}

// Validate validates this project version
func (m *ProjectVersion) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateActive(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBugTrackerEnabled(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateBugTrackerPluginID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCommitted(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrentState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueTemplateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueTemplateModifiedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueTemplateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLatestScanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMasterAttrGUID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOwner(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateProject(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateServerVersion(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSnapshotOutOfDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStaleIssueTemplate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ProjectVersion) validateActive(formats strfmt.Registry) error {

	if err := validate.Required("active", "body", m.Active); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateBugTrackerEnabled(formats strfmt.Registry) error {

	if err := validate.Required("bugTrackerEnabled", "body", m.BugTrackerEnabled); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateBugTrackerPluginID(formats strfmt.Registry) error {

	if err := validate.Required("bugTrackerPluginId", "body", m.BugTrackerPluginID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateCommitted(formats strfmt.Registry) error {

	if err := validate.Required("committed", "body", m.Committed); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateCreatedBy(formats strfmt.Registry) error {

	if err := validate.Required("createdBy", "body", m.CreatedBy); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateCreationDate(formats strfmt.Registry) error {

	if err := validate.Required("creationDate", "body", m.CreationDate); err != nil {
		return err
	}

	if err := validate.FormatOf("creationDate", "body", "date-time", m.CreationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateCurrentState(formats strfmt.Registry) error {

	if swag.IsZero(m.CurrentState) { // not required
		return nil
	}

	if m.CurrentState != nil {
		if err := m.CurrentState.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("currentState")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectVersion) validateDescription(formats strfmt.Registry) error {

	if err := validate.Required("description", "body", m.Description); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateIssueTemplateID(formats strfmt.Registry) error {

	if err := validate.Required("issueTemplateId", "body", m.IssueTemplateID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateIssueTemplateModifiedTime(formats strfmt.Registry) error {

	if err := validate.Required("issueTemplateModifiedTime", "body", m.IssueTemplateModifiedTime); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateIssueTemplateName(formats strfmt.Registry) error {

	if err := validate.Required("issueTemplateName", "body", m.IssueTemplateName); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateLatestScanID(formats strfmt.Registry) error {

	if err := validate.Required("latestScanId", "body", m.LatestScanID); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateMasterAttrGUID(formats strfmt.Registry) error {

	if err := validate.Required("masterAttrGuid", "body", m.MasterAttrGUID); err != nil {
		return err
	}

	return nil
}

var projectVersionTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["NONE","ASSESSMENT","BASIC","FULL"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectVersionTypeModePropEnum = append(projectVersionTypeModePropEnum, v)
	}
}

const (

	// ProjectVersionModeNONE captures enum value "NONE"
	ProjectVersionModeNONE string = "NONE"

	// ProjectVersionModeASSESSMENT captures enum value "ASSESSMENT"
	ProjectVersionModeASSESSMENT string = "ASSESSMENT"

	// ProjectVersionModeBASIC captures enum value "BASIC"
	ProjectVersionModeBASIC string = "BASIC"

	// ProjectVersionModeFULL captures enum value "FULL"
	ProjectVersionModeFULL string = "FULL"
)

// prop value enum
func (m *ProjectVersion) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, projectVersionTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProjectVersion) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateOwner(formats strfmt.Registry) error {

	if err := validate.Required("owner", "body", m.Owner); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateProject(formats strfmt.Registry) error {

	if swag.IsZero(m.Project) { // not required
		return nil
	}

	if m.Project != nil {
		if err := m.Project.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("project")
			}
			return err
		}
	}

	return nil
}

func (m *ProjectVersion) validateServerVersion(formats strfmt.Registry) error {

	if err := validate.Required("serverVersion", "body", m.ServerVersion); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateSnapshotOutOfDate(formats strfmt.Registry) error {

	if err := validate.Required("snapshotOutOfDate", "body", m.SnapshotOutOfDate); err != nil {
		return err
	}

	return nil
}

func (m *ProjectVersion) validateStaleIssueTemplate(formats strfmt.Registry) error {

	if err := validate.Required("staleIssueTemplate", "body", m.StaleIssueTemplate); err != nil {
		return err
	}

	return nil
}

var projectVersionTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ACTIVE","DELETING","ARCHIVED","COPYING_ISSUES"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		projectVersionTypeStatusPropEnum = append(projectVersionTypeStatusPropEnum, v)
	}
}

const (

	// ProjectVersionStatusACTIVE captures enum value "ACTIVE"
	ProjectVersionStatusACTIVE string = "ACTIVE"

	// ProjectVersionStatusDELETING captures enum value "DELETING"
	ProjectVersionStatusDELETING string = "DELETING"

	// ProjectVersionStatusARCHIVED captures enum value "ARCHIVED"
	ProjectVersionStatusARCHIVED string = "ARCHIVED"

	// ProjectVersionStatusCOPYINGISSUES captures enum value "COPYING_ISSUES"
	ProjectVersionStatusCOPYINGISSUES string = "COPYING_ISSUES"
)

// prop value enum
func (m *ProjectVersion) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, projectVersionTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *ProjectVersion) validateStatus(formats strfmt.Registry) error {

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
func (m *ProjectVersion) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ProjectVersion) UnmarshalBinary(b []byte) error {
	var res ProjectVersion
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
