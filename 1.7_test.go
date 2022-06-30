package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// RotateMatrix rotates the given matrix by 90 degrees clockwise.
func RotateMatrix(matrix [][]int) {
	newMatrix := make([][]int, len(matrix))

	for row := range newMatrix {
		newMatrix[row] = make([]int, len(matrix))
	}

	for row := range matrix {
		for col := range matrix[row] {
			newMatrix[row][len(matrix)-1-col] = matrix[col][row]
		}
	}

	for row := range matrix {
		for col := range matrix[row] {
			matrix[row][col] = newMatrix[row][col]
		}
	}
}

func TestRotateMatrix(t *testing.T) {
	tests := []struct {
		matrix   [][]int
		expected [][]int
	}{
		{
			nil,
			nil,
		},
		{
			[][]int{{0}},
			[][]int{{0}},
		},
		{
			[][]int{{0, 1}, {2, 3}},
			[][]int{{2, 0}, {3, 1}},
		},
		{
			[][]int{{0, 1, 2}, {3, 4, 5}, {6, 7, 8}},
			[][]int{{6, 3, 0}, {7, 4, 1}, {8, 5, 2}},
		},
	}

	for i, test := range tests {
		RotateMatrix(test.matrix)
		assert.True(t, assert.ObjectsAreEqualValues(test.expected, test.matrix), i)
	}
}
