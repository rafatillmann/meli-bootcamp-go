package positioner_test

import (
	"testdoubles/positioner"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGetLinearDistance(t *testing.T) {
	t.Run("negative coordinates", func(t *testing.T) {
		expectedLinearDistance := 14.142135623730951
		from := &positioner.Position{
			X: -10,
			Y: -10,
			Z: -10,
		}
		to := &positioner.Position{
			X: -10,
			Y: -20,
			Z: -20,
		}
		positioner := positioner.NewPositionerDefault()

		linearDistance := positioner.GetLinearDistance(from, to)

		require.Equal(t, expectedLinearDistance, linearDistance)
	})

	t.Run("positive coordinates", func(t *testing.T) {
		expectedLinearDistance := 28.284271247461902
		from := &positioner.Position{
			X: 30,
			Y: 20,
			Z: 10,
		}
		to := &positioner.Position{
			X: 10,
			Y: 20,
			Z: 30,
		}
		positioner := positioner.NewPositionerDefault()

		linearDistance := positioner.GetLinearDistance(from, to)

		require.Equal(t, expectedLinearDistance, linearDistance)
	})

	t.Run("linear distance without decimal", func(t *testing.T) {
		expectedLinearDistance := 10.00
		from := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 10,
		}
		to := &positioner.Position{
			X: 10,
			Y: 10,
			Z: 20,
		}
		positioner := positioner.NewPositionerDefault()

		linearDistance := positioner.GetLinearDistance(from, to)

		require.Equal(t, expectedLinearDistance, linearDistance)
	})
}
