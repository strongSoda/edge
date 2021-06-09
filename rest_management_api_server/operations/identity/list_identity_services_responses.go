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

package identity

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/openziti/edge/rest_model"
)

// ListIdentityServicesOKCode is the HTTP code returned for type ListIdentityServicesOK
const ListIdentityServicesOKCode int = 200

/*ListIdentityServicesOK A list of edge routers

swagger:response listIdentityServicesOK
*/
type ListIdentityServicesOK struct {

	/*
	  In: Body
	*/
	Payload *rest_model.ListEdgeRoutersEnvelope `json:"body,omitempty"`
}

// NewListIdentityServicesOK creates ListIdentityServicesOK with default headers values
func NewListIdentityServicesOK() *ListIdentityServicesOK {

	return &ListIdentityServicesOK{}
}

// WithPayload adds the payload to the list identity services o k response
func (o *ListIdentityServicesOK) WithPayload(payload *rest_model.ListEdgeRoutersEnvelope) *ListIdentityServicesOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identity services o k response
func (o *ListIdentityServicesOK) SetPayload(payload *rest_model.ListEdgeRoutersEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentityServicesOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentityServicesUnauthorizedCode is the HTTP code returned for type ListIdentityServicesUnauthorized
const ListIdentityServicesUnauthorizedCode int = 401

/*ListIdentityServicesUnauthorized The currently supplied session does not have the correct access rights to request this resource

swagger:response listIdentityServicesUnauthorized
*/
type ListIdentityServicesUnauthorized struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentityServicesUnauthorized creates ListIdentityServicesUnauthorized with default headers values
func NewListIdentityServicesUnauthorized() *ListIdentityServicesUnauthorized {

	return &ListIdentityServicesUnauthorized{}
}

// WithPayload adds the payload to the list identity services unauthorized response
func (o *ListIdentityServicesUnauthorized) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentityServicesUnauthorized {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identity services unauthorized response
func (o *ListIdentityServicesUnauthorized) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentityServicesUnauthorized) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(401)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// ListIdentityServicesNotFoundCode is the HTTP code returned for type ListIdentityServicesNotFound
const ListIdentityServicesNotFoundCode int = 404

/*ListIdentityServicesNotFound The requested resource does not exist

swagger:response listIdentityServicesNotFound
*/
type ListIdentityServicesNotFound struct {

	/*
	  In: Body
	*/
	Payload *rest_model.APIErrorEnvelope `json:"body,omitempty"`
}

// NewListIdentityServicesNotFound creates ListIdentityServicesNotFound with default headers values
func NewListIdentityServicesNotFound() *ListIdentityServicesNotFound {

	return &ListIdentityServicesNotFound{}
}

// WithPayload adds the payload to the list identity services not found response
func (o *ListIdentityServicesNotFound) WithPayload(payload *rest_model.APIErrorEnvelope) *ListIdentityServicesNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the list identity services not found response
func (o *ListIdentityServicesNotFound) SetPayload(payload *rest_model.APIErrorEnvelope) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *ListIdentityServicesNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}