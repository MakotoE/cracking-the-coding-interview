package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// TripleStep returns the number of ways a child can run up a staircase of size n.
func TripleStep(n int) int {
	return step(n, 0, 1) + step(n, 0, 2) + step(n, 0, 3)
}

func step(n int, current int, size int) int {
	newPos := current + size
	if newPos > n {
		return 0
	} else if newPos == n {
		return 1
	}

	return step(n, newPos, 1) + step(n, newPos, 2) + step(n, newPos, 3)
}

func TestTripleStep(t *testing.T) {
	tests := []struct {
		n        int
		expected int
	}{
		{
			0,
			0,
		},
		{
			1,
			1,
		},
		{
			2,
			2,
		},
		{
			3,
			4,
		},
		{
			4,
			7,
		},
		{
			5,
			13,
		},
	}

	for i, test := range tests {
		result := TripleStep(test.n)
		assert.Equal(t, test.expected, result, i)
	}
}
