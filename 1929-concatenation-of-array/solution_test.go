package main

import (
	"reflect"
	"testing"
)

func generateArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i + 1
	}
	return arr
}

func TestGetConcatenation(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected []int
	}{
		{
			name:     "Example 1",
			input:    []int{1, 2, 1},
			expected: []int{1, 2, 1, 1, 2, 1},
		},
		{
			name:     "Example 2",
			input:    []int{1, 3, 2, 1},
			expected: []int{1, 3, 2, 1, 1, 3, 2, 1},
		},
		{
			name:     "Single Element",
			input:    []int{5},
			expected: []int{5, 5},
		},
		{
			name:     "All Same Values",
			input:    []int{7, 7, 7},
			expected: []int{7, 7, 7, 7, 7, 7},
		},
		{
			name:     "Increasing Order",
			input:    []int{1, 2, 3, 4, 5},
			expected: []int{1, 2, 3, 4, 5, 1, 2, 3, 4, 5},
		},
		{
			name:     "Decreasing Order",
			input:    []int{5, 4, 3, 2, 1},
			expected: []int{5, 4, 3, 2, 1, 5, 4, 3, 2, 1},
		},
		{
			name:     "Mixed Values",
			input:    []int{10, 20, 30},
			expected: []int{10, 20, 30, 10, 20, 30},
		},
		{
			name:     "Max Constraint Size (small simulation)",
			input:    generateArray(10),
			expected: append(generateArray(10), generateArray(10)...),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := getConcatenation(tt.input)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FAILED %s\nExpected: %v\nGot: %v",
					tt.name, tt.expected, result)
			}
		})
	}
}
