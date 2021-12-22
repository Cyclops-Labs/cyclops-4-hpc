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
	"gitlab.com/cyclops-utilities/datamodels"
)

// Invoice invoice
//
// swagger:model Invoice
type Invoice struct {

	// amount invoiced
	AmountInvoiced float64 `json:"AmountInvoiced,omitempty" gorm:"type:numeric(23,13)"`

	// bill run ID
	// Format: uuid
	BillRunID strfmt.UUID `json:"BillRunID,omitempty"`

	// billing contact
	BillingContact string `json:"BillingContact,omitempty"`

	// currency
	// Enum: [CHF EUR USD]
	Currency *string `json:"Currency,omitempty"`

	// generation timestamp
	// Format: date-time
	GenerationTimestamp strfmt.DateTime `json:"GenerationTimestamp,omitempty" gorm:"type:timestamptz"`

	// ID
	// Format: uuid
	ID strfmt.UUID `json:"ID,omitempty" gorm:"type:uuid;primary_key;unique;default:md5(random()::text || clock_timestamp()::text)::uuid"`

	// items
	Items datamodels.JSONdb `json:"Items,omitempty" gorm:"type:jsonb"`

	// organization ID
	OrganizationID string `json:"OrganizationID,omitempty"`

	// organization name
	OrganizationName string `json:"OrganizationName,omitempty"`

	// organization type
	OrganizationType string `json:"OrganizationType,omitempty"`

	// payment deadline
	// Format: date
	PaymentDeadline strfmt.Date `json:"PaymentDeadline,omitempty" gorm:"type:date"`

	// payment status
	// Enum: [CANCELLED PAID UNPAID]
	PaymentStatus *string `json:"PaymentStatus,omitempty"`

	// period end date
	// Format: date
	PeriodEndDate strfmt.Date `json:"PeriodEndDate,omitempty" gorm:"type:date"`

	// period start date
	// Format: date
	PeriodStartDate strfmt.Date `json:"PeriodStartDate,omitempty" gorm:"type:date"`

	// status
	// Enum: [ERROR FINISHED NOT_PROCESSED PROCESSING]
	Status *string `json:"Status,omitempty" gorm:"default:NOT_PROCESSED"`
}

// Validate validates this invoice
func (m *Invoice) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateBillRunID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCurrency(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateGenerationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentDeadline(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePaymentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodEndDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePeriodStartDate(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Invoice) validateBillRunID(formats strfmt.Registry) error {

	if swag.IsZero(m.BillRunID) { // not required
		return nil
	}

	if err := validate.FormatOf("BillRunID", "body", "uuid", m.BillRunID.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceTypeCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CHF","EUR","USD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceTypeCurrencyPropEnum = append(invoiceTypeCurrencyPropEnum, v)
	}
}

const (

	// InvoiceCurrencyCHF captures enum value "CHF"
	InvoiceCurrencyCHF string = "CHF"

	// InvoiceCurrencyEUR captures enum value "EUR"
	InvoiceCurrencyEUR string = "EUR"

	// InvoiceCurrencyUSD captures enum value "USD"
	InvoiceCurrencyUSD string = "USD"
)

// prop value enum
func (m *Invoice) validateCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Invoice) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("Currency", "body", *m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateGenerationTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.GenerationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("GenerationTimestamp", "body", "date-time", m.GenerationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("ID", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validatePaymentDeadline(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentDeadline) { // not required
		return nil
	}

	if err := validate.FormatOf("PaymentDeadline", "body", "date", m.PaymentDeadline.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceTypePaymentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CANCELLED","PAID","UNPAID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceTypePaymentStatusPropEnum = append(invoiceTypePaymentStatusPropEnum, v)
	}
}

const (

	// InvoicePaymentStatusCANCELLED captures enum value "CANCELLED"
	InvoicePaymentStatusCANCELLED string = "CANCELLED"

	// InvoicePaymentStatusPAID captures enum value "PAID"
	InvoicePaymentStatusPAID string = "PAID"

	// InvoicePaymentStatusUNPAID captures enum value "UNPAID"
	InvoicePaymentStatusUNPAID string = "UNPAID"
)

// prop value enum
func (m *Invoice) validatePaymentStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceTypePaymentStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Invoice) validatePaymentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentStatusEnum("PaymentStatus", "body", *m.PaymentStatus); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validatePeriodEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PeriodEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PeriodEndDate", "body", "date", m.PeriodEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Invoice) validatePeriodStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PeriodStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PeriodStartDate", "body", "date", m.PeriodStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ERROR","FINISHED","NOT_PROCESSED","PROCESSING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceTypeStatusPropEnum = append(invoiceTypeStatusPropEnum, v)
	}
}

const (

	// InvoiceStatusERROR captures enum value "ERROR"
	InvoiceStatusERROR string = "ERROR"

	// InvoiceStatusFINISHED captures enum value "FINISHED"
	InvoiceStatusFINISHED string = "FINISHED"

	// InvoiceStatusNOTPROCESSED captures enum value "NOT_PROCESSED"
	InvoiceStatusNOTPROCESSED string = "NOT_PROCESSED"

	// InvoiceStatusPROCESSING captures enum value "PROCESSING"
	InvoiceStatusPROCESSING string = "PROCESSING"
)

// prop value enum
func (m *Invoice) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Invoice) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("Status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Invoice) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Invoice) UnmarshalBinary(b []byte) error {
	var res Invoice
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}