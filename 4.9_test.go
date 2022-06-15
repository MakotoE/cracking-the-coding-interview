package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// BSTSequence returns all possible sequences of binary tree insertions.
func BSTSequence(root *BinaryNode) [][]int {
	reversed := bstSequenceReversed(root)

	for _, arr := range reversed {
		reverse(arr)
	}

	return reversed
}

func reverse(arr []int) {
	i := 0
	j := len(arr) - 1

	for i < j {
		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
}

func bstSequenceReversed(node *BinaryNode) [][]int {
	if node == nil {
		return nil
	}

	if node.Left == nil && node.Right == nil {
		return [][]int{{node.Item}}
	}

	leftSequences := bstSequenceReversed(node.Left)
	rightSequences := bstSequenceReversed(node.Right)

	var permutation0 [][]int
	permutation1 := make([][]int, len(leftSequences)*len(rightSequences))

	if len(leftSequences) == 0 {
		permutation0 = make([][]int, len(rightSequences))
		for i, right := range rightSequences {
			permutation0[i] = append(permutation0[i], right...)
		}
	} else if len(rightSequences) == 0 {
		permutation0 = make([][]int, len(leftSequences))
		for i, left := range leftSequences {
			permutation0[i] = append(permutation0[i], left...)
		}
	} else {
		permutation0 = make([][]int, len(leftSequences)*len(rightSequences))
		i := 0
		for _, left := range leftSequences {
			for _, right := range rightSequences {
				permutation0[i] = append(permutation0[i], left...)
				permutation0[i] = append(permutation0[i], right...)

				permutation1[i] = append(permutation1[i], right...)
				permutation1[i] = append(permutation1[i], left...)

				i++
			}
		}
	}

	for i := range permutation0 {
		permutation0[i] = append(permutation0[i], node.Item)
	}

	for i := range permutation1 {
		permutation1[i] = append(permutation1[i], node.Item)
	}

	return append(permutation0, permutation1...)
}

func TestBSTSequence(t *testing.T) {
	tests := []struct {
		root     *BinaryNode
		expected [][]int
	}{
		{
			nil,
			nil,
		},
		{
			&BinaryNode{Item: 0},
			[][]int{
				{0},
			},
		},
		{
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{Item: 0},
			},
			[][]int{
				{1, 0},
			},
		},
		{
			&BinaryNode{
				Item:  0,
				Right: &BinaryNode{Item: 1},
			},
			[][]int{
				{0, 1},
			},
		},
		{
			&BinaryNode{
				Item:  1,
				Left:  &BinaryNode{Item: 0},
				Right: &BinaryNode{Item: 2},
			},
			[][]int{
				{1, 2, 0},
				{1, 0, 2},
			},
		},
		{
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{Item: 0},
				Right: &BinaryNode{
					Item:  2,
					Right: &BinaryNode{Item: 3},
				},
			},
			[][]int{
				{1, 2, 3, 0},
				{1, 0, 2, 3},
			},
		},
		{
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{Item: 0},
				Right: &BinaryNode{
					Item:  3,
					Left:  &BinaryNode{Item: 2},
					Right: &BinaryNode{Item: 4},
				},
			},
			[][]int{
				{1, 3, 4, 2, 0},
				{1, 3, 2, 4, 0},
				{1, 0, 3, 4, 2},
				{1, 0, 3, 2, 4},
			},
		},
		{
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{Item: 0},
				Right: &BinaryNode{
					Item:  3,
					Left:  &BinaryNode{Item: 2},
					Right: &BinaryNode{Item: 4},
				},
			},
			[][]int{
				{1, 3, 4, 2, 0},
				{1, 3, 2, 4, 0},
				{1, 0, 3, 4, 2},
				{1, 0, 3, 2, 4},
			},
		},
		{
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{
					Item:  -1,
					Left:  &BinaryNode{Item: -2},
					Right: &BinaryNode{Item: 0},
				},
				Right: &BinaryNode{
					Item:  3,
					Left:  &BinaryNode{Item: 2},
					Right: &BinaryNode{Item: 4},
				},
			},
			[][]int{
				{1, 3, 4, 2, -1, 0, -2},
				{1, 3, 2, 4, -1, 0, -2},
				{1, 3, 4, 2, -1, -2, 0},
				{1, 3, 2, 4, -1, -2, 0},
				{1, -1, 0, -2, 3, 4, 2},
				{1, -1, 0, -2, 3, 2, 4},
				{1, -1, -2, 0, 3, 4, 2},
				{1, -1, -2, 0, 3, 2, 4},
			},
		},
	}

	for i, test := range tests {
		result := BSTSequence(test.root)
		assert.Equal(t, test.expected, result, i)
	}
}
