package domain

// ServiceProduct is the interface that wraps the basic Product methods.
type ServiceProduct interface {
	FindAll() (p []Product, err error)
	Save(p *Product) (err error)
	BestSellers() ([]ProductBestSeller, error)
}
