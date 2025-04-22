package persistence

import (
	"slices"
	"sync"

	"github.com/tucuxi/invoicing/pkg/invoice"
)

type TransientRepository struct {
	mu       sync.Mutex
	invoices []*invoice.Invoice
}

func NewTransientRepository() *TransientRepository {
	return &TransientRepository{}
}

func (r *TransientRepository) CreateInvoice(i *invoice.Invoice) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	cp := *i
	r.invoices = append(r.invoices, &cp)
	return nil
}

func (r *TransientRepository) UpdateInvoice(i *invoice.Invoice) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := slices.IndexFunc(r.invoices, func(j *invoice.Invoice) bool { return j.ID == i.ID })
	if k == -1 {
		return invoice.ErrorInvoiceNotFound
	}
	cp := *i
	r.invoices[k] = &cp
	return nil
}

func (r *TransientRepository) FindInvoice(id string) (*invoice.Invoice, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := slices.IndexFunc(r.invoices, func(j *invoice.Invoice) bool { return j.ID == id })
	if k == -1 {
		return nil, invoice.ErrorInvoiceNotFound
	}
	cp := *r.invoices[k]
	return &cp, nil
}

func (r *TransientRepository) DeleteDraftInvoice(id string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	k := slices.IndexFunc(r.invoices, func(j *invoice.Invoice) bool { return j.ID == id })
	if k == -1 {
		return invoice.ErrorInvoiceNotFound
	}
	if r.invoices[k].Status != invoice.StatusDraft {
		return invoice.ErrorDeletionNotAllowed
	}
	r.invoices = append(r.invoices[:k], r.invoices[k+1:]...)
	return nil
}
