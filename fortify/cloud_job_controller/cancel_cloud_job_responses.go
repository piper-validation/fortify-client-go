// Code generated by go-swagger; DO NOT EDIT.

package cloud_job_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// CancelCloudJobReader is a Reader for the CancelCloudJob structure.
type CancelCloudJobReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CancelCloudJobReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewCancelCloudJobOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCancelCloudJobBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCancelCloudJobUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCancelCloudJobForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCancelCloudJobNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCancelCloudJobConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCancelCloudJobInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCancelCloudJobOK creates a CancelCloudJobOK with default headers values
func NewCancelCloudJobOK() *CancelCloudJobOK {
	return &CancelCloudJobOK{}
}

/*CancelCloudJobOK handles this case with default header values.

OK
*/
type CancelCloudJobOK struct {
	Payload *models.APIResultVoid
}

func (o *CancelCloudJobOK) Error() string {
	return fmt.Sprintf("[POST /cloudjobs/action/cancel][%d] cancelCloudJobOK  %+v", 200, o.Payload)
}

func (o *CancelCloudJobOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *CancelCloudJobOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelCloudJobBadRequest creates a CancelCloudJobBadRequest with default headers values
func NewCancelCloudJobBadRequest() *CancelCloudJobBadRequest {
	return &CancelCloudJobBadRequest{}
}

/*CancelCloudJobBadRequest handles this case with default header values.

Bad Request
*/
type CancelCloudJobBadRequest struct {
	Payload *models.APIResult
}

func (o *CancelCloudJobBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloudjobs/action/cancel][%d] cancelCloudJobBadRequest  %+v", 400, o.Payload)
}

func (o *CancelCloudJobBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelCloudJobBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelCloudJobUnauthorized creates a CancelCloudJobUnauthorized with default headers values
func NewCancelCloudJobUnauthorized() *CancelCloudJobUnauthorized {
	return &CancelCloudJobUnauthorized{}
}

/*CancelCloudJobUnauthorized handles this case with default header values.

Unauthorized
*/
type CancelCloudJobUnauthorized struct {
	Payload *models.APIResult
}

func (o *CancelCloudJobUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cloudjobs/action/cancel][%d] cancelCloudJobUnauthorized  %+v", 401, o.Payload)
}

func (o *CancelCloudJobUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelCloudJobUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelCloudJobForbidden creates a CancelCloudJobForbidden with default headers values
func NewCancelCloudJobForbidden() *CancelCloudJobForbidden {
	return &CancelCloudJobForbidden{}
}

/*CancelCloudJobForbidden handles this case with default header values.

Forbidden
*/
type CancelCloudJobForbidden struct {
	Payload *models.APIResult
}

func (o *CancelCloudJobForbidden) Error() string {
	return fmt.Sprintf("[POST /cloudjobs/action/cancel][%d] cancelCloudJobForbidden  %+v", 403, o.Payload)
}

func (o *CancelCloudJobForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelCloudJobForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelCloudJobNotFound creates a CancelCloudJobNotFound with default headers values
func NewCancelCloudJobNotFound() *CancelCloudJobNotFound {
	return &CancelCloudJobNotFound{}
}

/*CancelCloudJobNotFound handles this case with default header values.

Not Found
*/
type CancelCloudJobNotFound struct {
	Payload *models.APIResult
}

func (o *CancelCloudJobNotFound) Error() string {
	return fmt.Sprintf("[POST /cloudjobs/action/cancel][%d] cancelCloudJobNotFound  %+v", 404, o.Payload)
}

func (o *CancelCloudJobNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelCloudJobNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelCloudJobConflict creates a CancelCloudJobConflict with default headers values
func NewCancelCloudJobConflict() *CancelCloudJobConflict {
	return &CancelCloudJobConflict{}
}

/*CancelCloudJobConflict handles this case with default header values.

Conflict
*/
type CancelCloudJobConflict struct {
	Payload *models.APIResult
}

func (o *CancelCloudJobConflict) Error() string {
	return fmt.Sprintf("[POST /cloudjobs/action/cancel][%d] cancelCloudJobConflict  %+v", 409, o.Payload)
}

func (o *CancelCloudJobConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelCloudJobConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCancelCloudJobInternalServerError creates a CancelCloudJobInternalServerError with default headers values
func NewCancelCloudJobInternalServerError() *CancelCloudJobInternalServerError {
	return &CancelCloudJobInternalServerError{}
}

/*CancelCloudJobInternalServerError handles this case with default header values.

Internal Server Error
*/
type CancelCloudJobInternalServerError struct {
	Payload *models.APIResult
}

func (o *CancelCloudJobInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloudjobs/action/cancel][%d] cancelCloudJobInternalServerError  %+v", 500, o.Payload)
}

func (o *CancelCloudJobInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CancelCloudJobInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
