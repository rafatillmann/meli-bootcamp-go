package repository

import (
	"database/sql"

	"app/internal/apperrors"
	"app/internal/domain"
)

// NewInvoicesMySQL creates new mysql repository for invoice entity.
func NewInvoicesMySQL(db *sql.DB) *InvoicesMySQL {
	return &InvoicesMySQL{db}
}

type InvoicesMySQL struct {
	db *sql.DB
}

func (r *InvoicesMySQL) FindAll() (i []domain.Invoice, err error) {
	rows, err := r.db.Query("SELECT `id`, `datetime`, `total`, `customer_id` FROM invoices")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var iv domain.Invoice
		err := rows.Scan(&iv.Id, &iv.Datetime, &iv.Total, &iv.CustomerId)
		if err != nil {
			return nil, err
		}
		i = append(i, iv)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *InvoicesMySQL) Save(i *domain.Invoice) (err error) {
	res, err := r.db.Exec(
		"INSERT INTO invoices (`datetime`, `total`, `customer_id`) VALUES (?, ?, ?)",
		(*i).Datetime, (*i).Total, (*i).CustomerId,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	(*i).Id = int(id)

	return
}

func (r *InvoicesMySQL) RecalculateTotal(invoiceId int) (err error) {
	updateStmt := "UPDATE invoices i SET total = (SELECT SUM(p.price * s.quantity) FROM products p JOIN sales s ON p.id = s.product_id WHERE s.invoice_id = ?) WHERE i.id = ?;"
	result, err := r.db.Exec(updateStmt, invoiceId, invoiceId)
	if err != nil {
		err = apperrors.ErrDatabase
		return
	}
	if _, err = result.RowsAffected(); err != nil {
		err = apperrors.ErrDatabase
		return
	}
	return
}
