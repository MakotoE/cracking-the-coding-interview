package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// StackQueue is a queue that is implemented with two stacks.
type StackQueue struct {
	A []int
	B []int
}

func (s *StackQueue) Push(item int) {
	s.A = append(s.A, item)
}

func (s *StackQueue) Pop() (int, bool) {
	if len(s.B) == 0 {
		for len(s.A) > 0 {
			tmp := s.A[len(s.A)-1]
			s.A = s.A[:len(s.A)-1]
			s.B = append(s.B, tmp)
		}
	}

	if len(s.B) == 0 {
		return 0, false
	}

	tmp := s.B[len(s.B)-1]
	s.B = s.B[:len(s.B)-1]
	return tmp, true
}

func TestStackQueue(t *testing.T) {
	queue := StackQueue{}
	_, ok := queue.Pop()
	assert.False(t, ok)

	queue.Push(0)
	item, ok := queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 0, item)

	queue.Push(1)
	queue.Push(2)
	item, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 1, item)

	queue.Push(3)
	item, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 2, item)

	queue.Push(4)
	queue.Push(5)
	item, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 3, item)

	queue.Push(6)

	item, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 4, item)
	item, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 5, item)
	item, ok = queue.Pop()
	assert.True(t, ok)
	assert.Equal(t, 6, item)
}
