// Code generated by go-swagger; DO NOT EDIT.

package hello

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/kozmod/idea-tests/open-api/client/internal/generated/models"
)

// HelloWorldFullOKCode is the HTTP code returned for type HelloWorldFullOK
const HelloWorldFullOKCode int = 200

/*HelloWorldFullOK successful response

swagger:response helloWorldFullOK
*/
type HelloWorldFullOK struct {

	/*
	  In: Body
	*/
	Payload *models.HelloWorld `json:"body,omitempty"`
}

// NewHelloWorldFullOK creates HelloWorldFullOK with default headers values
func NewHelloWorldFullOK() *HelloWorldFullOK {

	return &HelloWorldFullOK{}
}

// WithPayload adds the payload to the hello world full o k response
func (o *HelloWorldFullOK) WithPayload(payload *models.HelloWorld) *HelloWorldFullOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the hello world full o k response
func (o *HelloWorldFullOK) SetPayload(payload *models.HelloWorld) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HelloWorldFullOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HelloWorldFullBadRequestCode is the HTTP code returned for type HelloWorldFullBadRequest
const HelloWorldFullBadRequestCode int = 400

/*HelloWorldFullBadRequest Bad request

swagger:response helloWorldFullBadRequest
*/
type HelloWorldFullBadRequest struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewHelloWorldFullBadRequest creates HelloWorldFullBadRequest with default headers values
func NewHelloWorldFullBadRequest() *HelloWorldFullBadRequest {

	return &HelloWorldFullBadRequest{}
}

// WithPayload adds the payload to the hello world full bad request response
func (o *HelloWorldFullBadRequest) WithPayload(payload *models.Error) *HelloWorldFullBadRequest {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the hello world full bad request response
func (o *HelloWorldFullBadRequest) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HelloWorldFullBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// HelloWorldFullInternalServerErrorCode is the HTTP code returned for type HelloWorldFullInternalServerError
const HelloWorldFullInternalServerErrorCode int = 500

/*HelloWorldFullInternalServerError Internal server error

swagger:response helloWorldFullInternalServerError
*/
type HelloWorldFullInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewHelloWorldFullInternalServerError creates HelloWorldFullInternalServerError with default headers values
func NewHelloWorldFullInternalServerError() *HelloWorldFullInternalServerError {

	return &HelloWorldFullInternalServerError{}
}

// WithPayload adds the payload to the hello world full internal server error response
func (o *HelloWorldFullInternalServerError) WithPayload(payload *models.Error) *HelloWorldFullInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the hello world full internal server error response
func (o *HelloWorldFullInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *HelloWorldFullInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
