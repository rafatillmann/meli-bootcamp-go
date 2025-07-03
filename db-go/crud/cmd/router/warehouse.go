package router

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func WarehouseRouter(db *sql.DB) http.Handler {
	rt := chi.NewRouter()

	rt.Get("/{id}", nil)
	rt.Post("/", nil)

	return rt
}
