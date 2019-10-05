// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Detail detail
// swagger:model Detail
type Detail struct {

	// The location of the field that threw the error.
	Location string `json:"location,omitempty"`

	// The name of the field that threw the error.
	Name string `json:"name,omitempty"`
}

// Validate validates this detail
func (m *Detail) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Detail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Detail) UnmarshalBinary(b []byte) error {
	var res Detail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}