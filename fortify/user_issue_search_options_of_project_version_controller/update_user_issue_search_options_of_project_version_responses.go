// Code generated by go-swagger; DO NOT EDIT.

package user_issue_search_options_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// UpdateUserIssueSearchOptionsOfProjectVersionReader is a Reader for the UpdateUserIssueSearchOptionsOfProjectVersion structure.
type UpdateUserIssueSearchOptionsOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateUserIssueSearchOptionsOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateUserIssueSearchOptionsOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateUserIssueSearchOptionsOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateUserIssueSearchOptionsOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateUserIssueSearchOptionsOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateUserIssueSearchOptionsOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateUserIssueSearchOptionsOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateUserIssueSearchOptionsOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateUserIssueSearchOptionsOfProjectVersionOK creates a UpdateUserIssueSearchOptionsOfProjectVersionOK with default headers values
func NewUpdateUserIssueSearchOptionsOfProjectVersionOK() *UpdateUserIssueSearchOptionsOfProjectVersionOK {
	return &UpdateUserIssueSearchOptionsOfProjectVersionOK{}
}

/*UpdateUserIssueSearchOptionsOfProjectVersionOK handles this case with default header values.

OK
*/
type UpdateUserIssueSearchOptionsOfProjectVersionOK struct {
	Payload *models.APIResultUserIssueSearchOptions
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/userIssueSearchOptions][%d] updateUserIssueSearchOptionsOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionOK) GetPayload() *models.APIResultUserIssueSearchOptions {
	return o.Payload
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultUserIssueSearchOptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserIssueSearchOptionsOfProjectVersionBadRequest creates a UpdateUserIssueSearchOptionsOfProjectVersionBadRequest with default headers values
func NewUpdateUserIssueSearchOptionsOfProjectVersionBadRequest() *UpdateUserIssueSearchOptionsOfProjectVersionBadRequest {
	return &UpdateUserIssueSearchOptionsOfProjectVersionBadRequest{}
}

/*UpdateUserIssueSearchOptionsOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateUserIssueSearchOptionsOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/userIssueSearchOptions][%d] updateUserIssueSearchOptionsOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserIssueSearchOptionsOfProjectVersionUnauthorized creates a UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized with default headers values
func NewUpdateUserIssueSearchOptionsOfProjectVersionUnauthorized() *UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized {
	return &UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized{}
}

/*UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/userIssueSearchOptions][%d] updateUserIssueSearchOptionsOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserIssueSearchOptionsOfProjectVersionForbidden creates a UpdateUserIssueSearchOptionsOfProjectVersionForbidden with default headers values
func NewUpdateUserIssueSearchOptionsOfProjectVersionForbidden() *UpdateUserIssueSearchOptionsOfProjectVersionForbidden {
	return &UpdateUserIssueSearchOptionsOfProjectVersionForbidden{}
}

/*UpdateUserIssueSearchOptionsOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type UpdateUserIssueSearchOptionsOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/userIssueSearchOptions][%d] updateUserIssueSearchOptionsOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserIssueSearchOptionsOfProjectVersionNotFound creates a UpdateUserIssueSearchOptionsOfProjectVersionNotFound with default headers values
func NewUpdateUserIssueSearchOptionsOfProjectVersionNotFound() *UpdateUserIssueSearchOptionsOfProjectVersionNotFound {
	return &UpdateUserIssueSearchOptionsOfProjectVersionNotFound{}
}

/*UpdateUserIssueSearchOptionsOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type UpdateUserIssueSearchOptionsOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/userIssueSearchOptions][%d] updateUserIssueSearchOptionsOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserIssueSearchOptionsOfProjectVersionConflict creates a UpdateUserIssueSearchOptionsOfProjectVersionConflict with default headers values
func NewUpdateUserIssueSearchOptionsOfProjectVersionConflict() *UpdateUserIssueSearchOptionsOfProjectVersionConflict {
	return &UpdateUserIssueSearchOptionsOfProjectVersionConflict{}
}

/*UpdateUserIssueSearchOptionsOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type UpdateUserIssueSearchOptionsOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/userIssueSearchOptions][%d] updateUserIssueSearchOptionsOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateUserIssueSearchOptionsOfProjectVersionInternalServerError creates a UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError with default headers values
func NewUpdateUserIssueSearchOptionsOfProjectVersionInternalServerError() *UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError {
	return &UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError{}
}

/*UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/userIssueSearchOptions][%d] updateUserIssueSearchOptionsOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateUserIssueSearchOptionsOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
