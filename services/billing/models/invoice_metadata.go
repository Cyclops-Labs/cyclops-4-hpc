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

// InvoiceMetadata invoice metadata
//
// swagger:model InvoiceMetadata
type InvoiceMetadata struct {

	// amount invoiced
	AmountInvoiced float64 `json:"AmountInvoiced,omitempty"`

	// currency
	// Enum: [CHF EUR USD]
	Currency *string `json:"Currency,omitempty"`

	// ID
	// Format: uuid
	ID strfmt.UUID `json:"ID,omitempty"`

	// organization ID
	OrganizationID string `json:"OrganizationID,omitempty"`

	// organization name
	OrganizationName string `json:"OrganizationName,omitempty"`

	// payment deadline
	// Format: date
	PaymentDeadline strfmt.Date `json:"PaymentDeadline,omitempty"`

	// payment status
	// Enum: [CANCELLED PAID UNPAID]
	PaymentStatus *string `json:"PaymentStatus,omitempty"`

	// period end date
	// Format: date
	PeriodEndDate strfmt.Date `json:"PeriodEndDate,omitempty"`

	// period start date
	// Format: date
	PeriodStartDate strfmt.Date `json:"PeriodStartDate,omitempty"`

	// status
	// Enum: [ERROR FINISHED NOT_PROCESSED PROCESSING]
	Status *string `json:"Status,omitempty"`
}

// Validate validates this invoice metadata
func (m *InvoiceMetadata) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCurrency(formats); err != nil {
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

var invoiceMetadataTypeCurrencyPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CHF","EUR","USD"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceMetadataTypeCurrencyPropEnum = append(invoiceMetadataTypeCurrencyPropEnum, v)
	}
}

const (

	// InvoiceMetadataCurrencyCHF captures enum value "CHF"
	InvoiceMetadataCurrencyCHF string = "CHF"

	// InvoiceMetadataCurrencyEUR captures enum value "EUR"
	InvoiceMetadataCurrencyEUR string = "EUR"

	// InvoiceMetadataCurrencyUSD captures enum value "USD"
	InvoiceMetadataCurrencyUSD string = "USD"
)

// prop value enum
func (m *InvoiceMetadata) validateCurrencyEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceMetadataTypeCurrencyPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceMetadata) validateCurrency(formats strfmt.Registry) error {

	if swag.IsZero(m.Currency) { // not required
		return nil
	}

	// value enum
	if err := m.validateCurrencyEnum("Currency", "body", *m.Currency); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceMetadata) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("ID", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceMetadata) validatePaymentDeadline(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentDeadline) { // not required
		return nil
	}

	if err := validate.FormatOf("PaymentDeadline", "body", "date", m.PaymentDeadline.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceMetadataTypePaymentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["CANCELLED","PAID","UNPAID"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceMetadataTypePaymentStatusPropEnum = append(invoiceMetadataTypePaymentStatusPropEnum, v)
	}
}

const (

	// InvoiceMetadataPaymentStatusCANCELLED captures enum value "CANCELLED"
	InvoiceMetadataPaymentStatusCANCELLED string = "CANCELLED"

	// InvoiceMetadataPaymentStatusPAID captures enum value "PAID"
	InvoiceMetadataPaymentStatusPAID string = "PAID"

	// InvoiceMetadataPaymentStatusUNPAID captures enum value "UNPAID"
	InvoiceMetadataPaymentStatusUNPAID string = "UNPAID"
)

// prop value enum
func (m *InvoiceMetadata) validatePaymentStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceMetadataTypePaymentStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceMetadata) validatePaymentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.PaymentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validatePaymentStatusEnum("PaymentStatus", "body", *m.PaymentStatus); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceMetadata) validatePeriodEndDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PeriodEndDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PeriodEndDate", "body", "date", m.PeriodEndDate.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *InvoiceMetadata) validatePeriodStartDate(formats strfmt.Registry) error {

	if swag.IsZero(m.PeriodStartDate) { // not required
		return nil
	}

	if err := validate.FormatOf("PeriodStartDate", "body", "date", m.PeriodStartDate.String(), formats); err != nil {
		return err
	}

	return nil
}

var invoiceMetadataTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ERROR","FINISHED","NOT_PROCESSED","PROCESSING"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		invoiceMetadataTypeStatusPropEnum = append(invoiceMetadataTypeStatusPropEnum, v)
	}
}

const (

	// InvoiceMetadataStatusERROR captures enum value "ERROR"
	InvoiceMetadataStatusERROR string = "ERROR"

	// InvoiceMetadataStatusFINISHED captures enum value "FINISHED"
	InvoiceMetadataStatusFINISHED string = "FINISHED"

	// InvoiceMetadataStatusNOTPROCESSED captures enum value "NOT_PROCESSED"
	InvoiceMetadataStatusNOTPROCESSED string = "NOT_PROCESSED"

	// InvoiceMetadataStatusPROCESSING captures enum value "PROCESSING"
	InvoiceMetadataStatusPROCESSING string = "PROCESSING"
)

// prop value enum
func (m *InvoiceMetadata) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, invoiceMetadataTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *InvoiceMetadata) validateStatus(formats strfmt.Registry) error {

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
func (m *InvoiceMetadata) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InvoiceMetadata) UnmarshalBinary(b []byte) error {
	var res InvoiceMetadata
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
