package domain

// RepositoryInvoice is the interface that wraps the basic methods that an invoice repository should implement.
type RepositoryInvoice interface {
	FindAll() (i []Invoice, err error)
	Save(i *Invoice) (err error)
	RecalculateTotal(invoiceId int) (err error)
}
