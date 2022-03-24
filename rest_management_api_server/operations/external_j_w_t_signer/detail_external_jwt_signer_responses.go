// Code generated by go-swagger; DO NOT EDIT.

//
// Copyright NetFoundry, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// __          __              _
// \ \        / /             (_)
//  \ \  /\  / /_ _ _ __ _ __  _ _ __   __ _
//   \ \/  \/ / _` | '__| '_ \| | '_ \ / _` |
//    \  /\  / (_| | |  | | | | | | | | (_| | : This file is generated, do not edit it.
//     \/  \/ \__,_|_|  |_| |_|_|_| |_|\__, |
//                                      __/ |
//                                     |___/

package external_j_w_t_signer

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// DetailExternalJwtSignerOKCode is the HTTP code returned for type DetailExternalJwtSignerOK
const DetailExternalJwtSignerOKCode int = 200

/*DetailExternalJwtSignerOK A singular External JWT Signer resource

swagger:response detailExternalJwtSignerOK
*/
type DetailExternalJwtSignerOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.DetailExternalJwtSignerEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJwtSignerOK creates DetailExternalJwtSignerOK with default headers values
func NewDetailExternalJwtSignerOK() *DetailExternalJwtSignerOK {

	return &DetailExternalJwtSignerOK{}
}

// WithPayload adds the payload to the detail external jwt signer o k response
func (o *DetailExternalJwtSignerOK) WithPayload(payload *rest_model.DetailExternalJwtSignerEnvelope) *DetailExternalJwtSignerOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external jwt signer o k response
func (o *DetailExternalJwtSignerOK) SetPayload(payload *rest_model.DetailExternalJwtSignerEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJwtSignerOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailExternalJwtSignerUnauthorizedCode is the HTTP code returned for type DetailExternalJwtSignerUnauthorized
const DetailExternalJwtSignerUnauthorizedCode int = 401

/*DetailExternalJwtSignerUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response detailExternalJwtSignerUnauthorized
*/
type DetailExternalJwtSignerUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJwtSignerUnauthorized creates DetailExternalJwtSignerUnauthorized with default headers values
func NewDetailExternalJwtSignerUnauthorized() *DetailExternalJwtSignerUnauthorized {

	return &DetailExternalJwtSignerUnauthorized{}
}

// WithPayload adds the payload to the detail external jwt signer unauthorized response
func (o *DetailExternalJwtSignerUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailExternalJwtSignerUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external jwt signer unauthorized response
func (o *DetailExternalJwtSignerUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJwtSignerUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// DetailExternalJwtSignerNotFoundCode is the HTTP code returned for type DetailExternalJwtSignerNotFound
const DetailExternalJwtSignerNotFoundCode int = 404

/*DetailExternalJwtSignerNotFound The requested resource does not exist

swagger:response detailExternalJwtSignerNotFound
*/
type DetailExternalJwtSignerNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewDetailExternalJwtSignerNotFound creates DetailExternalJwtSignerNotFound with default headers values
func NewDetailExternalJwtSignerNotFound() *DetailExternalJwtSignerNotFound {

	return &DetailExternalJwtSignerNotFound{}
}

// WithPayload adds the payload to the detail external jwt signer not found response
func (o *DetailExternalJwtSignerNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *DetailExternalJwtSignerNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the detail external jwt signer not found response
func (o *DetailExternalJwtSignerNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *DetailExternalJwtSignerNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}