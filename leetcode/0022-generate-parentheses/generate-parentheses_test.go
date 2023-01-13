package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_generateParenthesis(t *testing.T) {
	want := []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	got := generateParenthesis(3)
	require.Equal(t, want, got)

	want = []string{"()"}
	got = generateParenthesis(1)
	require.Equal(t, want, got)

	want = []string{"((()))", "(()())", "(())()", "()(())", "()()()"}
	got = generateParenthesis(3)
	require.Equal(t, want, got)
}
