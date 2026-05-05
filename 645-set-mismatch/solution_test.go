package main

import (
	"reflect"
	"testing"
)

func generateSetMismatch(n, duplicate, missing int) []int {
	arr := make([]int, n)
	idx := 0

	for i := 1; i <= n; i++ {
		if i == missing {
			continue
		}
		arr[idx] = i
		idx++
	}

	// replace one value with duplicate
	arr[idx-1] = duplicate

	return arr
}

func TestFindErrorNums(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		// Basic examples
		{"Example 1", []int{1, 2, 2, 4}, []int{2, 3}},
		{"Example 2", []int{1, 1}, []int{1, 2}},

		// Small edge cases
		{"Small n=2 (duplicate 2)", []int{2, 2}, []int{2, 1}},
		{"Small n=2 (duplicate 1)", []int{1, 1}, []int{1, 2}},

		// Duplicate in middle
		{"Duplicate in middle", []int{1, 3, 3, 4, 5}, []int{3, 2}},

		// Duplicate at start
		{"Duplicate at start", []int{1, 1, 3, 4, 5}, []int{1, 2}},

		// Duplicate at end
		{"Duplicate at end", []int{1, 2, 3, 4, 4}, []int{4, 5}},

		// Missing first number
		{"Missing 1", []int{2, 2, 3, 4, 5}, []int{2, 1}},

		// Missing last number
		{"Missing n", []int{1, 2, 3, 4, 4}, []int{4, 5}},

		// Random order
		{"Unsorted input", []int{3, 1, 2, 5, 3}, []int{3, 4}},

		// Larger array
		{"Larger case", []int{1, 2, 3, 4, 5, 6, 6, 8, 9, 10}, []int{6, 7}},

		// Multiple duplicates but only one valid pair
		{"Single valid duplicate", []int{4, 3, 6, 2, 1, 1}, []int{1, 5}},

		// Consecutive duplicate
		{"Consecutive duplicate", []int{1, 2, 3, 3, 5, 6}, []int{3, 4}},

		// All elements shifted
		{"Shifted pattern", []int{2, 3, 4, 5, 6, 6}, []int{6, 1}},

		// Max value duplication
		{"Max value duplicate", []int{1, 2, 3, 4, 5, 5}, []int{5, 6}},

		// Min value duplication
		{"Min value duplicate", []int{1, 1, 2, 3, 4, 5}, []int{1, 6}},

		// Larger simulated input
		{"Large input", generateSetMismatch(1000, 500, 700), []int{500, 700}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findErrorNums(tt.nums)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FAILED %s: expected %v, got %v",
					tt.name, tt.expected, result)
			}
		})
	}
}
