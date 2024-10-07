package problem_solving

import (
	"fmt"
	"testing"
)

func TestGroupOpenHours(t *testing.T) {
	tests := []struct {
		name      string
		openHours []map[string]string
		expected  []map[string]string
	}{
		{
			name: "Test 1",
			openHours: []map[string]string{
				{"day": "Monday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Tuesday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Wednesday", "open": "8:00 AM", "close": "6:00 PM"},
				{"day": "Thursday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Friday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Saturday", "open": "8:00 AM", "close": "4:00 PM"},
			},
			expected: []map[string]string{
				{"days": "Monday-Tuesday", "open": "8:00 AM", "close": "5:00 PM"},
				{"days": "Wednesday", "open": "8:00 AM", "close": "6:00 PM"},
				{"days": "Thursday-Friday", "open": "8:00 AM", "close": "5:00 PM"},
				{"days": "Saturday", "open": "8:00 AM", "close": "4:00 PM"},
				{"days": "Sunday", "open": "", "close": ""},
			},
		},
		{
			name: "Test 2",
			openHours: []map[string]string{
				{"day": "Monday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Tuesday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Wednesday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Thursday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Friday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Saturday", "open": "8:00 AM", "close": "5:00 PM"},
				{"day": "Sunday", "open": "8:00 AM", "close": "5:00 PM"},
			},
			expected: []map[string]string{
				{"days": "Monday-Sunday", "open": "8:00 AM", "close": "5:00 PM"},
			},
		},
		{
			name: "all closed week",
			openHours: []map[string]string{
				{"day": "Monday", "open": "", "close": ""},
				{"day": "Tuesday", "open": "", "close": ""},
				{"day": "Wednesday", "open": "", "close": ""},
				{"day": "Thursday", "open": "", "close": ""},
				{"day": "Friday", "open": "", "close": ""},
				{"day": "Saturday", "open": "", "close": ""},
				{"day": "Sunday", "open": "", "close": ""},
			},
			expected: []map[string]string{
				{"days": "Monday-Sunday", "open": "", "close": ""},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := GroupOpenHours(tt.openHours)
			if len(result) != len(tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, result)
			}
			for i := range result {
				if result[i]["days"] != tt.expected[i]["days"] {
					t.Errorf("Expected %v, got %v", tt.expected[i]["days"], result[i]["days"])
				}
				if result[i]["open"] != tt.expected[i]["open"] {
					t.Errorf("Expected %v, got %v", tt.expected[i]["open"], result[i]["open"])
				}
				if result[i]["close"] != tt.expected[i]["close"] {
					t.Errorf("Expected %v, got %v", tt.expected[i]["close"], result[i]["close"])
				}
			}
		})
	}
}

func ExampleGroupOpenHours() {
	openHours := []map[string]string{
		{"day": "Monday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Tuesday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Wednesday", "open": "8:00 AM", "close": "6:00 PM"},
		{"day": "Thursday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Friday", "open": "8:00 AM", "close": "5:00 PM"},
		{"day": "Saturday", "open": "8:00 AM", "close": "4:00 PM"},
	}

	result := GroupOpenHours(openHours)
	for _, r := range result {
		if r["open"] == "" && r["close"] == "" {
			fmt.Printf("%s\n", r["days"])
		} else {
			fmt.Printf("%s %s %s\n", r["days"], r["open"], r["close"])
		}
	}
	// Output:
	// Monday-Tuesday 8:00 AM 5:00 PM
	// Wednesday 8:00 AM 6:00 PM
	// Thursday-Friday 8:00 AM 5:00 PM
	// Saturday 8:00 AM 4:00 PM
	// Sunday
}
