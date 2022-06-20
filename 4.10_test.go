package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// CheckSubtree returns true if subtree is a sub-tree of tree.
func CheckSubtree(tree *BinaryNode, subtree *BinaryNode) bool {
	if tree == nil {
		return false
	}

	if isEqualTrees(tree, subtree) {
		return true
	}

	return CheckSubtree(tree.Left, subtree) || CheckSubtree(tree.Right, subtree)
}

func isEqualTrees(a *BinaryNode, b *BinaryNode) bool {
	if a == nil && b == nil {
		return true
	}

	if a == nil || b == nil || a.Item != b.Item {
		return false
	}

	return isEqualTrees(a.Left, b.Left) && isEqualTrees(a.Right, b.Right)
}

func TestCheckSubtree(t *testing.T) {
	tests := []struct {
		tree     *BinaryNode
		subtree  *BinaryNode
		expected bool
	}{
		{
			nil,
			nil,
			false,
		},
		{
			&BinaryNode{Item: 0},
			nil,
			false,
		},
		{
			nil,
			&BinaryNode{Item: 0},
			false,
		},
		{
			&BinaryNode{Item: 0},
			&BinaryNode{Item: 0},
			true,
		},
		{
			&BinaryNode{
				Item:  0,
				Left:  &BinaryNode{Item: 1},
				Right: &BinaryNode{Item: 2},
			},
			&BinaryNode{Item: 1},
			true,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{Item: 1},
			},
			&BinaryNode{
				Item:  0,
				Right: &BinaryNode{Item: 1},
			},
			false,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{Item: 3},
				},
				Right: &BinaryNode{Item: 2},
			},
			&BinaryNode{
				Item: 1,
				Left: &BinaryNode{Item: 3},
			},
			true,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{
						Item: 3,
						Left: &BinaryNode{
							Item: 1,
							Right: &BinaryNode{
								Item:  3,
								Right: &BinaryNode{Item: 6},
							},
						},
					},
				},
				Right: &BinaryNode{Item: 2},
			},
			&BinaryNode{
				Item: 1,
				Right: &BinaryNode{
					Item:  3,
					Right: &BinaryNode{Item: 6},
				},
			},
			true,
		},
	}

	for i, test := range tests {
		result := CheckSubtree(test.tree, test.subtree)
		assert.Equal(t, test.expected, result, i)
	}
}
