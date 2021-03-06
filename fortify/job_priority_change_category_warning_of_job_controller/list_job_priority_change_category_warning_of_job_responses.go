// Code generated by go-swagger; DO NOT EDIT.

package job_priority_change_category_warning_of_job_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListJobPriorityChangeCategoryWarningOfJobReader is a Reader for the ListJobPriorityChangeCategoryWarningOfJob structure.
type ListJobPriorityChangeCategoryWarningOfJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListJobPriorityChangeCategoryWarningOfJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListJobPriorityChangeCategoryWarningOfJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListJobPriorityChangeCategoryWarningOfJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListJobPriorityChangeCategoryWarningOfJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListJobPriorityChangeCategoryWarningOfJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListJobPriorityChangeCategoryWarningOfJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListJobPriorityChangeCategoryWarningOfJobConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListJobPriorityChangeCategoryWarningOfJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListJobPriorityChangeCategoryWarningOfJobOK creates a ListJobPriorityChangeCategoryWarningOfJobOK with default headers values
func NewListJobPriorityChangeCategoryWarningOfJobOK() *ListJobPriorityChangeCategoryWarningOfJobOK {
	return &ListJobPriorityChangeCategoryWarningOfJobOK{}
}

/*ListJobPriorityChangeCategoryWarningOfJobOK handles this case with default header values.

OK
*/
type ListJobPriorityChangeCategoryWarningOfJobOK struct {
	Payload *models.APIResultListJobPriorityChangeCategoryWarning
}

func (o *ListJobPriorityChangeCategoryWarningOfJobOK) Error() string {
	return fmt.Sprintf("[GET /jobs/{parentId}/warnings][%d] listJobPriorityChangeCategoryWarningOfJobOK  %+v", 200, o.Payload)
}

func (o *ListJobPriorityChangeCategoryWarningOfJobOK) GetPayload() *models.APIResultListJobPriorityChangeCategoryWarning {
	return o.Payload
}

func (o *ListJobPriorityChangeCategoryWarningOfJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListJobPriorityChangeCategoryWarning)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobPriorityChangeCategoryWarningOfJobBadRequest creates a ListJobPriorityChangeCategoryWarningOfJobBadRequest with default headers values
func NewListJobPriorityChangeCategoryWarningOfJobBadRequest() *ListJobPriorityChangeCategoryWarningOfJobBadRequest {
	return &ListJobPriorityChangeCategoryWarningOfJobBadRequest{}
}

/*ListJobPriorityChangeCategoryWarningOfJobBadRequest handles this case with default header values.

Bad Request
*/
type ListJobPriorityChangeCategoryWarningOfJobBadRequest struct {
	Payload *models.APIResult
}

func (o *ListJobPriorityChangeCategoryWarningOfJobBadRequest) Error() string {
	return fmt.Sprintf("[GET /jobs/{parentId}/warnings][%d] listJobPriorityChangeCategoryWarningOfJobBadRequest  %+v", 400, o.Payload)
}

func (o *ListJobPriorityChangeCategoryWarningOfJobBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListJobPriorityChangeCategoryWarningOfJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobPriorityChangeCategoryWarningOfJobUnauthorized creates a ListJobPriorityChangeCategoryWarningOfJobUnauthorized with default headers values
func NewListJobPriorityChangeCategoryWarningOfJobUnauthorized() *ListJobPriorityChangeCategoryWarningOfJobUnauthorized {
	return &ListJobPriorityChangeCategoryWarningOfJobUnauthorized{}
}

/*ListJobPriorityChangeCategoryWarningOfJobUnauthorized handles this case with default header values.

Unauthorized
*/
type ListJobPriorityChangeCategoryWarningOfJobUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListJobPriorityChangeCategoryWarningOfJobUnauthorized) Error() string {
	return fmt.Sprintf("[GET /jobs/{parentId}/warnings][%d] listJobPriorityChangeCategoryWarningOfJobUnauthorized  %+v", 401, o.Payload)
}

