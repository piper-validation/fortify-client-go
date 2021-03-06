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

// CreateAttributeDefinitionReader is a Reader for the CreateAttributeDefinition structure.
type CreateAttributeDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAttributeDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAttributeDefinitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAttributeDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAttributeDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAttributeDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAttributeDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAttributeDefinitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAttributeDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAttributeDefinitionCreated creates a CreateAttributeDefinitionCreated with default headers values
func NewCreateAttributeDefinitionCreated() *CreateAttributeDefinitionCreated {
	return &CreateAttributeDefinitionCreated{}
}

/*CreateAttributeDefinitionCreated handles this case with default header values.

Created
*/
type CreateAttributeDefinitionCreated struct {
	Payload *models.APIResultAttributeDefinition
}

func (o *CreateAttributeDefinitionCreated) Error() string {
	return fmt.Sprintf("[POST /attributeDefinitions][%d] createAttributeDefinitionCreated  %+v", 201, o.Payload)
}

func (o *CreateAttributeDefinitionCreated) GetPayload() *models.APIResultAttributeDefinition {
	return o.Payload
}

func (o *CreateAttributeDefinitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultAttributeDefinition)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeDefinitionBadRequest creates a CreateAttributeDefinitionBadRequest with default headers values
func NewCreateAttributeDefinitionBadRequest() *CreateAttributeDefinitionBadRequest {
	return &CreateAttributeDefinitionBadRequest{}
}

/*CreateAttributeDefinitionBadRequest handles this case with default header values.

Bad Request
*/
type CreateAttributeDefinitionBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateAttributeDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[POST /attributeDefinitions][%d] createAttributeDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAttributeDefinitionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeDefinitionUnauthorized creates a CreateAttributeDefinitionUnauthorized with default headers values
func NewCreateAttributeDefinitionUnauthorized() *CreateAttributeDefinitionUnauthorized {
	return &CreateAttributeDefinitionUnauthorized{}
}

/*CreateAttributeDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAttributeDefinitionUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateAttributeDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /attributeDefinitions][%d] createAttributeDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAttributeDefinitionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeDefinitionForbidden creates a CreateAttributeDefinitionForbidden with default headers values
func NewCreateAttributeDefinitionForbidden() *CreateAttributeDefinitionForbidden {
	return &CreateAttributeDefinitionForbidden{}
}

/*CreateAttributeDefinitionForbidden handles this case with default header values.

Forbidden
*/
type CreateAttributeDefinitionForbidden struct {
	Payload *models.APIResult
}

func (o *CreateAttributeDefinitionForbidden) Error() string {
	return fmt.Sprintf("[POST /attributeDefinitions][%d] createAttributeDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *CreateAttributeDefinitionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeDefinitionNotFound creates a CreateAttributeDefinitionNotFound with default headers values
func NewCreateAttributeDefinitionNotFound() *CreateAttributeDefinitionNotFound {
	return &CreateAttributeDefinitionNotFound{}
}

/*CreateAttributeDefinitionNotFound handles this case with default header values.

Not Found
*/
type CreateAttributeDefinitionNotFound struct {
	Payload *models.APIResult
}

func (o *CreateAttributeDefinitionNotFound) Error() string {
	return fmt.Sprintf("[POST /attributeDefinitions][%d] createAttributeDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *CreateAttributeDefinitionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeDefinitionConflict creates a CreateAttributeDefinitionConflict with default headers values
func NewCreateAttributeDefinitionConflict() *CreateAttributeDefinitionConflict {
	return &CreateAttributeDefinitionConflict{}
}

/*CreateAttributeDefinitionConflict handles this case with default header values.

Conflict
*/
type CreateAttributeDefinitionConflict struct {
	Payload *models.APIResult
}

func (o *CreateAttributeDefinitionConflict) Error() string {
	return fmt.Sprintf("[POST /attributeDefinitions][%d] createAttributeDefinitionConflict  %+v", 409, o.Payload)
}

func (o *CreateAttributeDefinitionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeDefinitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeDefinitionInternalServerError creates a CreateAttributeDefinitionInternalServerError with default headers values
func NewCreateAttributeDefinitionInternalServerError() *CreateAttributeDefinitionInternalServerError {
	return &CreateAttributeDefinitionInternalServerError{}
}

/*CreateAttributeDefinitionInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateAttributeDefinitionInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateAttributeDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /attributeDefinitions][%d] createAttributeDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAttributeDefinitionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
