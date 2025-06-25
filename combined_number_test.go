package main

import (
	"testing"
)

func TestCombinedNumber(t *testing.T) {
	tests := []struct {
		input    []int
		expected string
	}{
		{[]int{50, 2, 1, 9}, "95021"},
		{[]int{5, 50, 56}, "56550"},
		{[]int{420, 42, 423}, "42423420"},
		{[]int{0, 0}, "0"},
		{[]int{}, ""},
	}

	for _, tc := range tests {
		result := CombinedNumber(tc.input)
		if result != tc.expected {
			t.Errorf("For input %v: expected %s, got %s", tc.input, tc.expected, result)
		}
	}
}
