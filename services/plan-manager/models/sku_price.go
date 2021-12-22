// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SkuPrice sku price
//
// swagger:model SkuPrice
type SkuPrice struct {

	// accounting mode
	// Enum: [CREDIT CASH BOTH NONE]
	AccountingMode *string `json:"AccountingMode,omitempty" gorm:"index;default:CREDIT"`

	// discount
	Discount float64 `json:"Discount,omitempty" gorm:"type:numeric(23,13);default:0.0"`

	// ID
	ID string `json:"ID,omitempty" gorm:"primary_key;unique;default:md5(random()::text || clock_timestamp()::text)::uuid"`

	// plan ID
	// Required: true
	PlanID *string `json:"PlanID"`

	// sku ID
	// Required: true
	SkuID *string `json:"SkuID"`

	// sku name
	SkuName string `json:"SkuName,omitempty"`

	// unit
	// Required: true
	Unit *string `json:"Unit"`

	// unit credit price
	UnitCreditPrice float64 `json:"UnitCreditPrice,omitempty" gorm:"type:numeric(23,13)"`

	// unit price
	// Required: true
	UnitPrice *float64 `json:"UnitPrice" gorm:"type:numeric(23,13)"`
}

// Validate validates this sku price
func (m *SkuPrice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAccountingMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePlanID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSkuID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUnitPrice(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var skuPriceTypeAccountingModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CREDIT","CASH","BOTH","NONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		skuPriceTypeAccountingModePropEnum = append(skuPriceTypeAccountingModePropEnum, v)
	}
}

const (

	// SkuPriceAccountingModeCREDIT captures enum value "CREDIT"
	SkuPriceAccountingModeCREDIT string = "CREDIT"

	// SkuPriceAccountingModeCASH captures enum value "CASH"
	SkuPriceAccountingModeCASH string = "CASH"

	// SkuPriceAccountingModeBOTH captures enum value "BOTH"
	SkuPriceAccountingModeBOTH string = "BOTH"

	// SkuPriceAccountingModeNONE captures enum value "NONE"
	SkuPriceAccountingModeNONE string = "NONE"
)

// prop value enum
func (m *SkuPrice) validateAccountingModeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, skuPriceTypeAccountingModePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *SkuPrice) validateAccountingMode(formats strfmt.Registry) error {

	if swag.IsZero(m.AccountingMode) { // not required
		return nil
	}

	// value enum
	if err := m.validateAccountingModeEnum("AccountingMode", "body", *m.AccountingMode); err != nil {
		return err
	}

	return nil
}

func (m *SkuPrice) validatePlanID(formats strfmt.Registry) error {

	if err := validate.Required("PlanID", "body", m.PlanID); err != nil {
		return err
	}

	return nil
}

func (m *SkuPrice) validateSkuID(formats strfmt.Registry) error {

	if err := validate.Required("SkuID", "body", m.SkuID); err != nil {
		return err
	}

	return nil
}

func (m *SkuPrice) validateUnit(formats strfmt.Registry) error {

	if err := validate.Required("Unit", "body", m.Unit); err != nil {
		return err
	}

	return nil
}

func (m *SkuPrice) validateUnitPrice(formats strfmt.Registry) error {

	if err := validate.Required("UnitPrice", "body", m.UnitPrice); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SkuPrice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SkuPrice) UnmarshalBinary(b []byte) error {
	var res SkuPrice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
