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

// TestLdapServerReader is a Reader for the TestLdapServer structure.
type TestLdapServerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *TestLdapServerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewTestLdapServerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewTestLdapServerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewTestLdapServerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewTestLdapServerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewTestLdapServerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewTestLdapServerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewTestLdapServerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewTestLdapServerOK creates a TestLdapServerOK with default headers values
func NewTestLdapServerOK() *TestLdapServerOK {
	return &TestLdapServerOK{}
}

/*TestLdapServerOK handles this case with default header values.

OK
*/
type TestLdapServerOK struct {
	Payload *models.APIResultValidationStatus
}

func (o *TestLdapServerOK) Error() string {
	return fmt.Sprintf("[POST /ldapServers/action/test][%d] testLdapServerOK  %+v", 200, o.Payload)
}

func (o *TestLdapServerOK) GetPayload() *models.APIResultValidationStatus {
	return o.Payload
}

func (o *TestLdapServerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultValidationStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapServerBadRequest creates a TestLdapServerBadRequest with default headers values
func NewTestLdapServerBadRequest() *TestLdapServerBadRequest {
	return &TestLdapServerBadRequest{}
}

/*TestLdapServerBadRequest handles this case with default header values.

Bad Request
*/
type TestLdapServerBadRequest struct {
	Payload *models.APIResult
}

func (o *TestLdapServerBadRequest) Error() string {
	return fmt.Sprintf("[POST /ldapServers/action/test][%d] testLdapServerBadRequest  %+v", 400, o.Payload)
}

func (o *TestLdapServerBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *TestLdapServerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapServerUnauthorized creates a TestLdapServerUnauthorized with default headers values
func NewTestLdapServerUnauthorized() *TestLdapServerUnauthorized {
	return &TestLdapServerUnauthorized{}
}

/*TestLdapServerUnauthorized handles this case with default header values.

Unauthorized
*/
type TestLdapServerUnauthorized struct {
	Payload *models.APIResult
}

func (o *TestLdapServerUnauthorized) Error() string {
	return fmt.Sprintf("[POST /ldapServers/action/test][%d] testLdapServerUnauthorized  %+v", 401, o.Payload)
}

func (o *TestLdapServerUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *TestLdapServerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapServerForbidden creates a TestLdapServerForbidden with default headers values
func NewTestLdapServerForbidden() *TestLdapServerForbidden {
	return &TestLdapServerForbidden{}
}

/*TestLdapServerForbidden handles this case with default header values.

Forbidden
*/
type TestLdapServerForbidden struct {
	Payload *models.APIResult
}

func (o *TestLdapServerForbidden) Error() string {
	return fmt.Sprintf("[POST /ldapServers/action/test][%d] testLdapServerForbidden  %+v", 403, o.Payload)
}

func (o *TestLdapServerForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *TestLdapServerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapServerNotFound creates a TestLdapServerNotFound with default headers values
func NewTestLdapServerNotFound() *TestLdapServerNotFound {
	return &TestLdapServerNotFound{}
}

/*TestLdapServerNotFound handles this case with default header values.

Not Found
*/
type TestLdapServerNotFound struct {
	Payload *models.APIResult
}

func (o *TestLdapServerNotFound) Error() string {
	return fmt.Sprintf("[POST /ldapServers/action/test][%d] testLdapServerNotFound  %+v", 404, o.Payload)
}

func (o *TestLdapServerNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *TestLdapServerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapServerConflict creates a TestLdapServerConflict with default headers values
func NewTestLdapServerConflict() *TestLdapServerConflict {
	return &TestLdapServerConflict{}
}

/*TestLdapServerConflict handles this case with default header values.

Conflict
*/
type TestLdapServerConflict struct {
	Payload *models.APIResult
}

func (o *TestLdapServerConflict) Error() string {
	return fmt.Sprintf("[POST /ldapServers/action/test][%d] testLdapServerConflict  %+v", 409, o.Payload)
}

func (o *TestLdapServerConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *TestLdapServerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewTestLdapServerInternalServerError creates a TestLdapServerInternalServerError with default headers values
func NewTestLdapServerInternalServerError() *TestLdapServerInternalServerError {
	return &TestLdapServerInternalServerError{}
}

/*TestLdapServerInternalServerError handles this case with default header values.

Internal Server Error
*/
type TestLdapServerInternalServerError struct {
	Payload *models.APIResult
}

func (o *TestLdapServerInternalServerError) Error() string {
	return fmt.Sprintf("[POST /ldapServers/action/test][%d] testLdapServerInternalServerError  %+v", 500, o.Payload)
}

func (o *TestLdapServerInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *TestLdapServerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}