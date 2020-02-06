// Code generated by go-swagger; DO NOT EDIT.

package seed_bundle_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// UploadSeedBundleReader is a Reader for the UploadSeedBundle structure.
type UploadSeedBundleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UploadSeedBundleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUploadSeedBundleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUploadSeedBundleBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewUploadSeedBundleUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUploadSeedBundleForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUploadSeedBundleNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewUploadSeedBundleConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUploadSeedBundleInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUploadSeedBundleOK creates a UploadSeedBundleOK with default headers values
func NewUploadSeedBundleOK() *UploadSeedBundleOK {
	return &UploadSeedBundleOK{}
}

/*UploadSeedBundleOK handles this case with default header values.

OK
*/
type UploadSeedBundleOK struct {
	Payload *models.APIResultVoid
}

func (o *UploadSeedBundleOK) Error() string {
	return fmt.Sprintf("[POST /seedBundles][%d] uploadSeedBundleOK  %+v", 200, o.Payload)
}

func (o *UploadSeedBundleOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *UploadSeedBundleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSeedBundleBadRequest creates a UploadSeedBundleBadRequest with default headers values
func NewUploadSeedBundleBadRequest() *UploadSeedBundleBadRequest {
	return &UploadSeedBundleBadRequest{}
}

/*UploadSeedBundleBadRequest handles this case with default header values.

Bad Request
*/
type UploadSeedBundleBadRequest struct {
	Payload *models.APIResult
}

func (o *UploadSeedBundleBadRequest) Error() string {
	return fmt.Sprintf("[POST /seedBundles][%d] uploadSeedBundleBadRequest  %+v", 400, o.Payload)
}

func (o *UploadSeedBundleBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadSeedBundleBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSeedBundleUnauthorized creates a UploadSeedBundleUnauthorized with default headers values
func NewUploadSeedBundleUnauthorized() *UploadSeedBundleUnauthorized {
	return &UploadSeedBundleUnauthorized{}
}

/*UploadSeedBundleUnauthorized handles this case with default header values.

Unauthorized
*/
type UploadSeedBundleUnauthorized struct {
	Payload *models.APIResult
}

func (o *UploadSeedBundleUnauthorized) Error() string {
	return fmt.Sprintf("[POST /seedBundles][%d] uploadSeedBundleUnauthorized  %+v", 401, o.Payload)
}

func (o *UploadSeedBundleUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadSeedBundleUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSeedBundleForbidden creates a UploadSeedBundleForbidden with default headers values
func NewUploadSeedBundleForbidden() *UploadSeedBundleForbidden {
	return &UploadSeedBundleForbidden{}
}

/*UploadSeedBundleForbidden handles this case with default header values.

Forbidden
*/
type UploadSeedBundleForbidden struct {
	Payload *models.APIResult
}

func (o *UploadSeedBundleForbidden) Error() string {
	return fmt.Sprintf("[POST /seedBundles][%d] uploadSeedBundleForbidden  %+v", 403, o.Payload)
}

func (o *UploadSeedBundleForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadSeedBundleForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSeedBundleNotFound creates a UploadSeedBundleNotFound with default headers values
func NewUploadSeedBundleNotFound() *UploadSeedBundleNotFound {
	return &UploadSeedBundleNotFound{}
}

/*UploadSeedBundleNotFound handles this case with default header values.

Not Found
*/
type UploadSeedBundleNotFound struct {
	Payload *models.APIResult
}

func (o *UploadSeedBundleNotFound) Error() string {
	return fmt.Sprintf("[POST /seedBundles][%d] uploadSeedBundleNotFound  %+v", 404, o.Payload)
}

func (o *UploadSeedBundleNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadSeedBundleNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSeedBundleConflict creates a UploadSeedBundleConflict with default headers values
func NewUploadSeedBundleConflict() *UploadSeedBundleConflict {
	return &UploadSeedBundleConflict{}
}

/*UploadSeedBundleConflict handles this case with default header values.

Conflict
*/
type UploadSeedBundleConflict struct {
	Payload *models.APIResult
}

func (o *UploadSeedBundleConflict) Error() string {
	return fmt.Sprintf("[POST /seedBundles][%d] uploadSeedBundleConflict  %+v", 409, o.Payload)
}

func (o *UploadSeedBundleConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadSeedBundleConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUploadSeedBundleInternalServerError creates a UploadSeedBundleInternalServerError with default headers values
func NewUploadSeedBundleInternalServerError() *UploadSeedBundleInternalServerError {
	return &UploadSeedBundleInternalServerError{}
}

/*UploadSeedBundleInternalServerError handles this case with default header values.

Internal Server Error
*/
type UploadSeedBundleInternalServerError struct {
	Payload *models.APIResult
}

func (o *UploadSeedBundleInternalServerError) Error() string {
	return fmt.Sprintf("[POST /seedBundles][%d] uploadSeedBundleInternalServerError  %+v", 500, o.Payload)
}

func (o *UploadSeedBundleInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *UploadSeedBundleInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}