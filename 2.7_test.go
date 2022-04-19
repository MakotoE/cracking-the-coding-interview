package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
	"unsafe"
)

// getPointers returns all pointers to nodes.
func getPointers(list *IntNode) []uintptr {
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
func isIntersecting(a *IntNode, b *IntNode) bool {
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

func listLength(list *IntNode) int {
	result := 0
	curr := list
	for curr != nil {
		curr = curr.next
		result++
	}
	return result
}

func isIntersecting2(a *IntNode, b *IntNode) bool {
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
		assert.False(t, isIntersecting2(&IntNode{}, &IntNode{}))
	}
	{
		a := &IntNode{}
		b := &IntNode{next: a}
		assert.True(t, isIntersecting2(a, b))
		assert.True(t, isIntersecting2(b, a))
	}
	{
		a := &IntNode{}
		b := &IntNode{next: a}
		c := &IntNode{next: a}
		assert.True(t, isIntersecting2(b, c))
		assert.True(t, isIntersecting2(c, b))
	}
	{
		a := &IntNode{next: &IntNode{}}
		b := &IntNode{next: a}
		c := &IntNode{next: a}
		assert.True(t, isIntersecting2(b, c))
		assert.True(t, isIntersecting2(c, b))
	}
	{
		a := &IntNode{}
		b := &IntNode{next: &IntNode{next: a}}
		c := &IntNode{next: a}
		assert.True(t, isIntersecting2(b, c))
		assert.True(t, isIntersecting2(c, b))
	}
}
