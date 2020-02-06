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

// ListVariableHistoryOfProjectVersionReader is a Reader for the ListVariableHistoryOfProjectVersion structure.
type ListVariableHistoryOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListVariableHistoryOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListVariableHistoryOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListVariableHistoryOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListVariableHistoryOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListVariableHistoryOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListVariableHistoryOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListVariableHistoryOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListVariableHistoryOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListVariableHistoryOfProjectVersionOK creates a ListVariableHistoryOfProjectVersionOK with default headers values
func NewListVariableHistoryOfProjectVersionOK() *ListVariableHistoryOfProjectVersionOK {
	return &ListVariableHistoryOfProjectVersionOK{}
}

/*ListVariableHistoryOfProjectVersionOK handles this case with default header values.

OK
*/
type ListVariableHistoryOfProjectVersionOK struct {
	Payload *models.APIResultListVariableHistory
}

func (o *ListVariableHistoryOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories][%d] listVariableHistoryOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ListVariableHistoryOfProjectVersionOK) GetPayload() *models.APIResultListVariableHistory {
	return o.Payload
}

func (o *ListVariableHistoryOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListVariableHistory)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVariableHistoryOfProjectVersionBadRequest creates a ListVariableHistoryOfProjectVersionBadRequest with default headers values
func NewListVariableHistoryOfProjectVersionBadRequest() *ListVariableHistoryOfProjectVersionBadRequest {
	return &ListVariableHistoryOfProjectVersionBadRequest{}
}

/*ListVariableHistoryOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ListVariableHistoryOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ListVariableHistoryOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories][%d] listVariableHistoryOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ListVariableHistoryOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListVariableHistoryOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVariableHistoryOfProjectVersionUnauthorized creates a ListVariableHistoryOfProjectVersionUnauthorized with default headers values
func NewListVariableHistoryOfProjectVersionUnauthorized() *ListVariableHistoryOfProjectVersionUnauthorized {
	return &ListVariableHistoryOfProjectVersionUnauthorized{}
}

/*ListVariableHistoryOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ListVariableHistoryOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListVariableHistoryOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories][%d] listVariableHistoryOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListVariableHistoryOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListVariableHistoryOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVariableHistoryOfProjectVersionForbidden creates a ListVariableHistoryOfProjectVersionForbidden with default headers values
func NewListVariableHistoryOfProjectVersionForbidden() *ListVariableHistoryOfProjectVersionForbidden {
	return &ListVariableHistoryOfProjectVersionForbidden{}
}

/*ListVariableHistoryOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ListVariableHistoryOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ListVariableHistoryOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories][%d] listVariableHistoryOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ListVariableHistoryOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListVariableHistoryOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVariableHistoryOfProjectVersionNotFound creates a ListVariableHistoryOfProjectVersionNotFound with default headers values
func NewListVariableHistoryOfProjectVersionNotFound() *ListVariableHistoryOfProjectVersionNotFound {
	return &ListVariableHistoryOfProjectVersionNotFound{}
}

/*ListVariableHistoryOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ListVariableHistoryOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ListVariableHistoryOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories][%d] listVariableHistoryOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ListVariableHistoryOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListVariableHistoryOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVariableHistoryOfProjectVersionConflict creates a ListVariableHistoryOfProjectVersionConflict with default headers values
func NewListVariableHistoryOfProjectVersionConflict() *ListVariableHistoryOfProjectVersionConflict {
	return &ListVariableHistoryOfProjectVersionConflict{}
}

/*ListVariableHistoryOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ListVariableHistoryOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ListVariableHistoryOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories][%d] listVariableHistoryOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ListVariableHistoryOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListVariableHistoryOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListVariableHistoryOfProjectVersionInternalServerError creates a ListVariableHistoryOfProjectVersionInternalServerError with default headers values
func NewListVariableHistoryOfProjectVersionInternalServerError() *ListVariableHistoryOfProjectVersionInternalServerError {
	return &ListVariableHistoryOfProjectVersionInternalServerError{}
}

/*ListVariableHistoryOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListVariableHistoryOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListVariableHistoryOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/variableHistories][%d] listVariableHistoryOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListVariableHistoryOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListVariableHistoryOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}