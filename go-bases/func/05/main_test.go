package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAnimalDog(t *testing.T) {
	expected := 20.0

	result := animalDog(2)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestAnimalCat(t *testing.T) {
	expected := 15.0

	result := animalCat(3)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestAnimalHamster(t *testing.T) {
	expected := 1.0

	result := animalHamster(4)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

func TestAnimaTarantula(t *testing.T) {
	expected := 0.3

	result := animalTarantula(2)

	require.Equal(t, expected, result, "Verify the expected result is equal to the actual result")
}

