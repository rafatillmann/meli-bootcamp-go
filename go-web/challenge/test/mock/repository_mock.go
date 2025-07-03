package mock

import (
	"chanllenge/internal/domain"

	"github.com/stretchr/testify/mock"
)

type RepositoryTicketMock struct {
	mock.Mock
}

func NewRepositoryTicketMock() *RepositoryTicketMock {
	return &RepositoryTicketMock{}
}

func (r *RepositoryTicketMock) Get() (t map[int]domain.TicketAttributes, err error) {
	args := r.Mock.Called()
	return args.Get(0).(map[int]domain.TicketAttributes), args.Error(1)
}

func (r *RepositoryTicketMock) GetTicketsByDestinationCountry(country string) (t map[int]domain.TicketAttributes, err error) {
	args := r.Mock.Called()
	return args.Get(0).(map[int]domain.TicketAttributes), args.Error(1)
}
