package main

import "testing"

func generateLargeOnesTest(n int) []int {
	arr := make([]int, n)
	for i := range arr {
		arr[i] = 1
	}
	return arr
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		expected int
	}{
		// Basic examples
		{"Example 1", []int{1, 1, 0, 1, 1, 1}, 3},
		{"Example 2", []int{1, 0, 1, 1, 0, 1}, 2},

		// Edge cases
		{"Single element 1", []int{1}, 1},
		{"Single element 0", []int{0}, 0},

		// All same
		{"All ones", []int{1, 1, 1, 1, 1}, 5},
		{"All zeros", []int{0, 0, 0, 0}, 0},

		// Alternating pattern
		{"Alternating 1 and 0", []int{1, 0, 1, 0, 1, 0}, 1},
		{"Alternating starting with 0", []int{0, 1, 0, 1, 0, 1}, 1},

		// Ones at boundaries
		{"Ones at start", []int{1, 1, 1, 0, 0, 0}, 3},
		{"Ones at end", []int{0, 0, 0, 1, 1, 1}, 3},

		// Multiple segments
		{"Multiple segments", []int{1, 1, 0, 1, 1, 1, 0, 1}, 3},
		{"Equal segments", []int{1, 1, 0, 1, 1}, 2},

		// Complex patterns
		{"Mixed pattern 1", []int{0, 1, 1, 0, 1, 1, 1, 0, 1, 1}, 3},
		{"Mixed pattern 2", []int{1, 0, 1, 1, 1, 0, 1, 1, 0, 1}, 3},

		// Large continuous ones
		{"Long streak in middle", []int{0, 1, 1, 1, 1, 1, 1, 0}, 6},

		// No ones
		{"No ones at all", []int{0, 0, 0, 0, 0, 0}, 0},

		// Max constraint simulation (small version)
		{"Large input", generateLargeOnesTest(1000), 1000},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := findMaxConsecutiveOnes(tt.nums)

			if result != tt.expected {
				t.Errorf("FAILED %s: expected %d, got %d",
					tt.name, tt.expected, result)
			}
		})
	}
}