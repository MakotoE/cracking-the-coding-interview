package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"unsafe"
)

// getPointers returns all pointers to nodes.
func getPointers(list *Node) []uintptr {
	var result []uintptr

	curr := list
	for curr != nil {
		result = append(result, uintptr(unsafe.Pointer(curr)))
		curr = curr.next
	}

	return result
}

type UIntSlice []uintptr

func (u UIntSlice) Len() int {
	return len(u)
}

func (u UIntSlice) Less(i int, j int) bool {
	return u[i] < u[j]
}

func (u UIntSlice) Swap(i int, j int) {
	tmp := u[i]
	u[i] = u[j]
	u[j] = tmp
}

// Returns true if the linked lists have an intersection by reference.
func isIntersecting(a *Node, b *Node) bool {
	aPointers := getPointers(a)
	sort.Sort(UIntSlice(aPointers))

	curr := b
	for curr != nil {
		ptr := uintptr(unsafe.Pointer(curr))
		index := sort.Search(len(aPointers), func(i int) bool {
			return aPointers[i] >= ptr
		})
		if aPointers[index] == ptr {
			return true
		}
		curr = curr.next
	}
	return false
}

func listLength(list *Node) int {
	result := 0
	curr := list
	for curr != nil {
		curr = curr.next
		result++
	}
	return result
}

func isIntersecting2(a *Node, b *Node) bool {
	difference := listLength(a) - listLength(b)

	aStart := a
	bStart := b
	if difference < 0 {
		for i := 0; i < -difference; i++ {
			bStart = bStart.next
		}
	} else if difference > 0 {
		for i := 0; i < difference; i++ {
			aStart = aStart.next
		}
	}

	for aStart != nil {
		if aStart == bStart {
			return true
		}

		aStart = aStart.next
		bStart = bStart.next
	}

	return false
}

func TestIsIntersecting(t *testing.T) {
	{
		assert.False(t, isIntersecting2(nil, nil))
		assert.False(t, isIntersecting2(&Node{}, &Node{}))
	}
	{
		a := &Node{}
		b := &Node{next: a}
		assert.True(t, isIntersecting2(a, b))
		assert.True(t, isIntersecting2(b, a))
	}
	{
		a := &Node{}
		b := &Node{next: a}
		c := &Node{next: a}
		assert.True(t, isIntersecting2(b, c))
		assert.True(t, isIntersecting2(c, b))
	}
	{
		a := &Node{next: &Node{}}
		b := &Node{next: a}
		c := &Node{next: a}
		assert.True(t, isIntersecting2(b, c))
		assert.True(t, isIntersecting2(c, b))
	}
	{
		a := &Node{}
		b := &Node{next: &Node{next: a}}
		c := &Node{next: a}
		assert.True(t, isIntersecting2(b, c))
		assert.True(t, isIntersecting2(c, b))
	}
}