func (o *ListJobPriorityChangeCategoryWarningOfJobUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListJobPriorityChangeCategoryWarningOfJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobPriorityChangeCategoryWarningOfJobForbidden creates a ListJobPriorityChangeCategoryWarningOfJobForbidden with default headers values
func NewListJobPriorityChangeCategoryWarningOfJobForbidden() *ListJobPriorityChangeCategoryWarningOfJobForbidden {
	return &ListJobPriorityChangeCategoryWarningOfJobForbidden{}
}

/*ListJobPriorityChangeCategoryWarningOfJobForbidden handles this case with default header values.

Forbidden
*/
type ListJobPriorityChangeCategoryWarningOfJobForbidden struct {
	Payload *models.APIResult
}

func (o *ListJobPriorityChangeCategoryWarningOfJobForbidden) Error() string {
	return fmt.Sprintf("[GET /jobs/{parentId}/warnings][%d] listJobPriorityChangeCategoryWarningOfJobForbidden  %+v", 403, o.Payload)
}

func (o *ListJobPriorityChangeCategoryWarningOfJobForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListJobPriorityChangeCategoryWarningOfJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobPriorityChangeCategoryWarningOfJobNotFound creates a ListJobPriorityChangeCategoryWarningOfJobNotFound with default headers values
func NewListJobPriorityChangeCategoryWarningOfJobNotFound() *ListJobPriorityChangeCategoryWarningOfJobNotFound {
	return &ListJobPriorityChangeCategoryWarningOfJobNotFound{}
}

/*ListJobPriorityChangeCategoryWarningOfJobNotFound handles this case with default header values.

Not Found
*/
type ListJobPriorityChangeCategoryWarningOfJobNotFound struct {
	Payload *models.APIResult
}

func (o *ListJobPriorityChangeCategoryWarningOfJobNotFound) Error() string {
	return fmt.Sprintf("[GET /jobs/{parentId}/warnings][%d] listJobPriorityChangeCategoryWarningOfJobNotFound  %+v", 404, o.Payload)
}

func (o *ListJobPriorityChangeCategoryWarningOfJobNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListJobPriorityChangeCategoryWarningOfJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobPriorityChangeCategoryWarningOfJobConflict creates a ListJobPriorityChangeCategoryWarningOfJobConflict with default headers values
func NewListJobPriorityChangeCategoryWarningOfJobConflict() *ListJobPriorityChangeCategoryWarningOfJobConflict {
	return &ListJobPriorityChangeCategoryWarningOfJobConflict{}
}

/*ListJobPriorityChangeCategoryWarningOfJobConflict handles this case with default header values.

Conflict
*/
type ListJobPriorityChangeCategoryWarningOfJobConflict struct {
	Payload *models.APIResult
}

func (o *ListJobPriorityChangeCategoryWarningOfJobConflict) Error() string {
	return fmt.Sprintf("[GET /jobs/{parentId}/warnings][%d] listJobPriorityChangeCategoryWarningOfJobConflict  %+v", 409, o.Payload)
}

func (o *ListJobPriorityChangeCategoryWarningOfJobConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListJobPriorityChangeCategoryWarningOfJobConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListJobPriorityChangeCategoryWarningOfJobInternalServerError creates a ListJobPriorityChangeCategoryWarningOfJobInternalServerError with default headers values
func NewListJobPriorityChangeCategoryWarningOfJobInternalServerError() *ListJobPriorityChangeCategoryWarningOfJobInternalServerError {
	return &ListJobPriorityChangeCategoryWarningOfJobInternalServerError{}
}

/*ListJobPriorityChangeCategoryWarningOfJobInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListJobPriorityChangeCategoryWarningOfJobInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListJobPriorityChangeCategoryWarningOfJobInternalServerError) Error() string {
	return fmt.Sprintf("[GET /jobs/{parentId}/warnings][%d] listJobPriorityChangeCategoryWarningOfJobInternalServerError  %+v", 500, o.Payload)
}

func (o *ListJobPriorityChangeCategoryWarningOfJobInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListJobPriorityChangeCategoryWarningOfJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
