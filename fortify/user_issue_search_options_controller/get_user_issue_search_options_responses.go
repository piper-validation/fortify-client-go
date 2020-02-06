// Code generated by go-swagger; DO NOT EDIT.

package user_issue_search_options_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// GetUserIssueSearchOptionsReader is a Reader for the GetUserIssueSearchOptions structure.
type GetUserIssueSearchOptionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetUserIssueSearchOptionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetUserIssueSearchOptionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetUserIssueSearchOptionsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetUserIssueSearchOptionsUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetUserIssueSearchOptionsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetUserIssueSearchOptionsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetUserIssueSearchOptionsConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetUserIssueSearchOptionsInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetUserIssueSearchOptionsOK creates a GetUserIssueSearchOptionsOK with default headers values
func NewGetUserIssueSearchOptionsOK() *GetUserIssueSearchOptionsOK {
	return &GetUserIssueSearchOptionsOK{}
}

/*GetUserIssueSearchOptionsOK handles this case with default header values.

OK
*/
type GetUserIssueSearchOptionsOK struct {
	Payload *models.APIResultUserIssueSearchOptions
}

func (o *GetUserIssueSearchOptionsOK) Error() string {
	return fmt.Sprintf("[GET /userIssueSearchOptions][%d] getUserIssueSearchOptionsOK  %+v", 200, o.Payload)
}

func (o *GetUserIssueSearchOptionsOK) GetPayload() *models.APIResultUserIssueSearchOptions {
	return o.Payload
}

func (o *GetUserIssueSearchOptionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultUserIssueSearchOptions)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIssueSearchOptionsBadRequest creates a GetUserIssueSearchOptionsBadRequest with default headers values
func NewGetUserIssueSearchOptionsBadRequest() *GetUserIssueSearchOptionsBadRequest {
	return &GetUserIssueSearchOptionsBadRequest{}
}

/*GetUserIssueSearchOptionsBadRequest handles this case with default header values.

Bad Request
*/
type GetUserIssueSearchOptionsBadRequest struct {
	Payload *models.APIResult
}

func (o *GetUserIssueSearchOptionsBadRequest) Error() string {
	return fmt.Sprintf("[GET /userIssueSearchOptions][%d] getUserIssueSearchOptionsBadRequest  %+v", 400, o.Payload)
}

func (o *GetUserIssueSearchOptionsBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetUserIssueSearchOptionsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIssueSearchOptionsUnauthorized creates a GetUserIssueSearchOptionsUnauthorized with default headers values
func NewGetUserIssueSearchOptionsUnauthorized() *GetUserIssueSearchOptionsUnauthorized {
	return &GetUserIssueSearchOptionsUnauthorized{}
}

/*GetUserIssueSearchOptionsUnauthorized handles this case with default header values.

Unauthorized
*/
type GetUserIssueSearchOptionsUnauthorized struct {
	Payload *models.APIResult
}

func (o *GetUserIssueSearchOptionsUnauthorized) Error() string {
	return fmt.Sprintf("[GET /userIssueSearchOptions][%d] getUserIssueSearchOptionsUnauthorized  %+v", 401, o.Payload)
}

func (o *GetUserIssueSearchOptionsUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetUserIssueSearchOptionsUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIssueSearchOptionsForbidden creates a GetUserIssueSearchOptionsForbidden with default headers values
func NewGetUserIssueSearchOptionsForbidden() *GetUserIssueSearchOptionsForbidden {
	return &GetUserIssueSearchOptionsForbidden{}
}

/*GetUserIssueSearchOptionsForbidden handles this case with default header values.

Forbidden
*/
type GetUserIssueSearchOptionsForbidden struct {
	Payload *models.APIResult
}

func (o *GetUserIssueSearchOptionsForbidden) Error() string {
	return fmt.Sprintf("[GET /userIssueSearchOptions][%d] getUserIssueSearchOptionsForbidden  %+v", 403, o.Payload)
}

func (o *GetUserIssueSearchOptionsForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetUserIssueSearchOptionsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIssueSearchOptionsNotFound creates a GetUserIssueSearchOptionsNotFound with default headers values
func NewGetUserIssueSearchOptionsNotFound() *GetUserIssueSearchOptionsNotFound {
	return &GetUserIssueSearchOptionsNotFound{}
}

/*GetUserIssueSearchOptionsNotFound handles this case with default header values.

Not Found
*/
type GetUserIssueSearchOptionsNotFound struct {
	Payload *models.APIResult
}

func (o *GetUserIssueSearchOptionsNotFound) Error() string {
	return fmt.Sprintf("[GET /userIssueSearchOptions][%d] getUserIssueSearchOptionsNotFound  %+v", 404, o.Payload)
}

func (o *GetUserIssueSearchOptionsNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetUserIssueSearchOptionsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIssueSearchOptionsConflict creates a GetUserIssueSearchOptionsConflict with default headers values
func NewGetUserIssueSearchOptionsConflict() *GetUserIssueSearchOptionsConflict {
	return &GetUserIssueSearchOptionsConflict{}
}

/*GetUserIssueSearchOptionsConflict handles this case with default header values.

Conflict
*/
type GetUserIssueSearchOptionsConflict struct {
	Payload *models.APIResult
}

func (o *GetUserIssueSearchOptionsConflict) Error() string {
	return fmt.Sprintf("[GET /userIssueSearchOptions][%d] getUserIssueSearchOptionsConflict  %+v", 409, o.Payload)
}

func (o *GetUserIssueSearchOptionsConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetUserIssueSearchOptionsConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetUserIssueSearchOptionsInternalServerError creates a GetUserIssueSearchOptionsInternalServerError with default headers values
func NewGetUserIssueSearchOptionsInternalServerError() *GetUserIssueSearchOptionsInternalServerError {
	return &GetUserIssueSearchOptionsInternalServerError{}
}

/*GetUserIssueSearchOptionsInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetUserIssueSearchOptionsInternalServerError struct {
	Payload *models.APIResult
}

func (o *GetUserIssueSearchOptionsInternalServerError) Error() string {
	return fmt.Sprintf("[GET /userIssueSearchOptions][%d] getUserIssueSearchOptionsInternalServerError  %+v", 500, o.Payload)
}

func (o *GetUserIssueSearchOptionsInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetUserIssueSearchOptionsInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}