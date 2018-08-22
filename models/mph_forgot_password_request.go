// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// MphForgotPasswordRequest mph forgot password request
// swagger:model MphForgotPasswordRequest
type MphForgotPasswordRequest struct {

	// notification info
	NotificationInfo *NotificationInfo `json:"notificationInfo,omitempty"`

	// username
	Username string `json:"username,omitempty"`
}

// Validate validates this mph forgot password request
func (m *MphForgotPasswordRequest) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNotificationInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *MphForgotPasswordRequest) validateNotificationInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.NotificationInfo) { // not required
		return nil
	}

	if m.NotificationInfo != nil {
		if err := m.NotificationInfo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("notificationInfo")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *MphForgotPasswordRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *MphForgotPasswordRequest) UnmarshalBinary(b []byte) error {
	var res MphForgotPasswordRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}