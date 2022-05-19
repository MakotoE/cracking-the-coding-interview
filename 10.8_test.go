package cracking_the_coding_interview

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// FindDuplicates inputs integers where each integer is from 1 to 32000 and returns the numbers that
// are duplicates.
func FindDuplicates(arr []int) []int {
	bitmap := [32000]bool{}
	var result []int

	for _, n := range arr {
		if !(1 <= n && n <= 32000) {
			panic(fmt.Sprintf("n is invalid: %d", n))
		}

		index := n - 1
		if bitmap[index] {
			result = append(result, n)
		} else {
			bitmap[index] = true
		}
	}

	return result
}

func TestFindDuplicates(t *testing.T) {
	tests := []struct {
		arr      []int
		expected []int
	}{
		{
			nil,
			nil,
		},
		{
			[]int{1},
			nil,
		},
		{
			[]int{1, 2},
			nil,
		},
		{
			[]int{1, 1},
			[]int{1},
		},
		{
			[]int{1, 1, 2, 1},
			[]int{1, 1},
		},
	}

	for i, test := range tests {
		result := FindDuplicates(test.arr)
		assert.Equal(t, test.expected, result, i)
	}
}
