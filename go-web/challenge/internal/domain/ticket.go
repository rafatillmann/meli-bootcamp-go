package domain

type TicketAttributes struct {
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Country string  `json:"country"`
	Hour    string  `json:"hour"`
	Price   float64 `json:"price"`
}

type Ticket struct {
	Id         int              `json:"id"`
	Attributes TicketAttributes `json:"attributes"`
}

type ServiceTicket interface {
	GetTotalAmountTickets() (total int, err error)
	GetTicketsAmountByDestinationCountry(country string) (total int, err error)
	GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error)
}

type RepositoryTicket interface {
	Get() (t map[int]TicketAttributes, err error)
	GetTicketsByDestinationCountry(country string) (t map[int]TicketAttributes, err error)
}
