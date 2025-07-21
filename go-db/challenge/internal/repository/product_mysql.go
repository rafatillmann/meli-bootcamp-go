package repository

import (
	"database/sql"

	"app/internal/apperrors"
	"app/internal/domain"
)

// NewProductsMySQL creates new mysql repository for product entity.
func NewProductsMySQL(db *sql.DB) *ProductsMySQL {
	return &ProductsMySQL{db}
}

type ProductsMySQL struct {
	db *sql.DB
}

func (r *ProductsMySQL) FindAll() (p []domain.Product, err error) {
	rows, err := r.db.Query("SELECT `id`, `description`, `price` FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var pr domain.Product
		err := rows.Scan(&pr.Id, &pr.Description, &pr.Price)
		if err != nil {
			return nil, err
		}
		p = append(p, pr)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *ProductsMySQL) Save(p *domain.Product) (err error) {
	res, err := r.db.Exec(
		"INSERT INTO products (`description`, `price`) VALUES (?, ?)",
		(*p).Description, (*p).Price,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	(*p).Id = int(id)

	return
}

func (r *ProductsMySQL) BestSellers() ([]domain.ProductBestSeller, error) {
	rows, err := r.db.Query(`SELECT p.description, SUM(s.quantity) AS total FROM products p
                            JOIN sales s
                            ON s.product_id = p.id
                            GROUP BY p.description
                            ORDER BY total desc
                            LIMIT 5;`)
	if err != nil {
		return nil, apperrors.ErrDatabase
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, apperrors.ErrDatabase
	}

	var result []domain.ProductBestSeller
	for rows.Next() {
		var bestSeller domain.ProductBestSeller
		if err := rows.Scan(
			&bestSeller.Description,
			&bestSeller.Total,
		); err != nil {
			return nil, apperrors.ErrDatabase
		}
		result = append(result, bestSeller)
	}
	return result, nil
}
