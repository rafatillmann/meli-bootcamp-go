package repository_test

import (
	"app/internal"
	"app/internal/repository"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRepositorySearchProducts(t *testing.T) {
	t.Run("should return all products when query is empty", func(t *testing.T) {
		expected := map[int]internal.Product{
			1: {ID: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
			2: {ID: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 100, SellerId: 1}},
		}

		db := map[int]internal.Product{
			1: {ID: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
			2: {ID: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 100, SellerId: 1}},
		}
		repository := repository.NewProductsMap(db)

		result, err := repository.SearchProducts(internal.ProductQuery{})

		require.NoError(t, err)
		require.Equal(t, expected, result)
	})

	t.Run("should return products matching the query", func(t *testing.T) {
		expected := map[int]internal.Product{
			1: {ID: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
		}

		db := map[int]internal.Product{
			1: {ID: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
			2: {ID: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 100, SellerId: 1}},
		}
		repository := repository.NewProductsMap(db)

		result, err := repository.SearchProducts(internal.ProductQuery{ID: 1})

		require.NoError(t, err)
		require.Equal(t, expected, result)
	})
}
