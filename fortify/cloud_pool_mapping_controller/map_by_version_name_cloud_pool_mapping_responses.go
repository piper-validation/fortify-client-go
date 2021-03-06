// Code generated by go-swagger; DO NOT EDIT.

package cloud_pool_mapping_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// MapByVersionNameCloudPoolMappingReader is a Reader for the MapByVersionNameCloudPoolMapping structure.
type MapByVersionNameCloudPoolMappingReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *MapByVersionNameCloudPoolMappingReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewMapByVersionNameCloudPoolMappingOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewMapByVersionNameCloudPoolMappingBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewMapByVersionNameCloudPoolMappingUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewMapByVersionNameCloudPoolMappingForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewMapByVersionNameCloudPoolMappingNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewMapByVersionNameCloudPoolMappingConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewMapByVersionNameCloudPoolMappingInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewMapByVersionNameCloudPoolMappingOK creates a MapByVersionNameCloudPoolMappingOK with default headers values
func NewMapByVersionNameCloudPoolMappingOK() *MapByVersionNameCloudPoolMappingOK {
	return &MapByVersionNameCloudPoolMappingOK{}
}

/*MapByVersionNameCloudPoolMappingOK handles this case with default header values.

OK
*/
type MapByVersionNameCloudPoolMappingOK struct {
	Payload *models.APIResultCloudPoolMapping
}

func (o *MapByVersionNameCloudPoolMappingOK) Error() string {
	return fmt.Sprintf("[GET /cloudmappings/mapByVersionName][%d] mapByVersionNameCloudPoolMappingOK  %+v", 200, o.Payload)
}

func (o *MapByVersionNameCloudPoolMappingOK) GetPayload() *models.APIResultCloudPoolMapping {
	return o.Payload
}

func (o *MapByVersionNameCloudPoolMappingOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultCloudPoolMapping)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMapByVersionNameCloudPoolMappingBadRequest creates a MapByVersionNameCloudPoolMappingBadRequest with default headers values
func NewMapByVersionNameCloudPoolMappingBadRequest() *MapByVersionNameCloudPoolMappingBadRequest {
	return &MapByVersionNameCloudPoolMappingBadRequest{}
}

/*MapByVersionNameCloudPoolMappingBadRequest handles this case with default header values.

Bad Request
*/
type MapByVersionNameCloudPoolMappingBadRequest struct {
	Payload *models.APIResult
}

func (o *MapByVersionNameCloudPoolMappingBadRequest) Error() string {
	return fmt.Sprintf("[GET /cloudmappings/mapByVersionName][%d] mapByVersionNameCloudPoolMappingBadRequest  %+v", 400, o.Payload)
}

func (o *MapByVersionNameCloudPoolMappingBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MapByVersionNameCloudPoolMappingBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMapByVersionNameCloudPoolMappingUnauthorized creates a MapByVersionNameCloudPoolMappingUnauthorized with default headers values
func NewMapByVersionNameCloudPoolMappingUnauthorized() *MapByVersionNameCloudPoolMappingUnauthorized {
	return &MapByVersionNameCloudPoolMappingUnauthorized{}
}

/*MapByVersionNameCloudPoolMappingUnauthorized handles this case with default header values.

Unauthorized
*/
type MapByVersionNameCloudPoolMappingUnauthorized struct {
	Payload *models.APIResult
}

func (o *MapByVersionNameCloudPoolMappingUnauthorized) Error() string {
	return fmt.Sprintf("[GET /cloudmappings/mapByVersionName][%d] mapByVersionNameCloudPoolMappingUnauthorized  %+v", 401, o.Payload)
}

func (o *MapByVersionNameCloudPoolMappingUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MapByVersionNameCloudPoolMappingUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMapByVersionNameCloudPoolMappingForbidden creates a MapByVersionNameCloudPoolMappingForbidden with default headers values
func NewMapByVersionNameCloudPoolMappingForbidden() *MapByVersionNameCloudPoolMappingForbidden {
	return &MapByVersionNameCloudPoolMappingForbidden{}
}

/*MapByVersionNameCloudPoolMappingForbidden handles this case with default header values.

Forbidden
*/
type MapByVersionNameCloudPoolMappingForbidden struct {
	Payload *models.APIResult
}

func (o *MapByVersionNameCloudPoolMappingForbidden) Error() string {
	return fmt.Sprintf("[GET /cloudmappings/mapByVersionName][%d] mapByVersionNameCloudPoolMappingForbidden  %+v", 403, o.Payload)
}

func (o *MapByVersionNameCloudPoolMappingForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MapByVersionNameCloudPoolMappingForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMapByVersionNameCloudPoolMappingNotFound creates a MapByVersionNameCloudPoolMappingNotFound with default headers values
func NewMapByVersionNameCloudPoolMappingNotFound() *MapByVersionNameCloudPoolMappingNotFound {
	return &MapByVersionNameCloudPoolMappingNotFound{}
}

/*MapByVersionNameCloudPoolMappingNotFound handles this case with default header values.

Not Found
*/
type MapByVersionNameCloudPoolMappingNotFound struct {
	Payload *models.APIResult
}

func (o *MapByVersionNameCloudPoolMappingNotFound) Error() string {
	return fmt.Sprintf("[GET /cloudmappings/mapByVersionName][%d] mapByVersionNameCloudPoolMappingNotFound  %+v", 404, o.Payload)
}

func (o *MapByVersionNameCloudPoolMappingNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MapByVersionNameCloudPoolMappingNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMapByVersionNameCloudPoolMappingConflict creates a MapByVersionNameCloudPoolMappingConflict with default headers values
func NewMapByVersionNameCloudPoolMappingConflict() *MapByVersionNameCloudPoolMappingConflict {
	return &MapByVersionNameCloudPoolMappingConflict{}
}

/*MapByVersionNameCloudPoolMappingConflict handles this case with default header values.

Conflict
*/
type MapByVersionNameCloudPoolMappingConflict struct {
	Payload *models.APIResult
}

func (o *MapByVersionNameCloudPoolMappingConflict) Error() string {
	return fmt.Sprintf("[GET /cloudmappings/mapByVersionName][%d] mapByVersionNameCloudPoolMappingConflict  %+v", 409, o.Payload)
}

func (o *MapByVersionNameCloudPoolMappingConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MapByVersionNameCloudPoolMappingConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewMapByVersionNameCloudPoolMappingInternalServerError creates a MapByVersionNameCloudPoolMappingInternalServerError with default headers values
func NewMapByVersionNameCloudPoolMappingInternalServerError() *MapByVersionNameCloudPoolMappingInternalServerError {
	return &MapByVersionNameCloudPoolMappingInternalServerError{}
}

/*MapByVersionNameCloudPoolMappingInternalServerError handles this case with default header values.

Internal Server Error
*/
type MapByVersionNameCloudPoolMappingInternalServerError struct {
	Payload *models.APIResult
}

func (o *MapByVersionNameCloudPoolMappingInternalServerError) Error() string {
	return fmt.Sprintf("[GET /cloudmappings/mapByVersionName][%d] mapByVersionNameCloudPoolMappingInternalServerError  %+v", 500, o.Payload)
}

func (o *MapByVersionNameCloudPoolMappingInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *MapByVersionNameCloudPoolMappingInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
