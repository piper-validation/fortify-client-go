// Code generated by go-swagger; DO NOT EDIT.

package group_of_auth_entity_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListGroupOfAuthEntityReader is a Reader for the ListGroupOfAuthEntity structure.
type ListGroupOfAuthEntityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListGroupOfAuthEntityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListGroupOfAuthEntityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListGroupOfAuthEntityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListGroupOfAuthEntityUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListGroupOfAuthEntityForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListGroupOfAuthEntityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListGroupOfAuthEntityConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListGroupOfAuthEntityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListGroupOfAuthEntityOK creates a ListGroupOfAuthEntityOK with default headers values
func NewListGroupOfAuthEntityOK() *ListGroupOfAuthEntityOK {
	return &ListGroupOfAuthEntityOK{}
}

/*ListGroupOfAuthEntityOK handles this case with default header values.

OK
*/
type ListGroupOfAuthEntityOK struct {
	Payload *models.APIResultListAuthenticationEntity
}

func (o *ListGroupOfAuthEntityOK) Error() string {
	return fmt.Sprintf("[GET /authEntities/{parentId}/groups][%d] listGroupOfAuthEntityOK  %+v", 200, o.Payload)
}

func (o *ListGroupOfAuthEntityOK) GetPayload() *models.APIResultListAuthenticationEntity {
	return o.Payload
}

func (o *ListGroupOfAuthEntityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListAuthenticationEntity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupOfAuthEntityBadRequest creates a ListGroupOfAuthEntityBadRequest with default headers values
func NewListGroupOfAuthEntityBadRequest() *ListGroupOfAuthEntityBadRequest {
	return &ListGroupOfAuthEntityBadRequest{}
}

/*ListGroupOfAuthEntityBadRequest handles this case with default header values.

Bad Request
*/
type ListGroupOfAuthEntityBadRequest struct {
	Payload *models.APIResult
}

func (o *ListGroupOfAuthEntityBadRequest) Error() string {
	return fmt.Sprintf("[GET /authEntities/{parentId}/groups][%d] listGroupOfAuthEntityBadRequest  %+v", 400, o.Payload)
}

func (o *ListGroupOfAuthEntityBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListGroupOfAuthEntityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupOfAuthEntityUnauthorized creates a ListGroupOfAuthEntityUnauthorized with default headers values
func NewListGroupOfAuthEntityUnauthorized() *ListGroupOfAuthEntityUnauthorized {
	return &ListGroupOfAuthEntityUnauthorized{}
}

/*ListGroupOfAuthEntityUnauthorized handles this case with default header values.

Unauthorized
*/
type ListGroupOfAuthEntityUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListGroupOfAuthEntityUnauthorized) Error() string {
	return fmt.Sprintf("[GET /authEntities/{parentId}/groups][%d] listGroupOfAuthEntityUnauthorized  %+v", 401, o.Payload)
}

func (o *ListGroupOfAuthEntityUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListGroupOfAuthEntityUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupOfAuthEntityForbidden creates a ListGroupOfAuthEntityForbidden with default headers values
func NewListGroupOfAuthEntityForbidden() *ListGroupOfAuthEntityForbidden {
	return &ListGroupOfAuthEntityForbidden{}
}

/*ListGroupOfAuthEntityForbidden handles this case with default header values.

Forbidden
*/
type ListGroupOfAuthEntityForbidden struct {
	Payload *models.APIResult
}

func (o *ListGroupOfAuthEntityForbidden) Error() string {
	return fmt.Sprintf("[GET /authEntities/{parentId}/groups][%d] listGroupOfAuthEntityForbidden  %+v", 403, o.Payload)
}

func (o *ListGroupOfAuthEntityForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListGroupOfAuthEntityForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupOfAuthEntityNotFound creates a ListGroupOfAuthEntityNotFound with default headers values
func NewListGroupOfAuthEntityNotFound() *ListGroupOfAuthEntityNotFound {
	return &ListGroupOfAuthEntityNotFound{}
}

/*ListGroupOfAuthEntityNotFound handles this case with default header values.

Not Found
*/
type ListGroupOfAuthEntityNotFound struct {
	Payload *models.APIResult
}

func (o *ListGroupOfAuthEntityNotFound) Error() string {
	return fmt.Sprintf("[GET /authEntities/{parentId}/groups][%d] listGroupOfAuthEntityNotFound  %+v", 404, o.Payload)
}

func (o *ListGroupOfAuthEntityNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListGroupOfAuthEntityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupOfAuthEntityConflict creates a ListGroupOfAuthEntityConflict with default headers values
func NewListGroupOfAuthEntityConflict() *ListGroupOfAuthEntityConflict {
	return &ListGroupOfAuthEntityConflict{}
}

/*ListGroupOfAuthEntityConflict handles this case with default header values.

Conflict
*/
type ListGroupOfAuthEntityConflict struct {
	Payload *models.APIResult
}

func (o *ListGroupOfAuthEntityConflict) Error() string {
	return fmt.Sprintf("[GET /authEntities/{parentId}/groups][%d] listGroupOfAuthEntityConflict  %+v", 409, o.Payload)
}

func (o *ListGroupOfAuthEntityConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListGroupOfAuthEntityConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListGroupOfAuthEntityInternalServerError creates a ListGroupOfAuthEntityInternalServerError with default headers values
func NewListGroupOfAuthEntityInternalServerError() *ListGroupOfAuthEntityInternalServerError {
	return &ListGroupOfAuthEntityInternalServerError{}
}

/*ListGroupOfAuthEntityInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListGroupOfAuthEntityInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListGroupOfAuthEntityInternalServerError) Error() string {
	return fmt.Sprintf("[GET /authEntities/{parentId}/groups][%d] listGroupOfAuthEntityInternalServerError  %+v", 500, o.Payload)
}

func (o *ListGroupOfAuthEntityInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListGroupOfAuthEntityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}