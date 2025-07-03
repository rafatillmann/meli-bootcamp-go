package ticket_test

import (
	"chanllenge/internal/domain"
	"chanllenge/internal/ticket"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGet(t *testing.T) {
	t.Run("success to get tickets from db", func(t *testing.T) {
		expectedTickets := map[int]domain.TicketAttributes{
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
		db := map[int]domain.TicketAttributes{
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
		repository := ticket.NewRepository(db)

		total, err := repository.Get()

		require.Equal(t, total, expectedTickets)
		require.NoError(t, err)
	})
}

func TestGetTicketsByDestinationCountry(t *testing.T) {
	t.Run("success to get tickets from db", func(t *testing.T) {
		expectedTickets := map[int]domain.TicketAttributes{
			1: {
				Name:    "John Doe",
				Email:   "john.doe@email.com",
				Country: "Brazil",
				Hour:    "10:00",
				Price:   100.0,
			},
		}
		db := map[int]domain.TicketAttributes{
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
		repository := ticket.NewRepository(db)

		total, err := repository.GetTicketsByDestinationCountry("brazil")

		require.Equal(t, total, expectedTickets)
		require.NoError(t, err)
	})
}
