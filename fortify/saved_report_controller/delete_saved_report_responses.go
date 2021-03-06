// Code generated by go-swagger; DO NOT EDIT.

package saved_report_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// DeleteSavedReportReader is a Reader for the DeleteSavedReport structure.
type DeleteSavedReportReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteSavedReportReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDeleteSavedReportOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteSavedReportBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewDeleteSavedReportUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewDeleteSavedReportForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDeleteSavedReportNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewDeleteSavedReportConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewDeleteSavedReportInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteSavedReportOK creates a DeleteSavedReportOK with default headers values
func NewDeleteSavedReportOK() *DeleteSavedReportOK {
	return &DeleteSavedReportOK{}
}

/*DeleteSavedReportOK handles this case with default header values.

OK
*/
type DeleteSavedReportOK struct {
	Payload *models.APIResultVoid
}

func (o *DeleteSavedReportOK) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteSavedReportOK  %+v", 200, o.Payload)
}

func (o *DeleteSavedReportOK) GetPayload() *models.APIResultVoid {
	return o.Payload
}

func (o *DeleteSavedReportOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultVoid)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedReportBadRequest creates a DeleteSavedReportBadRequest with default headers values
func NewDeleteSavedReportBadRequest() *DeleteSavedReportBadRequest {
	return &DeleteSavedReportBadRequest{}
}

/*DeleteSavedReportBadRequest handles this case with default header values.

Bad Request
*/
type DeleteSavedReportBadRequest struct {
	Payload *models.APIResult
}

func (o *DeleteSavedReportBadRequest) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteSavedReportBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteSavedReportBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteSavedReportBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedReportUnauthorized creates a DeleteSavedReportUnauthorized with default headers values
func NewDeleteSavedReportUnauthorized() *DeleteSavedReportUnauthorized {
	return &DeleteSavedReportUnauthorized{}
}

/*DeleteSavedReportUnauthorized handles this case with default header values.

Unauthorized
*/
type DeleteSavedReportUnauthorized struct {
	Payload *models.APIResult
}

func (o *DeleteSavedReportUnauthorized) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteSavedReportUnauthorized  %+v", 401, o.Payload)
}

func (o *DeleteSavedReportUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteSavedReportUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedReportForbidden creates a DeleteSavedReportForbidden with default headers values
func NewDeleteSavedReportForbidden() *DeleteSavedReportForbidden {
	return &DeleteSavedReportForbidden{}
}

/*DeleteSavedReportForbidden handles this case with default header values.

Forbidden
*/
type DeleteSavedReportForbidden struct {
	Payload *models.APIResult
}

func (o *DeleteSavedReportForbidden) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteSavedReportForbidden  %+v", 403, o.Payload)
}

func (o *DeleteSavedReportForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteSavedReportForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedReportNotFound creates a DeleteSavedReportNotFound with default headers values
func NewDeleteSavedReportNotFound() *DeleteSavedReportNotFound {
	return &DeleteSavedReportNotFound{}
}

/*DeleteSavedReportNotFound handles this case with default header values.

Not Found
*/
type DeleteSavedReportNotFound struct {
	Payload *models.APIResult
}

func (o *DeleteSavedReportNotFound) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteSavedReportNotFound  %+v", 404, o.Payload)
}

func (o *DeleteSavedReportNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteSavedReportNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedReportConflict creates a DeleteSavedReportConflict with default headers values
func NewDeleteSavedReportConflict() *DeleteSavedReportConflict {
	return &DeleteSavedReportConflict{}
}

/*DeleteSavedReportConflict handles this case with default header values.

Conflict
*/
type DeleteSavedReportConflict struct {
	Payload *models.APIResult
}

func (o *DeleteSavedReportConflict) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteSavedReportConflict  %+v", 409, o.Payload)
}

func (o *DeleteSavedReportConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteSavedReportConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDeleteSavedReportInternalServerError creates a DeleteSavedReportInternalServerError with default headers values
func NewDeleteSavedReportInternalServerError() *DeleteSavedReportInternalServerError {
	return &DeleteSavedReportInternalServerError{}
}

/*DeleteSavedReportInternalServerError handles this case with default header values.

Internal Server Error
*/
type DeleteSavedReportInternalServerError struct {
	Payload *models.APIResult
}

func (o *DeleteSavedReportInternalServerError) Error() string {
	return fmt.Sprintf("[DELETE /reports/{id}][%d] deleteSavedReportInternalServerError  %+v", 500, o.Payload)
}

func (o *DeleteSavedReportInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *DeleteSavedReportInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
