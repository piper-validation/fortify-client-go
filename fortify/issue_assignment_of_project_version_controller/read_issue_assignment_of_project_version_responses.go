// Code generated by go-swagger; DO NOT EDIT.

package issue_assignment_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadIssueAssignmentOfProjectVersionReader is a Reader for the ReadIssueAssignmentOfProjectVersion structure.
type ReadIssueAssignmentOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadIssueAssignmentOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadIssueAssignmentOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadIssueAssignmentOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadIssueAssignmentOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadIssueAssignmentOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadIssueAssignmentOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadIssueAssignmentOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadIssueAssignmentOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadIssueAssignmentOfProjectVersionOK creates a ReadIssueAssignmentOfProjectVersionOK with default headers values
func NewReadIssueAssignmentOfProjectVersionOK() *ReadIssueAssignmentOfProjectVersionOK {
	return &ReadIssueAssignmentOfProjectVersionOK{}
}

/*ReadIssueAssignmentOfProjectVersionOK handles this case with default header values.

OK
*/
type ReadIssueAssignmentOfProjectVersionOK struct {
	Payload *models.APIResultIssueAssignment
}

func (o *ReadIssueAssignmentOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/issueAssignment/{id}][%d] readIssueAssignmentOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ReadIssueAssignmentOfProjectVersionOK) GetPayload() *models.APIResultIssueAssignment {
	return o.Payload
}

func (o *ReadIssueAssignmentOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueAssignment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAssignmentOfProjectVersionBadRequest creates a ReadIssueAssignmentOfProjectVersionBadRequest with default headers values
func NewReadIssueAssignmentOfProjectVersionBadRequest() *ReadIssueAssignmentOfProjectVersionBadRequest {
	return &ReadIssueAssignmentOfProjectVersionBadRequest{}
}

/*ReadIssueAssignmentOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ReadIssueAssignmentOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadIssueAssignmentOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/issueAssignment/{id}][%d] readIssueAssignmentOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ReadIssueAssignmentOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAssignmentOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAssignmentOfProjectVersionUnauthorized creates a ReadIssueAssignmentOfProjectVersionUnauthorized with default headers values
func NewReadIssueAssignmentOfProjectVersionUnauthorized() *ReadIssueAssignmentOfProjectVersionUnauthorized {
	return &ReadIssueAssignmentOfProjectVersionUnauthorized{}
}

/*ReadIssueAssignmentOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadIssueAssignmentOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadIssueAssignmentOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/issueAssignment/{id}][%d] readIssueAssignmentOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadIssueAssignmentOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAssignmentOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAssignmentOfProjectVersionForbidden creates a ReadIssueAssignmentOfProjectVersionForbidden with default headers values
func NewReadIssueAssignmentOfProjectVersionForbidden() *ReadIssueAssignmentOfProjectVersionForbidden {
	return &ReadIssueAssignmentOfProjectVersionForbidden{}
}

/*ReadIssueAssignmentOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ReadIssueAssignmentOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ReadIssueAssignmentOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/issueAssignment/{id}][%d] readIssueAssignmentOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ReadIssueAssignmentOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAssignmentOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAssignmentOfProjectVersionNotFound creates a ReadIssueAssignmentOfProjectVersionNotFound with default headers values
func NewReadIssueAssignmentOfProjectVersionNotFound() *ReadIssueAssignmentOfProjectVersionNotFound {
	return &ReadIssueAssignmentOfProjectVersionNotFound{}
}

/*ReadIssueAssignmentOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ReadIssueAssignmentOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ReadIssueAssignmentOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/issueAssignment/{id}][%d] readIssueAssignmentOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ReadIssueAssignmentOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAssignmentOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAssignmentOfProjectVersionConflict creates a ReadIssueAssignmentOfProjectVersionConflict with default headers values
func NewReadIssueAssignmentOfProjectVersionConflict() *ReadIssueAssignmentOfProjectVersionConflict {
	return &ReadIssueAssignmentOfProjectVersionConflict{}
}

/*ReadIssueAssignmentOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ReadIssueAssignmentOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ReadIssueAssignmentOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/issueAssignment/{id}][%d] readIssueAssignmentOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ReadIssueAssignmentOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAssignmentOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueAssignmentOfProjectVersionInternalServerError creates a ReadIssueAssignmentOfProjectVersionInternalServerError with default headers values
func NewReadIssueAssignmentOfProjectVersionInternalServerError() *ReadIssueAssignmentOfProjectVersionInternalServerError {
	return &ReadIssueAssignmentOfProjectVersionInternalServerError{}
}

/*ReadIssueAssignmentOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadIssueAssignmentOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadIssueAssignmentOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/issueAssignment/{id}][%d] readIssueAssignmentOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadIssueAssignmentOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueAssignmentOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
