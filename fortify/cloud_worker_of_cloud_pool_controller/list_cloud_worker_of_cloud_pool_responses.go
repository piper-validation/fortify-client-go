// Code generated by go-swagger; DO NOT EDIT.

package cloud_worker_of_cloud_pool_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListCloudWorkerOfCloudPoolReader is a Reader for the ListCloudWorkerOfCloudPool structure.
type ListCloudWorkerOfCloudPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCloudWorkerOfCloudPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCloudWorkerOfCloudPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCloudWorkerOfCloudPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCloudWorkerOfCloudPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCloudWorkerOfCloudPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCloudWorkerOfCloudPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListCloudWorkerOfCloudPoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCloudWorkerOfCloudPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCloudWorkerOfCloudPoolOK creates a ListCloudWorkerOfCloudPoolOK with default headers values
func NewListCloudWorkerOfCloudPoolOK() *ListCloudWorkerOfCloudPoolOK {
	return &ListCloudWorkerOfCloudPoolOK{}
}

/*ListCloudWorkerOfCloudPoolOK handles this case with default header values.

OK
*/
type ListCloudWorkerOfCloudPoolOK struct {
	Payload *models.APIResultListCloudWorker
}

func (o *ListCloudWorkerOfCloudPoolOK) Error() string {
	return fmt.Sprintf("[GET /cloudpools/{parentId}/workers][%d] listCloudWorkerOfCloudPoolOK  %+v", 200, o.Payload)
}

func (o *ListCloudWorkerOfCloudPoolOK) GetPayload() *models.APIResultListCloudWorker {
	return o.Payload
}

func (o *ListCloudWorkerOfCloudPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListCloudWorker)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerOfCloudPoolBadRequest creates a ListCloudWorkerOfCloudPoolBadRequest with default headers values
func NewListCloudWorkerOfCloudPoolBadRequest() *ListCloudWorkerOfCloudPoolBadRequest {
	return &ListCloudWorkerOfCloudPoolBadRequest{}
}

/*ListCloudWorkerOfCloudPoolBadRequest handles this case with default header values.

Bad Request
*/
type ListCloudWorkerOfCloudPoolBadRequest struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerOfCloudPoolBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudpools/{parentId}/workers][%d] listCloudWorkerOfCloudPoolBadRequest  %+v", 400, o.Payload)
}

func (o *ListCloudWorkerOfCloudPoolBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerOfCloudPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerOfCloudPoolUnauthorized creates a ListCloudWorkerOfCloudPoolUnauthorized with default headers values
func NewListCloudWorkerOfCloudPoolUnauthorized() *ListCloudWorkerOfCloudPoolUnauthorized {
	return &ListCloudWorkerOfCloudPoolUnauthorized{}
}

/*ListCloudWorkerOfCloudPoolUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCloudWorkerOfCloudPoolUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerOfCloudPoolUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cloudpools/{parentId}/workers][%d] listCloudWorkerOfCloudPoolUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCloudWorkerOfCloudPoolUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerOfCloudPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerOfCloudPoolForbidden creates a ListCloudWorkerOfCloudPoolForbidden with default headers values
func NewListCloudWorkerOfCloudPoolForbidden() *ListCloudWorkerOfCloudPoolForbidden {
	return &ListCloudWorkerOfCloudPoolForbidden{}
}

/*ListCloudWorkerOfCloudPoolForbidden handles this case with default header values.

Forbidden
*/
type ListCloudWorkerOfCloudPoolForbidden struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerOfCloudPoolForbidden) Error() string {
	return fmt.Sprintf("[GET /cloudpools/{parentId}/workers][%d] listCloudWorkerOfCloudPoolForbidden  %+v", 403, o.Payload)
}

func (o *ListCloudWorkerOfCloudPoolForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerOfCloudPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerOfCloudPoolNotFound creates a ListCloudWorkerOfCloudPoolNotFound with default headers values
func NewListCloudWorkerOfCloudPoolNotFound() *ListCloudWorkerOfCloudPoolNotFound {
	return &ListCloudWorkerOfCloudPoolNotFound{}
}

/*ListCloudWorkerOfCloudPoolNotFound handles this case with default header values.

Not Found
*/
type ListCloudWorkerOfCloudPoolNotFound struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerOfCloudPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /cloudpools/{parentId}/workers][%d] listCloudWorkerOfCloudPoolNotFound  %+v", 404, o.Payload)
}

func (o *ListCloudWorkerOfCloudPoolNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerOfCloudPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerOfCloudPoolConflict creates a ListCloudWorkerOfCloudPoolConflict with default headers values
func NewListCloudWorkerOfCloudPoolConflict() *ListCloudWorkerOfCloudPoolConflict {
	return &ListCloudWorkerOfCloudPoolConflict{}
}

/*ListCloudWorkerOfCloudPoolConflict handles this case with default header values.

Conflict
*/
type ListCloudWorkerOfCloudPoolConflict struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerOfCloudPoolConflict) Error() string {
	return fmt.Sprintf("[GET /cloudpools/{parentId}/workers][%d] listCloudWorkerOfCloudPoolConflict  %+v", 409, o.Payload)
}

func (o *ListCloudWorkerOfCloudPoolConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerOfCloudPoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudWorkerOfCloudPoolInternalServerError creates a ListCloudWorkerOfCloudPoolInternalServerError with default headers values
func NewListCloudWorkerOfCloudPoolInternalServerError() *ListCloudWorkerOfCloudPoolInternalServerError {
	return &ListCloudWorkerOfCloudPoolInternalServerError{}
}

/*ListCloudWorkerOfCloudPoolInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListCloudWorkerOfCloudPoolInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListCloudWorkerOfCloudPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudpools/{parentId}/workers][%d] listCloudWorkerOfCloudPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCloudWorkerOfCloudPoolInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudWorkerOfCloudPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
