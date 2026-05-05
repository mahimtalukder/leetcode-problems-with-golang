package main

import (
	"reflect"
	"testing"
)

func TestRotate(t *testing.T) {
	tests := []struct {
		name     string
		matrix   [][]int
		expected [][]int
	}{
		{
			name: "3x3 matrix",
			matrix: [][]int{
				{1, 2, 3},
				{4, 5, 6},
				{7, 8, 9},
			},
			expected: [][]int{
				{7, 4, 1},
				{8, 5, 2},
				{9, 6, 3},
			},
		},
		{
			name: "4x4 matrix",
			matrix: [][]int{
				{5, 1, 9, 11},
				{2, 4, 8, 10},
				{13, 3, 6, 7},
				{15, 14, 12, 16},
			},
			expected: [][]int{
				{15, 13, 2, 5},
				{14, 3, 4, 1},
				{12, 6, 8, 9},
				{16, 7, 10, 11},
			},
		},
		{
			name: "1x1 matrix",
			matrix: [][]int{
				{1},
			},
			expected: [][]int{
				{1},
			},
		},
		{
			name: "5x5 matrix",
			matrix: [][]int{
				{1, 2, 3, 4, 5},
				{6, 7, 8, 9, 10},
				{11, 12, 13, 14, 15},
				{16, 17, 18, 19, 20},
				{21, 22, 23, 24, 25},
			},
			expected: [][]int{
				{21, 16, 11, 6, 1},
				{22, 17, 12, 7, 2},
				{23, 18, 13, 8, 3},
				{24, 19, 14, 9, 4},
				{25, 20, 15, 10, 5},
			},
		},
		{
			name: "symmetric matrix",
			matrix: [][]int{
				{1, 2, 1},
				{2, 3, 2},
				{1, 2, 1},
			},
			expected: [][]int{
				{1, 2, 1},
				{2, 3, 2},
				{1, 2, 1},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotate(tt.matrix)

			if !reflect.DeepEqual(tt.matrix, tt.expected) {
				t.Errorf("expected %v but got %v", tt.expected, tt.matrix)
			}
		})
	}
}
