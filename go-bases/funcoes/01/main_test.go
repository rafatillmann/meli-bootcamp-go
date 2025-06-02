package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTaxLess50000(t *testing.T) {
	salary := 40000.0
	expected := 0.0

	result := tax(salary)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestTaxMore50000(t *testing.T) {
	salary := 60000.0
	expected := 10200.0

	result := tax(salary)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestTaxMore150000(t *testing.T) {
	salary := 160000.0
	expected := 27200.0

	result := tax(salary)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}
