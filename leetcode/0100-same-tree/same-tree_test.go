package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_isSameTree(t *testing.T) {
	t1 := types.Ints2TreeNode([]int{1, 2, 3})
	t2 := types.Ints2TreeNode([]int{1, 2, 3})
	got := isSameTree(t1, t2)
	require.True(t, got)

	t1 = types.Ints2TreeNode([]int{1, 2})
	t2 = types.Ints2TreeNode([]int{1, types.NULL, 3})
	got = isSameTree(t1, t2)
	require.False(t, got)

	t1 = types.Ints2TreeNode([]int{1, 2, 1})
	t2 = types.Ints2TreeNode([]int{1, 1, 2})
	got = isSameTree(t1, t2)
	require.False(t, got)
}
