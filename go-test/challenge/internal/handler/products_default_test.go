package handler_test

import (
	"app/internal"
	"app/internal/handler"
	"app/internal/repository/mocks"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestHandlerGet(t *testing.T) {
	t.Run("success - should return all products when query is empty", func(t *testing.T) {
		products := map[int]internal.Product{
			1: {ID: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
			2: {ID: 2, ProductAttributes: internal.ProductAttributes{Description: "Product 2", Price: 100, SellerId: 1}},
		}

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"message":"success","data":{"1":{"id":1,"description":"Product 1","price":100,"seller_id":1},"2":{"id":2,"description":"Product 2","price":100,"seller_id":1}}}`

		repository := mocks.NewProductsRepositoryMock()
		handler := handler.NewProductsDefault(repository)

		repository.On("SearchProducts", internal.ProductQuery{}).Return(products, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			`/products`,
			nil,
		)

		handler.Get().ServeHTTP(w, r)

		require.Equal(t, expectedCode, w.Code)
		require.Equal(t, expectedHeader, w.Header())
		require.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("success - should return products matching the query", func(t *testing.T) {
		products := map[int]internal.Product{
			1: {ID: 1, ProductAttributes: internal.ProductAttributes{Description: "Product 1", Price: 100, SellerId: 1}},
		}

		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"message":"success","data":{"1":{"id":1,"description":"Product 1","price":100,"seller_id":1}}}`

		repository := mocks.NewProductsRepositoryMock()
		handler := handler.NewProductsDefault(repository)

		repository.On("SearchProducts", internal.ProductQuery{ID: 1}).Return(products, nil)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf(`/products?id=%d`, 1),
			nil,
		)

		handler.Get().ServeHTTP(w, r)

		require.Equal(t, expectedCode, w.Code)
		require.Equal(t, expectedHeader, w.Header())
		require.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("err - invalid id", func(t *testing.T) {
		expectedCode := http.StatusBadRequest
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"message":"invalid id","status":"Bad Request"}`

		repository := mocks.NewProductsRepositoryMock()
		handler := handler.NewProductsDefault(repository)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf(`/products?id=%s`, "invalid"),
			nil,
		)

		handler.Get().ServeHTTP(w, r)

		require.Equal(t, expectedCode, w.Code)
		require.Equal(t, expectedHeader, w.Header())
		require.JSONEq(t, expectedBody, w.Body.String())
	})

	t.Run("err - internal error", func(t *testing.T) {
		expectedCode := http.StatusInternalServerError
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := `{"message":"internal error","status":"Internal Server Error"}`

		repository := mocks.NewProductsRepositoryMock()
		handler := handler.NewProductsDefault(repository)

		repository.On("SearchProducts", internal.ProductQuery{ID: 1}).Return(map[int]internal.Product{}, assert.AnError)

		w := httptest.NewRecorder()
		r := httptest.NewRequest(
			http.MethodGet,
			fmt.Sprintf(`/products?id=%d`, 1),
			nil,
		)

		handler.Get().ServeHTTP(w, r)

		require.Equal(t, expectedCode, w.Code)
		require.Equal(t, expectedHeader, w.Header())
		require.JSONEq(t, expectedBody, w.Body.String())
	})
}
