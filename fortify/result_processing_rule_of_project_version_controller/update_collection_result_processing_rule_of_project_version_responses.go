// Code generated by go-swagger; DO NOT EDIT.

package result_processing_rule_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// UpdateCollectionResultProcessingRuleOfProjectVersionReader is a Reader for the UpdateCollectionResultProcessingRuleOfProjectVersion structure.
type UpdateCollectionResultProcessingRuleOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCollectionResultProcessingRuleOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCollectionResultProcessingRuleOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCollectionResultProcessingRuleOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCollectionResultProcessingRuleOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCollectionResultProcessingRuleOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateCollectionResultProcessingRuleOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCollectionResultProcessingRuleOfProjectVersionOK creates a UpdateCollectionResultProcessingRuleOfProjectVersionOK with default headers values
func NewUpdateCollectionResultProcessingRuleOfProjectVersionOK() *UpdateCollectionResultProcessingRuleOfProjectVersionOK {
	return &UpdateCollectionResultProcessingRuleOfProjectVersionOK{}
}

/*UpdateCollectionResultProcessingRuleOfProjectVersionOK handles this case with default header values.

OK
*/
type UpdateCollectionResultProcessingRuleOfProjectVersionOK struct {
	Payload *models.APIResultListResultProcessingRule
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/resultProcessingRules][%d] updateCollectionResultProcessingRuleOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionOK) GetPayload() *models.APIResultListResultProcessingRule {
	return o.Payload
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListResultProcessingRule)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectionResultProcessingRuleOfProjectVersionBadRequest creates a UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest with default headers values
func NewUpdateCollectionResultProcessingRuleOfProjectVersionBadRequest() *UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest {
	return &UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest{}
}

/*UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/resultProcessingRules][%d] updateCollectionResultProcessingRuleOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized creates a UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized with default headers values
func NewUpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized() *UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized {
	return &UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized{}
}

/*UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/resultProcessingRules][%d] updateCollectionResultProcessingRuleOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectionResultProcessingRuleOfProjectVersionForbidden creates a UpdateCollectionResultProcessingRuleOfProjectVersionForbidden with default headers values
func NewUpdateCollectionResultProcessingRuleOfProjectVersionForbidden() *UpdateCollectionResultProcessingRuleOfProjectVersionForbidden {
	return &UpdateCollectionResultProcessingRuleOfProjectVersionForbidden{}
}

/*UpdateCollectionResultProcessingRuleOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type UpdateCollectionResultProcessingRuleOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/resultProcessingRules][%d] updateCollectionResultProcessingRuleOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectionResultProcessingRuleOfProjectVersionNotFound creates a UpdateCollectionResultProcessingRuleOfProjectVersionNotFound with default headers values
func NewUpdateCollectionResultProcessingRuleOfProjectVersionNotFound() *UpdateCollectionResultProcessingRuleOfProjectVersionNotFound {
	return &UpdateCollectionResultProcessingRuleOfProjectVersionNotFound{}
}

/*UpdateCollectionResultProcessingRuleOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type UpdateCollectionResultProcessingRuleOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/resultProcessingRules][%d] updateCollectionResultProcessingRuleOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectionResultProcessingRuleOfProjectVersionConflict creates a UpdateCollectionResultProcessingRuleOfProjectVersionConflict with default headers values
func NewUpdateCollectionResultProcessingRuleOfProjectVersionConflict() *UpdateCollectionResultProcessingRuleOfProjectVersionConflict {
	return &UpdateCollectionResultProcessingRuleOfProjectVersionConflict{}
}

/*UpdateCollectionResultProcessingRuleOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type UpdateCollectionResultProcessingRuleOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/resultProcessingRules][%d] updateCollectionResultProcessingRuleOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError creates a UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError with default headers values
func NewUpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError() *UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError {
	return &UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError{}
}

/*UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /projectVersions/{parentId}/resultProcessingRules][%d] updateCollectionResultProcessingRuleOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCollectionResultProcessingRuleOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}