package mock

import (
	"testdoubles/positioner"

	"github.com/stretchr/testify/mock"
)

type PositionerMock struct {
	mock.Mock
}

func NewPositionerMock() *PositionerMock {
	return &PositionerMock{}
}

func (p *PositionerMock) GetLinearDistance(from, to *positioner.Position) (linearDistance float64) {
	args := p.Called()
	return args.Get(0).(float64)
}
