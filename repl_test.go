package main

import (
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:	  "   hello, world!    ",
			expected: []string{"hello,", "world!"},
		},
		{
			input:    "Capitalized Words",
			expected: []string{"capitalized", "words"},
		},
		{
			input:    "Multiple   Spaces",
			expected: []string{"multiple", "spaces"},
		},
		{
			input:    "Special Characters! @#$%^&*()",
			expected: []string{"special", "characters!", "@#$%^&*()"},
		},
		{
			input:    "Mixed CASE and Numbers 123",
			expected: []string{"mixed", "case", "and", "numbers", "123"},
		},
	}
	for _, c := range cases {
		actual := cleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("For input '%s', expected length %d but got %d", c.input, len(c.expected), len(actual))
			continue
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("For input '%s', expected word '%s' but got '%s'", c.input, expectedWord, word)
			}
		}
	}
}
