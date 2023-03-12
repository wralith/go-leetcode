package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func TestRightSideView(t *testing.T) {
	root := types.Ints2TreeNode([]int{1, 2, 3, types.NULL, 5, types.NULL, 4})
	got := rightSideView(root)
	require.Equal(t, []int{1, 3, 4}, got)

	root = types.Ints2TreeNode([]int{1, types.NULL, 3})
	got = rightSideView(root)
	require.Equal(t, []int{1, 3}, got)

	root = types.Ints2TreeNode([]int{1, 2})
	got = rightSideView(root)
	require.Equal(t, []int{1, 2}, got)
}
