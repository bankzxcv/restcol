// Code generated by go-swagger; DO NOT EDIT.

package swagger

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/footprintai/restcol/api/go-openapiv2/models"
)

// RestColServiceGetSwaggerDocReader is a Reader for the RestColServiceGetSwaggerDoc structure.
type RestColServiceGetSwaggerDocReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestColServiceGetSwaggerDocReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestColServiceGetSwaggerDocOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestColServiceGetSwaggerDocDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestColServiceGetSwaggerDocOK creates a RestColServiceGetSwaggerDocOK with default headers values
func NewRestColServiceGetSwaggerDocOK() *RestColServiceGetSwaggerDocOK {
	return &RestColServiceGetSwaggerDocOK{}
}

/*
RestColServiceGetSwaggerDocOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestColServiceGetSwaggerDocOK struct {
	Payload *models.APIHTTPBody
}

// IsSuccess returns true when this rest col service get swagger doc o k response has a 2xx status code
func (o *RestColServiceGetSwaggerDocOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rest col service get swagger doc o k response has a 3xx status code
func (o *RestColServiceGetSwaggerDocOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rest col service get swagger doc o k response has a 4xx status code
func (o *RestColServiceGetSwaggerDocOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rest col service get swagger doc o k response has a 5xx status code
func (o *RestColServiceGetSwaggerDocOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rest col service get swagger doc o k response a status code equal to that given
func (o *RestColServiceGetSwaggerDocOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rest col service get swagger doc o k response
func (o *RestColServiceGetSwaggerDocOK) Code() int {
	return 200
}

func (o *RestColServiceGetSwaggerDocOK) Error() string {
	return fmt.Sprintf("[GET /v1/apidoc][%d] restColServiceGetSwaggerDocOK  %+v", 200, o.Payload)
}

func (o *RestColServiceGetSwaggerDocOK) String() string {
	return fmt.Sprintf("[GET /v1/apidoc][%d] restColServiceGetSwaggerDocOK  %+v", 200, o.Payload)
}

func (o *RestColServiceGetSwaggerDocOK) GetPayload() *models.APIHTTPBody {
	return o.Payload
}

func (o *RestColServiceGetSwaggerDocOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIHTTPBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestColServiceGetSwaggerDocDefault creates a RestColServiceGetSwaggerDocDefault with default headers values
func NewRestColServiceGetSwaggerDocDefault(code int) *RestColServiceGetSwaggerDocDefault {
	return &RestColServiceGetSwaggerDocDefault{
		_statusCode: code,
	}
}

/*
RestColServiceGetSwaggerDocDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestColServiceGetSwaggerDocDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this rest col service get swagger doc default response has a 2xx status code
func (o *RestColServiceGetSwaggerDocDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rest col service get swagger doc default response has a 3xx status code
func (o *RestColServiceGetSwaggerDocDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rest col service get swagger doc default response has a 4xx status code
func (o *RestColServiceGetSwaggerDocDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rest col service get swagger doc default response has a 5xx status code
func (o *RestColServiceGetSwaggerDocDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rest col service get swagger doc default response a status code equal to that given
func (o *RestColServiceGetSwaggerDocDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rest col service get swagger doc default response
func (o *RestColServiceGetSwaggerDocDefault) Code() int {
	return o._statusCode
}

func (o *RestColServiceGetSwaggerDocDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apidoc][%d] RestColService_GetSwaggerDoc default  %+v", o._statusCode, o.Payload)
}

func (o *RestColServiceGetSwaggerDocDefault) String() string {
	return fmt.Sprintf("[GET /v1/apidoc][%d] RestColService_GetSwaggerDoc default  %+v", o._statusCode, o.Payload)
}

func (o *RestColServiceGetSwaggerDocDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RestColServiceGetSwaggerDocDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}