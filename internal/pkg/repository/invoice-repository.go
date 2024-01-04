package repository

import (
	"errors"
	"slices"
	"sync"

	"github.com/tucuxi/invoicing/pkg/invoice"
)

type InvoiceRepository struct {
	mu       sync.Mutex
	invoices []*invoice.Invoice
}

var (
	ErrorInvoiceNotFound = errors.New("invoice not found")
)

func NewInvoiceRepository() *InvoiceRepository {
	return &InvoiceRepository{}
}

func (r *InvoiceRepository) CreateInvoice(i *invoice.Invoice) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.invoices = append(r.invoices, i)
	return nil
}

func (r *InvoiceRepository) FindInvoice(id string) (*invoice.Invoice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := slices.IndexFunc(r.invoices, func(i *invoice.Invoice) bool { return i.ID == id })
	if k == -1 {
		return nil, ErrorInvoiceNotFound
	}
	return r.invoices[k], nil
}