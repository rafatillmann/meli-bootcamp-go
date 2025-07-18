package hunter_test

import (
	"testdoubles/hunter"
	"testdoubles/positioner"
	"testdoubles/simulator"
	"testdoubles/test/mock"
	"testdoubles/test/stub"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHunt(t *testing.T) {
	t.Run("can hunt", func(t *testing.T) {
		prey := stub.NewPreyStub()
		preySubject := &simulator.Subject{
			Position: prey.GetPosition(),
			Speed:    prey.GetSpeed(),
		}
		simulatorMock := mock.NewCatchSimulatorMock()
		whiteSharkPosition := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 10,
		}
		whiteSharkSpeed := 10.00
		whiteShark := hunter.NewWhiteShark(whiteSharkSpeed, whiteSharkPosition, simulatorMock)
		whiteSharkSubject := &simulator.Subject{
			Position: whiteSharkPosition,
			Speed:    whiteSharkSpeed,
		}

		simulatorMock.On("CanCatch", whiteSharkSubject, preySubject).Return(true)

		err := whiteShark.Hunt(prey)

		require.NoError(t, err)
	})

	t.Run("can not hunt", func(t *testing.T) {
		prey := stub.NewPreyStub()
		preySubject := &simulator.Subject{
			Position: prey.GetPosition(),
			Speed:    prey.GetSpeed(),
		}
		simulatorMock := mock.NewCatchSimulatorMock()
		whiteSharkPosition := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 10,
		}
		whiteSharkSpeed := 10.00
		whiteShark := hunter.NewWhiteShark(whiteSharkSpeed, whiteSharkPosition, simulatorMock)
		whiteSharkSubject := &simulator.Subject{
			Position: whiteSharkPosition,
			Speed:    whiteSharkSpeed,
		}

		simulatorMock.On("CanCatch", whiteSharkSubject, preySubject).Return(false)

		err := whiteShark.Hunt(prey)

		require.Error(t, err)
		require.ErrorIs(t, err, hunter.ErrCanNotHunt)
	})
}
