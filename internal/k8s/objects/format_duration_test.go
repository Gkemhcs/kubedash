package objects

import (
	"testing"
	"time"
)

func TestFormatDuration(t *testing.T) {
	// Define test cases
	testCases := []struct {
		duration   time.Duration
		expected   string
	}{
		{
			duration: 30 * time.Second, // Test with seconds
			expected: "30s",
		},
		{
			duration: 3 * time.Minute, // Test with minutes
			expected: "3m",
		},
		{
			duration: 2 * time.Hour, // Test with hours
			expected: "2h",
		},
		{
			duration: 5 * 24 * time.Hour, // Test with days
			expected: "5d",
		},
		{
			duration: 45 * time.Minute, // Test with minutes less than an hour
			expected: "45m",
		},
		{
			duration: 25 * time.Hour, // Test with more than a day
			expected: "1d", // 25 hours should show as "1d"
		},
		{
			duration: 70 * time.Second, // Test with seconds over 60
			expected: "1m", // 70 seconds should round to 1 minute
		},
	}

	// Run tests
	for _, tc := range testCases {
		t.Run(tc.expected, func(t *testing.T) {
			got := formatDuration(tc.duration)
			if got != tc.expected {
				t.Errorf("formatDuration(%v) = %v, want %v", tc.duration, got, tc.expected)
			}
		})
	}
}
