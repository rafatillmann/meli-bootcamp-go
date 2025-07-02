package handler

import (
	"chanllenge/internal/domain"
	"net/http"

	"github.com/bootcamp-go/web/response"
	"github.com/go-chi/chi/v5"
)

type TicketHandler struct {
	sv domain.ServiceTicket
}

func NewTicketHandler(sv domain.ServiceTicket) *TicketHandler {
	return &TicketHandler{sv}
}

func (h *TicketHandler) GetAmount() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		total, err := h.sv.GetTotalAmountTickets()
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Unable to get amount of tickets")
		}
		response.JSON(w, http.StatusOK, map[string]int{"total": total})
	}
}

func (h *TicketHandler) GetAmountByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "country")
		total, err := h.sv.GetTicketsAmountByDestinationCountry(country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Unable to get tickets by destination country")
		}
		response.JSON(w, http.StatusOK, map[string]int{"total": total})
	}
}

func (h *TicketHandler) GetPercentageByCountry() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		country := chi.URLParam(r, "country")
		percentage, err := h.sv.GetPercentageTicketsByDestinationCountry(country)
		if err != nil {
			response.Error(w, http.StatusInternalServerError, "Unable to get percentage of tickets by destination country")
		}

		response.JSON(w, http.StatusOK, map[string]float64{"percentage": percentage})
	}
}
