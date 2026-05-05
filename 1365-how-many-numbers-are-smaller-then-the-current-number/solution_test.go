package main

import (
	"reflect"
	"testing"
)

func generateTestArray(n int) []int {
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		arr[i] = i % 10 // repeating values 0–9
	}
	return arr
}

func expectedForGenerated(n int) []int {
	arr := generateTestArray(n)
	result := make([]int, n)

	for i := range arr {
		count := 0
		for j := range arr {
			if arr[j] < arr[i] {
				count++
			}
		}
		result[i] = count
	}
	return result
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected []int
	}{
		// Given examples
		{"Example 1", []int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{"Example 2", []int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{"Example 3", []int{7, 7, 7, 7}, []int{0, 0, 0, 0}},

		// Edge cases
		{"Minimum size", []int{0, 1}, []int{0, 1}},
		{"Two equal elements", []int{5, 5}, []int{0, 0}},

		// All same values
		{"All zeros", []int{0, 0, 0}, []int{0, 0, 0}},
		{"All max values", []int{100, 100, 100}, []int{0, 0, 0}},

		// Increasing order
		{"Strictly increasing", []int{1, 2, 3, 4, 5}, []int{0, 1, 2, 3, 4}},

		// Decreasing order
		{"Strictly decreasing", []int{5, 4, 3, 2, 1}, []int{4, 3, 2, 1, 0}},

		// Mixed with duplicates
		{"Mixed duplicates 1", []int{3, 3, 1, 2}, []int{2, 2, 0, 1}},
		{"Mixed duplicates 2", []int{4, 2, 2, 1}, []int{3, 1, 1, 0}},

		// Zeros involved
		{"With zeros", []int{0, 1, 2, 0}, []int{0, 2, 3, 0}},

		// Random patterns
		{"Random 1", []int{9, 1, 5, 3, 7}, []int{4, 0, 2, 1, 3}},
		{"Random 2", []int{10, 0, 5, 5, 2}, []int{4, 0, 2, 2, 1}},

		// Boundary values
		{"Includes 0 and 100", []int{0, 100, 50}, []int{0, 2, 1}},

		// Larger case
		{"Larger case", []int{1, 3, 2, 5, 4, 2, 1}, []int{0, 4, 2, 6, 5, 2, 0}},

		// Repeating pattern
		{"Repeating pattern", []int{1, 2, 1, 2, 1, 2}, []int{0, 3, 0, 3, 0, 3}},

		// Max size simulation (small version)
		{"Large input", generateTestArray(100), expectedForGenerated(100)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := smallerNumbersThanCurrent(tt.nums)

			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("FAILED %s\nInput: %v\nExpected: %v\nGot: %v",
					tt.name,tt.nums , tt.expected, result)
			}
		})
	}
}