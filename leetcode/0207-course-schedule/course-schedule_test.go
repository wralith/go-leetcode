package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_canFinish(t *testing.T) {
	got := canFinish(2, [][]int{{1, 0}})
	require.True(t, got)

	got = canFinish(2, [][]int{{1, 0}, {0, 1}})
	require.False(t, got)
}
