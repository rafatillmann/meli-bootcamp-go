package simulator_test

import (
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testdoubles/test/mock"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCanCatch(t *testing.T) {
	t.Run("can catch", func(t *testing.T) {
		mockPositioner := mock.NewPositionerMock()
		hunter := &simulator.Subject{
			Position: &positioner.Position{},
			Speed:    20,
		}
		prey := &simulator.Subject{
			Position: &positioner.Position{},
			Speed:    10,
		}
		simulator := simulator.NewCatchSimulatorDefault(10.00, mockPositioner)
		mockPositioner.On("GetLinearDistance").Return(10.00)

		canCatch := simulator.CanCatch(hunter, prey)

		require.True(t, canCatch)
	})
	t.Run("can not catch, too slow", func(t *testing.T) {
		mockPositioner := mock.NewPositionerMock()
		hunter := &simulator.Subject{
			Position: &positioner.Position{},
			Speed:    10,
		}
		prey := &simulator.Subject{
			Position: &positioner.Position{},
			Speed:    20,
		}
		simulator := simulator.NewCatchSimulatorDefault(10.00, mockPositioner)
		mockPositioner.On("GetLinearDistance").Return(10.00)

		canCatch := simulator.CanCatch(hunter, prey)

		require.False(t, canCatch)
	})

	t.Run("can not catch, not enought time", func(t *testing.T) {
		mockPositioner := mock.NewPositionerMock()
		hunter := &simulator.Subject{
			Position: &positioner.Position{},
			Speed:    20,
		}
		prey := &simulator.Subject{
			Position: &positioner.Position{},
			Speed:    10,
		}
		simulator := simulator.NewCatchSimulatorDefault(10.00, mockPositioner)
		mockPositioner.On("GetLinearDistance").Return(1000.00)

		canCatch := simulator.CanCatch(hunter, prey)

		require.False(t, canCatch)
	})
}
