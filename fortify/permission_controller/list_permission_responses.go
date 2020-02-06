// Code generated by go-swagger; DO NOT EDIT.

package permission_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// ListPermissionReader is a Reader for the ListPermission structure.
type ListPermissionReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ListPermissionReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewListPermissionOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewListPermissionBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewListPermissionUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewListPermissionForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewListPermissionNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewListPermissionConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewListPermissionInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewListPermissionOK creates a ListPermissionOK with default headers values
func NewListPermissionOK() *ListPermissionOK {
	return &ListPermissionOK{}
}

/*ListPermissionOK handles this case with default header values.

OK
*/
type ListPermissionOK struct {
	Payload *models.APIResultListPermission
}

func (o *ListPermissionOK) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] listPermissionOK  %+v", 200, o.Payload)
}

func (o *ListPermissionOK) GetPayload() *models.APIResultListPermission {
	return o.Payload
}

func (o *ListPermissionOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultListPermission)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPermissionBadRequest creates a ListPermissionBadRequest with default headers values
func NewListPermissionBadRequest() *ListPermissionBadRequest {
	return &ListPermissionBadRequest{}
}

/*ListPermissionBadRequest handles this case with default header values.

Bad Request
*/
type ListPermissionBadRequest struct {
	Payload *models.APIResult
}

func (o *ListPermissionBadRequest) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] listPermissionBadRequest  %+v", 400, o.Payload)
}

func (o *ListPermissionBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPermissionBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPermissionUnauthorized creates a ListPermissionUnauthorized with default headers values
func NewListPermissionUnauthorized() *ListPermissionUnauthorized {
	return &ListPermissionUnauthorized{}
}

/*ListPermissionUnauthorized handles this case with default header values.

Unauthorized
*/
type ListPermissionUnauthorized struct {
	Payload *models.APIResult
}

func (o *ListPermissionUnauthorized) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] listPermissionUnauthorized  %+v", 401, o.Payload)
}

func (o *ListPermissionUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPermissionUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPermissionForbidden creates a ListPermissionForbidden with default headers values
func NewListPermissionForbidden() *ListPermissionForbidden {
	return &ListPermissionForbidden{}
}

/*ListPermissionForbidden handles this case with default header values.

Forbidden
*/
type ListPermissionForbidden struct {
	Payload *models.APIResult
}

func (o *ListPermissionForbidden) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] listPermissionForbidden  %+v", 403, o.Payload)
}

func (o *ListPermissionForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPermissionForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPermissionNotFound creates a ListPermissionNotFound with default headers values
func NewListPermissionNotFound() *ListPermissionNotFound {
	return &ListPermissionNotFound{}
}

/*ListPermissionNotFound handles this case with default header values.

Not Found
*/
type ListPermissionNotFound struct {
	Payload *models.APIResult
}

func (o *ListPermissionNotFound) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] listPermissionNotFound  %+v", 404, o.Payload)
}

func (o *ListPermissionNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPermissionNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPermissionConflict creates a ListPermissionConflict with default headers values
func NewListPermissionConflict() *ListPermissionConflict {
	return &ListPermissionConflict{}
}

/*ListPermissionConflict handles this case with default header values.

Conflict
*/
type ListPermissionConflict struct {
	Payload *models.APIResult
}

func (o *ListPermissionConflict) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] listPermissionConflict  %+v", 409, o.Payload)
}

func (o *ListPermissionConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPermissionConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewListPermissionInternalServerError creates a ListPermissionInternalServerError with default headers values
func NewListPermissionInternalServerError() *ListPermissionInternalServerError {
	return &ListPermissionInternalServerError{}
}

/*ListPermissionInternalServerError handles this case with default header values.

Internal Server Error
*/
type ListPermissionInternalServerError struct {
	Payload *models.APIResult
}

func (o *ListPermissionInternalServerError) Error() string {
	return fmt.Sprintf("[GET /permissions][%d] listPermissionInternalServerError  %+v", 500, o.Payload)
}

func (o *ListPermissionInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *ListPermissionInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}