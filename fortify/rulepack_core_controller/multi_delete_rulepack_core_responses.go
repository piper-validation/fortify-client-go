// Code generated by go-swagger; DO NOT EDIT.

package rulepack_core_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// MultiDeleteRulepackCoreReader is a Reader for the MultiDeleteRulepackCore structure.
type MultiDeleteRulepackCoreReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MultiDeleteRulepackCoreReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMultiDeleteRulepackCoreOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMultiDeleteRulepackCoreBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMultiDeleteRulepackCoreUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMultiDeleteRulepackCoreForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMultiDeleteRulepackCoreNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMultiDeleteRulepackCoreConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMultiDeleteRulepackCoreInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMultiDeleteRulepackCoreOK creates a MultiDeleteRulepackCoreOK with default headers values
func NewMultiDeleteRulepackCoreOK() *MultiDeleteRulepackCoreOK {
	return &MultiDeleteRulepackCoreOK{}
}

/*MultiDeleteRulepackCoreOK handles this case with default header values.

OK
*/
type MultiDeleteRulepackCoreOK struct {
	Payload *models.APIResultVoid
}

func (o *MultiDeleteRulepackCoreOK) Error() string {
	return fmt.Sprintf("[DELETE /coreRulepacks][%d] multiDeleteRulepackCoreOK  %+v", 200, o.Payload)
}

func (o *MultiDeleteRulepackCoreOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *MultiDeleteRulepackCoreOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteRulepackCoreBadRequest creates a MultiDeleteRulepackCoreBadRequest with default headers values
func NewMultiDeleteRulepackCoreBadRequest() *MultiDeleteRulepackCoreBadRequest {
	return &MultiDeleteRulepackCoreBadRequest{}
}

/*MultiDeleteRulepackCoreBadRequest handles this case with default header values.

Bad Request
*/
type MultiDeleteRulepackCoreBadRequest struct {
	Payload *models.APIResult
}

func (o *MultiDeleteRulepackCoreBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /coreRulepacks][%d] multiDeleteRulepackCoreBadRequest  %+v", 400, o.Payload)
}

func (o *MultiDeleteRulepackCoreBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteRulepackCoreBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteRulepackCoreUnauthorized creates a MultiDeleteRulepackCoreUnauthorized with default headers values
func NewMultiDeleteRulepackCoreUnauthorized() *MultiDeleteRulepackCoreUnauthorized {
	return &MultiDeleteRulepackCoreUnauthorized{}
}

/*MultiDeleteRulepackCoreUnauthorized handles this case with default header values.

Unauthorized
*/
type MultiDeleteRulepackCoreUnauthorized struct {
	Payload *models.APIResult
}

func (o *MultiDeleteRulepackCoreUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /coreRulepacks][%d] multiDeleteRulepackCoreUnauthorized  %+v", 401, o.Payload)
}

func (o *MultiDeleteRulepackCoreUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteRulepackCoreUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteRulepackCoreForbidden creates a MultiDeleteRulepackCoreForbidden with default headers values
func NewMultiDeleteRulepackCoreForbidden() *MultiDeleteRulepackCoreForbidden {
	return &MultiDeleteRulepackCoreForbidden{}
}

/*MultiDeleteRulepackCoreForbidden handles this case with default header values.

Forbidden
*/
type MultiDeleteRulepackCoreForbidden struct {
	Payload *models.APIResult
}

func (o *MultiDeleteRulepackCoreForbidden) Error() string {
	return fmt.Sprintf("[DELETE /coreRulepacks][%d] multiDeleteRulepackCoreForbidden  %+v", 403, o.Payload)
}

func (o *MultiDeleteRulepackCoreForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteRulepackCoreForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteRulepackCoreNotFound creates a MultiDeleteRulepackCoreNotFound with default headers values
func NewMultiDeleteRulepackCoreNotFound() *MultiDeleteRulepackCoreNotFound {
	return &MultiDeleteRulepackCoreNotFound{}
}

/*MultiDeleteRulepackCoreNotFound handles this case with default header values.

Not Found
*/
type MultiDeleteRulepackCoreNotFound struct {
	Payload *models.APIResult
}

func (o *MultiDeleteRulepackCoreNotFound) Error() string {
	return fmt.Sprintf("[DELETE /coreRulepacks][%d] multiDeleteRulepackCoreNotFound  %+v", 404, o.Payload)
}

func (o *MultiDeleteRulepackCoreNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteRulepackCoreNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteRulepackCoreConflict creates a MultiDeleteRulepackCoreConflict with default headers values
func NewMultiDeleteRulepackCoreConflict() *MultiDeleteRulepackCoreConflict {
	return &MultiDeleteRulepackCoreConflict{}
}

/*MultiDeleteRulepackCoreConflict handles this case with default header values.

Conflict
*/
type MultiDeleteRulepackCoreConflict struct {
	Payload *models.APIResult
}

func (o *MultiDeleteRulepackCoreConflict) Error() string {
	return fmt.Sprintf("[DELETE /coreRulepacks][%d] multiDeleteRulepackCoreConflict  %+v", 409, o.Payload)
}

func (o *MultiDeleteRulepackCoreConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteRulepackCoreConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMultiDeleteRulepackCoreInternalServerError creates a MultiDeleteRulepackCoreInternalServerError with default headers values
func NewMultiDeleteRulepackCoreInternalServerError() *MultiDeleteRulepackCoreInternalServerError {
	return &MultiDeleteRulepackCoreInternalServerError{}
}

/*MultiDeleteRulepackCoreInternalServerError handles this case with default header values.

Internal Server Error
*/
type MultiDeleteRulepackCoreInternalServerError struct {
	Payload *models.APIResult
}

func (o *MultiDeleteRulepackCoreInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /coreRulepacks][%d] multiDeleteRulepackCoreInternalServerError  %+v", 500, o.Payload)
}

func (o *MultiDeleteRulepackCoreInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MultiDeleteRulepackCoreInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}