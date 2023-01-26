package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_isSubtree(t *testing.T) {
	tree := types.Ints2TreeNode([]int{3, 4, 5, 1, 2})
	subTree := types.Ints2TreeNode([]int{4, 1, 2})

	got := isSubtree(tree, subTree)
	require.True(t, got)

	tree = types.Ints2TreeNode([]int{3, 4, 5, 1, 2, types.NULL, types.NULL, types.NULL, types.NULL, 0})
	subTree = types.Ints2TreeNode([]int{4, 1, 2})

	got = isSubtree(tree, subTree)
	require.False(t, got)
}
