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

package authenticator

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// DetailAuthenticatorReader is a Reader for the DetailAuthenticator structure.
type DetailAuthenticatorReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailAuthenticatorReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailAuthenticatorOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailAuthenticatorUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailAuthenticatorNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetailAuthenticatorOK creates a DetailAuthenticatorOK with default headers values
func NewDetailAuthenticatorOK() *DetailAuthenticatorOK {
	return &DetailAuthenticatorOK{}
}

/*DetailAuthenticatorOK handles this case with default header values.

A singular authenticator resource
*/
type DetailAuthenticatorOK struct {
	Payload *rest_model.DetailAuthenticatorEnvelope
}

func (o *DetailAuthenticatorOK) Error() string {
	return fmt.Sprintf("[GET /authenticators/{id}][%d] detailAuthenticatorOK  %+v", 200, o.Payload)
}

func (o *DetailAuthenticatorOK) GetPayload() *rest_model.DetailAuthenticatorEnvelope {
	return o.Payload
}

func (o *DetailAuthenticatorOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailAuthenticatorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailAuthenticatorUnauthorized creates a DetailAuthenticatorUnauthorized with default headers values
func NewDetailAuthenticatorUnauthorized() *DetailAuthenticatorUnauthorized {
	return &DetailAuthenticatorUnauthorized{}
}

/*DetailAuthenticatorUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DetailAuthenticatorUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailAuthenticatorUnauthorized) Error() string {
	return fmt.Sprintf("[GET /authenticators/{id}][%d] detailAuthenticatorUnauthorized  %+v", 401, o.Payload)
}

func (o *DetailAuthenticatorUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailAuthenticatorUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailAuthenticatorNotFound creates a DetailAuthenticatorNotFound with default headers values
func NewDetailAuthenticatorNotFound() *DetailAuthenticatorNotFound {
	return &DetailAuthenticatorNotFound{}
}

/*DetailAuthenticatorNotFound handles this case with default header values.

The requested resource does not exist
*/
type DetailAuthenticatorNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailAuthenticatorNotFound) Error() string {
	return fmt.Sprintf("[GET /authenticators/{id}][%d] detailAuthenticatorNotFound  %+v", 404, o.Payload)
}

func (o *DetailAuthenticatorNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailAuthenticatorNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}