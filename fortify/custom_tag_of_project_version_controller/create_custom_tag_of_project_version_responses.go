// Code generated by go-swagger; DO NOT EDIT.

package custom_tag_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CreateCustomTagOfProjectVersionReader is a Reader for the CreateCustomTagOfProjectVersion structure.
type CreateCustomTagOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCustomTagOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCustomTagOfProjectVersionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCustomTagOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCustomTagOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCustomTagOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCustomTagOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateCustomTagOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCustomTagOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCustomTagOfProjectVersionCreated creates a CreateCustomTagOfProjectVersionCreated with default headers values
func NewCreateCustomTagOfProjectVersionCreated() *CreateCustomTagOfProjectVersionCreated {
	return &CreateCustomTagOfProjectVersionCreated{}
}

/*CreateCustomTagOfProjectVersionCreated handles this case with default header values.

Created
*/
type CreateCustomTagOfProjectVersionCreated struct {
	Payload *models.APIResultCustomTag
}

func (o *CreateCustomTagOfProjectVersionCreated) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/customTags][%d] createCustomTagOfProjectVersionCreated  %+v", 201, o.Payload)
}

func (o *CreateCustomTagOfProjectVersionCreated) GetPayload() *models.APIResultCustomTag {
	return o.Payload
}

func (o *CreateCustomTagOfProjectVersionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultCustomTag)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomTagOfProjectVersionBadRequest creates a CreateCustomTagOfProjectVersionBadRequest with default headers values
func NewCreateCustomTagOfProjectVersionBadRequest() *CreateCustomTagOfProjectVersionBadRequest {
	return &CreateCustomTagOfProjectVersionBadRequest{}
}

/*CreateCustomTagOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type CreateCustomTagOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateCustomTagOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/customTags][%d] createCustomTagOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCustomTagOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCustomTagOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomTagOfProjectVersionUnauthorized creates a CreateCustomTagOfProjectVersionUnauthorized with default headers values
func NewCreateCustomTagOfProjectVersionUnauthorized() *CreateCustomTagOfProjectVersionUnauthorized {
	return &CreateCustomTagOfProjectVersionUnauthorized{}
}

/*CreateCustomTagOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCustomTagOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateCustomTagOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/customTags][%d] createCustomTagOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCustomTagOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCustomTagOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomTagOfProjectVersionForbidden creates a CreateCustomTagOfProjectVersionForbidden with default headers values
func NewCreateCustomTagOfProjectVersionForbidden() *CreateCustomTagOfProjectVersionForbidden {
	return &CreateCustomTagOfProjectVersionForbidden{}
}

/*CreateCustomTagOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type CreateCustomTagOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *CreateCustomTagOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/customTags][%d] createCustomTagOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *CreateCustomTagOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCustomTagOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomTagOfProjectVersionNotFound creates a CreateCustomTagOfProjectVersionNotFound with default headers values
func NewCreateCustomTagOfProjectVersionNotFound() *CreateCustomTagOfProjectVersionNotFound {
	return &CreateCustomTagOfProjectVersionNotFound{}
}

/*CreateCustomTagOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type CreateCustomTagOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *CreateCustomTagOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/customTags][%d] createCustomTagOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *CreateCustomTagOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCustomTagOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomTagOfProjectVersionConflict creates a CreateCustomTagOfProjectVersionConflict with default headers values
func NewCreateCustomTagOfProjectVersionConflict() *CreateCustomTagOfProjectVersionConflict {
	return &CreateCustomTagOfProjectVersionConflict{}
}

/*CreateCustomTagOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type CreateCustomTagOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *CreateCustomTagOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/customTags][%d] createCustomTagOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *CreateCustomTagOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCustomTagOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCustomTagOfProjectVersionInternalServerError creates a CreateCustomTagOfProjectVersionInternalServerError with default headers values
func NewCreateCustomTagOfProjectVersionInternalServerError() *CreateCustomTagOfProjectVersionInternalServerError {
	return &CreateCustomTagOfProjectVersionInternalServerError{}
}

/*CreateCustomTagOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateCustomTagOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateCustomTagOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/customTags][%d] createCustomTagOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCustomTagOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCustomTagOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}