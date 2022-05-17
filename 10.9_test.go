package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

type MatrixCoordinates struct {
	Row int
	Col int
}

// SearchMatrix returns the coordinates of element in sorted matrix if it exists.
func SearchMatrix(matrix [][]int, element int) (MatrixCoordinates, bool) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return MatrixCoordinates{}, false
	}

	for row := 0; row < len(matrix); row++ {
		if matrix[row][0] <= element && element <= matrix[row][len(matrix[row])-1] {
			col := sort.SearchInts(matrix[row], element)
			if col < len(matrix[row]) && matrix[row][col] == element {
				return MatrixCoordinates{Row: row, Col: col}, true
			}
		}
	}

	return MatrixCoordinates{}, false
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		matrix              [][]int
		element             int
		expectedCoordinates MatrixCoordinates
		expectedOk          bool
	}{
		{
			nil,
			0,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{},
			},
			0,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{0},
			},
			0,
			MatrixCoordinates{0, 0},
			true,
		},
		{
			[][]int{
				{0},
			},
			1,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{0},
			},
			-1,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{0, 1},
			},
			0,
			MatrixCoordinates{0, 0},
			true,
		},
		{
			[][]int{
				{0, 1},
			},
			1,
			MatrixCoordinates{0, 1},
			true,
		},
		{
			[][]int{
				{0, 1},
			},
			2,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{0},
				{1},
			},
			0,
			MatrixCoordinates{0, 0},
			true,
		},
		{
			[][]int{
				{0},
				{1},
			},
			1,
			MatrixCoordinates{1, 0},
			true,
		},
		{
			[][]int{
				{0},
				{1},
			},
			2,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{0, 1},
				{2, 3},
			},
			2,
			MatrixCoordinates{1, 0},
			true,
		},
		{
			[][]int{
				{0, 1},
				{2, 3},
			},
			3,
			MatrixCoordinates{1, 1},
			true,
		},
		{
			[][]int{
				{0, 1},
				{2, 3},
			},
			4,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{0, 2},
				{1, 3},
			},
			1,
			MatrixCoordinates{Row: 1, Col: 0},
			true,
		},
		{
			[][]int{
				{0, 2},
				{1, 3},
			},
			2,
			MatrixCoordinates{Row: 0, Col: 1},
			true,
		},
		{
			[][]int{
				{0, 2},
				{1, 3},
			},
			3,
			MatrixCoordinates{Row: 1, Col: 1},
			true,
		},
		{
			[][]int{
				{0, 3},
				{1, 3},
			},
			2,
			MatrixCoordinates{},
			false,
		},
		{
			[][]int{
				{0, 3},
				{1, 2},
			},
			2,
			MatrixCoordinates{Row: 1, Col: 1},
			true,
		},
		{
			[][]int{
				{0, 1},
				{2, 3},
			},
			2,
			MatrixCoordinates{Row: 1, Col: 0},
			true,
		},
	}

	for i, test := range tests {
		coordinates, ok := SearchMatrix(test.matrix, test.element)
		assert.Equal(t, test.expectedCoordinates, coordinates, i)
		assert.Equal(t, test.expectedOk, ok, i)
	}
}
