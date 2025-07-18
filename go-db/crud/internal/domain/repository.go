package domain

type RepositoryProduct interface {
	FindById(id int) (p Product, err error)
	Save(p *Product) (err error)
	Update(p *Product) (err error)
	Delete(id int) (err error)
}

type RepositoryWarehouse interface {
	FindById(id int) (p Warehouse, err error)
	Save(p *Warehouse) (err error)
	Update(p *Warehouse) (err error)
	Delete(id int) (err error)
}