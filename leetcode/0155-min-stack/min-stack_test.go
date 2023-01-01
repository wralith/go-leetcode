package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_MinStack(t *testing.T) {
	minStack := MinStackConstructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	got := minStack.GetMin() // return -3
	require.Equal(t, -3, got)
	minStack.Pop()
	got = minStack.Top() // return 0
	require.Equal(t, 0, got)
	got = minStack.GetMin() // return -2
	require.Equal(t, -2, got)
}
