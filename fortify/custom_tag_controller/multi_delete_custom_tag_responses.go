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

// MultiDeleteCustomTagReader is a Reader for the MultiDeleteCustomTag structure.
type MultiDeleteCustomTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiDeleteCustomTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiDeleteCustomTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMultiDeleteCustomTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMultiDeleteCustomTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMultiDeleteCustomTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMultiDeleteCustomTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMultiDeleteCustomTagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMultiDeleteCustomTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMultiDeleteCustomTagOK creates a MultiDeleteCustomTagOK with default headers values
func NewMultiDeleteCustomTagOK() *MultiDeleteCustomTagOK {
	return &MultiDeleteCustomTagOK{}
}

/*MultiDeleteCustomTagOK handles this case with default header values.

OK
*/
type MultiDeleteCustomTagOK struct {
	Payload *models.APIResultVoid
}

func (o *MultiDeleteCustomTagOK) Error() string {
	return fmt.Sprintf("[DELETE /customTags][%d] multiDeleteCustomTagOK  %+v", 200, o.Payload)
}

func (o *MultiDeleteCustomTagOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *MultiDeleteCustomTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteCustomTagBadRequest creates a MultiDeleteCustomTagBadRequest with default headers values
func NewMultiDeleteCustomTagBadRequest() *MultiDeleteCustomTagBadRequest {
	return &MultiDeleteCustomTagBadRequest{}
}

/*MultiDeleteCustomTagBadRequest handles this case with default header values.

Bad Request
*/
type MultiDeleteCustomTagBadRequest struct {
	Payload *models.APIResult
}

func (o *MultiDeleteCustomTagBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /customTags][%d] multiDeleteCustomTagBadRequest  %+v", 400, o.Payload)
}

func (o *MultiDeleteCustomTagBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteCustomTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteCustomTagUnauthorized creates a MultiDeleteCustomTagUnauthorized with default headers values
func NewMultiDeleteCustomTagUnauthorized() *MultiDeleteCustomTagUnauthorized {
	return &MultiDeleteCustomTagUnauthorized{}
}

/*MultiDeleteCustomTagUnauthorized handles this case with default header values.

Unauthorized
*/
type MultiDeleteCustomTagUnauthorized struct {
	Payload *models.APIResult
}

func (o *MultiDeleteCustomTagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /customTags][%d] multiDeleteCustomTagUnauthorized  %+v", 401, o.Payload)
}

func (o *MultiDeleteCustomTagUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteCustomTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteCustomTagForbidden creates a MultiDeleteCustomTagForbidden with default headers values
func NewMultiDeleteCustomTagForbidden() *MultiDeleteCustomTagForbidden {
	return &MultiDeleteCustomTagForbidden{}
}

/*MultiDeleteCustomTagForbidden handles this case with default header values.

Forbidden
*/
type MultiDeleteCustomTagForbidden struct {
	Payload *models.APIResult
}

func (o *MultiDeleteCustomTagForbidden) Error() string {
	return fmt.Sprintf("[DELETE /customTags][%d] multiDeleteCustomTagForbidden  %+v", 403, o.Payload)
}

func (o *MultiDeleteCustomTagForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteCustomTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteCustomTagNotFound creates a MultiDeleteCustomTagNotFound with default headers values
func NewMultiDeleteCustomTagNotFound() *MultiDeleteCustomTagNotFound {
	return &MultiDeleteCustomTagNotFound{}
}

/*MultiDeleteCustomTagNotFound handles this case with default header values.

Not Found
*/
type MultiDeleteCustomTagNotFound struct {
	Payload *models.APIResult
}

func (o *MultiDeleteCustomTagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /customTags][%d] multiDeleteCustomTagNotFound  %+v", 404, o.Payload)
}

func (o *MultiDeleteCustomTagNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteCustomTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteCustomTagConflict creates a MultiDeleteCustomTagConflict with default headers values
func NewMultiDeleteCustomTagConflict() *MultiDeleteCustomTagConflict {
	return &MultiDeleteCustomTagConflict{}
}

/*MultiDeleteCustomTagConflict handles this case with default header values.

Conflict
*/
type MultiDeleteCustomTagConflict struct {
	Payload *models.APIResult
}

func (o *MultiDeleteCustomTagConflict) Error() string {
	return fmt.Sprintf("[DELETE /customTags][%d] multiDeleteCustomTagConflict  %+v", 409, o.Payload)
}

func (o *MultiDeleteCustomTagConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteCustomTagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteCustomTagInternalServerError creates a MultiDeleteCustomTagInternalServerError with default headers values
func NewMultiDeleteCustomTagInternalServerError() *MultiDeleteCustomTagInternalServerError {
	return &MultiDeleteCustomTagInternalServerError{}
}

/*MultiDeleteCustomTagInternalServerError handles this case with default header values.

Internal Server Error
*/
type MultiDeleteCustomTagInternalServerError struct {
	Payload *models.APIResult
}

func (o *MultiDeleteCustomTagInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /customTags][%d] multiDeleteCustomTagInternalServerError  %+v", 500, o.Payload)
}

func (o *MultiDeleteCustomTagInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteCustomTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
