// Code generated by go-swagger; DO NOT EDIT.

package cloud_worker_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListCloudWorkerReader is a Reader for the ListCloudWorker structure.
type ListCloudWorkerReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCloudWorkerReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCloudWorkerOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCloudWorkerBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCloudWorkerUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCloudWorkerForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCloudWorkerNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListCloudWorkerConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCloudWorkerInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCloudWorkerOK creates a ListCloudWorkerOK with default headers values
func NewListCloudWorkerOK() *ListCloudWorkerOK {
	return &ListCloudWorkerOK{}
}

/*ListCloudWorkerOK handles this case with default header values.

OK
*/
type ListCloudWorkerOK struct {
	Payload *models.APIResultListCloudWorker
}

func (o *ListCloudWorkerOK) Error() string {
	return fmt.Sprintf("[GET /cloudworkers][%d] listCloudWorkerOK  %+v", 200, o.Payload)
}

func (o *ListCloudWorkerOK) GetPayload() *models.APIResultListCloudWorker {
	return o.Payload
}

func (o *ListCloudWorkerOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListCloudWorker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerBadRequest creates a ListCloudWorkerBadRequest with default headers values
func NewListCloudWorkerBadRequest() *ListCloudWorkerBadRequest {
	return &ListCloudWorkerBadRequest{}
}

/*ListCloudWorkerBadRequest handles this case with default header values.

Bad Request
*/
type ListCloudWorkerBadRequest struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudworkers][%d] listCloudWorkerBadRequest  %+v", 400, o.Payload)
}

func (o *ListCloudWorkerBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerUnauthorized creates a ListCloudWorkerUnauthorized with default headers values
func NewListCloudWorkerUnauthorized() *ListCloudWorkerUnauthorized {
	return &ListCloudWorkerUnauthorized{}
}

/*ListCloudWorkerUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCloudWorkerUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cloudworkers][%d] listCloudWorkerUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCloudWorkerUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerForbidden creates a ListCloudWorkerForbidden with default headers values
func NewListCloudWorkerForbidden() *ListCloudWorkerForbidden {
	return &ListCloudWorkerForbidden{}
}

/*ListCloudWorkerForbidden handles this case with default header values.

Forbidden
*/
type ListCloudWorkerForbidden struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerForbidden) Error() string {
	return fmt.Sprintf("[GET /cloudworkers][%d] listCloudWorkerForbidden  %+v", 403, o.Payload)
}

func (o *ListCloudWorkerForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerNotFound creates a ListCloudWorkerNotFound with default headers values
func NewListCloudWorkerNotFound() *ListCloudWorkerNotFound {
	return &ListCloudWorkerNotFound{}
}

/*ListCloudWorkerNotFound handles this case with default header values.

Not Found
*/
type ListCloudWorkerNotFound struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerNotFound) Error() string {
	return fmt.Sprintf("[GET /cloudworkers][%d] listCloudWorkerNotFound  %+v", 404, o.Payload)
}

func (o *ListCloudWorkerNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerConflict creates a ListCloudWorkerConflict with default headers values
func NewListCloudWorkerConflict() *ListCloudWorkerConflict {
	return &ListCloudWorkerConflict{}
}

/*ListCloudWorkerConflict handles this case with default header values.

Conflict
*/
type ListCloudWorkerConflict struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerConflict) Error() string {
	return fmt.Sprintf("[GET /cloudworkers][%d] listCloudWorkerConflict  %+v", 409, o.Payload)
}

func (o *ListCloudWorkerConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerInternalServerError creates a ListCloudWorkerInternalServerError with default headers values
func NewListCloudWorkerInternalServerError() *ListCloudWorkerInternalServerError {
	return &ListCloudWorkerInternalServerError{}
}

/*ListCloudWorkerInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListCloudWorkerInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudworkers][%d] listCloudWorkerInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCloudWorkerInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}