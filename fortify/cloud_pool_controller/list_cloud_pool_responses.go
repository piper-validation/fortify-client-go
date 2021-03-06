// Code generated by go-swagger; DO NOT EDIT.

package cloud_pool_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListCloudPoolReader is a Reader for the ListCloudPool structure.
type ListCloudPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListCloudPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListCloudPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListCloudPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListCloudPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListCloudPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListCloudPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListCloudPoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListCloudPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListCloudPoolOK creates a ListCloudPoolOK with default headers values
func NewListCloudPoolOK() *ListCloudPoolOK {
	return &ListCloudPoolOK{}
}

/*ListCloudPoolOK handles this case with default header values.

OK
*/
type ListCloudPoolOK struct {
	Payload *models.APIResultListCloudPool
}

func (o *ListCloudPoolOK) Error() string {
	return fmt.Sprintf("[GET /cloudpools][%d] listCloudPoolOK  %+v", 200, o.Payload)
}

func (o *ListCloudPoolOK) GetPayload() *models.APIResultListCloudPool {
	return o.Payload
}

func (o *ListCloudPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListCloudPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudPoolBadRequest creates a ListCloudPoolBadRequest with default headers values
func NewListCloudPoolBadRequest() *ListCloudPoolBadRequest {
	return &ListCloudPoolBadRequest{}
}

/*ListCloudPoolBadRequest handles this case with default header values.

Bad Request
*/
type ListCloudPoolBadRequest struct {
	Payload *models.APIResult
}

func (o *ListCloudPoolBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudpools][%d] listCloudPoolBadRequest  %+v", 400, o.Payload)
}

func (o *ListCloudPoolBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudPoolUnauthorized creates a ListCloudPoolUnauthorized with default headers values
func NewListCloudPoolUnauthorized() *ListCloudPoolUnauthorized {
	return &ListCloudPoolUnauthorized{}
}

/*ListCloudPoolUnauthorized handles this case with default header values.

Unauthorized
*/
type ListCloudPoolUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListCloudPoolUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cloudpools][%d] listCloudPoolUnauthorized  %+v", 401, o.Payload)
}

func (o *ListCloudPoolUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudPoolForbidden creates a ListCloudPoolForbidden with default headers values
func NewListCloudPoolForbidden() *ListCloudPoolForbidden {
	return &ListCloudPoolForbidden{}
}

/*ListCloudPoolForbidden handles this case with default header values.

Forbidden
*/
type ListCloudPoolForbidden struct {
	Payload *models.APIResult
}

func (o *ListCloudPoolForbidden) Error() string {
	return fmt.Sprintf("[GET /cloudpools][%d] listCloudPoolForbidden  %+v", 403, o.Payload)
}

func (o *ListCloudPoolForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudPoolNotFound creates a ListCloudPoolNotFound with default headers values
func NewListCloudPoolNotFound() *ListCloudPoolNotFound {
	return &ListCloudPoolNotFound{}
}

/*ListCloudPoolNotFound handles this case with default header values.

Not Found
*/
type ListCloudPoolNotFound struct {
	Payload *models.APIResult
}

func (o *ListCloudPoolNotFound) Error() string {
	return fmt.Sprintf("[GET /cloudpools][%d] listCloudPoolNotFound  %+v", 404, o.Payload)
}

func (o *ListCloudPoolNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudPoolConflict creates a ListCloudPoolConflict with default headers values
func NewListCloudPoolConflict() *ListCloudPoolConflict {
	return &ListCloudPoolConflict{}
}

/*ListCloudPoolConflict handles this case with default header values.

Conflict
*/
type ListCloudPoolConflict struct {
	Payload *models.APIResult
}

func (o *ListCloudPoolConflict) Error() string {
	return fmt.Sprintf("[GET /cloudpools][%d] listCloudPoolConflict  %+v", 409, o.Payload)
}

func (o *ListCloudPoolConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudPoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListCloudPoolInternalServerError creates a ListCloudPoolInternalServerError with default headers values
func NewListCloudPoolInternalServerError() *ListCloudPoolInternalServerError {
	return &ListCloudPoolInternalServerError{}
}

/*ListCloudPoolInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListCloudPoolInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListCloudPoolInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudpools][%d] listCloudPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *ListCloudPoolInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListCloudPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
