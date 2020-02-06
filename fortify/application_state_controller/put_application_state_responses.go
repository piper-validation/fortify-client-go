// Code generated by go-swagger; DO NOT EDIT.

package application_state_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// PutApplicationStateReader is a Reader for the PutApplicationState structure.
type PutApplicationStateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PutApplicationStateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPutApplicationStateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPutApplicationStateBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPutApplicationStateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPutApplicationStateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPutApplicationStateNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPutApplicationStateConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPutApplicationStateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPutApplicationStateOK creates a PutApplicationStateOK with default headers values
func NewPutApplicationStateOK() *PutApplicationStateOK {
	return &PutApplicationStateOK{}
}

/*PutApplicationStateOK handles this case with default header values.

OK
*/
type PutApplicationStateOK struct {
	Payload *models.APIResultApplicationState
}

func (o *PutApplicationStateOK) Error() string {
	return fmt.Sprintf("[PUT /applicationState][%d] putApplicationStateOK  %+v", 200, o.Payload)
}

func (o *PutApplicationStateOK) GetPayload() *models.APIResultApplicationState {
	return o.Payload
}

func (o *PutApplicationStateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultApplicationState)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationStateBadRequest creates a PutApplicationStateBadRequest with default headers values
func NewPutApplicationStateBadRequest() *PutApplicationStateBadRequest {
	return &PutApplicationStateBadRequest{}
}

/*PutApplicationStateBadRequest handles this case with default header values.

Bad Request
*/
type PutApplicationStateBadRequest struct {
	Payload *models.APIResult
}

func (o *PutApplicationStateBadRequest) Error() string {
	return fmt.Sprintf("[PUT /applicationState][%d] putApplicationStateBadRequest  %+v", 400, o.Payload)
}

func (o *PutApplicationStateBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PutApplicationStateBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationStateUnauthorized creates a PutApplicationStateUnauthorized with default headers values
func NewPutApplicationStateUnauthorized() *PutApplicationStateUnauthorized {
	return &PutApplicationStateUnauthorized{}
}

/*PutApplicationStateUnauthorized handles this case with default header values.

Unauthorized
*/
type PutApplicationStateUnauthorized struct {
	Payload *models.APIResult
}

func (o *PutApplicationStateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /applicationState][%d] putApplicationStateUnauthorized  %+v", 401, o.Payload)
}

func (o *PutApplicationStateUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PutApplicationStateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationStateForbidden creates a PutApplicationStateForbidden with default headers values
func NewPutApplicationStateForbidden() *PutApplicationStateForbidden {
	return &PutApplicationStateForbidden{}
}

/*PutApplicationStateForbidden handles this case with default header values.

Forbidden
*/
type PutApplicationStateForbidden struct {
	Payload *models.APIResult
}

func (o *PutApplicationStateForbidden) Error() string {
	return fmt.Sprintf("[PUT /applicationState][%d] putApplicationStateForbidden  %+v", 403, o.Payload)
}

func (o *PutApplicationStateForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PutApplicationStateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationStateNotFound creates a PutApplicationStateNotFound with default headers values
func NewPutApplicationStateNotFound() *PutApplicationStateNotFound {
	return &PutApplicationStateNotFound{}
}

/*PutApplicationStateNotFound handles this case with default header values.

Not Found
*/
type PutApplicationStateNotFound struct {
	Payload *models.APIResult
}

func (o *PutApplicationStateNotFound) Error() string {
	return fmt.Sprintf("[PUT /applicationState][%d] putApplicationStateNotFound  %+v", 404, o.Payload)
}

func (o *PutApplicationStateNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PutApplicationStateNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationStateConflict creates a PutApplicationStateConflict with default headers values
func NewPutApplicationStateConflict() *PutApplicationStateConflict {
	return &PutApplicationStateConflict{}
}

/*PutApplicationStateConflict handles this case with default header values.

Conflict
*/
type PutApplicationStateConflict struct {
	Payload *models.APIResult
}

func (o *PutApplicationStateConflict) Error() string {
	return fmt.Sprintf("[PUT /applicationState][%d] putApplicationStateConflict  %+v", 409, o.Payload)
}

func (o *PutApplicationStateConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PutApplicationStateConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPutApplicationStateInternalServerError creates a PutApplicationStateInternalServerError with default headers values
func NewPutApplicationStateInternalServerError() *PutApplicationStateInternalServerError {
	return &PutApplicationStateInternalServerError{}
}

/*PutApplicationStateInternalServerError handles this case with default header values.

Internal Server Error
*/
type PutApplicationStateInternalServerError struct {
	Payload *models.APIResult
}

func (o *PutApplicationStateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /applicationState][%d] putApplicationStateInternalServerError  %+v", 500, o.Payload)
}

func (o *PutApplicationStateInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PutApplicationStateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}