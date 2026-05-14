package main

import "testing"

func TestSeparateDigits(t *testing.T) {
    tests := []struct {
        name  string
        input []int
        output []int
    }{
        {
            name: "Test Case 1",
            input: []int{13, 25, 83, 77},
            output: []int{1, 3, 2, 5, 8, 3, 7, 7},
        },
        {
            name: "Test Case 2",
            input: []int{7, 1, 3, 9},
            output: []int{7, 1, 3, 9},
        },
        {
            name: "Test Case 3",
            input: []int{10921},
            output: []int{1, 0, 9, 2, 1},
        },
        {
            name: "Test Case 4",
            input: []int{111, 222, 333},
            output: []int{1, 1, 1, 2, 2, 2, 3, 3, 3},
        },
        {
            name: "Test Case 5",
            input: []int{9, 99, 999, 9999},
            output: []int{9, 9, 9, 9, 9, 9, 9, 9, 9, 9},
        },
        {
            name: "Test Case 6",
            input: []int{3},
            output: []int{3},
        },
        {
            name: "Test Case 7",
            input: []int{56, 14, 87, 90},
            output: []int{5, 6, 1, 4, 8, 7, 9, 0},
        },
        {
            name: "Test Case 8",
            input: []int{1000, 234, 8765},
            output: []int{1, 0, 0, 0, 2, 3, 4, 8, 7, 6, 5},
        },
        {
            name: "Test Case 9",
            input: []int{102, 56, 88, 97},
            output: []int{1, 0, 2, 5, 6, 8, 8, 9, 7},
        },
        {
            name: "Test Case 10",
            input: []int{10201, 45},
            output: []int{1, 0, 2, 0, 1, 4, 5},
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            got := separateDigits(tt.input)
            if !equal(got, tt.output) {
                t.Errorf("separateDigits() = %v, want %v", got, tt.output)
            }
        })
    }
}

func equal(a, b []int) bool {
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