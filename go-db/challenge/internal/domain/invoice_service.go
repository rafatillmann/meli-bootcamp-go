package domain

// ServiceInvoice is the interface that wraps the basic methods that an invoice service should implement.
type ServiceInvoice interface {
	FindAll() (i []Invoice, err error)
	Save(i *Invoice) (err error)
	RecalculateTotal(invoiceId int) (err error)
}
