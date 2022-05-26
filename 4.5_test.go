package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// IsBST returns true of given tree is a BST.
func IsBST(root *BinaryNode) bool {
	return nodeIsBST(root, math.MinInt, math.MaxInt)
}

func nodeIsBST(node *BinaryNode, min int, max int) bool {
	if node == nil {
		return true
	}

	if node.Left != nil && !(min <= node.Left.Item && node.Left.Item <= node.Item) {
		return false
	}

	if node.Right != nil && !(node.Item <= node.Right.Item && node.Right.Item <= max) {
		return false
	}

	if !nodeIsBST(node.Left, min, node.Item) ||
		!nodeIsBST(node.Right, node.Item, max) {
		return false
	}

	return true
}

func TestIsBST(t *testing.T) {
	tests := []struct {
		node     *BinaryNode
		expected bool
	}{
		{
			nil,
			true,
		},
		{
			&BinaryNode{
				Item: 0,
			},
			true,
		},
		{
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{Item: 0},
			},
			true,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{Item: 1},
			},
			false,
		},
		{
			&BinaryNode{
				Item:  1,
				Right: &BinaryNode{Item: 0},
			},
			false,
		},
		{
			&BinaryNode{
				Item:  1,
				Left:  &BinaryNode{Item: 0},
				Right: &BinaryNode{Item: 2},
			},
			true,
		},
		{
			&BinaryNode{
				Item: 2,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{Item: 0},
				},
				Right: &BinaryNode{
					Item:  3,
					Right: &BinaryNode{Item: 4},
				},
			},
			true,
		},
		{
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{
					Item:  0,
					Right: &BinaryNode{Item: 2},
				},
				Right: &BinaryNode{
					Item: 3,
				},
			},
			false,
		},
		{
			&BinaryNode{
				Item: 2,
				Right: &BinaryNode{
					Item: 3,
					Left: &BinaryNode{Item: 1},
				},
			},
			false,
		},
	}

	for i, test := range tests {
		result := IsBST(test.node)
		assert.Equal(t, test.expected, result, i)
	}
}
