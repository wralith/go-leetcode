package leetcode

import (
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/wralith/go-leetcode/types"
)

func Test_maxDepth(t *testing.T) {
	tree := types.Ints2TreeNode([]int{3, 9, 20, types.NULL, types.NULL, 15, 7})
	got := maxDepth(tree)
	require.Equal(t, 3, got)

	tree = types.Ints2TreeNode([]int{1, types.NULL, 2})
	got = maxDepth(tree)
	require.Equal(t, 2, got)
}
