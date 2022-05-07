package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type ColorMatrix struct {
	Arr    []uint8
	Width  int
	Height int
}

type Coordinate struct {
	X int
	Y int
}

func CoordinateToIndex(coordinate Coordinate, width int) int {
	return coordinate.Y*width + coordinate.X
}

func NewColorMatrixFromSlice(colors [][]uint8) *ColorMatrix {
	if len(colors) == 0 || len(colors[0]) == 0 {
		return &ColorMatrix{}
	}

	arr := make([]uint8, len(colors)*len(colors[0]))
	index := 0

	for x := 0; x < len(colors); x++ {
		if len(colors[x]) != len(colors[0]) {
			panic("invalid array length")
		}

		for y := 0; y < len(colors[0]); y++ {
			arr[index] = colors[x][y]
			index++
		}
	}

	return &ColorMatrix{
		Arr:    arr,
		Width:  len(colors),
		Height: len(colors[0]),
	}
}

// Fill fills the area with the same color containing the given coordinate with the given
// color.
func (c *ColorMatrix) Fill(color uint8, coordinate Coordinate) {

}

func (c *ColorMatrix) ToSlice() [][]uint8 {
	arr := make([][]uint8, c.Width)

	for i := range arr {
		arr[i] = make([]uint8, c.Height)
		copy(arr[i], c.Arr[i*c.Height:(i+1)*c.Height])
	}

	return arr
}

func TestNewColorMatrixFromSlice(t *testing.T) {
	assert.Panics(t, func() {
		NewColorMatrixFromSlice([][]uint8{{0}, {1, 2}})
	})

	tests := []struct {
		colors [][]uint8
	}{
		{
			nil,
		},
		{
			[][]uint8{
				{0},
			},
		},
		{
			[][]uint8{
				{0, 1},
			},
		},
		{
			[][]uint8{
				{0},
				{1},
			},
		},
		{
			[][]uint8{
				{0, 1},
				{2, 3},
			},
		},
	}

	for i, test := range tests {
		matrix := NewColorMatrixFromSlice(test.colors)
		result := matrix.ToSlice()
		assert.ElementsMatch(t, test.colors, result, i)
	}
}
