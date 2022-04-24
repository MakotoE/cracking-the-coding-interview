package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GraphNode struct {
	ID       int
	Children []*GraphNode
}

// IsConnected returns true if a and b are connected.
func IsConnected(a *GraphNode, b *GraphNode) bool {
	if a == nil || b == nil {
		return false
	}

	if a.ID == b.ID {
		return true
	}

	visitedA := map[int]bool{a.ID: true}
	visitedB := map[int]bool{b.ID: true}

	stackA := []*GraphNode{a}
	stackB := []*GraphNode{b}

	for {
		if len(stackA) > 0 {
			children := stackA[len(stackA)-1].Children
			stackA = stackA[:len(stackA)-1]
			for _, child := range children {
				if visitedB[child.ID] {
					return true
				}

				if !visitedA[child.ID] {
					stackA = append(stackA, child)
				}
			}
		}

		if len(stackB) > 0 {
			children := stackB[len(stackB)-1].Children
			stackB = stackB[:len(stackB)-1]
			for _, child := range children {
				if visitedA[child.ID] {
					return true
				}

				if !visitedB[child.ID] {
					stackB = append(stackB, child)
				}
			}
		} else if len(stackA) == 0 {
			return false
		}
	}
}

func TestIsConnected(t *testing.T) {
	{
		assert.False(t, IsConnected(nil, nil))
	}
	{
		a := &GraphNode{ID: 0}
		b := &GraphNode{ID: 1}
		assert.False(t, IsConnected(a, b))
	}
	{
		a := &GraphNode{
			ID: 0,
			Children: []*GraphNode{
				{
					ID: 1,
					Children: []*GraphNode{{
						ID:       2,
						Children: []*GraphNode{{ID: 3}},
					}},
				},
				{
					ID:       4,
					Children: []*GraphNode{{ID: 5}},
				},
			},
		}
		assert.True(t, IsConnected(a, a))
		assert.True(t, IsConnected(a, a.Children[0]))
		assert.True(t, IsConnected(a, a.Children[0].Children[0]))
		assert.True(t, IsConnected(a, a.Children[0].Children[0].Children[0]))
		assert.True(t, IsConnected(a, a.Children[1].Children[0]))
	}
	{
		a := &GraphNode{ID: 0}
		a.Children = []*GraphNode{a}
		b := &GraphNode{ID: 1}
		assert.False(t, IsConnected(a, b))
	}
	{
		a := &GraphNode{
			ID:       0,
			Children: []*GraphNode{{ID: 1}},
		}
		a.Children[0].Children = []*GraphNode{{ID: 2, Children: []*GraphNode{a}}}
		assert.True(t, IsConnected(a, a.Children[0].Children[0]))
	}
}
