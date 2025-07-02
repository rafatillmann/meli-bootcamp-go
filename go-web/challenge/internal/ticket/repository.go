package ticket

import (
	"chanllenge/internal/domain"
	"strings"
)

type RepositoryDefault struct {
	db     map[int]domain.TicketAttributes
	lastId int
}

func NewRepository(db map[int]domain.TicketAttributes) *RepositoryDefault {
	defaultLastId := 0
	for key := range db {
		if key > defaultLastId {
			defaultLastId = key
		}
	}

	return &RepositoryDefault{
		db:     db,
		lastId: defaultLastId,
	}
}

func (r *RepositoryDefault) Get() (t map[int]domain.TicketAttributes, err error) {
	t = make(map[int]domain.TicketAttributes, len(r.db))
	for k, v := range r.db {
		t[k] = v
	}

	return
}

func (r *RepositoryDefault) GetTicketsByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error) {
	t = make(map[int]domain.TicketAttributes)
	for k, v := range r.db {
		if strings.EqualFold(v.Country, country) {
			t[k] = v
		}
	}

	return
}
