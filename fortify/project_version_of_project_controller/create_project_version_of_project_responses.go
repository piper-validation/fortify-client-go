// Code generated by go-swagger; DO NOT EDIT.

package project_version_of_project_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CreateProjectVersionOfProjectReader is a Reader for the CreateProjectVersionOfProject structure.
type CreateProjectVersionOfProjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateProjectVersionOfProjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateProjectVersionOfProjectCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateProjectVersionOfProjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateProjectVersionOfProjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateProjectVersionOfProjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateProjectVersionOfProjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateProjectVersionOfProjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateProjectVersionOfProjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateProjectVersionOfProjectCreated creates a CreateProjectVersionOfProjectCreated with default headers values
func NewCreateProjectVersionOfProjectCreated() *CreateProjectVersionOfProjectCreated {
	return &CreateProjectVersionOfProjectCreated{}
}

/*CreateProjectVersionOfProjectCreated handles this case with default header values.

Created
*/
type CreateProjectVersionOfProjectCreated struct {
	Payload *models.APIResultProjectVersion
}

func (o *CreateProjectVersionOfProjectCreated) Error() string {
	return fmt.Sprintf("[POST /projects/{parentId}/versions][%d] createProjectVersionOfProjectCreated  %+v", 201, o.Payload)
}

func (o *CreateProjectVersionOfProjectCreated) GetPayload() *models.APIResultProjectVersion {
	return o.Payload
}

func (o *CreateProjectVersionOfProjectCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultProjectVersion)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectVersionOfProjectBadRequest creates a CreateProjectVersionOfProjectBadRequest with default headers values
func NewCreateProjectVersionOfProjectBadRequest() *CreateProjectVersionOfProjectBadRequest {
	return &CreateProjectVersionOfProjectBadRequest{}
}

/*CreateProjectVersionOfProjectBadRequest handles this case with default header values.

Bad Request
*/
type CreateProjectVersionOfProjectBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateProjectVersionOfProjectBadRequest) Error() string {
	return fmt.Sprintf("[POST /projects/{parentId}/versions][%d] createProjectVersionOfProjectBadRequest  %+v", 400, o.Payload)
}

func (o *CreateProjectVersionOfProjectBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateProjectVersionOfProjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectVersionOfProjectUnauthorized creates a CreateProjectVersionOfProjectUnauthorized with default headers values
func NewCreateProjectVersionOfProjectUnauthorized() *CreateProjectVersionOfProjectUnauthorized {
	return &CreateProjectVersionOfProjectUnauthorized{}
}

/*CreateProjectVersionOfProjectUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateProjectVersionOfProjectUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateProjectVersionOfProjectUnauthorized) Error() string {
	return fmt.Sprintf("[POST /projects/{parentId}/versions][%d] createProjectVersionOfProjectUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateProjectVersionOfProjectUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateProjectVersionOfProjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectVersionOfProjectForbidden creates a CreateProjectVersionOfProjectForbidden with default headers values
func NewCreateProjectVersionOfProjectForbidden() *CreateProjectVersionOfProjectForbidden {
	return &CreateProjectVersionOfProjectForbidden{}
}

/*CreateProjectVersionOfProjectForbidden handles this case with default header values.

Forbidden
*/
type CreateProjectVersionOfProjectForbidden struct {
	Payload *models.APIResult
}

func (o *CreateProjectVersionOfProjectForbidden) Error() string {
	return fmt.Sprintf("[POST /projects/{parentId}/versions][%d] createProjectVersionOfProjectForbidden  %+v", 403, o.Payload)
}

func (o *CreateProjectVersionOfProjectForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateProjectVersionOfProjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectVersionOfProjectNotFound creates a CreateProjectVersionOfProjectNotFound with default headers values
func NewCreateProjectVersionOfProjectNotFound() *CreateProjectVersionOfProjectNotFound {
	return &CreateProjectVersionOfProjectNotFound{}
}

/*CreateProjectVersionOfProjectNotFound handles this case with default header values.

Not Found
*/
type CreateProjectVersionOfProjectNotFound struct {
	Payload *models.APIResult
}

func (o *CreateProjectVersionOfProjectNotFound) Error() string {
	return fmt.Sprintf("[POST /projects/{parentId}/versions][%d] createProjectVersionOfProjectNotFound  %+v", 404, o.Payload)
}

func (o *CreateProjectVersionOfProjectNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateProjectVersionOfProjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectVersionOfProjectConflict creates a CreateProjectVersionOfProjectConflict with default headers values
func NewCreateProjectVersionOfProjectConflict() *CreateProjectVersionOfProjectConflict {
	return &CreateProjectVersionOfProjectConflict{}
}

/*CreateProjectVersionOfProjectConflict handles this case with default header values.

Conflict
*/
type CreateProjectVersionOfProjectConflict struct {
	Payload *models.APIResult
}

func (o *CreateProjectVersionOfProjectConflict) Error() string {
	return fmt.Sprintf("[POST /projects/{parentId}/versions][%d] createProjectVersionOfProjectConflict  %+v", 409, o.Payload)
}

func (o *CreateProjectVersionOfProjectConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateProjectVersionOfProjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateProjectVersionOfProjectInternalServerError creates a CreateProjectVersionOfProjectInternalServerError with default headers values
func NewCreateProjectVersionOfProjectInternalServerError() *CreateProjectVersionOfProjectInternalServerError {
	return &CreateProjectVersionOfProjectInternalServerError{}
}

/*CreateProjectVersionOfProjectInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateProjectVersionOfProjectInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateProjectVersionOfProjectInternalServerError) Error() string {
	return fmt.Sprintf("[POST /projects/{parentId}/versions][%d] createProjectVersionOfProjectInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateProjectVersionOfProjectInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateProjectVersionOfProjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
