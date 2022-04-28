package cracking_the_coding_interview

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// ConversionFlips returns the number of bit flips needed to convert a to b.
func ConversionFlips(a int32, b int32) int {
	flips := a ^ b
	count := 0
	for i := 0; i < 32; i++ {
		if flips&(1<<i) != 0 {
			count++
		}
	}
	return count
}

func TestConversionFlips(t *testing.T) {
	tests := []struct {
		a        int32
		b        int32
		expected int
	}{
		{
			0,
			0,
			0,
		},
		{
			0,
			1,
			1,
		},
		{
			0b00,
			0b10,
			1,
		},
		{
			0b00,
			0b11,
			2,
		},
		{
			0b11101,
			0b01111,
			2,
		},
	}

	for i, test := range tests {
		result := ConversionFlips(test.a, test.b)
		assert.Equal(t, test.expected, result, i)
	}
}
