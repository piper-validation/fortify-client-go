// Code generated by go-swagger; DO NOT EDIT.

package ldap_object_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// MultiDeleteLdapObjectReader is a Reader for the MultiDeleteLdapObject structure.
type MultiDeleteLdapObjectReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiDeleteLdapObjectReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiDeleteLdapObjectOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMultiDeleteLdapObjectBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMultiDeleteLdapObjectUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMultiDeleteLdapObjectForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMultiDeleteLdapObjectNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMultiDeleteLdapObjectConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMultiDeleteLdapObjectInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMultiDeleteLdapObjectOK creates a MultiDeleteLdapObjectOK with default headers values
func NewMultiDeleteLdapObjectOK() *MultiDeleteLdapObjectOK {
	return &MultiDeleteLdapObjectOK{}
}

/*MultiDeleteLdapObjectOK handles this case with default header values.

OK
*/
type MultiDeleteLdapObjectOK struct {
	Payload *models.APIResultVoid
}

func (o *MultiDeleteLdapObjectOK) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects][%d] multiDeleteLdapObjectOK  %+v", 200, o.Payload)
}

func (o *MultiDeleteLdapObjectOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *MultiDeleteLdapObjectOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteLdapObjectBadRequest creates a MultiDeleteLdapObjectBadRequest with default headers values
func NewMultiDeleteLdapObjectBadRequest() *MultiDeleteLdapObjectBadRequest {
	return &MultiDeleteLdapObjectBadRequest{}
}

/*MultiDeleteLdapObjectBadRequest handles this case with default header values.

Bad Request
*/
type MultiDeleteLdapObjectBadRequest struct {
	Payload *models.APIResult
}

func (o *MultiDeleteLdapObjectBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects][%d] multiDeleteLdapObjectBadRequest  %+v", 400, o.Payload)
}

func (o *MultiDeleteLdapObjectBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteLdapObjectBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteLdapObjectUnauthorized creates a MultiDeleteLdapObjectUnauthorized with default headers values
func NewMultiDeleteLdapObjectUnauthorized() *MultiDeleteLdapObjectUnauthorized {
	return &MultiDeleteLdapObjectUnauthorized{}
}

/*MultiDeleteLdapObjectUnauthorized handles this case with default header values.

Unauthorized
*/
type MultiDeleteLdapObjectUnauthorized struct {
	Payload *models.APIResult
}

func (o *MultiDeleteLdapObjectUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects][%d] multiDeleteLdapObjectUnauthorized  %+v", 401, o.Payload)
}

func (o *MultiDeleteLdapObjectUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteLdapObjectUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteLdapObjectForbidden creates a MultiDeleteLdapObjectForbidden with default headers values
func NewMultiDeleteLdapObjectForbidden() *MultiDeleteLdapObjectForbidden {
	return &MultiDeleteLdapObjectForbidden{}
}

/*MultiDeleteLdapObjectForbidden handles this case with default header values.

Forbidden
*/
type MultiDeleteLdapObjectForbidden struct {
	Payload *models.APIResult
}

func (o *MultiDeleteLdapObjectForbidden) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects][%d] multiDeleteLdapObjectForbidden  %+v", 403, o.Payload)
}

func (o *MultiDeleteLdapObjectForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteLdapObjectForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteLdapObjectNotFound creates a MultiDeleteLdapObjectNotFound with default headers values
func NewMultiDeleteLdapObjectNotFound() *MultiDeleteLdapObjectNotFound {
	return &MultiDeleteLdapObjectNotFound{}
}

/*MultiDeleteLdapObjectNotFound handles this case with default header values.

Not Found
*/
type MultiDeleteLdapObjectNotFound struct {
	Payload *models.APIResult
}

func (o *MultiDeleteLdapObjectNotFound) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects][%d] multiDeleteLdapObjectNotFound  %+v", 404, o.Payload)
}

func (o *MultiDeleteLdapObjectNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteLdapObjectNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteLdapObjectConflict creates a MultiDeleteLdapObjectConflict with default headers values
func NewMultiDeleteLdapObjectConflict() *MultiDeleteLdapObjectConflict {
	return &MultiDeleteLdapObjectConflict{}
}

/*MultiDeleteLdapObjectConflict handles this case with default header values.

Conflict
*/
type MultiDeleteLdapObjectConflict struct {
	Payload *models.APIResult
}

func (o *MultiDeleteLdapObjectConflict) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects][%d] multiDeleteLdapObjectConflict  %+v", 409, o.Payload)
}

func (o *MultiDeleteLdapObjectConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteLdapObjectConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteLdapObjectInternalServerError creates a MultiDeleteLdapObjectInternalServerError with default headers values
func NewMultiDeleteLdapObjectInternalServerError() *MultiDeleteLdapObjectInternalServerError {
	return &MultiDeleteLdapObjectInternalServerError{}
}

/*MultiDeleteLdapObjectInternalServerError handles this case with default header values.

Internal Server Error
*/
type MultiDeleteLdapObjectInternalServerError struct {
	Payload *models.APIResult
}

func (o *MultiDeleteLdapObjectInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /ldapObjects][%d] multiDeleteLdapObjectInternalServerError  %+v", 500, o.Payload)
}

func (o *MultiDeleteLdapObjectInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteLdapObjectInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
