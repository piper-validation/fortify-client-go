// Code generated by go-swagger; DO NOT EDIT.

package issue_aging_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListIssueAgingGroupReader is a Reader for the ListIssueAgingGroup structure.
type ListIssueAgingGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIssueAgingGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIssueAgingGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListIssueAgingGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListIssueAgingGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListIssueAgingGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListIssueAgingGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListIssueAgingGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListIssueAgingGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIssueAgingGroupOK creates a ListIssueAgingGroupOK with default headers values
func NewListIssueAgingGroupOK() *ListIssueAgingGroupOK {
	return &ListIssueAgingGroupOK{}
}

/*ListIssueAgingGroupOK handles this case with default header values.

OK
*/
type ListIssueAgingGroupOK struct {
	Payload *models.APIResultListIssueAgingGroupDto
}

func (o *ListIssueAgingGroupOK) Error() string {
	return fmt.Sprintf("[GET /issueaginggroup][%d] listIssueAgingGroupOK  %+v", 200, o.Payload)
}

func (o *ListIssueAgingGroupOK) GetPayload() *models.APIResultListIssueAgingGroupDto {
	return o.Payload
}

func (o *ListIssueAgingGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListIssueAgingGroupDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAgingGroupBadRequest creates a ListIssueAgingGroupBadRequest with default headers values
func NewListIssueAgingGroupBadRequest() *ListIssueAgingGroupBadRequest {
	return &ListIssueAgingGroupBadRequest{}
}

/*ListIssueAgingGroupBadRequest handles this case with default header values.

Bad Request
*/
type ListIssueAgingGroupBadRequest struct {
	Payload *models.APIResult
}

func (o *ListIssueAgingGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /issueaginggroup][%d] listIssueAgingGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListIssueAgingGroupBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAgingGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAgingGroupUnauthorized creates a ListIssueAgingGroupUnauthorized with default headers values
func NewListIssueAgingGroupUnauthorized() *ListIssueAgingGroupUnauthorized {
	return &ListIssueAgingGroupUnauthorized{}
}

/*ListIssueAgingGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type ListIssueAgingGroupUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListIssueAgingGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /issueaginggroup][%d] listIssueAgingGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *ListIssueAgingGroupUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAgingGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAgingGroupForbidden creates a ListIssueAgingGroupForbidden with default headers values
func NewListIssueAgingGroupForbidden() *ListIssueAgingGroupForbidden {
	return &ListIssueAgingGroupForbidden{}
}

/*ListIssueAgingGroupForbidden handles this case with default header values.

Forbidden
*/
type ListIssueAgingGroupForbidden struct {
	Payload *models.APIResult
}

func (o *ListIssueAgingGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /issueaginggroup][%d] listIssueAgingGroupForbidden  %+v", 403, o.Payload)
}

func (o *ListIssueAgingGroupForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAgingGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAgingGroupNotFound creates a ListIssueAgingGroupNotFound with default headers values
func NewListIssueAgingGroupNotFound() *ListIssueAgingGroupNotFound {
	return &ListIssueAgingGroupNotFound{}
}

/*ListIssueAgingGroupNotFound handles this case with default header values.

Not Found
*/
type ListIssueAgingGroupNotFound struct {
	Payload *models.APIResult
}

func (o *ListIssueAgingGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /issueaginggroup][%d] listIssueAgingGroupNotFound  %+v", 404, o.Payload)
}

func (o *ListIssueAgingGroupNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAgingGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAgingGroupConflict creates a ListIssueAgingGroupConflict with default headers values
func NewListIssueAgingGroupConflict() *ListIssueAgingGroupConflict {
	return &ListIssueAgingGroupConflict{}
}

/*ListIssueAgingGroupConflict handles this case with default header values.

Conflict
*/
type ListIssueAgingGroupConflict struct {
	Payload *models.APIResult
}

func (o *ListIssueAgingGroupConflict) Error() string {
	return fmt.Sprintf("[GET /issueaginggroup][%d] listIssueAgingGroupConflict  %+v", 409, o.Payload)
}

func (o *ListIssueAgingGroupConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAgingGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueAgingGroupInternalServerError creates a ListIssueAgingGroupInternalServerError with default headers values
func NewListIssueAgingGroupInternalServerError() *ListIssueAgingGroupInternalServerError {
	return &ListIssueAgingGroupInternalServerError{}
}

/*ListIssueAgingGroupInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListIssueAgingGroupInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListIssueAgingGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /issueaginggroup][%d] listIssueAgingGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListIssueAgingGroupInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueAgingGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}