package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSalaryA(t *testing.T) {
	expected := 10000.0

	result, err := salary("A", 600)

	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestSalaryB(t *testing.T) {
	expected := 18000.0

	result, err := salary("B", 600)

	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestSalaryC(t *testing.T) {
	expected := 45000.0

	result, err := salary("C", 600)

	require.NoError(t, err)
	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestSalaryErrorCategory(t *testing.T) {

	_, err := salary("D", 600)

	require.Error(t, err)
	require.EqualError(t, err, "Unexpected error")
}
