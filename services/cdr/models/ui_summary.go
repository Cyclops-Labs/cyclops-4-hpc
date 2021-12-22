// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"gitlab.com/cyclops-utilities/datamodels"
)

// UISummary UI summary
//
// swagger:model UISummary
type UISummary struct {

	// account Id
	AccountID string `json:"AccountId,omitempty"`

	// time from
	// Format: datetime
	TimeFrom strfmt.DateTime `json:"TimeFrom,omitempty"`

	// time to
	// Format: datetime
	TimeTo strfmt.DateTime `json:"TimeTo,omitempty"`

	// usage breakup
	UsageBreakup datamodels.JSONdb `json:"UsageBreakup,omitempty"`
}

// Validate validates this UI summary
func (m *UISummary) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimeFrom(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTimeTo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *UISummary) validateTimeFrom(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeFrom) { // not required
		return nil
	}

	if err := validate.FormatOf("TimeFrom", "body", "datetime", m.TimeFrom.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UISummary) validateTimeTo(formats strfmt.Registry) error {

	if swag.IsZero(m.TimeTo) { // not required
		return nil
	}

	if err := validate.FormatOf("TimeTo", "body", "datetime", m.TimeTo.String(), formats); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *UISummary) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UISummary) UnmarshalBinary(b []byte) error {
	var res UISummary
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}