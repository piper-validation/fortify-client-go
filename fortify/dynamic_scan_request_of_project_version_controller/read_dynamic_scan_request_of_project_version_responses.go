// Code generated by go-swagger; DO NOT EDIT.

package dynamic_scan_request_of_project_version_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadDynamicScanRequestOfProjectVersionReader is a Reader for the ReadDynamicScanRequestOfProjectVersion structure.
type ReadDynamicScanRequestOfProjectVersionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadDynamicScanRequestOfProjectVersionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadDynamicScanRequestOfProjectVersionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadDynamicScanRequestOfProjectVersionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadDynamicScanRequestOfProjectVersionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadDynamicScanRequestOfProjectVersionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadDynamicScanRequestOfProjectVersionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadDynamicScanRequestOfProjectVersionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadDynamicScanRequestOfProjectVersionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadDynamicScanRequestOfProjectVersionOK creates a ReadDynamicScanRequestOfProjectVersionOK with default headers values
func NewReadDynamicScanRequestOfProjectVersionOK() *ReadDynamicScanRequestOfProjectVersionOK {
	return &ReadDynamicScanRequestOfProjectVersionOK{}
}

/*ReadDynamicScanRequestOfProjectVersionOK handles this case with default header values.

OK
*/
type ReadDynamicScanRequestOfProjectVersionOK struct {
	Payload *models.APIResultDynamicScanRequest
}

func (o *ReadDynamicScanRequestOfProjectVersionOK) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] readDynamicScanRequestOfProjectVersionOK  %+v", 200, o.Payload)
}

func (o *ReadDynamicScanRequestOfProjectVersionOK) GetPayload() *models.APIResultDynamicScanRequest {
	return o.Payload
}

func (o *ReadDynamicScanRequestOfProjectVersionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultDynamicScanRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDynamicScanRequestOfProjectVersionBadRequest creates a ReadDynamicScanRequestOfProjectVersionBadRequest with default headers values
func NewReadDynamicScanRequestOfProjectVersionBadRequest() *ReadDynamicScanRequestOfProjectVersionBadRequest {
	return &ReadDynamicScanRequestOfProjectVersionBadRequest{}
}

/*ReadDynamicScanRequestOfProjectVersionBadRequest handles this case with default header values.

Bad Request
*/
type ReadDynamicScanRequestOfProjectVersionBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadDynamicScanRequestOfProjectVersionBadRequest) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] readDynamicScanRequestOfProjectVersionBadRequest  %+v", 400, o.Payload)
}

func (o *ReadDynamicScanRequestOfProjectVersionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadDynamicScanRequestOfProjectVersionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDynamicScanRequestOfProjectVersionUnauthorized creates a ReadDynamicScanRequestOfProjectVersionUnauthorized with default headers values
func NewReadDynamicScanRequestOfProjectVersionUnauthorized() *ReadDynamicScanRequestOfProjectVersionUnauthorized {
	return &ReadDynamicScanRequestOfProjectVersionUnauthorized{}
}

/*ReadDynamicScanRequestOfProjectVersionUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadDynamicScanRequestOfProjectVersionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadDynamicScanRequestOfProjectVersionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] readDynamicScanRequestOfProjectVersionUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadDynamicScanRequestOfProjectVersionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadDynamicScanRequestOfProjectVersionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDynamicScanRequestOfProjectVersionForbidden creates a ReadDynamicScanRequestOfProjectVersionForbidden with default headers values
func NewReadDynamicScanRequestOfProjectVersionForbidden() *ReadDynamicScanRequestOfProjectVersionForbidden {
	return &ReadDynamicScanRequestOfProjectVersionForbidden{}
}

/*ReadDynamicScanRequestOfProjectVersionForbidden handles this case with default header values.

Forbidden
*/
type ReadDynamicScanRequestOfProjectVersionForbidden struct {
	Payload *models.APIResult
}

func (o *ReadDynamicScanRequestOfProjectVersionForbidden) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] readDynamicScanRequestOfProjectVersionForbidden  %+v", 403, o.Payload)
}

func (o *ReadDynamicScanRequestOfProjectVersionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadDynamicScanRequestOfProjectVersionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDynamicScanRequestOfProjectVersionNotFound creates a ReadDynamicScanRequestOfProjectVersionNotFound with default headers values
func NewReadDynamicScanRequestOfProjectVersionNotFound() *ReadDynamicScanRequestOfProjectVersionNotFound {
	return &ReadDynamicScanRequestOfProjectVersionNotFound{}
}

/*ReadDynamicScanRequestOfProjectVersionNotFound handles this case with default header values.

Not Found
*/
type ReadDynamicScanRequestOfProjectVersionNotFound struct {
	Payload *models.APIResult
}

func (o *ReadDynamicScanRequestOfProjectVersionNotFound) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] readDynamicScanRequestOfProjectVersionNotFound  %+v", 404, o.Payload)
}

func (o *ReadDynamicScanRequestOfProjectVersionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadDynamicScanRequestOfProjectVersionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDynamicScanRequestOfProjectVersionConflict creates a ReadDynamicScanRequestOfProjectVersionConflict with default headers values
func NewReadDynamicScanRequestOfProjectVersionConflict() *ReadDynamicScanRequestOfProjectVersionConflict {
	return &ReadDynamicScanRequestOfProjectVersionConflict{}
}

/*ReadDynamicScanRequestOfProjectVersionConflict handles this case with default header values.

Conflict
*/
type ReadDynamicScanRequestOfProjectVersionConflict struct {
	Payload *models.APIResult
}

func (o *ReadDynamicScanRequestOfProjectVersionConflict) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] readDynamicScanRequestOfProjectVersionConflict  %+v", 409, o.Payload)
}

func (o *ReadDynamicScanRequestOfProjectVersionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadDynamicScanRequestOfProjectVersionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadDynamicScanRequestOfProjectVersionInternalServerError creates a ReadDynamicScanRequestOfProjectVersionInternalServerError with default headers values
func NewReadDynamicScanRequestOfProjectVersionInternalServerError() *ReadDynamicScanRequestOfProjectVersionInternalServerError {
	return &ReadDynamicScanRequestOfProjectVersionInternalServerError{}
}

/*ReadDynamicScanRequestOfProjectVersionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadDynamicScanRequestOfProjectVersionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadDynamicScanRequestOfProjectVersionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /projectVersions/{parentId}/dynamicScanRequests/{id}][%d] readDynamicScanRequestOfProjectVersionInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadDynamicScanRequestOfProjectVersionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadDynamicScanRequestOfProjectVersionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
