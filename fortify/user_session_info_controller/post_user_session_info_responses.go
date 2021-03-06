// Code generated by go-swagger; DO NOT EDIT.

package user_session_info_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// PostUserSessionInfoReader is a Reader for the PostUserSessionInfo structure.
type PostUserSessionInfoReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PostUserSessionInfoReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPostUserSessionInfoOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPostUserSessionInfoBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPostUserSessionInfoUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPostUserSessionInfoForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPostUserSessionInfoNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewPostUserSessionInfoConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPostUserSessionInfoInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewPostUserSessionInfoOK creates a PostUserSessionInfoOK with default headers values
func NewPostUserSessionInfoOK() *PostUserSessionInfoOK {
	return &PostUserSessionInfoOK{}
}

/*PostUserSessionInfoOK handles this case with default header values.

OK
*/
type PostUserSessionInfoOK struct {
	Payload *models.APIResultUserSessionInformation
}

func (o *PostUserSessionInfoOK) Error() string {
	return fmt.Sprintf("[POST /userSession/info][%d] postUserSessionInfoOK  %+v", 200, o.Payload)
}

func (o *PostUserSessionInfoOK) GetPayload() *models.APIResultUserSessionInformation {
	return o.Payload
}

func (o *PostUserSessionInfoOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultUserSessionInformation)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserSessionInfoBadRequest creates a PostUserSessionInfoBadRequest with default headers values
func NewPostUserSessionInfoBadRequest() *PostUserSessionInfoBadRequest {
	return &PostUserSessionInfoBadRequest{}
}

/*PostUserSessionInfoBadRequest handles this case with default header values.

Bad Request
*/
type PostUserSessionInfoBadRequest struct {
	Payload *models.APIResult
}

func (o *PostUserSessionInfoBadRequest) Error() string {
	return fmt.Sprintf("[POST /userSession/info][%d] postUserSessionInfoBadRequest  %+v", 400, o.Payload)
}

func (o *PostUserSessionInfoBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserSessionInfoBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserSessionInfoUnauthorized creates a PostUserSessionInfoUnauthorized with default headers values
func NewPostUserSessionInfoUnauthorized() *PostUserSessionInfoUnauthorized {
	return &PostUserSessionInfoUnauthorized{}
}

/*PostUserSessionInfoUnauthorized handles this case with default header values.

Unauthorized
*/
type PostUserSessionInfoUnauthorized struct {
	Payload *models.APIResult
}

func (o *PostUserSessionInfoUnauthorized) Error() string {
	return fmt.Sprintf("[POST /userSession/info][%d] postUserSessionInfoUnauthorized  %+v", 401, o.Payload)
}

func (o *PostUserSessionInfoUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserSessionInfoUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserSessionInfoForbidden creates a PostUserSessionInfoForbidden with default headers values
func NewPostUserSessionInfoForbidden() *PostUserSessionInfoForbidden {
	return &PostUserSessionInfoForbidden{}
}

/*PostUserSessionInfoForbidden handles this case with default header values.

Forbidden
*/
type PostUserSessionInfoForbidden struct {
	Payload *models.APIResult
}

func (o *PostUserSessionInfoForbidden) Error() string {
	return fmt.Sprintf("[POST /userSession/info][%d] postUserSessionInfoForbidden  %+v", 403, o.Payload)
}

func (o *PostUserSessionInfoForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserSessionInfoForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserSessionInfoNotFound creates a PostUserSessionInfoNotFound with default headers values
func NewPostUserSessionInfoNotFound() *PostUserSessionInfoNotFound {
	return &PostUserSessionInfoNotFound{}
}

/*PostUserSessionInfoNotFound handles this case with default header values.

Not Found
*/
type PostUserSessionInfoNotFound struct {
	Payload *models.APIResult
}

func (o *PostUserSessionInfoNotFound) Error() string {
	return fmt.Sprintf("[POST /userSession/info][%d] postUserSessionInfoNotFound  %+v", 404, o.Payload)
}

func (o *PostUserSessionInfoNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserSessionInfoNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserSessionInfoConflict creates a PostUserSessionInfoConflict with default headers values
func NewPostUserSessionInfoConflict() *PostUserSessionInfoConflict {
	return &PostUserSessionInfoConflict{}
}

/*PostUserSessionInfoConflict handles this case with default header values.

Conflict
*/
type PostUserSessionInfoConflict struct {
	Payload *models.APIResult
}

func (o *PostUserSessionInfoConflict) Error() string {
	return fmt.Sprintf("[POST /userSession/info][%d] postUserSessionInfoConflict  %+v", 409, o.Payload)
}

func (o *PostUserSessionInfoConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserSessionInfoConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPostUserSessionInfoInternalServerError creates a PostUserSessionInfoInternalServerError with default headers values
func NewPostUserSessionInfoInternalServerError() *PostUserSessionInfoInternalServerError {
	return &PostUserSessionInfoInternalServerError{}
}

/*PostUserSessionInfoInternalServerError handles this case with default header values.

Internal Server Error
*/
type PostUserSessionInfoInternalServerError struct {
	Payload *models.APIResult
}

func (o *PostUserSessionInfoInternalServerError) Error() string {
	return fmt.Sprintf("[POST /userSession/info][%d] postUserSessionInfoInternalServerError  %+v", 500, o.Payload)
}

func (o *PostUserSessionInfoInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *PostUserSessionInfoInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
