// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// Usage usage
//
// swagger:model Usage
type Usage struct {

	// account Id
	AccountID string `json:"AccountId,omitempty"`

	// time from
	TimeFrom int64 `json:"TimeFrom,omitempty"`

	// time to
	TimeTo int64 `json:"TimeTo,omitempty"`

	// usage
	Usage []*Use `json:"Usage"`
}

// Validate validates this usage
func (m *Usage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateUsage(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Usage) validateUsage(formats strfmt.Registry) error {

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
func (m *Usage) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Usage) UnmarshalBinary(b []byte) error {
	var res Usage
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
