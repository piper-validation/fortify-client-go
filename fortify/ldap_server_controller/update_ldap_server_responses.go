// Code generated by go-swagger; DO NOT EDIT.

package ldap_server_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// UpdateLdapServerReader is a Reader for the UpdateLdapServer structure.
type UpdateLdapServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateLdapServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateLdapServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateLdapServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateLdapServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateLdapServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateLdapServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateLdapServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateLdapServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateLdapServerOK creates a UpdateLdapServerOK with default headers values
func NewUpdateLdapServerOK() *UpdateLdapServerOK {
	return &UpdateLdapServerOK{}
}

/*UpdateLdapServerOK handles this case with default header values.

OK
*/
type UpdateLdapServerOK struct {
	Payload *models.APIResultLdapServerDto
}

func (o *UpdateLdapServerOK) Error() string {
	return fmt.Sprintf("[PUT /ldapServers/{id}][%d] updateLdapServerOK  %+v", 200, o.Payload)
}

func (o *UpdateLdapServerOK) GetPayload() *models.APIResultLdapServerDto {
	return o.Payload
}

func (o *UpdateLdapServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultLdapServerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapServerBadRequest creates a UpdateLdapServerBadRequest with default headers values
func NewUpdateLdapServerBadRequest() *UpdateLdapServerBadRequest {
	return &UpdateLdapServerBadRequest{}
}

/*UpdateLdapServerBadRequest handles this case with default header values.

Bad Request
*/
type UpdateLdapServerBadRequest struct {
	Payload *models.APIResult
}

func (o *UpdateLdapServerBadRequest) Error() string {
	return fmt.Sprintf("[PUT /ldapServers/{id}][%d] updateLdapServerBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateLdapServerBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateLdapServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapServerUnauthorized creates a UpdateLdapServerUnauthorized with default headers values
func NewUpdateLdapServerUnauthorized() *UpdateLdapServerUnauthorized {
	return &UpdateLdapServerUnauthorized{}
}

/*UpdateLdapServerUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateLdapServerUnauthorized struct {
	Payload *models.APIResult
}

func (o *UpdateLdapServerUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /ldapServers/{id}][%d] updateLdapServerUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateLdapServerUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateLdapServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapServerForbidden creates a UpdateLdapServerForbidden with default headers values
func NewUpdateLdapServerForbidden() *UpdateLdapServerForbidden {
	return &UpdateLdapServerForbidden{}
}

/*UpdateLdapServerForbidden handles this case with default header values.

Forbidden
*/
type UpdateLdapServerForbidden struct {
	Payload *models.APIResult
}

func (o *UpdateLdapServerForbidden) Error() string {
	return fmt.Sprintf("[PUT /ldapServers/{id}][%d] updateLdapServerForbidden  %+v", 403, o.Payload)
}

func (o *UpdateLdapServerForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateLdapServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapServerNotFound creates a UpdateLdapServerNotFound with default headers values
func NewUpdateLdapServerNotFound() *UpdateLdapServerNotFound {
	return &UpdateLdapServerNotFound{}
}

/*UpdateLdapServerNotFound handles this case with default header values.

Not Found
*/
type UpdateLdapServerNotFound struct {
	Payload *models.APIResult
}

func (o *UpdateLdapServerNotFound) Error() string {
	return fmt.Sprintf("[PUT /ldapServers/{id}][%d] updateLdapServerNotFound  %+v", 404, o.Payload)
}

func (o *UpdateLdapServerNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateLdapServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapServerConflict creates a UpdateLdapServerConflict with default headers values
func NewUpdateLdapServerConflict() *UpdateLdapServerConflict {
	return &UpdateLdapServerConflict{}
}

/*UpdateLdapServerConflict handles this case with default header values.

Conflict
*/
type UpdateLdapServerConflict struct {
	Payload *models.APIResult
}

func (o *UpdateLdapServerConflict) Error() string {
	return fmt.Sprintf("[PUT /ldapServers/{id}][%d] updateLdapServerConflict  %+v", 409, o.Payload)
}

func (o *UpdateLdapServerConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateLdapServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateLdapServerInternalServerError creates a UpdateLdapServerInternalServerError with default headers values
func NewUpdateLdapServerInternalServerError() *UpdateLdapServerInternalServerError {
	return &UpdateLdapServerInternalServerError{}
}

/*UpdateLdapServerInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateLdapServerInternalServerError struct {
	Payload *models.APIResult
}

func (o *UpdateLdapServerInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /ldapServers/{id}][%d] updateLdapServerInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateLdapServerInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateLdapServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
