// Code generated by go-swagger; DO NOT EDIT.

package collections

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/footprintai/restcol/api/go-openapiv2/models"
)

// RestColServiceCreateCollection2Reader is a Reader for the RestColServiceCreateCollection2 structure.
type RestColServiceCreateCollection2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestColServiceCreateCollection2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestColServiceCreateCollection2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestColServiceCreateCollection2Default(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestColServiceCreateCollection2OK creates a RestColServiceCreateCollection2OK with default headers values
func NewRestColServiceCreateCollection2OK() *RestColServiceCreateCollection2OK {
	return &RestColServiceCreateCollection2OK{}
}

/*
RestColServiceCreateCollection2OK describes a response with status code 200, with default header values.

A successful response.
*/
type RestColServiceCreateCollection2OK struct {
	Payload *models.APICreateCollectionResponse
}

// IsSuccess returns true when this rest col service create collection2 o k response has a 2xx status code
func (o *RestColServiceCreateCollection2OK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rest col service create collection2 o k response has a 3xx status code
func (o *RestColServiceCreateCollection2OK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rest col service create collection2 o k response has a 4xx status code
func (o *RestColServiceCreateCollection2OK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rest col service create collection2 o k response has a 5xx status code
func (o *RestColServiceCreateCollection2OK) IsServerError() bool {
	return false
}

// IsCode returns true when this rest col service create collection2 o k response a status code equal to that given
func (o *RestColServiceCreateCollection2OK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rest col service create collection2 o k response
func (o *RestColServiceCreateCollection2OK) Code() int {
	return 200
}

func (o *RestColServiceCreateCollection2OK) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{pid}/collections][%d] restColServiceCreateCollection2OK  %+v", 200, o.Payload)
}

func (o *RestColServiceCreateCollection2OK) String() string {
	return fmt.Sprintf("[POST /v1/projects/{pid}/collections][%d] restColServiceCreateCollection2OK  %+v", 200, o.Payload)
}

func (o *RestColServiceCreateCollection2OK) GetPayload() *models.APICreateCollectionResponse {
	return o.Payload
}

func (o *RestColServiceCreateCollection2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APICreateCollectionResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestColServiceCreateCollection2Default creates a RestColServiceCreateCollection2Default with default headers values
func NewRestColServiceCreateCollection2Default(code int) *RestColServiceCreateCollection2Default {
	return &RestColServiceCreateCollection2Default{
		_statusCode: code,
	}
}

/*
RestColServiceCreateCollection2Default describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestColServiceCreateCollection2Default struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this rest col service create collection2 default response has a 2xx status code
func (o *RestColServiceCreateCollection2Default) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rest col service create collection2 default response has a 3xx status code
func (o *RestColServiceCreateCollection2Default) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rest col service create collection2 default response has a 4xx status code
func (o *RestColServiceCreateCollection2Default) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rest col service create collection2 default response has a 5xx status code
func (o *RestColServiceCreateCollection2Default) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rest col service create collection2 default response a status code equal to that given
func (o *RestColServiceCreateCollection2Default) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rest col service create collection2 default response
func (o *RestColServiceCreateCollection2Default) Code() int {
	return o._statusCode
}

func (o *RestColServiceCreateCollection2Default) Error() string {
	return fmt.Sprintf("[POST /v1/projects/{pid}/collections][%d] RestColService_CreateCollection2 default  %+v", o._statusCode, o.Payload)
}

func (o *RestColServiceCreateCollection2Default) String() string {
	return fmt.Sprintf("[POST /v1/projects/{pid}/collections][%d] RestColService_CreateCollection2 default  %+v", o._statusCode, o.Payload)
}

func (o *RestColServiceCreateCollection2Default) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RestColServiceCreateCollection2Default) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}