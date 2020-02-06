// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// IssueAgingPortlet Issue aging portlet object contains application and issue metrics.
// swagger:model IssueAgingPortlet
type IssueAgingPortlet struct {

	// application versions
	ApplicationVersions int32 `json:"applicationVersions,omitempty"`

	// average days to remediate
	AverageDaysToRemediate int32 `json:"averageDaysToRemediate,omitempty"`

	// average days to review
	AverageDaysToReview int32 `json:"averageDaysToReview,omitempty"`

	// files scanned
	FilesScanned int64 `json:"filesScanned,omitempty"`

	// issues pending review
	IssuesPendingReview int64 `json:"issuesPendingReview,omitempty"`

	// issues remediated
	IssuesRemediated int64 `json:"issuesRemediated,omitempty"`

	// lines of code
	LinesOfCode int64 `json:"linesOfCode,omitempty"`

	// open issues
	OpenIssues int64 `json:"openIssues,omitempty"`

	// open issues reviewed
	OpenIssuesReviewed int64 `json:"openIssuesReviewed,omitempty"`
}

// Validate validates this issue aging portlet
func (m *IssueAgingPortlet) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *IssueAgingPortlet) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *IssueAgingPortlet) UnmarshalBinary(b []byte) error {
	var res IssueAgingPortlet
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}