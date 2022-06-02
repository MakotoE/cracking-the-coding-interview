package cracking_the_coding_interview

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type Dependency struct {
	DependsOn int
	Project   int
}

type GraphNodeFlag struct {
	ID            int
	Children      []*GraphNodeFlag
	Marked        bool
	TemporaryMark bool
}

// BuildOrder returns a valid build order if one exists.
func BuildOrder(projects []int, dependencies []Dependency) ([]int, bool) {
	graph := createGraph(projects, dependencies)

	var reversed []int

	for _, node := range graph {
		nodes, ok := visit(node)
		if !ok {
			return nil, false
		}
		reversed = append(reversed, nodes...)
	}

	result := make([]int, len(reversed))
	for i, n := range reversed {
		result[len(result)-1-i] = n
	}
	return result, true
}

func visit(node *GraphNodeFlag) ([]int, bool) {
	if node.Marked {
		return nil, true
	}
	if node.TemporaryMark {
		return nil, false
	}

	node.TemporaryMark = true

	var result []int

	for _, child := range node.Children {
		nodes, ok := visit(child)
		if !ok {
			return nil, false
		}
		result = append(result, nodes...)
	}

	node.TemporaryMark = false
	node.Marked = true
	return append(result, node.ID), true
}

func createGraph(projects []int, dependencies []Dependency) map[int]*GraphNodeFlag {
	nodes := make(map[int]*GraphNodeFlag)

	for _, project := range projects {
		nodes[project] = &GraphNodeFlag{ID: project}
	}

	for _, dependency := range dependencies {
		projectNode, ok := nodes[dependency.Project]
		if !ok {
			panic(fmt.Sprintf("node %d not found", dependency.Project))
		}

		dependsOnNode, ok := nodes[dependency.DependsOn]
		if !ok {
			panic(fmt.Sprintf("node %d not found", dependency.DependsOn))
		}

		dependsOnNode.Children = append(dependsOnNode.Children, projectNode)
	}

	return nodes
}

func testResult(t *testing.T, projects []int, dependencies []Dependency, result []int) {
	graph := createGraph(projects, dependencies)

	for _, n := range result {
		for _, child := range graph[n].Children {
			assert.False(t, child.Marked)
		}

		graph[n].Marked = true
	}
}

func TestBuildOrder(t *testing.T) {
	tests := []struct {
		projects     []int
		dependencies []Dependency
		expectedOk   bool
	}{
		{
			nil,
			nil,
			true,
		},
		{
			[]int{0},
			nil,
			true,
		},
		{
			[]int{0, 1, 2, 3, 4, 5, 6},
			[]Dependency{
				{0, 3},
				{6, 3},
				{1, 3},
				{6, 0},
				{3, 2},
			},
			true,
		},
		{
			[]int{0, 1},
			[]Dependency{
				{0, 1},
				{1, 0},
			},
			false,
		},
	}

	for i, test := range tests {
		result, ok := BuildOrder(test.projects, test.dependencies)
		assert.Equal(t, test.expectedOk, ok, i)

		if test.expectedOk {
			testResult(t, test.projects, test.dependencies, result)
		}
	}
}
