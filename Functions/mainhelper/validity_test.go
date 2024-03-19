package mainhelper

import (
	"testing"
)

func TestIsEmpty(t *testing.T) {
	testCases := []struct {
		message  []byte
		expected bool
	}{
		{[]byte{}, true},
		{[]byte{32}, true},
		{[]byte{32, 32, 32}, true},
		{[]byte{65, 66, 67}, false},
		{[]byte{32, 65, 66, 67}, false},
		{[]byte{32, 32, 32, 32}, true},
		{[]byte{32, 32, 32, 65, 10}, false},
		{[]byte{65, 10}, false},
		{[]byte{9, 10, 31, 12}, true},
	}

	for _, tc := range testCases {
		result := IsEmpty(tc.message)
		if result != tc.expected {
			t.Errorf("IsEmpty(%v) = %t, expected %t", tc.message, result, tc.expected)
		}
	}
}
func TestSignaltrapchecker(t *testing.T) {
	testCases := []struct {
		message  []byte
		expected bool
	}{
		{[]byte{27, 91, 65}, true},
		{[]byte{27, 91, 66}, true},
		{[]byte{27, 91, 68}, true},
		{[]byte{27, 91, 67}, true},
		{[]byte{27, 91, 72}, true},
		{[]byte{27, 91, 70}, true},
		{[]byte{27, 95, 65, 27, 91, 66}, true},
		{[]byte{27, 91, 65, 27, 91, 66, 27, 91, 68}, true},
		{[]byte{27, 91, 65, 27, 91, 66, 27, 91, 68, 27, 91, 67}, true},
		{[]byte{27, 96, 65, 27, 91, 66, 27, 91, 68, 27, 91, 67, 27, 91, 72}, true},
		{[]byte{27, 91, 65, 27, 91, 66, 27, 91, 68, 27, 91, 67, 27, 91, 72, 27, 91, 70}, true},
		{[]byte{27, 91, 65, 27, 91, 66, 27, 91, 68, 27, 91, 67, 27, 91, 72, 27, 91, 70, 65}, true},
		{[]byte{27, 93, 65, 27, 91, 66, 27, 91, 68, 27, 91, 67, 27, 91, 72, 27, 91, 70, 65, 66}, true},
		{[]byte{27, 96, 65, 27, 91, 66, 27, 91, 68, 27, 91, 67, 27, 91, 72, 27, 91, 70, 65, 66, 67}, true},
		{[]byte{65, 66, 65}, false},
	}

	for _, tc := range testCases {
		result := Signaltrapchecker(tc.message)
		if result != tc.expected {
			t.Errorf("Signaltrapchecker(%v) = %t, expected %t", tc.message, result, tc.expected)
		}
	}
}
