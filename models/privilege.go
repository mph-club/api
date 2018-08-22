// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Privilege privilege
// swagger:model Privilege
type Privilege struct {

	// application Id
	ApplicationID string `json:"applicationId,omitempty"`

	// privilege
	Privilege string `json:"privilege,omitempty"`
}

// Validate validates this privilege
func (m *Privilege) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Privilege) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Privilege) UnmarshalBinary(b []byte) error {
	var res Privilege
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
