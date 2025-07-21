package domain

// ServiceSale is the interface that wraps the basic ServiceSale methods.
type ServiceSale interface {
	FindAll() (s []Sale, err error)
	Save(s *Sale) (err error)
}