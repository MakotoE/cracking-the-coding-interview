package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"math"
	"sort"
	"testing"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

// SmallestDifference returns the smallest difference between any number in a and any number in b.
func SmallestDifference(a []int, b []int) int {
	if len(a) == 0 || len(b) == 0 {
		return 0
	}

	sort.Ints(a)
	sort.Ints(b)

	aIndex := 0
	bIndex := 0
	minDiff := math.MaxInt

	for {
		if abs(a[aIndex]-b[bIndex]) < minDiff {
			minDiff = abs(a[aIndex] - b[bIndex])
		}

		if aIndex == len(a)-1 && bIndex == len(b)-1 {
			return minDiff
		}

		if aIndex < len(a)-1 && a[aIndex] < b[bIndex] || bIndex == len(b)-1 {
			if aIndex+1 < len(a) {
				aIndex++
			}
		} else if bIndex+1 < len(b) {
			bIndex++
		}
	}
}

func TestSmallestDifference(t *testing.T) {
	tests := []struct {
		a        []int
		b        []int
		expected int
	}{
		{
			nil,
			nil,
			0,
		},
		{
			[]int{0},
			[]int{0},
			0,
		},
		{
			[]int{1},
			[]int{0},
			1,
		},
		{
			[]int{0},
			[]int{1},
			1,
		},
		{
			[]int{0, 1},
			[]int{0},
			0,
		},
		{
			[]int{2, 1},
			[]int{0},
			1,
		},
		{
			[]int{0},
			[]int{2, 1},
			1,
		},
		{
			[]int{0, 2},
			[]int{2, 1},
			0,
		},
		{
			[]int{1, 3, 15, 11, 2},
			[]int{23, 127, 235, 19, 8},
			3,
		},
		{
			[]int{4, 6, 8, 10},
			[]int{5, 9, 12},
			1,
		},
	}

	for i, test := range tests {
		result := SmallestDifference(test.a, test.b)
		assert.Equal(t, test.expected, result, i)
	}
}
