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

// Address Address Object
// swagger:model Address
type Address struct {

	// address line1
	AddressLine1 string `json:"addressLine1,omitempty"`

	// address line2
	AddressLine2 string `json:"addressLine2,omitempty"`

	// address type
	AddressType string `json:"addressType,omitempty"`

	// city
	// Required: true
	City *string `json:"city"`

	// country
	Country string `json:"country,omitempty"`

	// created by
	CreatedBy *UserAccountRef `json:"createdBy,omitempty"`

	// created time
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"createdTime,omitempty"`

	// full address
	FullAddress string `json:"fullAddress,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// latitude
	Latitude string `json:"latitude,omitempty"`

	// longitude
	Longitude string `json:"longitude,omitempty"`

	// place name
	PlaceName string `json:"placeName,omitempty"`

	// postal code
	PostalCode string `json:"postalCode,omitempty"`

	// province
	Province string `json:"province,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// updated by
	UpdatedBy *UserAccountRef `json:"updatedBy,omitempty"`

	// updated time
	// Format: date-time
	UpdatedTime strfmt.DateTime `json:"updatedTime,omitempty"`
}

// Validate validates this address
func (m *Address) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedTime(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Address) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateCreatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedBy) { // not required
		return nil
	}

	if m.CreatedBy != nil {
		if err := m.CreatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("createdBy")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Address) validateUpdatedBy(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedBy) { // not required
		return nil
	}

	if m.UpdatedBy != nil {
		if err := m.UpdatedBy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("updatedBy")
			}
			return err
		}
	}

	return nil
}

func (m *Address) validateUpdatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedTime", "body", "date-time", m.UpdatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Address) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Address) UnmarshalBinary(b []byte) error {
	var res Address
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
