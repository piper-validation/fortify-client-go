// Code generated by go-swagger; DO NOT EDIT.

package alert_definition_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CreateAlertDefinitionReader is a Reader for the CreateAlertDefinition structure.
type CreateAlertDefinitionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAlertDefinitionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAlertDefinitionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAlertDefinitionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAlertDefinitionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAlertDefinitionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAlertDefinitionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAlertDefinitionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAlertDefinitionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAlertDefinitionCreated creates a CreateAlertDefinitionCreated with default headers values
func NewCreateAlertDefinitionCreated() *CreateAlertDefinitionCreated {
	return &CreateAlertDefinitionCreated{}
}

/*CreateAlertDefinitionCreated handles this case with default header values.

Created
*/
type CreateAlertDefinitionCreated struct {
	Payload *models.APIResultAlertDefinitionDto
}

func (o *CreateAlertDefinitionCreated) Error() string {
	return fmt.Sprintf("[POST /alertDefinitions][%d] createAlertDefinitionCreated  %+v", 201, o.Payload)
}

func (o *CreateAlertDefinitionCreated) GetPayload() *models.APIResultAlertDefinitionDto {
	return o.Payload
}

func (o *CreateAlertDefinitionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultAlertDefinitionDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAlertDefinitionBadRequest creates a CreateAlertDefinitionBadRequest with default headers values
func NewCreateAlertDefinitionBadRequest() *CreateAlertDefinitionBadRequest {
	return &CreateAlertDefinitionBadRequest{}
}

/*CreateAlertDefinitionBadRequest handles this case with default header values.

Bad Request
*/
type CreateAlertDefinitionBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateAlertDefinitionBadRequest) Error() string {
	return fmt.Sprintf("[POST /alertDefinitions][%d] createAlertDefinitionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAlertDefinitionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAlertDefinitionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAlertDefinitionUnauthorized creates a CreateAlertDefinitionUnauthorized with default headers values
func NewCreateAlertDefinitionUnauthorized() *CreateAlertDefinitionUnauthorized {
	return &CreateAlertDefinitionUnauthorized{}
}

/*CreateAlertDefinitionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAlertDefinitionUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateAlertDefinitionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /alertDefinitions][%d] createAlertDefinitionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAlertDefinitionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAlertDefinitionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAlertDefinitionForbidden creates a CreateAlertDefinitionForbidden with default headers values
func NewCreateAlertDefinitionForbidden() *CreateAlertDefinitionForbidden {
	return &CreateAlertDefinitionForbidden{}
}

/*CreateAlertDefinitionForbidden handles this case with default header values.

Forbidden
*/
type CreateAlertDefinitionForbidden struct {
	Payload *models.APIResult
}

func (o *CreateAlertDefinitionForbidden) Error() string {
	return fmt.Sprintf("[POST /alertDefinitions][%d] createAlertDefinitionForbidden  %+v", 403, o.Payload)
}

func (o *CreateAlertDefinitionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAlertDefinitionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAlertDefinitionNotFound creates a CreateAlertDefinitionNotFound with default headers values
func NewCreateAlertDefinitionNotFound() *CreateAlertDefinitionNotFound {
	return &CreateAlertDefinitionNotFound{}
}

/*CreateAlertDefinitionNotFound handles this case with default header values.

Not Found
*/
type CreateAlertDefinitionNotFound struct {
	Payload *models.APIResult
}

func (o *CreateAlertDefinitionNotFound) Error() string {
	return fmt.Sprintf("[POST /alertDefinitions][%d] createAlertDefinitionNotFound  %+v", 404, o.Payload)
}

func (o *CreateAlertDefinitionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAlertDefinitionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAlertDefinitionConflict creates a CreateAlertDefinitionConflict with default headers values
func NewCreateAlertDefinitionConflict() *CreateAlertDefinitionConflict {
	return &CreateAlertDefinitionConflict{}
}

/*CreateAlertDefinitionConflict handles this case with default header values.

Conflict
*/
type CreateAlertDefinitionConflict struct {
	Payload *models.APIResult
}

func (o *CreateAlertDefinitionConflict) Error() string {
	return fmt.Sprintf("[POST /alertDefinitions][%d] createAlertDefinitionConflict  %+v", 409, o.Payload)
}

func (o *CreateAlertDefinitionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAlertDefinitionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAlertDefinitionInternalServerError creates a CreateAlertDefinitionInternalServerError with default headers values
func NewCreateAlertDefinitionInternalServerError() *CreateAlertDefinitionInternalServerError {
	return &CreateAlertDefinitionInternalServerError{}
}

/*CreateAlertDefinitionInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateAlertDefinitionInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateAlertDefinitionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /alertDefinitions][%d] createAlertDefinitionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAlertDefinitionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAlertDefinitionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
