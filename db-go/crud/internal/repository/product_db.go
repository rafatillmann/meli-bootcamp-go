package repository

import (
	"app/internal/domain"
	"database/sql"
)

func NewRepositoryDb(db *sql.DB) (r *RepositoryProductDb) {
	r = &RepositoryProductDb{db}
	return
}

type RepositoryProductDb struct {
	db *sql.DB
}

// Delete implements internal.RepositoryProduct.
func (r *RepositoryProductDb) Delete(id int) (err error) {
	panic("unimplemented")
}

// FindById implements internal.RepositoryProduct.
func (r *RepositoryProductDb) FindById(id int) (p domain.Product, err error) {
	panic("unimplemented")
}

// Save implements internal.RepositoryProduct.
func (r *RepositoryProductDb) Save(p *domain.Product) (err error) {
	panic("unimplemented")
}

// Update implements internal.RepositoryProduct.
func (r *RepositoryProductDb) Update(p *domain.Product) (err error) {
	panic("unimplemented")
}

// UpdateOrSave implements internal.RepositoryProduct.
func (r *RepositoryProductDb) UpdateOrSave(p *domain.Product) (err error) {
	panic("unimplemented")
}
