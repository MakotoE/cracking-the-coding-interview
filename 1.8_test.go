package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ZeroMatrix sets row and column of every zero element to zero.
func ZeroMatrix(matrix [][]int) {
	rows := make(map[int]bool)
	cols := make(map[int]bool)

	for row := range matrix {
		for col := range matrix[row] {
			if matrix[row][col] == 0 {
				rows[row] = true
				cols[col] = true
			}
		}
	}

	for row := range rows {
		for i := range matrix[row] {
			matrix[row][i] = 0
		}
	}

	for col := range cols {
		for _, row := range matrix {
			row[col] = 0
		}
	}
}

func TestZeroMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			nil,
			nil,
		},
		{
			[][]int{
				{0},
			},
			[][]int{
				{0},
			},
		},
		{
			[][]int{
				{1},
			},
			[][]int{
				{1},
			},
		},
		{
			[][]int{
				{1, 0},
			},
			[][]int{
				{0, 0},
			},
		},
		{
			[][]int{
				{1}, {0},
			},
			[][]int{
				{0}, {0},
			},
		},
		{
			[][]int{
				{0, 1},
				{1, 1},
			},
			[][]int{
				{0, 0},
				{0, 1},
			},
		},
		{
			[][]int{
				{1, 1, 1},
				{1, 0, 1},
				{1, 1, 1},
			},
			[][]int{
				{1, 0, 1},
				{0, 0, 0},
				{1, 0, 1},
			},
		},
	}

	for i, test := range tests {
		ZeroMatrix(test.matrix)
		assert.True(t, assert.ObjectsAreEqualValues(test.expected, test.matrix), i)
	}
}
