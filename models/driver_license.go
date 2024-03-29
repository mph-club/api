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

// DriverLicense Driver License Object
// swagger:model DriverLicense
type DriverLicense struct {

	// address
	Address *Address `json:"address,omitempty"`

	// city
	// Required: true
	City *string `json:"city"`

	// created by
	CreatedBy *UserAccountRef `json:"createdBy,omitempty"`

	// created time
	// Format: date-time
	CreatedTime strfmt.DateTime `json:"createdTime,omitempty"`

	// dl number
	DlNumber string `json:"dlNumber,omitempty"`

	// dl type
	DlType string `json:"dlType,omitempty"`

	// dob
	// Format: date-time
	Dob strfmt.DateTime `json:"dob,omitempty"`

	// endorsement
	Endorsement string `json:"endorsement,omitempty"`

	// expiration date
	// Format: date-time
	ExpirationDate strfmt.DateTime `json:"expirationDate,omitempty"`

	// gender
	Gender string `json:"gender,omitempty"`

	// id
	ID string `json:"id,omitempty"`

	// issue date
	// Format: date-time
	IssueDate strfmt.DateTime `json:"issueDate,omitempty"`

	// name
	Name *Name `json:"name,omitempty"`

	// photo back
	PhotoBack *Picture `json:"photoBack,omitempty"`

	// photo front
	PhotoFront *Picture `json:"photoFront,omitempty"`

	// photo with user
	PhotoWithUser *Picture `json:"photoWithUser,omitempty"`

	// state
	State string `json:"state,omitempty"`

	// status
	Status string `json:"status,omitempty"`

	// updated by
	UpdatedBy *UserAccountRef `json:"updatedBy,omitempty"`

	// updated time
	// Format: date-time
	UpdatedTime strfmt.DateTime `json:"updatedTime,omitempty"`
}

// Validate validates this driver license
func (m *DriverLicense) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAddress(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedBy(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedTime(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDob(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateExpirationDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIssueDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotoBack(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotoFront(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePhotoWithUser(formats); err != nil {
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

func (m *DriverLicense) validateAddress(formats strfmt.Registry) error {

	if swag.IsZero(m.Address) { // not required
		return nil
	}

	if m.Address != nil {
		if err := m.Address.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("address")
			}
			return err
		}
	}

	return nil
}

func (m *DriverLicense) validateCity(formats strfmt.Registry) error {

	if err := validate.Required("city", "body", m.City); err != nil {
		return err
	}

	return nil
}

func (m *DriverLicense) validateCreatedBy(formats strfmt.Registry) error {

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

func (m *DriverLicense) validateCreatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("createdTime", "body", "date-time", m.CreatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DriverLicense) validateDob(formats strfmt.Registry) error {

	if swag.IsZero(m.Dob) { // not required
		return nil
	}

	if err := validate.FormatOf("dob", "body", "date-time", m.Dob.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DriverLicense) validateExpirationDate(formats strfmt.Registry) error {

	if swag.IsZero(m.ExpirationDate) { // not required
		return nil
	}

	if err := validate.FormatOf("expirationDate", "body", "date-time", m.ExpirationDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DriverLicense) validateIssueDate(formats strfmt.Registry) error {

	if swag.IsZero(m.IssueDate) { // not required
		return nil
	}

	if err := validate.FormatOf("issueDate", "body", "date-time", m.IssueDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *DriverLicense) validateName(formats strfmt.Registry) error {

	if swag.IsZero(m.Name) { // not required
		return nil
	}

	if m.Name != nil {
		if err := m.Name.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("name")
			}
			return err
		}
	}

	return nil
}

func (m *DriverLicense) validatePhotoBack(formats strfmt.Registry) error {

	if swag.IsZero(m.PhotoBack) { // not required
		return nil
	}

	if m.PhotoBack != nil {
		if err := m.PhotoBack.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photoBack")
			}
			return err
		}
	}

	return nil
}

func (m *DriverLicense) validatePhotoFront(formats strfmt.Registry) error {

	if swag.IsZero(m.PhotoFront) { // not required
		return nil
	}

	if m.PhotoFront != nil {
		if err := m.PhotoFront.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photoFront")
			}
			return err
		}
	}

	return nil
}

func (m *DriverLicense) validatePhotoWithUser(formats strfmt.Registry) error {

	if swag.IsZero(m.PhotoWithUser) { // not required
		return nil
	}

	if m.PhotoWithUser != nil {
		if err := m.PhotoWithUser.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("photoWithUser")
			}
			return err
		}
	}

	return nil
}

func (m *DriverLicense) validateUpdatedBy(formats strfmt.Registry) error {

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

func (m *DriverLicense) validateUpdatedTime(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedTime) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedTime", "body", "date-time", m.UpdatedTime.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DriverLicense) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DriverLicense) UnmarshalBinary(b []byte) error {
	var res DriverLicense
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
