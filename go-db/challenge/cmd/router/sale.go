package router

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func SaleRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	repository := repository.NewSalesMySQL(db)
	service := service.NewSalesDefault(repository)
	handler := handler.NewSalesDefault(service)

	r.Get("/", handler.GetAll())
	r.Post("/", handler.Create())
	return r
}
