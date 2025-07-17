package mock

import (
	"testdoubles/simulator"

	"github.com/stretchr/testify/mock"
)

type CatchSimulatorMock struct {
	mock.Mock
}

func NewCatchSimulatorMock() *CatchSimulatorMock {
	return &CatchSimulatorMock{}
}

func (c *CatchSimulatorMock) CanCatch(hunter, prey *simulator.Subject) (canCatch bool) {
	args := c.Called()
	return args.Bool(0)
}
