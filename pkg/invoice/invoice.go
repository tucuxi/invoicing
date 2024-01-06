package invoice

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

type InvoiceType string

const (
	TypeEBill  InvoiceType = "ebill"
	TypePlain  InvoiceType = "plain"
	TypeQRBill InvoiceType = "qrbill"
)

var InvoiceTypes = []InvoiceType{TypeEBill, TypePlain, TypeQRBill}

type InvoiceStatus string

const (
	StatusDraft         InvoiceStatus = "draft"
	StatusOpen          InvoiceStatus = "open"
	StatusPaid          InvoiceStatus = "paid"
	StatusVoid          InvoiceStatus = "void"
	StatusUncollectible InvoiceStatus = "uncollectible"
)

type Invoice struct {
	ID                    string        `json:"id" form:"id"`
	Type                  InvoiceType   `json:"type" form:"type"`
	Recipient             string        `json:"recipient" form:"recipient"`
	Description           string        `json:"description,omitempty" form:"description"`
	Policy                string        `json:"policy,omitempty" form:"policy"`
	Currency              string        `json:"currency" form:"currency"`
	Total                 int64         `json:"total" form:"total"`
	TotalExcludingTax     int64         `json:"total_excluding_tax" form:"total_excluding_tax"`
	TotalTaxAmount        int64         `json:"total_tax_amount" form:"total_tax_amount"`
	StatementDescriptor   string        `json:"statement_descriptor,omitempty" form:"statement_descriptor"`
	PaymentOrder          string        `json:"payment_order,omitempty" form:"payment_order"`
	PeriodStart           int64         `json:"period_start,omitempty" form:"period_start"`
	PeriodEnd             int64         `json:"period_end,omitempty" form:"period_end"`
	DueDate               int64         `json:"due_date,omitempty" form:"due_date"`
	Status                InvoiceStatus `json:"status"`
	DraftedAt             int64         `json:"drafted_at,omitempty"`
	FinalizedAt           int64         `json:"finalized_at,omitempty"`
	MarkedUncollectibleAt int64         `json:"marked_uncollectible_at,omitempty"`
	PaidAt                int64         `json:"paid_at,omitempty"`
	VoidedAt              int64         `json:"voided_at,omitempty"`
	AmountPaid            int64         `json:"amount_paid,omitempty"`
	PaidOutOfBand         bool          `json:"paid_out_of_band,omitempty"`
	Source                string        `json:"source"`
}

func NewInvoiceID() string {
	return fmt.Sprintf("inv_%s", ksuid.New())
}
