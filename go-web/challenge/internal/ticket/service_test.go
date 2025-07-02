package ticket_test

import (
	"chanllenge/internal/domain"
	"chanllenge/internal/ticket"
	"chanllenge/test/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestService_GetTotalAmountTickets(t *testing.T) {
	t.Run("success to get total tickets", func(t *testing.T) {

		rp := mock.NewRepositoryTicketMock()

		rp.FuncGet = func() (t map[int]domain.TicketAttributes, err error) {
			t = map[int]domain.TicketAttributes{
				1: {
					Name:    "John",
					Email:   "johndoe@gmail.com",
					Country: "USA",
					Hour:    "10:00",
					Price:   100,
				},
			}
			return
		}

		sv := ticket.NewService(rp)

		total, err := sv.GetTotalAmountTickets()

		expectedTotal := 1
		require.NoError(t, err)
		require.Equal(t, expectedTotal, total)
	})
}
