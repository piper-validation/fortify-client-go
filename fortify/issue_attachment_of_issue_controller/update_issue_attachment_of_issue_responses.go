// Code generated by go-swagger; DO NOT EDIT.

package issue_attachment_of_issue_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// UpdateIssueAttachmentOfIssueReader is a Reader for the UpdateIssueAttachmentOfIssue structure.
type UpdateIssueAttachmentOfIssueReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIssueAttachmentOfIssueReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIssueAttachmentOfIssueOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateIssueAttachmentOfIssueBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateIssueAttachmentOfIssueUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateIssueAttachmentOfIssueForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateIssueAttachmentOfIssueNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateIssueAttachmentOfIssueConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateIssueAttachmentOfIssueInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateIssueAttachmentOfIssueOK creates a UpdateIssueAttachmentOfIssueOK with default headers values
func NewUpdateIssueAttachmentOfIssueOK() *UpdateIssueAttachmentOfIssueOK {
	return &UpdateIssueAttachmentOfIssueOK{}
}

/*UpdateIssueAttachmentOfIssueOK handles this case with default header values.

OK
*/
type UpdateIssueAttachmentOfIssueOK struct {
	Payload *models.APIResultIssueAttachment
}

func (o *UpdateIssueAttachmentOfIssueOK) Error() string {
	return fmt.Sprintf("[PUT /issues/{parentId}/attachments/{id}][%d] updateIssueAttachmentOfIssueOK  %+v", 200, o.Payload)
}

func (o *UpdateIssueAttachmentOfIssueOK) GetPayload() *models.APIResultIssueAttachment {
	return o.Payload
}

func (o *UpdateIssueAttachmentOfIssueOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueAttachment)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIssueAttachmentOfIssueBadRequest creates a UpdateIssueAttachmentOfIssueBadRequest with default headers values
func NewUpdateIssueAttachmentOfIssueBadRequest() *UpdateIssueAttachmentOfIssueBadRequest {
	return &UpdateIssueAttachmentOfIssueBadRequest{}
}

/*UpdateIssueAttachmentOfIssueBadRequest handles this case with default header values.

Bad Request
*/
type UpdateIssueAttachmentOfIssueBadRequest struct {
	Payload *models.APIResult
}

func (o *UpdateIssueAttachmentOfIssueBadRequest) Error() string {
	return fmt.Sprintf("[PUT /issues/{parentId}/attachments/{id}][%d] updateIssueAttachmentOfIssueBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateIssueAttachmentOfIssueBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateIssueAttachmentOfIssueBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIssueAttachmentOfIssueUnauthorized creates a UpdateIssueAttachmentOfIssueUnauthorized with default headers values
func NewUpdateIssueAttachmentOfIssueUnauthorized() *UpdateIssueAttachmentOfIssueUnauthorized {
	return &UpdateIssueAttachmentOfIssueUnauthorized{}
}

/*UpdateIssueAttachmentOfIssueUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateIssueAttachmentOfIssueUnauthorized struct {
	Payload *models.APIResult
}

func (o *UpdateIssueAttachmentOfIssueUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /issues/{parentId}/attachments/{id}][%d] updateIssueAttachmentOfIssueUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateIssueAttachmentOfIssueUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateIssueAttachmentOfIssueUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIssueAttachmentOfIssueForbidden creates a UpdateIssueAttachmentOfIssueForbidden with default headers values
func NewUpdateIssueAttachmentOfIssueForbidden() *UpdateIssueAttachmentOfIssueForbidden {
	return &UpdateIssueAttachmentOfIssueForbidden{}
}

/*UpdateIssueAttachmentOfIssueForbidden handles this case with default header values.

Forbidden
*/
type UpdateIssueAttachmentOfIssueForbidden struct {
	Payload *models.APIResult
}

func (o *UpdateIssueAttachmentOfIssueForbidden) Error() string {
	return fmt.Sprintf("[PUT /issues/{parentId}/attachments/{id}][%d] updateIssueAttachmentOfIssueForbidden  %+v", 403, o.Payload)
}

func (o *UpdateIssueAttachmentOfIssueForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateIssueAttachmentOfIssueForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIssueAttachmentOfIssueNotFound creates a UpdateIssueAttachmentOfIssueNotFound with default headers values
func NewUpdateIssueAttachmentOfIssueNotFound() *UpdateIssueAttachmentOfIssueNotFound {
	return &UpdateIssueAttachmentOfIssueNotFound{}
}

/*UpdateIssueAttachmentOfIssueNotFound handles this case with default header values.

Not Found
*/
type UpdateIssueAttachmentOfIssueNotFound struct {
	Payload *models.APIResult
}

func (o *UpdateIssueAttachmentOfIssueNotFound) Error() string {
	return fmt.Sprintf("[PUT /issues/{parentId}/attachments/{id}][%d] updateIssueAttachmentOfIssueNotFound  %+v", 404, o.Payload)
}

func (o *UpdateIssueAttachmentOfIssueNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateIssueAttachmentOfIssueNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIssueAttachmentOfIssueConflict creates a UpdateIssueAttachmentOfIssueConflict with default headers values
func NewUpdateIssueAttachmentOfIssueConflict() *UpdateIssueAttachmentOfIssueConflict {
	return &UpdateIssueAttachmentOfIssueConflict{}
}

/*UpdateIssueAttachmentOfIssueConflict handles this case with default header values.

Conflict
*/
type UpdateIssueAttachmentOfIssueConflict struct {
	Payload *models.APIResult
}

func (o *UpdateIssueAttachmentOfIssueConflict) Error() string {
	return fmt.Sprintf("[PUT /issues/{parentId}/attachments/{id}][%d] updateIssueAttachmentOfIssueConflict  %+v", 409, o.Payload)
}

func (o *UpdateIssueAttachmentOfIssueConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateIssueAttachmentOfIssueConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIssueAttachmentOfIssueInternalServerError creates a UpdateIssueAttachmentOfIssueInternalServerError with default headers values
func NewUpdateIssueAttachmentOfIssueInternalServerError() *UpdateIssueAttachmentOfIssueInternalServerError {
	return &UpdateIssueAttachmentOfIssueInternalServerError{}
}

/*UpdateIssueAttachmentOfIssueInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateIssueAttachmentOfIssueInternalServerError struct {
	Payload *models.APIResult
}

func (o *UpdateIssueAttachmentOfIssueInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /issues/{parentId}/attachments/{id}][%d] updateIssueAttachmentOfIssueInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateIssueAttachmentOfIssueInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateIssueAttachmentOfIssueInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
