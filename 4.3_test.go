package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Returns a list of items at each depth.
func ListOfDepths(root *BinaryNode) *Node[Node[int]] {
	if root == nil {
		return nil
	}

	nodes := make(map[int]*Node[int])
	getNodes(root, 0, nodes)

	result := &Node[Node[int]]{}
	curr := result

	i := 0
	for ; i < len(nodes)-1; i++ {
		curr.item = *nodes[i]
		curr.next = &Node[Node[int]]{}
		curr = curr.next
	}

	curr.item = *nodes[i]

	return result
}

func addToList(root *Node[int], item int) *Node[int] {
	return &Node[int]{
		item: item,
		next: root,
	}
}

func getNodes(node *BinaryNode, currentDepth int, items map[int]*Node[int]) {
	items[currentDepth] = addToList(items[currentDepth], node.Item)

	if node.Left != nil {
		getNodes(node.Left, currentDepth+1, items)
	}

	if node.Right != nil {
		getNodes(node.Right, currentDepth+1, items)
	}
}

func TestListOfDepths(t *testing.T) {
	tests := []struct {
		root         *BinaryNode
		expectedList [][]int
	}{
		{
			nil,
			nil,
		},
		{
			&BinaryNode{Item: 0},
			[][]int{
				{
					0,
				},
			},
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{Item: 1},
			},
			[][]int{
				{
					0,
				},
				{
					1,
				},
			},
		},
		{
			&BinaryNode{
				Item:  0,
				Left:  &BinaryNode{Item: 1},
				Right: &BinaryNode{Item: 2},
			},
			[][]int{
				{
					0,
				},
				{
					2, 1,
				},
			},
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
			[][]int{
				{
					0,
				},
				{
					2, 1,
				},
				{
					3,
				},
			},
		},
		{
			&BinaryNode{
				Item: 0,
				Left: &BinaryNode{
					Item: 1,
					Left: &BinaryNode{Item: 3},
				},
				Right: &BinaryNode{
					Item:  2,
					Right: &BinaryNode{Item: 4},
				},
			},
			[][]int{
				{
					0,
				},
				{
					2, 1,
				},
				{
					4, 3,
				},
			},
		},
	}

	for i, test := range tests {
		result := ListOfDepths(test.root)

		node := 0
		depthNode := result
		for depthNode != nil {
			itemIndex := 0
			curr := &depthNode.item
			for curr != nil {
				assert.Equal(t, test.expectedList[node][itemIndex], curr.item, i)
				curr = curr.next
				itemIndex++
			}
			depthNode = depthNode.next
			node++
		}
	}
}
