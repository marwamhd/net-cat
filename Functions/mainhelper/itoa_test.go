package mainhelper

import (
	"testing"
)

type itoaCase struct {
	input    int
	expected string
}

func TestItoa(t *testing.T) {
	// Test cases
	testCases := []itoaCase{
		{0, "0"},
		{123, "123"},
		{-456, "-456"},
		{789, "789"},
		{1234567890, "1234567890"},
		{-1, "-1"},
		{1, "1"},
		{-9999, "-9999"},
		{9999, "9999"},
		{-2147483648, "-2147483648"},
		{2147483647, "2147483647"},
		{2147483648, "2147483648"},
		{9223372036854775807, "9223372036854775807"},
		{-9223372036854775807, "-9223372036854775807"},
	}

	for _, tc := range testCases {
		result := Itoa(tc.input)
		if result != tc.expected {
			t.Errorf("Itoa(%d) = %s, expected %s", tc.input, result, tc.expected)
		}
	}
}
