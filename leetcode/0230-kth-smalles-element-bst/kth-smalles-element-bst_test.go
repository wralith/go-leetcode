package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_kthSmallest(t *testing.T) {
	tree := types.Ints2TreeNode([]int{3, 1, 4, types.NULL, 2})
	got := kthSmallest(tree, 1)
	require.Equal(t, 1, got)

	got = kthSmallest2(tree, 1)
	require.Equal(t, 1, got)

	tree = types.Ints2TreeNode([]int{5, 3, 6, 2, 4, types.NULL, types.NULL, 1})
	got = kthSmallest(tree, 3)
	require.Equal(t, 3, got)

	got = kthSmallest2(tree, 3)
	require.Equal(t, 3, got)
}
