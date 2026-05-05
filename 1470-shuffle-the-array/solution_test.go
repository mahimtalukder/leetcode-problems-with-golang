package main

import (
	"reflect"
	"testing"
)

func TestShuffle(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		n        int
		expected []int
	}{
		{
			name:     "Example 1",
			nums:     []int{2, 5, 1, 3, 4, 7},
			n:        3,
			expected: []int{2, 3, 5, 4, 1, 7},
		},
		{
			name:     "Example 2",
			nums:     []int{1, 2, 3, 4, 4, 3, 2, 1},
			n:        4,
			expected: []int{1, 4, 2, 3, 3, 2, 4, 1},
		},
		{
			name:     "Example 3",
			nums:     []int{1, 1, 2, 2},
			n:        2,
			expected: []int{1, 2, 1, 2},
		},
		{
			name:     "Minimum n",
			nums:     []int{10, 20},
			n:        1,
			expected: []int{10, 20},
		},
		{
			name:     "All same values",
			nums:     []int{7, 7, 7, 7, 7, 7},
			n:        3,
			expected: []int{7, 7, 7, 7, 7, 7},
		},
		{
			name:     "Increasing numbers",
			nums:     []int{1, 2, 3, 4, 5, 6},
			n:        3,
			expected: []int{1, 4, 2, 5, 3, 6},
		},
		{
			name:     "Two pairs",
			nums:     []int{100, 200, 300, 400},
			n:        2,
			expected: []int{100, 300, 200, 400},
		},
		{
			name:     "With max constraint values",
			nums:     []int{1000, 999, 1, 2},
			n:        2,
			expected: []int{1000, 1, 999, 2},
		},
		{
			name:     "Repeated pattern",
			nums:     []int{1, 2, 1, 2, 3, 4, 3, 4},
			n:        4,
			expected: []int{1, 3, 2, 4, 1, 3, 2, 4},
		},
		{
			name:     "Larger case",
			nums:     []int{11, 22, 33, 44, 55, 66, 77, 88, 99, 100},
			n:        5,
			expected: []int{11, 66, 22, 77, 33, 88, 44, 99, 55, 100},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := shuffle(tt.nums, tt.n)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("expected %v, got %v", tt.expected, got)
			}
		})
	}
}
