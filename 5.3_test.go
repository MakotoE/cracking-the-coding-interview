package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// LongestOneSequenceAfterFlip returns the length of the longest sequence of ones after flipping any
// bit.
func LongestOneSequenceAfterFlip(n int32) int {
	currCount := 0
	lastCount := 0
	max := 0

	for i := 0; i < 32; i++ {
		isOne := (n & (1 << i)) != 0
		if isOne {
			currCount++
		} else {
			if lastCount+currCount > max {
				max = lastCount + currCount
			}

			lastCount = currCount
			currCount = 0
		}
	}

	return max + 1
}

func TestLongestOneSequenceAfterFlip(t *testing.T) {
	tests := []struct {
		n        int32
		expected int
	}{
		{
			0,
			1,
		},
		{
			1,
			2,
		},
		{
			3, // 11
			3,
		},
		{
			5, // 101
			3,
		},
		{
			9, // 1001
			2,
		},
		{
			25, // 11001
			3,
		},
		{
			19, // 10011
			3,
		},
		{
			101, // 1100101
			3,
		},
		{
			1775, // 11011101111
			8,
		},
	}

	for i, test := range tests {
		result := LongestOneSequenceAfterFlip(test.n)
		assert.Equal(t, test.expected, result, i)
	}
}
