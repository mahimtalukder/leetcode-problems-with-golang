package main

import "testing"

func repeatString(s string, count int) string {
	result := ""
	for i := 0; i < count; i++ {
		result += s
	}
	return result
}

func TestRotateString(t *testing.T) {
	tests := []struct {
		S              string
		Goal           string
		ExpectedResult bool
	}{
		// 🔹 Basic cases
		{"abcde", "cdeab", true},
		{"abcde", "abced", false},

		// 🔹 Same string
		{"a", "a", true},
		{"aaaa", "aaaa", true},

		// 🔹 Single char mismatch
		{"a", "b", false},

		// 🔹 Two characters
		{"ab", "ba", true},
		{"ab", "ab", true},
		{"ab", "aa", false},

		// 🔹 Repeated characters
		{"aaab", "abaa", true},
		{"aaab", "aaba", true},
		{"aaab", "baaa", true},
		{"aaab", "abaa", true},
		{"aaab", "bbaa", false},

		// 🔹 Length mismatch
		{"abc", "abcd", false},
		{"abcd", "abc", false},

		// 🔹 Mid-size strings
		{"rotationtest", "testrotation", true},
		{"rotationtest", "tationtestro", true},
		{"rotationtest", "rotationtset", false},

		// 🔹 Complex patterns
		{"abababab", "babababa", true},
		{"abababab", "abababba", false},

		// 🔹 Palindrome-like
		{"racecar", "carrace", true},
		{"racecar", "racecra", false},

		// 🔹 Large strings
		{
			S:              "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			Goal:           "mnopqrstuvwxyzabcdefghijklmnopqrstuvwxyzabcdefghijkl",
			ExpectedResult: true,
		},
		{
			S:              "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxyz",
			Goal:           "abcdefghijklmnopqrstuvwxyzabcdefghijklmnopqrstuvwxzy",
			ExpectedResult: false,
		},

		// 🔹 Very large repeated string
		{
			S:              repeatString("a", 100),
			Goal:           repeatString("a", 100),
			ExpectedResult: true,
		},
		{
			S:              repeatString("a", 99) + "b",
			Goal:           "b" + repeatString("a", 99),
			ExpectedResult: true,
		},
		{
			S:              repeatString("a", 99) + "b",
			Goal:           repeatString("a", 100),
			ExpectedResult: false,
		},

		// 🔹 Edge tricky cases
		{"abcabc", "bcabca", true},
		{"abcabc", "cababc", false},
	}

	for _, test := range tests {
		got := rotateString(test.S, test.Goal)

		if test.ExpectedResult != got {
			t.Errorf("Expected %v but got %v for s=%s and goal=%s",
				test.ExpectedResult, got, test.S, test.Goal)
		}
	}
}
