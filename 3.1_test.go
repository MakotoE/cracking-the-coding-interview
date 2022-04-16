package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const interval = 64
const period = interval * 3

// ThreeInOne is a tuple of 3 stacks that are backed by 1 array.
type ThreeInOne struct {
	array []int
	aSize int
	bSize int
	cSize int
}

type StackID int

const (
	StackA StackID = iota
	StackB
	StackC
)

func StackIndexToArrayIndex(stack StackID, index int) int {
	periodNumber := index / interval
	return periodNumber*period + int(stack)*interval + index%interval
}

func (t *ThreeInOne) getSize(stack StackID) int {
	switch stack {
	case StackA:
		return t.aSize
	case StackB:
		return t.bSize
	case StackC:
		return t.cSize
	default:
		panic("unexpected")
	}
}

func (t *ThreeInOne) addIndex(stack StackID, delta int) {
	switch stack {
	case StackA:
		t.aSize += delta
	case StackB:
		t.bSize += delta
	case StackC:
		t.cSize += delta
	}
}

func (t *ThreeInOne) Push(stack StackID, item int) {
	arrayIndex := StackIndexToArrayIndex(stack, t.getSize(stack))
	if arrayIndex >= len(t.array) {
		tmp := make([]int, len(t.array)+period, len(t.array)+period)
		copy(tmp, t.array)
		t.array = tmp
	}

	t.array[arrayIndex] = item
	t.addIndex(stack, 1)
}

func (t *ThreeInOne) Pop(stack StackID) (int, bool) {
	stackSize := t.getSize(stack)
	if stackSize == 0 {
		return 0, false
	}

	arrayIndex := StackIndexToArrayIndex(stack, stackSize-1)
	tmp := t.array[arrayIndex]
	// To catch mistakes
	t.array[arrayIndex] = 0
	t.addIndex(stack, -1)

	return tmp, true
}

func (t *ThreeInOne) IsEmpty(stack StackID) bool {
	switch stack {
	case StackA:
		return t.aSize == 0
	case StackB:
		return t.bSize == 0
	case StackC:
		return t.cSize == 0
	default:
		panic("unexpected")
	}
}

func TestThreeInOne(t *testing.T) {
	{
		stack := ThreeInOne{}
		assert.True(t, stack.IsEmpty(StackA))
		_, ok := stack.Pop(0)
		assert.False(t, ok)

		stack.Push(StackA, 0)
		n, ok := stack.Pop(StackA)
		assert.True(t, ok)
		assert.Equal(t, 0, n)

		_, ok = stack.Pop(StackA)
		assert.False(t, ok)

		stack.Push(StackA, 1)
		stack.Push(StackA, 2)
		n, _ = stack.Pop(StackA)
		assert.Equal(t, 2, n)
		n, _ = stack.Pop(StackA)
		assert.Equal(t, 1, n)
		assert.True(t, stack.IsEmpty(StackA))
	}
	{
		stack := ThreeInOne{}
		stack.Push(StackA, 0)
		stack.Push(StackB, 1)
		stack.Push(StackC, 2)

		n, ok := stack.Pop(StackA)
		assert.True(t, ok)
		assert.Equal(t, 0, n)
		n, ok = stack.Pop(StackB)
		assert.True(t, ok)
		assert.Equal(t, 1, n)
		n, ok = stack.Pop(StackC)
		assert.True(t, ok)
		assert.Equal(t, 2, n)

		assert.True(t, stack.IsEmpty(StackA))
		assert.True(t, stack.IsEmpty(StackB))
		assert.True(t, stack.IsEmpty(StackC))
	}
	{
		stack := ThreeInOne{}
		for i := 0; i < interval+1; i++ {
			stack.Push(StackA, i)
		}

		stack.Push(StackB, 100)
		stack.Push(StackC, 101)

		n, ok := stack.Pop(StackA)
		assert.True(t, ok)
		assert.Equal(t, interval, n)

		n, ok = stack.Pop(StackB)
		assert.True(t, ok)
		assert.Equal(t, 100, n)
		n, ok = stack.Pop(StackC)
		assert.True(t, ok)
		assert.Equal(t, 101, n)

		for i := 0; i < interval+1; i++ {
			stack.Push(StackC, i)
		}
		n, ok = stack.Pop(StackC)
		assert.True(t, ok)
		assert.Equal(t, interval, n)
	}
}
