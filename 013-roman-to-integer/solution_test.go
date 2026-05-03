package main

import "testing"

func TestRomanToInt(t *testing.T){
	tests := []struct{
		Name string
		Value string
		ExpectedResult int
	}{
		{
			Name: "Example 1",
			Value: "III",
			ExpectedResult: 3,
		},
		{
			Name: "Example 2",
			Value: "LVIII",
			ExpectedResult: 58,
		},
		{
			Name: "Example 3",
			Value: "MCMXCIV",
			ExpectedResult: 1994,
		},
	}

	for _, test := range tests{
		got := romanToInt(test.Value)
		if got != test.ExpectedResult{
			t.Fatalf("Expected %v but got %v for %s", test.ExpectedResult, got, test.Value)
		}
	}
}