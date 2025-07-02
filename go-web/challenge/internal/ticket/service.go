package ticket

import "chanllenge/internal/domain"

// ServiceDefault represents the default service of the tickets
type ServiceDefault struct {
	rp domain.RepositoryTicket
}

func NewService(rp domain.RepositoryTicket) *ServiceDefault {
	return &ServiceDefault{
		rp: rp,
	}
}

func (s *ServiceDefault) GetTotalAmountTickets() (total int, err error) {
	tickets, err := s.rp.Get()
	if err != nil {
		return
	}

	total = len(tickets)
	return
}

func (s *ServiceDefault) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	ticketsByCountry, err := s.rp.GetTicketsByDestinationCountry(country)
	if err != nil {
		return
	}

	total = len(ticketsByCountry)
	return
}

func (s *ServiceDefault) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	total, err := s.GetTotalAmountTickets()
	if err != nil {
		return
	}
	totalCountry, err := s.GetTicketsAmountByDestinationCountry(country)
	if err != nil {
		return
	}

	percentage = float64(totalCountry) / float64(total) * 100
	return
}
