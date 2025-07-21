package repository

import (
	"database/sql"

	"app/internal/apperrors"
	"app/internal/domain"
)

// NewCustomersMySQL creates new mysql repository for customer entity.
func NewCustomersMySQL(db *sql.DB) *CustomersMySQL {
	return &CustomersMySQL{db}
}

type CustomersMySQL struct {
	db *sql.DB
}

func (r *CustomersMySQL) FindAll() (c []domain.Customer, err error) {
	rows, err := r.db.Query("SELECT `id`, `first_name`, `last_name`, `condition` FROM customers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var cs domain.Customer
		err := rows.Scan(&cs.Id, &cs.FirstName, &cs.LastName, &cs.Condition)
		if err != nil {
			return nil, err
		}
		c = append(c, cs)
	}
	err = rows.Err()
	if err != nil {
		return
	}

	return
}

func (r *CustomersMySQL) Save(c *domain.Customer) (err error) {
	res, err := r.db.Exec(
		"INSERT INTO customers (`first_name`, `last_name`, `condition`) VALUES (?, ?, ?)",
		(*c).FirstName, (*c).LastName, (*c).Condition,
	)
	if err != nil {
		return err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return err
	}

	(*c).Id = int(id)

	return
}

func (r *CustomersMySQL) TotalByCondition() ([]domain.CustomerCondition, error) {
	rows, err := r.db.Query(`SELECT c.condition, TRUNCATE(SUM(i.total), 2) AS total from customers c 
                            JOIN invoices i
                            ON i.customer_id = c.id
                            GROUP BY c.condition`)

	if err != nil {
		return nil, apperrors.ErrDatabase
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, apperrors.ErrDatabase
	}

	var result []domain.CustomerCondition
	for rows.Next() {
		var customerCondition domain.CustomerCondition

		if err := rows.Scan(
			&customerCondition.Condition,
			&customerCondition.Total,
		); err != nil {
			return nil, apperrors.ErrDatabase
		}
		result = append(result, customerCondition)
	}

	return result, nil
}

func (r *CustomersMySQL) Amount() ([]domain.CustomerAmount, error) {
	rows, err := r.db.Query(`SELECT c.first_name, c.last_name, TRUNCATE(SUM(i.total), 2) as amount from customers c
                            JOIN invoices i
                            ON i.customer_id = c.id
                            GROUP BY c.first_name, c.last_name
                            ORDER BY amount desc
                            LIMIT 5`)

	if err != nil {
		return nil, apperrors.ErrDatabase
	}
	defer rows.Close()

	if err := rows.Err(); err != nil {
		return nil, apperrors.ErrDatabase
	}

	var result []domain.CustomerAmount
	for rows.Next() {
		var customerAmount domain.CustomerAmount

		if err := rows.Scan(
			&customerAmount.FirstName,
			&customerAmount.LastName,
			&customerAmount.Amount,
		); err != nil {
			return nil, apperrors.ErrDatabase
		}
		result = append(result, customerAmount)
	}

	return result, nil
}
