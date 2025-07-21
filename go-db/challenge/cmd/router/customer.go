package router

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CustomerRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	repository := repository.NewCustomersMySQL(db)
	service := service.NewCustomersDefault(repository)
	handler := handler.NewCustomersDefault(service)

	r.Get("/", handler.GetAll())
	r.Post("/", handler.Create())
	r.Get("/total-condition", handler.TotalByCondition())
	r.Get("/amount", handler.Amount())
	return r
}
