// Code generated by go-swagger; DO NOT EDIT.

package configuration_controller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/piper-validation/fortify-client-go/models"
)

// GetConfigurationReader is a Reader for the GetConfiguration structure.
type GetConfigurationReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetConfigurationReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetConfigurationOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetConfigurationBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewGetConfigurationUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetConfigurationForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetConfigurationNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 409:
		result := NewGetConfigurationConflict()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetConfigurationInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetConfigurationOK creates a GetConfigurationOK with default headers values
func NewGetConfigurationOK() *GetConfigurationOK {
	return &GetConfigurationOK{}
}

/*GetConfigurationOK handles this case with default header values.

OK
*/
type GetConfigurationOK struct {
	Payload *models.APIResultConfigProperties
}

func (o *GetConfigurationOK) Error() string {
	return fmt.Sprintf("[GET /configuration][%d] getConfigurationOK  %+v", 200, o.Payload)
}

func (o *GetConfigurationOK) GetPayload() *models.APIResultConfigProperties {
	return o.Payload
}

func (o *GetConfigurationOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResultConfigProperties)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationBadRequest creates a GetConfigurationBadRequest with default headers values
func NewGetConfigurationBadRequest() *GetConfigurationBadRequest {
	return &GetConfigurationBadRequest{}
}

/*GetConfigurationBadRequest handles this case with default header values.

Bad Request
*/
type GetConfigurationBadRequest struct {
	Payload *models.APIResult
}

func (o *GetConfigurationBadRequest) Error() string {
	return fmt.Sprintf("[GET /configuration][%d] getConfigurationBadRequest  %+v", 400, o.Payload)
}

func (o *GetConfigurationBadRequest) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetConfigurationBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationUnauthorized creates a GetConfigurationUnauthorized with default headers values
func NewGetConfigurationUnauthorized() *GetConfigurationUnauthorized {
	return &GetConfigurationUnauthorized{}
}

/*GetConfigurationUnauthorized handles this case with default header values.

Unauthorized
*/
type GetConfigurationUnauthorized struct {
	Payload *models.APIResult
}

func (o *GetConfigurationUnauthorized) Error() string {
	return fmt.Sprintf("[GET /configuration][%d] getConfigurationUnauthorized  %+v", 401, o.Payload)
}

func (o *GetConfigurationUnauthorized) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetConfigurationUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationForbidden creates a GetConfigurationForbidden with default headers values
func NewGetConfigurationForbidden() *GetConfigurationForbidden {
	return &GetConfigurationForbidden{}
}

/*GetConfigurationForbidden handles this case with default header values.

Forbidden
*/
type GetConfigurationForbidden struct {
	Payload *models.APIResult
}

func (o *GetConfigurationForbidden) Error() string {
	return fmt.Sprintf("[GET /configuration][%d] getConfigurationForbidden  %+v", 403, o.Payload)
}

func (o *GetConfigurationForbidden) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetConfigurationForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationNotFound creates a GetConfigurationNotFound with default headers values
func NewGetConfigurationNotFound() *GetConfigurationNotFound {
	return &GetConfigurationNotFound{}
}

/*GetConfigurationNotFound handles this case with default header values.

Not Found
*/
type GetConfigurationNotFound struct {
	Payload *models.APIResult
}

func (o *GetConfigurationNotFound) Error() string {
	return fmt.Sprintf("[GET /configuration][%d] getConfigurationNotFound  %+v", 404, o.Payload)
}

func (o *GetConfigurationNotFound) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetConfigurationNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationConflict creates a GetConfigurationConflict with default headers values
func NewGetConfigurationConflict() *GetConfigurationConflict {
	return &GetConfigurationConflict{}
}

/*GetConfigurationConflict handles this case with default header values.

Conflict
*/
type GetConfigurationConflict struct {
	Payload *models.APIResult
}

func (o *GetConfigurationConflict) Error() string {
	return fmt.Sprintf("[GET /configuration][%d] getConfigurationConflict  %+v", 409, o.Payload)
}

func (o *GetConfigurationConflict) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetConfigurationConflict) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetConfigurationInternalServerError creates a GetConfigurationInternalServerError with default headers values
func NewGetConfigurationInternalServerError() *GetConfigurationInternalServerError {
	return &GetConfigurationInternalServerError{}
}

/*GetConfigurationInternalServerError handles this case with default header values.

Internal Server Error
*/
type GetConfigurationInternalServerError struct {
	Payload *models.APIResult
}

func (o *GetConfigurationInternalServerError) Error() string {
	return fmt.Sprintf("[GET /configuration][%d] getConfigurationInternalServerError  %+v", 500, o.Payload)
}

func (o *GetConfigurationInternalServerError) GetPayload() *models.APIResult {
	return o.Payload
}

func (o *GetConfigurationInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIResult)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
