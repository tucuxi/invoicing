package persistence

import (
	"errors"
	"slices"
	"sync"

	"github.com/tucuxi/invoicing/pkg/invoice"
)

type InvoiceRepository struct {
	mu       sync.Mutex
	invoices []invoice.Invoice
}

var (
	ErrorInvoiceNotFound    = errors.New("invoice not found")
	ErrorDeletionNotAllowed = errors.New("deletion not allowed")
)

func NewInvoiceRepository() *InvoiceRepository {
	return &InvoiceRepository{}
}

func (r *InvoiceRepository) CreateInvoice(i *invoice.Invoice) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.invoices = append(r.invoices, *i)
	return nil
}

func (r *InvoiceRepository) UpdateInvoice(i *invoice.Invoice) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := slices.IndexFunc(r.invoices, func(j invoice.Invoice) bool { return j.ID == i.ID })
	if k == -1 {
		return ErrorInvoiceNotFound
	}
	r.invoices[k] = *i
	return nil
}

func (r *InvoiceRepository) FindInvoice(id string) (*invoice.Invoice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := slices.IndexFunc(r.invoices, func(j invoice.Invoice) bool { return j.ID == id })
	if k == -1 {
		return nil, ErrorInvoiceNotFound
	}
	i := r.invoices[k]
	return &i, nil
}

func (r *InvoiceRepository) DeleteDraftInvoice(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := slices.IndexFunc(r.invoices, func(j invoice.Invoice) bool { return j.ID == id })
	if k == -1 {
		return ErrorInvoiceNotFound
	}
	if r.invoices[k].Status != invoice.StatusDraft {
		return ErrorDeletionNotAllowed
	}
	r.invoices = append(r.invoices[:k], r.invoices[k+1:]...)
	return nil
}
