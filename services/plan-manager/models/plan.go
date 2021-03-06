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

// Plan plan
//
// swagger:model Plan
type Plan struct {

	// discount
	Discount float64 `json:"Discount,omitempty" gorm:"type:numeric(23,13);default:0.0"`

	// ID
	ID string `json:"ID,omitempty" gorm:"primary_key;unique;default:md5(random()::text || clock_timestamp()::text)::uuid"`

	// name
	// Required: true
	Name *string `json:"Name"`

	// offered end date
	// Required: true
	// Format: date
	OfferedEndDate *strfmt.Date `json:"OfferedEndDate" gorm:"type:date"`

	// offered start date
	// Required: true
	// Format: date
	OfferedStartDate *strfmt.Date `json:"OfferedStartDate" gorm:"type:date"`

	// sku prices
	SkuPrices []*SkuPrice `json:"SkuPrices" gorm:"-"`
}

// Validate validates this plan
func (m *Plan) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferedEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOfferedStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkuPrices(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Plan) validateName(formats strfmt.Registry) error {

	if err := validate.Required("Name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateOfferedEndDate(formats strfmt.Registry) error {

	if err := validate.Required("OfferedEndDate", "body", m.OfferedEndDate); err != nil {
		return err
	}

	if err := validate.FormatOf("OfferedEndDate", "body", "date", m.OfferedEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateOfferedStartDate(formats strfmt.Registry) error {

	if err := validate.Required("OfferedStartDate", "body", m.OfferedStartDate); err != nil {
		return err
	}

	if err := validate.FormatOf("OfferedStartDate", "body", "date", m.OfferedStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Plan) validateSkuPrices(formats strfmt.Registry) error {

	if swag.IsZero(m.SkuPrices) { // not required
		return nil
	}

	for i := 0; i < len(m.SkuPrices); i++ {
		if swag.IsZero(m.SkuPrices[i]) { // not required
			continue
		}

		if m.SkuPrices[i] != nil {
			if err := m.SkuPrices[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("SkuPrices" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *Plan) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Plan) UnmarshalBinary(b []byte) error {
	var res Plan
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
