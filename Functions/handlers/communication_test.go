package handlers

import "testing"

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
		{[]byte{32, 32, 32, 10}, true},
		{[]byte{32, 32, 32, 65, 10}, false},
		{[]byte{65, 10}, false},
	}

	for _, tc := range testCases {
		result := IsEmpty(tc.message)
		if result != tc.expected {
			t.Errorf("IsEmpty(%v) = %t, expected %t", tc.message, result, tc.expected)
		}
	}
}
