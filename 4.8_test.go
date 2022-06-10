package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// CommonNode returns the common node for the node containing itemA and the node containing itemB.
func CommonAncestor(root *BinaryNode, itemA int, itemB int) (*BinaryNode, bool) {
	result := commonAncestor(root, itemA, itemB)
	if result.foundNodeA && result.foundNodeB {
		return result.commonAncestor, true
	}

	return nil, false
}

type commonAncestorResult struct {
	foundNodeA     bool
	foundNodeB     bool
	commonAncestor *BinaryNode
}

func commonAncestor(node *BinaryNode, itemA int, itemB int) commonAncestorResult {
	if node == nil {
		return commonAncestorResult{}
	}

	leftResult := commonAncestor(node.Left, itemA, itemB)
	if leftResult.foundNodeA && leftResult.foundNodeB {
		return leftResult
	}

	rightResult := commonAncestor(node.Right, itemA, itemB)

	foundNodeA := leftResult.foundNodeA || rightResult.foundNodeA || node.Item == itemA
	foundNodeB := leftResult.foundNodeB || rightResult.foundNodeB || node.Item == itemB

	return commonAncestorResult{
		foundNodeA:     foundNodeA,
		foundNodeB:     foundNodeB,
		commonAncestor: node,
	}
}

func TestCommonAncestor(t *testing.T) {
	tests := []struct {
		root                 *BinaryNode
		itemA                int
		itemB                int
		expectedAncestorItem int
		expectedOk           bool
	}{
		{
			nil,
			0,
			0,
			0,
			false,
		},
		{
			&BinaryNode{
				Item: 0,
			},
			0,
			0,
			0,
			true,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{Item: 1},
			},
			0,
			1,
			0,
			true,
		},
		{
			&BinaryNode{
				Item:  0,
				Left:  &BinaryNode{Item: 1},
				Right: &BinaryNode{Item: 2},
			},
			1,
			2,
			0,
			true,
		},
		{
			&BinaryNode{
				Item:  0,
				Left:  &BinaryNode{Item: 1},
				Right: &BinaryNode{Item: 2},
			},
			1,
			3,
			0,
			false,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{
						Item:  2,
						Left:  &BinaryNode{Item: 3},
						Right: &BinaryNode{Item: 4},
					},
				},
				Right: &BinaryNode{Item: 5},
			},
			3,
			4,
			2,
			true,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{
						Item:  2,
						Left:  &BinaryNode{Item: 3},
						Right: &BinaryNode{Item: 4},
					},
				},
				Right: &BinaryNode{Item: 5},
			},
			1,
			4,
			1,
			true,
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{
						Item:  2,
						Left:  &BinaryNode{Item: 3},
						Right: &BinaryNode{Item: 4},
					},
				},
				Right: &BinaryNode{
					Item: 5,
					Left: &BinaryNode{Item: 6},
				},
			},
			3,
			6,
			0,
			true,
		},
	}

	for i, test := range tests {
		result, ok := CommonAncestor(test.root, test.itemA, test.itemB)
		assert.Equal(t, test.expectedOk, ok, i)
		if test.expectedOk {
			assert.Equal(t, test.expectedAncestorItem, result.Item, i)
		}
	}
}
