// Code generated by go-swagger; DO NOT EDIT.

package issue_template_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadIssueTemplateReader is a Reader for the ReadIssueTemplate structure.
type ReadIssueTemplateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadIssueTemplateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadIssueTemplateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadIssueTemplateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadIssueTemplateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadIssueTemplateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadIssueTemplateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadIssueTemplateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadIssueTemplateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadIssueTemplateOK creates a ReadIssueTemplateOK with default headers values
func NewReadIssueTemplateOK() *ReadIssueTemplateOK {
	return &ReadIssueTemplateOK{}
}

/*ReadIssueTemplateOK handles this case with default header values.

OK
*/
type ReadIssueTemplateOK struct {
	Payload *models.APIResultIssueTemplate
}

func (o *ReadIssueTemplateOK) Error() string {
	return fmt.Sprintf("[GET /issueTemplates/{id}][%d] readIssueTemplateOK  %+v", 200, o.Payload)
}

func (o *ReadIssueTemplateOK) GetPayload() *models.APIResultIssueTemplate {
	return o.Payload
}

func (o *ReadIssueTemplateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultIssueTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueTemplateBadRequest creates a ReadIssueTemplateBadRequest with default headers values
func NewReadIssueTemplateBadRequest() *ReadIssueTemplateBadRequest {
	return &ReadIssueTemplateBadRequest{}
}

/*ReadIssueTemplateBadRequest handles this case with default header values.

Bad Request
*/
type ReadIssueTemplateBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadIssueTemplateBadRequest) Error() string {
	return fmt.Sprintf("[GET /issueTemplates/{id}][%d] readIssueTemplateBadRequest  %+v", 400, o.Payload)
}

func (o *ReadIssueTemplateBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueTemplateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueTemplateUnauthorized creates a ReadIssueTemplateUnauthorized with default headers values
func NewReadIssueTemplateUnauthorized() *ReadIssueTemplateUnauthorized {
	return &ReadIssueTemplateUnauthorized{}
}

/*ReadIssueTemplateUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadIssueTemplateUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadIssueTemplateUnauthorized) Error() string {
	return fmt.Sprintf("[GET /issueTemplates/{id}][%d] readIssueTemplateUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadIssueTemplateUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueTemplateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueTemplateForbidden creates a ReadIssueTemplateForbidden with default headers values
func NewReadIssueTemplateForbidden() *ReadIssueTemplateForbidden {
	return &ReadIssueTemplateForbidden{}
}

/*ReadIssueTemplateForbidden handles this case with default header values.

Forbidden
*/
type ReadIssueTemplateForbidden struct {
	Payload *models.APIResult
}

func (o *ReadIssueTemplateForbidden) Error() string {
	return fmt.Sprintf("[GET /issueTemplates/{id}][%d] readIssueTemplateForbidden  %+v", 403, o.Payload)
}

func (o *ReadIssueTemplateForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueTemplateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueTemplateNotFound creates a ReadIssueTemplateNotFound with default headers values
func NewReadIssueTemplateNotFound() *ReadIssueTemplateNotFound {
	return &ReadIssueTemplateNotFound{}
}

/*ReadIssueTemplateNotFound handles this case with default header values.

Not Found
*/
type ReadIssueTemplateNotFound struct {
	Payload *models.APIResult
}

func (o *ReadIssueTemplateNotFound) Error() string {
	return fmt.Sprintf("[GET /issueTemplates/{id}][%d] readIssueTemplateNotFound  %+v", 404, o.Payload)
}

func (o *ReadIssueTemplateNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueTemplateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueTemplateConflict creates a ReadIssueTemplateConflict with default headers values
func NewReadIssueTemplateConflict() *ReadIssueTemplateConflict {
	return &ReadIssueTemplateConflict{}
}

/*ReadIssueTemplateConflict handles this case with default header values.

Conflict
*/
type ReadIssueTemplateConflict struct {
	Payload *models.APIResult
}

func (o *ReadIssueTemplateConflict) Error() string {
	return fmt.Sprintf("[GET /issueTemplates/{id}][%d] readIssueTemplateConflict  %+v", 409, o.Payload)
}

func (o *ReadIssueTemplateConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueTemplateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadIssueTemplateInternalServerError creates a ReadIssueTemplateInternalServerError with default headers values
func NewReadIssueTemplateInternalServerError() *ReadIssueTemplateInternalServerError {
	return &ReadIssueTemplateInternalServerError{}
}

/*ReadIssueTemplateInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadIssueTemplateInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadIssueTemplateInternalServerError) Error() string {
	return fmt.Sprintf("[GET /issueTemplates/{id}][%d] readIssueTemplateInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadIssueTemplateInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadIssueTemplateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
