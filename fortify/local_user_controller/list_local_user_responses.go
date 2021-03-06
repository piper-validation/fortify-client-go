// Code generated by go-swagger; DO NOT EDIT.

package local_user_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListLocalUserReader is a Reader for the ListLocalUser structure.
type ListLocalUserReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListLocalUserReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListLocalUserOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListLocalUserBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListLocalUserUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListLocalUserForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListLocalUserNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListLocalUserConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListLocalUserInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListLocalUserOK creates a ListLocalUserOK with default headers values
func NewListLocalUserOK() *ListLocalUserOK {
	return &ListLocalUserOK{}
}

/*ListLocalUserOK handles this case with default header values.

OK
*/
type ListLocalUserOK struct {
	Payload *models.APIResultListLocalUser
}

func (o *ListLocalUserOK) Error() string {
	return fmt.Sprintf("[GET /localUsers][%d] listLocalUserOK  %+v", 200, o.Payload)
}

func (o *ListLocalUserOK) GetPayload() *models.APIResultListLocalUser {
	return o.Payload
}

func (o *ListLocalUserOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListLocalUser)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalUserBadRequest creates a ListLocalUserBadRequest with default headers values
func NewListLocalUserBadRequest() *ListLocalUserBadRequest {
	return &ListLocalUserBadRequest{}
}

/*ListLocalUserBadRequest handles this case with default header values.

Bad Request
*/
type ListLocalUserBadRequest struct {
	Payload *models.APIResult
}

func (o *ListLocalUserBadRequest) Error() string {
	return fmt.Sprintf("[GET /localUsers][%d] listLocalUserBadRequest  %+v", 400, o.Payload)
}

func (o *ListLocalUserBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListLocalUserBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalUserUnauthorized creates a ListLocalUserUnauthorized with default headers values
func NewListLocalUserUnauthorized() *ListLocalUserUnauthorized {
	return &ListLocalUserUnauthorized{}
}

/*ListLocalUserUnauthorized handles this case with default header values.

Unauthorized
*/
type ListLocalUserUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListLocalUserUnauthorized) Error() string {
	return fmt.Sprintf("[GET /localUsers][%d] listLocalUserUnauthorized  %+v", 401, o.Payload)
}

func (o *ListLocalUserUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListLocalUserUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalUserForbidden creates a ListLocalUserForbidden with default headers values
func NewListLocalUserForbidden() *ListLocalUserForbidden {
	return &ListLocalUserForbidden{}
}

/*ListLocalUserForbidden handles this case with default header values.

Forbidden
*/
type ListLocalUserForbidden struct {
	Payload *models.APIResult
}

func (o *ListLocalUserForbidden) Error() string {
	return fmt.Sprintf("[GET /localUsers][%d] listLocalUserForbidden  %+v", 403, o.Payload)
}

func (o *ListLocalUserForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListLocalUserForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalUserNotFound creates a ListLocalUserNotFound with default headers values
func NewListLocalUserNotFound() *ListLocalUserNotFound {
	return &ListLocalUserNotFound{}
}

/*ListLocalUserNotFound handles this case with default header values.

Not Found
*/
type ListLocalUserNotFound struct {
	Payload *models.APIResult
}

func (o *ListLocalUserNotFound) Error() string {
	return fmt.Sprintf("[GET /localUsers][%d] listLocalUserNotFound  %+v", 404, o.Payload)
}

func (o *ListLocalUserNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListLocalUserNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalUserConflict creates a ListLocalUserConflict with default headers values
func NewListLocalUserConflict() *ListLocalUserConflict {
	return &ListLocalUserConflict{}
}

/*ListLocalUserConflict handles this case with default header values.

Conflict
*/
type ListLocalUserConflict struct {
	Payload *models.APIResult
}

func (o *ListLocalUserConflict) Error() string {
	return fmt.Sprintf("[GET /localUsers][%d] listLocalUserConflict  %+v", 409, o.Payload)
}

func (o *ListLocalUserConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListLocalUserConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListLocalUserInternalServerError creates a ListLocalUserInternalServerError with default headers values
func NewListLocalUserInternalServerError() *ListLocalUserInternalServerError {
	return &ListLocalUserInternalServerError{}
}

/*ListLocalUserInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListLocalUserInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListLocalUserInternalServerError) Error() string {
	return fmt.Sprintf("[GET /localUsers][%d] listLocalUserInternalServerError  %+v", 500, o.Payload)
}

func (o *ListLocalUserInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListLocalUserInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
