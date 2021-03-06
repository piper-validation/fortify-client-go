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

// CreateLdapServerReader is a Reader for the CreateLdapServer structure.
type CreateLdapServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateLdapServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateLdapServerCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateLdapServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateLdapServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateLdapServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateLdapServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateLdapServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateLdapServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateLdapServerCreated creates a CreateLdapServerCreated with default headers values
func NewCreateLdapServerCreated() *CreateLdapServerCreated {
	return &CreateLdapServerCreated{}
}

/*CreateLdapServerCreated handles this case with default header values.

Created
*/
type CreateLdapServerCreated struct {
	Payload *models.APIResultLdapServerDto
}

func (o *CreateLdapServerCreated) Error() string {
	return fmt.Sprintf("[POST /ldapServers][%d] createLdapServerCreated  %+v", 201, o.Payload)
}

func (o *CreateLdapServerCreated) GetPayload() *models.APIResultLdapServerDto {
	return o.Payload
}

func (o *CreateLdapServerCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultLdapServerDto)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLdapServerBadRequest creates a CreateLdapServerBadRequest with default headers values
func NewCreateLdapServerBadRequest() *CreateLdapServerBadRequest {
	return &CreateLdapServerBadRequest{}
}

/*CreateLdapServerBadRequest handles this case with default header values.

Bad Request
*/
type CreateLdapServerBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateLdapServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /ldapServers][%d] createLdapServerBadRequest  %+v", 400, o.Payload)
}

func (o *CreateLdapServerBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateLdapServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLdapServerUnauthorized creates a CreateLdapServerUnauthorized with default headers values
func NewCreateLdapServerUnauthorized() *CreateLdapServerUnauthorized {
	return &CreateLdapServerUnauthorized{}
}

/*CreateLdapServerUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateLdapServerUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateLdapServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ldapServers][%d] createLdapServerUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateLdapServerUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateLdapServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLdapServerForbidden creates a CreateLdapServerForbidden with default headers values
func NewCreateLdapServerForbidden() *CreateLdapServerForbidden {
	return &CreateLdapServerForbidden{}
}

/*CreateLdapServerForbidden handles this case with default header values.

Forbidden
*/
type CreateLdapServerForbidden struct {
	Payload *models.APIResult
}

func (o *CreateLdapServerForbidden) Error() string {
	return fmt.Sprintf("[POST /ldapServers][%d] createLdapServerForbidden  %+v", 403, o.Payload)
}

func (o *CreateLdapServerForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateLdapServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLdapServerNotFound creates a CreateLdapServerNotFound with default headers values
func NewCreateLdapServerNotFound() *CreateLdapServerNotFound {
	return &CreateLdapServerNotFound{}
}

/*CreateLdapServerNotFound handles this case with default header values.

Not Found
*/
type CreateLdapServerNotFound struct {
	Payload *models.APIResult
}

func (o *CreateLdapServerNotFound) Error() string {
	return fmt.Sprintf("[POST /ldapServers][%d] createLdapServerNotFound  %+v", 404, o.Payload)
}

func (o *CreateLdapServerNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateLdapServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLdapServerConflict creates a CreateLdapServerConflict with default headers values
func NewCreateLdapServerConflict() *CreateLdapServerConflict {
	return &CreateLdapServerConflict{}
}

/*CreateLdapServerConflict handles this case with default header values.

Conflict
*/
type CreateLdapServerConflict struct {
	Payload *models.APIResult
}

func (o *CreateLdapServerConflict) Error() string {
	return fmt.Sprintf("[POST /ldapServers][%d] createLdapServerConflict  %+v", 409, o.Payload)
}

func (o *CreateLdapServerConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateLdapServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateLdapServerInternalServerError creates a CreateLdapServerInternalServerError with default headers values
func NewCreateLdapServerInternalServerError() *CreateLdapServerInternalServerError {
	return &CreateLdapServerInternalServerError{}
}

/*CreateLdapServerInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateLdapServerInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateLdapServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ldapServers][%d] createLdapServerInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateLdapServerInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateLdapServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
