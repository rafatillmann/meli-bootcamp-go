package repository_test

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSaleFindAll(t *testing.T) {
	t.Run("returns all sales", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		expected := []domain.Sale{{Id: 1, SaleAttributes: domain.SaleAttributes{ProductId: 1, InvoiceId: 1, Quantity: 3}}}
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (1, '2022-05-15 00:00:00', 1, 100)")
		db.Exec("INSERT INTO products (`id`, `description`, `price`) VALUES (1, 'French Pastry - Mini Chocolate', 97.01)")
		db.Exec("INSERT INTO sales (`id`, `product_id`, `invoice_id`, `quantity`) VALUES (1, 1, 1, 3)")

		repository := repository.NewSalesMySQL(db)

		sales, err := repository.FindAll()

		require.NoError(t, err)
		require.Equal(t, expected, sales)
	})
}

func TestSaleSave(t *testing.T) {
	t.Run("save sale", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		sale := domain.Sale{Id: 1, SaleAttributes: domain.SaleAttributes{ProductId: 1, InvoiceId: 1, Quantity: 3}}
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (1, '2022-05-15 00:00:00', 1, 100)")
		db.Exec("INSERT INTO products (`id`, `description`, `price`) VALUES (1, 'French Pastry - Mini Chocolate', 97.01)")

		repository := repository.NewSalesMySQL(db)

		err = repository.Save(&sale)

		require.NoError(t, err)
	})
}
