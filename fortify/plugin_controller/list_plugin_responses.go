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

// ListPluginReader is a Reader for the ListPlugin structure.
type ListPluginReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPluginReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPluginOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPluginBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListPluginUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPluginForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListPluginNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListPluginConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPluginInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPluginOK creates a ListPluginOK with default headers values
func NewListPluginOK() *ListPluginOK {
	return &ListPluginOK{}
}

/*ListPluginOK handles this case with default header values.

OK
*/
type ListPluginOK struct {
	Payload *models.APIResultListPluginMetaData
}

func (o *ListPluginOK) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] listPluginOK  %+v", 200, o.Payload)
}

func (o *ListPluginOK) GetPayload() *models.APIResultListPluginMetaData {
	return o.Payload
}

func (o *ListPluginOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListPluginMetaData)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPluginBadRequest creates a ListPluginBadRequest with default headers values
func NewListPluginBadRequest() *ListPluginBadRequest {
	return &ListPluginBadRequest{}
}

/*ListPluginBadRequest handles this case with default header values.

Bad Request
*/
type ListPluginBadRequest struct {
	Payload *models.APIResult
}

func (o *ListPluginBadRequest) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] listPluginBadRequest  %+v", 400, o.Payload)
}

func (o *ListPluginBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPluginBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPluginUnauthorized creates a ListPluginUnauthorized with default headers values
func NewListPluginUnauthorized() *ListPluginUnauthorized {
	return &ListPluginUnauthorized{}
}

/*ListPluginUnauthorized handles this case with default header values.

Unauthorized
*/
type ListPluginUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListPluginUnauthorized) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] listPluginUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPluginUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPluginUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPluginForbidden creates a ListPluginForbidden with default headers values
func NewListPluginForbidden() *ListPluginForbidden {
	return &ListPluginForbidden{}
}

/*ListPluginForbidden handles this case with default header values.

Forbidden
*/
type ListPluginForbidden struct {
	Payload *models.APIResult
}

func (o *ListPluginForbidden) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] listPluginForbidden  %+v", 403, o.Payload)
}

func (o *ListPluginForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPluginForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPluginNotFound creates a ListPluginNotFound with default headers values
func NewListPluginNotFound() *ListPluginNotFound {
	return &ListPluginNotFound{}
}

/*ListPluginNotFound handles this case with default header values.

Not Found
*/
type ListPluginNotFound struct {
	Payload *models.APIResult
}

func (o *ListPluginNotFound) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] listPluginNotFound  %+v", 404, o.Payload)
}

func (o *ListPluginNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPluginNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPluginConflict creates a ListPluginConflict with default headers values
func NewListPluginConflict() *ListPluginConflict {
	return &ListPluginConflict{}
}

/*ListPluginConflict handles this case with default header values.

Conflict
*/
type ListPluginConflict struct {
	Payload *models.APIResult
}

func (o *ListPluginConflict) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] listPluginConflict  %+v", 409, o.Payload)
}

func (o *ListPluginConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPluginConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPluginInternalServerError creates a ListPluginInternalServerError with default headers values
func NewListPluginInternalServerError() *ListPluginInternalServerError {
	return &ListPluginInternalServerError{}
}

/*ListPluginInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListPluginInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListPluginInternalServerError) Error() string {
	return fmt.Sprintf("[GET /plugins][%d] listPluginInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPluginInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPluginInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}