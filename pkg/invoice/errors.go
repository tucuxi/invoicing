package invoice

import "errors"

var (
	ErrorInvoiceNotFound    = errors.New("invoice not found")
	ErrorDeletionNotAllowed = errors.New("deletion not allowed")
)
