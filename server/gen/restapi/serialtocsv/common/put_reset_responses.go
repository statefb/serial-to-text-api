// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"local.packages/gen/models"
)

// PutResetOKCode is the HTTP code returned for type PutResetOK
const PutResetOKCode int = 200

/*PutResetOK OK

swagger:response putResetOK
*/
type PutResetOK struct {

	/*
	  In: Body
	*/
	Payload interface{} `json:"body,omitempty"`
}

// NewPutResetOK creates PutResetOK with default headers values
func NewPutResetOK() *PutResetOK {

	return &PutResetOK{}
}

// WithPayload adds the payload to the put reset o k response
func (o *PutResetOK) WithPayload(payload interface{}) *PutResetOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put reset o k response
func (o *PutResetOK) SetPayload(payload interface{}) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutResetOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*PutResetDefault Unexpected error

swagger:response putResetDefault
*/
type PutResetDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.ErrorV1 `json:"body,omitempty"`
}

// NewPutResetDefault creates PutResetDefault with default headers values
func NewPutResetDefault(code int) *PutResetDefault {
	if code <= 0 {
		code = 500
	}

	return &PutResetDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the put reset default response
func (o *PutResetDefault) WithStatusCode(code int) *PutResetDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the put reset default response
func (o *PutResetDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the put reset default response
func (o *PutResetDefault) WithPayload(payload *models.ErrorV1) *PutResetDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the put reset default response
func (o *PutResetDefault) SetPayload(payload *models.ErrorV1) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PutResetDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
