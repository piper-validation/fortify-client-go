// Code generated by go-swagger; DO NOT EDIT.

package unassigned_cloud_worker_of_cloud_pool_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListUnassignedCloudWorkerOfCloudPoolReader is a Reader for the ListUnassignedCloudWorkerOfCloudPool structure.
type ListUnassignedCloudWorkerOfCloudPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListUnassignedCloudWorkerOfCloudPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListUnassignedCloudWorkerOfCloudPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListUnassignedCloudWorkerOfCloudPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListUnassignedCloudWorkerOfCloudPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListUnassignedCloudWorkerOfCloudPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListUnassignedCloudWorkerOfCloudPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListUnassignedCloudWorkerOfCloudPoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListUnassignedCloudWorkerOfCloudPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListUnassignedCloudWorkerOfCloudPoolOK creates a ListUnassignedCloudWorkerOfCloudPoolOK with default headers values
func NewListUnassignedCloudWorkerOfCloudPoolOK() *ListUnassignedCloudWorkerOfCloudPoolOK {
	return &ListUnassignedCloudWorkerOfCloudPoolOK{}
}

/*ListUnassignedCloudWorkerOfCloudPoolOK handles this case with default header values.

OK
*/
type ListUnassignedCloudWorkerOfCloudPoolOK struct {
	Payload *models.APIResultListCloudWorker
}

func (o *ListUnassignedCloudWorkerOfCloudPoolOK) Error() string {
	return fmt.Sprintf("[GET /cloudpools/disabledWorkers][%d] listUnassignedCloudWorkerOfCloudPoolOK  %+v", 200, o.Payload)
}

func (o *ListUnassignedCloudWorkerOfCloudPoolOK) GetPayload() *models.APIResultListCloudWorker {
	return o.Payload
}

func (o *ListUnassignedCloudWorkerOfCloudPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListCloudWorker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUnassignedCloudWorkerOfCloudPoolBadRequest creates a ListUnassignedCloudWorkerOfCloudPoolBadRequest with default headers values
func NewListUnassignedCloudWorkerOfCloudPoolBadRequest() *ListUnassignedCloudWorkerOfCloudPoolBadRequest {
	return &ListUnassignedCloudWorkerOfCloudPoolBadRequest{}
}

/*ListUnassignedCloudWorkerOfCloudPoolBadRequest handles this case with default header values.

Bad Request
*/
type ListUnassignedCloudWorkerOfCloudPoolBadRequest struct {
	Payload *models.APIResult
}

func (o *ListUnassignedCloudWorkerOfCloudPoolBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudpools/disabledWorkers][%d] listUnassignedCloudWorkerOfCloudPoolBadRequest  %+v", 400, o.Payload)
}

func (o *ListUnassignedCloudWorkerOfCloudPoolBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListUnassignedCloudWorkerOfCloudPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUnassignedCloudWorkerOfCloudPoolUnauthorized creates a ListUnassignedCloudWorkerOfCloudPoolUnauthorized with default headers values
func NewListUnassignedCloudWorkerOfCloudPoolUnauthorized() *ListUnassignedCloudWorkerOfCloudPoolUnauthorized {
	return &ListUnassignedCloudWorkerOfCloudPoolUnauthorized{}
}

/*ListUnassignedCloudWorkerOfCloudPoolUnauthorized handles this case with default header values.

Unauthorized
*/
type ListUnassignedCloudWorkerOfCloudPoolUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListUnassignedCloudWorkerOfCloudPoolUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cloudpools/disabledWorkers][%d] listUnassignedCloudWorkerOfCloudPoolUnauthorized  %+v", 401, o.Payload)
}

func (o *ListUnassignedCloudWorkerOfCloudPoolUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListUnassignedCloudWorkerOfCloudPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUnassignedCloudWorkerOfCloudPoolForbidden creates a ListUnassignedCloudWorkerOfCloudPoolForbidden with default headers values
func NewListUnassignedCloudWorkerOfCloudPoolForbidden() *ListUnassignedCloudWorkerOfCloudPoolForbidden {
	return &ListUnassignedCloudWorkerOfCloudPoolForbidden{}
}

/*ListUnassignedCloudWorkerOfCloudPoolForbidden handles this case with default header values.

Forbidden
*/
type ListUnassignedCloudWorkerOfCloudPoolForbidden struct {
	Payload *models.APIResult
}

func (o *ListUnassignedCloudWorkerOfCloudPoolForbidden) Error() string {
	return fmt.Sprintf("[GET /cloudpools/disabledWorkers][%d] listUnassignedCloudWorkerOfCloudPoolForbidden  %+v", 403, o.Payload)
}

func (o *ListUnassignedCloudWorkerOfCloudPoolForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListUnassignedCloudWorkerOfCloudPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUnassignedCloudWorkerOfCloudPoolNotFound creates a ListUnassignedCloudWorkerOfCloudPoolNotFound with default headers values
func NewListUnassignedCloudWorkerOfCloudPoolNotFound() *ListUnassignedCloudWorkerOfCloudPoolNotFound {
	return &ListUnassignedCloudWorkerOfCloudPoolNotFound{}
}

/*ListUnassignedCloudWorkerOfCloudPoolNotFound handles this case with default header values.

Not Found
*/
type ListUnassignedCloudWorkerOfCloudPoolNotFound struct {
	Payload *models.APIResult
}

func (o *ListUnassignedCloudWorkerOfCloudPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /cloudpools/disabledWorkers][%d] listUnassignedCloudWorkerOfCloudPoolNotFound  %+v", 404, o.Payload)
}

func (o *ListUnassignedCloudWorkerOfCloudPoolNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListUnassignedCloudWorkerOfCloudPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUnassignedCloudWorkerOfCloudPoolConflict creates a ListUnassignedCloudWorkerOfCloudPoolConflict with default headers values
func NewListUnassignedCloudWorkerOfCloudPoolConflict() *ListUnassignedCloudWorkerOfCloudPoolConflict {
	return &ListUnassignedCloudWorkerOfCloudPoolConflict{}
}

/*ListUnassignedCloudWorkerOfCloudPoolConflict handles this case with default header values.

Conflict
*/
type ListUnassignedCloudWorkerOfCloudPoolConflict struct {
	Payload *models.APIResult
}

func (o *ListUnassignedCloudWorkerOfCloudPoolConflict) Error() string {
	return fmt.Sprintf("[GET /cloudpools/disabledWorkers][%d] listUnassignedCloudWorkerOfCloudPoolConflict  %+v", 409, o.Payload)
}

func (o *ListUnassignedCloudWorkerOfCloudPoolConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListUnassignedCloudWorkerOfCloudPoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListUnassignedCloudWorkerOfCloudPoolInternalServerError creates a ListUnassignedCloudWorkerOfCloudPoolInternalServerError with default headers values
func NewListUnassignedCloudWorkerOfCloudPoolInternalServerError() *ListUnassignedCloudWorkerOfCloudPoolInternalServerError {
	return &ListUnassignedCloudWorkerOfCloudPoolInternalServerError{}
}

/*ListUnassignedCloudWorkerOfCloudPoolInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListUnassignedCloudWorkerOfCloudPoolInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListUnassignedCloudWorkerOfCloudPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudpools/disabledWorkers][%d] listUnassignedCloudWorkerOfCloudPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *ListUnassignedCloudWorkerOfCloudPoolInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListUnassignedCloudWorkerOfCloudPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
