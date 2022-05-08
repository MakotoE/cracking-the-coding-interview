package cracking_the_coding_interview

import (
	"errors"
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

func coordinateToIndex(coordinate Coordinate, width int) int {
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
func (c *ColorMatrix) Fill(coordinate Coordinate, newColor uint8) error {
	if coordinate.X < 0 || coordinate.X >= c.Width || coordinate.Y < 0 || coordinate.Y >= c.Height {
		return errors.New("invalid coordinate")
	}

	startingIndex := coordinateToIndex(coordinate, c.Width)
	color := c.Arr[startingIndex]
	c.Arr[startingIndex] = newColor

	stack := []Coordinate{coordinate}

	for len(stack) > 0 {
		currentCoordinate := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if currentCoordinate.X > 0 {
			left := currentCoordinate
			left.X -= 1

			index := coordinateToIndex(left, c.Width)
			if c.Arr[index] == color {
				c.Arr[index] = newColor
				stack = append(stack, left)
			}
		}
		if currentCoordinate.X < c.Width-1 {
			right := currentCoordinate
			right.X += 1

			index := coordinateToIndex(right, c.Width)
			if c.Arr[index] == color {
				c.Arr[index] = newColor
				stack = append(stack, right)
			}
		}
		if currentCoordinate.Y > 0 {
			up := currentCoordinate
			up.Y -= 1

			index := coordinateToIndex(up, c.Width)
			if c.Arr[index] == color {
				c.Arr[index] = newColor
				stack = append(stack, up)
			}
		}
		if currentCoordinate.Y < c.Height-1 {
			down := currentCoordinate
			down.Y += 1

			index := coordinateToIndex(down, c.Width)
			if c.Arr[index] == color {
				c.Arr[index] = newColor
				stack = append(stack, down)
			}
		}
	}
	return nil
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

func TestColorMatrix_Fill(t *testing.T) {
	assert.Error(t, NewColorMatrixFromSlice(nil).Fill(Coordinate{}, 0))

	tests := []struct {
		colors     [][]uint8
		coordinate Coordinate
		newColor   uint8
		expected   [][]uint8
	}{
		{
			[][]uint8{
				{0},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1},
			},
		},
		{
			[][]uint8{
				{0, 0},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1, 1},
			},
		},
		{
			[][]uint8{
				{0, 0},
			},
			Coordinate{0, 1},
			1,
			[][]uint8{
				{1, 1},
			},
		},
		{
			[][]uint8{
				{0, 2},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1, 2},
			},
		},
		{
			[][]uint8{
				{0},
				{0},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1},
				{1},
			},
		},
		{
			[][]uint8{
				{0},
				{0},
			},
			Coordinate{1, 0},
			1,
			[][]uint8{
				{1},
				{1},
			},
		},
		{
			[][]uint8{
				{0, 2, 0},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1, 2, 0},
			},
		},
		{
			[][]uint8{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			},
			Coordinate{1, 1},
			1,
			[][]uint8{
				{1, 1, 1},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			[][]uint8{
				{2, 2, 2},
				{0, 0, 0},
				{0, 0, 0},
			},
			Coordinate{1, 1},
			1,
			[][]uint8{
				{2, 2, 2},
				{1, 1, 1},
				{1, 1, 1},
			},
		},
		{
			[][]uint8{
				{0, 0, 2},
				{0, 0, 2},
				{2, 2, 2},
			},
			Coordinate{1, 1},
			1,
			[][]uint8{
				{1, 1, 2},
				{1, 1, 2},
				{2, 2, 2},
			},
		},
		{
			[][]uint8{
				{0, 0, 0},
				{2, 2, 2},
				{0, 0, 0},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1, 1, 1},
				{2, 2, 2},
				{0, 0, 0},
			},
		},
		{
			[][]uint8{
				{0, 2, 0},
				{2, 2, 0},
				{0, 0, 0},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1, 2, 0},
				{2, 2, 0},
				{0, 0, 0},
			},
		},
		{
			[][]uint8{
				{0, 0, 0},
				{2, 2, 0},
				{0, 0, 0},
			},
			Coordinate{0, 0},
			1,
			[][]uint8{
				{1, 1, 1},
				{2, 2, 1},
				{1, 1, 1},
			},
		},
	}

	for i, test := range tests {
		matrix := NewColorMatrixFromSlice(test.colors)
		err := matrix.Fill(test.coordinate, test.newColor)
		assert.NoError(t, err, i)

		result := matrix.ToSlice()
		assert.Equal(t, test.expected, result, i)
	}
}
