package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_searchMatrix(t *testing.T) {
	got := searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 3)
	require.True(t, got)

	got = searchMatrix([][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}, 13)
	require.False(t, got)
}
