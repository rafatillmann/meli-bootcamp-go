package ticket_test

import "testing"

// Tests for ServiceTicketDefault.GetTotalAmountTickets
func TestServiceTicketDefault_GetTotalAmountTickets(t *testing.T) {
	// t.Run("success to get total tickets", func(t *testing.T) {
	// 	// arrange
	// 	// - repository: mock
	// 	rp := ticket.NewRepositoryTicketMock()
	// 	// - repository: set-up
	// 	rp.FuncGet = func() (t map[int]domain.TicketAttributes, err error) {
	// 		t = map[int]domain.TicketAttributes{
	// 			1: {
	// 				Name:    "John",
	// 				Email:   "johndoe@gmail.com",
	// 				Country: "USA",
	// 				Hour:    "10:00",
	// 				Price:   100,
	// 			},
	// 		}
	// 		return
	// 	}

	// 	// - service
	// 	sv := ticket.NewServiceTicketDefault(rp)

	// 	// act
	// 	total, err := sv.GetTotalAmountTickets()

	// 	// assert
	// 	expectedTotal := 1
	// 	require.NoError(t, err)
	// 	require.Equal(t, expectedTotal, total)
	// })
}
