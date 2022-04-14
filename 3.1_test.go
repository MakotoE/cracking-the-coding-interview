package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const interval = 64

// ThreeInOne is a tuple of 3 stacks that are backed by 1 array.
type ThreeInOne struct {
	array []int
}

func (t *ThreeInOne) Push(array int, item int) {
	t.array = append(t.array, item)
}

func (t *ThreeInOne) Pop(array int) (int, bool) {
	if len(t.array) == 0 {
		return 0, false
	}

	tmp := t.array[len(t.array)-1]
	t.array = t.array[:len(t.array)-1]
	return tmp, true
}

func TestThreeInOne(t *testing.T) {
	{
		stack := ThreeInOne{}
		_, ok := stack.Pop(0)
		assert.False(t, ok)

		stack.Push(0, 0)
		n, ok := stack.Pop(0)
		assert.True(t, ok)
		assert.Equal(t, 0, n)

		_, ok = stack.Pop(0)
		assert.False(t, ok)

		stack.Push(0, 1)
		stack.Push(0, 2)
		n, _ = stack.Pop(0)
		assert.Equal(t, 2, n)
		n, _ = stack.Pop(0)
		assert.Equal(t, 1, n)
	}
}
