package repository_test

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCustomerFindAll(t *testing.T) {
	t.Run("returns all customers", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		expected := []domain.Customer{{Id: 1, CustomerAttributes: domain.CustomerAttributes{FirstName: "Fifield", LastName: "Ike", Condition: 1}}}
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")

		repository := repository.NewCustomersMySQL(db)

		customers, err := repository.FindAll()

		require.NoError(t, err)
		require.Equal(t, expected, customers)
	})
}

func TestCustomerSave(t *testing.T) {
	t.Run("save customer", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		customer := domain.Customer{CustomerAttributes: domain.CustomerAttributes{FirstName: "Fifield", LastName: "Ike", Condition: 1}}

		repository := repository.NewCustomersMySQL(db)

		err = repository.Save(&customer)

		require.NoError(t, err)
	})
}

func TestCustomerAmount(t *testing.T) {
	t.Run("returns customers amount", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		expected := []domain.Customer{
			{Id: 1, CustomerAttributes: domain.CustomerAttributes{FirstName: "Fifield", LastName: "Ike", Condition: 1}},
			{Id: 2, CustomerAttributes: domain.CustomerAttributes{FirstName: "Cowland", LastName: "Brannon", Condition: 1}},
		}
		// Fifield
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (1, '2022-05-15', 1, 200)")
		// Cowland
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (2, 'Cowland', 'Brannon', 1)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (2, '2022-05-15', 2, 100)")

		repository := repository.NewCustomersMySQL(db)

		customers, err := repository.FindAll()

		require.NoError(t, err)
		require.Equal(t, expected, customers)
	})
}

func TestCustomerTotalCondition(t *testing.T) {
	t.Run("total by condition", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		expected := []domain.CustomerCondition{
			{Condition: 1, Total: 200},
			{Condition: 0, Total: 400},
		}
		// Fifield
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (1, '2022-05-15', 1, 200)")
		// Cowland
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (2, 'Cowland', 'Brannon', 0)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (2, '2022-05-15', 2, 200)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (3, '2022-05-15', 2, 200)")

		repository := repository.NewCustomersMySQL(db)

		result, err := repository.TotalByCondition()

		require.NoError(t, err)
		require.Equal(t, expected, result)
	})
}
