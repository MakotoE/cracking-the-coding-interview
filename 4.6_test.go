package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type BinaryNodeWithParent struct {
	Item   int
	Left   *BinaryNodeWithParent
	Right  *BinaryNodeWithParent
	Parent *BinaryNodeWithParent
}

// Successor returns the next in-order node of given BST node. Returns nil if there is no successor.
func Successor(node *BinaryNodeWithParent) *BinaryNodeWithParent {
	if node == nil {
		return nil
	}

	if node.Right != nil {
		return leftMostNode(node.Right)
	}

	curr := node

	for {
		if curr.Parent == nil {
			return nil
		}

		if curr.Parent.Left == curr {
			return curr.Parent
		}

		curr = curr.Parent
	}
}

func leftMostNode(node *BinaryNodeWithParent) *BinaryNodeWithParent {
	curr := node

	for curr.Left != nil {
		curr = curr.Left
	}

	return curr
}

func TestSuccessor(t *testing.T) {
	{
		assert.Nil(t, Successor(nil))
		assert.Nil(t, Successor(&BinaryNodeWithParent{Item: 1}))
	}
	{
		tree := &BinaryNodeWithParent{
			Item: 1,
			Left: &BinaryNodeWithParent{
				Item: 0,
			},
		}
		tree.Left.Parent = tree

		assert.Nil(t, Successor(tree))
		assert.Equal(t, tree, Successor(tree.Left))
	}
	{
		tree := &BinaryNodeWithParent{
			Item: 1,
			Right: &BinaryNodeWithParent{
				Item: 2,
			},
		}
		tree.Right.Parent = tree

		assert.Equal(t, tree.Right, Successor(tree))
		assert.Nil(t, Successor(tree.Right))
	}
	{
		tree := &BinaryNodeWithParent{
			Item: 2,
			Left: &BinaryNodeWithParent{
				Item: 0,
				Right: &BinaryNodeWithParent{
					Item: 1,
				},
			},
		}
		tree.Left.Parent = tree
		tree.Left.Right.Parent = tree.Left

		assert.Equal(t, tree.Left.Right, Successor(tree.Left))
		assert.Equal(t, tree, Successor(tree.Left.Right))
	}
	{
		tree := &BinaryNodeWithParent{
			Item: 0,
			Right: &BinaryNodeWithParent{
				Item: 1,
				Right: &BinaryNodeWithParent{
					Item: 2,
				},
			},
		}
		tree.Right.Parent = tree
		tree.Right.Right.Parent = tree.Right

		assert.Equal(t, tree.Right.Right, Successor(tree.Right))
		assert.Nil(t, Successor(tree.Right.Right))
	}
	{
		tree := &BinaryNodeWithParent{
			Item: 3,
			Left: &BinaryNodeWithParent{
				Item: 0,
				Right: &BinaryNodeWithParent{
					Item: 1,
					Right: &BinaryNodeWithParent{
						Item: 2,
					},
				},
			},
		}
		tree.Left.Parent = tree
		tree.Left.Right.Parent = tree.Left
		tree.Left.Right.Right.Parent = tree.Left.Right

		assert.Equal(t, tree, Successor(tree.Left.Right.Right))
	}
	{
		tree := &BinaryNodeWithParent{
			Item: 0,
			Right: &BinaryNodeWithParent{
				Item: 2,
				Left: &BinaryNodeWithParent{
					Item: 1,
				},
			},
		}
		tree.Right.Parent = tree
		tree.Right.Left.Parent = tree.Right

		assert.Equal(t, tree.Right.Left, Successor(tree))
	}
}
