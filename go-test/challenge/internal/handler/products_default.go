package handler

import (
	"app/internal"
	"app/platform/web/response"
	"net/http"
	"strconv"
)

func NewProductsDefault(rp internal.RepositoryProducts) *ProductsDefault {
	return &ProductsDefault{
		rp: rp,
	}
}

type ProductsDefault struct {
	rp internal.RepositoryProducts
}

type ProductJSON struct {
	ID          int     `json:"id"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	SellerId    int     `json:"seller_id"`
}

func (h *ProductsDefault) Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var query internal.ProductQuery
		if r.URL.Query().Has("id") {
			var err error
			query.ID, err = strconv.Atoi(r.URL.Query().Get("id"))
			if err != nil {
				response.Error(w, http.StatusBadRequest, "invalid id")
				return
			}
		}

		p, err := h.rp.SearchProducts(query)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "internal error")
			return
		}

		data := make(map[int]ProductJSON)
		for k, v := range p {
			data[k] = ProductJSON{
				ID:          v.ID,
				Description: v.Description,
				Price:       v.Price,
				SellerId:    v.SellerId,
			}
		}
		response.JSON(w, http.StatusOK, map[string]any{
			"message": "success",
			"data":    data,
		})
	}
}
