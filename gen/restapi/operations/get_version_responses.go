// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/mwinters-stuff/halo-one-thing/gen/models"
)

// GetVersionOKCode is the HTTP code returned for type GetVersionOK
const GetVersionOKCode int = 200

/*
GetVersionOK OK

swagger:response getVersionOK
*/
type GetVersionOK struct {

	/*
	  In: Body
	*/
	Payload *models.Version `json:"body,omitempty"`
}

// NewGetVersionOK creates GetVersionOK with default headers values
func NewGetVersionOK() *GetVersionOK {

	return &GetVersionOK{}
}

// WithPayload adds the payload to the get version o k response
func (o *GetVersionOK) WithPayload(payload *models.Version) *GetVersionOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get version o k response
func (o *GetVersionOK) SetPayload(payload *models.Version) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetVersionConflictCode is the HTTP code returned for type GetVersionConflict
const GetVersionConflictCode int = 409

/*
GetVersionConflict Failed

swagger:response getVersionConflict
*/
type GetVersionConflict struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetVersionConflict creates GetVersionConflict with default headers values
func NewGetVersionConflict() *GetVersionConflict {

	return &GetVersionConflict{}
}

// WithPayload adds the payload to the get version conflict response
func (o *GetVersionConflict) WithPayload(payload *models.Error) *GetVersionConflict {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get version conflict response
func (o *GetVersionConflict) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetVersionConflict) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(409)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
