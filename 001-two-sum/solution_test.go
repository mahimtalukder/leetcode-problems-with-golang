package main

import "testing"

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{
			name:   "basic example",
			nums:   []int{2, 7, 11, 15},
			target: 9,
			want:   []int{0, 1},
		},
		{
			name:   "duplicate values",
			nums:   []int{3, 3},
			target: 6,
			want:   []int{0, 1},
		},
		{
			name:   "no solution",
			nums:   []int{1, 2, 3},
			target: 7,
			want:   []int{},
		},
	}

	for _, tt := range tests {
		got := twoSum(tt.nums, tt.target)
		if !sameSlice(got, tt.want) {
			t.Fatalf("%s: got %v, want %v", tt.name, got, tt.want)
		}
	}
}

func sameSlice(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}
