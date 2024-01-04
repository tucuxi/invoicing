package invoice

import (
	"fmt"

	"github.com/segmentio/ksuid"
)

type InvoiceType string

const (
	EBill  = "ebill"
	Plain  = "plain"
	QRBill = "qrbill"
)

var InvoiceTypes = []InvoiceType{EBill, Plain, QRBill}

type InvoiceStatus string

const (
	Draft         = "draft"
	Open          = "open"
	Paid          = "paid"
	Void          = "void"
	Uncollectible = "uncollectible"
)

type Invoice struct {
	ID                    string        `json:"id" query:"id"`
	Type                  InvoiceType   `json:"type" query:"type"`
	Recipient             string        `json:"recipient" query:"recipient"`
	Description           string        `json:"description,omitempty" query:"description"`
	Policy                string        `json:"policy,omitempty" query:"policy"`
	Currency              string        `json:"currency" query:"currency"`
	Total                 int64         `json:"total" query:"total"`
	TotalExcludingTax     int64         `json:"total_excluding_tax" query:"total_excluding_tax"`
	TotalTaxAmount        int64         `json:"total_tax_amount" query:"total_tax_amount"`
	StatementDescriptor   string        `json:"statement_descriptor,omitempty" query:"statement_descriptor"`
	PaymentOrder          string        `json:"payment_order,omitempty" query:"payment_order"`
	PeriodStart           int64         `json:"period_start,omitempty" query:"period_start"`
	PeriodEnd             int64         `json:"period_end,omitempty" query:"period_end"`
	DueDate               int64         `json:"due_date,omitempty" query:"due_date"`
	Status                InvoiceStatus `json:"status"`
	DraftedAt             int64         `json:"drafted_at,omitempty"`
	FinalizedAt           int64         `json:"finalized_at,omitempty"`
	MarkedUncollectibleAt int64         `json:"marked_uncollectible_at,omitempty"`
	PaidAt                int64         `json:"paid_at,omitempty"`
	VoidedAt              int64         `json:"voided_at,omitempty"`
	Source                string        `json:"source"`
}

func NewInvoiceID() string {
	return fmt.Sprintf("inv_%s", ksuid.New())
}
