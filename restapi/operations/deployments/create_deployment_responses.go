package deployments

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"io"
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/luizalabs/teresa-api/models"
)

/*CreateDeploymentOK App

swagger:response createDeploymentOK
*/
type CreateDeploymentOK struct {

	// In: body
	Payload io.ReadCloser `json:"body,omitempty"`
}

// NewCreateDeploymentOK creates CreateDeploymentOK with default headers values
func NewCreateDeploymentOK() *CreateDeploymentOK {
	return &CreateDeploymentOK{}
}

// WithPayload adds the payload to the create deployment o k response
func (o *CreateDeploymentOK) WithPayload(payload io.ReadCloser) *CreateDeploymentOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create deployment o k response
func (o *CreateDeploymentOK) SetPayload(payload io.ReadCloser) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeploymentOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if err := producer.Produce(rw, o.Payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}

}

/*CreateDeploymentDefault Error

swagger:response createDeploymentDefault
*/
type CreateDeploymentDefault struct {
	_statusCode int

	// In: body
	Payload *models.Error `json:"body,omitempty"`
}

// NewCreateDeploymentDefault creates CreateDeploymentDefault with default headers values
func NewCreateDeploymentDefault(code int) *CreateDeploymentDefault {
	if code <= 0 {
		code = 500
	}

	return &CreateDeploymentDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the create deployment default response
func (o *CreateDeploymentDefault) WithStatusCode(code int) *CreateDeploymentDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the create deployment default response
func (o *CreateDeploymentDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WithPayload adds the payload to the create deployment default response
func (o *CreateDeploymentDefault) WithPayload(payload *models.Error) *CreateDeploymentDefault {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the create deployment default response
func (o *CreateDeploymentDefault) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *CreateDeploymentDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(o._statusCode)
	if o.Payload != nil {
		if err := producer.Produce(rw, o.Payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}