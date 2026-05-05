package main

import (
	"reflect"
	"testing"
)

// helper: create linked list from slice
func buildList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	head := &ListNode{Val: nums[0]}
	current := head
	for _, v := range nums[1:] {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}
	return head
}

// helper: convert linked list to slice
func toSlice(head *ListNode) []int {
	result := []int{} // important: empty slice, not nil

	for head != nil {
		result = append(result, head.Val)
		head = head.Next
	}

	return result
}

func TestRotateRight(t *testing.T) {

	tests := []struct {
		name     string
		input    []int
		k        int
		expected []int
	}{
		// 🔹 basic cases
		{"Example 1", []int{1, 2, 3, 4, 5}, 2, []int{4, 5, 1, 2, 3}},
		{"Example 2", []int{0, 1, 2}, 4, []int{2, 0, 1}},

		// 🔹 edge cases
		{"Empty list", []int{}, 5, []int{}},
		{"Single node", []int{1}, 10, []int{1}},
		{"k = 0", []int{1, 2, 3}, 0, []int{1, 2, 3}},
		{"k = length", []int{1, 2, 3}, 3, []int{1, 2, 3}},

		// 🔹 k > length
		{"k greater than length", []int{1, 2, 3}, 4, []int{3, 1, 2}},
		{"k much larger", []int{1, 2, 3, 4}, 100, []int{1, 2, 3, 4}}, // 100 % 4 = 0

		// 🔹 small variations
		{"Two nodes rotate 1", []int{1, 2}, 1, []int{2, 1}},
		{"Two nodes rotate 2", []int{1, 2}, 2, []int{1, 2}},

		// 🔹 full rotations
		{"Rotate full cycle", []int{1, 2, 3, 4}, 8, []int{1, 2, 3, 4}},

		// 🔹 complex cases
		{"Rotate 1", []int{1, 2, 3, 4, 5}, 1, []int{5, 1, 2, 3, 4}},
		{"Rotate 3", []int{1, 2, 3, 4, 5}, 3, []int{3, 4, 5, 1, 2}},

		// 🔹 duplicates
		{"With duplicates", []int{1, 1, 2, 2, 3}, 2, []int{2, 3, 1, 1, 2}},

		// 🔹 negative values
		{"Negative values", []int{-1, -2, -3, -4}, 2, []int{-3, -4, -1, -2}},

		// 🔹 long list
		{"Long list", []int{1,2,3,4,5,6,7,8,9,10}, 3, []int{8,9,10,1,2,3,4,5,6,7}},

		// 🔹 max constraint style
		{"Large k edge", []int{1,2,3}, 2000000000, []int{2,3,1}}, // k % 3 = 2
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			head := buildList(tt.input)
			result := rotateRight(head, tt.k)
			got := toSlice(result)

			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("FAILED %s: expected %v, got %v", tt.name, tt.expected, got)
			}
		})
	}
}