package repository

import (
	"database/sql"

	"app/internal/domain"
)

// NewSalesMySQL creates new mysql repository for sale entity.
func NewSalesMySQL(db *sql.DB) *SalesMySQL {
	return &SalesMySQL{db}
}

type SalesMySQL struct {
	db *sql.DB
}

func (r *SalesMySQL) FindAll() (s []domain.Sale, err error) {
	rows, err := r.db.Query("SELECT `id`, `quantity`, `product_id`, `invoice_id` FROM sales")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var sa domain.Sale
		err := rows.Scan(&sa.Id, &sa.Quantity, &sa.ProductId, &sa.InvoiceId)
		if err != nil {
			return nil, err
		}
		s = append(s, sa)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *SalesMySQL) Save(s *domain.Sale) (err error) {
	res, err := r.db.Exec(
		"INSERT INTO sales (`quantity`, `product_id`, `invoice_id`) VALUES (?, ?, ?)",
		(*s).Quantity, (*s).ProductId, (*s).InvoiceId,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	(*s).Id = int(id)

	return
}
