package cracking_the_coding_interview

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Returns all possible combinations of parentheses.
func Parens(n int) map[string]bool {
	if n == 0 {
		return map[string]bool{}
	}

	return parentheses(n, 1)
}

func parentheses(n int, i int) map[string]bool {
	if n == i {
		return map[string]bool{"()": true}
	}

	result := make(map[string]bool)

	for s := range parentheses(n, i+1) {
		result[fmt.Sprintf("%s()", s)] = true
		result[fmt.Sprintf("(%s)", s)] = true
		result[fmt.Sprintf("()%s", s)] = true
	}

	return result
}

func TestParens(t *testing.T) {
	tests := []struct {
		n               int
		expectedStrings []string
	}{
		{
			0,
			nil,
		},
		{
			1,
			[]string{"()"},
		},
		{
			2,
			[]string{"()()", "(())", "()()"},
		},
		{
			3,
			[]string{"()()()", "(())()", "(()())", "((()))", "()(())"},
		},
		{
			4,
			[]string{
				"(((())))",
				"((()()))",
				"((())())",
				"((()))()",
				"(()(()))",
				"(()()())",
				"(()())()",
				"(())()()",
				"()((()))",
				"()(()())",
				"()(())()",
				"()()(())",
				"()()()()",
			},
		},
	}

	for i, test := range tests {
		result := Parens(test.n)
		expected := map[string]bool{}
		for _, s := range test.expectedStrings {
			expected[s] = true
		}
		assert.Equal(t, expected, result, i)
	}
}
