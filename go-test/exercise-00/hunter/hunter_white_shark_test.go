package hunter_test

import (
	"testdoubles/hunter"
	"testdoubles/test/mock"
	"testdoubles/test/stub"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHunt(t *testing.T) {
	t.Run("can hunt", func(t *testing.T) {
		prey := stub.NewPreyStub()
		simulatorMock := mock.NewCatchSimulatorMock()
		whiteShark := hunter.CreateWhiteShark(simulatorMock)

		simulatorMock.On("CanCatch").Return(true)

		err := whiteShark.Hunt(prey)

		require.NoError(t, err)
	})

	t.Run("can not hunt", func(t *testing.T) {
		prey := stub.NewPreyStub()
		simulatorMock := mock.NewCatchSimulatorMock()
		whiteShark := hunter.CreateWhiteShark(simulatorMock)

		simulatorMock.On("CanCatch").Return(false)

		err := whiteShark.Hunt(prey)

		require.Error(t, err)
		require.ErrorIs(t, err, hunter.ErrCanNotHunt)
	})
}
