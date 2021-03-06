// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// ErrorV1 Error
//
// A standard error object.
//
// swagger:model Error.v1
type ErrorV1 struct {

	// code
	// Required: true
	Code *string `json:"code"`

	// message
	// Required: true
	Message *string `json:"message"`
}

// UnmarshalJSON unmarshals this object while disallowing additional properties from JSON
func (m *ErrorV1) UnmarshalJSON(data []byte) error {
	var props struct {

		// code
		// Required: true
		Code *string `json:"code"`

		// message
		// Required: true
		Message *string `json:"message"`
	}

	dec := json.NewDecoder(bytes.NewReader(data))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&props); err != nil {
		return err
	}

	m.Code = props.Code
	m.Message = props.Message
	return nil
}

// Validate validates this error v1
func (m *ErrorV1) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMessage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ErrorV1) validateCode(formats strfmt.Registry) error {

	if err := validate.Required("code", "body", m.Code); err != nil {
		return err
	}

	return nil
}

func (m *ErrorV1) validateMessage(formats strfmt.Registry) error {

	if err := validate.Required("message", "body", m.Message); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this error v1 based on context it is used
func (m *ErrorV1) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *ErrorV1) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ErrorV1) UnmarshalBinary(b []byte) error {
	var res ErrorV1
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
