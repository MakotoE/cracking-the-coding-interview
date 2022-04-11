package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

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

type DLNode struct {
	next *DLNode
	prev *DLNode
	item int
}

func newDoubleLinkedList(items []int) *DLNode {
	if len(items) == 0 {
		return nil
	}

	head := &DLNode{}
	curr := head

	for _, item := range items {
		curr.item = item
		curr.next = &DLNode{prev: curr}
		curr = curr.next
	}

	curr.prev.next = nil
	return head
}

func getEnd(head *DLNode) *DLNode {
	curr := head
	for curr.next != nil {
		curr = curr.next
	}
	return curr
}

// palindromeDoubleLinked returns true if the linked list is a palindrome.
func palindromeDoubleLinked(head *DLNode) bool {
	if head == nil {
		return true
	}

	start := head
	end := getEnd(head)

	for start != end {
		if start.item != end.item {
			return false
		}
		start = start.next
		end = end.prev
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
		{
			&Node{
				&Node{
					&Node{
						&Node{
							&Node{
								nil,
								0,
							},
							2,
						},
						2,
					},
					1,
				},
				0,
			},
			false,
		},
	}

	for i, test := range tests {
		assert.Equal(t, test.expected, palindromeSingleLinked(test.head), i)
	}
}

func TestPalindromeDoubleLinked(t *testing.T) {
	tests := []struct {
		head     *DLNode
		expected bool
	}{
		{
			nil,
			true,
		},
		{
			newDoubleLinkedList([]int{0}),
			true,
		},
		{
			newDoubleLinkedList([]int{0, 1}),
			false,
		},
		{
			newDoubleLinkedList([]int{0, 1, 0}),
			true,
		},
		{
			newDoubleLinkedList([]int{0, 1, 2}),
			false,
		},
		{
			newDoubleLinkedList([]int{0, 1, 1, 0}),
			true,
		},
		{
			newDoubleLinkedList([]int{0, 1, 2, 0}),
			false,
		},
		{
			newDoubleLinkedList([]int{0, 1, 2, 1, 0}),
			true,
		},
		{
			newDoubleLinkedList([]int{0, 1, 2, 2, 1, 0}),
			true,
		},
		{
			newDoubleLinkedList([]int{0, 1, 2, 2, 3, 0}),
			false,
		},
	}

	for i, test := range tests {
		assert.Equal(t, test.expected, palindromeDoubleLinked(test.head), i)
	}
}
