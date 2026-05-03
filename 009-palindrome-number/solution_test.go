package main

import "testing"

func TestIsPalindrome(t *testing.T){
	tests := []struct {
		Name string
		Value int
		ExpectedResult bool
	}{
		{
			Name: "Example 1",
			Value: 121,
			ExpectedResult: true,
		},
		{
			Name: "Example 2",
			Value: -121,
			ExpectedResult: false,
		},
		{
			Name: "Example 3",
			Value: 10,
			ExpectedResult: false,
		},
	}

	for _, test := range tests{
		got := isPalindrome(test.Value)
		if test.ExpectedResult != got {
			t.Fatalf("Expected %v But Got %v for test name %s", test.ExpectedResult, got, test.Name)
		}
	}
}