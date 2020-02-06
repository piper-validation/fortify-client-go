// Code generated by go-swagger; DO NOT EDIT.

package custom_tag_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// DeleteCustomTagReader is a Reader for the DeleteCustomTag structure.
type DeleteCustomTagReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteCustomTagReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteCustomTagOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteCustomTagBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteCustomTagUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteCustomTagForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteCustomTagNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteCustomTagConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteCustomTagInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteCustomTagOK creates a DeleteCustomTagOK with default headers values
func NewDeleteCustomTagOK() *DeleteCustomTagOK {
	return &DeleteCustomTagOK{}
}

/*DeleteCustomTagOK handles this case with default header values.

OK
*/
type DeleteCustomTagOK struct {
	Payload *models.APIResultVoid
}

func (o *DeleteCustomTagOK) Error() string {
	return fmt.Sprintf("[DELETE /customTags/{id}][%d] deleteCustomTagOK  %+v", 200, o.Payload)
}

func (o *DeleteCustomTagOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *DeleteCustomTagOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomTagBadRequest creates a DeleteCustomTagBadRequest with default headers values
func NewDeleteCustomTagBadRequest() *DeleteCustomTagBadRequest {
	return &DeleteCustomTagBadRequest{}
}

/*DeleteCustomTagBadRequest handles this case with default header values.

Bad Request
*/
type DeleteCustomTagBadRequest struct {
	Payload *models.APIResult
}

func (o *DeleteCustomTagBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /customTags/{id}][%d] deleteCustomTagBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteCustomTagBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteCustomTagBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomTagUnauthorized creates a DeleteCustomTagUnauthorized with default headers values
func NewDeleteCustomTagUnauthorized() *DeleteCustomTagUnauthorized {
	return &DeleteCustomTagUnauthorized{}
}

/*DeleteCustomTagUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteCustomTagUnauthorized struct {
	Payload *models.APIResult
}

func (o *DeleteCustomTagUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /customTags/{id}][%d] deleteCustomTagUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteCustomTagUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteCustomTagUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomTagForbidden creates a DeleteCustomTagForbidden with default headers values
func NewDeleteCustomTagForbidden() *DeleteCustomTagForbidden {
	return &DeleteCustomTagForbidden{}
}

/*DeleteCustomTagForbidden handles this case with default header values.

Forbidden
*/
type DeleteCustomTagForbidden struct {
	Payload *models.APIResult
}

func (o *DeleteCustomTagForbidden) Error() string {
	return fmt.Sprintf("[DELETE /customTags/{id}][%d] deleteCustomTagForbidden  %+v", 403, o.Payload)
}

func (o *DeleteCustomTagForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteCustomTagForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomTagNotFound creates a DeleteCustomTagNotFound with default headers values
func NewDeleteCustomTagNotFound() *DeleteCustomTagNotFound {
	return &DeleteCustomTagNotFound{}
}

/*DeleteCustomTagNotFound handles this case with default header values.

Not Found
*/
type DeleteCustomTagNotFound struct {
	Payload *models.APIResult
}

func (o *DeleteCustomTagNotFound) Error() string {
	return fmt.Sprintf("[DELETE /customTags/{id}][%d] deleteCustomTagNotFound  %+v", 404, o.Payload)
}

func (o *DeleteCustomTagNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteCustomTagNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomTagConflict creates a DeleteCustomTagConflict with default headers values
func NewDeleteCustomTagConflict() *DeleteCustomTagConflict {
	return &DeleteCustomTagConflict{}
}

/*DeleteCustomTagConflict handles this case with default header values.

Conflict
*/
type DeleteCustomTagConflict struct {
	Payload *models.APIResult
}

func (o *DeleteCustomTagConflict) Error() string {
	return fmt.Sprintf("[DELETE /customTags/{id}][%d] deleteCustomTagConflict  %+v", 409, o.Payload)
}

func (o *DeleteCustomTagConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteCustomTagConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteCustomTagInternalServerError creates a DeleteCustomTagInternalServerError with default headers values
func NewDeleteCustomTagInternalServerError() *DeleteCustomTagInternalServerError {
	return &DeleteCustomTagInternalServerError{}
}

/*DeleteCustomTagInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteCustomTagInternalServerError struct {
	Payload *models.APIResult
}

func (o *DeleteCustomTagInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /customTags/{id}][%d] deleteCustomTagInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteCustomTagInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteCustomTagInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}