// Code generated by go-swagger; DO NOT EDIT.

package plugin_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ReadPluginReader is a Reader for the ReadPlugin structure.
type ReadPluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ReadPluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewReadPluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewReadPluginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewReadPluginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewReadPluginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewReadPluginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewReadPluginConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewReadPluginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewReadPluginOK creates a ReadPluginOK with default headers values
func NewReadPluginOK() *ReadPluginOK {
	return &ReadPluginOK{}
}

/*ReadPluginOK handles this case with default header values.

OK
*/
type ReadPluginOK struct {
	Payload *models.APIResultPluginMetaData
}

func (o *ReadPluginOK) Error() string {
	return fmt.Sprintf("[GET /plugins/{id}][%d] readPluginOK  %+v", 200, o.Payload)
}

func (o *ReadPluginOK) GetPayload() *models.APIResultPluginMetaData {
	return o.Payload
}

func (o *ReadPluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultPluginMetaData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPluginBadRequest creates a ReadPluginBadRequest with default headers values
func NewReadPluginBadRequest() *ReadPluginBadRequest {
	return &ReadPluginBadRequest{}
}

/*ReadPluginBadRequest handles this case with default header values.

Bad Request
*/
type ReadPluginBadRequest struct {
	Payload *models.APIResult
}

func (o *ReadPluginBadRequest) Error() string {
	return fmt.Sprintf("[GET /plugins/{id}][%d] readPluginBadRequest  %+v", 400, o.Payload)
}

func (o *ReadPluginBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPluginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPluginUnauthorized creates a ReadPluginUnauthorized with default headers values
func NewReadPluginUnauthorized() *ReadPluginUnauthorized {
	return &ReadPluginUnauthorized{}
}

/*ReadPluginUnauthorized handles this case with default header values.

Unauthorized
*/
type ReadPluginUnauthorized struct {
	Payload *models.APIResult
}

func (o *ReadPluginUnauthorized) Error() string {
	return fmt.Sprintf("[GET /plugins/{id}][%d] readPluginUnauthorized  %+v", 401, o.Payload)
}

func (o *ReadPluginUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPluginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPluginForbidden creates a ReadPluginForbidden with default headers values
func NewReadPluginForbidden() *ReadPluginForbidden {
	return &ReadPluginForbidden{}
}

/*ReadPluginForbidden handles this case with default header values.

Forbidden
*/
type ReadPluginForbidden struct {
	Payload *models.APIResult
}

func (o *ReadPluginForbidden) Error() string {
	return fmt.Sprintf("[GET /plugins/{id}][%d] readPluginForbidden  %+v", 403, o.Payload)
}

func (o *ReadPluginForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPluginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPluginNotFound creates a ReadPluginNotFound with default headers values
func NewReadPluginNotFound() *ReadPluginNotFound {
	return &ReadPluginNotFound{}
}

/*ReadPluginNotFound handles this case with default header values.

Not Found
*/
type ReadPluginNotFound struct {
	Payload *models.APIResult
}

func (o *ReadPluginNotFound) Error() string {
	return fmt.Sprintf("[GET /plugins/{id}][%d] readPluginNotFound  %+v", 404, o.Payload)
}

func (o *ReadPluginNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPluginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPluginConflict creates a ReadPluginConflict with default headers values
func NewReadPluginConflict() *ReadPluginConflict {
	return &ReadPluginConflict{}
}

/*ReadPluginConflict handles this case with default header values.

Conflict
*/
type ReadPluginConflict struct {
	Payload *models.APIResult
}

func (o *ReadPluginConflict) Error() string {
	return fmt.Sprintf("[GET /plugins/{id}][%d] readPluginConflict  %+v", 409, o.Payload)
}

func (o *ReadPluginConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPluginConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewReadPluginInternalServerError creates a ReadPluginInternalServerError with default headers values
func NewReadPluginInternalServerError() *ReadPluginInternalServerError {
	return &ReadPluginInternalServerError{}
}

/*ReadPluginInternalServerError handles this case with default header values.

Internal Server Error
*/
type ReadPluginInternalServerError struct {
	Payload *models.APIResult
}

func (o *ReadPluginInternalServerError) Error() string {
	return fmt.Sprintf("[GET /plugins/{id}][%d] readPluginInternalServerError  %+v", 500, o.Payload)
}

func (o *ReadPluginInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ReadPluginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}