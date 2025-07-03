package domain

type RepositoryProduct interface {
	FindById(id int) (p Product, err error)
	Save(p *Product) (err error)
	Update(p *Product) (err error)
	Delete(id int) (err error)
}
