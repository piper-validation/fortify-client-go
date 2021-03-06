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

// ListAttributeOfProjectVersionReader is a Reader for the ListAttributeOfProjectVersion structure.
type ListAttributeOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAttributeOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAttributeOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAttributeOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAttributeOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAttributeOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListAttributeOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListAttributeOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAttributeOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAttributeOfProjectVersionOK creates a ListAttributeOfProjectVersionOK with default headers values
func NewListAttributeOfProjectVersionOK() *ListAttributeOfProjectVersionOK {
	return &ListAttributeOfProjectVersionOK{}
}

/*ListAttributeOfProjectVersionOK handles this case with default header values.

OK
*/
type ListAttributeOfProjectVersionOK struct {
	Payload *models.APIResultListAttribute
}

func (o *ListAttributeOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/attributes][%d] listAttributeOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ListAttributeOfProjectVersionOK) GetPayload() *models.APIResultListAttribute {
	return o.Payload
}

func (o *ListAttributeOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListAttribute)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAttributeOfProjectVersionBadRequest creates a ListAttributeOfProjectVersionBadRequest with default headers values
func NewListAttributeOfProjectVersionBadRequest() *ListAttributeOfProjectVersionBadRequest {
	return &ListAttributeOfProjectVersionBadRequest{}
}

/*ListAttributeOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ListAttributeOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ListAttributeOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/attributes][%d] listAttributeOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ListAttributeOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAttributeOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAttributeOfProjectVersionUnauthorized creates a ListAttributeOfProjectVersionUnauthorized with default headers values
func NewListAttributeOfProjectVersionUnauthorized() *ListAttributeOfProjectVersionUnauthorized {
	return &ListAttributeOfProjectVersionUnauthorized{}
}

/*ListAttributeOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAttributeOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListAttributeOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/attributes][%d] listAttributeOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAttributeOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAttributeOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAttributeOfProjectVersionForbidden creates a ListAttributeOfProjectVersionForbidden with default headers values
func NewListAttributeOfProjectVersionForbidden() *ListAttributeOfProjectVersionForbidden {
	return &ListAttributeOfProjectVersionForbidden{}
}

/*ListAttributeOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ListAttributeOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ListAttributeOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/attributes][%d] listAttributeOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ListAttributeOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAttributeOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAttributeOfProjectVersionNotFound creates a ListAttributeOfProjectVersionNotFound with default headers values
func NewListAttributeOfProjectVersionNotFound() *ListAttributeOfProjectVersionNotFound {
	return &ListAttributeOfProjectVersionNotFound{}
}

/*ListAttributeOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ListAttributeOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ListAttributeOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/attributes][%d] listAttributeOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ListAttributeOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAttributeOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAttributeOfProjectVersionConflict creates a ListAttributeOfProjectVersionConflict with default headers values
func NewListAttributeOfProjectVersionConflict() *ListAttributeOfProjectVersionConflict {
	return &ListAttributeOfProjectVersionConflict{}
}

/*ListAttributeOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ListAttributeOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ListAttributeOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/attributes][%d] listAttributeOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ListAttributeOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAttributeOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAttributeOfProjectVersionInternalServerError creates a ListAttributeOfProjectVersionInternalServerError with default headers values
func NewListAttributeOfProjectVersionInternalServerError() *ListAttributeOfProjectVersionInternalServerError {
	return &ListAttributeOfProjectVersionInternalServerError{}
}

/*ListAttributeOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAttributeOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListAttributeOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/attributes][%d] listAttributeOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAttributeOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAttributeOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
