package repository

import (
	"app/internal/apperrors"
	"app/internal/domain"
	"database/sql"
	"errors"
)

func NewRepositoryProductSql(db *sql.DB) (r *RepositoryProductSql) {
	r = &RepositoryProductSql{db}
	return
}

type RepositoryProductSql struct {
	db *sql.DB
}

func (r *RepositoryProductSql) Delete(id int) (err error) {
	if _, err := r.db.Exec("delete from users u where u.id = ?", id); err != nil {
		return apperrors.ErrRepositoryDeleteData
	}

	return nil
}

func (r *RepositoryProductSql) FindById(id int) (p domain.Product, err error) {
	row := r.db.QueryRow("select p.id, p.name, p.quantity, p.code_value, p.is_published, p.expiration, p.price from products p where p.id = ?", id)

	if err := row.Err(); err != nil {
		return domain.Product{}, apperrors.ErrRepositoryFetchData
	}

	if err := row.Scan(&p.Id, &p.Name, &p.Quantity, &p.CodeValue, &p.IsPublished, &p.Expiration, &p.Price); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return domain.Product{}, apperrors.ErrRepositoryProductNotFound
		} else {
			return domain.Product{}, apperrors.ErrRepositoryFetchData
		}
	}
	return p, nil
}

func (r *RepositoryProductSql) Save(p *domain.Product) (err error) {
	result, err := r.db.Exec("insert into products (name, quantity, code_value, is_published, expiration, price) values (?, ?, ?, ?, ?, ?)",
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price)

	if err != nil {
		return apperrors.ErrRepositorySaveData
	}

	lastId, err := result.LastInsertId()
	if err != nil {
		return apperrors.ErrRepositoryGetLastId
	}

	p.Id = int(lastId)
	return nil
}

func (r *RepositoryProductSql) Update(p *domain.Product) (err error) {
	if _, err := r.db.Exec("update products set name = ?, quantity = ?, code_value = ?, is_published = ?, expiration = ?, price = ? where id = ?",
		p.Name, p.Quantity, p.CodeValue, p.IsPublished, p.Expiration, p.Price, p.Id); err != nil {
		return apperrors.ErrRepositoryUpdateData
	}

	return nil
}
