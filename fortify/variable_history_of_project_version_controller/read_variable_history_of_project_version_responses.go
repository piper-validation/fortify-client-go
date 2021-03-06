// Code generated by go-swagger; DO NOT EDIT.

package variable_history_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadVariableHistoryOfProjectVersionReader is a Reader for the ReadVariableHistoryOfProjectVersion structure.
type ReadVariableHistoryOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadVariableHistoryOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadVariableHistoryOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadVariableHistoryOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadVariableHistoryOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadVariableHistoryOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadVariableHistoryOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadVariableHistoryOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadVariableHistoryOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadVariableHistoryOfProjectVersionOK creates a ReadVariableHistoryOfProjectVersionOK with default headers values
func NewReadVariableHistoryOfProjectVersionOK() *ReadVariableHistoryOfProjectVersionOK {
	return &ReadVariableHistoryOfProjectVersionOK{}
}

/*ReadVariableHistoryOfProjectVersionOK handles this case with default header values.

OK
*/
type ReadVariableHistoryOfProjectVersionOK struct {
	Payload *models.APIResultVariableHistory
}

func (o *ReadVariableHistoryOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories/{id}][%d] readVariableHistoryOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ReadVariableHistoryOfProjectVersionOK) GetPayload() *models.APIResultVariableHistory {
	return o.Payload
}

func (o *ReadVariableHistoryOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVariableHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVariableHistoryOfProjectVersionBadRequest creates a ReadVariableHistoryOfProjectVersionBadRequest with default headers values
func NewReadVariableHistoryOfProjectVersionBadRequest() *ReadVariableHistoryOfProjectVersionBadRequest {
	return &ReadVariableHistoryOfProjectVersionBadRequest{}
}

/*ReadVariableHistoryOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ReadVariableHistoryOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadVariableHistoryOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories/{id}][%d] readVariableHistoryOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ReadVariableHistoryOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadVariableHistoryOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVariableHistoryOfProjectVersionUnauthorized creates a ReadVariableHistoryOfProjectVersionUnauthorized with default headers values
func NewReadVariableHistoryOfProjectVersionUnauthorized() *ReadVariableHistoryOfProjectVersionUnauthorized {
	return &ReadVariableHistoryOfProjectVersionUnauthorized{}
}

/*ReadVariableHistoryOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadVariableHistoryOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadVariableHistoryOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories/{id}][%d] readVariableHistoryOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadVariableHistoryOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadVariableHistoryOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVariableHistoryOfProjectVersionForbidden creates a ReadVariableHistoryOfProjectVersionForbidden with default headers values
func NewReadVariableHistoryOfProjectVersionForbidden() *ReadVariableHistoryOfProjectVersionForbidden {
	return &ReadVariableHistoryOfProjectVersionForbidden{}
}

/*ReadVariableHistoryOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ReadVariableHistoryOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ReadVariableHistoryOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories/{id}][%d] readVariableHistoryOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ReadVariableHistoryOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadVariableHistoryOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVariableHistoryOfProjectVersionNotFound creates a ReadVariableHistoryOfProjectVersionNotFound with default headers values
func NewReadVariableHistoryOfProjectVersionNotFound() *ReadVariableHistoryOfProjectVersionNotFound {
	return &ReadVariableHistoryOfProjectVersionNotFound{}
}

/*ReadVariableHistoryOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ReadVariableHistoryOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ReadVariableHistoryOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories/{id}][%d] readVariableHistoryOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ReadVariableHistoryOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadVariableHistoryOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVariableHistoryOfProjectVersionConflict creates a ReadVariableHistoryOfProjectVersionConflict with default headers values
func NewReadVariableHistoryOfProjectVersionConflict() *ReadVariableHistoryOfProjectVersionConflict {
	return &ReadVariableHistoryOfProjectVersionConflict{}
}

/*ReadVariableHistoryOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ReadVariableHistoryOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ReadVariableHistoryOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories/{id}][%d] readVariableHistoryOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ReadVariableHistoryOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadVariableHistoryOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadVariableHistoryOfProjectVersionInternalServerError creates a ReadVariableHistoryOfProjectVersionInternalServerError with default headers values
func NewReadVariableHistoryOfProjectVersionInternalServerError() *ReadVariableHistoryOfProjectVersionInternalServerError {
	return &ReadVariableHistoryOfProjectVersionInternalServerError{}
}

/*ReadVariableHistoryOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadVariableHistoryOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadVariableHistoryOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories/{id}][%d] readVariableHistoryOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadVariableHistoryOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadVariableHistoryOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
