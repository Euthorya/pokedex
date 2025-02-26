package main

import (
	"fmt"
	"testing"
	//"fmt"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
		{
			input:    "",
			expected: []string{},
		},
		{
			input:    "hello  world  DFDSFSD       ",
			expected: []string{"hello", "world", "dfdsfsd"},
		},
	}

	for _, c := range cases {
		actual := cleanInput(c.input)

		if len(actual) != len(c.expected) {
			t.Errorf("Length mismatch between %v and %v", actual, c.expected)
		}
		fmt.Println(actual)
		fmt.Println(c.expected)
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("Error, %v != %v", word, expectedWord)
			}
		}
	}
}
