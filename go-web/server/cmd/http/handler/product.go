package handler

import (
	"encoding/json"
	"net/http"
	"server/internal/domain"
	"server/internal/product"
	"server/pkg/response"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type ProductHandler struct {
	repository product.Repository
}

func NewHandler(repository product.Repository) *ProductHandler {
	return &ProductHandler{repository}
}

func (h *ProductHandler) Products() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(h.repository.Get())
	}
}

func (h *ProductHandler) ProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			body := response.ResponseError{Message: "ID must be a valid integer"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(body)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(h.repository.GetByID(ID))
	}
}

func (h *ProductHandler) SearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)

		if err != nil {
			body := response.ResponseError{Message: "priceGt is required"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(body)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(h.repository.GetFilterByPrice(priceGt))
	}
}


// TODO - Adicionar as validacões para os atributos
func (h *ProductHandler) AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var requestProduct domain.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&requestProduct); err != nil {
			body := response.ResponseError{Message: "Bad request"}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(body)
			return
		}

		product := h.repository.Create(requestProduct)

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}
