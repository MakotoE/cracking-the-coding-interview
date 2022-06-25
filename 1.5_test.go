package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// OneAway returns true if b is one edit away from a.
func OneAway(a string, b string) bool {
	diff := len(a) - len(b)
	if diff != 1 && diff != 0 && diff != -1 {
		return false
	}

	editFound := false
	aIndex := 0
	bIndex := 0

	for aIndex < len(a) && bIndex < len(b) {
		if a[aIndex] != b[bIndex] {
			if editFound {
				return false
			}

			editFound = true

			if aIndex < len(a)-1 && a[aIndex+1] == b[bIndex] {
				aIndex++
			} else if bIndex < len(b)-1 && a[aIndex] == b[bIndex+1] {
				bIndex++
			} else if aIndex < len(a)-1 && bIndex < len(b)-1 && a[aIndex+1] == b[bIndex+1] {
				aIndex++
				bIndex++
			} else if (aIndex == len(a)-2 && bIndex <= len(b)-1) || (aIndex <= len(a)-1 && bIndex == len(b)-2) {
				return false
			}
		}

		aIndex++
		bIndex++
	}

	return true
}

func TestOneAway(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected bool
	}{
		{
			"",
			"",
			true,
		},
		{
			"a",
			"",
			true,
		},
		{
			"a",
			"a",
			true,
		},
		{
			"a",
			"b",
			true,
		},
		{
			"aa",
			"a",
			true,
		},
		{
			"ab",
			"aa",
			true,
		},
		{
			"ba",
			"aa",
			true,
		},
		{
			"bb",
			"aa",
			false,
		},
		{
			"aa",
			"",
			false,
		},
		{
			"b",
			"aa",
			false,
		},
		{
			"aaa",
			"a",
			false,
		},
		{
			"b",
			"aaa",
			false,
		},
		{
			"bbb",
			"a",
			false,
		},
	}

	for i, test := range tests {
		{
			result := OneAway(test.a, test.b)
			assert.Equal(t, test.expected, result, i)
		}

		// Test symmetric relation
		{
			result := OneAway(test.b, test.a)
			assert.Equal(t, test.expected, result, i)
		}
	}
}
