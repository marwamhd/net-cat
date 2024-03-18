package mainhelper

import (
	"testing"
)

type atoiCase struct {
	input    string
	expected int
}

func TestAtoi(t *testing.T) {
	// Test cases
	testCases := []atoiCase{
		// Valid inputs
		{"123", 123},
		{"-456", -456},
		{"+789", 789},
		{"1234567890", 1234567890},
		{"0", 0},
		{"-1", -1},                   // negative single digit
		{"+1", 1},                    // positive single digit
		{"-9999", -9999},             // negative four digits
		{"+9999", 9999},              // positive four digits
		{"-2147483648", -2147483648}, // minimum int32 value
		{"2147483647", 2147483647},   // maximum int32 value
		{"2147483648", 2147483648},   // maximum int32 value
		{"9223372036854775807", 9223372036854775807},   // maximum int64 value
		{"-9223372036854775808", -9223372036854775808}, // minimum int64 value

		// Invalid inputs
		{"abc", 0},   // non-digit character
		{"12a34", 0}, // non-digit character
		{"-0", 0},    // negative zero value
		{"", 0},      // empty input
		{"-", 0},     // sign input
		{"+", 0},     // sign input
		{"+0", 0},    // positive zero value

		// {"9223372036854775807", 9223372036854775807},   // overflow
		// {"-9223372036854775809", -9223372036854775808}, // undeflow
	}

	for _, tc := range testCases {
		result, _ := Atoi(tc.input)
		if result != tc.expected {
			t.Errorf("Atoi(%s) = %d, expected %d", tc.input, result, tc.expected)
		}
	}
}
