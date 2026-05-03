package main

import (
	"testing"
)

func TestRotatedDigits(t *testing.T) {
	tests := []struct {
		Value          int
		ExpectedResult int
	}{
		{
			Value:          10,
			ExpectedResult: 4,
		},
		{
			Value:          1,
			ExpectedResult: 0,
		},
		{
			Value:          78,
			ExpectedResult: 29,
		},
	}

	for _, test := range tests {
		got := rotatedDigits(test.Value)
		if got != test.ExpectedResult {
			t.Errorf("Expected %v but got %v for range %v", test.ExpectedResult, got, test.Value)
		}
	}
}
