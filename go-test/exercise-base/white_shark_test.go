package hunt_test

import (
	hunt "testdoubles"
	"testing"

	"github.com/stretchr/testify/require"
)

// Tests for the WhiteShark implementation - Hunt method
func TestWhiteSharkHunt(t *testing.T) {
	t.Run("case 1: white shark hunts successfully", func(t *testing.T) {
		tuna := hunt.NewTuna("Tuna", 50)
		whiteShark := hunt.NewWhiteShark(true, false, 60)

		err := whiteShark.Hunt(tuna)

		require.NoError(t, err)
		require.False(t, whiteShark.Hungry)
		require.True(t, whiteShark.Tired)
	})

	t.Run("case 2: white shark is not hungry", func(t *testing.T) {
		tuna := hunt.NewTuna("Tuna", 50)
		whiteShark := hunt.NewWhiteShark(false, false, 60)

		err := whiteShark.Hunt(tuna)

		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrSharkIsNotHungry)
	})

	t.Run("case 3: white shark is tired", func(t *testing.T) {
		tuna := hunt.NewTuna("Tuna", 50)
		whiteShark := hunt.NewWhiteShark(true, true, 60)

		err := whiteShark.Hunt(tuna)

		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrSharkIsTired)
	})

	t.Run("case 4: white shark is slower than the tuna", func(t *testing.T) {
		tuna := hunt.NewTuna("Tuna", 50)
		whiteShark := hunt.NewWhiteShark(true, false, 40)

		err := whiteShark.Hunt(tuna)

		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrSharkIsSlower)
	})

	t.Run("case 5: tuna is nil", func(t *testing.T) {
		whiteShark := hunt.NewWhiteShark(true, false, 40)

		err := whiteShark.Hunt(nil)

		require.Error(t, err)
		require.ErrorIs(t, err, hunt.ErrTunaIsNil)
	})
}
