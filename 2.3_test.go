package cracking_the_coding_interview

import (
	"github.com/MakotoE/cracking-the-coding-interview/sll"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Deletes the item at given index. The given index is valid and is not 0 or the last item.
func deleteMiddleNode(head *sll.Node, index int) {
	parent := head
	for i := 0; i < index-1; i++ {
		parent = parent.Next
	}

	parent.Next = parent.Next.Next
}

func TestDeleteMiddleNode(t *testing.T) {
	tests := []struct {
		head     *sll.Node
		index    int
		expected *sll.Node
	}{
		{
			&sll.Node{
				Next: &sll.Node{
					Next: &sll.Node{Item: 2},
					Item: 1,
				},
				Item: 0,
			},
			1,
			&sll.Node{
				Next: &sll.Node{Item: 2},
				Item: 0,
			},
		},
		{
			&sll.Node{
				Next: &sll.Node{
					Next: &sll.Node{
						Next: &sll.Node{Item: 3},
						Item: 2,
					},
					Item: 1,
				},
				Item: 0,
			},
			1,
			&sll.Node{
				Next: &sll.Node{
					Next: &sll.Node{Item: 3},
					Item: 2,
				},
				Item: 0,
			},
		},
		{
			&sll.Node{
				Next: &sll.Node{
					Next: &sll.Node{
						Next: &sll.Node{Item: 3},
						Item: 2,
					},
					Item: 1,
				},
				Item: 0,
			},
			2,
			&sll.Node{
				Next: &sll.Node{
					Next: &sll.Node{Item: 3},
					Item: 1,
				},
				Item: 0,
			},
		},
	}

	for i, test := range tests {
		deleteMiddleNode(test.head, test.index)
		assert.EqualValues(t, test.expected, test.head, i)
	}
}
