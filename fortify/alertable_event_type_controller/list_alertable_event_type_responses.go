// Code generated by go-swagger; DO NOT EDIT.

package alertable_event_type_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListAlertableEventTypeReader is a Reader for the ListAlertableEventType structure.
type ListAlertableEventTypeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListAlertableEventTypeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListAlertableEventTypeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListAlertableEventTypeBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListAlertableEventTypeUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListAlertableEventTypeForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListAlertableEventTypeNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListAlertableEventTypeConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListAlertableEventTypeInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListAlertableEventTypeOK creates a ListAlertableEventTypeOK with default headers values
func NewListAlertableEventTypeOK() *ListAlertableEventTypeOK {
	return &ListAlertableEventTypeOK{}
}

/*ListAlertableEventTypeOK handles this case with default header values.

OK
*/
type ListAlertableEventTypeOK struct {
	Payload *models.APIResultListAlertableEventType
}

func (o *ListAlertableEventTypeOK) Error() string {
	return fmt.Sprintf("[GET /alertableEventTypes][%d] listAlertableEventTypeOK  %+v", 200, o.Payload)
}

func (o *ListAlertableEventTypeOK) GetPayload() *models.APIResultListAlertableEventType {
	return o.Payload
}

func (o *ListAlertableEventTypeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListAlertableEventType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlertableEventTypeBadRequest creates a ListAlertableEventTypeBadRequest with default headers values
func NewListAlertableEventTypeBadRequest() *ListAlertableEventTypeBadRequest {
	return &ListAlertableEventTypeBadRequest{}
}

/*ListAlertableEventTypeBadRequest handles this case with default header values.

Bad Request
*/
type ListAlertableEventTypeBadRequest struct {
	Payload *models.APIResult
}

func (o *ListAlertableEventTypeBadRequest) Error() string {
	return fmt.Sprintf("[GET /alertableEventTypes][%d] listAlertableEventTypeBadRequest  %+v", 400, o.Payload)
}

func (o *ListAlertableEventTypeBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAlertableEventTypeBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlertableEventTypeUnauthorized creates a ListAlertableEventTypeUnauthorized with default headers values
func NewListAlertableEventTypeUnauthorized() *ListAlertableEventTypeUnauthorized {
	return &ListAlertableEventTypeUnauthorized{}
}

/*ListAlertableEventTypeUnauthorized handles this case with default header values.

Unauthorized
*/
type ListAlertableEventTypeUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListAlertableEventTypeUnauthorized) Error() string {
	return fmt.Sprintf("[GET /alertableEventTypes][%d] listAlertableEventTypeUnauthorized  %+v", 401, o.Payload)
}

func (o *ListAlertableEventTypeUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAlertableEventTypeUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlertableEventTypeForbidden creates a ListAlertableEventTypeForbidden with default headers values
func NewListAlertableEventTypeForbidden() *ListAlertableEventTypeForbidden {
	return &ListAlertableEventTypeForbidden{}
}

/*ListAlertableEventTypeForbidden handles this case with default header values.

Forbidden
*/
type ListAlertableEventTypeForbidden struct {
	Payload *models.APIResult
}

func (o *ListAlertableEventTypeForbidden) Error() string {
	return fmt.Sprintf("[GET /alertableEventTypes][%d] listAlertableEventTypeForbidden  %+v", 403, o.Payload)
}

func (o *ListAlertableEventTypeForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAlertableEventTypeForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlertableEventTypeNotFound creates a ListAlertableEventTypeNotFound with default headers values
func NewListAlertableEventTypeNotFound() *ListAlertableEventTypeNotFound {
	return &ListAlertableEventTypeNotFound{}
}

/*ListAlertableEventTypeNotFound handles this case with default header values.

Not Found
*/
type ListAlertableEventTypeNotFound struct {
	Payload *models.APIResult
}

func (o *ListAlertableEventTypeNotFound) Error() string {
	return fmt.Sprintf("[GET /alertableEventTypes][%d] listAlertableEventTypeNotFound  %+v", 404, o.Payload)
}

func (o *ListAlertableEventTypeNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAlertableEventTypeNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlertableEventTypeConflict creates a ListAlertableEventTypeConflict with default headers values
func NewListAlertableEventTypeConflict() *ListAlertableEventTypeConflict {
	return &ListAlertableEventTypeConflict{}
}

/*ListAlertableEventTypeConflict handles this case with default header values.

Conflict
*/
type ListAlertableEventTypeConflict struct {
	Payload *models.APIResult
}

func (o *ListAlertableEventTypeConflict) Error() string {
	return fmt.Sprintf("[GET /alertableEventTypes][%d] listAlertableEventTypeConflict  %+v", 409, o.Payload)
}

func (o *ListAlertableEventTypeConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAlertableEventTypeConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListAlertableEventTypeInternalServerError creates a ListAlertableEventTypeInternalServerError with default headers values
func NewListAlertableEventTypeInternalServerError() *ListAlertableEventTypeInternalServerError {
	return &ListAlertableEventTypeInternalServerError{}
}

/*ListAlertableEventTypeInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListAlertableEventTypeInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListAlertableEventTypeInternalServerError) Error() string {
	return fmt.Sprintf("[GET /alertableEventTypes][%d] listAlertableEventTypeInternalServerError  %+v", 500, o.Payload)
}

func (o *ListAlertableEventTypeInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListAlertableEventTypeInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
