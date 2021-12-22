// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// UReport u report
//
// swagger:model UReport
type UReport struct {

	// account Id
	AccountID string `json:"AccountId,omitempty" gorm:"index"`

	// time from
	// Format: datetime
	TimeFrom strfmt.DateTime `json:"TimeFrom,omitempty" gorm:"index;type:timestamptz"`

	// time to
	// Format: datetime
	TimeTo strfmt.DateTime `json:"TimeTo,omitempty" gorm:"index;type:timestamptz"`

	// usage
	Usage []*UDRReport `json:"Usage" gorm:"-"`
}

// Validate validates this u report
func (m *UReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeTo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UReport) validateTimeFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("TimeFrom", "body", "datetime", m.TimeFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UReport) validateTimeTo(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeTo) { // not required
		return nil
	}

	if err := validate.FormatOf("TimeTo", "body", "datetime", m.TimeTo.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UReport) validateUsage(formats strfmt.Registry) error {

	if swag.IsZero(m.Usage) { // not required
		return nil
	}

	for i := 0; i < len(m.Usage); i++ {
		if swag.IsZero(m.Usage[i]) { // not required
			continue
		}

		if m.Usage[i] != nil {
			if err := m.Usage[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("Usage" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *UReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UReport) UnmarshalBinary(b []byte) error {
	var res UReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}