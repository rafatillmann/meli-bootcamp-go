package domain

// RepositoryProduct is the interface that wraps the basic methods that a product repository must have.
type RepositoryProduct interface {
	FindAll() (p []Product, err error)
	Save(p *Product) (err error)
	BestSellers() ([]ProductBestSeller, error)
}
