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

// CreateCloudPoolReader is a Reader for the CreateCloudPool structure.
type CreateCloudPoolReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreateCloudPoolReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewCreateCloudPoolCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewCreateCloudPoolBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewCreateCloudPoolUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewCreateCloudPoolForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewCreateCloudPoolNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewCreateCloudPoolConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCreateCloudPoolInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreateCloudPoolCreated creates a CreateCloudPoolCreated with default headers values
func NewCreateCloudPoolCreated() *CreateCloudPoolCreated {
	return &CreateCloudPoolCreated{}
}

/*CreateCloudPoolCreated handles this case with default header values.

Created
*/
type CreateCloudPoolCreated struct {
	Payload *models.APIResultCloudPool
}

func (o *CreateCloudPoolCreated) Error() string {
	return fmt.Sprintf("[POST /cloudpools][%d] createCloudPoolCreated  %+v", 201, o.Payload)
}

func (o *CreateCloudPoolCreated) GetPayload() *models.APIResultCloudPool {
	return o.Payload
}

func (o *CreateCloudPoolCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultCloudPool)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudPoolBadRequest creates a CreateCloudPoolBadRequest with default headers values
func NewCreateCloudPoolBadRequest() *CreateCloudPoolBadRequest {
	return &CreateCloudPoolBadRequest{}
}

/*CreateCloudPoolBadRequest handles this case with default header values.

Bad Request
*/
type CreateCloudPoolBadRequest struct {
	Payload *models.APIResult
}

func (o *CreateCloudPoolBadRequest) Error() string {
	return fmt.Sprintf("[POST /cloudpools][%d] createCloudPoolBadRequest  %+v", 400, o.Payload)
}

func (o *CreateCloudPoolBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCloudPoolBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudPoolUnauthorized creates a CreateCloudPoolUnauthorized with default headers values
func NewCreateCloudPoolUnauthorized() *CreateCloudPoolUnauthorized {
	return &CreateCloudPoolUnauthorized{}
}

/*CreateCloudPoolUnauthorized handles this case with default header values.

Unauthorized
*/
type CreateCloudPoolUnauthorized struct {
	Payload *models.APIResult
}

func (o *CreateCloudPoolUnauthorized) Error() string {
	return fmt.Sprintf("[POST /cloudpools][%d] createCloudPoolUnauthorized  %+v", 401, o.Payload)
}

func (o *CreateCloudPoolUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCloudPoolUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudPoolForbidden creates a CreateCloudPoolForbidden with default headers values
func NewCreateCloudPoolForbidden() *CreateCloudPoolForbidden {
	return &CreateCloudPoolForbidden{}
}

/*CreateCloudPoolForbidden handles this case with default header values.

Forbidden
*/
type CreateCloudPoolForbidden struct {
	Payload *models.APIResult
}

func (o *CreateCloudPoolForbidden) Error() string {
	return fmt.Sprintf("[POST /cloudpools][%d] createCloudPoolForbidden  %+v", 403, o.Payload)
}

func (o *CreateCloudPoolForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCloudPoolForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudPoolNotFound creates a CreateCloudPoolNotFound with default headers values
func NewCreateCloudPoolNotFound() *CreateCloudPoolNotFound {
	return &CreateCloudPoolNotFound{}
}

/*CreateCloudPoolNotFound handles this case with default header values.

Not Found
*/
type CreateCloudPoolNotFound struct {
	Payload *models.APIResult
}

func (o *CreateCloudPoolNotFound) Error() string {
	return fmt.Sprintf("[POST /cloudpools][%d] createCloudPoolNotFound  %+v", 404, o.Payload)
}

func (o *CreateCloudPoolNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCloudPoolNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudPoolConflict creates a CreateCloudPoolConflict with default headers values
func NewCreateCloudPoolConflict() *CreateCloudPoolConflict {
	return &CreateCloudPoolConflict{}
}

/*CreateCloudPoolConflict handles this case with default header values.

Conflict
*/
type CreateCloudPoolConflict struct {
	Payload *models.APIResult
}

func (o *CreateCloudPoolConflict) Error() string {
	return fmt.Sprintf("[POST /cloudpools][%d] createCloudPoolConflict  %+v", 409, o.Payload)
}

func (o *CreateCloudPoolConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCloudPoolConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewCreateCloudPoolInternalServerError creates a CreateCloudPoolInternalServerError with default headers values
func NewCreateCloudPoolInternalServerError() *CreateCloudPoolInternalServerError {
	return &CreateCloudPoolInternalServerError{}
}

/*CreateCloudPoolInternalServerError handles this case with default header values.

Internal Server Error
*/
type CreateCloudPoolInternalServerError struct {
	Payload *models.APIResult
}

func (o *CreateCloudPoolInternalServerError) Error() string {
	return fmt.Sprintf("[POST /cloudpools][%d] createCloudPoolInternalServerError  %+v", 500, o.Payload)
}

func (o *CreateCloudPoolInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *CreateCloudPoolInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}