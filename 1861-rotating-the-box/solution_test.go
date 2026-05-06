package main

import (
	"testing"
)

func TestRotatedBoxGrid(t *testing.T) {
	tests := []struct {
		name     string
		boxGrid  [][]byte
		expected [][]byte
	}{
		{
			name: "Test 1",
			boxGrid: [][]byte{
				{'#', '.', '#'},
			},
			expected: [][]byte{
				{'.'},
				{'#'},
				{'#'},
			},
		},
		{
			name: "Test 2",
			boxGrid: [][]byte{
				{'#', '.', '*', '.'},
				{'#', '#', '*', '.'},
			},
			expected: [][]byte{
				{'#', '.'},
				{'#', '#'},
				{'*', '*'},
				{'.', '.'},
			},
		},
		{
			name: "Test 3",
			boxGrid: [][]byte{
				{'#', '#', '*', '.', '*', '.'},
				{'#', '#', '#', '*', '.', '.'},
				{'#', '#', '#', '.', '#', '.'},
			},
			expected: [][]byte{
				{'.', '#', '#'},
				{'.', '#', '#'},
				{'#', '#', '*'},
				{'#', '*', '.'},
				{'#', '.', '*'},
				{'#', '.', '.'},
			},
		},
		{
			name: "Test 4 (Edge Case - 1x1 grid)",
			boxGrid: [][]byte{
				{'#'},
			},
			expected: [][]byte{
				{'#'},
			},
		},
		{
			name: "Test 5 (All stones)",
			boxGrid: [][]byte{
				{'#', '#', '#'},
				{'#', '#', '#'},
				{'#', '#', '#'},
			},
			expected: [][]byte{
				{'#', '#', '#'},
				{'#', '#', '#'},
				{'#', '#', '#'},
			},
		},
		{
			name: "Test 6 (All obstacles)",
			boxGrid: [][]byte{
				{'*', '*', '*'},
				{'*', '*', '*'},
				{'*', '*', '*'},
			},
			expected: [][]byte{
				{'*', '*', '*'},
				{'*', '*', '*'},
				{'*', '*', '*'},
			},
		},
		{
			name: "Test 7 (Empty grid)",
			boxGrid: [][]byte{
				{'.', '.', '.'},
				{'.', '.', '.'},
				{'.', '.', '.'},
			},
			expected: [][]byte{
				{'.', '.', '.'},
				{'.', '.', '.'},
				{'.', '.', '.'},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Call the function with boxGrid as input
			got := rotateTheBox(tt.boxGrid)
			// Check if the result matches the expected value
			for i := 0; i < len(got); i++ {
				for j := 0; j < len(got[i]); j++ {
					if got[i][j] != tt.expected[i][j] {
						t.Errorf("rotateTheBox() = %v, want %v", got, tt.expected)
					}
				}
			}
		})
	}
}
