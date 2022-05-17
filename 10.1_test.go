package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// SortedMerge merges sorted array b into a. a must have capacity hold all elements of b. Returns
// merged array.
func SortedMerge(a []int, b []int) []int {
	if !sort.IntsAreSorted(a) || !sort.IntsAreSorted(b) {
		panic("a or b is unsorted")
	}

	i := 0
	aIndex := len(a) - 1
	bIndex := len(b) - 1

	a = a[:len(a)+len(b)]

	for aIndex >= 0 && bIndex >= 0 {
		if a[aIndex] > b[bIndex] {
			a[len(a)-1-i] = a[aIndex]
			if aIndex == 0 {
				i++
				break
			}

			aIndex--
		} else {
			a[len(a)-1-i] = b[bIndex]
			bIndex--
		}

		i++
	}

	for j := 0; j <= bIndex; j++ {
		a[len(a)-1-i-j] = b[bIndex-j]
	}

	return a
}

func TestSortedMerge(t *testing.T) {
	tests := []struct {
		aArray   []int
		b        []int
		expected []int
	}{
		{
			nil,
			nil,
			[]int{},
		},
		{
			[]int{0},
			nil,
			[]int{0},
		},
		{
			nil,
			[]int{0},
			[]int{0},
		},
		{
			[]int{0},
			[]int{0},
			[]int{0, 0},
		},
		{
			[]int{1},
			[]int{0},
			[]int{0, 1},
		},
		{
			[]int{0},
			[]int{1},
			[]int{0, 1},
		},
		{
			[]int{0, 1},
			nil,
			[]int{0, 1},
		},
		{
			nil,
			[]int{0, 1},
			[]int{0, 1},
		},
		{
			[]int{0, 2},
			[]int{1},
			[]int{0, 1, 2},
		},
		{
			[]int{0},
			[]int{1, 2},
			[]int{0, 1, 2},
		},
		{
			[]int{0, 1},
			[]int{2, 3},
			[]int{0, 1, 2, 3},
		},
		{
			[]int{2, 3},
			[]int{0, 1},
			[]int{0, 1, 2, 3},
		},
		{
			[]int{0, 2},
			[]int{1, 3},
			[]int{0, 1, 2, 3},
		},
		{
			[]int{0, 1, 2, 4},
			[]int{3, 5},
			[]int{0, 1, 2, 3, 4, 5},
		},
		{
			[]int{3, 5},
			[]int{0, 1, 2, 4},
			[]int{0, 1, 2, 3, 4, 5},
		},
	}

	for i, test := range tests {
		a := make([]int, len(test.aArray), len(test.aArray)+len(test.b))
		copy(a, test.aArray)
		result := SortedMerge(a, test.b)
		assert.Equal(t, test.expected, result, i)
	}
}
