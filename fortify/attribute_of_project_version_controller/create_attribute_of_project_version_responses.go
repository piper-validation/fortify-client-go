// Code generated by go-swagger; DO NOT EDIT.

package attribute_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CreateAttributeOfProjectVersionReader is a Reader for the CreateAttributeOfProjectVersion structure.
type CreateAttributeOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateAttributeOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateAttributeOfProjectVersionCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateAttributeOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateAttributeOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateAttributeOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateAttributeOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateAttributeOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateAttributeOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateAttributeOfProjectVersionCreated creates a CreateAttributeOfProjectVersionCreated with default headers values
func NewCreateAttributeOfProjectVersionCreated() *CreateAttributeOfProjectVersionCreated {
	return &CreateAttributeOfProjectVersionCreated{}
}

/*CreateAttributeOfProjectVersionCreated handles this case with default header values.

Created
*/
type CreateAttributeOfProjectVersionCreated struct {
	Payload *models.APIResultAttribute
}

func (o *CreateAttributeOfProjectVersionCreated) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/attributes][%d] createAttributeOfProjectVersionCreated  %+v", 201, o.Payload)
}

func (o *CreateAttributeOfProjectVersionCreated) GetPayload() *models.APIResultAttribute {
	return o.Payload
}

func (o *CreateAttributeOfProjectVersionCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultAttribute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeOfProjectVersionBadRequest creates a CreateAttributeOfProjectVersionBadRequest with default headers values
func NewCreateAttributeOfProjectVersionBadRequest() *CreateAttributeOfProjectVersionBadRequest {
	return &CreateAttributeOfProjectVersionBadRequest{}
}

/*CreateAttributeOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type CreateAttributeOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateAttributeOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/attributes][%d] createAttributeOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *CreateAttributeOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeOfProjectVersionUnauthorized creates a CreateAttributeOfProjectVersionUnauthorized with default headers values
func NewCreateAttributeOfProjectVersionUnauthorized() *CreateAttributeOfProjectVersionUnauthorized {
	return &CreateAttributeOfProjectVersionUnauthorized{}
}

/*CreateAttributeOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateAttributeOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateAttributeOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/attributes][%d] createAttributeOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateAttributeOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeOfProjectVersionForbidden creates a CreateAttributeOfProjectVersionForbidden with default headers values
func NewCreateAttributeOfProjectVersionForbidden() *CreateAttributeOfProjectVersionForbidden {
	return &CreateAttributeOfProjectVersionForbidden{}
}

/*CreateAttributeOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type CreateAttributeOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *CreateAttributeOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/attributes][%d] createAttributeOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *CreateAttributeOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeOfProjectVersionNotFound creates a CreateAttributeOfProjectVersionNotFound with default headers values
func NewCreateAttributeOfProjectVersionNotFound() *CreateAttributeOfProjectVersionNotFound {
	return &CreateAttributeOfProjectVersionNotFound{}
}

/*CreateAttributeOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type CreateAttributeOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *CreateAttributeOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/attributes][%d] createAttributeOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *CreateAttributeOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeOfProjectVersionConflict creates a CreateAttributeOfProjectVersionConflict with default headers values
func NewCreateAttributeOfProjectVersionConflict() *CreateAttributeOfProjectVersionConflict {
	return &CreateAttributeOfProjectVersionConflict{}
}

/*CreateAttributeOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type CreateAttributeOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *CreateAttributeOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/attributes][%d] createAttributeOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *CreateAttributeOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateAttributeOfProjectVersionInternalServerError creates a CreateAttributeOfProjectVersionInternalServerError with default headers values
func NewCreateAttributeOfProjectVersionInternalServerError() *CreateAttributeOfProjectVersionInternalServerError {
	return &CreateAttributeOfProjectVersionInternalServerError{}
}

/*CreateAttributeOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateAttributeOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateAttributeOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projectVersions/{parentId}/attributes][%d] createAttributeOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateAttributeOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateAttributeOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}