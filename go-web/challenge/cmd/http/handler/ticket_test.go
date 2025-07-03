package handler_test

import (
	"chanllenge/cmd/http/handler"
	"chanllenge/test/mock"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetAmount(t *testing.T) {
	t.Run("sucess to get amount", func(t *testing.T) {
		total := 10
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := fmt.Sprintf(`{"total":%d}`, total)

		service := mock.NewServiceTicketMock()
		handler := handler.NewTicketHandler(service)

		service.On("GetTotalAmountTickets").Return(total, nil)

		r := httptest.NewRequest(http.MethodGet, "/tickets/amount", nil)
		w := httptest.NewRecorder()

		handler.GetAmount().ServeHTTP(w, r)

		require.Equal(t, expectedCode, w.Code)
		require.Equal(t, expectedHeader, w.Header())
		require.Equal(t, expectedBody, w.Body.String())
	})
}

func TestGetAmountByDestinationCountry(t *testing.T) {
	t.Run("sucess to get amount by country", func(t *testing.T) {
		total := 5
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := fmt.Sprintf(`{"total":%d}`, total)

		service := mock.NewServiceTicketMock()
		handler := handler.NewTicketHandler(service)

		service.On("GetTicketsAmountByDestinationCountry").Return(total, nil)

		r := httptest.NewRequest(http.MethodGet, "/tickets/amount/brazil", nil)
		w := httptest.NewRecorder()

		handler.GetAmountByDestinationCountry().ServeHTTP(w, r)

		require.Equal(t, expectedCode, w.Code)
		require.Equal(t, expectedHeader, w.Header())
		require.Equal(t, expectedBody, w.Body.String())
	})
}

func TestGetPercentageByDestinationCountry(t *testing.T) {
	t.Run("sucess to get percentage by country", func(t *testing.T) {
		percentage := 50.0
		expectedCode := http.StatusOK
		expectedHeader := http.Header{"Content-Type": []string{"application/json"}}
		expectedBody := fmt.Sprintf(`{"percentage":%.0f}`, percentage)

		service := mock.NewServiceTicketMock()
		handler := handler.NewTicketHandler(service)

		service.On("GetPercentageTicketsByDestinationCountry").Return(percentage, nil)

		r := httptest.NewRequest(http.MethodGet, "/tickets/percentage/brazil", nil)
		w := httptest.NewRecorder()

		handler.GetPercentageByDestinationCountry().ServeHTTP(w, r)

		require.Equal(t, expectedCode, w.Code)
		require.Equal(t, expectedHeader, w.Header())
		require.Equal(t, expectedBody, w.Body.String())
	})
}
