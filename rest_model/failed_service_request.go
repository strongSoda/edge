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

package rest_model

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// FailedServiceRequest failed service request
//
// swagger:model failedServiceRequest
type FailedServiceRequest struct {

	// api session Id
	APISessionID string `json:"apiSessionId,omitempty"`

	// policy failures
	PolicyFailures []*PolicyFailure `json:"policyFailures"`

	// service Id
	ServiceID string `json:"serviceId,omitempty"`

	// service name
	ServiceName string `json:"serviceName,omitempty"`

	// session type
	SessionType DialBind `json:"sessionType,omitempty"`

	// when
	// Format: date-time
	When strfmt.DateTime `json:"when,omitempty"`
}

// Validate validates this failed service request
func (m *FailedServiceRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validatePolicyFailures(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSessionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateWhen(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FailedServiceRequest) validatePolicyFailures(formats strfmt.Registry) error {
	if swag.IsZero(m.PolicyFailures) { // not required
		return nil
	}

	for i := 0; i < len(m.PolicyFailures); i++ {
		if swag.IsZero(m.PolicyFailures[i]) { // not required
			continue
		}

		if m.PolicyFailures[i] != nil {
			if err := m.PolicyFailures[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyFailures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FailedServiceRequest) validateSessionType(formats strfmt.Registry) error {
	if swag.IsZero(m.SessionType) { // not required
		return nil
	}

	if err := m.SessionType.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sessionType")
		}
		return err
	}

	return nil
}

func (m *FailedServiceRequest) validateWhen(formats strfmt.Registry) error {
	if swag.IsZero(m.When) { // not required
		return nil
	}

	if err := validate.FormatOf("when", "body", "date-time", m.When.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this failed service request based on the context it is used
func (m *FailedServiceRequest) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidatePolicyFailures(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateSessionType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *FailedServiceRequest) contextValidatePolicyFailures(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.PolicyFailures); i++ {

		if m.PolicyFailures[i] != nil {
			if err := m.PolicyFailures[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("policyFailures" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *FailedServiceRequest) contextValidateSessionType(ctx context.Context, formats strfmt.Registry) error {

	if err := m.SessionType.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("sessionType")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *FailedServiceRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *FailedServiceRequest) UnmarshalBinary(b []byte) error {
	var res FailedServiceRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}