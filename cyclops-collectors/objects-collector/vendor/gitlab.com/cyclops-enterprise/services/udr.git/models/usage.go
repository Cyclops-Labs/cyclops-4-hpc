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

// Usage usage
//
// swagger:model Usage
type Usage struct {

	// account
	Account string `json:"Account,omitempty" gorm:"index"`

	// metadata
	Metadata datamodels.JSONdb `json:"Metadata,omitempty" gorm:"type:jsonb"`

	// resource ID
	ResourceID string `json:"ResourceID,omitempty"`

	// resource name
	ResourceName string `json:"ResourceName,omitempty"`

	// resource type
	ResourceType string `json:"ResourceType,omitempty"`

	// time
	Time int64 `json:"Time,omitempty"`

	// timedate
	// Format: datetime
	Timedate strfmt.DateTime `json:"Timedate,omitempty" gorm:"index;type:timestamptz"`

	// unit
	Unit string `json:"Unit,omitempty"`

	// usage
	Usage float64 `json:"Usage,omitempty" gorm:"type:numeric(23,13);default:0.0"`
}

// Validate validates this usage
func (m *Usage) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTimedate(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Usage) validateTimedate(formats strfmt.Registry) error {

	if swag.IsZero(m.Timedate) { // not required
		return nil
	}

	if err := validate.FormatOf("Timedate", "body", "datetime", m.Timedate.String(), formats); err != nil {
		return err
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
