package router

import (
	"app/internal/handler"
	"app/internal/repository"
	"app/internal/service"
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func InvoiceRouter(db *sql.DB) http.Handler {
	r := chi.NewRouter()
	repository := repository.NewInvoicesMySQL(db)
	service := service.NewInvoicesDefault(repository)
	handler := handler.NewInvoicesDefault(service)

	r.Get("/", handler.GetAll())
	r.Post("/", handler.Create())
	r.Post("/recalculate-total/{invoiceId}", handler.RecalcuteTotal())
	return r
}
