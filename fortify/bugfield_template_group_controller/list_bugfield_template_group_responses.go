// Code generated by go-swagger; DO NOT EDIT.

package bugfield_template_group_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListBugfieldTemplateGroupReader is a Reader for the ListBugfieldTemplateGroup structure.
type ListBugfieldTemplateGroupReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListBugfieldTemplateGroupReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListBugfieldTemplateGroupOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListBugfieldTemplateGroupBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListBugfieldTemplateGroupUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListBugfieldTemplateGroupForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListBugfieldTemplateGroupNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListBugfieldTemplateGroupConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListBugfieldTemplateGroupInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListBugfieldTemplateGroupOK creates a ListBugfieldTemplateGroupOK with default headers values
func NewListBugfieldTemplateGroupOK() *ListBugfieldTemplateGroupOK {
	return &ListBugfieldTemplateGroupOK{}
}

/*ListBugfieldTemplateGroupOK handles this case with default header values.

OK
*/
type ListBugfieldTemplateGroupOK struct {
	Payload *models.APIResultListBugfieldTemplateGroupDto
}

func (o *ListBugfieldTemplateGroupOK) Error() string {
	return fmt.Sprintf("[GET /bugfieldTemplateGroups][%d] listBugfieldTemplateGroupOK  %+v", 200, o.Payload)
}

func (o *ListBugfieldTemplateGroupOK) GetPayload() *models.APIResultListBugfieldTemplateGroupDto {
	return o.Payload
}

func (o *ListBugfieldTemplateGroupOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListBugfieldTemplateGroupDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBugfieldTemplateGroupBadRequest creates a ListBugfieldTemplateGroupBadRequest with default headers values
func NewListBugfieldTemplateGroupBadRequest() *ListBugfieldTemplateGroupBadRequest {
	return &ListBugfieldTemplateGroupBadRequest{}
}

/*ListBugfieldTemplateGroupBadRequest handles this case with default header values.

Bad Request
*/
type ListBugfieldTemplateGroupBadRequest struct {
	Payload *models.APIResult
}

func (o *ListBugfieldTemplateGroupBadRequest) Error() string {
	return fmt.Sprintf("[GET /bugfieldTemplateGroups][%d] listBugfieldTemplateGroupBadRequest  %+v", 400, o.Payload)
}

func (o *ListBugfieldTemplateGroupBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListBugfieldTemplateGroupBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBugfieldTemplateGroupUnauthorized creates a ListBugfieldTemplateGroupUnauthorized with default headers values
func NewListBugfieldTemplateGroupUnauthorized() *ListBugfieldTemplateGroupUnauthorized {
	return &ListBugfieldTemplateGroupUnauthorized{}
}

/*ListBugfieldTemplateGroupUnauthorized handles this case with default header values.

Unauthorized
*/
type ListBugfieldTemplateGroupUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListBugfieldTemplateGroupUnauthorized) Error() string {
	return fmt.Sprintf("[GET /bugfieldTemplateGroups][%d] listBugfieldTemplateGroupUnauthorized  %+v", 401, o.Payload)
}

func (o *ListBugfieldTemplateGroupUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListBugfieldTemplateGroupUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBugfieldTemplateGroupForbidden creates a ListBugfieldTemplateGroupForbidden with default headers values
func NewListBugfieldTemplateGroupForbidden() *ListBugfieldTemplateGroupForbidden {
	return &ListBugfieldTemplateGroupForbidden{}
}

/*ListBugfieldTemplateGroupForbidden handles this case with default header values.

Forbidden
*/
type ListBugfieldTemplateGroupForbidden struct {
	Payload *models.APIResult
}

func (o *ListBugfieldTemplateGroupForbidden) Error() string {
	return fmt.Sprintf("[GET /bugfieldTemplateGroups][%d] listBugfieldTemplateGroupForbidden  %+v", 403, o.Payload)
}

func (o *ListBugfieldTemplateGroupForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListBugfieldTemplateGroupForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBugfieldTemplateGroupNotFound creates a ListBugfieldTemplateGroupNotFound with default headers values
func NewListBugfieldTemplateGroupNotFound() *ListBugfieldTemplateGroupNotFound {
	return &ListBugfieldTemplateGroupNotFound{}
}

/*ListBugfieldTemplateGroupNotFound handles this case with default header values.

Not Found
*/
type ListBugfieldTemplateGroupNotFound struct {
	Payload *models.APIResult
}

func (o *ListBugfieldTemplateGroupNotFound) Error() string {
	return fmt.Sprintf("[GET /bugfieldTemplateGroups][%d] listBugfieldTemplateGroupNotFound  %+v", 404, o.Payload)
}

func (o *ListBugfieldTemplateGroupNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListBugfieldTemplateGroupNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBugfieldTemplateGroupConflict creates a ListBugfieldTemplateGroupConflict with default headers values
func NewListBugfieldTemplateGroupConflict() *ListBugfieldTemplateGroupConflict {
	return &ListBugfieldTemplateGroupConflict{}
}

/*ListBugfieldTemplateGroupConflict handles this case with default header values.

Conflict
*/
type ListBugfieldTemplateGroupConflict struct {
	Payload *models.APIResult
}

func (o *ListBugfieldTemplateGroupConflict) Error() string {
	return fmt.Sprintf("[GET /bugfieldTemplateGroups][%d] listBugfieldTemplateGroupConflict  %+v", 409, o.Payload)
}

func (o *ListBugfieldTemplateGroupConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListBugfieldTemplateGroupConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListBugfieldTemplateGroupInternalServerError creates a ListBugfieldTemplateGroupInternalServerError with default headers values
func NewListBugfieldTemplateGroupInternalServerError() *ListBugfieldTemplateGroupInternalServerError {
	return &ListBugfieldTemplateGroupInternalServerError{}
}

/*ListBugfieldTemplateGroupInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListBugfieldTemplateGroupInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListBugfieldTemplateGroupInternalServerError) Error() string {
	return fmt.Sprintf("[GET /bugfieldTemplateGroups][%d] listBugfieldTemplateGroupInternalServerError  %+v", 500, o.Payload)
}

func (o *ListBugfieldTemplateGroupInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListBugfieldTemplateGroupInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
