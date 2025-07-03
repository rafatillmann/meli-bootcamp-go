package ticket_test

import (
	"chanllenge/internal/domain"
	"chanllenge/internal/ticket"
	"chanllenge/test/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestServiceGetTotalAmountTickets(t *testing.T) {
	t.Run("success to get amount tickets", func(t *testing.T) {

		tickets := map[int]domain.TicketAttributes{
			1: {
				Name:    "John Doe",
				Email:   "john.doe@email.com",
				Country: "Brazil",
				Hour:    "10:00",
				Price:   100.0,
			},
		}
		expectedTotal := 1

		repository := mock.NewRepositoryTicketMock()
		service := ticket.NewService(repository)

		repository.On("Get").Return(tickets, nil)

		total, err := service.GetTotalAmountTickets()

		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

func TestGetTicketsAmountByDestinationCountry(t *testing.T) {
	t.Run("success to get amount tickets by country", func(t *testing.T) {

		tickets := map[int]domain.TicketAttributes{
			1: {
				Name:    "John Doe",
				Email:   "john.doe@email.com",
				Country: "Brazil",
				Hour:    "10:00",
				Price:   100.0,
			},
		}
		expectedTotal := 1

		repository := mock.NewRepositoryTicketMock()
		service := ticket.NewService(repository)

		repository.On("GetTicketsByDestinationCountry").Return(tickets, nil)

		total, err := service.GetTicketsAmountByDestinationCountry("Brazil")

		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}

func TestGetPercentageTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get percentage of tickets by country", func(t *testing.T) {

		tickets := map[int]domain.TicketAttributes{
			1: {
				Name:    "John Doe",
				Email:   "john.doe@email.com",
				Country: "Brazil",
				Hour:    "10:00",
				Price:   100.0,
			},
			2: {
				Name:    "Amanda Smith",
				Email:   "amanda.smith@email.cim",
				Country: "China",
				Hour:    "11:00",
				Price:   200.0,
			},
		}

		ticketsByCountry := map[int]domain.TicketAttributes{
			1: {
				Name:    "John Doe",
				Email:   "john.doe@email.com",
				Country: "Brazil",
				Hour:    "10:00",
				Price:   100.0,
			},
		}

		expectedPercentage := 50.0

		repository := mock.NewRepositoryTicketMock()
		service := ticket.NewService(repository)

		repository.On("Get").Return(tickets, nil)
		repository.On("GetTicketsByDestinationCountry").Return(ticketsByCountry, nil)

		percentage, err := service.GetPercentageTicketsByDestinationCountry("Brazil")

		require.NoError(t, err)
		require.Equal(t, expectedPercentage, percentage)
	})
}
