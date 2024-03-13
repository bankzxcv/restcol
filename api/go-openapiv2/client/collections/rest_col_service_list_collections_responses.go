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

// RestColServiceListCollectionsReader is a Reader for the RestColServiceListCollections structure.
type RestColServiceListCollectionsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RestColServiceListCollectionsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRestColServiceListCollectionsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		result := NewRestColServiceListCollectionsDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewRestColServiceListCollectionsOK creates a RestColServiceListCollectionsOK with default headers values
func NewRestColServiceListCollectionsOK() *RestColServiceListCollectionsOK {
	return &RestColServiceListCollectionsOK{}
}

/*
RestColServiceListCollectionsOK describes a response with status code 200, with default header values.

A successful response.
*/
type RestColServiceListCollectionsOK struct {
	Payload models.APIListCollectionsResponse
}

// IsSuccess returns true when this rest col service list collections o k response has a 2xx status code
func (o *RestColServiceListCollectionsOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this rest col service list collections o k response has a 3xx status code
func (o *RestColServiceListCollectionsOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this rest col service list collections o k response has a 4xx status code
func (o *RestColServiceListCollectionsOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this rest col service list collections o k response has a 5xx status code
func (o *RestColServiceListCollectionsOK) IsServerError() bool {
	return false
}

// IsCode returns true when this rest col service list collections o k response a status code equal to that given
func (o *RestColServiceListCollectionsOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the rest col service list collections o k response
func (o *RestColServiceListCollectionsOK) Code() int {
	return 200
}

func (o *RestColServiceListCollectionsOK) Error() string {
	return fmt.Sprintf("[GET /v1/collections][%d] restColServiceListCollectionsOK  %+v", 200, o.Payload)
}

func (o *RestColServiceListCollectionsOK) String() string {
	return fmt.Sprintf("[GET /v1/collections][%d] restColServiceListCollectionsOK  %+v", 200, o.Payload)
}

func (o *RestColServiceListCollectionsOK) GetPayload() models.APIListCollectionsResponse {
	return o.Payload
}

func (o *RestColServiceListCollectionsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response payload
	if err := consumer.Consume(response.Body(), &o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRestColServiceListCollectionsDefault creates a RestColServiceListCollectionsDefault with default headers values
func NewRestColServiceListCollectionsDefault(code int) *RestColServiceListCollectionsDefault {
	return &RestColServiceListCollectionsDefault{
		_statusCode: code,
	}
}

/*
RestColServiceListCollectionsDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type RestColServiceListCollectionsDefault struct {
	_statusCode int

	Payload *models.RPCStatus
}

// IsSuccess returns true when this rest col service list collections default response has a 2xx status code
func (o *RestColServiceListCollectionsDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this rest col service list collections default response has a 3xx status code
func (o *RestColServiceListCollectionsDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this rest col service list collections default response has a 4xx status code
func (o *RestColServiceListCollectionsDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this rest col service list collections default response has a 5xx status code
func (o *RestColServiceListCollectionsDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this rest col service list collections default response a status code equal to that given
func (o *RestColServiceListCollectionsDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the rest col service list collections default response
func (o *RestColServiceListCollectionsDefault) Code() int {
	return o._statusCode
}

func (o *RestColServiceListCollectionsDefault) Error() string {
	return fmt.Sprintf("[GET /v1/collections][%d] RestColService_ListCollections default  %+v", o._statusCode, o.Payload)
}

func (o *RestColServiceListCollectionsDefault) String() string {
	return fmt.Sprintf("[GET /v1/collections][%d] RestColService_ListCollections default  %+v", o._statusCode, o.Payload)
}

func (o *RestColServiceListCollectionsDefault) GetPayload() *models.RPCStatus {
	return o.Payload
}

func (o *RestColServiceListCollectionsDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RPCStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}