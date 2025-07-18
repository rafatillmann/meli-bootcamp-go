package router

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func ProductRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	repository := repository.NewProductsMySQL(db)
	service := service.NewProductsDefault(repository)
	handler := handler.NewProductsDefault(service)

	r.Get("/", handler.GetAll())
	r.Post("/", handler.Create())
	return r
}
