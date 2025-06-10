package router

import (
	"net/http"
	"server/cmd/http/handler"
	"server/internal/product"

	"github.com/go-chi/chi/v5"
)

func ProductRouters() http.Handler {
	repository := product.NewRepository()
	handler := handler.NewHandler(repository)

	r := chi.NewRouter()

	r.Get("/", handler.Products())
	r.Get("/{id}", handler.ProductById())
	r.Get("/search", handler.SearchProducts())

    r.Post("/", handler.AddProduct())

	return r
}
