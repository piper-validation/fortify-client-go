// Code generated by go-swagger; DO NOT EDIT.

package folder_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListFolderOfProjectVersionReader is a Reader for the ListFolderOfProjectVersion structure.
type ListFolderOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListFolderOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListFolderOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListFolderOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListFolderOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListFolderOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListFolderOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListFolderOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListFolderOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListFolderOfProjectVersionOK creates a ListFolderOfProjectVersionOK with default headers values
func NewListFolderOfProjectVersionOK() *ListFolderOfProjectVersionOK {
	return &ListFolderOfProjectVersionOK{}
}

/*ListFolderOfProjectVersionOK handles this case with default header values.

OK
*/
type ListFolderOfProjectVersionOK struct {
	Payload *models.APIResultListFolder
}

func (o *ListFolderOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/folders][%d] listFolderOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ListFolderOfProjectVersionOK) GetPayload() *models.APIResultListFolder {
	return o.Payload
}

func (o *ListFolderOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListFolder)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFolderOfProjectVersionBadRequest creates a ListFolderOfProjectVersionBadRequest with default headers values
func NewListFolderOfProjectVersionBadRequest() *ListFolderOfProjectVersionBadRequest {
	return &ListFolderOfProjectVersionBadRequest{}
}

/*ListFolderOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ListFolderOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ListFolderOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/folders][%d] listFolderOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ListFolderOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListFolderOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFolderOfProjectVersionUnauthorized creates a ListFolderOfProjectVersionUnauthorized with default headers values
func NewListFolderOfProjectVersionUnauthorized() *ListFolderOfProjectVersionUnauthorized {
	return &ListFolderOfProjectVersionUnauthorized{}
}

/*ListFolderOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ListFolderOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListFolderOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/folders][%d] listFolderOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListFolderOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListFolderOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFolderOfProjectVersionForbidden creates a ListFolderOfProjectVersionForbidden with default headers values
func NewListFolderOfProjectVersionForbidden() *ListFolderOfProjectVersionForbidden {
	return &ListFolderOfProjectVersionForbidden{}
}

/*ListFolderOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ListFolderOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ListFolderOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/folders][%d] listFolderOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ListFolderOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListFolderOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFolderOfProjectVersionNotFound creates a ListFolderOfProjectVersionNotFound with default headers values
func NewListFolderOfProjectVersionNotFound() *ListFolderOfProjectVersionNotFound {
	return &ListFolderOfProjectVersionNotFound{}
}

/*ListFolderOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ListFolderOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ListFolderOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/folders][%d] listFolderOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ListFolderOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListFolderOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFolderOfProjectVersionConflict creates a ListFolderOfProjectVersionConflict with default headers values
func NewListFolderOfProjectVersionConflict() *ListFolderOfProjectVersionConflict {
	return &ListFolderOfProjectVersionConflict{}
}

/*ListFolderOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ListFolderOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ListFolderOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/folders][%d] listFolderOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ListFolderOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListFolderOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListFolderOfProjectVersionInternalServerError creates a ListFolderOfProjectVersionInternalServerError with default headers values
func NewListFolderOfProjectVersionInternalServerError() *ListFolderOfProjectVersionInternalServerError {
	return &ListFolderOfProjectVersionInternalServerError{}
}

/*ListFolderOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListFolderOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListFolderOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/folders][%d] listFolderOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListFolderOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListFolderOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}