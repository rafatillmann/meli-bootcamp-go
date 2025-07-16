package prey_test

import (
	"testdoubles/positioner"
	"testdoubles/prey"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetSpeed(t *testing.T) {
	t.Run("speed is nil", func(t *testing.T) {
		expected := 0.00
		tuna := prey.NewTuna(expected, nil)

		result := tuna.GetSpeed()

		require.Equal(t, expected, result)
	})

	t.Run("speed is not nil", func(t *testing.T) {
		expected := 50.00
		tuna := prey.NewTuna(expected, nil)

		result := tuna.GetSpeed()

		require.Equal(t, expected, result)
	})
}

func TestGetPosition(t *testing.T) {
	t.Run("position is nil", func(t *testing.T) {
		expected := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 10,
		}
		tuna := prey.NewTuna(0.00, expected)

		result := tuna.GetPosition()

		require.Equal(t, expected, result)
	})

	t.Run("position is not nil", func(t *testing.T) {
		tuna := prey.NewTuna(0.00, nil)

		result := tuna.GetPosition()

		require.Nil(t, result)
	})
}
