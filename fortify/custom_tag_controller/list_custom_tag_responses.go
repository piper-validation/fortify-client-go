// Code generated by go-swagger; DO NOT EDIT.

package custom_tag_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListCustomTagReader is a Reader for the ListCustomTag structure.
type ListCustomTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCustomTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCustomTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCustomTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCustomTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCustomTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCustomTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListCustomTagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCustomTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCustomTagOK creates a ListCustomTagOK with default headers values
func NewListCustomTagOK() *ListCustomTagOK {
	return &ListCustomTagOK{}
}

/*ListCustomTagOK handles this case with default header values.

OK
*/
type ListCustomTagOK struct {
	Payload *models.APIResultListCustomTag
}

func (o *ListCustomTagOK) Error() string {
	return fmt.Sprintf("[GET /customTags][%d] listCustomTagOK  %+v", 200, o.Payload)
}

func (o *ListCustomTagOK) GetPayload() *models.APIResultListCustomTag {
	return o.Payload
}

func (o *ListCustomTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListCustomTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCustomTagBadRequest creates a ListCustomTagBadRequest with default headers values
func NewListCustomTagBadRequest() *ListCustomTagBadRequest {
	return &ListCustomTagBadRequest{}
}

/*ListCustomTagBadRequest handles this case with default header values.

Bad Request
*/
type ListCustomTagBadRequest struct {
	Payload *models.APIResult
}

func (o *ListCustomTagBadRequest) Error() string {
	return fmt.Sprintf("[GET /customTags][%d] listCustomTagBadRequest  %+v", 400, o.Payload)
}

func (o *ListCustomTagBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCustomTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCustomTagUnauthorized creates a ListCustomTagUnauthorized with default headers values
func NewListCustomTagUnauthorized() *ListCustomTagUnauthorized {
	return &ListCustomTagUnauthorized{}
}

/*ListCustomTagUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCustomTagUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListCustomTagUnauthorized) Error() string {
	return fmt.Sprintf("[GET /customTags][%d] listCustomTagUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCustomTagUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCustomTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCustomTagForbidden creates a ListCustomTagForbidden with default headers values
func NewListCustomTagForbidden() *ListCustomTagForbidden {
	return &ListCustomTagForbidden{}
}

/*ListCustomTagForbidden handles this case with default header values.

Forbidden
*/
type ListCustomTagForbidden struct {
	Payload *models.APIResult
}

func (o *ListCustomTagForbidden) Error() string {
	return fmt.Sprintf("[GET /customTags][%d] listCustomTagForbidden  %+v", 403, o.Payload)
}

func (o *ListCustomTagForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCustomTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCustomTagNotFound creates a ListCustomTagNotFound with default headers values
func NewListCustomTagNotFound() *ListCustomTagNotFound {
	return &ListCustomTagNotFound{}
}

/*ListCustomTagNotFound handles this case with default header values.

Not Found
*/
type ListCustomTagNotFound struct {
	Payload *models.APIResult
}

func (o *ListCustomTagNotFound) Error() string {
	return fmt.Sprintf("[GET /customTags][%d] listCustomTagNotFound  %+v", 404, o.Payload)
}

func (o *ListCustomTagNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCustomTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCustomTagConflict creates a ListCustomTagConflict with default headers values
func NewListCustomTagConflict() *ListCustomTagConflict {
	return &ListCustomTagConflict{}
}

/*ListCustomTagConflict handles this case with default header values.

Conflict
*/
type ListCustomTagConflict struct {
	Payload *models.APIResult
}

func (o *ListCustomTagConflict) Error() string {
	return fmt.Sprintf("[GET /customTags][%d] listCustomTagConflict  %+v", 409, o.Payload)
}

func (o *ListCustomTagConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCustomTagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCustomTagInternalServerError creates a ListCustomTagInternalServerError with default headers values
func NewListCustomTagInternalServerError() *ListCustomTagInternalServerError {
	return &ListCustomTagInternalServerError{}
}

/*ListCustomTagInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListCustomTagInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListCustomTagInternalServerError) Error() string {
	return fmt.Sprintf("[GET /customTags][%d] listCustomTagInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCustomTagInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCustomTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
