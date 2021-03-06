// Code generated by go-swagger; DO NOT EDIT.

package feature_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadFeatureReader is a Reader for the ReadFeature structure.
type ReadFeatureReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadFeatureReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadFeatureOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadFeatureBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadFeatureUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadFeatureForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadFeatureNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadFeatureConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadFeatureInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadFeatureOK creates a ReadFeatureOK with default headers values
func NewReadFeatureOK() *ReadFeatureOK {
	return &ReadFeatureOK{}
}

/*ReadFeatureOK handles this case with default header values.

OK
*/
type ReadFeatureOK struct {
	Payload *models.APIResultFeature
}

func (o *ReadFeatureOK) Error() string {
	return fmt.Sprintf("[GET /features/{id}][%d] readFeatureOK  %+v", 200, o.Payload)
}

func (o *ReadFeatureOK) GetPayload() *models.APIResultFeature {
	return o.Payload
}

func (o *ReadFeatureOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultFeature)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadFeatureBadRequest creates a ReadFeatureBadRequest with default headers values
func NewReadFeatureBadRequest() *ReadFeatureBadRequest {
	return &ReadFeatureBadRequest{}
}

/*ReadFeatureBadRequest handles this case with default header values.

Bad Request
*/
type ReadFeatureBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadFeatureBadRequest) Error() string {
	return fmt.Sprintf("[GET /features/{id}][%d] readFeatureBadRequest  %+v", 400, o.Payload)
}

func (o *ReadFeatureBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadFeatureBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadFeatureUnauthorized creates a ReadFeatureUnauthorized with default headers values
func NewReadFeatureUnauthorized() *ReadFeatureUnauthorized {
	return &ReadFeatureUnauthorized{}
}

/*ReadFeatureUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadFeatureUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadFeatureUnauthorized) Error() string {
	return fmt.Sprintf("[GET /features/{id}][%d] readFeatureUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadFeatureUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadFeatureUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadFeatureForbidden creates a ReadFeatureForbidden with default headers values
func NewReadFeatureForbidden() *ReadFeatureForbidden {
	return &ReadFeatureForbidden{}
}

/*ReadFeatureForbidden handles this case with default header values.

Forbidden
*/
type ReadFeatureForbidden struct {
	Payload *models.APIResult
}

func (o *ReadFeatureForbidden) Error() string {
	return fmt.Sprintf("[GET /features/{id}][%d] readFeatureForbidden  %+v", 403, o.Payload)
}

func (o *ReadFeatureForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadFeatureForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadFeatureNotFound creates a ReadFeatureNotFound with default headers values
func NewReadFeatureNotFound() *ReadFeatureNotFound {
	return &ReadFeatureNotFound{}
}

/*ReadFeatureNotFound handles this case with default header values.

Not Found
*/
type ReadFeatureNotFound struct {
	Payload *models.APIResult
}

func (o *ReadFeatureNotFound) Error() string {
	return fmt.Sprintf("[GET /features/{id}][%d] readFeatureNotFound  %+v", 404, o.Payload)
}

func (o *ReadFeatureNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadFeatureNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadFeatureConflict creates a ReadFeatureConflict with default headers values
func NewReadFeatureConflict() *ReadFeatureConflict {
	return &ReadFeatureConflict{}
}

/*ReadFeatureConflict handles this case with default header values.

Conflict
*/
type ReadFeatureConflict struct {
	Payload *models.APIResult
}

func (o *ReadFeatureConflict) Error() string {
	return fmt.Sprintf("[GET /features/{id}][%d] readFeatureConflict  %+v", 409, o.Payload)
}

func (o *ReadFeatureConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadFeatureConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadFeatureInternalServerError creates a ReadFeatureInternalServerError with default headers values
func NewReadFeatureInternalServerError() *ReadFeatureInternalServerError {
	return &ReadFeatureInternalServerError{}
}

/*ReadFeatureInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadFeatureInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadFeatureInternalServerError) Error() string {
	return fmt.Sprintf("[GET /features/{id}][%d] readFeatureInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadFeatureInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadFeatureInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
