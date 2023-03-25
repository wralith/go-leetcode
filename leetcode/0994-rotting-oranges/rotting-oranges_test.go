package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrangesRotting(t *testing.T) {
	got := orangesRotting([][]int{
		{2, 1, 1},
		{1, 1, 0},
		{0, 1, 1},
	})
	require.Equal(t, 4, got)

	got = orangesRotting([][]int{
		{2, 1, 1},
		{0, 1, 1},
		{1, 0, 1},
	})
	require.Equal(t, -1, got)

	got = orangesRotting([][]int{
		{0, 2},
	})
	require.Equal(t, 0, got)

	got = orangesRotting([][]int{
		{1},
		{2},
		{2},
	})
	require.Equal(t, 1, got)
}
