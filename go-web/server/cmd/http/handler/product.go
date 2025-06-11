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
		json.NewEncoder(w).Encode(h.repository.GetAll())
	}
}

func (h *ProductHandler) ProductById() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Error(w, "ID must be a valid integer", http.StatusBadRequest)
			return
		}

		product, err := h.repository.GetByID(ID)
		if err != nil {
			response.Error(w, "Product not found", http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func (h *ProductHandler) SearchProducts() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		priceGt, err := strconv.ParseFloat(r.URL.Query().Get("priceGt"), 64)

		if err != nil {
			response.Error(w, "priceGt is required", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(h.repository.GetFilterByPrice(priceGt))
	}
}

func (h *ProductHandler) AddProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var request domain.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		product := domain.Product{
			Name:        request.Name,
			Quantity:    request.Quantity,
			CodeValue:   request.CodeValue,
			IsPublished: request.IsPublished,
			Expiration:  request.Expiration,
			Price:       request.Price,
		}
		if err := h.repository.Create(&product); err != nil {
			response.Error(w, "Unable to add product", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(product)
	}
}

func (h *ProductHandler) UpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Error(w, "ID must be a valid integer", http.StatusBadRequest)
			return
		}

		var request domain.ProductRequest
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			response.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		product := domain.Product{
			ID:          ID,
			Name:        request.Name,
			Quantity:    request.Quantity,
			CodeValue:   request.CodeValue,
			IsPublished: request.IsPublished,
			Expiration:  request.Expiration,
			Price:       request.Price,
		}
		if err := h.repository.Update(&product); err != nil {
			response.Error(w, "Unable to update product", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

// Verificar Reflection
func (h *ProductHandler) PartialUpdateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Error(w, "ID must be a valid integer", http.StatusBadRequest)
			return
		}

		product, err := h.repository.GetByID(ID)
		if err != nil {
			response.Error(w, "Product not found", http.StatusNotFound)
		}

		if err := json.NewDecoder(r.Body).Decode(product); err != nil {
			response.Error(w, "Bad request", http.StatusBadRequest)
			return
		}

		if err := h.repository.Update(product); err != nil {
			response.Error(w, "Unable to update product", http.StatusBadRequest)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(product)
	}
}

func (h *ProductHandler) DeleteProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		ID, err := strconv.Atoi(chi.URLParam(r, "id"))

		if err != nil {
			response.Error(w, "ID must be a valid integer", http.StatusBadRequest)
			return
		}

		if err := h.repository.Delete(ID); err != nil {
			response.Error(w, "Product not found", http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusNoContent)
		json.NewEncoder(w).Encode(nil)
	}
}
