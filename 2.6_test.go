package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func getEnd(head *Node) *Node {
	curr := head
	for {
		curr = curr.next

		if curr.next == nil {
			return curr
		}
	}
}

func listToArray(head *Node) []int {
	if head == nil {
		return nil
	}

	var items []int

	curr := head
	for curr != nil {
		items = append(items, curr.item)
		curr = curr.next
	}
	return items
}

// palindromeSingleLinked returns true if the linked list is a palindrome.
func palindromeSingleLinked(head *Node) bool {
	items := listToArray(head)

	for i := 0; i < len(items)/2; i++ {
		if items[i] != items[len(items)-1-i] {
			return false
		}
	}
	return true
}

func TestPalindromeSingleLinked(t *testing.T) {
	tests := []struct {
		head     *Node
		expected bool
	}{
		{
			nil,
			true,
		},
		{
			&Node{nil, 0},
			true,
		},
		{
			&Node{
				&Node{nil, 1},
				0,
			},
			false,
		},
		{
			&Node{
				&Node{
					&Node{
						nil,
						0,
					},
					1,
				},
				0,
			},
			true,
		},
		{
			&Node{
				&Node{
					&Node{
						nil,
						2,
					},
					1,
				},
				0,
			},
			false,
		},
		{
			&Node{
				&Node{
					&Node{
						&Node{
							nil,
							0,
						},
						1,
					},
					1,
				},
				0,
			},
			true,
		},
		{
			&Node{
				&Node{
					&Node{
						&Node{
							nil,
							0,
						},
						2,
					},
					1,
				},
				0,
			},
			false,
		},
		{
			&Node{
				&Node{
					&Node{
						&Node{
							&Node{
								nil,
								0,
							},
							1,
						},
						2,
					},
					1,
				},
				0,
			},
			true,
		},
	}

	for i, test := range tests {
		assert.Equal(t, test.expected, palindromeSingleLinked(test.head), i)
	}
}
