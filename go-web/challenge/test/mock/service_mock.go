package mock

import "github.com/stretchr/testify/mock"

type ServiceTicketMock struct {
	mock.Mock
}

func NewServiceTicketMock() *ServiceTicketMock {
	return &ServiceTicketMock{}
}

func (s *ServiceTicketMock) GetTotalAmountTickets() (total int, err error) {
	args := s.Mock.Called()
	return args.Get(0).(int), args.Error(1)
}

func (s *ServiceTicketMock) GetTicketsAmountByDestinationCountry(country string) (total int, err error) {
	args := s.Mock.Called()
	return args.Get(0).(int), args.Error(1)
}

func (s *ServiceTicketMock) GetPercentageTicketsByDestinationCountry(country string) (percentage float64, err error) {
	args := s.Mock.Called()
	return args.Get(0).(float64), args.Error(1)
}
