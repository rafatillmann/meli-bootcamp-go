package router

import (
	"net/http"
	"server/cmd/http/handler"
	"server/internal/product"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
)

func Routers() http.Handler {
	validator := validator.New(validator.WithRequiredStructEnabled())
	repository := product.NewRepository()
	handler := handler.NewHandler(repository, validator)

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
