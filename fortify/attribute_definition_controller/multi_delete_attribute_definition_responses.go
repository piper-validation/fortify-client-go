// Code generated by go-swagger; DO NOT EDIT.

package attribute_definition_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// MultiDeleteAttributeDefinitionReader is a Reader for the MultiDeleteAttributeDefinition structure.
type MultiDeleteAttributeDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiDeleteAttributeDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiDeleteAttributeDefinitionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMultiDeleteAttributeDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMultiDeleteAttributeDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMultiDeleteAttributeDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMultiDeleteAttributeDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMultiDeleteAttributeDefinitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMultiDeleteAttributeDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMultiDeleteAttributeDefinitionOK creates a MultiDeleteAttributeDefinitionOK with default headers values
func NewMultiDeleteAttributeDefinitionOK() *MultiDeleteAttributeDefinitionOK {
	return &MultiDeleteAttributeDefinitionOK{}
}

/*MultiDeleteAttributeDefinitionOK handles this case with default header values.

OK
*/
type MultiDeleteAttributeDefinitionOK struct {
	Payload *models.APIResultVoid
}

func (o *MultiDeleteAttributeDefinitionOK) Error() string {
	return fmt.Sprintf("[DELETE /attributeDefinitions][%d] multiDeleteAttributeDefinitionOK  %+v", 200, o.Payload)
}

func (o *MultiDeleteAttributeDefinitionOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *MultiDeleteAttributeDefinitionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAttributeDefinitionBadRequest creates a MultiDeleteAttributeDefinitionBadRequest with default headers values
func NewMultiDeleteAttributeDefinitionBadRequest() *MultiDeleteAttributeDefinitionBadRequest {
	return &MultiDeleteAttributeDefinitionBadRequest{}
}

/*MultiDeleteAttributeDefinitionBadRequest handles this case with default header values.

Bad Request
*/
type MultiDeleteAttributeDefinitionBadRequest struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAttributeDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /attributeDefinitions][%d] multiDeleteAttributeDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *MultiDeleteAttributeDefinitionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAttributeDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAttributeDefinitionUnauthorized creates a MultiDeleteAttributeDefinitionUnauthorized with default headers values
func NewMultiDeleteAttributeDefinitionUnauthorized() *MultiDeleteAttributeDefinitionUnauthorized {
	return &MultiDeleteAttributeDefinitionUnauthorized{}
}

/*MultiDeleteAttributeDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type MultiDeleteAttributeDefinitionUnauthorized struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAttributeDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /attributeDefinitions][%d] multiDeleteAttributeDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *MultiDeleteAttributeDefinitionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAttributeDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAttributeDefinitionForbidden creates a MultiDeleteAttributeDefinitionForbidden with default headers values
func NewMultiDeleteAttributeDefinitionForbidden() *MultiDeleteAttributeDefinitionForbidden {
	return &MultiDeleteAttributeDefinitionForbidden{}
}

/*MultiDeleteAttributeDefinitionForbidden handles this case with default header values.

Forbidden
*/
type MultiDeleteAttributeDefinitionForbidden struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAttributeDefinitionForbidden) Error() string {
	return fmt.Sprintf("[DELETE /attributeDefinitions][%d] multiDeleteAttributeDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *MultiDeleteAttributeDefinitionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAttributeDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAttributeDefinitionNotFound creates a MultiDeleteAttributeDefinitionNotFound with default headers values
func NewMultiDeleteAttributeDefinitionNotFound() *MultiDeleteAttributeDefinitionNotFound {
	return &MultiDeleteAttributeDefinitionNotFound{}
}

/*MultiDeleteAttributeDefinitionNotFound handles this case with default header values.

Not Found
*/
type MultiDeleteAttributeDefinitionNotFound struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAttributeDefinitionNotFound) Error() string {
	return fmt.Sprintf("[DELETE /attributeDefinitions][%d] multiDeleteAttributeDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *MultiDeleteAttributeDefinitionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAttributeDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAttributeDefinitionConflict creates a MultiDeleteAttributeDefinitionConflict with default headers values
func NewMultiDeleteAttributeDefinitionConflict() *MultiDeleteAttributeDefinitionConflict {
	return &MultiDeleteAttributeDefinitionConflict{}
}

/*MultiDeleteAttributeDefinitionConflict handles this case with default header values.

Conflict
*/
type MultiDeleteAttributeDefinitionConflict struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAttributeDefinitionConflict) Error() string {
	return fmt.Sprintf("[DELETE /attributeDefinitions][%d] multiDeleteAttributeDefinitionConflict  %+v", 409, o.Payload)
}

func (o *MultiDeleteAttributeDefinitionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAttributeDefinitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteAttributeDefinitionInternalServerError creates a MultiDeleteAttributeDefinitionInternalServerError with default headers values
func NewMultiDeleteAttributeDefinitionInternalServerError() *MultiDeleteAttributeDefinitionInternalServerError {
	return &MultiDeleteAttributeDefinitionInternalServerError{}
}

/*MultiDeleteAttributeDefinitionInternalServerError handles this case with default header values.

Internal Server Error
*/
type MultiDeleteAttributeDefinitionInternalServerError struct {
	Payload *models.APIResult
}

func (o *MultiDeleteAttributeDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /attributeDefinitions][%d] multiDeleteAttributeDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *MultiDeleteAttributeDefinitionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteAttributeDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
