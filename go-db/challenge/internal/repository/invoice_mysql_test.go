package repository_test

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInvoiceFindAll(t *testing.T) {
	t.Run("returns all invoices", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		expected := []domain.Invoice{
			{Id: 1, InvoiceAttributes: domain.InvoiceAttributes{Datetime: "2022-05-15 00:00:00", CustomerId: 1, Total: 100}},
			{Id: 2, InvoiceAttributes: domain.InvoiceAttributes{Datetime: "2022-05-15 00:00:00", CustomerId: 1, Total: 100}},
		}
		// Invoices
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (1, '2022-05-15 00:00:00', 1, 100)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (2, '2022-05-15 00:00:00', 1, 100)")

		repository := repository.NewInvoicesMySQL(db)

		invoices, err := repository.FindAll()

		require.NoError(t, err)
		require.Equal(t, expected, invoices)
	})
}

func TestInvoiceSave(t *testing.T) {
	t.Run("save invoice", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		invoice := domain.Invoice{Id: 1, InvoiceAttributes: domain.InvoiceAttributes{Datetime: "2022-05-15 00:00:00", CustomerId: 1, Total: 100}}
		// Customer
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")

		repository := repository.NewInvoicesMySQL(db)

		err = repository.Save(&invoice)

		require.NoError(t, err)
	})
}
