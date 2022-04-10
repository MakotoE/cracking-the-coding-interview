package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"strconv"
	"strings"
	"testing"
)

// compressString returns the compressed form of the given string.
func compressString(s string) string {
	/*
		Track last char and its index
		Count number of equal chars
		Replace duplicates with the count
	*/
	if len(s) == 0 {
		return ""
	}

	type CharacterInfo struct {
		character rune
		count     int
	}

	chars := []CharacterInfo{
		{[]rune(s)[0], 1},
	}

	for _, c := range s[1:] {
		if chars[len(chars)-1].character == c {
			chars[len(chars)-1].count += 1
		} else {
			chars = append(chars, CharacterInfo{c, 1})
		}
	}

	result := strings.Builder{}
	for _, info := range chars {
		result.WriteRune(info.character)
		result.WriteString(strconv.Itoa(info.count))
	}

	return result.String()
}

func TestCompressString(t *testing.T) {
	tests := []struct {
		s        string
		expected string
	}{
		{
			"",
			"",
		},
		{
			"a",
			"a1",
		},
		{
			"ab",
			"a1b1",
		},
		{
			"aa",
			"a2",
		},
		{
			"aab",
			"a2b1",
		},
		{
			"baa",
			"b1a2",
		},
		{
			"aba",
			"a1b1a1",
		},
		{
			"abc",
			"a1b1c1",
		},
		{
			"aabcccccaaa",
			"a2b1c5a3",
		},
		{
			"1",
			"11",
		},
	}

	for i, test := range tests {
		assert.Equal(t, test.expected, compressString(test.s), i)
	}
}
