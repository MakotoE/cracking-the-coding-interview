package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// isUnique returns true if all characters are unique.
func isUnique(s string) bool {
	var chars []int
	for _, c := range s {
		chars = append(chars, int(c))
	}
	sort.Ints(chars)

	if len(chars) == 0 {
		return true
	}

	last := chars[0]
	for _, c := range chars[1:] {
		if c == last {
			return false
		}
		last = c
	}

	return true
}

func TestIsUnique(t *testing.T) {
	tests := []struct {
		s        string
		expected bool
	}{
		{
			"",
			true,
		},
		{
			"a",
			true,
		},
		{
			"ab",
			true,
		},
		{
			"aa",
			false,
		},
		{
			"aba",
			false,
		},
		{
			"abca",
			false,
		},
		{
			"abcA",
			true,
		},
	}

	for i, test := range tests {
		assert.Equal(t, test.expected, isUnique(test.s), i)
	}
}
