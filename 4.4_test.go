package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// IsBalanced returns true if the binary tree is balanced.
func IsBalanced(root *BinaryNode) bool {
	if root == nil {
		return true
	}

	diff := height(root.Left) - height(root.Right)
	return -1 <= diff && diff <= 1
}

func height(node *BinaryNode) int {
	if node == nil {
		return 0
	}

	leftHeight := height(node.Left)
	rightHeight := height(node.Right)

	if leftHeight > rightHeight {
		return leftHeight + 1
	}

	return rightHeight + 1
}

func TestIsBalanced(t *testing.T) {
	tests := []struct {
		root     *BinaryNode
		expected bool
	}{
		{
			nil,
			true,
		},
		{
			&BinaryNode{
				Left: &BinaryNode{},
			},
			true,
		},
		{
			&BinaryNode{
				Right: &BinaryNode{},
			},
			true,
		},
		{
			&BinaryNode{
				Left:  &BinaryNode{},
				Right: &BinaryNode{},
			},
			true,
		},
		{
			&BinaryNode{
				Left: &BinaryNode{
					Left: &BinaryNode{},
				},
			},
			false,
		},
		{
			&BinaryNode{
				Right: &BinaryNode{
					Right: &BinaryNode{},
				},
			},
			false,
		},
	}

	for i, test := range tests {
		result := IsBalanced(test.root)
		assert.Equal(t, test.expected, result, i)
	}
}
