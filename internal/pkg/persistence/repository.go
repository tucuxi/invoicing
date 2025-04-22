package persistence

import (
	"github.com/tucuxi/invoicing/pkg/invoice"
)

type InvoiceRepository interface {
	CreateInvoice(i *invoice.Invoice) error
	UpdateInvoice(i *invoice.Invoice) error
	FindInvoice(id string) (*invoice.Invoice, error)
	DeleteDraftInvoice(id string) error
}
