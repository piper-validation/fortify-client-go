// Code generated by go-swagger; DO NOT EDIT.

package issue_view_template_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListIssueViewTemplateReader is a Reader for the ListIssueViewTemplate structure.
type ListIssueViewTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListIssueViewTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListIssueViewTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListIssueViewTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListIssueViewTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListIssueViewTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListIssueViewTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListIssueViewTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListIssueViewTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListIssueViewTemplateOK creates a ListIssueViewTemplateOK with default headers values
func NewListIssueViewTemplateOK() *ListIssueViewTemplateOK {
	return &ListIssueViewTemplateOK{}
}

/*ListIssueViewTemplateOK handles this case with default header values.

OK
*/
type ListIssueViewTemplateOK struct {
	Payload *models.APIResultListIssueViewTemplate
}

func (o *ListIssueViewTemplateOK) Error() string {
	return fmt.Sprintf("[GET /issueViewTemplates][%d] listIssueViewTemplateOK  %+v", 200, o.Payload)
}

func (o *ListIssueViewTemplateOK) GetPayload() *models.APIResultListIssueViewTemplate {
	return o.Payload
}

func (o *ListIssueViewTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListIssueViewTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueViewTemplateBadRequest creates a ListIssueViewTemplateBadRequest with default headers values
func NewListIssueViewTemplateBadRequest() *ListIssueViewTemplateBadRequest {
	return &ListIssueViewTemplateBadRequest{}
}

/*ListIssueViewTemplateBadRequest handles this case with default header values.

Bad Request
*/
type ListIssueViewTemplateBadRequest struct {
	Payload *models.APIResult
}

func (o *ListIssueViewTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /issueViewTemplates][%d] listIssueViewTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *ListIssueViewTemplateBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueViewTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueViewTemplateUnauthorized creates a ListIssueViewTemplateUnauthorized with default headers values
func NewListIssueViewTemplateUnauthorized() *ListIssueViewTemplateUnauthorized {
	return &ListIssueViewTemplateUnauthorized{}
}

/*ListIssueViewTemplateUnauthorized handles this case with default header values.

Unauthorized
*/
type ListIssueViewTemplateUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListIssueViewTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /issueViewTemplates][%d] listIssueViewTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *ListIssueViewTemplateUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueViewTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueViewTemplateForbidden creates a ListIssueViewTemplateForbidden with default headers values
func NewListIssueViewTemplateForbidden() *ListIssueViewTemplateForbidden {
	return &ListIssueViewTemplateForbidden{}
}

/*ListIssueViewTemplateForbidden handles this case with default header values.

Forbidden
*/
type ListIssueViewTemplateForbidden struct {
	Payload *models.APIResult
}

func (o *ListIssueViewTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /issueViewTemplates][%d] listIssueViewTemplateForbidden  %+v", 403, o.Payload)
}

func (o *ListIssueViewTemplateForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueViewTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueViewTemplateNotFound creates a ListIssueViewTemplateNotFound with default headers values
func NewListIssueViewTemplateNotFound() *ListIssueViewTemplateNotFound {
	return &ListIssueViewTemplateNotFound{}
}

/*ListIssueViewTemplateNotFound handles this case with default header values.

Not Found
*/
type ListIssueViewTemplateNotFound struct {
	Payload *models.APIResult
}

func (o *ListIssueViewTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /issueViewTemplates][%d] listIssueViewTemplateNotFound  %+v", 404, o.Payload)
}

func (o *ListIssueViewTemplateNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueViewTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueViewTemplateConflict creates a ListIssueViewTemplateConflict with default headers values
func NewListIssueViewTemplateConflict() *ListIssueViewTemplateConflict {
	return &ListIssueViewTemplateConflict{}
}

/*ListIssueViewTemplateConflict handles this case with default header values.

Conflict
*/
type ListIssueViewTemplateConflict struct {
	Payload *models.APIResult
}

func (o *ListIssueViewTemplateConflict) Error() string {
	return fmt.Sprintf("[GET /issueViewTemplates][%d] listIssueViewTemplateConflict  %+v", 409, o.Payload)
}

func (o *ListIssueViewTemplateConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueViewTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListIssueViewTemplateInternalServerError creates a ListIssueViewTemplateInternalServerError with default headers values
func NewListIssueViewTemplateInternalServerError() *ListIssueViewTemplateInternalServerError {
	return &ListIssueViewTemplateInternalServerError{}
}

/*ListIssueViewTemplateInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListIssueViewTemplateInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListIssueViewTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /issueViewTemplates][%d] listIssueViewTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *ListIssueViewTemplateInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListIssueViewTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
