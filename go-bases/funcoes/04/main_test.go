package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestMinOperation(t *testing.T) {
	expected := 1.0

	result := minOperation(2, 1, 3, 4)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestAverageOperation(t *testing.T) {
	expected := 5.25

	result := averageOperation(5, 4, 4, 8)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestMaxOperation(t *testing.T) {
	expected := 7.0

	result := maxOperation(5, 3, 7, 2)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}
