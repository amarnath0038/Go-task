package service

import (
	"testing"
	"time"
)

func TestCalculateAge(t *testing.T) {
	tests := []struct {
		name     string
		dob      time.Time
		expected int
	}{
		{
			name:     "birthday already passed this year",
			dob:      time.Now().AddDate(-25, -1, 0),
			expected: 25,
		},
		{
			name:     "birthday not yet reached this year",
			dob:      time.Now().AddDate(-25, 1, 0),
			expected: 24,
		},
		{
			name:     "birthday is today",
			dob:      time.Now().AddDate(-30, 0, 0),
			expected: 30,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			age := CalculateAge(tt.dob)
			if age != tt.expected {
				t.Fatalf("expected age %d, got %d", tt.expected, age)
			}
		})
	}
}
