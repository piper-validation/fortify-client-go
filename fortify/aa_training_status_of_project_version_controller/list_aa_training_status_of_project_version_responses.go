// Code generated by go-swagger; DO NOT EDIT.

package aa_training_status_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListAaTrainingStatusOfProjectVersionReader is a Reader for the ListAaTrainingStatusOfProjectVersion structure.
type ListAaTrainingStatusOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAaTrainingStatusOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAaTrainingStatusOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAaTrainingStatusOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAaTrainingStatusOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAaTrainingStatusOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListAaTrainingStatusOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListAaTrainingStatusOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAaTrainingStatusOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAaTrainingStatusOfProjectVersionOK creates a ListAaTrainingStatusOfProjectVersionOK with default headers values
func NewListAaTrainingStatusOfProjectVersionOK() *ListAaTrainingStatusOfProjectVersionOK {
	return &ListAaTrainingStatusOfProjectVersionOK{}
}

/*ListAaTrainingStatusOfProjectVersionOK handles this case with default header values.

OK
*/
type ListAaTrainingStatusOfProjectVersionOK struct {
	Payload *models.APIResultListAATrainingStatus
}

func (o *ListAaTrainingStatusOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/auditAssistantTrainingStatus][%d] listAaTrainingStatusOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ListAaTrainingStatusOfProjectVersionOK) GetPayload() *models.APIResultListAATrainingStatus {
	return o.Payload
}

func (o *ListAaTrainingStatusOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListAATrainingStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAaTrainingStatusOfProjectVersionBadRequest creates a ListAaTrainingStatusOfProjectVersionBadRequest with default headers values
func NewListAaTrainingStatusOfProjectVersionBadRequest() *ListAaTrainingStatusOfProjectVersionBadRequest {
	return &ListAaTrainingStatusOfProjectVersionBadRequest{}
}

/*ListAaTrainingStatusOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ListAaTrainingStatusOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ListAaTrainingStatusOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/auditAssistantTrainingStatus][%d] listAaTrainingStatusOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ListAaTrainingStatusOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAaTrainingStatusOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAaTrainingStatusOfProjectVersionUnauthorized creates a ListAaTrainingStatusOfProjectVersionUnauthorized with default headers values
func NewListAaTrainingStatusOfProjectVersionUnauthorized() *ListAaTrainingStatusOfProjectVersionUnauthorized {
	return &ListAaTrainingStatusOfProjectVersionUnauthorized{}
}

/*ListAaTrainingStatusOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAaTrainingStatusOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListAaTrainingStatusOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/auditAssistantTrainingStatus][%d] listAaTrainingStatusOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAaTrainingStatusOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAaTrainingStatusOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAaTrainingStatusOfProjectVersionForbidden creates a ListAaTrainingStatusOfProjectVersionForbidden with default headers values
func NewListAaTrainingStatusOfProjectVersionForbidden() *ListAaTrainingStatusOfProjectVersionForbidden {
	return &ListAaTrainingStatusOfProjectVersionForbidden{}
}

/*ListAaTrainingStatusOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ListAaTrainingStatusOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ListAaTrainingStatusOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/auditAssistantTrainingStatus][%d] listAaTrainingStatusOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ListAaTrainingStatusOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAaTrainingStatusOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAaTrainingStatusOfProjectVersionNotFound creates a ListAaTrainingStatusOfProjectVersionNotFound with default headers values
func NewListAaTrainingStatusOfProjectVersionNotFound() *ListAaTrainingStatusOfProjectVersionNotFound {
	return &ListAaTrainingStatusOfProjectVersionNotFound{}
}

/*ListAaTrainingStatusOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ListAaTrainingStatusOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ListAaTrainingStatusOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/auditAssistantTrainingStatus][%d] listAaTrainingStatusOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ListAaTrainingStatusOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAaTrainingStatusOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAaTrainingStatusOfProjectVersionConflict creates a ListAaTrainingStatusOfProjectVersionConflict with default headers values
func NewListAaTrainingStatusOfProjectVersionConflict() *ListAaTrainingStatusOfProjectVersionConflict {
	return &ListAaTrainingStatusOfProjectVersionConflict{}
}

/*ListAaTrainingStatusOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ListAaTrainingStatusOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ListAaTrainingStatusOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/auditAssistantTrainingStatus][%d] listAaTrainingStatusOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ListAaTrainingStatusOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAaTrainingStatusOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAaTrainingStatusOfProjectVersionInternalServerError creates a ListAaTrainingStatusOfProjectVersionInternalServerError with default headers values
func NewListAaTrainingStatusOfProjectVersionInternalServerError() *ListAaTrainingStatusOfProjectVersionInternalServerError {
	return &ListAaTrainingStatusOfProjectVersionInternalServerError{}
}

/*ListAaTrainingStatusOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAaTrainingStatusOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListAaTrainingStatusOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/auditAssistantTrainingStatus][%d] listAaTrainingStatusOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAaTrainingStatusOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAaTrainingStatusOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}