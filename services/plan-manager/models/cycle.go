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

// Cycle cycle
//
// swagger:model Cycle
type Cycle struct {

	// ID
	ID string `json:"ID,omitempty" gorm:"primary_key;unique;default:md5(random()::text || clock_timestamp()::text)::uuid"`

	// resource type
	// Required: true
	ResourceType *string `json:"ResourceType"`

	// sku list
	// Required: true
	SkuList datamodels.JSONdb `json:"SkuList" gorm:"type:jsonb"`

	// state
	// Required: true
	State *string `json:"State"`
}

// Validate validates this cycle
func (m *Cycle) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResourceType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkuList(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Cycle) validateResourceType(formats strfmt.Registry) error {

	if err := validate.Required("ResourceType", "body", m.ResourceType); err != nil {
		return err
	}

	return nil
}

func (m *Cycle) validateSkuList(formats strfmt.Registry) error {

	if err := validate.Required("SkuList", "body", m.SkuList); err != nil {
		return err
	}

	return nil
}

func (m *Cycle) validateState(formats strfmt.Registry) error {

	if err := validate.Required("State", "body", m.State); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Cycle) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Cycle) UnmarshalBinary(b []byte) error {
	var res Cycle
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
