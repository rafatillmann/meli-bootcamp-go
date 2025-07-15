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

func (r *CatchSimulatorMock) CanCatch(hunter, prey *simulator.Subject) (canCatch bool) {
	args := r.Mock.Called()
	return args.Bool(0)
}
