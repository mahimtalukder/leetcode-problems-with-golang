package main

import "testing"

func TestMaximumJumps(t *testing.T) {
	tests := []struct {
		name     string
		nums     []int
		target   int
		expected int
	}{
		{
			name:     "example 1",
			nums:     []int{1, 3, 6, 4, 1, 2},
			target:   2,
			expected: 3,
		},
		{
			name:     "example 2",
			nums:     []int{1, 3, 6, 4, 1, 2},
			target:   3,
			expected: 5,
		},
		{
			name:     "example 3",
			nums:     []int{1, 3, 6, 4, 1, 2},
			target:   0,
			expected: -1,
		},
		{
			name:     "only two elements reachable",
			nums:     []int{1, 2},
			target:   1,
			expected: 1,
		},
		{
			name:     "only two elements not reachable",
			nums:     []int{1, 5},
			target:   2,
			expected: -1,
		},
		{
			name:     "all adjacent jumps possible",
			nums:     []int{1, 2, 3, 4, 5},
			target:   1,
			expected: 4,
		},
		{
			name:     "can jump directly but max jumps is larger",
			nums:     []int{1, 2, 3, 4},
			target:   10,
			expected: 3,
		},
		{
			name:     "strictly decreasing reachable",
			nums:     []int{5, 4, 3, 2, 1},
			target:   1,
			expected: 4,
		},
		{
			name:     "strictly decreasing not reachable",
			nums:     []int{10, 7, 4, 1},
			target:   2,
			expected: -1,
		},
		{
			name:     "all same numbers with zero target",
			nums:     []int{4, 4, 4, 4},
			target:   0,
			expected: 3,
		},
		{
			name:     "zero target but values differ",
			nums:     []int{4, 4, 5, 4},
			target:   0,
			expected: 2,
		},
		{
			name:     "negative numbers reachable",
			nums:     []int{-5, -4, -3, -2, -1},
			target:   1,
			expected: 4,
		},
		{
			name:     "mixed negative and positive reachable",
			nums:     []int{-3, -1, 1, 3},
			target:   2,
			expected: 3,
		},
		{
			name:     "mixed negative and positive not reachable",
			nums:     []int{-10, -5, 0, 5},
			target:   4,
			expected: -1,
		},
		{
			name:     "skip blocked middle values",
			nums:     []int{1, 100, 2, 3},
			target:   2,
			expected: 2,
		},
		{
			name:     "last index unreachable although middle reachable",
			nums:     []int{1, 2, 3, 100},
			target:   2,
			expected: -1,
		},
		{
			name:     "multiple paths choose maximum jumps",
			nums:     []int{1, 2, 1, 2, 1},
			target:   1,
			expected: 4,
		},
		{
			name:     "large values reachable",
			nums:     []int{-1000000000, 0, 1000000000},
			target:   1000000000,
			expected: 2,
		},
		{
			name:     "large values not reachable",
			nums:     []int{-1000000000, 1000000000},
			target:   999999999,
			expected: -1,
		},
		{
			name:     "target maximum allows all jumps",
			nums:     []int{-1000000000, 500, -300, 1000000000},
			target:   2000000000,
			expected: 3,
		},
		{
			name:     "reachable only by direct jump",
			nums:     []int{5, 100, 200, 6},
			target:   1,
			expected: 1,
		},
		{
			name:     "cannot move from start",
			nums:     []int{1, 10, 20, 30},
			target:   5,
			expected: -1,
		},
		{
			name:     "duplicate values with zero target",
			nums:     []int{7, 8, 7, 7, 9, 7},
			target:   0,
			expected: 3,
		},
		{
			name:     "zigzag values all reachable",
			nums:     []int{1, 3, 2, 4, 3, 5},
			target:   2,
			expected: 5,
		},
		{
			name:     "zigzag values partially blocked",
			nums:     []int{1, 4, 2, 8, 3},
			target:   2,
			expected: 2,
		},
		{
			name:     "direct jump only because middle path blocked",
			nums:     []int{1, 0, 2},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump only with decreasing middle value",
			nums:     []int{5, 3, 6},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump only because middle to last invalid",
			nums:     []int{10, 8, 11},
			target:   1,
			expected: 1,
		},
		{
			name:     "middle reachable but cannot continue to last",
			nums:     []int{4, 3, 5},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump valid but two step path invalid",
			nums:     []int{2, 1, 3},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump only with negative numbers",
			nums:     []int{-1, -3, 0},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump only although first middle is reachable",
			nums:     []int{7, 6, 8},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump valid but middle value too far from last",
			nums:     []int{100, 99, 101},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump valid but all middle routes blocked",
			nums:     []int{1, 0, 5, 2},
			target:   1,
			expected: 1,
		},
		{
			name:     "direct jump valid after several invalid middle options",
			nums:     []int{-533985778,-424626669,794071124,694501105,-651162637,-789411200,773124493,-655591953,205086705,-421668572},
			target:   1194793065,
			expected: 6,
		},
		{
			name:     "leetcode test no 2304",
			nums:     []int{1, 0, 3, 4, 2},
			target:   2,
			expected: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maximumJumps(tt.nums, tt.target)

			if got != tt.expected {
				t.Errorf("maximumJumps(%v, %d) = %d, expected %d",
					tt.nums, tt.target, got, tt.expected)
			}
		})
	}
}
