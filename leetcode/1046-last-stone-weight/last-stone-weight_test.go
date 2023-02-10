package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_lastStoneWeight(t *testing.T) {
	got := lastStoneWeight([]int{2, 7, 4, 1, 8, 1})
	require.Equal(t, 1, got)

	got = lastStoneWeight([]int{1})
	require.Equal(t, 1, got)
}
