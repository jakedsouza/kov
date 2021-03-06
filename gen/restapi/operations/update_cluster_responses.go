///////////////////////////////////////////////////////////////////////
// Copyright (C) 2017 VMware, Inc. All rights reserved.
// -- VMware Confidential
///////////////////////////////////////////////////////////////////////
package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/supervised-io/kov/gen/models"
)

// UpdateClusterAcceptedCode is the HTTP code returned for type UpdateClusterAccepted
const UpdateClusterAcceptedCode int = 202

/*UpdateClusterAccepted update cluster task has been accepted

swagger:response updateClusterAccepted
*/
type UpdateClusterAccepted struct {

	/*
	  In: Body
	*/
	Payload models.TaskID `json:"body,omitempty"`
}

// NewUpdateClusterAccepted creates UpdateClusterAccepted with default headers values
func NewUpdateClusterAccepted() *UpdateClusterAccepted {
	return &UpdateClusterAccepted{}
}

// WithPayload adds the payload to the update cluster accepted response
func (o *UpdateClusterAccepted) WithPayload(payload models.TaskID) *UpdateClusterAccepted {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster accepted response
func (o *UpdateClusterAccepted) SetPayload(payload models.TaskID) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterAccepted) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(202)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

// UpdateClusterNotFoundCode is the HTTP code returned for type UpdateClusterNotFound
const UpdateClusterNotFoundCode int = 404

/*UpdateClusterNotFound The cluster was not found

swagger:response updateClusterNotFound
*/
type UpdateClusterNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterNotFound creates UpdateClusterNotFound with default headers values
func NewUpdateClusterNotFound() *UpdateClusterNotFound {
	return &UpdateClusterNotFound{}
}

// WithPayload adds the payload to the update cluster not found response
func (o *UpdateClusterNotFound) WithPayload(payload *models.Error) *UpdateClusterNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster not found response
func (o *UpdateClusterNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

/*UpdateClusterDefault Error

swagger:response updateClusterDefault
*/
type UpdateClusterDefault struct {
	_statusCode int

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewUpdateClusterDefault creates UpdateClusterDefault with default headers values
func NewUpdateClusterDefault(code int) *UpdateClusterDefault {
	if code <= 0 {
		code = 500
	}

	return &UpdateClusterDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the update cluster default response
func (o *UpdateClusterDefault) WithStatusCode(code int) *UpdateClusterDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the update cluster default response
func (o *UpdateClusterDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the update cluster default response
func (o *UpdateClusterDefault) WithPayload(payload *models.Error) *UpdateClusterDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the update cluster default response
func (o *UpdateClusterDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *UpdateClusterDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
