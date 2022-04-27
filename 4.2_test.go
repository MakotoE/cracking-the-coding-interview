package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type BinaryNode struct {
	Item  int
	Left  *BinaryNode
	Right *BinaryNode
}

// MinimalTree inputs a sorted array and outputs a binary search tree.
func MinimalTree(arr []int) *BinaryNode {
	if !sort.IntsAreSorted(arr) {
		panic("arr is not sorted")
	}
	return createNode(arr, 0, len(arr))
}

func createNode(arr []int, start int, end int) *BinaryNode {
	if start >= end {
		return nil
	}

	return &BinaryNode{
		Item:  arr[start+(end-start)/2],
		Left:  createNode(arr, start, start+(end-start)/2),
		Right: createNode(arr, start+(end-start)/2+1, end),
	}
}

func TestMinimalTree(t *testing.T) {
	tests := []struct {
		arr      []int
		expected *BinaryNode
	}{
		{
			nil,
			nil,
		},
		{
			[]int{0},
			&BinaryNode{Item: 0},
		},
		{
			[]int{0, 1},
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{Item: 0},
			},
		},
		{
			[]int{0, 0},
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{Item: 0},
			},
		},
		{
			[]int{0, 1, 2},
			&BinaryNode{
				Item:  1,
				Left:  &BinaryNode{Item: 0},
				Right: &BinaryNode{Item: 2},
			},
		},
		{
			[]int{0, 0, 0},
			&BinaryNode{
				Item:  0,
				Left:  &BinaryNode{Item: 0},
				Right: &BinaryNode{Item: 0},
			},
		},
		{
			[]int{0, 1, 2, 3},
			&BinaryNode{
				Item: 2,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{Item: 0},
				},
				Right: &BinaryNode{Item: 3},
			},
		},
		{
			[]int{0, 1, 2, 3, 4},
			&BinaryNode{
				Item: 2,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{Item: 0},
				},
				Right: &BinaryNode{
					Item: 4,
					Left: &BinaryNode{Item: 3},
				},
			},
		},
		{
			[]int{0, 1, 2, 3, 4, 5},
			&BinaryNode{
				Item: 3,
				Left: &BinaryNode{
					Item:  1,
					Left:  &BinaryNode{Item: 0},
					Right: &BinaryNode{Item: 2},
				},
				Right: &BinaryNode{
					Item: 5,
					Left: &BinaryNode{Item: 4},
				},
			},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			&BinaryNode{
				Item: 3,
				Left: &BinaryNode{
					Item:  1,
					Left:  &BinaryNode{Item: 0},
					Right: &BinaryNode{Item: 2},
				},
				Right: &BinaryNode{
					Item:  5,
					Left:  &BinaryNode{Item: 4},
					Right: &BinaryNode{Item: 6},
				},
			},
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6, 7},
			&BinaryNode{
				Item: 4,
				Left: &BinaryNode{
					Item: 2,
					Left: &BinaryNode{
						Item: 1,
						Left: &BinaryNode{Item: 0},
					},
					Right: &BinaryNode{Item: 3},
				},
				Right: &BinaryNode{
					Item:  6,
					Left:  &BinaryNode{Item: 5},
					Right: &BinaryNode{Item: 7},
				},
			},
		},
	}

	for i, test := range tests {
		result := MinimalTree(test.arr)
		assert.True(t, assert.ObjectsAreEqualValues(test.expected, result), i)
	}
}
