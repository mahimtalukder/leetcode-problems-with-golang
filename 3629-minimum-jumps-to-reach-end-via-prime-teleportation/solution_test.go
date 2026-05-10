package main

import "testing"

func TestMinJumps(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "Example 1 - teleport by 2",
			input:    []int{1, 2, 4, 6},
			expected: 2,
		},
		{
			name:     "Example 2 - teleport by 3",
			input:    []int{2, 3, 4, 7, 9},
			expected: 2,
		},
		{
			name:     "Example 3 - no useful teleport",
			input:    []int{4, 6, 5, 8},
			expected: 3,
		},
		{
			name:     "Single element",
			input:    []int{1},
			expected: 0,
		},
		{
			name:     "Two elements adjacent only",
			input:    []int{1, 2},
			expected: 1,
		},
		{
			name:     "Direct prime teleport",
			input:    []int{2, 4},
			expected: 1,
		},
		{
			name:     "Direct prime teleport by 7",
			input:    []int{7, 14},
			expected: 1,
		},
		{
			name:     "Prime teleport skips middle",
			input:    []int{7, 1, 1, 1, 14},
			expected: 1,
		},
		{
			name:     "No prime start, adjacent only",
			input:    []int{6, 10, 15},
			expected: 2,
		},
		{
			name:     "All powers of 2",
			input:    []int{2, 4, 8, 16, 32},
			expected: 1,
		},
		{
			name:     "All multiples of 3",
			input:    []int{3, 6, 9, 12, 15},
			expected: 1,
		},
		{
			name:     "All multiples of 5",
			input:    []int{5, 10, 15, 20, 25},
			expected: 1,
		},
		{
			name:     "All multiples of 11",
			input:    []int{11, 22, 33, 44, 55},
			expected: 1,
		},
		{
			name:     "Prime in middle can teleport",
			input:    []int{4, 2, 9},
			expected: 2,
		},
		{
			name:     "Prime in middle but target not divisible",
			input:    []int{8, 2, 4, 9},
			expected: 3,
		},
		{
			name:     "Move then teleport by 5",
			input:    []int{6, 5, 10, 15, 30},
			expected: 2,
		},
		{
			name:     "No teleport possible",
			input:    []int{1, 4, 6, 8, 10},
			expected: 4,
		},
		{
			name:     "Large prime direct teleport",
			input:    []int{13, 26, 39, 52, 65, 78},
			expected: 1,
		},
		{
			name:     "Multiple primes",
			input:    []int{2, 9, 3, 6, 15},
			expected: 2,
		},
		{
			name:     "Given testcase with answer 5",
			input:    []int{7, 17, 19, 23, 39, 11, 29, 31, 13, 14, 47, 53, 59, 61, 55},
			expected: 5,
		},
		{
			name:     "Given big testcase",
			input:    []int{893, 786, 607, 137, 69, 381, 790, 233, 15, 42, 7, 764, 890, 269, 84, 262, 870, 514, 514, 650, 269, 485, 760, 181, 489, 107, 585, 428, 862, 563},
			expected: 21,
		},
		{
			name:     "All ones",
			input:    []int{1, 1, 1, 1},
			expected: 3,
		},
		{
			name:     "Prime start but no useful divisor",
			input:    []int{2, 1, 1, 1, 1},
			expected: 4,
		},
		{
			name:     "Prime 17 with many multiples",
			input:    []int{17, 34, 51, 68, 85, 102, 119},
			expected: 1,
		},
		{
			name:     "Teleport from index 1",
			input:    []int{10, 7, 20, 35, 49, 77},
			expected: 2,
		},
		{
			name:     "Need move then teleport",
			input:    []int{2, 6, 5, 10, 25},
			expected: 2,
		},
		{
			name:     "Composite only",
			input:    []int{15, 21, 35, 77, 143},
			expected: 4,
		},
		{
			name:     "Large prime direct",
			input:    []int{97, 194, 291, 388, 485},
			expected: 1,
		},
		{
			name:     "Prime 11 near start",
			input:    []int{9, 11, 22, 33, 44, 55},
			expected: 2,
		},
		{
			name:     "Max value divisible by 2",
			input:    []int{2, 1000000},
			expected: 1,
		},
		{
			name:     "Same large prime duplicate",
			input:    []int{999983, 999983},
			expected: 1,
		},
		{
			name:     "Move then same large prime duplicate",
			input:    []int{1, 999983, 999983},
			expected: 2,
		},
		{
			name:     "Squares composite only",
			input:    []int{4, 9, 25, 49, 121},
			expected: 4,
		},
		{
			name:     "All primes but no teleport target",
			input:    []int{2, 3, 5, 7, 11, 13},
			expected: 5,
		},
		{
			name:     "Mixed many teleports",
			input:    []int{14, 2, 28, 3, 42, 5, 70},
			expected: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := minJumps(tt.input)

			if result != tt.expected {
				t.Errorf("FAILED %s\nInput: %v\nExpected: %v\nGot: %v",
					tt.name, tt.input, tt.expected, result)
			}
		})
	}
}
