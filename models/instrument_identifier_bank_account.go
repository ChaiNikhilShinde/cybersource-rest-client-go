// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// InstrumentIdentifierBankAccount instrument identifier bank account
// swagger:model InstrumentIdentifierBankAccount
type InstrumentIdentifierBankAccount struct {

	// Checking account number.
	// Max Length: 19
	// Min Length: 1
	Number string `json:"number,omitempty"`

	// Routing number.
	// Max Length: 9
	// Min Length: 1
	RoutingNumber string `json:"routingNumber,omitempty"`
}

// Validate validates this instrument identifier bank account
func (m *InstrumentIdentifierBankAccount) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNumber(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRoutingNumber(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InstrumentIdentifierBankAccount) validateNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.Number) { // not required
		return nil
	}

	if err := validate.MinLength("number", "body", string(m.Number), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("number", "body", string(m.Number), 19); err != nil {
		return err
	}

	return nil
}

func (m *InstrumentIdentifierBankAccount) validateRoutingNumber(formats strfmt.Registry) error {

	if swag.IsZero(m.RoutingNumber) { // not required
		return nil
	}

	if err := validate.MinLength("routingNumber", "body", string(m.RoutingNumber), 1); err != nil {
		return err
	}

	if err := validate.MaxLength("routingNumber", "body", string(m.RoutingNumber), 9); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InstrumentIdentifierBankAccount) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InstrumentIdentifierBankAccount) UnmarshalBinary(b []byte) error {
	var res InstrumentIdentifierBankAccount
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}