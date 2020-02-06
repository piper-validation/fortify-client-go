// Code generated by go-swagger; DO NOT EDIT.

package plugin_image_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// GetPluginImageReader is a Reader for the GetPluginImage structure.
type GetPluginImageReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetPluginImageReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetPluginImageOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetPluginImageBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetPluginImageUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetPluginImageForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetPluginImageNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetPluginImageConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetPluginImageInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetPluginImageOK creates a GetPluginImageOK with default headers values
func NewGetPluginImageOK() *GetPluginImageOK {
	return &GetPluginImageOK{}
}

/*GetPluginImageOK handles this case with default header values.

OK
*/
type GetPluginImageOK struct {
	Payload strfmt.Base64
}

func (o *GetPluginImageOK) Error() string {
	return fmt.Sprintf("[GET /pluginimage][%d] getPluginImageOK  %+v", 200, o.Payload)
}

func (o *GetPluginImageOK) GetPayload() strfmt.Base64 {
	return o.Payload
}

func (o *GetPluginImageOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginImageBadRequest creates a GetPluginImageBadRequest with default headers values
func NewGetPluginImageBadRequest() *GetPluginImageBadRequest {
	return &GetPluginImageBadRequest{}
}

/*GetPluginImageBadRequest handles this case with default header values.

Bad Request
*/
type GetPluginImageBadRequest struct {
	Payload *models.APIResult
}

func (o *GetPluginImageBadRequest) Error() string {
	return fmt.Sprintf("[GET /pluginimage][%d] getPluginImageBadRequest  %+v", 400, o.Payload)
}

func (o *GetPluginImageBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetPluginImageBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginImageUnauthorized creates a GetPluginImageUnauthorized with default headers values
func NewGetPluginImageUnauthorized() *GetPluginImageUnauthorized {
	return &GetPluginImageUnauthorized{}
}

/*GetPluginImageUnauthorized handles this case with default header values.

Unauthorized
*/
type GetPluginImageUnauthorized struct {
	Payload *models.APIResult
}

func (o *GetPluginImageUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pluginimage][%d] getPluginImageUnauthorized  %+v", 401, o.Payload)
}

func (o *GetPluginImageUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetPluginImageUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginImageForbidden creates a GetPluginImageForbidden with default headers values
func NewGetPluginImageForbidden() *GetPluginImageForbidden {
	return &GetPluginImageForbidden{}
}

/*GetPluginImageForbidden handles this case with default header values.

Forbidden
*/
type GetPluginImageForbidden struct {
	Payload *models.APIResult
}

func (o *GetPluginImageForbidden) Error() string {
	return fmt.Sprintf("[GET /pluginimage][%d] getPluginImageForbidden  %+v", 403, o.Payload)
}

func (o *GetPluginImageForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetPluginImageForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginImageNotFound creates a GetPluginImageNotFound with default headers values
func NewGetPluginImageNotFound() *GetPluginImageNotFound {
	return &GetPluginImageNotFound{}
}

/*GetPluginImageNotFound handles this case with default header values.

Not Found
*/
type GetPluginImageNotFound struct {
	Payload *models.APIResult
}

func (o *GetPluginImageNotFound) Error() string {
	return fmt.Sprintf("[GET /pluginimage][%d] getPluginImageNotFound  %+v", 404, o.Payload)
}

func (o *GetPluginImageNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetPluginImageNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginImageConflict creates a GetPluginImageConflict with default headers values
func NewGetPluginImageConflict() *GetPluginImageConflict {
	return &GetPluginImageConflict{}
}

/*GetPluginImageConflict handles this case with default header values.

Conflict
*/
type GetPluginImageConflict struct {
	Payload *models.APIResult
}

func (o *GetPluginImageConflict) Error() string {
	return fmt.Sprintf("[GET /pluginimage][%d] getPluginImageConflict  %+v", 409, o.Payload)
}

func (o *GetPluginImageConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetPluginImageConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetPluginImageInternalServerError creates a GetPluginImageInternalServerError with default headers values
func NewGetPluginImageInternalServerError() *GetPluginImageInternalServerError {
	return &GetPluginImageInternalServerError{}
}

/*GetPluginImageInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetPluginImageInternalServerError struct {
	Payload *models.APIResult
}

func (o *GetPluginImageInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pluginimage][%d] getPluginImageInternalServerError  %+v", 500, o.Payload)
}

func (o *GetPluginImageInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetPluginImageInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}