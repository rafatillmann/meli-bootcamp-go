package router

import (
	"net/http"
	"server/cmd/http/handler"
	"server/internal/product"

	"github.com/go-chi/chi/v5"
)

func Routers() http.Handler {
	repository := product.NewRepository()
	handler := handler.NewHandler(repository)

	r := chi.NewRouter()

	r.Get("/", handler.Products())
	r.Get("/{id}", handler.ProductById())
	r.Get("/search", handler.SearchProducts())

	r.Post("/", handler.AddProduct())
	r.Put("/{id}", handler.UpdateProduct())
	r.Patch("/{id}", handler.PartialUpdateProduct())
	r.Delete("/{id}", handler.DeleteProduct())

	return r
}
