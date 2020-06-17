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

package config

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/openziti/edge/rest_model"
)

// DetailConfigReader is a Reader for the DetailConfig structure.
type DetailConfigReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DetailConfigReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDetailConfigOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewDetailConfigUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewDetailConfigNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDetailConfigOK creates a DetailConfigOK with default headers values
func NewDetailConfigOK() *DetailConfigOK {
	return &DetailConfigOK{}
}

/*DetailConfigOK handles this case with default header values.

A singular config resource
*/
type DetailConfigOK struct {
	Payload *rest_model.DetailConfigEnvelope
}

func (o *DetailConfigOK) Error() string {
	return fmt.Sprintf("[GET /configs/{id}][%d] detailConfigOK  %+v", 200, o.Payload)
}

func (o *DetailConfigOK) GetPayload() *rest_model.DetailConfigEnvelope {
	return o.Payload
}

func (o *DetailConfigOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.DetailConfigEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailConfigUnauthorized creates a DetailConfigUnauthorized with default headers values
func NewDetailConfigUnauthorized() *DetailConfigUnauthorized {
	return &DetailConfigUnauthorized{}
}

/*DetailConfigUnauthorized handles this case with default header values.

The currently supplied session does not have the correct access rights to request this resource
*/
type DetailConfigUnauthorized struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailConfigUnauthorized) Error() string {
	return fmt.Sprintf("[GET /configs/{id}][%d] detailConfigUnauthorized  %+v", 401, o.Payload)
}

func (o *DetailConfigUnauthorized) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailConfigUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewDetailConfigNotFound creates a DetailConfigNotFound with default headers values
func NewDetailConfigNotFound() *DetailConfigNotFound {
	return &DetailConfigNotFound{}
}

/*DetailConfigNotFound handles this case with default header values.

The requested resource does not exist
*/
type DetailConfigNotFound struct {
	Payload *rest_model.APIErrorEnvelope
}

func (o *DetailConfigNotFound) Error() string {
	return fmt.Sprintf("[GET /configs/{id}][%d] detailConfigNotFound  %+v", 404, o.Payload)
}

func (o *DetailConfigNotFound) GetPayload() *rest_model.APIErrorEnvelope {
	return o.Payload
}

func (o *DetailConfigNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(rest_model.APIErrorEnvelope)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}