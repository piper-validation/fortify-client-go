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

// UpdateCloudPoolReader is a Reader for the UpdateCloudPool structure.
type UpdateCloudPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateCloudPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateCloudPoolOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateCloudPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUpdateCloudPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateCloudPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateCloudPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUpdateCloudPoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateCloudPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateCloudPoolOK creates a UpdateCloudPoolOK with default headers values
func NewUpdateCloudPoolOK() *UpdateCloudPoolOK {
	return &UpdateCloudPoolOK{}
}

/*UpdateCloudPoolOK handles this case with default header values.

OK
*/
type UpdateCloudPoolOK struct {
	Payload *models.APIResultCloudPool
}

func (o *UpdateCloudPoolOK) Error() string {
	return fmt.Sprintf("[PUT /cloudpools/{uuid}][%d] updateCloudPoolOK  %+v", 200, o.Payload)
}

func (o *UpdateCloudPoolOK) GetPayload() *models.APIResultCloudPool {
	return o.Payload
}

func (o *UpdateCloudPoolOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultCloudPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudPoolBadRequest creates a UpdateCloudPoolBadRequest with default headers values
func NewUpdateCloudPoolBadRequest() *UpdateCloudPoolBadRequest {
	return &UpdateCloudPoolBadRequest{}
}

/*UpdateCloudPoolBadRequest handles this case with default header values.

Bad Request
*/
type UpdateCloudPoolBadRequest struct {
	Payload *models.APIResult
}

func (o *UpdateCloudPoolBadRequest) Error() string {
	return fmt.Sprintf("[PUT /cloudpools/{uuid}][%d] updateCloudPoolBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateCloudPoolBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCloudPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudPoolUnauthorized creates a UpdateCloudPoolUnauthorized with default headers values
func NewUpdateCloudPoolUnauthorized() *UpdateCloudPoolUnauthorized {
	return &UpdateCloudPoolUnauthorized{}
}

/*UpdateCloudPoolUnauthorized handles this case with default header values.

Unauthorized
*/
type UpdateCloudPoolUnauthorized struct {
	Payload *models.APIResult
}

func (o *UpdateCloudPoolUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /cloudpools/{uuid}][%d] updateCloudPoolUnauthorized  %+v", 401, o.Payload)
}

func (o *UpdateCloudPoolUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCloudPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudPoolForbidden creates a UpdateCloudPoolForbidden with default headers values
func NewUpdateCloudPoolForbidden() *UpdateCloudPoolForbidden {
	return &UpdateCloudPoolForbidden{}
}

/*UpdateCloudPoolForbidden handles this case with default header values.

Forbidden
*/
type UpdateCloudPoolForbidden struct {
	Payload *models.APIResult
}

func (o *UpdateCloudPoolForbidden) Error() string {
	return fmt.Sprintf("[PUT /cloudpools/{uuid}][%d] updateCloudPoolForbidden  %+v", 403, o.Payload)
}

func (o *UpdateCloudPoolForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCloudPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudPoolNotFound creates a UpdateCloudPoolNotFound with default headers values
func NewUpdateCloudPoolNotFound() *UpdateCloudPoolNotFound {
	return &UpdateCloudPoolNotFound{}
}

/*UpdateCloudPoolNotFound handles this case with default header values.

Not Found
*/
type UpdateCloudPoolNotFound struct {
	Payload *models.APIResult
}

func (o *UpdateCloudPoolNotFound) Error() string {
	return fmt.Sprintf("[PUT /cloudpools/{uuid}][%d] updateCloudPoolNotFound  %+v", 404, o.Payload)
}

func (o *UpdateCloudPoolNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCloudPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudPoolConflict creates a UpdateCloudPoolConflict with default headers values
func NewUpdateCloudPoolConflict() *UpdateCloudPoolConflict {
	return &UpdateCloudPoolConflict{}
}

/*UpdateCloudPoolConflict handles this case with default header values.

Conflict
*/
type UpdateCloudPoolConflict struct {
	Payload *models.APIResult
}

func (o *UpdateCloudPoolConflict) Error() string {
	return fmt.Sprintf("[PUT /cloudpools/{uuid}][%d] updateCloudPoolConflict  %+v", 409, o.Payload)
}

func (o *UpdateCloudPoolConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCloudPoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateCloudPoolInternalServerError creates a UpdateCloudPoolInternalServerError with default headers values
func NewUpdateCloudPoolInternalServerError() *UpdateCloudPoolInternalServerError {
	return &UpdateCloudPoolInternalServerError{}
}

/*UpdateCloudPoolInternalServerError handles this case with default header values.

Internal Server Error
*/
type UpdateCloudPoolInternalServerError struct {
	Payload *models.APIResult
}

func (o *UpdateCloudPoolInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /cloudpools/{uuid}][%d] updateCloudPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *UpdateCloudPoolInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UpdateCloudPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
