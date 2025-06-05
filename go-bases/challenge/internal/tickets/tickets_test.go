package tickets_test

import (
	"io"
	"math"
	"os"
	"testing"

	"github.com/rafatillmann/meli-bootcamp-go/go-bases/challenge/internal/tickets"
	"github.com/stretchr/testify/require"
)

func mockCSV(t *testing.T) func(name string) error {
	content := `1,John Doe,john@example.com,Brazil,20:00,1000
                2,Jane Doe,jane@example.com,China,17:00,1000
                3,Jim Beam,jim@example.com,Brazil,03:00,2000
                4,Jack Daniels,jack@example.com,China,10:00,2500
                5,Alice Smith,alice@example.com,Brazil,13:00,2000`

	tmpFile, err := os.Create("tickets.csv")
	require.NoError(t, err)
	_, err = io.WriteString(tmpFile, content)
	require.NoError(t, err)
	err = tmpFile.Close()
	require.NoError(t, err)
	return os.Remove
}

func TestGetTotalTickets(t *testing.T) {
	defer mockCSV(t)("tickets.csv")
	testCases := []struct {
		destination string
		expected    int
	}{
		{"Brazil", 3},
		{"China", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.destination, func(t *testing.T) {
			result, _ := tickets.GetTotalTickets(tc.destination)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestGetCountByPeriod(t *testing.T) {
	defer mockCSV(t)("tickets.csv")
	testCases := []struct {
		period   string
		expected int
	}{
		{"early-morning", 1},
		{"morning", 1},
		{"afternoon", 2},
		{"night", 1},
	}

	for _, tc := range testCases {
		t.Run(tc.period, func(t *testing.T) {
			result, _ := tickets.GetCountByPeriod(tc.period)
			require.Equal(t, tc.expected, result)
		})
	}
}

func TestAverageDestination(t *testing.T) {
	defer mockCSV(t)("tickets.csv")
	testCases := []struct {
		destination string
		expected    float64
	}{
		{"Brazil", 60.00},
		{"China", 40.00},
	}

	for _, tc := range testCases {
		t.Run(tc.destination, func(t *testing.T) {
			result, _ := tickets.AverageDestination(tc.destination)
			percentage := math.Round(result) // Round to facilitate comparison
			require.Equal(t, tc.expected, percentage)
		})
	}
}

func TestGetTotalTicketsError(t *testing.T) {
	_, err := tickets.GetTotalTickets("")

	require.Error(t, err)
	require.EqualError(t, err, "unspecified destination")
}

func TestGetCountByPeriodError(t *testing.T) {
	_, err := tickets.GetCountByPeriod("")

	require.Error(t, err)
	require.EqualError(t, err, "unspecified period")
}

func TestAverageDestinationError(t *testing.T) {
	_, err := tickets.AverageDestination("")

	require.Error(t, err)
	require.EqualError(t, err, "unspecified destination")
}
