package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type Node struct {
	next *Node
	item int
}

// Deletes the item at given index. The given index is valid and is not 0 or the last item.
func deleteMiddleNode(head *Node, index int) {
	parent := head
	for i := 0; i < index-1; i++ {
		parent = parent.next
	}

	parent.next = parent.next.next
}

func TestDeleteMiddleNode(t *testing.T) {
	tests := []struct {
		head     *Node
		index    int
		expected *Node
	}{
		{
			&Node{
				&Node{
					&Node{nil, 2},
					1,
				},
				0,
			},
			1,
			&Node{
				&Node{
					nil,
					2,
				},
				0,
			},
		},
		{
			&Node{
				&Node{
					&Node{
						&Node{nil, 3},
						2,
					},
					1,
				},
				0,
			},
			1,
			&Node{
				&Node{
					&Node{nil, 3},
					2,
				},
				0,
			},
		},
		{
			&Node{
				&Node{
					&Node{
						&Node{nil, 3},
						2,
					},
					1,
				},
				0,
			},
			2,
			&Node{
				&Node{
					&Node{nil, 3},
					1,
				},
				0,
			},
		},
	}

	for i, test := range tests {
		deleteMiddleNode(test.head, test.index)
		assert.EqualValues(t, test.expected, test.head, i)
	}
}
