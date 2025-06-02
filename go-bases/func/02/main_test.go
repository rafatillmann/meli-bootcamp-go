package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {
	expected := 5.0

	result := average(5, 6, 4)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}
