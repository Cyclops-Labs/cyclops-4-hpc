// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
	"github.com/lib/pq"
)

// BillRunReport bill run report
//
// swagger:model BillRunReport
type BillRunReport struct {

	// amount invoiced
	AmountInvoiced float64 `json:"AmountInvoiced,omitempty"`

	// creation timestamp
	// Format: date-time
	CreationTimestamp strfmt.DateTime `json:"CreationTimestamp,omitempty"`

	// ID
	// Format: uuid
	ID strfmt.UUID `json:"ID,omitempty"`

	// invoices count
	InvoicesCount int64 `json:"InvoicesCount,omitempty"`

	// invoices error count
	InvoicesErrorCount int64 `json:"InvoicesErrorCount,omitempty"`

	// invoices error list
	InvoicesErrorList pq.StringArray `json:"InvoicesErrorList,omitempty" gorm:"type:text[]"`

	// invoices list
	InvoicesList []*InvoiceMetadata `json:"InvoicesList" gorm:"-"`

	// invoices processed count
	InvoicesProcessedCount int64 `json:"InvoicesProcessedCount,omitempty"`

	// status
	// Enum: [ERROR FINISHED PROCESSING QUEUED]
	Status *string `json:"Status,omitempty"`
}

// Validate validates this bill run report
func (m *BillRunReport) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreationTimestamp(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInvoicesList(formats); err != nil {
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

func (m *BillRunReport) validateCreationTimestamp(formats strfmt.Registry) error {

	if swag.IsZero(m.CreationTimestamp) { // not required
		return nil
	}

	if err := validate.FormatOf("CreationTimestamp", "body", "date-time", m.CreationTimestamp.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillRunReport) validateID(formats strfmt.Registry) error {

	if swag.IsZero(m.ID) { // not required
		return nil
	}

	if err := validate.FormatOf("ID", "body", "uuid", m.ID.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *BillRunReport) validateInvoicesList(formats strfmt.Registry) error {

	if swag.IsZero(m.InvoicesList) { // not required
		return nil
	}

	for i := 0; i < len(m.InvoicesList); i++ {
		if swag.IsZero(m.InvoicesList[i]) { // not required
			continue
		}

		if m.InvoicesList[i] != nil {
			if err := m.InvoicesList[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("InvoicesList" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

var billRunReportTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["ERROR","FINISHED","PROCESSING","QUEUED"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		billRunReportTypeStatusPropEnum = append(billRunReportTypeStatusPropEnum, v)
	}
}

const (

	// BillRunReportStatusERROR captures enum value "ERROR"
	BillRunReportStatusERROR string = "ERROR"

	// BillRunReportStatusFINISHED captures enum value "FINISHED"
	BillRunReportStatusFINISHED string = "FINISHED"

	// BillRunReportStatusPROCESSING captures enum value "PROCESSING"
	BillRunReportStatusPROCESSING string = "PROCESSING"

	// BillRunReportStatusQUEUED captures enum value "QUEUED"
	BillRunReportStatusQUEUED string = "QUEUED"
)

// prop value enum
func (m *BillRunReport) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, billRunReportTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *BillRunReport) validateStatus(formats strfmt.Registry) error {

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
func (m *BillRunReport) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *BillRunReport) UnmarshalBinary(b []byte) error {
	var res BillRunReport
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}