package repository_test

import (
	"app/internal/domain"
	"app/internal/repository"
	"app/test"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestProductFindAll(t *testing.T) {
	t.Run("returns all products", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		expected := []domain.Product{{Id: 1, ProductAttributes: domain.ProductAttributes{Description: "French Pastry - Mini Chocolate", Price: 97.01}}}
		db.Exec("INSERT INTO products (`id`, `description`, `price`) VALUES (1, 'French Pastry - Mini Chocolate', 97.01)")

		repository := repository.NewProductsMySQL(db)

		products, err := repository.FindAll()

		require.NoError(t, err)
		require.Equal(t, expected, products)
	})
}

func TestProductSave(t *testing.T) {
	t.Run("save product", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		product := domain.Product{Id: 1, ProductAttributes: domain.ProductAttributes{Description: "French Pastry - Mini Chocolate", Price: 97.01}}

		repository := repository.NewProductsMySQL(db)

		err = repository.Save(&product)

		require.NoError(t, err)
	})
}

func TestProductBestSeller(t *testing.T) {
	t.Run("returns best sellers products", func(t *testing.T) {
		db, err := test.GetTxdb()
		require.NoError(t, err)
		defer db.Close()

		expected := []domain.ProductBestSeller{
			{Description: "Beans - Soya Bean", Total: 5},
			{Description: "French Pastry - Mini Chocolate", Total: 3},
		}
		// 01
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (1, 'Fifield', 'Ike', 1)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (1, '2022-05-15 00:00:00', 1, 100)")
		db.Exec("INSERT INTO products (`id`, `description`, `price`) VALUES (1, 'French Pastry - Mini Chocolate', 97.01)")
		db.Exec("INSERT INTO sales (`id`, `product_id`, `invoice_id`, `quantity`) VALUES (1, 1, 1, 3)")
		// 02
		db.Exec("INSERT INTO customers (`id`, `first_name`, `last_name`, `condition`) VALUES (2, 'Cowland', 'Brannon', 0)")
		db.Exec("INSERT INTO invoices (`id`, `datetime`, `customer_id`, `total`) VALUES (2, '2022-05-15 00:00:00', 2, 100)")
		db.Exec("INSERT INTO products (`id`, `description`, `price`) VALUES (2, 'Beans - Soya Bean', 12.89)")
		db.Exec("INSERT INTO sales (`id`, `product_id`, `invoice_id`, `quantity`) VALUES (2, 2, 2, 5)")

		repository := repository.NewProductsMySQL(db)

		products, err := repository.BestSellers()

		require.NoError(t, err)
		require.Equal(t, expected, products)
	})
}
